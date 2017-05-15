package hello

import (
	"fmt"
	"math"
)

type MyFloat1 float64

type Vertex1 struct {
	X float64
	Y float64
}

type Abser interface {
	Abs() float64
}

func (v *Vertex1) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat1) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func describe(a Abser) string {
	return fmt.Sprintf("(%v, %T)", a, a)
}
