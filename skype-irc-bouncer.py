#!/usr/bin/python

# TODOs
# Add support for:
#  - WHOIS/WHO/WHOWAS
#  - AWAY
#  - creating chats
#  - TOPIC (setting)
import sys
import re
import ssl
import time
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
READ_HANDLE = "READ"
SRV_WELCOME = "Welcome to %s v%s." % (__name__, client.VERSION)


def camelcase_sentence(value):
    """
    inspired by http://stackoverflow.com/a/4306777/87207
    """
    def camelcase():
        yield str.lower
        while True:
            yield str.capitalize

    def camelcase_u():
        yield unicode.lower
        while True:
            yield unicode.capitalize

    if isinstance(value, unicode):
        c = camelcase_u()
    else:
        c = camelcase()
    return "".join(c.next()(x) if x else ' ' for x in value.split(' '))


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


def ensure_client_joined_unread_chats_operation(client):
        for chat in client.get_unread_chats():
            client._ensure_joined_chat(chat)


class IRCClientModule(object):
    def __init__(self, client):
        super(IRCClientModule, self).__init__()
        self._client = client

    def get_irc_handlers(self):
        """
        Return sequence of ("irc-command", function)
        """
        return []

    def get_bang_handlers(self):
        """
        Return sequence of ("bang-command", function)
        """
        return []


class HandleUnreadModule(IRCClientModule):
    def __init__(self, client):
        super(HandleUnreadModule, self).__init__(client)

    def get_irc_handlers(self):
        return [("unread", self._handle_unread)]

    def _handle_unread(self, params):
        self._client.queue_irc_message(321, "COUNT :NAME")
        for chat in self._client.get_unread_chats():
            self._client.queue_irc_message(322, "(%d)   [%s]" % (len([m for m in self._client.get_unread_messages(chat)]), self._client.get_friendly_channelname_from_chat(chat)))
        self._client.queue_irc_message(323, ":End of !unread")


class HandleClearUnreadModule(IRCClientModule):
    def __init__(self, client):
        super(HandleClearUnreadModule, self).__init__(client)

    def get_irc_handlers(self):
        return [("clear_unread", self._handle_clear_unread)]

    def _handle_clear_unread(self, params):
        self._client.queue_irc_message(321, "Beginning :to clear unread messages")
        for chat in self._client.get_unread_chats():
            for message in self._client.get_unread_messages(chat):
                message.MarkAsSeen()
            self._client.queue_irc_message(322, "  cleared: %s" % (self._client.get_friendly_channelname_from_chat(chat)))
        self._client.queue_irc_message(323, ":End of clear unread")


class HandleHistoryModule(IRCClientModule):
    def __init__(self, client):
        super(HandleHistoryModule, self).__init__(client)

    def get_bang_handlers(self):
        return [("history", self._handle_history)]

    def _handle_history(self, chat, message):
        parts = message.split(" ")
        friendlychannelname = self._client.get_friendly_channelname_from_chat(chat)
        num_messages = 20
        if len(parts) > 1:
            try:
                num_messages = int(parts[1])
            except ValueError:
                pass

        messages = sorted([m for m in chat.Messages], key=lambda m: m.Timestamp)
        if len(messages) > 0:
            self._client.queue_mod_message(chat, "Begin of HISTORY(%d)" % (num_messages))
            for message in messages[-num_messages:]:
                ts = datetime.datetime.fromtimestamp(message.Timestamp)
                # TODO(wb): handle long messages
                # TODO(wb): handle messages with newlines
                for part in message.Body.split("\n"):
                    if part.rstrip("\r\n ") == "":
                        continue
                    m = "<%s %s> %s" % (ts.isoformat(), message.FromHandle, part)
                    if message.Status == "RECEIVED":
                        handle = UNREAD_HANDLE
                    else:
                        handle = READ_HANDLE
                    composed = ":%s NOTICE %s :%s" % (handle, friendlychannelname, m)
                    self._client.queue_message(composed)
            self._client.queue_mod_message(chat, "End of HISTORY(%d)" % (num_messages))


class HandleListModule(IRCClientModule):
    def __init__(self, client):
        super(HandleListModule, self).__init__(client)

    def get_irc_handlers(self):
        return [("list", self._handle_list)]

    def _handle_list(self, params):
        self._client.queue_irc_message(321, "NAME :FRIENDLYNAME TIMESTAMP")

        sortable_chats = []
        unsortable_chats = []

        for chat in self._client.server.skype.Chats:
            self._client.queue_irc_message(322, "[%s] : | [%s]" % \
               (self._client.get_friendly_channelname_from_chat(chat), chat.FriendlyName))
        self._client.queue_irc_message(323, ":End of /LIST") 


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
        self._logger = logging.getLogger("IRCClient")
        self.user = None
        self.host = client_address   # Client's hostname / ip.
        self.realname = None         # Client's real name
        self.nick = None             # Client's currently registered nickname
        self.send_queue = []         # Messages to send to client (strings)
        self._joined_chats = set()   # set(chat.Name), Skype chats that the client has interacted with
        self._chat_channelnames = []  # list(chat.Name, irc_channel_name)
        self._supported_irc_commands = {
            "nick": self._handle_nick,
            "user": self._handle_user,
            "ping": self._handle_ping,
            "join": self._handle_join,
            "privmsg": self._handle_privmsg,
            "mode": self._handle_mode,
            }
        self._supported_bang_commands = { }

        self.install_client_module(HandleUnreadModule)
        self.install_client_module(HandleClearUnreadModule)
        self.install_client_module(HandleListModule)

        self.install_client_module(HandleHistoryModule)

        six.moves.socketserver.BaseRequestHandler.__init__(self, request,
                                                           client_address,
                                                           server)

    def install_client_module(self, module_cls):
        module = module_cls(self)
        for cmd, fn in module.get_irc_handlers():
            self._supported_irc_commands[cmd] = fn
        for cmd, fn in module.get_bang_handlers():
            self._supported_bang_commands[cmd] = fn

    def handle(self):
        self._logger.info('Client connected: %s', self.client_ident())
        self.buffer = buffer.LineBuffer()

        try:
            while True:
                if not self._handle_one():
                    time.sleep(0.1)
        except self.Disconnect:
            self.request.close()

    def _handle_one(self):
        """
        Handle one read/write cycle.
        """
        ready_to_read, ready_to_write, in_error = select.select(
            [self.request], [self.request], [self.request], 0.1)

        did_handle = False
        if in_error:
            raise self.Disconnect()

        # Write any commands to the client
        while self.send_queue and ready_to_write:
            msg = self.send_queue.pop(0)
            self._send(msg)
            did_handle = True

        # See if the client has any commands for us.
        if ready_to_read:
            self._handle_incoming()
            did_handle = True

        return did_handle

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

            try:
                handler = self._supported_irc_commands[command.lower()]
            except KeyError:
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

    def _ensure_joined_unread_chats(self):
        for chat in self.get_unread_chats():
            self._ensure_joined_chat(chat)

    def _handle_nick(self, params):
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
            self.queue_message(response)
            response = ':%s 376 %s :End of MOTD command.' % (
                self.server.servername, self.nick)
            self.queue_message(response)
            self._ensure_joined_unread_chats()
            return

        # Nick is available. Change the nick.
        message = ':%s NICK :%s' % (self.client_ident(), nick)

        self.server.clients.pop(self.nick)
        self.nick = nick
        self.server.clients[self.nick] = self

        # Send a notification of the nick change to the client itself
        return message

    def _handle_user(self, params):
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

    def _handle_ping(self, params):
        """
        Handle client PING requests to keep the connection alive.
        """
        response = ':%s PONG :%s' % (self.server.servername, self.server.servername)
        return response

    def get_unread_chats(self):
        ret = []
        for chat in self.server.skype.Chats:
            for message in self.get_unread_messages(chat):
                ret.append(chat)
                break
        return ret

    def get_unread_messages(self, chat):
        for message in chat.Messages:
            if message.Status == "RECEIVED":
                yield message

    def _ensure_joined_chat(self, chat):
        if chat.Name not in self._joined_chats:
            friendlychannelname = self.get_friendly_channelname_from_chat(chat)
            self._logger.info("joining to %s (%s)" % (chat.Name, friendlychannelname))
            message = ":%s JOIN :%s" % (self.client_ident(), friendlychannelname) 
            self.queue_message(message)

            message = ":%s TOPIC %s :%s" % (MOD_HANDLE, friendlychannelname, chat.FriendlyName)
            self.queue_message(message)

            nicks = [m.Handle for m in chat.Members]
            self.queue_irc_message(353, "= %s :%s" % (friendlychannelname, " ".join(nicks)))
            self.queue_irc_message(366, "%s :End of /NAMES list" % (friendlychannelname))

            unread_messages = sorted([m for m in self.get_unread_messages(chat)], key=lambda m: m.Timestamp)
            if len(unread_messages) > 0:
                self.queue_mod_message(chat, "Begin of UNREAD messages")
                for message in unread_messages:
                    ts = datetime.datetime.fromtimestamp(message.Timestamp)
                    # TODO(wb): handle long messages
                    # TODO(wb): handle messages with newlines
                    for part in message.Body.split("\n"):
                        if part.rstrip("\r\n ") == "":
                            continue
                        m = "<%s %s> %s" % (ts.isoformat(), message.FromHandle, part)
                        composed = ":%s NOTICE %s :%s" % (UNREAD_HANDLE, friendlychannelname, m)
                        self.queue_message(composed)
                self.queue_mod_message(chat, "End of UNREAD messages")
            self._joined_chats.add(chat.Name)

    def _handle_join(self, params):
        channel_names = params.split(' ', 1)[0]
        for channel_name in channel_names.split(','):
            channel_name = channel_name.strip()
            self._logger.info("Request to JOIN %s" % (channel_name))

            if channel_name.startswith('#') or channel_name.startswith('$'):
                chats = self._guess_chats_from_user_channelname(channel_name)
                if len(chats) == 0:
                    raise IRCError.from_name('nosuchchannel',
                                             '%s :Cannot join to channel (DNE)' % channel_name)
                elif len(chats) > 1:
                    chat_list = ", ".join(map(lambda c: c.Name, chats))
                    raise IRCError.from_name('nosuchchannel',
                                             '%s :Cannot join to channel %s, ambiguous: [%s]' % (channel_name, channel_name, chat_list))
                else:
                    the_chat = chats[0]
                self._ensure_joined_chat(the_chat)
            else:
                self._logger.info("user join [UNSUPPORTED!] to %s" % (channel_name))

    def _mark_all_chat_messages_read(self, chat):
        for message in chat.Messages:
            if message.Status == "RECEIVED":
                message.MarkAsSeen()

    def _handle_privmsg(self, params):
        target, sep, msg = params.partition(' ')
        if not msg:
            raise IRCError.from_name('needmoreparams', 'PRIVMSG :Not enough parameters')
        msg = msg.lstrip(":")

        if target.startswith('#') or target.startswith('$'):
            self._logger.info("message to %s: %s" % (target, msg))

            chats = self._guess_chats_from_user_channelname(target)
            if len(chats) == 0:
                raise IRCError.from_name('cannotsendtochan',
                                         '%s :Cannot send to channel (DNE)' % target)
            elif len(chats) > 1:
                chat_list = ", ".join(map(lambda c: c.Name, chats))
                raise IRCError.from_name('cannotsendtochan',
                                         '%s :Cannot send to channel %s, ambiguous: [%s]' % (target, chat_list))
            else:
                the_chat = chats[0]

            for nick, client in self.server.clients.items():
                if nick == self.nick:
                    continue
                # TODO: untested
                client.send_queue.append(params)

            if msg[0] == "!" and msg[1:].partition(" ")[0].lower() in self._supported_bang_commands:
                handler = self._supported_bang_commands[msg[1:].partition(" ")[0].lower()]
                self.queue_mod_message(the_chat, "Intercepted operator command: " + msg)
                handler(the_chat, msg)
            elif msg[0] == "!":
                self.queue_mod_message(the_chat, "Intercepted operator command: " + msg)
            else:
                the_chat.SendMessage(msg)
                self._mark_all_chat_messages_read(the_chat)
        else:
            self._logger.info("user message [UNSUPPORTED!] to %s: %s" % (target, msg))

    def queue_message(self, message):
        self._logger.info("SENDING: %s" % (message))
        self.send_queue.append(message)

    def queue_irc_message(self, id_, message):
        composed = ":%s %d %s %s" % (self.server.servername, id_, self.client_ident(), message)
        self.queue_message(composed)

    def queue_mod_message(self, chat, message):
        composed = ":%s NOTICE %s :%s" % (MOD_HANDLE, self.get_friendly_channelname_from_chat(chat), message)
        self.queue_message(composed)

    def get_friendly_channelname_from_chat(self, chat):
        # first, try to find a previously generated name
        for skypename, channelname in self._chat_channelnames:
            if chat.Name == skypename:
                return channelname

        # next, try to create a nice name
        channelname = None
        members = [m for m in chat.Members]
        if len(members) == 2:
            # this is a private chat
            other_user = members[0]
            if other_user == self.server.skype.User():
                other_user = members[1]
            if other_user.FullName.strip(" \t\n") != "":
                channelname = other_user.FullName
            else:
                channelname = other_user.Handle
            channelname += "-priv"
            channelname = camelcase_sentence(channelname)
            # example: u"WilliBallenthin-priv"
        else:
            # generic name is a Skype chat ID that simply contains
            #  the name of one of the participants and the first
            #  line of the chat:
            #  for example:
            #   u'Nick Pelletier | I think I found a good topic for MIRcon'
            is_generic_name = False
            for m in members:
                if m.FullName in chat.FriendlyName:
                    is_generic_name = True
                    break

            if is_generic_name:
                channelname = chat.FriendlyName.replace("|", "-")
                channelname = camelcase_sentence(channelname)
                # example: u'NickPelletier-IThinkIFoundAGoodTopicForMIRcon'
            elif chat.FriendlyName != "":
                channelname = camelcase_sentence(chat.FriendlyName)
                # example: u'OlympicCommittee(RIP09/17/2013)'
            elif chat.Topic != "":
                channelname = camelcase_sentence(chat.Topic)
                # example: u'PebbleWatchesWatch'
            else:
                channelname = chat.Name
                # example: '#williballenthin/$xolot1;bb755632147dedddfe3'

        if channelname[0] != "#":
            channelname = "#" + channelname

        channelname = channelname.replace(",", "-")

        if len(self._chat_channelnames) != 0:
            for skypename, existing_channelname in self._chat_channelnames:
                if existing_channelname == channelname and \
                        skypename == chat.Name:
                    # must be a race with another channel, take the other one
                    return channelname
                elif existing_channelname == channelname:
                    channelname += chat.Name[-16:]

        self._chat_channelnames.append((chat.Name, channelname))
        return channelname

    def _precompute_all_friendly_channelnames(self):
        for chat in self.server.skype.Chats:
            self.get_friendly_channelname_from_chat(chat)

    def _get_skypename_from_friendly_channelname(self, friendlychannelname):
        # TODO: consider caching here
        self._precompute_all_friendly_channelnames()

        for skypename, channelname in self._chat_channelnames:
            if channelname == friendlychannelname:
                return skypename
        raise KeyError("Unable to determine skypename from channelname: " + friendlychannelname)

    def _get_chat_by_channelname(self, channelname):
        """
        aka, by Skype.Name
        """
        # TODO: consider caching here
        for chat in self.server.skype.Chats:
            if chat.Name == channelname:
                return chat
        raise KeyError("Unable to find Chat from channelname: " + channelname)

    def _get_chat_from_friendly_channelname(self, friendlychannelname):
        """
        @raises KeyError: if the Chat cannot be found.
        """
        # TODO: consider caching here
        channelname = self._get_skypename_from_friendly_channelname(friendlychannelname)
        return self._get_chat_by_channelname(channelname)

    def _guess_skypenames_from_user_channelname(self, userchannelname):
        """
        Given a channel name provided by an IRC client, get a list of
          matching Skype channel names.
        """
        # TODO: consider caching here
        try:
            ret = [self._get_chat_from_friendly_channelname(userchannelname).Name]
            return ret
        except KeyError:
            pass

        matching_skypenames = []
        matching_friendlynames = []
        for skypename, channelname in self._chat_channelnames:
            if userchannelname == skypename:
                return skypename
            if userchannelname.lstrip("#") in skypename:
                matching_skypenames.append(skypename)
            if userchannelname == channelname:
                return skypename
            if userchannelname.lstrip("#") in channelname:
                matching_friendlynames.append(skypename)

        matching_skypenames.extend(matching_friendlynames)
        return matching_skypenames

    def _guess_chats_from_user_channelname(self, userchannelname):
        """
        Given a channel name provided by an IRC client, get a list of
          matching Skype Chats.
        """
        # TODO: consider caching here
        ret = []
        for channelname in self._guess_skypenames_from_user_channelname(userchannelname):
            try:
                ret.append(self._get_chat_by_channelname(channelname))
            except KeyError:
                pass
        return ret
    def _handle_mode(self, params):
        pass

    def client_ident(self):
        """
        Return the client identifier as included in many command replies.
        """
        return client.NickMask.from_params(self.nick, self.user, self.server.servername)

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

    def _handle_skype_incoming(self, message, status):
        self._ensure_joined_chat(message.Chat)

        # TODO(wb): handle long messages
        self._logger.info("incoming message: %s (%s) %s %s" % (message.Chat.Name,
                                                               self.get_friendly_channelname_from_chat(message.Chat),
                                                               message.FromHandle, message.Body))
        for part in message.Body.split("\n"):
            if part.rstrip("\r\n ") == "":
                continue
            m = ":%s PRIVMSG %s :%s" % (message.FromHandle, self.get_friendly_channelname_from_chat(message.Chat), part)
            self.queue_message(m)



class IRCServer(six.moves.socketserver.ThreadingMixIn,
                six.moves.socketserver.TCPServer):
    daemon_threads = True
    allow_reuse_address = True

    clients = {}
    "Connected clients (IRCClient instances) by nick name"

    def __init__(self, *args, **kwargs):
        self._logger = logging.getLogger("IRCServer")

        self.servername = '0.0.0.0'
        self.clients = {}  # map(nick --> client)

        self.skype = Skype()
        self.skype.Attach()
        self.skype.OnUserAuthorizationRequestReceived = self._handle_auth_req
        self.skype.OnMessageStatus = self._handle_recv_message

        disable_ssl = kwargs.pop("disable_ssl")
        six.moves.socketserver.TCPServer.__init__(self, *args, **kwargs)
        if not disable_ssl:
            self._logger.info("Enabled SSL")
            self.socket = ssl.wrap_socket(self.socket,
                                          keyfile="./ssl/server.key",
                                          certfile="./ssl/serverca.pem",
                                          server_side=True,
                                          cert_reqs=ssl.CERT_REQUIRED,
                                          ca_certs="./ssl/clientca.pem",
                                          do_handshake_on_connect=True)
        else:
            self._logger.info("Disabled SSL")

    def _handle_auth_req(self, user):
        self._logger.info("Received authorization request: %s" % (user))

    def _handle_recv_message(self, message, status):
        if status == Skype4Py.cmsReceived:
            for client in self.clients.values():
                client._handle_skype_incoming(message, status)


def main():
    logger = logging.getLogger("main")
    logging.basicConfig(format='%(asctime)s %(message)s', level=logging.INFO)

    address = "0.0.0.0"
    port = 6667

    try:
        ircserver = IRCServer((address, port), IRCClient, disable_ssl="--disable_ssl" in sys.argv)
        logger.info("Server started on %s:%d" % (address, port))
        ircserver.serve_forever()
    except socket.error as e:
        logger.error(repr(e))
        raise SystemExit(-2)

if __name__ == "__main__":
    if "--profile" in sys.argv:
        import cProfile
        cProfile.run("main()")
    else:
        main()
