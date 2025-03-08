package asm

import (
	"fmt"
	"github.com/crookdc/nand2tetris/internal/chip"
	"strconv"
)

func Assemble(src string) ([][16]chip.Signal, error) {
	mem, err := buildMemoryMap(src)
	if err != nil {
		return nil, err
	}
	var program [][16]chip.Signal
	ps := parser{
		lexer: lexer{
			src: src,
		},
	}
	for ps.more() {
		ins, err := ps.next()
		if err != nil {
			return nil, err
		}
		switch v := ins.(type) {
		case load:
			var val int
			if v.value.variant == integer {
				val, err = strconv.Atoi(v.value.literal)
			} else if v.value.variant == identifier {
				val = mem[v.value.literal]
			} else {
				return nil, fmt.Errorf("unexpected load token %+v", v.value)
			}
			program = append(program, chip.WrapUint16(uint16(val)).Copy())
		default:
		}
	}
	return program, nil
}

func buildMemoryMap(src string) (map[string]int, error) {
	mem := make(map[string]int)
	ps := parser{
		lexer: lexer{
			src: src,
		},
	}
	cur := 16
	for ps.more() {
		ins, err := ps.next()
		if err != nil {
			return nil, err
		}
		switch v := ins.(type) {
		case label:
			if _, ok := mem[v.value.literal]; !ok {
				mem[v.value.literal] = cur + 1
				cur++
			}
		case load:
			if _, ok := mem[v.value.literal]; !ok {
				mem[v.value.literal] = cur + 1
				cur++
			}
		default:
		}
	}
	return mem, nil
}
