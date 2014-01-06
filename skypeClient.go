package gsi

import "os"
import "io"
import "log"
import "net"
import "fmt"
import "time"
import "bufio"
import "strings"
import "crypto/tls"

type EventHandler chan string

type EventDispatcher struct {
	handlers map[string][]EventHandler
}

func MakeEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (self *EventDispatcher) RegisterHandler(eventName string, handler EventHandler) {
	if _, ok := self.handlers[eventName]; !ok {
		self.handlers[eventName] = make([]EventHandler, 5)
	}

	for _, existing_handler := range self.handlers[eventName] {
		if existing_handler == handler {
			return
		}
	}
	self.handlers[eventName] = append(self.handlers[eventName], handler)
}

func (self *EventDispatcher) UnregisterHandler(eventName string, handler EventHandler) {
	if _, ok := self.handlers[eventName]; !ok {
		return
	}

	for i, existing_handler := range self.handlers[eventName] {
		if existing_handler == handler {
			// ugly, but from http://code.google.com/p/go-wiki/wiki/SliceTricks
			copy(self.handlers[eventName][i:], self.handlers[eventName][i+1:])
			self.handlers[eventName][len(self.handlers[eventName])-1] = nil
			self.handlers[eventName] = self.handlers[eventName][:len(self.handlers[eventName])-1]
			return
		}
	}
}

func (self *EventDispatcher) Emit(eventName string, eventData string) {
	if _, ok := self.handlers[eventName]; !ok {
		return
	}

	for _, handler := range self.handlers[eventName] {
		if handler != nil {
			go func() {
				handler <- eventData
			}()
		}
	}
}

type SkypeConnection struct {
	incoming chan string
	outgoing chan string
}

func MakeSkypeConnection() *SkypeConnection {
	return &SkypeConnection{
		incoming: make(chan string),
		outgoing: make(chan string),
	}
}

// skypeMessageReader parses lines from an io.Reader and places them into a SkypeConnection.
// from the network to this client.
func skypeMessageReader(sc *SkypeConnection, r io.Reader) error {
	rr := bufio.NewReader(r)
	for {
		line, e := rr.ReadString('\n')
		if e == io.EOF {
			log.Printf("reader: eof")
			return nil
		}
		if e != nil {
			log.Fatal(e.Error())
			return e
		}
		
		line = strings.TrimRight(line, "\r\n")
		sc.incoming <- line
	}
	return nil
}

// skypeMessageWriter places outgoing messages from an io.Writer into a SkypeConnection.
// from this client to the network.
func skypeMessageWriter(sc *SkypeConnection, w io.Writer) error {
	ww := bufio.NewWriter(w)
	for line := range sc.outgoing {
		towrite := fmt.Sprintf("%s\r\n", strings.TrimRight(line, "\r\n"))
		n, e := ww.WriteString(towrite)
		if e != nil {
			log.Fatal(e.Error())
			return e
		} else {
			e = ww.Flush()
			if e != nil {
				log.Fatal(e.Error())
				return e
			}
			// TODO(wb): don't log here
			log.Printf("<<[%d] '%s'", n, strings.TrimRight(towrite, "\r\n"))
		}
	}
	return nil
}

func MakeTLSSkypeConnection(host string, port int) (*SkypeConnection, error) {
	connString := fmt.Sprintf("%s:%d", host, port)
	conn, e := tls.Dial("tcp", connString, &tls.Config{InsecureSkipVerify: true}) // TODO(wb):verify
	if e != nil {
		return nil, e
	}
	
	sc := MakeSkypeConnection()
	go skypeMessageReader(sc, conn)
	go skypeMessageWriter(sc, conn)
	return sc, nil
}

// MakeTCPSkypeConnection creates a connection to a Skype proxy using unencrypted TCP.
func MakeTCPSkypeConnection(host string, port int) (*SkypeConnection, error) {
	connString := fmt.Sprintf("%s:%d", host, port)
	conn, e := net.Dial("tcp", connString)
	if e != nil {
		return nil, e
	}
	
	sc := MakeSkypeConnection()
	go skypeMessageReader(sc, conn)
	go skypeMessageWriter(sc, conn)
	return sc, nil
}

func MakeFileStubbedSkypeConnection(filename string) (*SkypeConnection, error) {
	f, e := os.Open(filename)
	if e != nil {
		return nil, e
	}
	sc := MakeSkypeConnection()

	go skypeMessageReader(sc, f)
	go func() {
//		for _ = range sc.outgoing {
		for line := range sc.outgoing {
			log.Printf("<<[%d] (stubbed write) '%s'", len(line), strings.TrimRight(line, "\r\n"))
		}
	}()
	return sc, nil
}

type Client struct {
	conn *SkypeConnection

	users        map[string]*User
	groups       map[string]*Group
	chats        map[string]*Chat
	chatmessages map[string]*Chatmessage // TODO: get rid of this

	events *EventDispatcher
}

func MakeClient(config *Config, conn *SkypeConnection) (*Client, error) {
	client := Client{
		conn:         conn,
		users:        make(map[string]*User),
		groups:       make(map[string]*Group),
		chats:        make(map[string]*Chat),
		chatmessages: make(map[string]*Chatmessage),
		events:       MakeEventDispatcher(),
	}
	
	if e := client.setupInternalHandlers(); e != nil {
		return nil, e
	}

	return &client, nil
}

func (self *Client) WriteLine(line string) error {
	self.conn.outgoing <- line + "\n"
	return nil
}

func (self *Client) Authenticate(username string, password string) error {
	self.WriteLine("USERNAME " + username)
	self.WriteLine("PASSWORD " + password)
	return nil
}

func makeHandler(fn func(string)) EventHandler {
	Q := make(EventHandler)

	go func() {
		for eventData := range Q {
			fn(eventData)
		}
	}()

	return Q
}


func (self *Client) setupInternalHandlers() error {
	/***
	 * Here are the events and their associated data schema we use here:
	 *   - "recv" : entire line from Skype4Py
	 *   - "recv.PING" : entire line from Skype4Py
	 *   - "recv.USER" : entire line from Skype4Py
	 *   - "recv.USER.new" : the new User id
	 *
	 */

	self.events.RegisterHandler("recv", makeHandler(func(line string) {
		log.Printf(">>[%d] '%s'", len(line), strings.TrimRight(line, "\r\n"))
	}))

	self.events.RegisterHandler("recv", makeHandler(func(line string) {
		if len(line) > 4 && line[:4] == "PING" {
			self.events.Emit("recv.PING", line)
		} else if len(line) > 4 && line[:4] == "USER" {
			self.events.Emit("recv.USER", line)
		} else if len(line) > 5 && line[:5] == "GROUP" {
			self.events.Emit("recv.GROUP", line)
		} else if len(line) > 11 && line[:11] == "CHATMESSAGE" {
			self.events.Emit("recv.CHATMESSAGE", line)
		} else if len(line) > 5 && line[:5] == "CHAT " {
			self.events.Emit("recv.CHAT", line)
		}
	}))

	self.events.RegisterHandler("recv.PING", makeHandler(func(line string) {
		self.WriteLine("PONG")
	}))

	self.events.RegisterHandler("recv.USER", makeHandler(func(line string) {
		var id string
		if n, e := fmt.Sscanf(line, "USER %s", &id); e != nil {
			return
		} else if n != 1 {
			return
		}

		if _, ok := self.users[id]; !ok {
			self.users[id] = &User{Id: id}
			self.events.Emit("recv.USER.new", id)
		}
		user := self.users[id]
		user.parseSet(line)
	}))

	self.events.RegisterHandler("recv.USER.new", makeHandler(func(id string) {
		user, ok := self.users[id]
		if !ok {
			return
		}

		cmds, e := user.getFetchAllFieldsCommands()
		if e != nil {
			// TODO: get rid of the error on this fetch
			return
		}
		for _, cmd := range cmds {
			self.WriteLine(cmd)
		}
	}))

	self.events.RegisterHandler("recv.GROUP", makeHandler(func(line string) {
		var id string
		if n, e := fmt.Sscanf(line, "GROUP %s", &id); e != nil {
			return
		} else if n != 1 {
			return
		}

		if _, ok := self.groups[id]; !ok {
			self.groups[id] = &Group{Id: id}
			self.events.Emit("recv.GROUP.new", id)
		}
		group := self.groups[id]
		group.parseSet(line)
	}))

	self.events.RegisterHandler("recv.GROUP.new", makeHandler(func(id string) {
		group, ok := self.groups[id]
		if !ok {
			return
		}

		cmds, e := group.getFetchAllFieldsCommands()
		if e != nil {
			// TODO: get rid of the error on this fetch
			return
		}
		for _, cmd := range cmds {
			self.WriteLine(cmd)
		}
	}))

	self.events.RegisterHandler("recv.CHATMESSAGE", makeHandler(func(line string) {
		var id string
		if n, e := fmt.Sscanf(line, "CHATMESSAGE %s", &id); e != nil {
			return
		} else if n != 1 {
			return
		}

		if _, ok := self.chatmessages[id]; !ok {
			self.chatmessages[id] = &Chatmessage{Id: id}
			self.events.Emit("recv.CHATMESSAGE.new", id)
		}
		chatmessage := self.chatmessages[id]
		chatmessage.parseSet(line)
	}))

	self.events.RegisterHandler("recv.CHATMESSAGE.new", makeHandler(func(id string) {
		chatmessage, ok := self.chatmessages[id]
		if !ok {
			return
		}

		cmds, e := chatmessage.getFetchAllFieldsCommands()
		if e != nil {
			// TODO: get rid of the error on this fetch
			return
		}
		for _, cmd := range cmds {
			self.WriteLine(cmd)
		}
	}))

	self.events.RegisterHandler("recv.CHAT", makeHandler(func(line string) {
		var id string
		if n, e := fmt.Sscanf(line, "CHAT %s", &id); e != nil {
			return
		} else if n != 1 {
			return
		}

		if _, ok := self.chats[id]; !ok {
			self.chats[id] = &Chat{Id: id}
			self.events.Emit("recv.CHAT.new", id)
		}
		chat := self.chats[id]
		chat.parseSet(line)
	}))

	self.events.RegisterHandler("recv.CHAT.new", makeHandler(func(id string) {
		chat, ok := self.chats[id]
		if !ok {
			return
		}

		cmds, e := chat.getFetchAllFieldsCommands()
		if e != nil {
			// TODO: get rid of the error on this fetch
			return
		}
		for _, cmd := range cmds {
			self.WriteLine(cmd)
		}
	}))
	return nil
}

func (self *Client) ServeForever() error {
	for line := range self.conn.incoming {
		self.events.Emit("recv", line)
	}
	return nil
}

func (self *Client) ServeForDuration(duration time.Duration) error {
	var line string
	for {
		select {
		case line = <- self.conn.incoming:
			log.Printf("Emitting line: " + line)
			self.events.Emit("recv", line)
		case <- time.After(duration):
			return nil
		}
	}
	return nil
}


func LoadConfig(file string) (*Config, error) {
	// TODO: validation here
	return ReadConfig(file)
}
