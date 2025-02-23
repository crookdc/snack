package chip

import (
	"testing"
)

func TestCPU_Out(t *testing.T) {
	var assertions = []struct {
		name  string
		instr [16]Pin
		imem  [16]Pin
		rst   Pin
		ain   [16]Pin
		din   [16]Pin

		aout [16]Signal
		dout [16]Signal
		pc   [15]Signal
		omem [16]Signal
		wmem Signal
		addr [15]Signal
	}{
		{
			name:  "0",
			instr: NewPin16(split16(0b1110101010000000)),
			omem:  split16(0),
			pc:    split15(1),
		},
		{
			name:  "1",
			instr: NewPin16(split16(0b1110111111000000)),
			omem:  split16(1),
			pc:    split15(1),
		},
		{
			name:  "-1",
			instr: NewPin16(split16(0b1110111010000000)),
			omem:  split16(0xFFFF),
			pc:    split15(1),
		},
		{
			name:  "D",
			instr: NewPin16(split16(0b1110001100000000)),
			din:   NewPin16(split16(0xF00F)),
			dout:  split16(0xF00F),
			omem:  split16(0xF00F),
			pc:    split15(1),
		},
		{
			name:  "A",
			instr: NewPin16(split16(0b1110110000000000)),
			ain:   NewPin16(split16(0xF00F)),
			aout:  split16(0xF00F),
			omem:  split16(0xF00F),
			addr:  split15(0xF00F),
			pc:    split15(1),
		},
		{
			name:  "M",
			instr: NewPin16(split16(0b1111110000000000)),
			imem:  NewPin16(split16(0xF00F)),
			omem:  split16(0xF00F),
			pc:    split15(1),
		},
		{
			name:  "!D",
			instr: NewPin16(split16(0b1110001101000000)),
			din:   NewPin16(split16(0xF00F)),
			dout:  split16(0xF00F),
			omem:  split16(0x0FF0),
			pc:    split15(1),
		},
		{
			name:  "!A",
			instr: NewPin16(split16(0b1110110001000000)),
			ain:   NewPin16(split16(0xF00F)),
			aout:  split16(0xF00F),
			omem:  split16(0x0FF0),
			addr:  split15(0xF00F),
			pc:    split15(1),
		},
		{
			name:  "!M",
			instr: NewPin16(split16(0b1111110001000000)),
			imem:  NewPin16(split16(0xF00F)),
			omem:  split16(0x0FF0),
			pc:    split15(1),
		},
		{
			name:  "-D",
			instr: NewPin16(split16(0b1110001111000000)),
			din:   NewPin16(split16(2)),
			dout:  split16(2),
			omem:  split16(0xFFFE),
			pc:    split15(1),
		},
		{
			name:  "-A",
			instr: NewPin16(split16(0b1110110011000000)),
			ain:   NewPin16(split16(2)),
			aout:  split16(2),
			addr:  split15(2),
			omem:  split16(0xFFFE),
			pc:    split15(1),
		},
		{
			name:  "-M",
			instr: NewPin16(split16(0b1111110011000000)),
			imem:  NewPin16(split16(2)),
			omem:  split16(0xFFFE),
			pc:    split15(1),
		},
		{
			name:  "D+1",
			instr: NewPin16(split16(0b1110011111000000)),
			din:   NewPin16(split16(5)),
			dout:  split16(5),
			omem:  split16(6),
			pc:    split15(1),
		},
		{
			name:  "A+1",
			instr: NewPin16(split16(0b1110110111000000)),
			ain:   NewPin16(split16(5)),
			aout:  split16(5),
			omem:  split16(6),
			addr:  split15(5),
			pc:    split15(1),
		},
		{
			name:  "M+1",
			instr: NewPin16(split16(0b1111110111000000)),
			imem:  NewPin16(split16(5)),
			omem:  split16(6),
			pc:    split15(1),
		},
		{
			name:  "D-1",
			instr: NewPin16(split16(0b1110001110000000)),
			din:   NewPin16(split16(5)),
			dout:  split16(5),
			omem:  split16(4),
			pc:    split15(1),
		},
		{
			name:  "A-1",
			instr: NewPin16(split16(0b1110110010000000)),
			ain:   NewPin16(split16(5)),
			aout:  split16(5),
			omem:  split16(4),
			addr:  split15(5),
			pc:    split15(1),
		},
		{
			name:  "M-1",
			instr: NewPin16(split16(0b1111110010000000)),
			imem:  NewPin16(split16(5)),
			omem:  split16(4),
			pc:    split15(1),
		},
		{
			name:  "D+A",
			instr: NewPin16(split16(0b1110000010000000)),
			ain:   NewPin16(split16(5)),
			aout:  split16(5),
			din:   NewPin16(split16(4)),
			dout:  split16(4),
			omem:  split16(9),
			addr:  split15(5),
			pc:    split15(1),
		},
		{
			name:  "D+M",
			instr: NewPin16(split16(0b1111000010000000)),
			imem:  NewPin16(split16(10)),
			din:   NewPin16(split16(4)),
			dout:  split16(4),
			omem:  split16(14),
			pc:    split15(1),
		},
		{
			name:  "D-A",
			instr: NewPin16(split16(0b1110010011000000)),
			ain:   NewPin16(split16(4)),
			aout:  split16(4),
			din:   NewPin16(split16(6)),
			dout:  split16(6),
			omem:  split16(2),
			addr:  split15(4),
			pc:    split15(1),
		},
		{
			name:  "D-M",
			instr: NewPin16(split16(0b1111010011000000)),
			imem:  NewPin16(split16(6)),
			din:   NewPin16(split16(10)),
			dout:  split16(10),
			omem:  split16(4),
			pc:    split15(1),
		},
		{
			name:  "A-D",
			instr: NewPin16(split16(0b1110000111000000)),
			ain:   NewPin16(split16(6)),
			aout:  split16(6),
			din:   NewPin16(split16(4)),
			dout:  split16(4),
			omem:  split16(2),
			addr:  split15(6),
			pc:    split15(1),
		},
		{
			name:  "M-D",
			instr: NewPin16(split16(0b1111000111000000)),
			imem:  NewPin16(split16(19)),
			din:   NewPin16(split16(7)),
			dout:  split16(7),
			omem:  split16(12),
			pc:    split15(1),
		},
		{
			name:  "D&A",
			instr: NewPin16(split16(0b1110000000000000)),
			ain:   NewPin16(split16(0xF00F)),
			aout:  split16(0xF00F),
			din:   NewPin16(split16(0xFF00)),
			dout:  split16(0xFF00),
			omem:  split16(0xF000),
			addr:  split15(0xF00F),
			pc:    split15(1),
		},
		{
			name:  "D&M",
			instr: NewPin16(split16(0b1111000000000000)),
			imem:  NewPin16(split16(0x0F0F)),
			din:   NewPin16(split16(0xFF00)),
			dout:  split16(0xFF00),
			omem:  split16(0x0F00),
			pc:    split15(1),
		},
		{
			name:  "D|A",
			instr: NewPin16(split16(0b1110010101000000)),
			ain:   NewPin16(split16(0xF00F)),
			aout:  split16(0xF00F),
			din:   NewPin16(split16(0xF0F0)),
			dout:  split16(0xF0F0),
			omem:  split16(0xF0FF),
			addr:  split15(0xF00F),
			pc:    split15(1),
		},
		{
			name:  "D|M",
			instr: NewPin16(split16(0b1111010101000000)),
			imem:  NewPin16(split16(0xF00F)),
			din:   NewPin16(split16(0xF0F0)),
			dout:  split16(0xF0F0),
			omem:  split16(0xF0FF),
			pc:    split15(1),
		},
		{
			name:  "A = !A",
			instr: NewPin16(split16(0b1110110001100000)),
			aout:  split16(0xFFFF),
			omem:  split16(0xFFFF),
			pc:    split15(1),
		},
		{
			name:  "D = !A",
			instr: NewPin16(split16(0b1110110001010000)),
			aout:  split16(0),
			dout:  split16(0xFFFF),
			omem:  split16(0xFFFF),
			pc:    split15(1),
		},
		{
			name:  "D, A = !A",
			instr: NewPin16(split16(0b1110110001110000)),
			ain:   NewPin16(split16(0xF00F)),
			aout:  split16(0x0FF0),
			dout:  split16(0x0FF0),
			omem:  split16(0x0FF0),
			addr:  split15(0xF00F),
			pc:    split15(1),
		},
		{
			name:  "D == 0, JGT",
			instr: NewPin16(split16(0b1110001100000001)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			pc:    split15(1),
		},
		{
			name:  "D < 0, JGT",
			instr: NewPin16(split16(0b1110001100000001)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   NewPin16(split16(0xF000)),
			dout:  split16(0xF000),
			omem:  split16(0xF000),
			pc:    split15(1),
		},
		{
			name:  "D > 0, JGT",
			instr: NewPin16(split16(0b1110001100000001)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   NewPin16(split16(512)),
			dout:  split16(512),
			omem:  split16(512),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D == 0, JEQ",
			instr: NewPin16(split16(0b1110001100000010)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D < 0, JEQ",
			instr: NewPin16(split16(0b1110001100000010)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   NewPin16(split16(0xF000)),
			dout:  split16(0xF000),
			omem:  split16(0xF000),
			pc:    split15(1),
		},
		{
			name:  "D > 0, JEQ",
			instr: NewPin16(split16(0b1110001100000010)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   NewPin16(split16(512)),
			dout:  split16(512),
			omem:  split16(512),
			pc:    split15(1),
		},
		{
			name:  "D == 0, JGE",
			instr: NewPin16(split16(0b1110001100000011)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D < 0, JGE",
			instr: NewPin16(split16(0b1110001100000011)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   NewPin16(split16(0xF000)),
			dout:  split16(0xF000),
			omem:  split16(0xF000),
			pc:    split15(1),
		},
		{
			name:  "D > 0, JGE",
			instr: NewPin16(split16(0b1110001100000011)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   NewPin16(split16(512)),
			dout:  split16(512),
			omem:  split16(512),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D == 0, JLT",
			instr: NewPin16(split16(0b1110001100000100)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			pc:    split15(1),
		},
		{
			name:  "D < 0, JLT",
			instr: NewPin16(split16(0b1110001100000100)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   NewPin16(split16(0xF000)),
			dout:  split16(0xF000),
			omem:  split16(0xF000),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D > 0, JLT",
			instr: NewPin16(split16(0b1110001100000100)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   NewPin16(split16(512)),
			dout:  split16(512),
			omem:  split16(512),
			pc:    split15(1),
		},
		{
			name:  "D == 0, JNE",
			instr: NewPin16(split16(0b1110001100000101)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			pc:    split15(1),
		},
		{
			name:  "D < 0, JNE",
			instr: NewPin16(split16(0b1110001100000101)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   NewPin16(split16(0xF000)),
			dout:  split16(0xF000),
			omem:  split16(0xF000),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D > 0, JNE",
			instr: NewPin16(split16(0b1110001100000101)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   NewPin16(split16(512)),
			dout:  split16(512),
			omem:  split16(512),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D == 0, JLE",
			instr: NewPin16(split16(0b1110001100000110)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D < 0, JLE",
			instr: NewPin16(split16(0b1110001100000110)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   NewPin16(split16(0xF000)),
			dout:  split16(0xF000),
			omem:  split16(0xF000),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D > 0, JLE",
			instr: NewPin16(split16(0b1110001100000110)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   NewPin16(split16(512)),
			dout:  split16(512),
			omem:  split16(512),
			pc:    split15(1),
		},
		{
			name:  "D == 0, JMP",
			instr: NewPin16(split16(0b1110001100000111)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D < 0, JMP",
			instr: NewPin16(split16(0b1110001100000111)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   NewPin16(split16(0xF000)),
			dout:  split16(0xF000),
			omem:  split16(0xF000),
			pc:    split15(0xFFFF),
		},
		{
			name:  "D > 0, JMP",
			instr: NewPin16(split16(0b1110001100000111)),
			ain:   NewPin16(split16(0xFFFF)),
			aout:  split16(0xFFFF),
			addr:  split15(0xFFFF),
			din:   NewPin16(split16(512)),
			dout:  split16(512),
			omem:  split16(512),
			pc:    split15(0xFFFF),
		},
	}
	for _, a := range assertions {
		t.Run(a.name, func(t *testing.T) {
			cpu := CPU{}
			cpu.a.Out(NewPin(1), a.ain)
			cpu.d.Out(NewPin(1), a.din)

			omem, wmem, addr, pc := cpu.Out(a.instr, a.imem, a.rst)
			areg := cpu.a.Out(NewPin(0), [16]Pin{})
			if a.aout != areg {
				t.Errorf("expected aout register to contain %v but found %v", a.aout, areg)
			}
			dreg := cpu.d.Out(NewPin(0), [16]Pin{})
			if a.dout != dreg {
				t.Errorf("expected dout register to contain %v but found %v", a.dout, dreg)
			}
			if a.pc != pc {
				t.Errorf("expected program counter to contain %v but found %v", a.pc, pc)
			}
			pcreg16 := cpu.pc.Out(NewPin(0), NewPin(0), NewPin(0), [16]Pin{})
			pcreg := [15]Signal{}
			copy(pcreg[:], pcreg16[1:])
			if a.pc != pcreg {
				t.Errorf("expected pcreg to contain %v but got %v", a.pc, pcreg)
			}
			if a.omem != omem {
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
