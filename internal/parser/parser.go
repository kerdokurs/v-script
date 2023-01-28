package parser

import (
	"fmt"
	"io"

	"github.com/kerdokurs/v-script/internal/ast"
	"github.com/kerdokurs/v-script/internal/lexer"
)

type Parser struct {
	lexer *lexer.Lexer

	tokens       []lexer.Token
	currentIndex int
}

func NewParser(input io.Reader) *Parser {
	return &Parser{
		lexer: lexer.New(input),
	}
}

func (p *Parser) Parse() (*ast.ScriptStmt, error) {
	tokens, err := p.lexer.Lex()
	if err != nil {
		return nil, err
	}
	p.tokens = tokens

	scriptStmt := ast.ScriptStmt{
		TypeDecls:      make([]ast.TypeDecl, 0),
		ValidatorDecls: make([]ast.ValidatorDecl, 0),
	}

	for {
		if p.peek().Type == lexer.EOF {
			break
		}

		decl, err := p.decl()
		if err != nil {
			return nil, err
		}

		switch val := decl.(type) {
		case ast.TypeDecl:
			scriptStmt.TypeDecls = append(scriptStmt.TypeDecls, val)
		case ast.ValidatorDecl:
			scriptStmt.ValidatorDecls = append(scriptStmt.ValidatorDecls, val)
		}
	}

	return &scriptStmt, err
}

func (p *Parser) decl() (ast.Decl, error) {
	if err := p.expectType(lexer.Keyword); err != nil {
		return nil, err
	}

	tok := p.consume()
	switch tok.Data {
	case "type":
		return p.typeDecl()
	case "validate":
		fmt.Println("validate")
	}

	return nil, fmt.Errorf("invalid keyword: %s\n", tok.Data)
}

func (p *Parser) typeDecl() (ast.TypeDecl, error) {
	var typeDecl ast.TypeDecl

	var err error
	if err = p.expectType(lexer.Ident); err != nil {
		return typeDecl, err
	}

	nameTok := p.consume()
	typeDecl.Name = nameTok.Data.(string)

	if _, err = p.expectAndConsume(lexer.LCurly); err != nil {
		return ast.TypeDecl{}, err
	}

	if typeDecl.Fields, err = p.fieldDecls(); err != nil {
		return ast.TypeDecl{}, err
	}

	if _, err = p.expectAndConsume(lexer.RCurly); err != nil {
		return ast.TypeDecl{}, err
	}

	return typeDecl, nil
}

func (p *Parser) fieldDecls() ([]ast.FieldDecl, error) {
	fields := make([]ast.FieldDecl, 0)

	for {
		if p.peek().Type == lexer.RCurly {
			break
		}

		if fieldDecl, err := p.fieldDecl(); err != nil {
			return nil, err
		} else {
			fields = append(fields, fieldDecl)
		}
	}

	return fields, nil
}

func (p *Parser) fieldDecl() (ast.FieldDecl, error) {
	typeTok, err := p.expectAndConsume(lexer.Ident)
	if err != nil {
		return ast.FieldDecl{}, err
	}

	nameTok, err := p.expectAndConsume(lexer.Ident)
	if err != nil {
		return ast.FieldDecl{}, err
	}

	if _, err := p.expectAndConsume(lexer.Semi); err != nil {
		return ast.FieldDecl{}, err
	}

	return ast.FieldDecl{
		Name: nameTok.Data.(string),
		Type: typeTok.Data.(string),
	}, nil
}

func (p *Parser) expectType(tokenType lexer.TokenType) error {
	if p.peek().Type != tokenType {
		return fmt.Errorf("expected token %s, got %s", tokenType.String(), p.peek().Type.String())
	}

	return nil
}

func (p *Parser) expectAndConsume(tokenType lexer.TokenType) (*lexer.Token, error) {
	if err := p.expectType(tokenType); err != nil {
		return nil, err
	}

	return p.consume(), nil
}

func (p *Parser) peek() *lexer.Token {
	return &p.tokens[p.currentIndex]
}

func (p *Parser) consume() *lexer.Token {
	p.currentIndex += 1
	return &p.tokens[p.currentIndex-1]
}
