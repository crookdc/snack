package asm

import (
	"errors"
	"github.com/crookdc/nand2tetris/internal/chip"
)

func Assemble(src string) ([][16]chip.Signal, error) {
	lex := lexer{src: src}
	var bin [][16]chip.Signal
	for lex.more() {
		tok, err := lex.next()
		if err != nil {
			return nil, err
		}
		switch tok.variant {
		case eof:
			return bin, nil
		}
	}
	// The token stream **should** end with an EOF token, that case is handled in the above loop. If we exit the loop
	// due to exhaustion of the lexers input without entering there then something is wrong.
	return nil, errors.New("unexpected end of file")
}
