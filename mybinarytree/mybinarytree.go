package mybinarytree

import (
	"fmt"
	"math/rand"
	"reflect"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func NewRandomTree(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = Insert(t, (1+v)*k)
	}
	return t
}

func Insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{Value: v}
	}
	if v < t.Value {
		t.Left = Insert(t.Left, v)
	} else {
		t.Right = Insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

func Slice(t *Tree) []int {
	if t == nil {
		return []int{}
	}
	l := []int{t.Value}
	if t.Left != nil {
		ll := Slice(t.Left)
		l = append(ll, l...)
	}
	if t.Right != nil {
		rl := Slice(t.Right)
		l = append(l, rl...)
	}
	return l
}

func Walk(t *Tree, out chan<- int) {
	defer close(out)
	l := Slice(t)
	for _, v := range l {
		out <- v
	}
}

func Same(t1, t2 *Tree) bool {
	l1 := Slice(t1)
	l2 := Slice(t2)
	return reflect.DeepEqual(l1, l2)
}
