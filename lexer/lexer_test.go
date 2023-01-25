package lexer

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TypesTestSuite struct {
	suite.Suite
}

func (s *TypesTestSuite) TestValid() {
	input := `type Payload {
		string item_name;
		float weight;
		bool is_fragile;
		Metadata metadata;
	}`
	expected := []Token{
		{Keyword, "type"},
		{Ident, "Payload"},
		{LCurly, "{"},
		{Ident, "string"},
		{Ident, "item_name"},
		{Semi, ";"},
		{Ident, "float"},
		{Ident, "weight"},
		{Semi, ";"},
		{Ident, "bool"},
		{Ident, "is_fragile"},
		{Semi, ";"},
		{Ident, "Metadata"},
		{Ident, "metadata"},
		{Semi, ";"},
		{RCurly, "}"},
		{EOF, nil},
	}
	lexer := New(strings.NewReader(input))
	tokens, err := lexer.Lex()
	if !s.Nil(err) {
		s.FailNow("got error")
	}
	s.Equalf(len(expected), len(tokens), "expected %d tokens, got %d", len(expected), len(tokens))
	for i, expectedToken := range expected {
		gotToken := tokens[i]
		if !s.Equal(expectedToken, gotToken) {
			s.FailNow("tokens did not match")
		}
	}
}

type ValidationsTestSuite struct {
	suite.Suite
}

func (s *ValidationsTestSuite) TestValid() {
	input := `validate on Payload (version) if version > 1 and weight > 5.5 {
		allow if item_name is not empty -> "Item name is empty";
		allow if is_fragile is false;
		allow if validate metadata (version, 0);
	}`
	expected := []Token{
		{Keyword, "validate"},
		{Keyword, "on"},
		{Ident, "Payload"},
		{LParen, "("},
		{Ident, "version"},
		{RParen, ")"},
		{Keyword, "if"},
		{Ident, "version"},
		{Operator, ">"},
		{Int, "1"},
		{Keyword, "and"},
		{Ident, "weight"},
		{Operator, ">"},
		{Float, "5.5"},
		{LCurly, "{"},
		{Keyword, "allow"},
		{Keyword, "if"},
		{Ident, "item_name"},
		{Keyword, "is"},
		{Keyword, "not"},
		{Keyword, "empty"},
		{Operator, "->"},
		{String, "Item name is empty"},
		{Semi, ";"},
		{Keyword, "allow"},
		{Keyword, "if"},
		{Ident, "is_fragile"},
		{Keyword, "is"},
		{Bool, "false"},
		{Semi, ";"},
		{Keyword, "allow"},
		{Keyword, "if"},
		{Keyword, "validate"},
		{Ident, "metadata"},
		{LParen, "("},
		{Ident, "version"},
		{Comma, ","},
		{Int, "0"},
		{RParen, ")"},
		{Semi, ";"},
		{RCurly, "}"},
		{EOF, nil},
	}
	lexer := New(strings.NewReader(input))
	tokens, err := lexer.Lex()
	if !s.Nil(err) {
		s.FailNow("got error")
	}
	s.Equalf(len(expected), len(tokens), "expected %d tokens, got %d", len(expected), len(tokens))
	for i, expectedToken := range expected {
		gotToken := tokens[i]
		if !s.Equal(expectedToken, gotToken) {
			s.FailNow("tokens did not match")
		}
	}
}

func (s *ValidationsTestSuite) TestInvalid() {
	input1 := `"Test`
	lexer := New(strings.NewReader(input1))
	tokens, err := lexer.Lex()
	s.NotNil(err)
	s.Nil(tokens)

	input2 := `-`
	lexer = New(strings.NewReader(input2))
	tokens, err = lexer.Lex()
	s.NotNil(err)
	s.Nil(tokens)

	input3 := `00`
	lexer = New(strings.NewReader(input3))
	tokens, err = lexer.Lex()
	s.NotNil(err)
	s.Nil(tokens)
}

func TestLexer(t *testing.T) {
	suite.Run(t, &TypesTestSuite{})
	suite.Run(t, &ValidationsTestSuite{})
}
