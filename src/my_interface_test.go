package hello

import (
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	var a Abser
	f := MyFloat1(-math.Sqrt2)
	v := Vertex1{3.0, 4.0}

	a = f

	if a.Abs() != math.Sqrt2 {
		t.Error("Expected math.Sqrt2 but got", a.Abs())
	}

	a = &v

	if a.Abs() != 5.0 {
		t.Error("Expected 5.0 but got", a.Abs())
	}
}

func TestAbsWithImplicitInterface(t *testing.T) {
	var a Abser = &Vertex1{3.0, 4.0}

	if a.Abs() != 5.0 {
		t.Error("Expected 5.0 but got", a.Abs())
	}
}

func TestDescribeVertex1(t *testing.T) {
	var v Abser = &Vertex1{3.0, 4.0}
	var description = describe(v)
	var expected = "(&{3 4}, *hello.Vertex1)"

	if description != expected {
		t.Error("Expected", expected, "but got", description)
	}
}

func TestDescribeMyFloat1(t *testing.T) {
	var f Abser = MyFloat1(1.0)
	var description = describe(f)
	var expected = "(1, hello.MyFloat1)"

	if description != expected {
		t.Error("Expected", expected, "but got", description)
	}
}
