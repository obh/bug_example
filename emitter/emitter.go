package emitter

type Event struct {
	MID      int64
	URI      string
	Request  string
	Response string
}

type Emitter interface {
	Emit(event Event) error
}
