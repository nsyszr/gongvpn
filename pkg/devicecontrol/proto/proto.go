package proto

type MessageType int

const (
	MessageTypeUnknown   MessageType = 0
	MessageTypeHello                 = 1
	MessageTypeWelcome               = 2
	MessageTypeAbort                 = 3
	MessageTypePing                  = 4
	MessageTypePong                  = 5
	MessageTypeError                 = 9
	MessageTypeCall                  = 10
	MessageTypeResult                = 11
	MessageTypePublish               = 20
	MessageTypePublished             = 21
)

type AbortReason int

const (
	AbortReasonNoSuchRealm AbortReason = iota
	AbortReasonProtocolViolation
	AbortReasonInvalidSession
	AbortReasonUnknownException
	AbortReasonTechnicalException
)

func (reason AbortReason) String() string {
	names := []string{
		"ERR_NO_SUCH_REALM",
		"ERR_PROTOCOL_VIOLATION",
		"ERR_INVALID_SESSION",
		"ERR_UNKNOWN_EXCEPTION",
		"ERR_TECHNICAL_EXCEPTION"}

	if reason < AbortReasonNoSuchRealm || reason > AbortReasonTechnicalException {
		return "ERR_UNKNOWN_REASON"
	}

	return names[reason]
}

func int64ToMessageType(msgType int64) MessageType {
	switch msgType {
	case 1:
		return MessageTypeHello
	case 2:
		return MessageTypeWelcome
	case 3:
		return MessageTypeAbort
	case 4:
		return MessageTypePing
	case 5:
		return MessageTypePong
	case 9:
		return MessageTypeError
	case 10:
		return MessageTypeCall
	case 11:
		return MessageTypeResult
	case 20:
		return MessageTypePublish
	case 21:
		return MessageTypePublished
	default:
		return MessageTypeUnknown
	}
}

type Envelope interface {
	MessageType() MessageType
	Message() interface{}
}

type RawMessage struct {
	messageType MessageType
	message     []interface{}
}

func (m RawMessage) MessageType() MessageType {
	return m.messageType
}

func (m RawMessage) Message() interface{} {
	return nil
}

func NewAbortMessage(reason AbortReason, details interface{}) []interface{} {
	var msg []interface{}
	msg = append(msg, MessageTypeAbort)
	msg = append(msg, reason.String())
	msg = append(msg, details)
	return msg
}

/*type MessageHello interface {
	Message
	Realm() string
	Details() interface{}
}

type MessageWelcome interface {
	Message
	SessionID() string
	SetSessionID(sessID string)
	Details() interface{}
	SetDetails(details interface{})
}

type MessageAbort interface {
	Message
	Reason() string
	SetReason(reason string)
	Details() interface{}
	SetDetails(details interface{})
}*/
