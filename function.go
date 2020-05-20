package gosdlisp

type Function interface {
	Atom
	funCall(arguments List) T
}

func regist(name string, fun Function) {
	AddSymbolFunc(name, fun)
}

func registSystemFunctions() {
	AddSymbolFunc("CAR", &Car{})
	AddSymbolFunc("CDR", &Cdr{})
	AddSymbolFunc("CONS", &FunCons{})
	AddSymbolFunc("EQ", &Eq{})
	AddSymbolFunc("+", &Add{})
	AddSymbolFunc("-", &Sub{})
}

type Car struct{}

func (c Car) String() string {
	return "#<SYSTEM-FUNCTION Car>"
}

func (c *Car) funCall(arguments List) T {
	return nil
}

func (c *Car) A() {}

type Cdr struct{}

func (c *Cdr) A() {}

func (c Cdr) String() string {
	return "#<SYSTEM-FUNCTION Cdr>"
}

func (c *Cdr) funCall(arguments List) T {
	return nil
}

type FunCons struct{}

func (c *FunCons) A() {}

func (c FunCons) String() string {
	return "#<SYSTEM-FUNCTION FunCons>"
}

func (c *FunCons) funCall(arguments List) T {
	return nil
}

type Eq struct{}

func (c *Eq) A() {}

func (c Eq) String() string {
	return "#<SYSTEM-FUNCTION Eq>"
}

func (c *Eq) funCall(arguments List) T {
	return nil
}

type Add struct{}

func (c *Add) A() {}

func (c Add) String() string {
	return "#<SYSTEM-FUNCTION Add>"
}

func (c *Add) funCall(arguments List) T {
	eval := NewEval()
	arg1 := eval.Evaluate((arguments.(*Cons)).Car)
	arg2 := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	return arg1.(*Integer).add(arg2.(*Integer))
}

type Sub struct{}

func (c *Sub) A() {}

func (c Sub) String() string {
	return "#<SYSTEM-FUNCTION Sub>"
}

func (c *Sub) funCall(arguments List) T {
	eval := NewEval()
	arg1 := eval.Evaluate((arguments.(*Cons)).Car)
	arg2 := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	return arg1.(*Integer).sub(arg2.(*Integer))
}
