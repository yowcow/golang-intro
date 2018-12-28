package mywire

type Greeter struct {
	msg Message
}

func NewGreeter(m Message) (Greeter, error) {
	return Greeter{m}, nil
}

func (g Greeter) Greet() Message {
	return g.msg
}
