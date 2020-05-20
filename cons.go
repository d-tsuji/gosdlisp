package gosdlisp

import (
	"log"
	"strings"
)

type Cons struct {
	car T
	cdr T
}

func NewCons(car, cdr T) *Cons {
	return &Cons{car, cdr}
}

func (c Cons) String() string {
	var str strings.Builder
	str.WriteString("(")

	list := &c

	for {
		str.WriteString(list.car.String())
		if list.cdr == nil {
			str.WriteString(")")
			break
		} else if a, ok := list.cdr.(Atom); ok {
			str.WriteString(" . ")
			str.WriteString(a.String())
			str.WriteString(")")
			break
		} else {
			str.WriteString(" ")
			l, ok := list.cdr.(*Cons)
			if !ok {
				log.Fatalf("cannot convert Cons: %v", list.cdr)
			}
			list = l
		}
	}

	return str.String()
}
