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
	String
	Ident
	Keyword
	LCurly
	RCurly
	LParen
	RParen
	Semi
	Comma
	Operator
)

var tokenNames = map[TokenType]string{
	EOF:      "EOF",
	Int:      "Int",
	Float:    "Float",
	Bool:     "Bool",
	String:   "String",
	Ident:    "Ident",
	Keyword:  "Keyword",
	LCurly:   "LCurly",
	RCurly:   "RCurly",
	LParen:   "LParen",
	RParen:   "RParen",
	Semi:     "Semi",
	Comma:    "Comma",
	Operator: "Operator",
}

type Token struct {
	Type TokenType
	Data any
}

func (t Token) String() string {
	return fmt.Sprintf("%s (%v)", t.Type, t.Data)
}
