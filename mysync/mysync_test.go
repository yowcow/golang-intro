package mysync

import (
	"context"
	"fmt"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

func TestErrGroupWithContext(t *testing.T) {
	msgs := []string{}

	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		msgs = append(msgs, "go1: zzz...")
		time.Sleep(1000 * time.Millisecond)
		msgs = append(msgs, "go1: woke up!")
		return nil
	})

	g.Go(func() error {
		msgs = append(msgs, "go2: zzz...")
		time.Sleep(100 * time.Millisecond)
		return fmt.Errorf("go2 has gone wrong!!")
	})

	g.Go(func() error {
		msgs = append(msgs, "go3: zzz...")
		time.Sleep(1 * time.Millisecond)
		return fmt.Errorf("go3 has gone wrong!!")
	})

	err := g.Wait()

	if msgs[len(msgs)-1] != "go1: woke up!" {
		t.Error("expected 'go1: woke up!' but got", msgs[len(msgs)-1])
	}
	if err.Error() != "go3 has gone wrong!!" {
		t.Error("expected 'go3 has gone wrong!!' but got", err.Error())
	}
}
