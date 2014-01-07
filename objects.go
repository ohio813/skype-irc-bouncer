package gsi

import "fmt"
import "errors"
import "strings"

var UnexpectedNumberOfFieldsError = errors.New("Encountered an unexpected number of fields to parse")
var ParserForFieldDoesNotExist = errors.New("Parser for the given field does not exist")

func getSkypeMessageSetSimpleString(t string, f string, s string) (string, error) {
	var __ string
	var ret string
	if n, e := fmt.Sscanf(s, t + " %s " + f + " %s", &__, &ret); e != nil {
		return "", e
	} else if n != 2 {
		return "", UnexpectedNumberOfFieldsError
	}
	return ret, nil
}

func getSkypeMessageSetComplexString(t string, f string, s string) (string, error) {
	var id string
	if n, e := fmt.Sscanf(s, t + " %s " + f + " ", &id); e != nil {
		return "", e
	} else if n != 1 {
		return "", UnexpectedNumberOfFieldsError
	}

	remainder := s[len(fmt.Sprintf(t + " %s " + f + " ", id)):]
	return strings.TrimRight(remainder, "\r\n\t "), nil
}

func getSkypeMessageSetList(t string, f string, s string, sep string) ([]string, error) {
	var id string
	if n, e := fmt.Sscanf(s, t + " %s " + f + " ", &id); e != nil {
		return []string{}, e
	} else if n != 1 {
		return []string{}, UnexpectedNumberOfFieldsError
	}

	remainder := s[len(fmt.Sprintf(t + " %s " + f + " ", id)):]
	return strings.Split(remainder, sep), nil
}

type Chat struct {
	Id                 string
	DialogPartner      string
	Passwordhint       string // string may contain space
	Options            string // integer
	Applicants         string // TODO(wb): could be a list?
	Bookmarked         string
	Recentchatmessages []string  // ,-separated list
	Chatname           string // string may contain space
	Adder              string
	ActivityTimestamp  string
	Posters            []string  // space-separated list
	Status             string
	Guidelines         string // TODO(wb): could be a list?
	Topicxml           string // string may contain space
	Mystatus           string
	Memberobjects      string // TODO(wb): could be a list?
	Friendlyname       string // string may contain space
	Activemembers      []string  // space-separated list
	Description        string // string may contain space
	Timestamp          string
	Chatmessages       []string  // ,-separated list
	Topic              string // string may contain space
	Role               string
	Blob               string
	Members            []string  // space-separated list
	Myrole             string
}

func (self *Chat) getFetchAllFieldsCommands() ([]string, error) {
	return []string{
		"GET CHAT " + self.Id + " DIALOG_PARTNER",
		"GET CHAT " + self.Id + " PASSWORDHINT",
		"GET CHAT " + self.Id + " OPTIONS",
		"GET CHAT " + self.Id + " APPLICANTS",
		"GET CHAT " + self.Id + " BOOKMARKED",
		"GET CHAT " + self.Id + " RECENTCHATMESSAGES",
		"GET CHAT " + self.Id + " CHATNAME",
		"GET CHAT " + self.Id + " ADDER",
		"GET CHAT " + self.Id + " ACTIVITY_TIMESTAMP",
		"GET CHAT " + self.Id + " POSTERS",
		"GET CHAT " + self.Id + " STATUS",
		"GET CHAT " + self.Id + " GUIDELINES",
		"GET CHAT " + self.Id + " TOPICXML",
		"GET CHAT " + self.Id + " MYSTATUS",
		"GET CHAT " + self.Id + " MEMBEROBJECTS",
		"GET CHAT " + self.Id + " FRIENDLYNAME",
		"GET CHAT " + self.Id + " ACTIVEMEMBERS",
		"GET CHAT " + self.Id + " DESCRIPTION",
		"GET CHAT " + self.Id + " TIMESTAMP",
		"GET CHAT " + self.Id + " CHATMESSAGES",
		"GET CHAT " + self.Id + " TOPIC",
		"GET CHAT " + self.Id + " ROLE",
		"GET CHAT " + self.Id + " BLOB",
		"GET CHAT " + self.Id + " MEMBERS",
		"GET CHAT " + self.Id + " MYROLE"}, nil
}

func (self *Chat) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "CHAT "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return UnexpectedNumberOfFieldsError
	}
	
	switch field_to_set {
	case "DIALOG_PARTNER":
		ret, e := getSkypeMessageSetSimpleString("CHAT", "DIALOG_PARTNER", s)
		if e != nil {
			return e
		}
		self.DialogPartner = ret
	case "PASSWORDHINT":
		ret, e := getSkypeMessageSetComplexString("CHAT", "PASSWORDHINT", s)
		if e != nil {
			return e
		}
		self.Passwordhint = ret
	case "OPTIONS":
		ret, e := getSkypeMessageSetSimpleString("CHAT", "OPTIONS", s)
		if e != nil {
			return e
		}
		self.Options = ret
	case "Applicants":
		ret, e := getSkypeMessageSetSimpleString("CHAT", "APPLICANTS", s)
		if e != nil {
			return e
		}
		self.Applicants = ret
	case "BOOKMARKED":
		ret, e := getSkypeMessageSetSimpleString("CHAT", "BOOKMARKED", s)
		if e != nil {
			return e
		}
		self.Bookmarked = ret
	case "RECENTCHATMESSAGES":
		ret, e := getSkypeMessageSetList("CHAT", "RECENTCHATMESSAGES", s, ", ")
		if e != nil {
			return e
		}
		self.Recentchatmessages = ret
	case "CHATNAME":
		ret, e := getSkypeMessageSetSimpleString("CHAT", "CHATNAME", s)
		if e != nil {
			return e
		}
		self.Chatname = ret
	case "ADDER":
		ret, e := getSkypeMessageSetSimpleString("CHAT", "ADDER", s)
		if e != nil {
			return e
		}
		self.Adder = ret
	case "ACTIVITY_TIMESTAMP":
		ret, e := getSkypeMessageSetSimpleString("CHAT", "ACTIVITY_TIMESTAMP", s)
		if e != nil {
			return e
		}
		self.ActivityTimestamp = ret
	case "POSTERS":
		ret, e := getSkypeMessageSetList("CHAT", "POSTERS", s, " ")
		if e != nil {
			return e
		}
		self.Posters = ret
	case "STATUS":
		ret, e := getSkypeMessageSetSimpleString("CHAT", "STATUS", s)
		if e != nil {
			return e
		}
		self.Status = ret
	case "GUIDELINES":
		ret, e := getSkypeMessageSetSimpleString("CHAT", "GUIDELINES", s)
		if e != nil {
			return e
		}
		self.Guidelines = ret
	case "TOPICXML":
		ret, e := getSkypeMessageSetComplexString("CHAT", "TOPICXML", s)
		if e != nil {
			return e
		}
		self.Topicxml = ret
	case "MYSTATUS":
		ret, e := getSkypeMessageSetSimpleString("CHAT", "MYSTATUS", s)
		if e != nil {
			return e
		}
		self.Mystatus = ret
	case "MEMBEROBJECTS":
		ret, e := getSkypeMessageSetSimpleString("CHAT", "MEMBEROBJECTS", s)
		if e != nil {
			return e
		}
		self.Memberobjects = ret
	case "FRIENDLYNAME":
		ret, e := getSkypeMessageSetComplexString("CHAT", "FRIENDLYNAME", s)
		if e != nil {
			return e
		}
		self.Friendlyname = ret
	case "ACTIVEMEMBERS":
		ret, e := getSkypeMessageSetList("CHAT", "ACTIVEMEMBERS", s, " ")
		if e != nil {
			return e
		}
		self.Activemembers = ret
	case "DESCRIPTION":
		ret, e := getSkypeMessageSetComplexString("CHAT", "DESCRIPTION", s)
		if e != nil {
			return e
		}
		self.Description = ret
	case "TIMESTAMP":
		ret, e := getSkypeMessageSetSimpleString("CHAT", "TIMESTAMP", s)
		if e != nil {
			return e
		}
		self.Timestamp = ret
	case "CHATMESSAGES":
		ret, e := getSkypeMessageSetList("CHAT", "CHATMESSAGES", s, ", ")
		if e != nil {
			return e
		}
		self.Chatmessages = ret
	case "TOPIC":
		ret, e := getSkypeMessageSetComplexString("CHAT", "TOPIC", s)
		if e != nil {
			return e
		}
		self.Topic = ret
	case "ROLE":
		ret, e := getSkypeMessageSetSimpleString("CHAT", "ROLE", s)
		if e != nil {
			return e
		}
		self.Role = ret
	case "BLOB":
		ret, e := getSkypeMessageSetSimpleString("CHAT", "BLOB", s)
		if e != nil {
			return e
		}
		self.Blob = ret
	case "MEMBERS":
		ret, e := getSkypeMessageSetList("CHAT", "MEMBERS", s, " ")
		if e != nil {
			return e
		}
		self.Members = ret
	case "MYROLE":
		ret, e := getSkypeMessageSetSimpleString("CHAT", "MYROLE", s)
		if e != nil {
			return e
		}
		self.Myrole = ret
	default:
		return ParserForFieldDoesNotExist
	}
	return nil
}


type Chatmember struct {
	Id       string
	Role     string
	IsActive string
	Identity string  // TODO(wb): what is this?
	Chatname string  // may contain a space
}

func (self *Chatmember) getFetchAllFieldsCommands() ([]string, error) {
	return []string{"GET CHATMEMBER " + self.Id + " ROLE",
		"GET CHATMEMBER " + self.Id + " IS_ACTIVE",
		"GET CHATMEMBER " + self.Id + " IDENTITY",
		"GET CHATMEMBER " + self.Id + " CHATNAME"}, nil
}

func (self *Chatmember) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "CHATMEMBER "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return UnexpectedNumberOfFieldsError
	}
	
	switch field_to_set {
	case "ROLE":
		ret, e := getSkypeMessageSetSimpleString("CHATMEMBER", "ROLE", s)
		if e != nil {
			return e
		}
		self.Role = ret
	case "IS_ACTIVE":
		ret, e := getSkypeMessageSetSimpleString("CHATMEMBER", "IS_ACTIVE", s)
		if e != nil {
			return e
		}
		self.IsActive = ret
	case "IDENTITY":
		ret, e := getSkypeMessageSetSimpleString("CHATMEMBER", "IDENTITY", s)
		if e != nil {
			return e
		}
		self.Identity = ret
	case "CHATNAME":
		ret, e := getSkypeMessageSetComplexString("CHATMEMBER", "CHATNAME", s)
		if e != nil {
			return e
		}
		self.Chatname = ret
	default:
		return ParserForFieldDoesNotExist
	}
	return nil
}

type Chatmessage struct {
	Id              string
	Body            string  // may contain a space
	Status          string
	EditedTimestamp string
	EditedBy        string
	Users           []string  // ,-separated list
	Timestamp       string
	FromHandle      string
	Chatname        string  // simple string
	IsEditable      string
	Leavereason     string
	FromDispname    string  // may contain a space
	ChatmessageType string
}

func (self *Chatmessage) getFetchAllFieldsCommands() ([]string, error) {
	return []string{"GET CHATMESSAGE " + self.Id + " BODY",
		"GET CHATMESSAGE " + self.Id + " STATUS",
		"GET CHATMESSAGE " + self.Id + " EDITED_TIMESTAMP",
		"GET CHATMESSAGE " + self.Id + " EDITED_BY",
		"GET CHATMESSAGE " + self.Id + " USERS",
		"GET CHATMESSAGE " + self.Id + " TIMESTAMP",
		"GET CHATMESSAGE " + self.Id + " FROM_HANDLE",
		"GET CHATMESSAGE " + self.Id + " CHATNAME",
		"GET CHATMESSAGE " + self.Id + " IS_EDITABLE",
		"GET CHATMESSAGE " + self.Id + " LEAVEREASON",
		"GET CHATMESSAGE " + self.Id + " FROM_DISPNAME",
		"GET CHATMESSAGE " + self.Id + " TYPE"}, nil
}

func (self *Chatmessage) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "CHATMESSAGE "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return UnexpectedNumberOfFieldsError
	}

	switch field_to_set {
	case "BODY":
		ret, e := getSkypeMessageSetComplexString("CHATMESSAGE", "BODY", s)
		if e != nil {
			return e
		}
		self.Body = ret
	case "STATUS":
		ret, e := getSkypeMessageSetSimpleString("CHATMESSAGE", "STATUS", s)
		if e != nil {
			return e
		}
		self.Status = ret
	case "EDITED_TIMESTAMP":
		ret, e := getSkypeMessageSetSimpleString("CHATMESSAGE", "EDITED_TIMESTAMP", s)
		if e != nil {
			return e
		}
		self.EditedTimestamp = ret
	case "EDITED_BY":
		ret, e := getSkypeMessageSetSimpleString("CHATMESSAGE", "EDITED_BY", s)
		if e != nil {
			return e
		}
		self.EditedBy = ret
	case "USERS":
		ret, e := getSkypeMessageSetList("CHATMESSAGE", "USERS", s, ", ")
		if e != nil {
			return e
		}
		self.Users = ret
	case "TIMESTAMP":
		ret, e := getSkypeMessageSetSimpleString("CHATMESSAGE", "TIMESTAMP", s)
		if e != nil {
			return e
		}
		self.Timestamp = ret
	case "FROM_HANDLE":
		ret, e := getSkypeMessageSetSimpleString("CHATMESSAGE", "FROM_HANDLE", s)
		if e != nil {
			return e
		}
		self.FromHandle = ret
	case "CHATNAME":
		ret, e := getSkypeMessageSetSimpleString("CHATMESSAGE", "CHATNAME", s)
		if e != nil {
			return e
		}
		self.Chatname = ret
	case "IS_EDITABLE":
		ret, e := getSkypeMessageSetSimpleString("CHATMESSAGE", "IS_EDITABLE", s)
		if e != nil {
			return e
		}
		self.IsEditable = ret
	case "LEAVEREASON":
		ret, e := getSkypeMessageSetSimpleString("CHATMESSAGE", "LEAVEREASON", s)
		if e != nil {
			return e
		}
		self.Leavereason = ret
	case "FROM_DISPNAME":
		ret, e := getSkypeMessageSetComplexString("CHATMESSAGE", "FROM_DISPNAME", s)
		if e != nil {
			return e
		}
		self.FromDispname = ret
	case "TYPE":
		ret, e := getSkypeMessageSetSimpleString("CHATMESSAGE", "TYPE", s)
		if e != nil {
			return e
		}
		self.ChatmessageType = ret
	default:
		return ParserForFieldDoesNotExist
	}
	return nil
}

type Filetransfer struct {
	Id               string
	Finishtime       string
	Status           string
	PartnerHandle    string
	Filepath         string
	Bytespersecond   string
	Filesize         string
	Starttime        string
	PartnerDispname  string  // probably can contain a space
	FiletransferType string
	Bytestransferred string
	Failurereason    string
}

func (self *Filetransfer) getFetchAllFieldsCommands() ([]string, error) {
	return []string{"GET FILETRANSFER " + self.Id + " FINISHTIME",
		"GET FILETRANSFER " + self.Id + " STATUS",
		"GET FILETRANSFER " + self.Id + " PARTNER_HANDLE",
		"GET FILETRANSFER " + self.Id + " FILEPATH",
		"GET FILETRANSFER " + self.Id + " BYTESPERSECOND",
		"GET FILETRANSFER " + self.Id + " FILESIZE",
		"GET FILETRANSFER " + self.Id + " STARTTIME",
		"GET FILETRANSFER " + self.Id + " PARTNER_DISPNAME",
		"GET FILETRANSFER " + self.Id + " TYPE",
		"GET FILETRANSFER " + self.Id + " BYTESTRANSFERRED",
		"GET FILETRANSFER " + self.Id + " FAILUREREASON"}, nil
}

func (self *Filetransfer) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "FILETRANSFER "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return UnexpectedNumberOfFieldsError
	}

	switch field_to_set {
	case "FINISHTIME":
		ret, e := getSkypeMessageSetSimpleString("FILETRANSFER", "FINISHTIME", s)
		if e != nil {
			return e
		}
		self.Finishtime = ret
	case "STATUS":
		ret, e := getSkypeMessageSetSimpleString("FILETRANSFER", "STATUS", s)
		if e != nil {
			return e
		}
		self.Status = ret
	case "PARTNER_HANDLE":
		ret, e := getSkypeMessageSetSimpleString("FILETRANSFER", "PARTNER_HANDLE", s)
		if e != nil {
			return e
		}
		self.PartnerHandle = ret
	case "FILEPATH":
		ret, e := getSkypeMessageSetSimpleString("FILETRANSFER", "FILEPATH", s)
		if e != nil {
			return e
		}
		self.Filepath = ret
	case "BYTESPERSECOND":
		ret, e := getSkypeMessageSetSimpleString("FILETRANSFER", "BYTESPERSECOND", s)
		if e != nil {
			return e
		}
		self.Bytespersecond = ret
	case "FILESIZE":
		ret, e := getSkypeMessageSetSimpleString("FILETRANSFER", "FILESIZE", s)
		if e != nil {
			return e
		}
		self.Filesize = ret
	case "STARTTIME":
		ret, e := getSkypeMessageSetSimpleString("FILETRANSFER", "STARTTIME", s)
		if e != nil {
			return e
		}
		self.Starttime = ret
	case "PARTNER_DISPNAME":
		ret, e := getSkypeMessageSetComplexString("FILETRANSFER", "PARTNER_DISPNAME", s)
		if e != nil {
			return e
		}
		self.PartnerDispname = ret
	case "TYPE":
		ret, e := getSkypeMessageSetSimpleString("FILETRANSFER", "TYPE", s)
		if e != nil {
			return e
		}
		self.FiletransferType = ret
	case "BYTESTRANSFERRED":
		ret, e := getSkypeMessageSetSimpleString("FILETRANSFER", "BYTESTRANSFERRED", s)
		if e != nil {
			return e
		}
		self.Bytestransferred = ret
	case "FAILUREREASON":
		ret, e := getSkypeMessageSetSimpleString("FILETRANSFER", "FAILUREREASON", s)
		if e != nil {
			return e
		}
		self.Failurereason = ret
	default:
		return ParserForFieldDoesNotExist
	}
	return nil
}

type Group struct {
	Id            string
	Displayname   string
	Users         []string
	Expanded      string
	CustomGroupId string
	Visible       string
	GroupType     string
	NrofUsers     string
}

func (self *Group) getFetchAllFieldsCommands() ([]string, error) {
	return []string{"GET GROUP " + self.Id + " DISPLAYNAME",
		"GET GROUP " + self.Id + " USERS",
		"GET GROUP " + self.Id + " EXPANDED",
		"GET GROUP " + self.Id + " CUSTOM_GROUP_ID",
		"GET GROUP " + self.Id + " VISIBLE",
		"GET GROUP " + self.Id + " TYPE",
		"GET GROUP " + self.Id + " NROFUSERS"}, nil
}

func (self *Group) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "GROUP "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return UnexpectedNumberOfFieldsError
	}

	switch field_to_set {
	case "USERS":
		ret, e := getSkypeMessageSetList("GROUP", "USERS", s, ", ")
		if e != nil {
			return e
		}
		self.Users = ret
	case "EXPANDED":
		ret, e := getSkypeMessageSetComplexString("GROUP", "EXPANDED", s)
		if e != nil {
			return e
		}
		self.Expanded = ret
	case "CUSTOM_GROUP_ID":
		ret, e := getSkypeMessageSetComplexString("GROUP", "CUSTOM_GROUP_ID", s)
		if e != nil {
			return e
		}
		self.CustomGroupId = ret
	case "VISIBLE":
		ret, e := getSkypeMessageSetSimpleString("GROUP", "VISIBLE", s)
		if e != nil {
			return e
		}
		self.Visible = ret
	case "TYPE":
		ret, e := getSkypeMessageSetSimpleString("GROUP", "TYPE", s)
		if e != nil {
			return e
		}
		self.GroupType = ret
	case "NROFUSERS":
		ret, e := getSkypeMessageSetSimpleString("GROUP", "NROFUSERS", s)
		if e != nil {
			return e
		}
		self.NrofUsers = ret
	default:
		return ParserForFieldDoesNotExist
	}
	return nil
}

type User struct {
	Id                  string
	Province            string  // may have a space?
	About               string  // may have a space
	PhoneOffice         string  // may have a space?
	Country             string  // may have a space
	Birthday            string
	IsCfActive          string
	Timezone            string
	Speeddial           string
	Displayname         string  // not sure what this looks like
	Language            string  // may have a space
	Isblocked           string
	Onlinestatus        string
	Sex                 string
	CanLeaveVm          string
	MoodText            string  // may have a space
	Homepage            string  // may have a space
	Aliases             string  // don't know what this looks like
	IsVideoCapable      string
	Lastonlinetimestamp string
	Buddystatus         string
	Hascallequipment    string
	NrofAuthedBuddies   string
	Receivedauthrequest string
	City                string  // may have a space
	Isauthorized        string
	IsVoicemailCapable  string
	PhoneHome           string  // may have a space?
	Avatar              string  // not sure what this looks like
	RichMoodText        string  // may have a space
	Fullname            string  // may have a space
	PhoneMobile         string  // may have a space
}

func (self *User) getFetchAllFieldsCommands() ([]string, error) {
	return []string{"GET USER " + self.Id + " PROVINCE",
		"GET USER " + self.Id + " ABOUT",
		"GET USER " + self.Id + " PHONE_OFFICE",
		"GET USER " + self.Id + " COUNTRY",
		"GET USER " + self.Id + " BIRTHDAY",
		"GET USER " + self.Id + " IS_CF_ACTIVE",
		"GET USER " + self.Id + " TIMEZONE",
		"GET USER " + self.Id + " SPEEDDIAL",
		"GET USER " + self.Id + " DISPLAYNAME",
		"GET USER " + self.Id + " LANGUAGE",
		"GET USER " + self.Id + " ISBLOCKED",
		"GET USER " + self.Id + " ONLINESTATUS",
		"GET USER " + self.Id + " SEX",
		"GET USER " + self.Id + " CAN_LEAVE_VM",
		"GET USER " + self.Id + " MOOD_TEXT",
		"GET USER " + self.Id + " HOMEPAGE",
		"GET USER " + self.Id + " ALIASES",
		"GET USER " + self.Id + " IS_VIDEO_CAPABLE",
		"GET USER " + self.Id + " LASTONLINETIMESTAMP",
		"GET USER " + self.Id + " BUDDYSTATUS",
		"GET USER " + self.Id + " HASCALLEQUIPMENT",
		"GET USER " + self.Id + " NROF_AUTHED_BUDDIES",
		"GET USER " + self.Id + " RECEIVEDAUTHREQUEST",
		"GET USER " + self.Id + " CITY",
		"GET USER " + self.Id + " ISAUTHORIZED",
		"GET USER " + self.Id + " IS_VOICEMAIL_CAPABLE",
		"GET USER " + self.Id + " PHONE_HOME",
		"GET USER " + self.Id + " AVATAR",
		"GET USER " + self.Id + " RICH_MOOD_TEXT",
		"GET USER " + self.Id + " FULLNAME",
		"GET USER " + self.Id + " PHONE_MOBILE"}, nil
}

func (self *User) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "USER "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return UnexpectedNumberOfFieldsError
	}

	switch field_to_set {
	case "PROVINCE":
		ret, e := getSkypeMessageSetComplexString("USER", "PROVINCE", s)
		if e != nil {
			return e
		}
		self.Province = ret
	case "ABOUT":
		ret, e := getSkypeMessageSetComplexString("USER", "ABOUT", s)
		if e != nil {
			return e
		}
		self.About = ret
	case "PHONE_OFFICE":
		ret, e := getSkypeMessageSetComplexString("USER", "PHONE_OFFICE", s)
		if e != nil {
			return e
		}
		self.PhoneOffice = ret
	case "COUNTRY":
		ret, e := getSkypeMessageSetComplexString("USER", "COUNTRY", s)
		if e != nil {
			return e
		}
		self.Country = ret
	case "BIRTHDAY":
		ret, e := getSkypeMessageSetSimpleString("USER", "BIRTHDAY", s)
		if e != nil {
			return e
		}
		self.Birthday = ret
	case "IS_CF_ACTIVE":
		ret, e := getSkypeMessageSetSimpleString("USER", "IS_CF_ACTIVE", s)
		if e != nil {
			return e
		}
		self.IsCfActive = ret
	case "TIMEZONE":
		ret, e := getSkypeMessageSetSimpleString("USER", "TIMEZONE", s)
		if e != nil {
			return e
		}
		self.Timezone = ret
	case "SPEEDDIAL":
		ret, e := getSkypeMessageSetSimpleString("USER", "SPEEDDIAL", s)
		if e != nil {
			return e
		}
		self.Speeddial = ret
	case "DISPLAYNAME":
		ret, e := getSkypeMessageSetComplexString("USER", "DISPLAYNAME", s)
		if e != nil {
			return e
		}
		self.Displayname = ret
	case "LANGUAGE":
		ret, e := getSkypeMessageSetComplexString("USER", "LANGUAGE", s)
		if e != nil {
			return e
		}
		self.Language = ret
	case "ISBLOCKED":
		ret, e := getSkypeMessageSetSimpleString("USER", "ISBLOCKED", s)
		if e != nil {
			return e
		}
		self.Isblocked = ret
	case "ONLINESTATUS":
		ret, e := getSkypeMessageSetSimpleString("USER", "ONLINESTATUS", s)
		if e != nil {
			return e
		}
		self.Onlinestatus = ret
	case "SEX":
		ret, e := getSkypeMessageSetSimpleString("USER", "SEX", s)
		if e != nil {
			return e
		}
		self.Sex = ret
	case "CAN_LEAVE_VM":
		ret, e := getSkypeMessageSetSimpleString("USER", "CAN_LEAVE_VM", s)
		if e != nil {
			return e
		}
		self.CanLeaveVm = ret
	case "MOOD_TEXT":
		ret, e := getSkypeMessageSetComplexString("USER", "MOOD_TEXT", s)
		if e != nil {
			return e
		}
		self.MoodText = ret
	case "HOMEPAGE":
		ret, e := getSkypeMessageSetComplexString("USER", "HOMEPAGE", s)
		if e != nil {
			return e
		}
		self.Homepage = ret
	case "ALIASES":
		ret, e := getSkypeMessageSetSimpleString("USER", "ALIASES", s)
		if e != nil {
			return e
		}
		self.Aliases = ret
	case "IS_VIDEO_CAPABLE":
		ret, e := getSkypeMessageSetSimpleString("USER", "IS_VIDEO_CAPABLE", s)
		if e != nil {
			return e
		}
		self.IsVideoCapable = ret
	case "LASTONLINETIMESTAMP":
		ret, e := getSkypeMessageSetSimpleString("USER", "LASTONLINETIMESTAMP", s)
		if e != nil {
			return e
		}
		self.Lastonlinetimestamp = ret
	case "BUDDYSTATUS":
		ret, e := getSkypeMessageSetSimpleString("USER", "BUDDYSTATUS", s)
		if e != nil {
			return e
		}
		self.Buddystatus = ret
	case "HASCALLEQUIPMENT":
		ret, e := getSkypeMessageSetSimpleString("USER", "HASCALLEQUIPMENT", s)
		if e != nil {
			return e
		}
		self.Hascallequipment = ret
	case "NROF_AUTHED_BUDDIES":
		ret, e := getSkypeMessageSetSimpleString("USER", "NROF_AUTHED_BUDDIES", s)
		if e != nil {
			return e
		}
		self.NrofAuthedBuddies = ret
	case "RECEIVEDAUTHREQUEST":
		ret, e := getSkypeMessageSetSimpleString("USER", "RECEIVEDAUTHREQUEST", s)
		if e != nil {
			return e
		}
		self.Receivedauthrequest = ret
	case "CITY":
		ret, e := getSkypeMessageSetComplexString("USER", "CITY", s)
		if e != nil {
			return e
		}
		self.City = ret
	case "ISAUTHORIZED":
		ret, e := getSkypeMessageSetSimpleString("USER", "ISAUTHORIZED", s)
		if e != nil {
			return e
		}
		self.Isauthorized = ret
	case "IS_VOICEMAIL_CAPABLE":
		ret, e := getSkypeMessageSetSimpleString("USER", "IS_VOICEMAIL_CAPABLE", s)
		if e != nil {
			return e
		}
		self.IsVoicemailCapable = ret
	case "PHONE_HOME":
		ret, e := getSkypeMessageSetComplexString("USER", "PHONE_HOME", s)
		if e != nil {
			return e
		}
		self.PhoneHome = ret
	case "AVATAR":
		ret, e := getSkypeMessageSetComplexString("USER", "AVATAR", s)
		if e != nil {
			return e
		}
		self.Avatar = ret
	case "RICH_MOOD_TEXT":
		ret, e := getSkypeMessageSetComplexString("USER", "RICH_MOOD_TEXT", s)
		if e != nil {
			return e
		}
		self.RichMoodText = ret
	case "FULLNAME":
		ret, e := getSkypeMessageSetComplexString("USER", "FULLNAME", s)
		if e != nil {
			return e
		}
		self.Fullname = ret
	case "PHONE_MOBILE":
		ret, e := getSkypeMessageSetComplexString("USER", "PHONE_MOBILE", s)
		if e != nil {
			return e
		}
		self.PhoneMobile = ret
	default:
		return ParserForFieldDoesNotExist
	}
	return nil
}

