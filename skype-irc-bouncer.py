#!/usr/bin/python

# TODOs
# Add support for:
#  - WHOIS/WHO/WHOWAS
#  - AWAY
#  - creating chats
#  - TOPIC (setting)

import re
import ssl
import socket
import select
import logging
import datetime

from Skype4Py import Skype
import Skype4Py
import six
from irc import events
from irc import client
from irc import buffer

MOD_HANDLE = "[MOD]"
UNREAD_HANDLE = "UNREAD"
SRV_WELCOME = "Welcome to %s v%s." % (__name__, client.VERSION)


class IRCError(Exception):
    """
    Exception thrown by IRC command handlers to notify
     client of a server/client error.
    """
    def __init__(self, code, value):
        self.code = code
        self.value = value

    def __str__(self):
        return repr(self.value)

    @classmethod
    def from_name(cls, name, value):
        return cls(events.codes[name], value)


class IRCClient(six.moves.socketserver.BaseRequestHandler):
    """
    IRC client connect and command handling. Client connection is handled by
    the `handle` method which sets up a two-way communication with the client.
    It then handles commands sent by the client by dispatching them to the
    handle_ methods.
    """
    class Disconnect(BaseException):
        pass

    def __init__(self, request, client_address, server):
        self.user = None
        self.host = client_address  # Client's hostname / ip.
        self.realname = None        # Client's real name
        self.nick = None            # Client's currently registered nickname
        self.send_queue = []        # Messages to send to client (strings)

        self._joined_chats = set()  # Skype chats that the client has interacted with
        
        self._chat_channelnames = []  # list(chat.Name, irc_channel_name)

        self._logger = logging.getLogger("IRCClient")

        six.moves.socketserver.BaseRequestHandler.__init__(self, request,
            client_address, server)

    def handle(self):
        self._logger.info('Client connected: %s', self.client_ident())
        self.buffer = buffer.LineBuffer()

        try:
            while True:
                self._handle_one()
        except self.Disconnect:
            self.request.close()

    def _handle_one(self):
        """
        Handle one read/write cycle.
        """
        ready_to_read, ready_to_write, in_error = select.select(
            [self.request], [self.request], [self.request], 0.1)

        if in_error:
            raise self.Disconnect()

        # Write any commands to the client
        while self.send_queue and ready_to_write:
            msg = self.send_queue.pop(0)
            self._send(msg)

        # See if the client has any commands for us.
        if ready_to_read:
            self._handle_incoming()

    def _handle_incoming(self):
        try:
            data = self.request.recv(1024)
        except Exception:
            raise self.Disconnect()

        if not data:
            raise self.Disconnect()

        self.buffer.feed(data)
        for line in self.buffer:
            line = line.decode('utf-8')
            self._handle_line(line)

    def _handle_line(self, line):
        try:
            self._logger.debug('from %s: %s' % (self.client_ident(), line))
            command, sep, params = line.partition(' ')
            handler = getattr(self, 'handle_%s' % command.lower(), None)
            if not handler:
                self._logger.info('No handler for command: %s. '
                    'Full line: %s' % (command, line))
                raise IRCError.from_name('unknowncommand',
                    '%s :Unknown command' % command)
            response = handler(params)
        except AttributeError as e:
            self._logger.error(six.text_type(e))
            raise
        except IRCError as e:
            response = ':%s %s %s' % (self.server.servername, e.code, e.value)
            self._logger.error(response)
        except Exception as e:
            response = ':%s ERROR %r' % (self.server.servername, e)
            self._logger.error(response)
            raise

        if response:
            self._send(response)

    def _send(self, msg):
        self._logger.debug('to %s: %s', self.client_ident(), msg)
        self.request.send(msg.encode('utf-8') + b'\r\n')

    def handle_nick(self, params):
        """
        Handle the initial setting of the user's nickname and nick changes.
        """
        nick = params

        # Valid nickname?
        if re.search('[^a-zA-Z0-9\-\[\]\'`^{}_]', nick):
            raise IRCError.from_name('erroneusnickname', ':%s' % nick)

        if self.server.clients.get(nick, None) == self:
            # Already registered to user
            return

        if nick in self.server.clients:
            # Someone else is using the nick
            raise IRCError.from_name('nicknameinuse', 'NICK :%s' % (nick))

        if not self.nick:
            # New connection and nick is available; register and send welcome
            # and MOTD.
            self.nick = nick
            self.server.clients[nick] = self
            response = ':%s %s %s :%s' % (self.server.servername,
                events.codes['welcome'], self.nick, SRV_WELCOME)
            self.send_queue.append(response)
            response = ':%s 376 %s :End of MOTD command.' % (
                self.server.servername, self.nick)
            self.send_queue.append(response)
            return

        # Nick is available. Change the nick.
        message = ':%s NICK :%s' % (self.client_ident(), nick)

        self.server.clients.pop(self.nick)
        self.nick = nick
        self.server.clients[self.nick] = self

        # Send a notification of the nick change to the client itself
        return message

    def handle_user(self, params):
        """
        Handle the USER command which identifies the user to the server.
        """
        params = params.split(' ', 3)

        if len(params) != 4:
            raise IRCError.from_name('needmoreparams',
                'USER :Not enough parameters')

        user, mode, unused, realname = params
        self.user = user
        self.mode = mode
        self.realname = realname
        return ''

    def handle_ping(self, params):
        """
        Handle client PING requests to keep the connection alive.
        """
        response = ':%s PONG :%s' % (self.server.servername, self.server.servername)
        return response

    def handle_join(self, params):
        # TODO: this is a disaster!
        channel_names = params.split(' ', 1)[0]  # Ignore keys
        for channel_name in channel_names.split(','):
            r_channel_name = channel_name.strip()

            self._logger.info("Request to JOIN %s" % (r_channel_name))

            found = False
            for chat in self.server.skype.Chats:
                if chat.Name == r_channel_name:
                    found = True
                    self._ensure_client_joined(chat)
                    # TODO: i don't like the early returns, but need a refactor first
                    return
            if not found:
                stripped_channel_name = r_channel_name.lstrip("#")
                matches = []
                for chat in self.server.skype.Chats:
                    if stripped_channel_name in chat.FriendlyName:
                        if stripped_channel_name == chat.FriendlyName:
                            self._ensure_client_joined(chat)
                            return
                        else:
                            matches.append(chat)
                        
                if len(matches) == 1:
                    self._ensure_client_joined(matches[0])
                    return
                elif len(matches) > 1:
                    message = "%s :Ambiguous channel name matches [%s]" % \
                        (stripped_channel_name, ", ".join(map(lambda c: c.FriendlyName, matches)))
                    raise IRCError.from_name('nosuchchannel', message)
                else:
                    self.build_all_friendly_channelnames()
                    matches = []
                    for skypename, channelname in self._chat_channelnames:
                        print stripped_channel_name, channelname
                        if stripped_channel_name in channelname:
                            if stripped_channel_name == channelname:
                                # perfect match
                                
                                # TODO: refactor this out
                                found = False
                                for chat in self.server.skype.Chats:
                                    if chat.Name == skypename:
                                        self._ensure_client_joined(chat)
                                        return
                                if not found:
                                    raise IRCError.from_name('nosuchchannel',
                                                         '%s :No such channel (0) ' % stripped_channel_name)                                
                            else:
                                matches.append((skypename, channelname))
                    if len(matches) == 1:
                        found = False
                        for chat in self.server.skype.Chats:
                            if chat.Name == matches[0][0]:
                                self._ensure_client_joined(chat)
                                return
                        if not found:
                            raise IRCError.from_name('nosuchchannel',
                                                 '%s :No such channel (0) ' % stripped_channel_name)
                    elif len(matches) == 0:
                        raise IRCError.from_name('nosuchchannel',
                                                 '%s :No such channel (1) ' % stripped_channel_name)
                    else:
                        message = "%s :Ambiguous channel name matches [%s]" % \
                            (stripped_channel_name, ", ".join(map(lambda p: p[1], matches)))
                        raise IRCError.from_name('nosuchchannel', message)   

    def _mark_all_chat_messages_read(self, chat):
        for message in chat.Messages:
            if message.Status == "RECEIVED":
                message.MarkAsSeen()

    def handle_privmsg(self, params):
        # TODO(wb): send privmsg to all other connected clients

        target, sep, msg = params.partition(' ')
        if not msg:
            raise IRCError.from_name('needmoreparams',
                'PRIVMSG :Not enough parameters')
        msg = msg.lstrip(":")

        if target.startswith('#') or target.startswith('$'):
            self._logger.info("message to %s: %s" % (target, msg))

            the_chat = None
            for chat in self.server.skype.Chats:
                if chat.Name == target:
                    the_chat = chat

            if the_chat is None:
                names = self.get_skypename_from_friendly_channelname(target)
                if len(names) == 0:
                    raise IRCError.from_name('cannotsendtochan',
                        '%s :Cannot send to channel' % target)
                elif len(names) > 0:
                    raise IRCError.from_name('cannotsendtochan',
                        '%s :Cannot send to channel %s, ambiguous: [%s]' % (target, ", ".join(names)))
                else:
                    the_chat = names

            the_chat.SendMessage(msg)
            self._mark_all_chat_messages_read(the_chat)
        else:
            self._logger.info("user message [UNSUPPORTED!] to %s: %s" % (target, msg))

    def _queue_client_message(self, id_, message):
        composed = ":%s %d %s %s" % (self.server.servername, id_, self.client_ident(), message)
        self._logger.info("SENDING: %s" % (composed))
        self.send_queue.append(composed)

    def _queue_client_mod_message(self, chat, message):
        composed = ":%s NOTICE %s :%s" % (MOD_HANDLE, chat.Name, message)
        self._logger.info("SENDING: %s" % (composed))
        self.send_queue.append(composed)
            
    def get_friendly_channelname_from_skypename(self, chat):
        # first, try to find a previously generated name
        for skypename, channelname in self._chat_channelnames:
            if chat.Name == skypename:
                return channelname
                
        # next, try to create a nice name
        channelname = None
        members = [m for m in chat.Members]
        if len(members) == 2:
            if members[0] == self.server.skype.User():
                channelname = members[1].Handle + "-priv"
            else:
                channelname = members[0].Handle + "-priv"
        else:
            is_generic_name = False
            for m in members:
                if m.FullName in chat.FriendlyName:
                    is_generic_name = True
                    break
            
            if not is_generic_name:
                channelname = chat.FriendlyName.replace(" ", "-")[:16]
            elif chat.Topic != "":
                channelname = chat.Topic.replace(" ", "-")[:16]
            else:
                channelname = chat.Name
        while True:
            if len(self._chat_channelnames) == 0:
                self._chat_channelnames.append((chat.Name, channelname))
                return channelname 
            has_conflict = False
            for skypename, existing_channelname in self._chat_channelnames:
                if existing_channelname == channelname and skypename == chat.Name:
                    # race with other channel, take the other one
                    return channelname
                elif existing_channelname == channelname:
                    # name conflict, change ours and retry
                    channelname += "Q"
                    has_conflict = True
                    break
            
            if not has_conflict:
                self._chat_channelnames.append((chat.Name, channelname))
                return channelname       
    
    def build_all_friendly_channelnames(self):
        for chat in self.server.skype.Chats:
            self.get_friendly_channelname_from_skypename(chat)
        
    def get_skypename_from_friendly_channelname(self, friendlychannelname):
        skypenames = []
        for skypename, channelname in self._chat_channelnames:
            if channelname == friendlychannelname:
                skypenames.append(skypename)
                
        return skypenames
        
    def handle_list(self, params):
        self._queue_client_message(321, "NAME :FRIENDLYNAME TIMESTAMP")
        
        sortable_chats = []
        unsortable_chats = []
        
        for c in self.server.skype.Chats:
            try:
                ts = c.Messagse[0].Timestamp
                sortable_chats.append(c)
            except Exception:
                self._logger.warn("Unable to get TS from chat")
                unsortable_chats.append(c)
        
        for chat in sorted(sortable_chats, key=lambda c: c.Messages[0].Timestamp):
            ts = datetime.datetime.fromtimestamp(chat.Messages[0].Timestamp)
            #self._queue_client_message(322, "%s :%s %s" % (chat.Name, chat.FriendlyName, ts))
            self._queue_client_message(322, "%s : | %s |  %s" % (self.get_friendly_channelname_from_skypename(chat), chat.FriendlyName, ts))
        for chat in unsortable_chats:
            #self._queue_client_message(322, "%s :%s %s" % (chat.Name, chat.FriendlyName, "UNKNOWN"))
            self._queue_client_message(322, "%s : | %s | %s" % (self.get_friendly_channelname_from_skypename(chat), chat.FriendlyName, "UNKNOWN"))
        self._queue_client_message(323, ":End of /LIST")
        print self._chat_channelnames
        
    def handle_mode(self, params):
        pass

    def client_ident(self):
        """
        Return the client identifier as included in many command replies.
        """
        return client.NickMask.from_params(self.nick, self.user,
            self.server.servername)

    def finish(self):
        """
        The client conection is finished. Do some cleanup to ensure that the
        client doesn't linger around in any channel or the client list, in case
        the client didn't properly close the connection with PART and QUIT.
        """
        self._logger.info('Client disconnected: %s', self.client_ident())
        if self.nick is not None:
            self.server.clients.pop(self.nick)
        else:
            print self.server.clients
        self._logger.info('Connection finished: %s', self.client_ident())

    def __repr__(self):
        """
        Return a user-readable description of the client
        """
        return '<%s %s!%s@%s (%s)>' % (
            self.__class__.__name__,
            self.nick,
            self.user,
            self.host[0],
            self.realname,
            )

    def _ensure_client_joined(self, chat):
        if chat.Name not in self._joined_chats:
            self._logger.info("joining to %s" % (chat.Name))
            message = ":%s JOIN :%s" % (self.client_ident(), chat.Name)
            self.send_queue.append(message)

            message = ":%s TOPIC %s :%s" % (MOD_HANDLE, chat.Name, chat.FriendlyName)
            self.send_queue.append(message)

            nicks = [m.Handle for m in chat.Members]
            self._queue_client_message(353, "= %s :%s" % (chat.Name, " ".join(nicks)))
            self._queue_client_message(366, "%s :End of /NAMES list" % (chat.Name))

            unread_messages = sorted([m for m in chat.Messages if m.Status == "RECEIVED"], key=lambda m: m.Timestamp)
            if len(unread_messages) > 0:
                self._queue_client_mod_message(chat, "Begin of UNREAD messages")            
                for message in unread_messages:
                    self._logger.info("unread message: %s %s %s" % (message.Timestamp, message.FromHandle, message.Body))
                    ts = datetime.datetime.fromtimestamp(message.Timestamp)
                    # TODO(wb): handle long messages
                    # TODO(wb): handle messages with newlines
                    for part in message.Body.split("\n"):
                        if part.rstrip("\r\n ") == "":
                            continue
                        m = "<%s %s> %s" % (ts.isoformat(), message.FromHandle, part)
                        composed = ":%s NOTICE %s :%s" % (UNREAD_HANDLE, chat.Name, m)
                        self.send_queue.append(composed)
                self._queue_client_mod_message(chat, "End of UNREAD messages")

            self._joined_chats.add(chat.Name)

    def handle_skype_incoming(self, message, status): #chat_name, user_handle, message):
        self._ensure_client_joined(message.Chat)

        # TODO(wb): handle long messages
        # TODO(wb): handle messages with newlines
        self._logger.info("incoming message: %s %s %s" % (message.Chat.Name, message.FromHandle, message.Body))
        for part in message.Body.split("\n"):
            if part.rstrip("\r\n ") == "":
                continue
            m = ":%s PRIVMSG %s :%s" % (message.FromHandle, message.Chat.Name, part)
            self.send_queue.append(m)



class IRCServer(six.moves.socketserver.ThreadingMixIn,
                six.moves.socketserver.TCPServer):
    daemon_threads = True
    allow_reuse_address = True

    clients = {}
    "Connected clients (IRCClient instances) by nick name"

    def __init__(self, *args, **kwargs):
        self._logger = logging.getLogger("IRCServer")

        self.servername = '0.0.0.0'
        self.clients = {}

        self.skype = Skype()
        self.skype.Attach()
        self.skype.OnUserAuthorizationRequestReceived = self._handle_auth_req
        self.skype.OnMessageStatus = self._handle_recv_message

        six.moves.socketserver.TCPServer.__init__(self, *args, **kwargs)
        self.socket = ssl.wrap_socket(self.socket, 
                                      keyfile="./ssl/server.key",
                                      certfile="./ssl/serverca.pem", 
                                      server_side=True, 
                                      cert_reqs=ssl.CERT_REQUIRED, 
                                      ca_certs="./ssl/clientca.pem", 
                                      do_handshake_on_connect=True)


    def _handle_auth_req(self, user):
        self._logger.info("Received authorization request: %s" % (user))

    def _handle_recv_message(self, message, status):
        if status == Skype4Py.cmsReceived:
            for client in self.clients.values():
                client.handle_skype_incoming(message, status)


def main():
    logger = logging.getLogger("main")
    logging.basicConfig(format='%(asctime)s %(message)s', level=logging.INFO)

    address = "0.0.0.0"
    port = 6667

    try:
        ircserver = IRCServer((address, port), IRCClient)
        logger.info("Server started on %s:%d" % (address, port))
        ircserver.serve_forever()
    except socket.error as e:
        logger.error(repr(e))
        raise SystemExit(-2)

if __name__ == "__main__":
    import sys
    if len(sys.argv) > 1 and sys.argv[1] == "profile":
        import cProfile
        cProfile.run("main()")
    else:
        main()
