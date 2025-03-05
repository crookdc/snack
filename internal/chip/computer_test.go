package chip

import (
	"testing"
)

func TestCPU_Out(t *testing.T) {
	var assertions = []struct {
		name  string
		instr [16]Signal
		imem  [16]Signal
		rst   Signal
		ain   [16]Signal
		din   [16]Signal

		aout [16]Signal
		dout [16]Signal
		pc   [15]Signal
		omem [16]Signal
		wmem Signal
		addr [15]Signal
	}{
		{
			name:  "0",
			instr: split16(0b1110101010000000),
			omem:  split16(0),
			pc:    split15(1),
		},
		{
			name:  "1",
			instr: split16(0b1110111111000000),
			omem:  split16(1),
			pc:    split15(1),
		},
		{
			name:  "-1",
			instr: split16(0b1110111010000000),
			omem:  split16(0xFFFF),
			pc:    split15(1),
		},
		{
			name:  "D",
			instr: split16(0b1110001100000000),
			din:   split16(0xF00F),
			dout:  split16(0xF00F),
			omem:  split16(0xF00F),
			pc:    split15(1),
		},
		{
			name:  "A",
			instr: split16(0b1110110000000000),
			ain:   split16(0xF00F),
			aout:  split16(0xF00F),
			omem:  split16(0xF00F),
			addr:  split15(0xF00F),
			pc:    split15(1),
		},
		{
			name:  "M",
			instr: split16(0b1111110000000000),
			imem:  split16(0xF00F),
			omem:  split16(0xF00F),
			pc:    split15(1),
		},
		{
			name:  "!D",
			instr: split16(0b1110001101000000),
			din:   split16(0xF00F),
			dout:  split16(0xF00F),
			omem:  split16(0x0FF0),
			pc:    split15(1),
		},
		{
			name:  "!A",
			instr: split16(0b1110110001000000),
			ain:   split16(0xF00F),
			aout:  split16(0xF00F),
			omem:  split16(0x0FF0),
			addr:  split15(0xF00F),
			pc:    split15(1),
		},
		{
			name:  "!M",
			instr: split16(0b1111110001000000),
			imem:  split16(0xF00F),
			omem:  split16(0x0FF0),
			pc:    split15(1),
		},
		{
			name:  "-D",
			instr: split16(0b1110001111000000),
			din:   split16(2),
			dout:  split16(2),
			omem:  split16(0xFFFE),
			pc:    split15(1),
		},
		{
			name:  "-A",
			instr: split16(0b1110110011000000),
			ain:   split16(2),
			aout:  split16(2),
			addr:  split15(2),
			omem:  split16(0xFFFE),
			pc:    split15(1),
		},
		{
			name:  "-M",
			instr: split16(0b1111110011000000),
			imem:  split16(2),
			omem:  split16(0xFFFE),
			pc:    split15(1),
		},
		{
			name:  "D+1",
			instr: split16(0b1110011111000000),
			din:   split16(5),
			dout:  split16(5),
			omem:  split16(6),
			pc:    split15(1),
		},
		{
			name:  "A+1",
			instr: split16(0b1110110111000000),
			ain:   split16(5),
			aout:  split16(5),
			omem:  split16(6),
			addr:  split15(5),
			pc:    split15(1),
		},
		{
			name:  "M+1",
			instr: split16(0b1111110111000000),
			imem:  split16(5),
			omem:  split16(6),
			pc:    split15(1),
		},
		{
			name:  "D-1",
			instr: split16(0b1110001110000000),
			din:   split16(5),
			dout:  split16(5),
			omem:  split16(4),
			pc:    split15(1),
		},
		{
			name:  "A-1",
			instr: split16(0b1110110010000000),
			ain:   split16(5),
			aout:  split16(5),
			omem:  split16(4),
			addr:  split15(5),
			pc:    split15(1),
		},
		{
			name:  "M-1",
			instr: split16(0b1111110010000000),
			imem:  split16(5),
			omem:  split16(4),
			pc:    split15(1),
		},
		{
			name:  "D+A",
			instr: split16(0b1110000010000000),
			ain:   split16(5),
			aout:  split16(5),
			din:   split16(4),
			dout:  split16(4),
			omem:  split16(9),
			addr:  split15(5),
			pc:    split15(1),
		},
		{
			name:  "D+M",
			instr: split16(0b1111000010000000),
			imem:  split16(10),
			din:   split16(4),
			dout:  split16(4),
			omem:  split16(14),
			pc:    split15(1),
		},
		{
			name:  "D-A",
			instr: split16(0b1110010011000000),
			ain:   split16(4),
			aout:  split16(4),
			din:   split16(6),
			dout:  split16(6),
			omem:  split16(2),
			addr:  split15(4),
			pc:    split15(1),
		},
		{
			name:  "D-M",
			instr: split16(0b1111010011000000),
			imem:  split16(6),
			din:   split16(10),
			dout:  split16(10),
			omem:  split16(4),
			pc:    split15(1),
		},
		{
			name:  "A-D",
			instr: split16(0b1110000111000000),
			ain:   split16(6),
			aout:  split16(6),
			din:   split16(4),
			dout:  split16(4),
			omem:  split16(2),
			addr:  split15(6),
			pc:    split15(1),
		},
		{
			name:  "M-D",
			instr: split16(0b1111000111000000),
			imem:  split16(19),
			din:   split16(7),
			dout:  split16(7),
			omem:  split16(12),
			pc:    split15(1),
		},
		{
			name:  "D&A",
			instr: split16(0b1110000000000000),
			ain:   split16(0xF00F),
			aout:  split16(0xF00F),
			din:   split16(0xFF00),
			dout:  split16(0xFF00),
			omem:  split16(0xF000),
			addr:  split15(0xF00F),
			pc:    split15(1),
		},
		{
			name:  "D&M",
			instr: split16(0b1111000000000000),
			imem:  split16(0x0F0F),
			din:   split16(0xFF00),
			dout:  split16(0xFF00),
			omem:  split16(0x0F00),
			pc:    split15(1),
		},
		{
			name:  "D|A",
			instr: split16(0b1110010101000000),
			ain:   split16(0xF00F),
			aout:  split16(0xF00F),
			din:   split16(0xF0F0),
			dout:  split16(0xF0F0),
			omem:  split16(0xF0FF),
			addr:  split15(0xF00F),
			pc:    split15(1),
		},
		{
			name:  "D|M",
			instr: split16(0b1111010101000000),
			imem:  split16(0xF00F),
			din:   split16(0xF0F0),
			dout:  split16(0xF0F0),
			omem:  split16(0xF0FF),
			pc:    split15(1),
		},
		{
			name:  "A = !A",
			instr: split16(0b1110110001100000),
			aout:  split16(0xFFFF),
			omem:  split16(0xFFFF),
			pc:    split15(1),
		},
		{
			name:  "D = !A",
			instr: split16(0b1110110001010000),
			aout:  split16(0),
			dout:  split16(0xFFFF),
			omem:  split16(0xFFFF),
			pc:    split15(1),
		},
		{
			name:  "D, A = !A",
			instr: split16(0b1110110001110000),
			ain:   split16(0xF00F),
			aout:  split16(0x0FF0),
			dout:  split16(0x0FF0),
			omem:  split16(0x0FF0),
			addr:  split15(0xF00F),
			pc:    split15(1),
		},
		{
			name:  "D == 0, JGT",
			instr: split16(0b1110001100000001),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			pc:    split15(1),
		},
		{
			name:  "D < 0, JGT",
			instr: split16(0b1110001100000001),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   split16(0xF000),
			dout:  split16(0xF000),
			omem:  split16(0xF000),
			pc:    split15(1),
		},
		{
			name:  "D > 0, JGT",
			instr: split16(0b1110001100000001),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   split16(512),
			dout:  split16(512),
			omem:  split16(512),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D == 0, JEQ",
			instr: split16(0b1110001100000010),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D < 0, JEQ",
			instr: split16(0b1110001100000010),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   split16(0xF000),
			dout:  split16(0xF000),
			omem:  split16(0xF000),
			pc:    split15(1),
		},
		{
			name:  "D > 0, JEQ",
			instr: split16(0b1110001100000010),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   split16(512),
			dout:  split16(512),
			omem:  split16(512),
			pc:    split15(1),
		},
		{
			name:  "D == 0, JGE",
			instr: split16(0b1110001100000011),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D < 0, JGE",
			instr: split16(0b1110001100000011),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   split16(0xF000),
			dout:  split16(0xF000),
			omem:  split16(0xF000),
			pc:    split15(1),
		},
		{
			name:  "D > 0, JGE",
			instr: split16(0b1110001100000011),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   split16(512),
			dout:  split16(512),
			omem:  split16(512),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D == 0, JLT",
			instr: split16(0b1110001100000100),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			pc:    split15(1),
		},
		{
			name:  "D < 0, JLT",
			instr: split16(0b1110001100000100),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   split16(0xF000),
			dout:  split16(0xF000),
			omem:  split16(0xF000),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D > 0, JLT",
			instr: split16(0b1110001100000100),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   split16(512),
			dout:  split16(512),
			omem:  split16(512),
			pc:    split15(1),
		},
		{
			name:  "D == 0, JNE",
			instr: split16(0b1110001100000101),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			pc:    split15(1),
		},
		{
			name:  "D < 0, JNE",
			instr: split16(0b1110001100000101),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   split16(0xF000),
			dout:  split16(0xF000),
			omem:  split16(0xF000),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D > 0, JNE",
			instr: split16(0b1110001100000101),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   split16(512),
			dout:  split16(512),
			omem:  split16(512),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D == 0, JLE",
			instr: split16(0b1110001100000110),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D < 0, JLE",
			instr: split16(0b1110001100000110),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   split16(0xF000),
			dout:  split16(0xF000),
			omem:  split16(0xF000),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D > 0, JLE",
			instr: split16(0b1110001100000110),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   split16(512),
			dout:  split16(512),
			omem:  split16(512),
			pc:    split15(1),
		},
		{
			name:  "D == 0, JMP",
			instr: split16(0b1110001100000111),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D < 0, JMP",
			instr: split16(0b1110001100000111),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   split16(0xF000),
			dout:  split16(0xF000),
			omem:  split16(0xF000),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D > 0, JMP",
			instr: split16(0b1110001100000111),
			ain:   split16(0xFFFF),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   split16(512),
			dout:  split16(512),
			omem:  split16(512),
			pc:    split15(0xFFFF),
		},
	}
	for _, a := range assertions {
		t.Run(a.name, func(t *testing.T) {
			cpu := CPU{}
			cpu.a.Out(Active, Wrap(&a.ain))
			cpu.d.Out(Active, Wrap(&a.din))

			omem, wmem, addr := cpu.Out(Wrap(&a.instr), Wrap(&a.imem), a.rst)
			areg := cpu.a.Out(0, NullWord)
			if a.aout != areg.Copy() {
				t.Errorf("expected aout register to contain %v but found %v", a.aout, areg)
			}
			dreg := cpu.d.Out(0, NullWord)
			if a.dout != dreg.Copy() {
				t.Errorf("expected dout register to contain %v but found %v", a.dout, dreg)
			}
			pcreg := cpu.pc.Out(0, 0, 0, NullWord).Address()
			if a.pc != pcreg {
				t.Errorf("expected pcreg to contain %v but got %v", a.pc, pcreg)
			}
			if a.omem != omem.Copy() {
				t.Errorf("expected omem to contain %v but found %v", a.omem, omem)
			}
			if a.wmem != wmem {
				t.Errorf("expected wmem to contain %v but found %v", a.wmem, wmem)
			}
			if a.addr != addr {
				t.Errorf("expected addr to contain %v but found %v", a.addr, addr)
			}
		})
	}
}

func TestComputer_Tick(t *testing.T) {
	var assertions = []struct {
		name    string
		program [][16]Signal
		mem     map[uint16][16]Signal
	}{
		{
			name: "add two integers and store in RAM",
			program: [][16]Signal{
				split16(0b0000_0000_0001_0100), // @20
				split16(0b1110_1100_0001_0000), // D = A
				split16(0b0000_0000_0011_1100), // @60
				split16(0b1110_0000_1001_0000), // D = A + D
				split16(0b0000_0000_0000_0101), // @5
				split16(0b1110_0011_0000_1000), // M = D
			},
			mem: map[uint16][16]Signal{
				5: split16(80),
			},
		},
		{
			name: "multiply an integer with itself",
			program: [][16]Signal{
				split16(0b0000000000000001), // @1
				split16(0b1110101010001000), // M=0
				split16(0b0000000000000100), // @4
				split16(0b1110110000010000), // D=A
				split16(0b0000000000000000), // @0
				split16(0b1110001100001000), // M=D
				split16(0b0000000000000001), // @1
				split16(0b1111110000010000), // D=M
				split16(0b0000000000000100), // @4
				split16(0b1110000010010000), // D=D+A
				split16(0b0000000000000001), // @1
				split16(0b1110001100001000), // M=D
				split16(0b0000000000000000), // @0
				split16(0b1111110000010000), // D=M
				split16(0b1110001110010000), // D=D-1
				split16(0b1110001100001000), // M=D
				split16(0b0000000000000100), // @4
				split16(0b1110001100000001), // D;JGT
			},
			mem: map[uint16][16]Signal{
				1: split16(16),
			},
		},
	}
	for _, a := range assertions {
		t.Run(a.name, func(t *testing.T) {
			c := Computer{
				mem: &RAM{},
			}
			c.rom = ROM(a.program)
			pc := c.cpu.pc.Out(Inactive, Inactive, Inactive, NullWord)
			for pc.Uint16() < uint16(len(a.program)) {
				c.Tick(Inactive)
				pc = c.cpu.pc.Out(Inactive, Inactive, Inactive, NullWord)
			}
			for address, value := range a.mem {
				if out := c.mem.Out(Inactive, split15(address), NullWord); out.Copy() != value {
					t.Errorf("expected RAM[%v] to contain %v but got %v", address, value, out)
				}
			}
		})
	}
}
