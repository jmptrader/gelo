package gelo

import "bytes"

var Null Symbol = _iSymbol([]byte("")) //TODO replace with 0

func StrToSym(s string) Symbol {
	return _dSymbol([]byte(s))
}

func BytesToSym(s []byte) Symbol {
	return _dSymbol(s)
}

func RuneToSym(s []rune) Symbol {
	return _dSymbol([]byte(string(s)))
}

func StrEqualsSym(a string, b Symbol) bool {
	return bytes.Equal([]byte(a), b.Bytes())
}

//returns true if w is the Symbol ""
func IsNullString(w Word) bool {
	s, ok := w.(Symbol)
	if !ok {
		return false
	}
	return len(s.Bytes()) == 0
}

func intern(s []byte) Symbol {
	if len(s) == 0 {
		return Null
	}
	return _iSymbol(s)
}

func interns(s string) Symbol {
	return intern([]byte(s))
}

//noncopying for internal use when we absolutely know that it will not end
//in horror
func bytesof(s Symbol) []byte {
	if s.interned() {
		return []byte(s.(_iSymbol))
	}
	return []byte(s.(_dSymbol))
}

func stringof(s Symbol) string {
	return string(bytesof(s))
}

//dynamic, ie not interned, symbols
type _dSymbol []byte

func (s _dSymbol) Bytes() []byte {
	return dup([]byte(s))
}

func (s _dSymbol) String() string {
	return string([]byte(s))
}

func (s _dSymbol) Runes() []rune {
	return []rune(string([]byte(s)))
}

func (_dSymbol) interned() bool {
	return false
}

func (s _dSymbol) Ser() Symbol {
	return s
}

func (s _dSymbol) Equals(w Word) bool {
	return bytes.Equal([]byte(s), w.Ser().Bytes())
}

func (s _dSymbol) Copy() Word {
	sb := []byte(s)
	if len(sb) == 0 {
		return Null
	}
	return _dSymbol(dup(sb))
}

func (s _dSymbol) DeepCopy() Word {
	return s.Copy()
}

func (_dSymbol) Type() Symbol {
	return interns("*SYMBOL*")
}

//This type is to be replaced by an index into an intern pool
type _iSymbol []byte

func (s _iSymbol) Bytes() []byte {
	return dup([]byte(s))
}

func (s _iSymbol) String() string {
	return string([]byte(s))
}

func (s _iSymbol) Runes() []rune {
	return []rune(string([]byte(s)))
}

func (_iSymbol) interned() bool {
	return true
}

func (s _iSymbol) Ser() Symbol {
	return s
}

func (s _iSymbol) Equals(w Word) bool {
	return bytes.Equal([]byte(s), w.Ser().Bytes())
}

func (s _iSymbol) Copy() Word {
	sb := []byte(s)
	if len(s) == 0 {
		return Null
	}
	return _dSymbol(dup(sb)) //not a typo
}

func (s _iSymbol) DeepCopy() Word {
	return s.Copy()
}

func (_iSymbol) Type() Symbol {
	return interns("*SYMBOL*")
}
