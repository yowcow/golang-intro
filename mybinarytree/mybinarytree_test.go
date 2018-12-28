package mybinarytree_test

import (
	"reflect"
	"testing"

	bt "github.com/yowcow/golang-intro/mybinarytree"
)

func TestInsert(t *testing.T) {
	var root *bt.Tree
	var tr *bt.Tree

	root = bt.Insert(root, 1)
	if root.Value != 1 {
		t.Error("expected 1 but got", root.Value)
	}

	tr = bt.Insert(root, 3)
	if tr.Value != 1 {
		t.Error("expected 1 but got", tr.Value)
	}

	tr = bt.Insert(root, 2)
	if tr.Value != 1 {
		t.Error("expected 1 but got", tr.Value)
	}
}

func TestString(t *testing.T) {
	root := bt.Insert(nil, 3)
	bt.Insert(root, 5)
	bt.Insert(root, 6)
	bt.Insert(root, 1)
	bt.Insert(root, 4)
	bt.Insert(root, 2)

	expected := `((1 (2)) 3 ((4) 5 (6)))`
	actual := root.String()

	if actual != expected {
		t.Error("expected", expected, "but got", actual)
	}
}

func TestSlice(t *testing.T) {
	root := bt.Insert(nil, 3)
	bt.Insert(root, 5)
	bt.Insert(root, 6)
	bt.Insert(root, 1)
	bt.Insert(root, 4)
	bt.Insert(root, 2)

	expected := []int{1, 2, 3, 4, 5, 6}
	actual := bt.Slice(root)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %#v but got %#v", expected, actual)
	}
}

func TestWalk(t *testing.T) {
	root := bt.Insert(nil, 3)
	bt.Insert(root, 5)
	bt.Insert(root, 6)
	bt.Insert(root, 1)
	bt.Insert(root, 4)
	bt.Insert(root, 2)

	out := make(chan int)
	go bt.Walk(root, out)

	actual := make([]int, 0, 6)
	for v := range out {
		actual = append(actual, v)
	}

	expected := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %#v but got %#v", expected, actual)
	}
}

func TestSame(t *testing.T) {
	cases := []struct {
		subtest  string
		tr1      *bt.Tree
		tr2      *bt.Tree
		expected bool
	}{
		{
			"true for trees with the same values",
			bt.NewRandomTree(1),
			bt.NewRandomTree(1),
			true,
		},
		{
			"false for trees with different values",
			bt.NewRandomTree(1),
			bt.NewRandomTree(2),
			false,
		},
	}

	for _, c := range cases {
		t.Run(c.subtest, func(t *testing.T) {
			actual := bt.Same(c.tr1, c.tr2)
			if c.expected != actual {
				t.Error("expected", c.expected, "but got", actual)
			}
		})
	}
}
