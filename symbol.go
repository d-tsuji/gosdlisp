package gosdlisp

var symbolTable map[string]*Symbol

func init() {
	symbolTable = make(map[string]*Symbol)

	// Initialization of Symbol T
	symbolT := &Symbol{name: "T"}
	symbolT.value = symbolT
	symbolTable["T"] = symbolT

	// Initialization of Symbol QUIT
	symbolQuit := &Symbol{name: "QUIT"}
	symbolTable["QUIT"] = symbolQuit
}

type Symbol struct {
	name     string
	value    T
	function T
}

func NewSymbol(name string) *Symbol {
	symbol, exists := symbolTable[name]
	if !exists {
		symbolTable[name] = &Symbol{name: name}
	}
	return symbol
}

func AddSymbolFunc(name string, f Function) {
	s := NewSymbol(name)
	s.function = f
	symbolTable[name] = s
}

func (s *Symbol) String() string {
	return s.name
}
