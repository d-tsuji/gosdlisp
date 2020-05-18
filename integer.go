package gosdlisp

import "strconv"

type Integer struct {
	value int
}

func NewInteger(value int) *Integer {
	return &Integer{value: value}
}

func (i *Integer) String() string {
	return strconv.Itoa(i.value)
}

func (i *Integer) add(arg2 Integer) T {
	return &Integer{value: i.value + arg2.value}
}

func (i *Integer) sub(arg2 Integer) T {
	return &Integer{value: i.value - arg2.value}
}

func (i *Integer) mul(arg2 Integer) T {
	return &Integer{value: i.value * arg2.value}
}

func (i *Integer) div(arg2 Integer) T {
	return &Integer{value: i.value / arg2.value}
}

func (i *Integer) ge(arg2 Integer) T {
	// TODO
	return nil
}

func (i *Integer) le(arg2 Integer) T {
	// TODO
	return nil
}

func (i *Integer) gt(arg2 Integer) T {
	// TODO
	return nil
}

func (i *Integer) lt(arg2 Integer) T {
	// TODO
	return nil
}
