package main

import (
	"container/list"
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	l := list.New()

	l.PushFront("hoge")
	l.PushBack("fuga")
	l.PushFront("foo")
	l.PushBack("bar")

	if l.Len() != 4 {
		t.Error("expected 4 but got", l.Len())
	}

	actual := make([]string, l.Len())
	i := 0
	for l.Len() > 0 {
		el := l.Front()
		actual[i] = el.Value.(string)
		l.Remove(el)
		i++
	}

	expected := []string{"foo", "hoge", "fuga", "bar"}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %#v but got %#v", expected, actual)
	}
}

func TestListAndList(t *testing.T) {
	l1 := list.New()
	l1.PushBack("hoge")
	l1.PushBack("fuga")

	l2 := list.New()
	l2.PushBack("foo")
	l2.PushBack("bar")

	l := list.New()
	l.PushFrontList(l1)
	l.PushBackList(l2)

	actual := make([]string, l.Len())
	i := 0
	for l.Len() > 0 {
		el := l.Front()
		actual[i] = el.Value.(string)
		l.Remove(el)
		i++
	}

	expected := []string{"hoge", "fuga", "foo", "bar"}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %#v but got %#v", expected, actual)
	}
}
