package lexer

import "fmt"

type TokenType int

func (t TokenType) String() string {
	return tokenNames[t]
}

const (
	EOF TokenType = iota

	Int
	Float

	Bool

	Ident
	Keyword

	LCurly
	RCurly

	Semi
)

var tokenNames = map[TokenType]string{
	EOF:     "EOF",
	Int:     "Int",
	Float:   "Float",
	Bool:    "Bool",
	Ident:   "Ident",
	Keyword: "Keyword",
	LCurly:  "LCurly",
	RCurly:  "RCurly",
	Semi:    "Semi",
}

type Token struct {
	Type TokenType
	Data any
}

func (t Token) String() string {
	return fmt.Sprintf("%s (%v)", t.Type, t.Data)
}
