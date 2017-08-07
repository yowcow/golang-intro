package hello

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestHello(t *testing.T) {
	str := GetHello()
	assert.Equal(t, "Hello", str)
}

func TestAddInt(t *testing.T) {
	sum := AddInt(1, 2)
	if sum != 3 {
		t.Error("Expected 3 but got", sum)
	}
}

func TestAddFloat(t *testing.T) {
	sum := AddFloat(1.0, 2.0)
	if sum != 3.0 {
		t.Error("Expected 3.0 but got", sum)
	}
}

func TestSumInt(t *testing.T) {
	x := []int{1, 2, 3, 4, 5}
	sum := SumInt(x)
	if sum != 15 {
		t.Error("Expected 15 but got", sum)
	}
}

func TestAppendStr(t *testing.T) {
	str := AppendStr("Hello ", "world")
	if str != "Hello world" {
		t.Error("Expected 'Hello world' but got", str)
	}
}

func TestGetVertex(t *testing.T) {
	v := GetVertex(1, 2)
	if v.X != 1 {
		t.Error("Expected v.X to be 1 but got", v.X)
	}
	if v.Y != 2 {
		t.Error("Expected v.Y to be 2 but got", v.Y)
	}
}

func TestVertexPtr(t *testing.T) {
	v := GetVertexPtr(3, 4)
	if v.X != 3 {
		t.Error("Expected v.X to be 3 but got", v.X)
	}
	if v.Y != 4 {
		t.Error("Expected v.Y to be 4 but got", v.Y)
	}
}

func TestVertexAdd(t *testing.T) {
	v := GetVertexPtr(3, 4)
	sum := v.Add()
	if sum != 7 {
		t.Error("Expected 7 but got", sum)
	}
}

func TestVertexPtrScale(t *testing.T) {
	v := GetVertexPtr(1, 2)
	v.Scale(2)
	if v.X != 2 {
		t.Error("Expected 2 but got", v.X)
	}
	if v.Y != 4 {
		t.Error("Expected 4 but got", v.Y)
	}
}

func TestMyFloatAbs(t *testing.T) {
	var result float64

	f := MyFloat(-math.Sqrt2)
	result = f.Abs()

	if result != math.Sqrt2 {
		t.Error("Expected math.Sqrt2 but got", result)
	}
}
