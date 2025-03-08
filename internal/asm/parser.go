package asm

import (
	"fmt"
)

type instruction interface {
	Literal() string
}

type load struct {
	value token
}

func (l load) Literal() string {
	return fmt.Sprintf("@%s", l.value.literal)
}

type compute struct {
	dest *token
	comp string
	jump *token
}

func (c compute) Literal() string {
	var str string
	if c.dest != nil {
		str += fmt.Sprintf("%s=", c.dest.literal)
	}
	str += c.comp
	if c.jump != nil {
		str += fmt.Sprintf(";%s", c.jump.literal)
	}
	return str
}

type label struct {
	value token
}

func (l label) Literal() string {
	return fmt.Sprintf("(%s)", l.value.literal)
}

type parser struct {
	lexer lexer
}

func (p *parser) next() (instruction, error) {
	tok, err := p.lexer.peek()
	if err != nil {
		return nil, err
	}
	switch tok.variant {
	case eof:
		return nil, nil
	case at:
		return p.a()
	case lparen:
		return p.label()
	default:
		return p.c()
	}
}

func (p *parser) a() (load, error) {
	if _, err := p.want(at); err != nil {
		return load{}, err
	}
	tok, err := p.lexer.next()
	if err != nil {
		return load{}, err
	}
	if tok.variant != integer && tok.variant != identifier {
		return load{}, fmt.Errorf("unexpected token for A-instruction '%v'", tok)
	}
	if err := p.seek(linefeed); err != nil {
		return load{}, err
	}
	return load{value: tok}, nil
}

func (p *parser) label() (label, error) {
	if _, err := p.want(lparen); err != nil {
		return label{}, err
	}
	name, err := p.want(identifier)
	if err != nil {
		return label{}, err
	}
	if _, err := p.want(rparen); err != nil {
		return label{}, err
	}
	if err := p.seek(linefeed); err != nil {
		return label{}, err
	}
	return label{value: name}, nil
}

func (p *parser) c() (comp compute, err error) {
	tok, err := p.lexer.next()
	if err != nil {
		return compute{}, err
	}
	next, err := p.lexer.peek()
	if err != nil {
		return compute{}, err
	}
	if next.variant == equals {
		_, _ = p.want(equals)
		comp.dest = &token{
			variant: tok.variant,
			literal: tok.literal,
		}
		// Fetch the next token for parsing the compute field
		tok, err = p.lexer.next()
		if err != nil {
			return compute{}, err
		}
	}
	for tok.variant != semicolon && tok.variant != linefeed {
		comp.comp += tok.literal
		tok, err = p.lexer.next()
		if err != nil {
			return compute{}, err
		}
	}
	if tok.variant == semicolon {
		jmp, err := p.lexer.next()
		if err != nil {
			return compute{}, err
		}
		comp.jump = &token{
			variant: jmp.variant,
			literal: jmp.literal,
		}
		if err := p.seek(linefeed); err != nil {
			return compute{}, err
		}
	}
	return comp, nil
}

func (p *parser) seek(v variant) error {
	tok, err := p.lexer.next()
	if err != nil {
		return err
	}
	for tok.variant != v {
		tok, err = p.lexer.next()
		if err != nil {
			return err
		}
		if tok.variant == eof && v != eof {
			return fmt.Errorf("could not find token of variant '%v' before eof", v)
		}
	}
	return nil
}

func (p *parser) want(v variant) (token, error) {
	tok, err := p.lexer.next()
	if err != nil {
		return token{}, err
	}
	if tok.variant != v {
		return token{}, fmt.Errorf("expected '%v' token but found '%v'", v, tok.variant)
	}
	return tok, nil
}
