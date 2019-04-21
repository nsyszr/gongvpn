package devicecontrol

type State int

const (
	StateEstablished State = iota
	StateAnnounced
	StateRegistered
	StateAborted
	StateClosed
)

func (state State) String() string {
	names := []string{
		"ESTABLISHED",
		"ANNOUNCED",
		"REGISTERED",
		"ABORTED",
		"CLOSED"}

	if state < StateEstablished || state > StateClosed {
		return "UNKNOWN"
	}

	return names[state]
}

type Session struct {
	ManagedDeviceID string
	State           State
}

func NewSession() *Session {
	return &Session{State: StateEstablished}
}
