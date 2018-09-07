package errors

import (
	"testing"

	"github.com/pkg/errors"
)

var (
	_ error = (*MyError)(nil)
	_ error = (*MyError2)(nil)
)

type MyError struct {
	s string
}

func (e *MyError) Error() string {
	return e.s
}

type MyError2 struct {
	s string
}

func (e *MyError2) Error() string {
	return e.s
}

func fails(str string) error {
	return &MyError{str}
}

func TestMyError(t *testing.T) {
	err := fails("hogefuga")

	if serr, ok := err.(*MyError); !ok {
		t.Errorf("expected err to be *MyError but got something-else: %#v", serr)
	}

	if serr, ok := err.(*MyError2); ok {
		t.Errorf("expected err to be not *MyError2 but got *MyError: %#v", serr)
	}
}

func TestWrapMyError(t *testing.T) {
	err := fails("hoge")
	err = errors.Wrap(err, "fuga")
	err = errors.Wrap(err, "foo")
	err = errors.Wrap(err, "bar")

	if err.Error() != "bar: foo: fuga: hoge" {
		t.Errorf("expected wrapped error but got %+v", err)
	}
}
