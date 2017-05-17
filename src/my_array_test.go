package hello

import (
	"reflect"
	"testing"
)

func TestFlatSlice(t *testing.T) {
	scores := []int{1, 2, 3}
	if scores[0] != 1 {
		t.Error("Expected 1 but got", scores[0])
	}
	if scores[1] != 2 {
		t.Error("Expected 2 but got", scores[1])
	}
	if scores[2] != 3 {
		t.Error("Expected 3 but got", scores[2])
	}
}

func TestSliceOfSlices(t *testing.T) {
	items := [][]int{
		{1, 2},
		{3, 4},
	}
	if !reflect.DeepEqual(items[0], []int{1, 2}) {
		t.Error("Expected [1, 2] but got", items[0])
	}
	if !reflect.DeepEqual(items[1], []int{3, 4}) {
		t.Error("Expected [3, 4] but got", items[1])
	}
}
