package gosdlisp

type Null struct {
	List
}

func (n Null) String() string {
	return "NIL"
}
