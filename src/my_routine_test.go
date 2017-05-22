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

func TestSimpleSum(t *testing.T) {
	assert := assert.New(t)

	doSumArray := func(ch chan int, in []int) {
		res := sumArray(in)
		ch <- res
	}

	chan1 := make(chan int)
	chan2 := make(chan int)

	go doSumArray(chan1, []int{1, 2, 3, 4, 5})
	go doSumArray(chan2, []int{2, 3, 4, 5, 6})

	assert.Equal(15, <-chan1)
	assert.Equal(20, <-chan2)
}

func TestSumLocked(t *testing.T) {
	assert := assert.New(t)

	doSumArray := func(ch chan int, m *sync.Mutex, in []int) {
		m.Lock()
		defer m.Unlock()

		res := sumArray(in)
		ch <- res
	}

	c := make(chan int, 1)
	m := new(sync.Mutex)

	go doSumArray(c, m, []int{1, 2, 3, 4, 5})
	res1 := <-c

	go doSumArray(c, m, []int{2, 3, 4, 5, 6})
	res2 := <-c

	assert.Equal(15, res1)
	assert.Equal(20, res2)
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
	ret := []int{}

	go fibRunner(c, q, r)
	fib(c, q)
	ret = <-r

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
