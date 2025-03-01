package asm

import (
	"errors"
	"fmt"
)

var (
	symbols = map[uint8]variant{
		'\n': linefeed,
		'@':  at,
		'-':  minus,
		'+':  plus,
		'&':  and,
		'|':  or,
		';':  semicolon,
		'(':  lparen,
		')':  rparen,
		'=':  equals,
	}
	keywords = map[string]variant{
		"JGT": jgt,
		"JEQ": jeq,
		"JGE": jge,
		"JLT": jlt,
		"JNE": jne,
		"JLE": jle,
		"JMP": jmp,
	}
)

const (
	eof variant = iota
	at
	minus
	plus
	and
	or
	equals
	identifier
	integer
	lparen
	rparen
	semicolon
	jgt
	jeq
	jge
	jlt
	jne
	jle
	jmp
	linefeed
)

type variant int

type token struct {
	variant variant
	literal string
}

type lexer struct {
	src    string
	cursor int
}

// next returns the next token that can be extracted from the underlying source code from the current cursor position
// and places cursor at the next character after all characters belonging to the most recently processed token. Calling
// next after the cursor has reached the end of the underlying source code is safe and will only result in an eof token
// being returned together with a nil error.
func (l *lexer) next() (token, error) {
	if l.cursor >= len(l.src) {
		return token{
			variant: eof,
		}, nil
	}
	char := l.src[l.cursor]
	if symbol, ok := symbols[char]; ok {
		l.cursor++
		return token{
			variant: symbol,
			literal: string(char),
		}, nil
	}
	if char == '/' {
		return l.comment()
	}
	if !alphanumerical(char) {
		// At this point all the valid non-alphanumerical characters should have been processed. Therefore, if the
		// current character is not alphanumerical it should be considered invalid.
		return token{}, fmt.Errorf("invalid token '%s'", string(char))
	}
	// Identifiers cannot start with a digit, therefore we must first check if the current character is an integer to
	// decide whether to regard this as an integer literal
	if numerical(char) {
		return token{
			variant: integer,
			literal: l.literal(numerical),
		}, nil
	}
	literal := l.literal(alphanumerical)
	keyword, ok := keywords[literal]
	if ok {
		return token{
			variant: keyword,
			literal: literal,
		}, nil
	}
	return token{
		variant: identifier,
		literal: literal,
	}, nil
}

func (l *lexer) seek(c uint8) error {
	for ; l.cursor < len(l.src) && l.src[l.cursor] != c; l.cursor++ {
	}
	if l.cursor == len(l.src) {
		return fmt.Errorf("character '%s' not found", string(c))
	}
	return nil
}

func (l *lexer) literal(fn func(uint8) bool) string {
	literal := ""
	for ; l.cursor < len(l.src) && fn(l.src[l.cursor]); l.cursor++ {
		literal += string(l.src[l.cursor])
	}
	return literal
}

func (l *lexer) comment() (token, error) {
	if l.src[l.cursor+1] != '/' {
		return token{}, errors.New("invalid token '/'")
	}
	l.cursor += 2
	if err := l.seek('\n'); err != nil {
		return token{}, err
	}
	l.cursor += 1
	return l.next()
}

func alphanumerical(c uint8) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func numerical(c uint8) bool {
	return c >= '0' && c <= '9'
}
