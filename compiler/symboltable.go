package compiler

type SymbolScope string

const (
	GlobalScope   SymbolScope = "GLOBAL"
	LocalScope    SymbolScope = "LOCAL"
	BuiltinScope  SymbolScope = "BUILTIN"
	FreeScope     SymbolScope = "FREE"
	FunctionScope SymbolScope = "FUNCTION"
)

type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

type SymbolTable struct {
	Outer          *SymbolTable
	store          map[string]Symbol
	numDefinitions int
	FreeSymbols    []Symbol
}

func NewEnclosedSymbolTable(outer *SymbolTable) *SymbolTable {
	s := NewSymbolTable()
	s.Outer = outer
	return s
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		store:       make(map[string]Symbol),
		FreeSymbols: []Symbol{},
	}
}

func (s *SymbolTable) Define(name string) Symbol {
	symbol := Symbol{
		Name:  name,
		Index: s.numDefinitions,
	}

	if s.Outer == nil {
		symbol.Scope = GlobalScope
	} else {
		symbol.Scope = LocalScope
	}

	s.store[name] = symbol
	s.numDefinitions++
	return symbol
}

func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.store[name]
	if !ok && s.Outer != nil {
		obj, ok = s.Outer.Resolve(name)
		if !ok {
			return obj, ok
		}

		if obj.Scope == GlobalScope || obj.Scope == BuiltinScope {
			return obj, ok
		}

		free := s.defineFree(obj)
		return free, true
	}
	return obj, ok
}

func (s *SymbolTable) DefineBuiltin(index int, name string) Symbol {
	symbol := Symbol{
		Name:  name,
		Scope: BuiltinScope,
		Index: index,
	}
	s.store[name] = symbol
	return symbol
}

func (s *SymbolTable) defineFree(original Symbol) Symbol {
	s.FreeSymbols = append(s.FreeSymbols, original)

	symbol := Symbol{
		Name:  original.Name,
		Scope: FreeScope,
		Index: len(s.FreeSymbols) - 1,
	}

	s.store[original.Name] = symbol

	return symbol
}

func (s *SymbolTable) DefineFunctionName(name string) Symbol {
	symbol := Symbol{
		Name:  name,
		Scope: FunctionScope,
		Index: 0,
	}
	s.store[name] = symbol
	return symbol
}
