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

func ThreeDots(a ...interface{}) []interface{} {
	return a
}

func CountNoInterface(a ...interface{}) int {
	return len(a)
}

func GetFromNoInterface(idx int, a ...interface{}) (interface{}, error) {
	if idx >= len(a) {
		return nil, fmt.Errorf("index %d does not exist", idx)
	}
	return a[idx], nil
}

type Person struct {
	ID   int
	Name string
}

type PersonInterface interface {
	GetId() int
	GetName() string
	String() string
}

func NewPerson(id int, name string) PersonInterface {
	person := Person{id, name}
	return person
}

func (p Person) GetId() int {
	return p.ID
}

func (p Person) GetName() string {
	return p.Name
}

func (p Person) String() string {
	return fmt.Sprintf("Hi, I'm %s (ID: %d)", p.GetName(), p.GetId())
}
