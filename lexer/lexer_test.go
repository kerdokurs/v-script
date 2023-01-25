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
		string first_name;
		int age;
		Metadata metadata;
	}`
	expected := []Token{
		{Keyword, "type"},
		{Ident, "Payload"},
		{LCurly, "{"},
		{Ident, "string"},
		{Ident, "first_name"},
		{Semi, ";"},
		{Ident, "int"},
		{Ident, "age"},
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

func TestLexer(t *testing.T) {
	suite.Run(t, &TypesTestSuite{})
}
