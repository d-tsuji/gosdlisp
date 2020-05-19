package gosdlisp

func init() {
	AddSymbolFunc("CAR", &Car{})
	AddSymbolFunc("CDR", &Cdr{})
}

type Function interface {
	Atom
	funCall(arguments List) T
}

func regist(name string, fun Function) {
	AddSymbolFunc(name, fun)
}

type Car struct{}

func (c Car) String() string {
	return "#<SYSTEM-FUNCTION Car>"
}

func (c *Car) funCall(arguments List) T {
	return nil
}

type Cdr struct{}

func (c Cdr) String() string {
	return "#<SYSTEM-FUNCTION Cdr>"
}

func (c *Cdr) funCall(arguments List) T {
	return nil
}
