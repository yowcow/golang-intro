package hello

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type Fruit struct {
	Name  string
	Count int
}

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

func TestSliceOfStructs(t *testing.T) {
	assert := assert.New(t)

	a := []Fruit{
		Fruit{"apple", 1},
		Fruit{"banana", 2},
	}

	a = append(a, Fruit{"orange", 3})

	assert.Equal("apple", a[0].Name)
	assert.Equal(1, a[0].Count)

	assert.Equal("banana", a[1].Name)
	assert.Equal(2, a[1].Count)

	assert.Equal("orange", a[2].Name)
	assert.Equal(3, a[2].Count)
}

func TestSliceOfStructPtrs(t *testing.T) {
	assert := assert.New(t)

	a := []*Fruit{
		&Fruit{"apple", 1},
		&Fruit{"banana", 2},
	}

	a = append(a, &Fruit{"orange", 3})

	assert.Equal("apple", a[0].Name)
	assert.Equal(1, a[0].Count)

	assert.Equal("banana", a[1].Name)
	assert.Equal(2, a[1].Count)

	assert.Equal("orange", a[2].Name)
	assert.Equal(3, a[2].Count)
}
