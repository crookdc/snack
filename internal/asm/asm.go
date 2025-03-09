package asm

import (
	"fmt"
	"github.com/crookdc/nand2tetris/internal/chip"
	"strconv"
)

var (
	computations = map[string]int{
		"0":   0b0101010,
		"1":   0b0111111,
		"-1":  0b0111010,
		"D":   0b0001100,
		"A":   0b0110000,
		"M":   0b1110000,
		"!D":  0b0001101,
		"!A":  0b0110001,
		"!M":  0b1110001,
		"-D":  0b0001111,
		"-A":  0b0110011,
		"-M":  0b1110011,
		"D+1": 0b0011111,
		"A+1": 0b0110111,
		"M+1": 0b1110111,
		"D-1": 0b0001110,
		"A-1": 0b0110010,
		"M-1": 0b1110010,
		"D+A": 0b0000010,
		"D+M": 0b1000010,
		"D-A": 0b0010011,
		"D-M": 0b1010011,
		"A-D": 0b0000111,
		"M-D": 0b1000111,
		"D&A": 0b0000000,
		"D&M": 0b1000000,
		"D|A": 0b0010101,
		"D|M": 0b1010101,
	}
	jumps = map[string]int{
		"JGT": 0b001,
		"JEQ": 0b010,
		"JGE": 0b011,
		"JLT": 0b100,
		"JNE": 0b101,
		"JLE": 0b110,
		"JMP": 0b111,
	}
	destinations = map[string]int{
		"M":   0b001,
		"D":   0b010,
		"DM":  0b011,
		"A":   0b100,
		"AM":  0b101,
		"AD":  0b110,
		"ADM": 0b111,
	}
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
			var bin int
			if v.value.variant == integer {
				bin, err = strconv.Atoi(v.value.literal)
			} else if v.value.variant == identifier {
				bin = mem[v.value.literal]
			} else {
				return nil, fmt.Errorf("unexpected load token %+v", v.value)
			}
			bin = bin & 0b0111_1111_1111_1111
			program = append(program, chip.WrapUint16(uint16(bin)).Copy())
		case compute:
			bin, ok := computations[v.comp]
			if !ok {
				return nil, fmt.Errorf("unexpected computational segment %s", v.comp)
			}
			bin = bin << 6
			if v.dest != nil {
				dest, ok := destinations[v.dest.literal]
				if !ok {
					return nil, fmt.Errorf("invalid destination %+v", v.dest)
				}
				bin = bin | (dest << 3)
			}
			if v.jump != nil {
				jump, ok := jumps[v.jump.literal]
				if !ok {
					return nil, fmt.Errorf("invalid jump %+v", v.jump)
				}
				bin = bin | jump
			}
			bin = bin | 0b1110_0000_0000_0000
			program = append(program, chip.WrapUint16(uint16(bin)).Copy())
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
	for line := 0; ps.more(); line++ {
		ins, err := ps.next()
		if err != nil {
			return nil, err
		}
		switch v := ins.(type) {
		case label:
			if _, ok := mem[v.value.literal]; !ok {
				mem[v.value.literal] = line
				cur++
			}
			line--
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
