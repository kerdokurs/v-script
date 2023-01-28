package ast

import (
	"fmt"
	"strings"

	"github.com/kerdokurs/v-script/utils"
)

type node interface {
	fmt.Stringer
}

type expr interface {
	node
	exprNode()
}

type stmt interface {
	node
	stmtNode()
}

type Decl interface {
	node
	declNode()
}

type ScriptStmt struct {
	TypeDecls      []TypeDecl
	ValidatorDecls []ValidatorDecl
}

func (s ScriptStmt) String() string {
	typesStr := strings.Join(utils.Map(s.TypeDecls, TypeDecl.String), "\n\t\t")
	validatorsStr := strings.Join(utils.Map(s.ValidatorDecls, ValidatorDecl.String), "\n\t\t")
	return fmt.Sprintf("Script(\n\t%s\n\t%s\n)", typesStr, validatorsStr)
}

func (s ScriptStmt) stmtNode() {}

type TypeDecl struct {
	Name   string
	Fields []FieldDecl
}

func (d TypeDecl) String() string {
	fieldsStr := strings.Join(utils.Map(d.Fields, FieldDecl.String), "\n\t\t")
	return fmt.Sprintf("TypeDecl(\n\t\t%s,\n\t\t%s\n\t)", d.Name, fieldsStr)
}

func (d TypeDecl) declNode() {}

type FieldDecl struct {
	Type string
	Name string
}

func (d FieldDecl) String() string {
	return fmt.Sprintf("FieldDecl(%s, %s)", d.Type, d.Name)
}

func (d FieldDecl) declNode() {}

type ValidatorDecl struct {
}

func (d ValidatorDecl) String() string {
	return "ValidatorDecl()"
}

func (d ValidatorDecl) declNode() {}
