package gosdlisp

import "strconv"

type Integer struct {
	Value int
}

func (i *Integer) A() {}

func NewInteger(value int) *Integer {
	return &Integer{Value: value}
}

func (i *Integer) String() string {
	return strconv.Itoa(i.Value)
}

func (i *Integer) add(arg2 *Integer) T {
	return &Integer{Value: i.Value + arg2.Value}
}

func (i *Integer) sub(arg2 *Integer) T {
	return &Integer{Value: i.Value - arg2.Value}
}

func (i *Integer) mul(arg2 *Integer) T {
	return &Integer{Value: i.Value * arg2.Value}
}

func (i *Integer) div(arg2 *Integer) T {
	return &Integer{Value: i.Value / arg2.Value}
}

func (i *Integer) ge(arg2 *Integer) T {
	if i.Value >= arg2.Value {
		return NewSymbol("T")
	}
	return nil
}

func (i *Integer) le(arg2 *Integer) T {
	if i.Value <= arg2.Value {
		return NewSymbol("T")
	}
	return nil
}

func (i *Integer) gt(arg2 *Integer) T {
	if i.Value > arg2.Value {
		return NewSymbol("T")
	}
	return nil
}

func (i *Integer) lt(arg2 *Integer) T {
	if i.Value < arg2.Value {
		return NewSymbol("T")
	}
	return nil
}

func (i *Integer) numberEqual(arg2 *Integer) T {
	if i.Value == arg2.Value {
		return NewSymbol("T")
	}
	return nil
}
