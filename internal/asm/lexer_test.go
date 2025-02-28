package asm

import (
	"fmt"
	"testing"
)

func TestLexer_next(t *testing.T) {
	var assertions = []struct {
		src    string
		tokens []token
	}{
		{
			src: "@17\n",
			tokens: []token{
				{
					variant: at,
					literal: "@",
				},
				{
					variant: integer,
					literal: "17",
				},
				{
					variant: linefeed,
					literal: "\n",
				},
			},
		},
		{
			src: "//this is a comment\n@2000",
			tokens: []token{
				{
					variant: at,
					literal: "@",
				},
				{
					variant: integer,
					literal: "2000",
				},
			},
		},
		{
			src: "A;D;M",
			tokens: []token{
				{
					variant: identifier,
					literal: "A",
				},
				{
					variant: semicolon,
					literal: ";",
				},
				{
					variant: identifier,
					literal: "D",
				},
				{
					variant: semicolon,
					literal: ";",
				},
				{
					variant: identifier,
					literal: "M",
				},
			},
		},
		{
			src: "D;JGT",
			tokens: []token{
				{
					variant: identifier,
					literal: "D",
				},
				{
					variant: semicolon,
					literal: ";",
				},
				{
					variant: jgt,
					literal: "JGT",
				},
			},
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given src %v", a.src), func(t *testing.T) {
			l := lexer{src: a.src}
			i := 0
			var err error
			var tok token
			for tok, err = l.next(); tok.variant != eof && err == nil; {
				if a.tokens[i].variant != tok.variant {
					t.Errorf("expected tok variant %v but got %v", a.tokens[i].variant, tok.variant)
				}
				if a.tokens[i].literal != tok.literal {
					t.Errorf("expected tok variant %v but got %v", a.tokens[i].variant, tok.variant)
				}
				tok, err = l.next()
				i += 1
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
