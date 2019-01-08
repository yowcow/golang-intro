package mysync

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

func TestOnce(t *testing.T) {
	msgs := []string{}
	f := func() {
		msgs = append(msgs, "must be once")
	}

	var once sync.Once
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(o *sync.Once, w *sync.WaitGroup) {
			defer w.Done()
			o.Do(f)
		}(&once, &wg)
	}
	wg.Wait()

	expected := []string{"must be once"}

	if !reflect.DeepEqual(expected, msgs) {
		t.Errorf("expected %#v but got %#v", expected, msgs)
	}
}

func TestGoOnce(t *testing.T) {
	msgs := []string{}
	f := func() {
		msgs = append(msgs, "must be once")
	}

	var once sync.Once
	var g errgroup.Group

	for i := 0; i < 10; i++ {
		g.Go(func() error {
			once.Do(f)
			return nil
		})
	}
	err := g.Wait()

	if err != nil {
		t.Error("expected nil but got", err)
	}

	expected := []string{"must be once"}

	if !reflect.DeepEqual(expected, msgs) {
		t.Errorf("expected %#v but got %#v", expected, msgs)
	}
}

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
