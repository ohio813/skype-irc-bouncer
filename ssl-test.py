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

import six

from irc import buffer


class ClientHandler(six.moves.socketserver.BaseRequestHandler):
    class Disconnect(BaseException):
        pass

    def __init__(self, request, client_address, server):
        self.send_queue = []
        self._logger = logging.getLogger("IRCClient")

        six.moves.socketserver.BaseRequestHandler.__init__(self, request,
                                                           client_address, server)

    def handle(self):
        self._logger.info('Client connected.')
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
            self._logger.error(in_error)
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
        except Exception as e:
            self._logger.error(e)
            raise self.Disconnect()

        if not data:
            self._logger.error("no data")
            raise self.Disconnect()

        self.buffer.feed(data)
        for line in self.buffer:
            line = line.decode('utf-8')
            self._handle_line(line)

    def _handle_line(self, line):
        try:
            self._logger.debug('> %s' % (line))
            self.send_queue.append(line)
        except Exception as e:
            response = 'ERROR %r' % (e)
            self._logger.error(response)
            raise

    def _send(self, msg):
        self._logger.debug('< %s', msg)
        self.request.send(msg.encode('utf-8') + b'\r\n')


class Server(six.moves.socketserver.ThreadingMixIn,
             six.moves.socketserver.TCPServer):
    daemon_threads = True
    allow_reuse_address = True

    def __init__(self, *args, **kwargs):
        self._logger = logging.getLogger("Server")

        six.moves.socketserver.TCPServer.__init__(self, *args, **kwargs)

        comment = """"""
        self.socket = ssl.wrap_socket(self.socket, 
                                      keyfile="./ssl/server.key",
                                      certfile="./ssl/serverca.pem", 
                                      server_side=True, 
                                      cert_reqs=ssl.CERT_REQUIRED, 
                                      ca_certs="./ssl/clientca.pem", 
                                      do_handshake_on_connect=True)

def main():
    logger = logging.getLogger("main")
    logging.basicConfig(format='%(asctime)s %(message)s', level=logging.INFO)

    address = "0.0.0.0"
    port = 6667

    try:
        server = Server((address, port), ClientHandler)
        logger.info("Server started on %s:%d" % (address, port))
        server.serve_forever()
    except socket.error as e:
        logger.error(repr(e))
        raise SystemExit(-2)

if __name__ == "__main__":
    main()
