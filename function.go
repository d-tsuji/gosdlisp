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
	AddSymbolFunc("*", &Mul{})
	AddSymbolFunc("/", &Div{})
	AddSymbolFunc("DEFUN", &Defun{})
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

func (a *Add) A() {}

func (a Add) String() string {
	return "#<SYSTEM-FUNCTION Add>"
}

func (a *Add) funCall(arguments List) T {
	eval := NewEval()
	arg1 := eval.Evaluate((arguments.(*Cons)).Car)
	arg2 := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	return arg1.(*Integer).add(arg2.(*Integer))
}

type Sub struct{}

func (s *Sub) A() {}

func (s Sub) String() string {
	return "#<SYSTEM-FUNCTION Sub>"
}

func (s *Sub) funCall(arguments List) T {
	eval := NewEval()
	arg1 := eval.Evaluate((arguments.(*Cons)).Car)
	arg2 := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	return arg1.(*Integer).sub(arg2.(*Integer))
}

type Mul struct{}

func (m *Mul) A() {}

func (m Mul) String() string {
	return "#<SYSTEM-FUNCTION Mul>"
}

func (m *Mul) funCall(arguments List) T {
	eval := NewEval()
	arg1 := eval.Evaluate((arguments.(*Cons)).Car)
	arg2 := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	return arg1.(*Integer).mul(arg2.(*Integer))
}

type Div struct{}

func (d *Div) A() {}

func (d Div) String() string {
	return "#<SYSTEM-FUNCTION Div>"
}

func (d *Div) funCall(arguments List) T {
	eval := NewEval()
	arg1 := eval.Evaluate((arguments.(*Cons)).Car)
	arg2 := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	return arg1.(*Integer).div(arg2.(*Integer))
}

type Defun struct{}

func (d *Defun) A() {}

func (d *Defun) String() string {
	return "#<SYSTEM-FUNCTION Defun>"
}

func (d *Defun) funCall(arguments List) T {
	arg1 := (arguments.(*Cons)).Car
	args := (arguments.(*Cons)).Cdr
	fun := arg1.(*Symbol)
	lambda := NewCons(NewSymbol("LAMBDA"), args)
	fun.Function = lambda
	return fun
}
