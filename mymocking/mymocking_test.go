package mymocking

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_DoSomethingReal_returns_false(t *testing.T) {
	o := new(MyObject)
	o.number = 50
	result, err := doSomethingReal(o, 49)

	assert.False(t, result)
	assert.Nil(t, err)
}

func Test_DoSomethingReal_returns_true(t *testing.T) {
	o := new(MyObject)
	o.number = 50
	result, err := doSomethingReal(o, 50)

	assert.True(t, result)
	assert.Nil(t, err)
}

type MyMockedObject struct {
	mock.Mock
	number int
}

func (m *MyMockedObject) DoSomething(number int) (bool, error) {
	args := m.Called(number)
	return args.Bool(0), args.Error(1)
}

func TestSomething(t *testing.T) {
	testObj := new(MyMockedObject)
	testObj.number = 50
	testObj.On("DoSomething", 50).Return(false, nil)

	result, err := doSomethingReal(testObj, 50)

	assert.False(t, result)
	assert.Nil(t, err)
}
