package hello

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func sumArray(a []int) int {
	r := 0
	for _, v := range a {
		r += v
	}
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
