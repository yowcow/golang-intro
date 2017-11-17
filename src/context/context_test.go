package context

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func generator(ctx context.Context, quit chan<- bool) <-chan int {
	out := make(chan int)
	n := 0
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(out)
				quit <- true
				return
			case out <- n:
				n++
			}
		}
	}()
	return out
}

func Test_WithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	var last int
	quit := make(chan bool)
	in := generator(ctx, quit)
	for last = range in {
		if last == 5 {
			break
		}
	}
	cancel()
	done := <-quit
	close(quit)

	assert.True(t, done)
	assert.Equal(t, 5, last)
}

func Test_WithDeadline(t *testing.T) {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	quit := make(chan bool)
	in := generator(ctx, quit)
	for _ = range in {
	}
	done := <-quit
	close(quit)
	cancel()

	assert.True(t, done)
}

func Test_WithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)

	quit := make(chan bool)
	in := generator(ctx, quit)
	for _ = range in {
	}
	done := <-quit
	close(quit)
	cancel()

	assert.True(t, done)
}

func Test_WithValue(t *testing.T) {
	type favContextKey string
	type notFavContextKey string

	f := func(ctx context.Context, k favContextKey) (interface{}, bool) {
		if v := ctx.Value(k); v != nil {
			return v, true
		}
		return "", false
	}
	ctx := context.WithValue(context.Background(), favContextKey("language"), "Go")

	type Case struct {
		input         favContextKey
		expectedValue string
		expectedOK    bool
		message       string
	}
	cases := []Case{
		{favContextKey("language"), "Go", true, "exact match"},
		{favContextKey("hgoe"), "", false, "no match"},
	}

	for _, c := range cases {
		t.Run(c.message, func(t *testing.T) {
			v, ok := f(ctx, c.input)

			assert.Equal(t, c.expectedOK, ok)
			assert.Equal(t, c.expectedValue, v)
		})
	}
}
