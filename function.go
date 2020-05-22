package gosdlisp

import "reflect"

type Function interface {
	Atom
	funCall(arguments List) T
}

func regist(name string, fun Function) {
	AddSymbolFunc(name, fun)
}

func registSystemFunctions() {
	regist("CAR", &Car{})
	regist("CDR", &Cdr{})
	regist("CONS", &FunCons{})
	regist("EQ", &Eq{})
	regist("+", &Add{})
	regist("-", &Sub{})
	regist("*", &Mul{})
	regist("/", &Div{})
	regist("MOD", &Mod{})
	regist(">=", &Ge{})
	regist("<=", &Le{})
	regist(">", &Gt{})
	regist("<", &Lt{})
	regist("=", &NumberEqual{})
	regist("SETQ", &Setq{})
	regist("QUOTE", &Quote{})
	regist("DEFUN", &Defun{})
	regist("IF", &If{})
	regist("SYMBOL-FUNCTION", &SymbolFunction{})
}

type Car struct{}

func (c Car) String() string {
	return "#<SYSTEM-FUNCTION Car>"
}

func (c *Car) funCall(arguments List) T {
	eval := NewEval()
	arg1 := eval.Evaluate((arguments.(*Cons)).Car)
	if arg1 != nil {
		return (arg1.(*Cons)).Car
	}
	return nil
}

func (c *Car) A() {}

type Cdr struct{}

func (c *Cdr) A() {}

func (c Cdr) String() string {
	return "#<SYSTEM-FUNCTION Cdr>"
}

func (c *Cdr) funCall(arguments List) T {
	eval := NewEval()
	arg1 := eval.Evaluate((arguments.(*Cons)).Car)
	if arg1 != nil {
		return (arg1.(*Cons)).Cdr
	}
	return nil
}

type FunCons struct{}

func (c *FunCons) A() {}

func (c FunCons) String() string {
	return "#<SYSTEM-FUNCTION FunCons>"
}

func (c *FunCons) funCall(arguments List) T {
	eval := NewEval()
	arg1 := eval.Evaluate((arguments.(*Cons)).Car)
	arg2 := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	return NewCons(arg1, arg2)
}

type Eq struct{}

func (c *Eq) A() {}

func (c Eq) String() string {
	return "#<SYSTEM-FUNCTION Eq>"
}

func (c *Eq) funCall(arguments List) T {
	eval := NewEval()
	arg1 := eval.Evaluate((arguments.(*Cons)).Car)
	arg2 := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	if reflect.DeepEqual(arg1, arg2) {
		return NewSymbol("T")
	}
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

type Mod struct{}

func (*Mod) A() {}

func (Mod) String() string {
	return "#<SYSTEM-FUNCTION Mod>"
}

func (*Mod) funCall(arguments List) T {
	eval := NewEval()
	arg1 := eval.Evaluate((arguments.(*Cons)).Car)
	arg2 := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	return arg1.(*Integer).mod(arg2.(*Integer))
}

type Ge struct{}

func (g *Ge) A() {}

func (g Ge) String() string {
	return "#<SYSTEM-FUNCTION Ge>"
}

func (g *Ge) funCall(arguments List) T {
	eval := NewEval()
	arg1 := eval.Evaluate((arguments.(*Cons)).Car)
	arg2 := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	return arg1.(*Integer).ge(arg2.(*Integer))
}

type Le struct{}

func (l *Le) A() {}

func (l Le) String() string {
	return "#<SYSTEM-FUNCTION Le>"
}

func (l *Le) funCall(arguments List) T {
	eval := NewEval()
	arg1 := eval.Evaluate((arguments.(*Cons)).Car)
	arg2 := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	return arg1.(*Integer).le(arg2.(*Integer))
}

type Gt struct{}

func (g *Gt) A() {}

func (g Gt) String() string {
	return "#<SYSTEM-FUNCTION Gt>"
}

func (g *Gt) funCall(arguments List) T {
	eval := NewEval()
	arg1 := eval.Evaluate((arguments.(*Cons)).Car)
	arg2 := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	return arg1.(*Integer).gt(arg2.(*Integer))
}

type Lt struct{}

func (l *Lt) A() {}

func (l Lt) String() string {
	return "#<SYSTEM-FUNCTION Lt>"
}

func (l *Lt) funCall(arguments List) T {
	eval := NewEval()
	arg1 := eval.Evaluate((arguments.(*Cons)).Car)
	arg2 := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	return arg1.(*Integer).lt(arg2.(*Integer))
}

type NumberEqual struct{}

func (n *NumberEqual) A() {}

func (n NumberEqual) String() string {
	return "#<SYSTEM-FUNCTION NumberEqual>"
}

func (n *NumberEqual) funCall(arguments List) T {
	eval := NewEval()
	arg1 := eval.Evaluate((arguments.(*Cons)).Car)
	arg2 := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	return arg1.(*Integer).numberEqual(arg2.(*Integer))
}

type Setq struct{}

func (*Setq) A() {}

func (Setq) String() string {
	return "#<SYSTEM-FUNCTION Setq>"
}

func (*Setq) funCall(arguments List) T {
	e := NewEval()
	arg1 := (arguments.(*Cons)).Car
	arg2 := e.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	sym := arg1.(*Symbol)
	sym.Value = e.Evaluate(arg2)
	return sym
}

type Quote struct{}

func (*Quote) A() {}

func (Quote) String() string {
	return "#<SYSTEM-FUNCTION Quote>"
}

func (*Quote) funCall(arguments List) T {
	return arguments.(*Cons).Car
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

type If struct{}

func (i *If) A() {}

func (i *If) String() string {
	return "#<SYSTEM-FUNCTION If>"
}

func (i *If) funCall(arguments List) T {
	arg1 := (arguments.(*Cons)).Car
	args := (arguments.(*Cons)).Cdr
	arg2 := (args.(*Cons)).Car
	var arg3 T
	if (args.(*Cons)).Cdr != nil {
		arg3 = (((args.(*Cons)).Cdr).(*Cons)).Car
	}
	e := NewEval()
	if e.Evaluate(arg1) != nil {
		return e.Evaluate(arg2)
	} else {
		return e.Evaluate(arg3)
	}
}

type SymbolFunction struct{}

func (*SymbolFunction) A() {}

func (*SymbolFunction) String() string {
	return "#<SYSTEM-FUNCTION If>"
}

func (*SymbolFunction) funCall(arguments List) T {
	e := NewEval()
	arg1 := e.Evaluate((arguments.(*Cons)).Car)
	return (arg1.(*Symbol)).Function
}
