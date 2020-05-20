package gosdlisp

type T interface {
	String() string
}

type Number interface {
	Atom
}

type List interface {
	T
}

type Atom interface {
	T
	A()
}
