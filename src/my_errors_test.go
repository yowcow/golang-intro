package hello

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func recoversRuntime(a, b int, ch chan string) {
	defer func() {
		if err := recover(); err != nil {
			ch <- "died"
		}
	}()

	_ = a / b

	ch <- "lived"
}

func TestOnRuntimeError(t *testing.T) {
	assert := assert.New(t)

	ch := make(chan string)

	go recoversRuntime(1, 0, ch)

	assert.Equal("died", <-ch)

	go recoversRuntime(1, 1, ch)

	assert.Equal("lived", <-ch)
}
