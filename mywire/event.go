package mywire

import (
	"fmt"
	"io"
)

type Event struct {
	greeter Greeter
	w       io.Writer
}

func NewEvent(g Greeter, w io.Writer) (Event, error) {
	return Event{g, w}, nil // suppose it may return an error
}

func (e Event) Start() {
	fmt.Fprintln(e.w, e.greeter.Greet())
}
