package main

import (
	"container/heap"
	_ "fmt"
	"reflect"
	"testing"
)

type Item struct {
	value    string
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].priority > (*pq)[j].priority
}

func (pq *PriorityQueue) Pop() interface{} {
	item := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	item.index = -1
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	n := pq.Len()
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index = i
	(*pq)[j].index = j
}

func TestHeapInterface(t *testing.T) {
	pq := make(PriorityQueue, 1)
	heap.Init(&pq)
}

func createPriorityQueue(items *map[string]int) *PriorityQueue {
	pq := make(PriorityQueue, len(*items))
	i := 0
	for value, priority := range *items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
	return &pq
}

func TestPriorityQueueOrderAtInit(t *testing.T) {
	items := map[string]int{
		"banana": 3,
		"apple":  2,
		"pear":   4,
	}
	pq := createPriorityQueue(&items)

	actual := make([]string, pq.Len())
	i := 0
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		actual[i] = item.value
		i++
	}

	expected := []string{"pear", "banana", "apple"}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %#v but got %#v", expected, actual)
	}
}

func TestPriorityQueueOrderAfterPush(t *testing.T) {
	items := map[string]int{
		"banana": 3,
		"apple":  2,
		"pear":   4,
	}
	pq := createPriorityQueue(&items)

	// push a new item with the highest priority
	item := &Item{
		value:    "orange",
		priority: 5,
	}
	heap.Push(pq, item)

	actual := make([]string, pq.Len())
	i := 0
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		actual[i] = item.value
		i++
	}

	expected := []string{"orange", "pear", "banana", "apple"}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %#v but got %#v", expected, actual)
	}
}

func TestPriorityQueueAfterFix(t *testing.T) {
	items := map[string]int{
		"banana": 3,
		"apple":  2,
		"pear":   4,
	}
	pq := createPriorityQueue(&items)

	// push a new item with the lowest priority
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(pq, item)

	// then modify its priority to the lowest
	item.priority = 5
	heap.Fix(pq, item.index)

	actual := make([]string, pq.Len())
	i := 0
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		actual[i] = item.value
		i++
	}

	expected := []string{"orange", "pear", "banana", "apple"}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %#v but got %#v", expected, actual)
	}
}
