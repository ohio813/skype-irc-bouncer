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

func (self *Client) getUser(id string) (user *User, didCreate bool, e error) {
	// TODO(wb): locking
	if user, ok := self.users[id]; ok {
		return user, false, nil
	} else {
		self.users[id] = &User{
			Id: id,
		}
		user = self.users[id]
		for _, cmd := range user.getFetchAllFieldsCommands() {
			self.WriteLine(cmd)
		}
		return user, true, nil
	}
}

func (self *Client) touchUser(id string) (didCreate bool, e error) {
	_, didCreate, e = self.getUser(id)
	return didCreate, e
}

func (self *Client) getGroup(id string) (group *Group, didCreate bool, e error) {
	// TODO(wb): locking
	if group, ok := self.groups[id]; ok {
		return group, false, nil
	} else {
		self.groups[id] = &Group{
			Id: id,
		}

		group = self.groups[id]
		for _, cmd := range group.getFetchAllFieldsCommands() {
			self.WriteLine(cmd)
		}
		return group, true, nil
	}
}

func (self *Client) touchGroup(id string) (didCreate bool, e error) {
	_, didCreate, e = self.getGroup(id)
	return didCreate, e
}

func (self *Client) getChat(id string) (chat *Chat, didCreate bool, e error) {
	// TODO(wb): locking
	if chat, ok := self.chats[id]; ok {
		return chat, false, nil
	} else {
		self.chats[id] = &Chat{
			Id: id,
		}

		chat = self.chats[id]
		for _, cmd := range chat.getFetchAllFieldsCommands() {
			self.WriteLine(cmd)
		}
		return chat, true, nil
	}
}

func (self *Client) touchChat(id string) (didCreate bool, e error) {
	_, didCreate, e = self.getChat(id)
	return didCreate, e
}

func (self *Client) getChatmessage(id string) (chatmessage *Chatmessage, didCreate bool, e error) {
	// TODO(wb): locking
	if chatmessage, ok := self.chatmessages[id]; ok {
		return chatmessage, false, nil
	} else {
		self.chatmessages[id] = &Chatmessage{
			Id: id,
		}

		chatmessage = self.chatmessages[id]
		for _, cmd := range chatmessage.getFetchAllFieldsCommands() {
			self.WriteLine(cmd)
		}
		return chatmessage, true, nil
	}

}

func (self *Client) touchChatmessage(id string) (didCreate bool, e error) {
	_, didCreate, e = self.getChatmessage(id)
	return didCreate, e
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

// create an EventHandler and attach a function to it that handles each event
func makeHandler(fn func(string)) EventHandler {
	Q := make(EventHandler)

	go func() {
		for eventData := range Q {
			fn(eventData)
		}
	}()

	return Q
}

// get the second field from the Skype message
func getSkypeMessageObjectId(objectType string, s string) (string, error) {
	var id string
	if n, e := fmt.Sscanf(s, objectType + " %s", &id); e != nil {
		return "", e
	} else if n != 1 {
		return "", UnexpectedNumberOfFieldsError
	}
	return id, nil
}

// get the third field from the Skype message
func getSkypeMessageFieldName(objectType string, s string) (string, error) {
	var id string
	var fieldName string
	if n, e := fmt.Sscanf(s, objectType + " %s %s", &id, &fieldName); e != nil {
		return "", e
	} else if n != 2 {
		return "", UnexpectedNumberOfFieldsError
	}
	return fieldName, nil
}


func (self *Client) setupInternalHandlers() error {
	/***
	 * Here are the events and their associated data schema we use here:
	 *   - "recv" : entire line from Skype4Py
	 *   - "recv.PING" : entire line from Skype4Py
	 *   - "recv.USER" : entire line from Skype4Py
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
		id, e := getSkypeMessageObjectId("USER", line)
		if e != nil {
			return
		}

		user, didCreate, e := self.getUser(id)
		if e != nil {
			return 
		}
		if didCreate {

		}
		user.parseSet(line)

		fieldName, e := getSkypeMessageFieldName("USER", line)
		if e != nil {
			return
		}
		self.events.Emit("recv.USER." + fieldName, line)
	}))

	self.events.RegisterHandler("recv.GROUP", makeHandler(func(line string) {
		id, e := getSkypeMessageObjectId("GROUP", line)
		if e != nil {
			return
		}

		group, didCreate,  e := self.getGroup(id)
		if e != nil {
			return
		}
		if didCreate {

		}
		group.parseSet(line)

		fieldName, e := getSkypeMessageFieldName("GROUP", line)
		if e != nil {
			return
		}
		self.events.Emit("recv.GROUP." + fieldName, line)
	}))

	self.events.RegisterHandler("recv.GROUP.USERS", makeHandler(func(line string) {
		id, e := getSkypeMessageObjectId("GROUP", line)
		if e != nil {
			return
		}

		group, _, e := self.getGroup(id)
		if e != nil {
			return
		}
		for _, userid := range group.Users {
			self.touchUser(userid)
		}
	}))

	self.events.RegisterHandler("recv.CHATMESSAGE", makeHandler(func(line string) {
		id, e := getSkypeMessageObjectId("CHATMESSAGE", line)
		if e != nil {
			return
		}

		chatmessage, didCreate, e := self.getChatmessage(id)
		if e != nil {
			return
		}
		if didCreate {

		}
		chatmessage.parseSet(line)

		fieldName, e := getSkypeMessageFieldName("CHATMESSAGE", line)
		if e != nil {
			return
		}
		self.events.Emit("recv.CHATMESSAGE." + fieldName, line)
	}))

	self.events.RegisterHandler("recv.CHATMESSAGE.FROM_HANDLE", makeHandler(func(line string) {
		id, e := getSkypeMessageObjectId("CHATMESSAGE", line)
		if e != nil {
			return
		}

		chatmessage, _, e := self.getChatmessage(id)
		if e != nil {
			return
		}
		self.touchUser(chatmessage.FromHandle)
	}))

	self.events.RegisterHandler("recv.CHATMESSAGE.CHATNAME", makeHandler(func(line string) {
		id, e := getSkypeMessageObjectId("CHATMESSAGE", line)
		if e != nil {
			return
		}

		chatmessage, _, e := self.getChatmessage(id)
		if e != nil {
			return
		}
		self.touchChat(chatmessage.Chatname)
	}))

	self.events.RegisterHandler("recv.CHAT", makeHandler(func(line string) {
		id, e := getSkypeMessageObjectId("CHAT", line)
		if e != nil {
			return
		}

		chat, didCreate, e := self.getChat(id)
		if e != nil {
			return
		}
		if didCreate {

		}
		chat.parseSet(line)

		fieldName, e := getSkypeMessageFieldName("CHAT", line)
		if e != nil {
			return
		}
		self.events.Emit("recv.CHAT." + fieldName, line)
	}))

	self.events.RegisterHandler("recv.CHAT.RECENTCHATMESSAGES", makeHandler(func(line string) {
		id, e := getSkypeMessageObjectId("CHAT", line)
		if e != nil {
			return
		}

		chat, _, e := self.getChat(id)
		if e != nil {
			return
		}
		for _, chatmessageid := range chat.Recentchatmessages {
			self.touchChatmessage(chatmessageid)
		}
	}))

	self.events.RegisterHandler("recv.CHAT.MESSAGES", makeHandler(func(line string) {
		id, e := getSkypeMessageObjectId("CHAT", line)
		if e != nil {
			return
		}

		chat, _, e := self.getChat(id)
		if e != nil {
			return
		}
		for _, chatmessageid := range chat.Chatmessages {
			self.touchChatmessage(chatmessageid)
		}
	}))

	self.events.RegisterHandler("recv.CHAT.MEMBERS", makeHandler(func(line string) {
		id, e := getSkypeMessageObjectId("CHAT", line)
		if e != nil {
			return
		}

		chat, _, e := self.getChat(id)
		if e != nil {
			return
		}
		for _, userid := range chat.Members {
			self.touchUser(userid)
		}
	}))

	self.events.RegisterHandler("recv.CHAT.POSTERS", makeHandler(func(line string) {
		id, e := getSkypeMessageObjectId("CHAT", line)
		if e != nil {
			return
		}

		chat, _, e := self.getChat(id)
		if e != nil {
			return
		}
		for _, userid := range chat.Posters {
			self.touchUser(userid)
		}
	}))

	self.events.RegisterHandler("recv.CHAT.ACTIVEMEMBERS", makeHandler(func(line string) {
		id, e := getSkypeMessageObjectId("CHAT", line)
		if e != nil {
			return
		}

		chat, _, e := self.getChat(id)
		if e != nil {
			return
		}
		for _, userid := range chat.Activemembers {
			self.touchUser(userid)
		}
	}))

	self.events.RegisterHandler("recv.CHAT.DialogPartner", makeHandler(func(line string) {
		id, e := getSkypeMessageObjectId("CHAT", line)
		if e != nil {
			return
		}

		chat, _, e := self.getChat(id)
		if e != nil {
			return
		}
		self.touchUser(chat.DialogPartner)
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
