package main

import (
	"bufio"
	"fmt"
	"github.com/crookdc/nand2tetris/internal/chip"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		panic("missing path to program")
	}
	program, err := loadProgram(os.Args[1])
	if err != nil {
		panic(err)
	}
	computer := chip.NewComputer(program)
	clock := time.Now()
	for {
		delta := time.Now().Sub(clock)
		for delta.Milliseconds() < 16 {
			delta = time.Now().Sub(clock)
		}
		computer.Tick(chip.Inactive)
		clock = time.Now()
	}
}

func loadProgram(file string) ([][16]chip.Signal, error) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var program [][16]chip.Signal
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		if len(line) != 16 {
			return nil, fmt.Errorf("invalid line length '%s'", line)
		}
		instruction, err := parseInstruction(line)
		if err != nil {
			return nil, err
		}
		program = append(program, instruction)
	}
	return program, nil
}

func parseInstruction(line string) ([16]chip.Signal, error) {
	var instruction [16]chip.Signal
	for i := range 16 {
		bit, err := strconv.Atoi(string(line[i]))
		if err != nil {
			return [16]chip.Signal{}, err
		}
		instruction[i] = chip.Signal(bit)
	}
	return instruction, nil
}
