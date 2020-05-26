package gosdlisp

import (
	"fmt"
)

var (
	defaultMaxStackSize = 65536
)

type Eval struct {
	stack  []T
	stackP int
}

func NewEval() *Eval {
	return &Eval{
		stack:  make([]T, defaultMaxStackSize),
		stackP: 0,
	}
}

func (e Eval) Evaluate(form T) (T, error) {
	// Evaluation of Symbols
	f, ok := form.(*Symbol)
	if ok {
		symbolValue := f.Value
		if symbolValue == nil {
			return nil, fmt.Errorf("unbound variable: %v", form)
		}
		return symbolValue, nil
	}

	// Evaluating when Atom is not a symbol
	if form == nil {
		return form, nil
	}
	if _, ok := form.(Atom); ok {
		return form, nil
	}

	// Evaluation of a list (evaluation of a function)
	car := (form.(*Cons)).Car
	_, ok = car.(*Symbol)
	if !ok {
		return nil, fmt.Errorf("not a symbol: %v", car)
	}
	fun := car.(*Symbol).Function
	if fun == nil {
		return nil, fmt.Errorf("undefined function: %v", car)
	}
	switch fun.(type) {
	case Function:
		argumentList := (form.(*Cons)).Cdr
		return fun.(Function).funCall(argumentList.(List))
	case *Cons:
		cdr := ((fun.(*Cons)).Cdr).(*Cons)
		lambdaList := cdr.Car
		body := (cdr.Cdr).(*Cons)
		if lambdaList == nil {
			return e.evalBody(body)
		}
		return e.bindEvalBody(lambdaList.(*Cons), body, ((form.(*Cons)).Cdr).(*Cons))
	default:
		return nil, fmt.Errorf("not a function: %v, type: %T", fun, fun)
	}
}

func (e Eval) bindEvalBody(lambda, body, form *Cons) (T, error) {
	// (1) Argument evaluation in a pre-bound environment
	// (using a temporary stack to store evaluated values)
	oldStackP := e.stackP
	for {
		ret, err := e.Evaluate(form.Car)
		if err != nil {
			return nil, err
		}
		e.stack[e.stackP] = ret
		e.stackP++
		if form.Cdr == nil {
			break
		}
		form = form.Cdr.(*Cons)
	}

	// (2) Binding (move the past value of the symbol to the stack
	// and put the value evaluated in (1) into the symbol (swap))
	argList := lambda
	sp := oldStackP
	for {
		sym := (argList.Car).(*Symbol)
		sym.Value, e.stack[sp] = e.stack[sp], sym.Value
		sp++
		if argList.Cdr == nil {
			break
		}
		argList = (argList.Cdr).(*Cons)
	}

	// Evaluate body
	ret, err := e.evalBody(body)
	if err != nil {
		return nil, err
	}

	// Return the previous value from the stack
	argList = lambda
	e.stackP = oldStackP
	for {
		sym := (argList.Car).(*Symbol)
		sym.Value = e.stack[oldStackP]
		oldStackP++
		if argList.Cdr == nil {
			break
		}
		argList = (argList.Cdr).(*Cons)
	}

	// return value
	return ret, nil
}

func (e Eval) evalBody(body *Cons) (T, error) {
	var ret T
	var err error
	for {
		ret, err = e.Evaluate(body.Car)
		if err != nil {
			return nil, err
		}
		if body.Cdr == nil {
			break
		}
		body = (body.Cdr).(*Cons)
	}
	return ret, nil
}
