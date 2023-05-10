package events

type Fetcher interface {
	Fetch() ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type   Type
	Text   string
	UserId int
}

type ViewKeyBoard struct {
	OneTime bool       `json:"one_time"`
	Buttons [][]Button `json:"buttons"`
}

type Button struct {
	Action Action `json:"action"`
}

type Action struct {
	TypeAction string `json:"type"`
	Label      string `json:"label"`
}
