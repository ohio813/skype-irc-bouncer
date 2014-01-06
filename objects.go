package gsi

import "fmt"
import "strings"

type UnexpectedNumberOfFieldsError struct {
	s string
}

func (e *UnexpectedNumberOfFieldsError) Error() string {
	return "Unexpected number of fields, line: '" + e.s + "'"
}

type ParserForFieldDoesNotExist struct {
	field string
}

func (e *ParserForFieldDoesNotExist) Error() string {
	return "Parser for field '" + e.field + "' does not exist"
}

type Aec struct {
	Id string
}

func (self *Aec) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *Aec) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "AEC "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Agc struct {
	Id string
}

func (self *Agc) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *Agc) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "AGC "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Application struct {
	Id          string
	Connecting  string
	Sending     string
	Streams     string
	Connectable string
	Received    string
}

func (self *Application) parseSetConnecting(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "APPLICATION %s CONNECTING %s", &__, &self.Connecting); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Application) parseSetSending(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "APPLICATION %s SENDING %s", &__, &self.Sending); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Application) parseSetStreams(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "APPLICATION %s STREAMS %s", &__, &self.Streams); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Application) parseSetConnectable(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "APPLICATION %s CONNECTABLE %s", &__, &self.Connectable); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Application) parseSetReceived(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "APPLICATION %s RECEIVED %s", &__, &self.Received); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Application) getFetchAllFieldsCommands() ([]string, error) {
	return []string{"GET APPLICATION " + self.Id + " CONNECTING",
		"GET APPLICATION " + self.Id + " SENDING",
		"GET APPLICATION " + self.Id + " STREAMS",
		"GET APPLICATION " + self.Id + " CONNECTABLE",
		"GET APPLICATION " + self.Id + " RECEIVED"}, nil
}

func (self *Application) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "APPLICATION "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	if "CONNECTING" == field_to_set {
		return self.parseSetConnecting(s)
	}
	if "SENDING" == field_to_set {
		return self.parseSetSending(s)
	}
	if "STREAMS" == field_to_set {
		return self.parseSetStreams(s)
	}
	if "CONNECTABLE" == field_to_set {
		return self.parseSetConnectable(s)
	}
	if "RECEIVED" == field_to_set {
		return self.parseSetReceived(s)
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type AudioIn struct {
	Id string
}

func (self *AudioIn) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *AudioIn) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "AUDIO_IN "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type AudioOut struct {
	Id string
}

func (self *AudioOut) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *AudioOut) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "AUDIO_OUT "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Autoaway struct {
	Id string
}

func (self *Autoaway) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *Autoaway) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "AUTOAWAY "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Avatar struct {
	Id string
}

func (self *Avatar) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *Avatar) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "AVATAR "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Call struct {
	Id                    string
	VideoReceiveStatus    string
	VideoStatus           string
	CanTransfer           string
	RateCurrency          string
	RatePrecision         string
	Rate                  string
	TransferStatus        string
	Seen                  string
	Duration              string
	PstnStatus            string
	VmAllowedDuration     string
	Failurereason         string
	VideoSendStatus       string
	TargetIdentity        string
	TransferActive        string
	TransferredBy         string
	Status                string
	VaaInputStatus        string
	PartnerHandle         string
	ForwardedBy           string
	ConfParticipantsCount string
	ConfParticipant       string
	ConfId                string
	TransferredTo         string
	Input                 string
	Timestamp             string
	CaptureMic            string
	PstnNumber            string
	PartnerDispname       string
	Output                string
	VmDuration            string
	CallType              string
	Subject               string
}

func (self *Call) parseSetVideoReceiveStatus(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s VIDEO_RECEIVE_STATUS %s", &__, &self.VideoReceiveStatus); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetVideoStatus(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s VIDEO_STATUS %s", &__, &self.VideoStatus); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetCanTransfer(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s CAN_TRANSFER %s", &__, &self.CanTransfer); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetRateCurrency(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s RATE_CURRENCY %s", &__, &self.RateCurrency); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetRatePrecision(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s RATE_PRECISION %s", &__, &self.RatePrecision); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetRate(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s RATE %s", &__, &self.Rate); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetTransferStatus(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s TRANSFER_STATUS %s", &__, &self.TransferStatus); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetSeen(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s SEEN %s", &__, &self.Seen); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetDuration(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s DURATION %s", &__, &self.Duration); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetPstnStatus(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s PSTN_STATUS %s", &__, &self.PstnStatus); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetVmAllowedDuration(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s VM_ALLOWED_DURATION %s", &__, &self.VmAllowedDuration); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetFailurereason(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s FAILUREREASON %s", &__, &self.Failurereason); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetVideoSendStatus(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s VIDEO_SEND_STATUS %s", &__, &self.VideoSendStatus); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetTargetIdentity(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s TARGET_IDENTITY %s", &__, &self.TargetIdentity); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetTransferActive(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s TRANSFER_ACTIVE %s", &__, &self.TransferActive); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetTransferredBy(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s TRANSFERRED_BY %s", &__, &self.TransferredBy); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetStatus(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s STATUS %s", &__, &self.Status); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetVaaInputStatus(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s VAA_INPUT_STATUS %s", &__, &self.VaaInputStatus); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetPartnerHandle(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s PARTNER_HANDLE %s", &__, &self.PartnerHandle); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetForwardedBy(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s FORWARDED_BY %s", &__, &self.ForwardedBy); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetConfParticipantsCount(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s CONF_PARTICIPANTS_COUNT %s", &__, &self.ConfParticipantsCount); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetConfParticipant(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s CONF_PARTICIPANT %s", &__, &self.ConfParticipant); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetConfId(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s CONF_ID %s", &__, &self.ConfId); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetTransferredTo(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s TRANSFERRED_TO %s", &__, &self.TransferredTo); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetInput(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s INPUT %s", &__, &self.Input); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetTimestamp(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s TIMESTAMP %s", &__, &self.Timestamp); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetCaptureMic(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s CAPTURE_MIC %s", &__, &self.CaptureMic); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetPstnNumber(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s PSTN_NUMBER %s", &__, &self.PstnNumber); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetPartnerDispname(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s PARTNER_DISPNAME %s", &__, &self.PartnerDispname); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetOutput(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s OUTPUT %s", &__, &self.Output); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetVmDuration(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s VM_DURATION %s", &__, &self.VmDuration); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetCallType(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s TYPE %s", &__, &self.CallType); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) parseSetSubject(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CALL %s SUBJECT %s", &__, &self.Subject); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Call) getFetchAllFieldsCommands() ([]string, error) {
	return []string{"GET CALL " + self.Id + " VIDEO_RECEIVE_STATUS",
		"GET CALL " + self.Id + " VIDEO_STATUS",
		"GET CALL " + self.Id + " CAN_TRANSFER",
		"GET CALL " + self.Id + " RATE_CURRENCY",
		"GET CALL " + self.Id + " RATE_PRECISION",
		"GET CALL " + self.Id + " RATE",
		"GET CALL " + self.Id + " TRANSFER_STATUS",
		"GET CALL " + self.Id + " SEEN",
		"GET CALL " + self.Id + " DURATION",
		"GET CALL " + self.Id + " PSTN_STATUS",
		"GET CALL " + self.Id + " VM_ALLOWED_DURATION",
		"GET CALL " + self.Id + " FAILUREREASON",
		"GET CALL " + self.Id + " VIDEO_SEND_STATUS",
		"GET CALL " + self.Id + " TARGET_IDENTITY",
		"GET CALL " + self.Id + " TRANSFER_ACTIVE",
		"GET CALL " + self.Id + " TRANSFERRED_BY",
		"GET CALL " + self.Id + " STATUS",
		"GET CALL " + self.Id + " VAA_INPUT_STATUS",
		"GET CALL " + self.Id + " PARTNER_HANDLE",
		"GET CALL " + self.Id + " FORWARDED_BY",
		"GET CALL " + self.Id + " CONF_PARTICIPANTS_COUNT",
		"GET CALL " + self.Id + " CONF_PARTICIPANT",
		"GET CALL " + self.Id + " CONF_ID",
		"GET CALL " + self.Id + " TRANSFERRED_TO",
		"GET CALL " + self.Id + " INPUT",
		"GET CALL " + self.Id + " TIMESTAMP",
		"GET CALL " + self.Id + " CAPTURE_MIC",
		"GET CALL " + self.Id + " PSTN_NUMBER",
		"GET CALL " + self.Id + " PARTNER_DISPNAME",
		"GET CALL " + self.Id + " OUTPUT",
		"GET CALL " + self.Id + " VM_DURATION",
		"GET CALL " + self.Id + " TYPE",
		"GET CALL " + self.Id + " SUBJECT"}, nil
}

func (self *Call) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "CALL "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	if "VIDEO_RECEIVE_STATUS" == field_to_set {
		return self.parseSetVideoReceiveStatus(s)
	}
	if "VIDEO_STATUS" == field_to_set {
		return self.parseSetVideoStatus(s)
	}
	if "CAN_TRANSFER" == field_to_set {
		return self.parseSetCanTransfer(s)
	}
	if "RATE_CURRENCY" == field_to_set {
		return self.parseSetRateCurrency(s)
	}
	if "RATE_PRECISION" == field_to_set {
		return self.parseSetRatePrecision(s)
	}
	if "RATE" == field_to_set {
		return self.parseSetRate(s)
	}
	if "TRANSFER_STATUS" == field_to_set {
		return self.parseSetTransferStatus(s)
	}
	if "SEEN" == field_to_set {
		return self.parseSetSeen(s)
	}
	if "DURATION" == field_to_set {
		return self.parseSetDuration(s)
	}
	if "PSTN_STATUS" == field_to_set {
		return self.parseSetPstnStatus(s)
	}
	if "VM_ALLOWED_DURATION" == field_to_set {
		return self.parseSetVmAllowedDuration(s)
	}
	if "FAILUREREASON" == field_to_set {
		return self.parseSetFailurereason(s)
	}
	if "VIDEO_SEND_STATUS" == field_to_set {
		return self.parseSetVideoSendStatus(s)
	}
	if "TARGET_IDENTITY" == field_to_set {
		return self.parseSetTargetIdentity(s)
	}
	if "TRANSFER_ACTIVE" == field_to_set {
		return self.parseSetTransferActive(s)
	}
	if "TRANSFERRED_BY" == field_to_set {
		return self.parseSetTransferredBy(s)
	}
	if "STATUS" == field_to_set {
		return self.parseSetStatus(s)
	}
	if "VAA_INPUT_STATUS" == field_to_set {
		return self.parseSetVaaInputStatus(s)
	}
	if "PARTNER_HANDLE" == field_to_set {
		return self.parseSetPartnerHandle(s)
	}
	if "FORWARDED_BY" == field_to_set {
		return self.parseSetForwardedBy(s)
	}
	if "CONF_PARTICIPANTS_COUNT" == field_to_set {
		return self.parseSetConfParticipantsCount(s)
	}
	if "CONF_PARTICIPANT" == field_to_set {
		return self.parseSetConfParticipant(s)
	}
	if "CONF_ID" == field_to_set {
		return self.parseSetConfId(s)
	}
	if "TRANSFERRED_TO" == field_to_set {
		return self.parseSetTransferredTo(s)
	}
	if "INPUT" == field_to_set {
		return self.parseSetInput(s)
	}
	if "TIMESTAMP" == field_to_set {
		return self.parseSetTimestamp(s)
	}
	if "CAPTURE_MIC" == field_to_set {
		return self.parseSetCaptureMic(s)
	}
	if "PSTN_NUMBER" == field_to_set {
		return self.parseSetPstnNumber(s)
	}
	if "PARTNER_DISPNAME" == field_to_set {
		return self.parseSetPartnerDispname(s)
	}
	if "OUTPUT" == field_to_set {
		return self.parseSetOutput(s)
	}
	if "VM_DURATION" == field_to_set {
		return self.parseSetVmDuration(s)
	}
	if "TYPE" == field_to_set {
		return self.parseSetCallType(s)
	}
	if "SUBJECT" == field_to_set {
		return self.parseSetSubject(s)
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
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

func (self *Chat) parseSetDialogPartner(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHAT %s DIALOG_PARTNER %s", &__, &self.DialogPartner); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chat) parseSetPasswordhint(s string) error {
	var id string
	if n, e := fmt.Sscanf(s, "CHAT %s PASSWORDHINT ", &id); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}

	remainder := s[len(fmt.Sprintf("CHAT %s PASSWORDHINT ", id)):]
	self.Passwordhint = strings.TrimRight(remainder, "\r\n\t ")

	return nil
}

func (self *Chat) parseSetOptions(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHAT %s OPTIONS %s", &__, &self.Options); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chat) parseSetApplicants(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHAT %s APPLICANTS %s", &__, &self.Applicants); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chat) parseSetBookmarked(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHAT %s BOOKMARKED %s", &__, &self.Bookmarked); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chat) parseSetRecentchatmessages(s string) error {
	var id string
	if n, e := fmt.Sscanf(s, "CHAT %s RECENTCHATMESSAGES ", &id); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}

	remainder := s[len(fmt.Sprintf("CHAT %s RECENTCHATMESSAGES ", id)):]
	self.Recentchatmessages = strings.Split(remainder, ", ")	

	return nil
}

func (self *Chat) parseSetChatname(s string) error {
	var id string
	if n, e := fmt.Sscanf(s, "CHAT %s CHATNAME ", &id); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}

	remainder := s[len(fmt.Sprintf("CHAT %s CHATNAME ", id)):]
	self.Chatname = strings.TrimRight(remainder, "\r\n\t ")

	return nil
}

func (self *Chat) parseSetAdder(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHAT %s ADDER %s", &__, &self.Adder); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chat) parseSetActivityTimestamp(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHAT %s ACTIVITY_TIMESTAMP %s", &__, &self.ActivityTimestamp); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chat) parseSetPosters(s string) error {
	var id string
	if n, e := fmt.Sscanf(s, "CHAT %s POSTERS ", &id); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}

	remainder := s[len(fmt.Sprintf("CHAT %s POSTERS ", id)):]
	self.Posters = strings.Split(remainder, ", ")	

	return nil
}

func (self *Chat) parseSetStatus(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHAT %s STATUS %s", &__, &self.Status); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chat) parseSetGuidelines(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHAT %s GUIDELINES %s", &__, &self.Guidelines); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chat) parseSetTopicxml(s string) error {
	var id string
	if n, e := fmt.Sscanf(s, "CHAT %s TOPICXML ", &id); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}

	remainder := s[len(fmt.Sprintf("CHAT %s TOPICXML ", id)):]
	self.Topicxml = strings.TrimRight(remainder, "\r\n\t ")

	return nil
}

func (self *Chat) parseSetMystatus(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHAT %s MYSTATUS %s", &__, &self.Mystatus); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chat) parseSetMemberobjects(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHAT %s MEMBEROBJECTS %s", &__, &self.Memberobjects); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chat) parseSetFriendlyname(s string) error {
	var id string
	if n, e := fmt.Sscanf(s, "CHAT %s FRIENDLYNAME ", &id); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}

	remainder := s[len(fmt.Sprintf("CHAT %s FRIENDLYNAME ", id)):]
	self.Friendlyname = strings.TrimRight(remainder, "\r\n\t ")

	return nil
}

func (self *Chat) parseSetActivemembers(s string) error {
	var id string
	if n, e := fmt.Sscanf(s, "CHAT %s ACTIVEMEMBERS ", &id); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}

	remainder := s[len(fmt.Sprintf("CHAT %s ACTIVEMEMBERS ", id)):]
	self.Activemembers = strings.Split(remainder, " ")	

	return nil
}

func (self *Chat) parseSetDescription(s string) error {
	var id string
	if n, e := fmt.Sscanf(s, "CHAT %s DESCRIPTION ", &id); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}

	remainder := s[len(fmt.Sprintf("CHAT %s DESCRIPTION ", id)):]
	self.Description = strings.TrimRight(remainder, "\r\n\t ")

	return nil
}

func (self *Chat) parseSetTimestamp(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHAT %s TIMESTAMP %s", &__, &self.Timestamp); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chat) parseSetChatmessages(s string) error {
	var id string
	if n, e := fmt.Sscanf(s, "CHAT %s CHATMESSAGES ", &id); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}

	remainder := s[len(fmt.Sprintf("CHAT %s CHATMESSAGES ", id)):]
	self.Chatmessages = strings.Split(remainder, ", ")	

	return nil
}

func (self *Chat) parseSetTopic(s string) error {
	var id string
	if n, e := fmt.Sscanf(s, "CHAT %s TOPIC ", &id); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}

	remainder := s[len(fmt.Sprintf("CHAT %s TOPIC ", id)):]
	self.Topic = strings.TrimRight(remainder, "\r\n\t ")

	return nil
}

func (self *Chat) parseSetRole(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHAT %s ROLE %s", &__, &self.Role); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chat) parseSetBlob(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHAT %s BLOB %s", &__, &self.Blob); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chat) parseSetMembers(s string) error {
	var id string
	if n, e := fmt.Sscanf(s, "CHAT %s MEMBERS ", &id); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}

	remainder := s[len(fmt.Sprintf("CHAT %s MEMBERS ", id)):]
	self.Members = strings.Split(remainder, " ")	

	return nil
}

func (self *Chat) parseSetMyrole(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHAT %s MYROLE %s", &__, &self.Myrole); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
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
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	if "DIALOG_PARTNER" == field_to_set {
		return self.parseSetDialogPartner(s)
	}
	if "PASSWORDHINT" == field_to_set {
		return self.parseSetPasswordhint(s)
	}
	if "OPTIONS" == field_to_set {
		return self.parseSetOptions(s)
	}
	if "APPLICANTS" == field_to_set {
		return self.parseSetApplicants(s)
	}
	if "BOOKMARKED" == field_to_set {
		return self.parseSetBookmarked(s)
	}
	if "RECENTCHATMESSAGES" == field_to_set {
		return self.parseSetRecentchatmessages(s)
	}
	if "CHATNAME" == field_to_set {
		return self.parseSetChatname(s)
	}
	if "ADDER" == field_to_set {
		return self.parseSetAdder(s)
	}
	if "ACTIVITY_TIMESTAMP" == field_to_set {
		return self.parseSetActivityTimestamp(s)
	}
	if "POSTERS" == field_to_set {
		return self.parseSetPosters(s)
	}
	if "STATUS" == field_to_set {
		return self.parseSetStatus(s)
	}
	if "GUIDELINES" == field_to_set {
		return self.parseSetGuidelines(s)
	}
	if "TOPICXML" == field_to_set {
		return self.parseSetTopicxml(s)
	}
	if "MYSTATUS" == field_to_set {
		return self.parseSetMystatus(s)
	}
	if "MEMBEROBJECTS" == field_to_set {
		return self.parseSetMemberobjects(s)
	}
	if "FRIENDLYNAME" == field_to_set {
		return self.parseSetFriendlyname(s)
	}
	if "ACTIVEMEMBERS" == field_to_set {
		return self.parseSetActivemembers(s)
	}
	if "DESCRIPTION" == field_to_set {
		return self.parseSetDescription(s)
	}
	if "TIMESTAMP" == field_to_set {
		return self.parseSetTimestamp(s)
	}
	if "CHATMESSAGES" == field_to_set {
		return self.parseSetChatmessages(s)
	}
	if "TOPIC" == field_to_set {
		return self.parseSetTopic(s)
	}
	if "ROLE" == field_to_set {
		return self.parseSetRole(s)
	}
	if "BLOB" == field_to_set {
		return self.parseSetBlob(s)
	}
	if "MEMBERS" == field_to_set {
		return self.parseSetMembers(s)
	}
	if "MYROLE" == field_to_set {
		return self.parseSetMyrole(s)
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Chatmember struct {
	Id       string
	Role     string
	IsActive string
	Identity string
	Chatname string
}

func (self *Chatmember) parseSetRole(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHATMEMBER %s ROLE %s", &__, &self.Role); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chatmember) parseSetIsActive(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHATMEMBER %s IS_ACTIVE %s", &__, &self.IsActive); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chatmember) parseSetIdentity(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHATMEMBER %s IDENTITY %s", &__, &self.Identity); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chatmember) parseSetChatname(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHATMEMBER %s CHATNAME %s", &__, &self.Chatname); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
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
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	if "ROLE" == field_to_set {
		return self.parseSetRole(s)
	}
	if "IS_ACTIVE" == field_to_set {
		return self.parseSetIsActive(s)
	}
	if "IDENTITY" == field_to_set {
		return self.parseSetIdentity(s)
	}
	if "CHATNAME" == field_to_set {
		return self.parseSetChatname(s)
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Chatmessage struct {
	Id              string
	Body            string
	Status          string
	EditedTimestamp string
	EditedBy        string
	Users           string
	Timestamp       string
	FromHandle      string
	Chatname        string
	IsEditable      string
	Leavereason     string
	FromDispname    string
	ChatmessageType string
}

func (self *Chatmessage) parseSetBody(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHATMESSAGE %s BODY %s", &__, &self.Body); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chatmessage) parseSetStatus(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHATMESSAGE %s STATUS %s", &__, &self.Status); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chatmessage) parseSetEditedTimestamp(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHATMESSAGE %s EDITED_TIMESTAMP %s", &__, &self.EditedTimestamp); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chatmessage) parseSetEditedBy(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHATMESSAGE %s EDITED_BY %s", &__, &self.EditedBy); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chatmessage) parseSetUsers(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHATMESSAGE %s USERS %s", &__, &self.Users); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chatmessage) parseSetTimestamp(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHATMESSAGE %s TIMESTAMP %s", &__, &self.Timestamp); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chatmessage) parseSetFromHandle(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHATMESSAGE %s FROM_HANDLE %s", &__, &self.FromHandle); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chatmessage) parseSetChatname(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHATMESSAGE %s CHATNAME %s", &__, &self.Chatname); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chatmessage) parseSetIsEditable(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHATMESSAGE %s IS_EDITABLE %s", &__, &self.IsEditable); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chatmessage) parseSetLeavereason(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHATMESSAGE %s LEAVEREASON %s", &__, &self.Leavereason); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chatmessage) parseSetFromDispname(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHATMESSAGE %s FROM_DISPNAME %s", &__, &self.FromDispname); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Chatmessage) parseSetChatmessageType(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "CHATMESSAGE %s TYPE %s", &__, &self.ChatmessageType); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
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
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	if "BODY" == field_to_set {
		return self.parseSetBody(s)
	}
	if "STATUS" == field_to_set {
		return self.parseSetStatus(s)
	}
	if "EDITED_TIMESTAMP" == field_to_set {
		return self.parseSetEditedTimestamp(s)
	}
	if "EDITED_BY" == field_to_set {
		return self.parseSetEditedBy(s)
	}
	if "USERS" == field_to_set {
		return self.parseSetUsers(s)
	}
	if "TIMESTAMP" == field_to_set {
		return self.parseSetTimestamp(s)
	}
	if "FROM_HANDLE" == field_to_set {
		return self.parseSetFromHandle(s)
	}
	if "CHATNAME" == field_to_set {
		return self.parseSetChatname(s)
	}
	if "IS_EDITABLE" == field_to_set {
		return self.parseSetIsEditable(s)
	}
	if "LEAVEREASON" == field_to_set {
		return self.parseSetLeavereason(s)
	}
	if "FROM_DISPNAME" == field_to_set {
		return self.parseSetFromDispname(s)
	}
	if "TYPE" == field_to_set {
		return self.parseSetChatmessageType(s)
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Connstatus struct {
	Id string
}

func (self *Connstatus) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *Connstatus) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "CONNSTATUS "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type ContactsFocused struct {
	Id string
}

func (self *ContactsFocused) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *ContactsFocused) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "CONTACTS_FOCUSED "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Currentuserhandle struct {
	Id string
}

func (self *Currentuserhandle) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *Currentuserhandle) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "CURRENTUSERHANDLE "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
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
	PartnerDispname  string
	FiletransferType string
	Bytestransferred string
	Failurereason    string
}

func (self *Filetransfer) parseSetFinishtime(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "FILETRANSFER %s FINISHTIME %s", &__, &self.Finishtime); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Filetransfer) parseSetStatus(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "FILETRANSFER %s STATUS %s", &__, &self.Status); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Filetransfer) parseSetPartnerHandle(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "FILETRANSFER %s PARTNER_HANDLE %s", &__, &self.PartnerHandle); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Filetransfer) parseSetFilepath(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "FILETRANSFER %s FILEPATH %s", &__, &self.Filepath); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Filetransfer) parseSetBytespersecond(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "FILETRANSFER %s BYTESPERSECOND %s", &__, &self.Bytespersecond); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Filetransfer) parseSetFilesize(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "FILETRANSFER %s FILESIZE %s", &__, &self.Filesize); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Filetransfer) parseSetStarttime(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "FILETRANSFER %s STARTTIME %s", &__, &self.Starttime); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Filetransfer) parseSetPartnerDispname(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "FILETRANSFER %s PARTNER_DISPNAME %s", &__, &self.PartnerDispname); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Filetransfer) parseSetFiletransferType(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "FILETRANSFER %s TYPE %s", &__, &self.FiletransferType); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Filetransfer) parseSetBytestransferred(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "FILETRANSFER %s BYTESTRANSFERRED %s", &__, &self.Bytestransferred); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Filetransfer) parseSetFailurereason(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "FILETRANSFER %s FAILUREREASON %s", &__, &self.Failurereason); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
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
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	if "FINISHTIME" == field_to_set {
		return self.parseSetFinishtime(s)
	}
	if "STATUS" == field_to_set {
		return self.parseSetStatus(s)
	}
	if "PARTNER_HANDLE" == field_to_set {
		return self.parseSetPartnerHandle(s)
	}
	if "FILEPATH" == field_to_set {
		return self.parseSetFilepath(s)
	}
	if "BYTESPERSECOND" == field_to_set {
		return self.parseSetBytespersecond(s)
	}
	if "FILESIZE" == field_to_set {
		return self.parseSetFilesize(s)
	}
	if "STARTTIME" == field_to_set {
		return self.parseSetStarttime(s)
	}
	if "PARTNER_DISPNAME" == field_to_set {
		return self.parseSetPartnerDispname(s)
	}
	if "TYPE" == field_to_set {
		return self.parseSetFiletransferType(s)
	}
	if "BYTESTRANSFERRED" == field_to_set {
		return self.parseSetBytestransferred(s)
	}
	if "FAILUREREASON" == field_to_set {
		return self.parseSetFailurereason(s)
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Group struct {
	Id            string
	Displayname   string
	Users         string
	Expanded      string
	CustomGroupId string
	Visible       string
	GroupType     string
	NrofUsers     string
}

func (self *Group) parseSetDisplayname(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "GROUP %s DISPLAYNAME %s", &__, &self.Displayname); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Group) parseSetUsers(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "GROUP %s USERS %s", &__, &self.Users); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Group) parseSetExpanded(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "GROUP %s EXPANDED %s", &__, &self.Expanded); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Group) parseSetCustomGroupId(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "GROUP %s CUSTOM_GROUP_ID %s", &__, &self.CustomGroupId); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Group) parseSetVisible(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "GROUP %s VISIBLE %s", &__, &self.Visible); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Group) parseSetGroupType(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "GROUP %s TYPE %s", &__, &self.GroupType); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Group) parseSetNrofUsers(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "GROUP %s NROFUSERS %s", &__, &self.NrofUsers); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
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
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	if "DISPLAYNAME" == field_to_set {
		return self.parseSetDisplayname(s)
	}
	if "USERS" == field_to_set {
		return self.parseSetUsers(s)
	}
	if "EXPANDED" == field_to_set {
		return self.parseSetExpanded(s)
	}
	if "CUSTOM_GROUP_ID" == field_to_set {
		return self.parseSetCustomGroupId(s)
	}
	if "VISIBLE" == field_to_set {
		return self.parseSetVisible(s)
	}
	if "TYPE" == field_to_set {
		return self.parseSetGroupType(s)
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Mute struct {
	Id string
}

func (self *Mute) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *Mute) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "MUTE "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Pcspeaker struct {
	Id string
}

func (self *Pcspeaker) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *Pcspeaker) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "PCSPEAKER "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type PredictiveDialerCountry struct {
	Id string
}

func (self *PredictiveDialerCountry) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *PredictiveDialerCountry) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "PREDICTIVE_DIALER_COUNTRY "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Privilege struct {
	Id string
}

func (self *Privilege) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *Privilege) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "PRIVILEGE "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Profile struct {
	Id string
}

func (self *Profile) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *Profile) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "PROFILE "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Ringer struct {
	Id string
}

func (self *Ringer) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *Ringer) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "RINGER "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Ringtone struct {
	Id     string
	Status string
}

func (self *Ringtone) parseSetStatus(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "RINGTONE %s STATUS %s", &__, &self.Status); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Ringtone) getFetchAllFieldsCommands() ([]string, error) {
	return []string{"GET RINGTONE " + self.Id + " STATUS"}, nil
}

func (self *Ringtone) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "RINGTONE "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	if "STATUS" == field_to_set {
		return self.parseSetStatus(s)
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type SilentMode struct {
	Id string
}

func (self *SilentMode) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *SilentMode) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "SILENT_MODE "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Skypeversion struct {
	Id string
}

func (self *Skypeversion) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *Skypeversion) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "SKYPEVERSION "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Sms struct {
	Id             string
	Body           string
	PriceCurrency  string
	TargetNumbers  string
	Status         string
	ReplyToNumber  string
	Chunking       string
	Price          string
	TargetStatuses string
	IsFailedUnseen string
	PricePrecision string
	Timestamp      string
	SmsType        string
	Chunk          string
	Failurereason  string
}

func (self *Sms) parseSetBody(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "SMS %s BODY %s", &__, &self.Body); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Sms) parseSetPriceCurrency(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "SMS %s PRICE_CURRENCY %s", &__, &self.PriceCurrency); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Sms) parseSetTargetNumbers(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "SMS %s TARGET_NUMBERS %s", &__, &self.TargetNumbers); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Sms) parseSetStatus(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "SMS %s STATUS %s", &__, &self.Status); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Sms) parseSetReplyToNumber(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "SMS %s REPLY_TO_NUMBER %s", &__, &self.ReplyToNumber); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Sms) parseSetChunking(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "SMS %s CHUNKING %s", &__, &self.Chunking); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Sms) parseSetPrice(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "SMS %s PRICE %s", &__, &self.Price); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Sms) parseSetTargetStatuses(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "SMS %s TARGET_STATUSES %s", &__, &self.TargetStatuses); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Sms) parseSetIsFailedUnseen(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "SMS %s IS_FAILED_UNSEEN %s", &__, &self.IsFailedUnseen); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Sms) parseSetPricePrecision(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "SMS %s PRICE_PRECISION %s", &__, &self.PricePrecision); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Sms) parseSetTimestamp(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "SMS %s TIMESTAMP %s", &__, &self.Timestamp); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Sms) parseSetSmsType(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "SMS %s TYPE %s", &__, &self.SmsType); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Sms) parseSetChunk(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "SMS %s CHUNK %s", &__, &self.Chunk); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Sms) parseSetFailurereason(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "SMS %s FAILUREREASON %s", &__, &self.Failurereason); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Sms) getFetchAllFieldsCommands() ([]string, error) {
	return []string{"GET SMS " + self.Id + " BODY",
		"GET SMS " + self.Id + " PRICE_CURRENCY",
		"GET SMS " + self.Id + " TARGET_NUMBERS",
		"GET SMS " + self.Id + " STATUS",
		"GET SMS " + self.Id + " REPLY_TO_NUMBER",
		"GET SMS " + self.Id + " CHUNKING",
		"GET SMS " + self.Id + " PRICE",
		"GET SMS " + self.Id + " TARGET_STATUSES",
		"GET SMS " + self.Id + " IS_FAILED_UNSEEN",
		"GET SMS " + self.Id + " PRICE_PRECISION",
		"GET SMS " + self.Id + " TIMESTAMP",
		"GET SMS " + self.Id + " TYPE",
		"GET SMS " + self.Id + " CHUNK",
		"GET SMS " + self.Id + " FAILUREREASON"}, nil
}

func (self *Sms) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "SMS "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	if "BODY" == field_to_set {
		return self.parseSetBody(s)
	}
	if "PRICE_CURRENCY" == field_to_set {
		return self.parseSetPriceCurrency(s)
	}
	if "TARGET_NUMBERS" == field_to_set {
		return self.parseSetTargetNumbers(s)
	}
	if "STATUS" == field_to_set {
		return self.parseSetStatus(s)
	}
	if "REPLY_TO_NUMBER" == field_to_set {
		return self.parseSetReplyToNumber(s)
	}
	if "CHUNKING" == field_to_set {
		return self.parseSetChunking(s)
	}
	if "PRICE" == field_to_set {
		return self.parseSetPrice(s)
	}
	if "TARGET_STATUSES" == field_to_set {
		return self.parseSetTargetStatuses(s)
	}
	if "IS_FAILED_UNSEEN" == field_to_set {
		return self.parseSetIsFailedUnseen(s)
	}
	if "PRICE_PRECISION" == field_to_set {
		return self.parseSetPricePrecision(s)
	}
	if "TIMESTAMP" == field_to_set {
		return self.parseSetTimestamp(s)
	}
	if "TYPE" == field_to_set {
		return self.parseSetSmsType(s)
	}
	if "CHUNK" == field_to_set {
		return self.parseSetChunk(s)
	}
	if "FAILUREREASON" == field_to_set {
		return self.parseSetFailurereason(s)
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Spam struct {
	Id string
}

func (self *Spam) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *Spam) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "SPAM "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type UiLanguage struct {
	Id string
}

func (self *UiLanguage) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *UiLanguage) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "UI_LANGUAGE "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type User struct {
	Id                  string
	Province            string
	About               string
	PhoneOffice         string
	Country             string
	Birthday            string
	IsCfActive          string
	Timezone            string
	Speeddial           string
	Displayname         string
	Language            string
	Isblocked           string
	Onlinestatus        string
	Sex                 string
	CanLeaveVm          string
	MoodText            string
	Homepage            string
	Aliases             string
	IsVideoCapable      string
	Lastonlinetimestamp string
	Buddystatus         string
	Hascallequipment    string
	NrofAuthedBuddies   string
	Receivedauthrequest string
	City                string
	Isauthorized        string
	IsVoicemailCapable  string
	PhoneHome           string
	Avatar              string
	RichMoodText        string
	Fullname            string
	PhoneMobile         string
}

func (self *User) parseSetProvince(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s PROVINCE %s", &__, &self.Province); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetAbout(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s ABOUT %s", &__, &self.About); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetPhoneOffice(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s PHONE_OFFICE %s", &__, &self.PhoneOffice); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetCountry(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s COUNTRY %s", &__, &self.Country); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetBirthday(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s BIRTHDAY %s", &__, &self.Birthday); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetIsCfActive(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s IS_CF_ACTIVE %s", &__, &self.IsCfActive); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetTimezone(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s TIMEZONE %s", &__, &self.Timezone); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetSpeeddial(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s SPEEDDIAL %s", &__, &self.Speeddial); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetDisplayname(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s DISPLAYNAME %s", &__, &self.Displayname); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetLanguage(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s LANGUAGE %s", &__, &self.Language); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetIsblocked(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s ISBLOCKED %s", &__, &self.Isblocked); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetOnlinestatus(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s ONLINESTATUS %s", &__, &self.Onlinestatus); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetSex(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s SEX %s", &__, &self.Sex); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetCanLeaveVm(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s CAN_LEAVE_VM %s", &__, &self.CanLeaveVm); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetMoodText(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s MOOD_TEXT %s", &__, &self.MoodText); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetHomepage(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s HOMEPAGE %s", &__, &self.Homepage); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetAliases(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s ALIASES %s", &__, &self.Aliases); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetIsVideoCapable(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s IS_VIDEO_CAPABLE %s", &__, &self.IsVideoCapable); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetLastonlinetimestamp(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s LASTONLINETIMESTAMP %s", &__, &self.Lastonlinetimestamp); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetBuddystatus(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s BUDDYSTATUS %s", &__, &self.Buddystatus); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetHascallequipment(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s HASCALLEQUIPMENT %s", &__, &self.Hascallequipment); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetNrofAuthedBuddies(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s NROF_AUTHED_BUDDIES %s", &__, &self.NrofAuthedBuddies); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetReceivedauthrequest(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s RECEIVEDAUTHREQUEST %s", &__, &self.Receivedauthrequest); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetCity(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s CITY %s", &__, &self.City); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetIsauthorized(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s ISAUTHORIZED %s", &__, &self.Isauthorized); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetIsVoicemailCapable(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s IS_VOICEMAIL_CAPABLE %s", &__, &self.IsVoicemailCapable); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetPhoneHome(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s PHONE_HOME %s", &__, &self.PhoneHome); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetAvatar(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s AVATAR %s", &__, &self.Avatar); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetRichMoodText(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s RICH_MOOD_TEXT %s", &__, &self.RichMoodText); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetFullname(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s FULLNAME %s", &__, &self.Fullname); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *User) parseSetPhoneMobile(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "USER %s PHONE_MOBILE %s", &__, &self.PhoneMobile); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
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
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	if "PROVINCE" == field_to_set {
		return self.parseSetProvince(s)
	}
	if "ABOUT" == field_to_set {
		return self.parseSetAbout(s)
	}
	if "PHONE_OFFICE" == field_to_set {
		return self.parseSetPhoneOffice(s)
	}
	if "COUNTRY" == field_to_set {
		return self.parseSetCountry(s)
	}
	if "BIRTHDAY" == field_to_set {
		return self.parseSetBirthday(s)
	}
	if "IS_CF_ACTIVE" == field_to_set {
		return self.parseSetIsCfActive(s)
	}
	if "TIMEZONE" == field_to_set {
		return self.parseSetTimezone(s)
	}
	if "SPEEDDIAL" == field_to_set {
		return self.parseSetSpeeddial(s)
	}
	if "DISPLAYNAME" == field_to_set {
		return self.parseSetDisplayname(s)
	}
	if "LANGUAGE" == field_to_set {
		return self.parseSetLanguage(s)
	}
	if "ISBLOCKED" == field_to_set {
		return self.parseSetIsblocked(s)
	}
	if "ONLINESTATUS" == field_to_set {
		return self.parseSetOnlinestatus(s)
	}
	if "SEX" == field_to_set {
		return self.parseSetSex(s)
	}
	if "CAN_LEAVE_VM" == field_to_set {
		return self.parseSetCanLeaveVm(s)
	}
	if "MOOD_TEXT" == field_to_set {
		return self.parseSetMoodText(s)
	}
	if "HOMEPAGE" == field_to_set {
		return self.parseSetHomepage(s)
	}
	if "ALIASES" == field_to_set {
		return self.parseSetAliases(s)
	}
	if "IS_VIDEO_CAPABLE" == field_to_set {
		return self.parseSetIsVideoCapable(s)
	}
	if "LASTONLINETIMESTAMP" == field_to_set {
		return self.parseSetLastonlinetimestamp(s)
	}
	if "BUDDYSTATUS" == field_to_set {
		return self.parseSetBuddystatus(s)
	}
	if "HASCALLEQUIPMENT" == field_to_set {
		return self.parseSetHascallequipment(s)
	}
	if "NROF_AUTHED_BUDDIES" == field_to_set {
		return self.parseSetNrofAuthedBuddies(s)
	}
	if "RECEIVEDAUTHREQUEST" == field_to_set {
		return self.parseSetReceivedauthrequest(s)
	}
	if "CITY" == field_to_set {
		return self.parseSetCity(s)
	}
	if "ISAUTHORIZED" == field_to_set {
		return self.parseSetIsauthorized(s)
	}
	if "IS_VOICEMAIL_CAPABLE" == field_to_set {
		return self.parseSetIsVoicemailCapable(s)
	}
	if "PHONE_HOME" == field_to_set {
		return self.parseSetPhoneHome(s)
	}
	if "AVATAR" == field_to_set {
		return self.parseSetAvatar(s)
	}
	if "RICH_MOOD_TEXT" == field_to_set {
		return self.parseSetRichMoodText(s)
	}
	if "FULLNAME" == field_to_set {
		return self.parseSetFullname(s)
	}
	if "PHONE_MOBILE" == field_to_set {
		return self.parseSetPhoneMobile(s)
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Userstatus struct {
	Id string
}

func (self *Userstatus) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *Userstatus) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "USERSTATUS "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type VideoIn struct {
	Id string
}

func (self *VideoIn) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *VideoIn) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "VIDEO_IN "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Voicemail struct {
	Id              string
	Status          string
	PartnerHandle   string
	Timestamp       string
	AllowedDuration string
	CaptureMic      string
	PartnerDispname string
	Output          string
	Duration        string
	Input           string
	VoicemailType   string
	Failurereason   string
}

func (self *Voicemail) parseSetStatus(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "VOICEMAIL %s STATUS %s", &__, &self.Status); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Voicemail) parseSetPartnerHandle(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "VOICEMAIL %s PARTNER_HANDLE %s", &__, &self.PartnerHandle); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Voicemail) parseSetTimestamp(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "VOICEMAIL %s TIMESTAMP %s", &__, &self.Timestamp); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Voicemail) parseSetAllowedDuration(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "VOICEMAIL %s ALLOWED_DURATION %s", &__, &self.AllowedDuration); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Voicemail) parseSetCaptureMic(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "VOICEMAIL %s CAPTURE_MIC %s", &__, &self.CaptureMic); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Voicemail) parseSetPartnerDispname(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "VOICEMAIL %s PARTNER_DISPNAME %s", &__, &self.PartnerDispname); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Voicemail) parseSetOutput(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "VOICEMAIL %s OUTPUT %s", &__, &self.Output); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Voicemail) parseSetDuration(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "VOICEMAIL %s DURATION %s", &__, &self.Duration); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Voicemail) parseSetInput(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "VOICEMAIL %s INPUT %s", &__, &self.Input); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Voicemail) parseSetVoicemailType(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "VOICEMAIL %s TYPE %s", &__, &self.VoicemailType); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Voicemail) parseSetFailurereason(s string) error {
	var __ string
	if n, e := fmt.Sscanf(s, "VOICEMAIL %s FAILUREREASON %s", &__, &self.Failurereason); e != nil {
		return e
	} else if n != 2 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return nil
}

func (self *Voicemail) getFetchAllFieldsCommands() ([]string, error) {
	return []string{"GET VOICEMAIL " + self.Id + " STATUS",
		"GET VOICEMAIL " + self.Id + " PARTNER_HANDLE",
		"GET VOICEMAIL " + self.Id + " TIMESTAMP",
		"GET VOICEMAIL " + self.Id + " ALLOWED_DURATION",
		"GET VOICEMAIL " + self.Id + " CAPTURE_MIC",
		"GET VOICEMAIL " + self.Id + " PARTNER_DISPNAME",
		"GET VOICEMAIL " + self.Id + " OUTPUT",
		"GET VOICEMAIL " + self.Id + " DURATION",
		"GET VOICEMAIL " + self.Id + " INPUT",
		"GET VOICEMAIL " + self.Id + " TYPE",
		"GET VOICEMAIL " + self.Id + " FAILUREREASON"}, nil
}

func (self *Voicemail) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "VOICEMAIL "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	if "STATUS" == field_to_set {
		return self.parseSetStatus(s)
	}
	if "PARTNER_HANDLE" == field_to_set {
		return self.parseSetPartnerHandle(s)
	}
	if "TIMESTAMP" == field_to_set {
		return self.parseSetTimestamp(s)
	}
	if "ALLOWED_DURATION" == field_to_set {
		return self.parseSetAllowedDuration(s)
	}
	if "CAPTURE_MIC" == field_to_set {
		return self.parseSetCaptureMic(s)
	}
	if "PARTNER_DISPNAME" == field_to_set {
		return self.parseSetPartnerDispname(s)
	}
	if "OUTPUT" == field_to_set {
		return self.parseSetOutput(s)
	}
	if "DURATION" == field_to_set {
		return self.parseSetDuration(s)
	}
	if "INPUT" == field_to_set {
		return self.parseSetInput(s)
	}
	if "TYPE" == field_to_set {
		return self.parseSetVoicemailType(s)
	}
	if "FAILUREREASON" == field_to_set {
		return self.parseSetFailurereason(s)
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Wallpaper struct {
	Id string
}

func (self *Wallpaper) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *Wallpaper) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "WALLPAPER "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type Windowstate struct {
	Id string
}

func (self *Windowstate) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *Windowstate) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "WINDOWSTATE "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}

type In struct {
	Id string
}

func (self *In) getFetchAllFieldsCommands() ([]string, error) {
	return []string{}, nil
}

func (self *In) parseSet(s string) error {
	var field_to_set string
	if n, e := fmt.Sscanf(s, "in "+self.Id+" %s", &field_to_set); e != nil {
		return e
	} else if n != 1 {
		return &UnexpectedNumberOfFieldsError{s: s}
	}
	return &ParserForFieldDoesNotExist{field: field_to_set}
}
