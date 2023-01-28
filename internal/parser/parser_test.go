package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/kerdokurs/v-script/internal/ast"
	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	input := `type Payload {
    string name;
    int age;
  }`
	expected := ast.ScriptStmt{
		TypeDecls: []ast.TypeDecl{
			{
				Name: "Payload",
				Fields: []ast.FieldDecl{
					{"string", "name"},
					{"int", "age"},
				},
			},
		},
		ValidatorDecls: []ast.ValidatorDecl{},
	}
	parser := NewParser(strings.NewReader(input))
	stmt, err := parser.Parse()
	assert.Nil(t, err)
	assert.NotNil(t, stmt)

	fmt.Println(expected)
	assert.Equal(t, expected.String(), stmt.String())
}
