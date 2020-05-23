package gosdlisp

import (
	"strings"
)

type T interface{ String() string }

type Number interface{ Atom }

type List interface{ T }

type Atom interface {
	T
	A()
}

type Cons struct{ Car, Cdr T }

func NewCons(car, cdr T) *Cons { return &Cons{car, cdr} }

func (c Cons) String() string {
	var str strings.Builder
	str.WriteString("(")

	list := &c

	for {
		str.WriteString(list.Car.String())
		if list.Cdr == nil {
			str.WriteString(")")
			break
		} else if a, ok := list.Cdr.(Atom); ok {
			str.WriteString(" . ")
			str.WriteString(a.String())
			str.WriteString(")")
			break
		} else {
			str.WriteString(" ")
			list = list.Cdr.(*Cons)
		}
	}

	return str.String()
}
