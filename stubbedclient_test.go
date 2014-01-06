package gsi

import "os"
import "io"
import "time"
import "bufio"
import "strings"
import "testing"

func Test(t *testing.T) {
	
	conn, e := MakeFileStubbedSkypeConnection("test/test.txt")
	if e != nil {
		t.Error(e.Error())
		return
	}

	config, e := LoadConfig("test/gsi.cfg")
	if e != nil {
		t.Error(e.Error())
		return
	}

	client, e := MakeClient(config, conn)
	if e != nil {
		t.Error(e.Error())
		return
	}

	client.ServeForDuration(5 * time.Second)

	client.DumpUsers(os.Stdout)
	client.DumpGroups(os.Stdout)
	client.DumpChatmessages(os.Stdout)
	client.DumpChats(os.Stdout)
}


func (self *Client) TriggerUserTest() {
	self.users["asekhar"] = &User{
		Id: "asekhar",
	}

	self.events.Emit("recv.USER.new", "asekhar")
}

func (self *Client) TriggerGroupTest() {
	self.groups["234"] = &Group{
		Id: "234",
	}
	self.events.Emit("recv.GROUP.new", "234")
}

func (self *Client) TriggerChatmessageTest() {
	self.chatmessages["211977"] = &Chatmessage{
		Id: "211977",
	}
	self.events.Emit("recv.CHATMESSAGE.new", "211977")
}

func (self *Client) TriggerChatTest() {
	id := "#ejeangeo/$williballenthin;bf4e426f52c0b1f6"
	self.chats[id] = &Chat{
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
		w.WriteString("    Id: " + chat.Id + "\n")
		w.WriteString("    DialogPartner: " + chat.DialogPartner + "\n")
		w.WriteString("    Passwordhint: '" + chat.Passwordhint + "'\n")
		w.WriteString("    Options: " + chat.Options + "\n")
		w.WriteString("    Applicants: " + chat.Applicants + "\n")
		w.WriteString("    Bookmarked: " + chat.Bookmarked + "\n")
		w.WriteString("    Recentchatmessages: " + strings.Join(chat.Recentchatmessages, ", ") + "\n")
		w.WriteString("    Chatname: '" + chat.Chatname + "'\n")
		w.WriteString("    Adder: " + chat.Adder + "\n")
		w.WriteString("    ActivityTimestamp: " + chat.ActivityTimestamp + "\n")
		w.WriteString("    Posters: " + strings.Join(chat.Posters, ", ") + "\n")
		w.WriteString("    Status: " + chat.Status + "\n")
		w.WriteString("    Guidelines: " + chat.Guidelines + "\n")
		w.WriteString("    Topicxml: '" + chat.Topicxml + "'\n")
		w.WriteString("    Mystatus: " + chat.Mystatus + "\n")
		w.WriteString("    Memberobjects: " + chat.Memberobjects + "\n")
		w.WriteString("    Friendlyname: '" + chat.Friendlyname + "'\n")
		w.WriteString("    Activemembers: " + strings.Join(chat.Activemembers, ", ") + "\n")
		w.WriteString("    Description: '" + chat.Description + "'\n")
		w.WriteString("    Timestamp: " + chat.Timestamp + "\n")
		w.WriteString("    Chatmessages: " + strings.Join(chat.Chatmessages, ", ") + "\n")
		w.WriteString("    Topic: '" + chat.Topic + "'\n")
		w.WriteString("    Role: " + chat.Role + "\n")
		w.WriteString("    Blob: " + chat.Blob + "\n")
		w.WriteString("    Members: " + strings.Join(chat.Members, ", ") + "\n")
		w.WriteString("    Myrole: " + chat.Myrole + "\n")
	}
	w.Flush()
}

