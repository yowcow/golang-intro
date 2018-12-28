package mymocking

type MyObject struct {
	number int
}

type MyObjectInterface interface {
	DoSomething(int) (bool, error)
}

func (o MyObject) DoSomething(number int) (bool, error) {
	sum := o.number + number
	return sum >= 100, nil
}

func doSomethingReal(o MyObjectInterface, number int) (bool, error) {
	return o.DoSomething(number)
}
