package hello

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func sumArray(a []int) int {
	fmt.Println("Begin summing an array:", a)
	time.Sleep(200 * time.Millisecond)
	r := 0
	for _, v := range a {
		r += v
	}
	fmt.Println("Done summing an array:", r)
	return r
}

func sumArrayRunner(ch chan int, in []int) {
	res := sumArray(in)
	ch <- res
}

func TestSimpleSum(t *testing.T) {
	assert := assert.New(t)

	chan1 := make(chan int)
	chan2 := make(chan int)

	go sumArrayRunner(chan1, []int{1, 2, 3, 4, 5})
	go sumArrayRunner(chan2, []int{2, 3, 4, 5, 6})

	assert.Equal(15, <-chan1)
	assert.Equal(20, <-chan2)
}

func TestSumWaited(t *testing.T) {
	assert := assert.New(t)

	sumArrayWaited := func(ch chan int, w *sync.WaitGroup, in []int) {
		res := sumArray(in)
		ch <- res
		w.Done()
	}

	c := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go sumArrayWaited(c, wg, []int{1, 2, 3, 4, 5})
	wg.Wait()

	wg.Add(1)
	go sumArrayWaited(c, wg, []int{2, 3, 4, 5, 6})
	wg.Wait()

	ret1 := <-c
	ret2 := <-c

	assert.Equal(15, ret1)
	assert.Equal(20, ret2)
}

func fib(c, q chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-q:
			return
		}
	}
}

func fibRunner(c, q chan int, r chan []int) {
	result := []int{}
	for i := 0; i < 10; i++ {
		result = append(result, <-c)
	}
	q <- 0
	r <- result
}

func TestReceiverToQuitSender(t *testing.T) {
	assert := assert.New(t)

	c := make(chan int)
	q := make(chan int)
	r := make(chan []int)

	go fibRunner(c, q, r)
	fib(c, q)
	ret := <-r

	assert.EqualValues([]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}, ret)
}

func boomAfterTick(c chan []string) {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(300 * time.Millisecond)
	ret := []string{}

	for {
		select {
		case <-tick:
			ret = append(ret, "tick")
		case <-boom:
			c <- append(ret, "boom")
		default:
			ret = append(ret, ".")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func TestBoomAfterTick(t *testing.T) {
	assert := assert.New(t)

	ch := make(chan []string)

	go boomAfterTick(ch)

	ret := <-ch

	assert.True(9 <= len(ret))
}

func incrementPtr(name string, counter *int, count int, mx *sync.Mutex, wg *sync.WaitGroup) {
	incr := func() {
		mx.Lock()
		defer mx.Unlock()

		c := *counter
		time.Sleep(100 * time.Millisecond)
		*counter = c + 1
	}

	for i := 0; i < count; i++ {
		incr()
		fmt.Printf("(%s) counter is now %d\n", name, *counter)
	}

	wg.Done()
}

func TestIncrementLocked(t *testing.T) {
	assert := assert.New(t)

	counter := 0

	mx := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go incrementPtr("worker1", &counter, 5, mx, wg)

	wg.Add(1)
	go incrementPtr("worker2", &counter, 5, mx, wg)

	wg.Wait()

	assert.Equal(10, counter)
}
