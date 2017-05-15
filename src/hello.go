package hello

func GetHello() string {
	return "Hello"
}

func AddInt(x, y int) int {
	return x + y
}

func AddFloat(x, y float64) float64 {
	return x + y
}

func SumInt(x []int) int {
	total := 0
	for _, value := range x {
		total += value
	}
	return total
}

func AppendStr(x, y string) string {
	return x + y
}

type Vertex struct {
	X int
	Y int
}

func GetVertex(x, y int) Vertex {
	return Vertex{
		X: x,
		Y: y,
	}
}

func GetVertexPtr(x, y int) *Vertex {
	return &Vertex{
		X: x,
		Y: y,
	}
}

func (v *Vertex) Add() int {
	return v.X + v.Y
}

func (v *Vertex) Scale(s int) {
	v.X = v.X * s
	v.Y = v.Y * s
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-1 * f)
	}
	return float64(f)
}
