package gosdlisp

import "log"

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

func (e Eval) Evaluate(form T) T {
	f, ok := form.(*Symbol)
	if ok {
		symbolValue := f.Value
		if symbolValue == nil {
			log.Fatalf("Unbound Variable Error: %v", symbolValue)
		}
		return symbolValue
	}

	if form == nil {
		return form
	}
	if _, ok := form.(Atom); ok {
		return form
	}
	car := form.(*Cons).Car
	_, ok = car.(*Symbol)
	if !ok {
		log.Fatalf("Not a Symbol: %v", car)
	}
	fun := car.(*Symbol).Function
	if fun == nil {
		log.Fatalf("Undefined Function Error: %v", car)
	}
	switch fun.(type) {
	case Function:
		argumentList := (form.(*Cons)).Cdr
		return fun.(Function).funCall(argumentList.(List))
	case Cons:
		cdr := ((fun.(*Cons)).Cdr).(*Cons)
		lambdaList := cdr.Car
		body := (cdr.Cdr).(*Cons)
		if lambdaList == nil {
			return e.evalBody(body)
		}
		return e.bindEvalBody(lambdaList.(*Cons), body, ((form.(*Cons)).Cdr).(*Cons))
	default:
		log.Fatalf("Not a Function: %v", fun)
	}
	return nil
}

func (e Eval) bindEvalBody(lambda, body, form *Cons) T {
	// (1) Argument evaluation in a pre-bound environment
	// (using a temporary stack to store evaluated values)
	oldStackP := e.stackP
	for {
		ret := e.Evaluate(form.Car)
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
		sym := (argList.Cdr).(*Symbol)
		swap := sym.Value
		sym.Value = e.stack[sp]
		e.stack[sp] = swap
		sp++
		if argList.Cdr == nil {
			break
		}
		argList = (argList.Cdr).(*Cons)
	}

	// Evaluate body
	ret := e.evalBody(body)

	// Return the previous value from the stack
	argList = lambda
	e.stackP = oldStackP
	for {
		sym := (argList.Cdr).(*Symbol)
		sym.Value = e.stack[oldStackP]
		oldStackP++
		if argList.Cdr == nil {
			break
		}
		argList = (argList.Cdr).(*Cons)
	}

	// return value
	return ret
}

func (e Eval) evalBody(body *Cons) T {
	var ret T
	for {
		ret = e.Evaluate(body.Car)
		if body.Cdr == nil {
			break
		}
		body = (body.Cdr).(*Cons)
	}
	return ret
}
