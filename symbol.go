package gosdlisp

var symbolTable map[string]*Symbol

func init() {
	symbolTable = make(map[string]*Symbol)

	// Initialization of Symbol T
	symbolT := &Symbol{Name: "T"}
	// TODO: require check symbolT.Value
	//symbolT.Value = symbolT
	symbolTable["T"] = symbolT

	// Initialization of Symbol QUIT
	symbolQuit := &Symbol{Name: "QUIT"}
	symbolTable["QUIT"] = symbolQuit

	// Initialization of system functions
	registSystemFunctions()
}

type Symbol struct {
	Name     string
	Value    T
	Function T
}

func NewSymbol(name string) *Symbol {
	_, exists := symbolTable[name]
	if !exists {
		symbolTable[name] = &Symbol{Name: name}
	}
	return symbolTable[name]
}

func AddSymbolFunc(name string, f Function) {
	s := NewSymbol(name)
	s.Function = f
	symbolTable[name] = s
}

func (s Symbol) String() string {
	return s.Name
}

func (s Symbol) A() {}
