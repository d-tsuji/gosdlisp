package gosdlisp

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
