package mypanic

import (
	"errors"
	"testing"
)

func alwaysPanics(in chan<- int) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	close(in)
	return
}

func TestPanicAndRecover(t *testing.T) {
	var in chan int

	err := alwaysPanics(in)
	expected := errors.New("close of nil channel")

	if err.Error() != expected.Error() {
		t.Errorf("expected '%s' but got '%s'", expected, err)
	}
}
