package lexer

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"

	"github.com/kerdokurs/v-script/utils"
)

var keywords = []string{
	"type",
	"validate", "on",
	"allow", "disallow",
	"if",
	"is", "not",
	"empty",
	"and", "or",
}

var singleChars = map[rune]TokenType{
	'{': LCurly,
	'}': RCurly,
	';': Semi,
	'(': LParen,
	')': RParen,
	',': Comma,
}

var operators = map[rune][]rune{
	'-': {0, '>'},
	'+': {0},
	'>': {0, '='},
	'<': {0, '='},
	'=': {0},
	'!': {'='},
	'%': {0},
}

type Lexer struct {
	r *bufio.Reader
}

func New(reader io.Reader) *Lexer {
	return &Lexer{
		r: bufio.NewReader(reader),
	}
}

func (l *Lexer) Lex() ([]Token, error) {
	tokens := make([]Token, 0)

	for {
		tok, err := l.nextToken()
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, tok)
		if tok.Type == EOF {
			break
		}
	}

	return tokens, nil
}

func (l *Lexer) nextToken() (Token, error) {
	for {
		rs, err := l.r.Peek(1)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return Token{Type: EOF}, nil
			}

			return Token{}, err
		}

		r := rune(rs[0])
		if unicode.IsSpace(r) {
			l.r.ReadRune()
			continue
		} else if unicode.IsDigit(r) {
			return l.readNumericLiteral()
		} else if unicode.IsLetter(r) || r == '_' {
			return l.readIdentOrKeyword()
		} else if t, ok := singleChars[r]; ok {
			return l.readSingleChar(t)
		} else if _, ok := operators[r]; ok {
			return l.readOperator()
		} else if r == '"' {
			return l.readStringLiteral()
		}

		return Token{}, fmt.Errorf("invalid token: %s", string(r))
	}
}

func (l *Lexer) readNumericLiteral() (Token, error) {
	var tokenType TokenType = Int
	var sb strings.Builder

	for {
		r, _, err := l.r.ReadRune()
		if err != nil {
			return Token{}, err
		}

		if r == '.' {
			tokenType = Float
		} else if !unicode.IsDigit(r) {
			err = l.r.UnreadRune()
			if err != nil {
				return Token{}, fmt.Errorf("error unreading from input: %v", err)
			}
			break
		}

		sb.WriteRune(r)
	}

	data := sb.String()

	if len(data) >= 2 && data[0] == '0' && data[1] == '0' {
		return Token{}, fmt.Errorf("invalid number: %v", data)
	}

	return Token{
		Type: tokenType,
		Data: data,
	}, nil
}

func (l *Lexer) readIdentOrKeyword() (Token, error) {
	var sb strings.Builder

	for {
		r, _, err := l.r.ReadRune()
		if err != nil {
			return Token{}, err
		}

		if !(unicode.IsDigit(r) || unicode.IsLetter(r) || r == '_') {
			err = l.r.UnreadRune()
			if err != nil {
				return Token{}, fmt.Errorf("error unreading from input: %v", err)
			}
			break
		}

		sb.WriteRune(r)
	}

	var tokenType TokenType = Ident
	data := sb.String()

	if data == "true" || data == "false" {
		tokenType = Bool
	} else if utils.Contains(keywords, data) {
		tokenType = Keyword
	}

	return Token{
		Type: tokenType,
		Data: data,
	}, nil
}

func (l *Lexer) readSingleChar(tokenType TokenType) (Token, error) {
	r, _, err := l.r.ReadRune()
	if err != nil {
		return Token{}, err
	}

	return Token{
		Type: tokenType,
		Data: string(r),
	}, nil
}

func (l *Lexer) readOperator() (Token, error) {
	var sb strings.Builder

	r, _, err := l.r.ReadRune()
	if err != nil {
		return Token{}, err
	}

	sb.WriteRune(r)

	extended := operators[r]

	if len(extended) > 1 {
		if r, _, err = l.r.ReadRune(); err != nil {
			return Token{}, err
		}

		if utils.Contains(extended, r) {
			sb.WriteRune(r)
		} else {
			if err = l.r.UnreadRune(); err != nil {
				return Token{}, err
			}
		}
	}

	return Token{
		Type: Operator,
		Data: sb.String(),
	}, nil
}

func (l *Lexer) readStringLiteral() (Token, error) {
	if _, _, err := l.r.ReadRune(); err != nil {
		return Token{}, err
	}

	var sb strings.Builder

	for {
		r, _, err := l.r.ReadRune()
		if err != nil {
			return Token{}, err
		}

		if r == '"' {
			break
		}

		sb.WriteRune(r)
	}

	return Token{
		Type: String,
		Data: sb.String(),
	}, nil
}
