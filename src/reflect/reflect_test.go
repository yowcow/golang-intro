package main

import (
	"reflect"
	"testing"
)

type MyReflect struct {
	FName string `hoge:"fname"`
	LName string `hoge:"lname"`
	Age   int    `hoge:"age"`
}

func TestTag(t *testing.T) {
	rf := new(MyReflect)
	tp := reflect.TypeOf(*rf)

	type Elem struct {
		Name string
		Tag  string
	}

	actual := make([]Elem, 0, 3)
	for i := 0; i < tp.NumField(); i++ {
		f := tp.Field(i)
		actual = append(actual, Elem{
			Name: f.Name,
			Tag:  f.Tag.Get("hoge"),
		})
	}

	expected := []Elem{
		{"FName", "fname"},
		{"LName", "lname"},
		{"Age", "age"},
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}
