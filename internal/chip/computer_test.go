package chip

import (
	"fmt"
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
			cpu.a.Out(Active, a.ain)
			cpu.d.Out(Active, a.din)

			omem, wmem, addr, pc := cpu.Out(a.instr, a.imem, a.rst)
			areg := cpu.a.Out(0, [16]Signal{})
			if a.aout != areg {
				t.Errorf("expected aout register to contain %v but found %v", a.aout, areg)
			}
			dreg := cpu.d.Out(0, [16]Signal{})
			if a.dout != dreg {
				t.Errorf("expected dout register to contain %v but found %v", a.dout, dreg)
			}
			if a.pc != pc {
				t.Errorf("expected program counter to contain %v but found %v", a.pc, pc)
			}
			pcreg16 := cpu.pc.Out(0, 0, 0, [16]Signal{})
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

func TestMemory_Out(t *testing.T) {
	ramAddress := func(n uint16) [14]Signal {
		res := [14]Signal{}
		for i := range 14 {
			res[i] = Signal(uint8(n>>(13-i)) & 1)
		}
		p := [14]Signal{}
		for i := range 14 {
			p[i] = res[i]
		}
		return p
	}
	var ramw = []struct {
		address uint16
		value   uint16
		r       uint16
	}{
		{
			address: 0,
			value:   0xF0FF,
			r:       0xF0FF,
		},
		{
			address: 12345,
			value:   0x0FFF,
			r:       0x0FFF,
		},
		{
			address: 16383,
			value:   0xFFFF,
			r:       0xFFFF,
		},
		{
			address: 16384, // The address for this test is outside the RAM range
			value:   0xFFFF,
			r:       0,
		},
	}
	for _, w := range ramw {
		t.Run(fmt.Sprintf("writing %v to %v address", w.value, w.address), func(t *testing.T) {
			mem := Memory{}
			mem.Out(Active, split15(w.address), split16(w.value))
			n := mem.ram.Out(Inactive, ramAddress(w.address), [16]Signal{})
			if n != split16(w.r) {
				t.Errorf("expected RAM n to be %v but got %v", split16(w.r), n)
			}
		})
	}

	screenAddress := func(n uint16) [13]Signal {
		res := [13]Signal{}
		for i := range 13 {
			res[i] = Signal(uint8(n>>(12-i)) & 1)
		}
		p := [13]Signal{}
		for i := range 13 {
			p[i] = res[i]
		}
		return p
	}
	var screenw = []struct {
		addr  uint16
		value uint16
		r     uint16
	}{
		{
			addr:  0b011_0100_0110_0011,
			value: 0xF0FF,
			r:     0,
		},
		{
			addr:  0b100_0100_0110_0011,
			value: 0x0FFF,
			r:     0x0FFF,
		},
		{
			addr:  0b101_1100_0110_0011,
			value: 0xFFFF,
			r:     0xFFFF,
		},
		{
			addr:  0b111_1100_0110_0011,
			value: 0xFFFF,
			r:     0,
		},
	}
	for _, w := range screenw {
		t.Run(fmt.Sprintf("writing %v to %v address", w.value, w.addr), func(t *testing.T) {
			mem := Memory{}
			mem.Out(Active, split15(w.addr), split16(w.value))
			n := mem.screen.Out(Inactive, screenAddress(w.addr), [16]Signal{})
			if n != split16(w.r) {
				t.Errorf("expected screen n to be %v but got %v", split16(w.r), n)
			}
		})
	}

	t.Run("reading and writing keyboard", func(t *testing.T) {
		addr := uint16(24576)
		mem := Memory{}
		mem.Out(Active, split15(addr), split16(1012))
		n := mem.keyboard.Out(Inactive, [16]Signal{})
		if n != split16(1012) {
			t.Errorf("expected keyboard to be %v but got %v", split16(1012), n)
		}
	})
}

func TestROM32K_Out(t *testing.T) {
	equals := func(a [16]Signal, b [16]Signal) bool {
		converted := [16]Signal{}
		for i := range a {
			converted[i] = a[i]
		}
		return converted == b
	}
	address := func(n int) [15]Signal {
		n = n >> 14
		return [15]Signal{
			Signal(n >> 0 & 1),
			Signal(n >> 1 & 1),
		}
	}
	rom := ROM32K{}
	for i := 0; i < 32768; i += 16384 {
		addr := address(i)
		t.Run(fmt.Sprintf("reading address %v", addr), func(t *testing.T) {
			// Reach in and set the ROM on the provided address, we cannot use the load bit for this like we have in the
			// RAM testing since the ROM does not allow writes
			nxt := [14]Signal{}
			copy(nxt[:], addr[1:])
			Mux2Way16(
				addr[0],
				rom.chips[0].Out(Active, nxt, split16(uint16(i))),
				rom.chips[1].Out(Active, nxt, split16(uint16(i))),
			)
			n := rom.Out(addr)
			n = rom.Out(addr)
			if !equals(split16(uint16(i)), n) {
				t.Errorf("expected %v but got %v", i, n)
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
				split16(20),                    // @20
				split16(0b1110_1100_0001_0000), // D = A
				split16(60),                    // @60
				split16(0b1110_0000_1001_0000), // D = A + D
				split16(5),                     // @5
				split16(0b1110_0011_0000_1000), // M = D
			},
			mem: map[uint16][16]Signal{
				5: split16(80),
			},
		},
	}
	for _, a := range assertions {
		t.Run(a.name, func(t *testing.T) {
			c := Computer{}
			c.rom.write(a.program)
			// During normal (non-test) execution the computer will just continue looping forever
			for range len(a.program) {
				c.Tick(Inactive)
			}
			for address, value := range a.mem {
				if out := c.mem.Out(Inactive, split15(address), split16(0)); out != value {
					t.Errorf("expected RAM[%v] to contain %v but got %v", address, value, out)
				}
			}
		})
	}
}
