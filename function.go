package gosdlisp

import "reflect"

type Function interface {
	Atom
	funCall(arguments List) (T, error)
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

func (*Car) String() string {
	return "#<SYSTEM-FUNCTION Car>"
}

func (*Car) funCall(arguments List) (T, error) {
	eval := NewEval()
	arg1, err := eval.Evaluate((arguments.(*Cons)).Car)
	if err != nil {
		return nil, err
	}
	if arg1 != nil {
		return (arg1.(*Cons)).Car, nil
	}
	return nil, nil
}

func (*Car) A() {}

type Cdr struct{}

func (*Cdr) A() {}

func (*Cdr) String() string {
	return "#<SYSTEM-FUNCTION Cdr>"
}

func (*Cdr) funCall(arguments List) (T, error) {
	eval := NewEval()
	arg1, err := eval.Evaluate((arguments.(*Cons)).Car)
	if err != nil {
		return nil, err
	}
	if arg1 != nil {
		return (arg1.(*Cons)).Cdr, nil
	}
	return nil, nil
}

type FunCons struct{}

func (*FunCons) A() {}

func (*FunCons) String() string {
	return "#<SYSTEM-FUNCTION FunCons>"
}

func (*FunCons) funCall(arguments List) (T, error) {
	eval := NewEval()
	arg1, err := eval.Evaluate((arguments.(*Cons)).Car)
	if err != nil {
		return nil, err
	}
	arg2, err := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	if err != nil {
		return nil, err
	}
	return NewCons(arg1, arg2), nil
}

type Eq struct{}

func (*Eq) A() {}

func (*Eq) String() string {
	return "#<SYSTEM-FUNCTION Eq>"
}

func (*Eq) funCall(arguments List) (T, error) {
	eval := NewEval()
	arg1, err := eval.Evaluate((arguments.(*Cons)).Car)
	if err != nil {
		return nil, err
	}
	arg2, err := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	if err != nil {
		return nil, err
	}
	if reflect.DeepEqual(arg1, arg2) {
		return NewSymbol("T"), nil
	}
	return nil, nil
}

type Add struct{}

func (*Add) A() {}

func (*Add) String() string {
	return "#<SYSTEM-FUNCTION Add>"
}

func (*Add) funCall(arguments List) (T, error) {
	eval := NewEval()
	arg1, err := eval.Evaluate((arguments.(*Cons)).Car)
	if err != nil {
		return nil, err
	}
	arg2, err := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	if err != nil {
		return nil, err
	}
	return arg1.(*Integer).add(arg2.(*Integer)), nil
}

type Sub struct{}

func (*Sub) A() {}

func (*Sub) String() string {
	return "#<SYSTEM-FUNCTION Sub>"
}

func (*Sub) funCall(arguments List) (T, error) {
	eval := NewEval()
	arg1, err := eval.Evaluate((arguments.(*Cons)).Car)
	if err != nil {
		return nil, err
	}
	arg2, err := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	if err != nil {
		return nil, err
	}
	return arg1.(*Integer).sub(arg2.(*Integer)), nil
}

type Mul struct{}

func (*Mul) A() {}

func (*Mul) String() string {
	return "#<SYSTEM-FUNCTION Mul>"
}

func (*Mul) funCall(arguments List) (T, error) {
	eval := NewEval()
	arg1, err := eval.Evaluate((arguments.(*Cons)).Car)
	if err != nil {
		return nil, err
	}
	arg2, err := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	if err != nil {
		return nil, err
	}
	return arg1.(*Integer).mul(arg2.(*Integer)), nil
}

type Div struct{}

func (*Div) A() {}

func (*Div) String() string {
	return "#<SYSTEM-FUNCTION Div>"
}

func (*Div) funCall(arguments List) (T, error) {
	eval := NewEval()
	arg1, err := eval.Evaluate((arguments.(*Cons)).Car)
	if err != nil {
		return nil, err
	}
	arg2, err := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	if err != nil {
		return nil, err
	}
	return arg1.(*Integer).div(arg2.(*Integer)), nil
}

type Mod struct{}

func (*Mod) A() {}

func (*Mod) String() string {
	return "#<SYSTEM-FUNCTION Mod>"
}

func (*Mod) funCall(arguments List) (T, error) {
	eval := NewEval()
	arg1, err := eval.Evaluate((arguments.(*Cons)).Car)
	if err != nil {
		return nil, err
	}
	arg2, err := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	if err != nil {
		return nil, err
	}
	return arg1.(*Integer).mod(arg2.(*Integer)), nil
}

type Ge struct{}

func (*Ge) A() {}

func (*Ge) String() string {
	return "#<SYSTEM-FUNCTION Ge>"
}

func (*Ge) funCall(arguments List) (T, error) {
	eval := NewEval()
	arg1, err := eval.Evaluate((arguments.(*Cons)).Car)
	if err != nil {
		return nil, err
	}
	arg2, err := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	if err != nil {
		return nil, err
	}
	return arg1.(*Integer).ge(arg2.(*Integer)), nil
}

type Le struct{}

func (*Le) A() {}

func (*Le) String() string {
	return "#<SYSTEM-FUNCTION Le>"
}

func (*Le) funCall(arguments List) (T, error) {
	eval := NewEval()
	arg1, err := eval.Evaluate((arguments.(*Cons)).Car)
	if err != nil {
		return nil, err
	}
	arg2, err := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	if err != nil {
		return nil, err
	}
	return arg1.(*Integer).le(arg2.(*Integer)), nil
}

type Gt struct{}

func (*Gt) A() {}

func (*Gt) String() string {
	return "#<SYSTEM-FUNCTION Gt>"
}

func (*Gt) funCall(arguments List) (T, error) {
	eval := NewEval()
	arg1, err := eval.Evaluate((arguments.(*Cons)).Car)
	if err != nil {
		return nil, err
	}
	arg2, err := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	if err != nil {
		return nil, err
	}
	return arg1.(*Integer).gt(arg2.(*Integer)), nil
}

type Lt struct{}

func (*Lt) A() {}

func (*Lt) String() string {
	return "#<SYSTEM-FUNCTION Lt>"
}

func (*Lt) funCall(arguments List) (T, error) {
	eval := NewEval()
	arg1, err := eval.Evaluate((arguments.(*Cons)).Car)
	if err != nil {
		return nil, err
	}
	arg2, err := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	if err != nil {
		return nil, err
	}
	return arg1.(*Integer).lt(arg2.(*Integer)), nil
}

type NumberEqual struct{}

func (*NumberEqual) A() {}

func (*NumberEqual) String() string {
	return "#<SYSTEM-FUNCTION NumberEqual>"
}

func (*NumberEqual) funCall(arguments List) (T, error) {
	eval := NewEval()
	arg1, err := eval.Evaluate((arguments.(*Cons)).Car)
	if err != nil {
		return nil, err
	}
	arg2, err := eval.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	if err != nil {
		return nil, err
	}
	return arg1.(*Integer).numberEqual(arg2.(*Integer)), nil
}

type Setq struct{}

func (*Setq) A() {}

func (*Setq) String() string {
	return "#<SYSTEM-FUNCTION Setq>"
}

func (*Setq) funCall(arguments List) (T, error) {
	e := NewEval()
	arg1 := (arguments.(*Cons)).Car
	arg2, err := e.Evaluate(((arguments.(*Cons)).Cdr).(*Cons).Car)
	if err != nil {
		return nil, err
	}
	sym := arg1.(*Symbol)
	sym.Value, err = e.Evaluate(arg2)
	if err != nil {
		return nil, err
	}
	return sym, nil
}

type Quote struct{}

func (*Quote) A() {}

func (*Quote) String() string {
	return "#<SYSTEM-FUNCTION Quote>"
}

func (*Quote) funCall(arguments List) (T, error) {
	return arguments.(*Cons).Car, nil
}

type Defun struct{}

func (*Defun) A() {}

func (*Defun) String() string {
	return "#<SYSTEM-FUNCTION Defun>"
}

func (*Defun) funCall(arguments List) (T, error) {
	arg1 := (arguments.(*Cons)).Car
	args := (arguments.(*Cons)).Cdr
	fun := arg1.(*Symbol)
	lambda := NewCons(NewSymbol("LAMBDA"), args)
	fun.Function = lambda
	return fun, nil
}

type If struct{}

func (*If) A() {}

func (*If) String() string {
	return "#<SYSTEM-FUNCTION If>"
}

func (*If) funCall(arguments List) (T, error) {
	arg1 := (arguments.(*Cons)).Car
	args := (arguments.(*Cons)).Cdr
	arg2 := (args.(*Cons)).Car
	var arg3 T
	if (args.(*Cons)).Cdr != nil {
		arg3 = (((args.(*Cons)).Cdr).(*Cons)).Car
	}
	e := NewEval()
	v, err := e.Evaluate(arg1)
	if err != nil {
		return nil, err
	}
	if v != nil {
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

func (*SymbolFunction) funCall(arguments List) (T, error) {
	e := NewEval()
	arg1, err := e.Evaluate((arguments.(*Cons)).Car)
	if err != nil {
		return nil, err
	}
	return (arg1.(*Symbol)).Function, nil
}
