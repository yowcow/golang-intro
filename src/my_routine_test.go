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
