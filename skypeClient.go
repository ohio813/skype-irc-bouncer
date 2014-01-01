package gsi

import "io"
import "log"
import "net"
import "fmt"
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

func (self *EventDispatcher) deleteHandler() {
}

func (self *EventDispatcher) RegisterHandler(eventName string, handler EventHandler) {
  if _, ok := self.handlers[eventName]; ! ok {
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
  if _, ok := self.handlers[eventName]; ! ok {
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
  if _, ok := self.handlers[eventName]; ! ok {
    return
  }

  for _, handler := range self.handlers[eventName] {
    if handler != nil {
      handler <- eventData
    }
  }
}

type Client struct {
  conn net.Conn
  incoming chan string
  outgoing chan string

  users map[string]*User
  groups map[string]*Group
  chats map[string]*Chat
  chatmessages map[string]*Chatmessage // TODO: get rid of this

  events *EventDispatcher
}

func MakeClient(config *Config) (*Client, error) {
  host := config.GetIfDefined("Client", "host", "localhost")
  port := config.GetIfDefined("Client", "port", "2727")
  connString := host + ":" + port
  conn, e := tls.Dial("tcp", connString, &tls.Config{InsecureSkipVerify: true})
  //  conn, e := net.Dial("tcp", connString)
  if e != nil {
    return nil, e
  }

  client := Client{
    conn: conn,
    incoming: make(chan string),
    outgoing: make(chan string),
    users: make(map[string]*User),
    groups: make(map[string]*Group),
    chats: make(map[string]*Chat),
    chatmessages: make(map[string]*Chatmessage),
    events: MakeEventDispatcher(),
  }

  go client.handleReads()
  go client.handleWrites()

  return &client, nil
}

func (self *Client) handleReads() {
  r := bufio.NewReader(self.conn)
  for {
    line, e := r.ReadString('\n')
    if e == io.EOF {
      log.Fatal(e.Error())
    }
    if e != nil {
      log.Fatal(e.Error())
    }

    line = strings.TrimRight(line, "\r\n")
    self.incoming <- line
  }
}

func (self *Client) handleWrites() {
  w := bufio.NewWriter(self.conn)
  for line := range self.outgoing {
    n, e := w.WriteString(line)
    if  e != nil {
      log.Fatal(e.Error())
    } else {
      e = w.Flush()
      if e != nil {
        log.Fatal(e.Error())
      }
      log.Printf("<<[%d] '%s'", n, strings.TrimRight(line, "\r\n"))
    }
  }
}

func (self *Client) WriteLine(line string) error {
  self.outgoing <- line + "\n"
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

func (self *Client) Serve() error {

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
    if line[:4] == "PING" {
      self.events.Emit("recv.PING", line)
    } else if line[:4] == "USER" {
      self.events.Emit("recv.USER", line)
    } else if line[:5] == "GROUP" {
      self.events.Emit("recv.GROUP", line)
    } else if line[:11] == "CHATMESSAGE" {
      self.events.Emit("recv.CHATMESSAGE", line)
    } else if line[:5] == "CHAT " {
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

    if _, ok := self.users[id]; ! ok {
      self.users[id] = &User{Id: id}
      self.events.Emit("recv.USER.new", id)
    }
    user := self.users[id]
    user.parseSet(line)
  }))

  self.events.RegisterHandler("recv.USER.new", makeHandler(func(id string) {
    user, ok := self.users[id]
    if ! ok {
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

    if _, ok := self.groups[id]; ! ok {
      self.groups[id] = &Group{Id: id}
      self.events.Emit("recv.GROUP.new", id)
    }
    group := self.groups[id]
    group.parseSet(line)
  }))

  self.events.RegisterHandler("recv.GROUP.new", makeHandler(func(id string) {
    group, ok := self.groups[id]
    if ! ok {
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

    if _, ok := self.chatmessages[id]; ! ok {
      self.chatmessages[id] = &Chatmessage{Id: id}
      self.events.Emit("recv.CHATMESSAGE.new", id)
    }
    chatmessage := self.chatmessages[id]
    chatmessage.parseSet(line)
  }))

  self.events.RegisterHandler("recv.CHATMESSAGE.new", makeHandler(func(id string) {
    chatmessage, ok := self.chatmessages[id]
    if ! ok {
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

    if _, ok := self.chats[id]; ! ok {
      self.chats[id] = &Chat{Id: id}
      self.events.Emit("recv.CHAT.new", id)
    }
    chat := self.chats[id]
    chat.parseSet(line)
  }))

  self.events.RegisterHandler("recv.CHAT.new", makeHandler(func(id string) {
    chat, ok := self.chats[id]
    if ! ok {
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


  for line := range self.incoming {
    self.events.Emit("recv", line)
  }
  return nil
}

func (self *Client) TriggerUserTest() {
  self.users["asekhar"] = &User{
    Id: "asekhar",
  }
  
  self.events.Emit("recv.USER.new", "asekhar")
}

func (self *Client) TriggerGroupTest() {
  self.groups["234"] = &Group {
    Id: "234",
  }
  self.events.Emit("recv.GROUP.new", "234")
}

func (self *Client) TriggerChatmessageTest() {
  self.chatmessages["211977"] = &Chatmessage {
    Id: "211977",
  }
  self.events.Emit("recv.CHATMESSAGE.new", "211977")
}

func (self *Client) TriggerChatTest() {
  id := "#ejeangeo/$williballenthin;bf4e426f52c0b1f6"
  self.chats[id] = &Chat {
    Id: id,
  }
  self.events.Emit("recv.CHAT.new", id)
}

func (self *Client) DumpUsers(writer io.Writer) {
  w := bufio.NewWriter(writer)
  w.WriteString("USERS" + "\n")
  for id, user := range self.users {
    w.WriteString("  " + id + "\n")
    w.WriteString("    About:" + user.About + "\n")
    w.WriteString("    Country:" + user.Country + "\n")
    w.WriteString("    Birthday:" + user.Birthday + "\n")
    w.WriteString("    Displayname:" + user.Displayname + "\n")
    w.WriteString("    Language:" + user.Language + "\n")
    w.WriteString("    Onlinestatus:" + user.Onlinestatus + "\n")
    w.WriteString("    Sex:" + user.Sex + "\n")
    w.WriteString("    MoodText:" + user.MoodText + "\n")
    w.WriteString("    Aliases:" + user.Aliases + "\n")
    w.WriteString("    Lastonlinetimestamp:" + user.Lastonlinetimestamp + "\n")
    w.WriteString("    Buddystatus:" + user.Buddystatus + "\n")
    w.WriteString("    NrofAuthedBuddies:" + user.NrofAuthedBuddies + "\n")
    w.WriteString("    City:" + user.City + "\n")
    w.WriteString("    Avatar:" + user.Avatar + "\n")
    w.WriteString("    RichMoodText:" + user.RichMoodText + "\n")
    w.WriteString("    Fullname:" + user.Fullname + "\n")
  }
  w.Flush()
}

func (self *Client) DumpGroups(writer io.Writer) {
  w := bufio.NewWriter(writer)
  w.WriteString("GROUPS" + "\n")
  for id, group := range self.groups {
    w.WriteString("  " + id + "\n")
    w.WriteString("    Displayname:" + group.Displayname + "\n")
    w.WriteString("    Users:" + group.Users + "\n")
    w.WriteString("    Expanded:" + group.Expanded + "\n")
    w.WriteString("    CustomGroupId:" + group.CustomGroupId + "\n")
    w.WriteString("    Visible:" + group.Visible + "\n")
    w.WriteString("    GroupType:" + group.GroupType + "\n")
    w.WriteString("    NrofUsers:" + group.NrofUsers + "\n")
  }
  w.Flush()
}

func (self *Client) DumpChatmessages(writer io.Writer) {
  w := bufio.NewWriter(writer)
  w.WriteString("CHATMESSAGES" + "\n")
  for id, chatmessage := range self.chatmessages {
    w.WriteString("  " + id + "\n")
    w.WriteString("    Body:" + chatmessage.Body + "\n")
    w.WriteString("    Status:" + chatmessage.Status + "\n")
    w.WriteString("    EditedTimestamp:" + chatmessage.EditedTimestamp + "\n")
    w.WriteString("    EditedBy:" + chatmessage.EditedBy + "\n")
    w.WriteString("    Users:" + chatmessage.Users + "\n")
    w.WriteString("    Timestamp:" + chatmessage.Timestamp + "\n")
    w.WriteString("    FromHandle:" + chatmessage.FromHandle + "\n")
    w.WriteString("    Chatname:" + chatmessage.Chatname + "\n")
    w.WriteString("    IsEditable:" + chatmessage.IsEditable + "\n")
    w.WriteString("    Leavereason:" + chatmessage.Leavereason + "\n")
    w.WriteString("    FromDispname:" + chatmessage.FromDispname + "\n")
    w.WriteString("    ChatmessageType:" + chatmessage.ChatmessageType + "\n")
  }
  w.Flush()
}

func (self *Client) DumpChats(writer io.Writer) {
  w := bufio.NewWriter(writer)
  w.WriteString("CHAT" + "\n")
  for id, chat := range self.chats {
    w.WriteString("  " + id + "\n")
    w.WriteString("    Id :" + chat.Id  + "\n")
    w.WriteString("    DialogPartner :" + chat.DialogPartner  + "\n")
    w.WriteString("    Passwordhint :" + chat.Passwordhint  + "\n")
    w.WriteString("    Options :" + chat.Options  + "\n")
    w.WriteString("    Applicants :" + chat.Applicants  + "\n")
    w.WriteString("    Bookmarked :" + chat.Bookmarked  + "\n")
    w.WriteString("    Recentchatmessages :" + chat.Recentchatmessages  + "\n")
    w.WriteString("    Chatname :" + chat.Chatname  + "\n")
    w.WriteString("    Adder :" + chat.Adder  + "\n")
    w.WriteString("    ActivityTimestamp :" + chat.ActivityTimestamp  + "\n")
    w.WriteString("    Posters :" + chat.Posters  + "\n")
    w.WriteString("    Status :" + chat.Status  + "\n")
    w.WriteString("    Guidelines :" + chat.Guidelines  + "\n")
    w.WriteString("    Topicxml :" + chat.Topicxml  + "\n")
    w.WriteString("    Mystatus :" + chat.Mystatus  + "\n")
    w.WriteString("    Memberobjects :" + chat.Memberobjects  + "\n")
    w.WriteString("    Friendlyname :" + chat.Friendlyname  + "\n")
    w.WriteString("    Activemembers :" + chat.Activemembers  + "\n")
    w.WriteString("    Description :" + chat.Description  + "\n")
    w.WriteString("    Timestamp :" + chat.Timestamp  + "\n")
    w.WriteString("    Chatmessages :" + chat.Chatmessages  + "\n")
    w.WriteString("    Topic :" + chat.Topic  + "\n")
    w.WriteString("    Role :" + chat.Role  + "\n")
    w.WriteString("    Blob :" + chat.Blob  + "\n")
    w.WriteString("    Members :" + chat.Members  + "\n")
    w.WriteString("    Myrole :" + chat.Myrole  + "\n")
  }
  w.Flush()
}




func LoadConfig(file string) (*Config, error) {
  // TODO: validation here
	return ReadConfig(file)
}


