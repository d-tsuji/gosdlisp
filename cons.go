package gosdlisp

import "strings"

type Cons struct {
	car T
	cdr T
}

func NewCons(car, cdr T) Cons {
	return Cons{car, cdr}
}

func (c Cons) String() string {
	var str strings.Builder
	str.WriteString("(")

	next := c.cdr

	for {
		str.WriteString(c.car.String())
		if next == nil {
			str.WriteString(")")
			break
		} else if a, ok := next.(Atom); ok {
			str.WriteString(" . ")
			str.WriteString(a.String())
			str.WriteString(")")
			break
		} else {
			str.WriteString(" ")
			next = c.cdr
		}
	}

	return str.String()
}
