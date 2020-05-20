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
		// TODO
	default:
		log.Fatalf("Not a Function: %v", fun)
	}
	// TODO: will be deleted.
	return nil
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
