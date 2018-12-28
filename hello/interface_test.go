package hello

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
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

func TestTreeDots(t *testing.T) {
	res := ThreeDots("hoge", "fuga", "foo", "bar")
	assert.Equal(t, []interface{}{"hoge", "fuga", "foo", "bar"}, res)
}

func TestThreeDots2(t *testing.T) {
	in := []interface{}{"hoge", "fuga"}
	res := ThreeDots(in...)
	assert.Equal(t, in, res)
}

func TestCountNoInterface(t *testing.T) {
	count := CountNoInterface(1, "two", 3, "four")
	assert.Equal(t, 4, count)
}

func TestGetFromNoInterface(t *testing.T) {
	assert := assert.New(t)

	res1, err1 := GetFromNoInterface(0, 1, "two", 3, "four")
	res2, err2 := GetFromNoInterface(1, 1, "two", 3, "four")
	res3, err3 := GetFromNoInterface(4, 1, "two", 3, "four")

	assert.Equal(1, res1)
	assert.Nil(err1)

	assert.Equal("two", res2)
	assert.Nil(err2)

	assert.Nil(res3)
	assert.EqualValues(fmt.Errorf("index 4 does not exist"), err3)
}

func TestNewPerson(t *testing.T) {
	assert := assert.New(t)

	p := NewPerson(123, "hoge")

	assert.Equal(123, p.GetID())
	assert.Equal("hoge", p.GetName())
	assert.Equal("Hi, I'm hoge (ID: 123)", fmt.Sprintf("%s", p))
}

type HogeWriter interface {
	Write(string)
}
type HogeReader interface {
	Read() string
}
type HogeResetter interface {
	Reset()
}
type HogeThinger interface {
	HogeWriter
	HogeReader
	HogeResetter
}
type HogeThing struct {
	buffer *bytes.Buffer
}

func newHogeThing() HogeThinger {
	return &HogeThing{&bytes.Buffer{}}
}
func (h *HogeThing) Write(s string) {
	h.buffer.WriteString(s)
}
func (h *HogeThing) Read() string {
	return h.buffer.String()
}
func (h *HogeThing) Reset() {
	h.buffer.Reset()
}

func TestMultipleInterfaces(t *testing.T) {
	assert := assert.New(t)

	hoge := newHogeThing()

	hoge.Write("hoge")
	hoge.Write("fuga")
	hoge.Write("ほげ")
	hoge.Write("ふが")

	assert.Equal("hogefugaほげふが", hoge.Read())

	hoge.Reset()

	assert.Equal("", hoge.Read())
}
