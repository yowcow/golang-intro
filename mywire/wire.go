//+build wireinject

package mywire

import (
	"io"

	"github.com/google/wire"
)

func InitializeEvent(phrase string, w io.Writer) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
