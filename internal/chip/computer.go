package chip

type Computer struct {
	rom ROM32K
	cpu CPU
	mem Memory
}

// CPU represents the central processing unit of the Computer. It is responsible for executing instructions coming from
// the ROM, which in turn has side effects on the RAM and internal state of the CPU.
type CPU struct {
	a   Register
	d   Register
	alu ALU
	pc  ProgramCounter
}

func (c *CPU) Out(instr [16]Pin, imem [16]Pin, rst Pin) (omem [16]Signal, wmem Signal, addr [15]Signal, pc [15]Signal) {
	a := Mux2Way16(instr[0].Signal(), read16(instr), c.a.Out(NewPin(Inactive), [16]Pin{}))
	a = c.a.Out(instr[0], NewPin16(a))

	c.alu.ZX.Set(instr[4].Signal())
	c.alu.NX.Set(instr[5].Signal())
	c.alu.ZY.Set(instr[6].Signal())
	c.alu.NY.Set(instr[7].Signal())
	c.alu.F.Set(instr[8].Signal())
	c.alu.NO.Set(instr[9].Signal())
	opa := Mux2Way16(instr[3].Signal(), a, read16(imem))   // Choose between the A register and incoming memory data for the first operand
	opb := c.d.Out(NewPin(Inactive), NewPin16(split16(0))) // Just read the D register to act as the second operand
	omem, zr, ng := c.alu.Out(opb, opa)

	c.a.Out(instr[10], NewPin16(omem))
	c.d.Out(instr[11], NewPin16(omem))
	wmem = instr[12].Signal()
	copy(addr[:], a[1:])

	jgt := And(instr[15].Signal(), And(Not(zr), Not(ng)))
	jeq := And(instr[14].Signal(), zr)
	jge := And(Not(ng), And(instr[15].Signal(), instr[14].Signal()))
	jlt := And(instr[13].Signal(), ng)
	jne := And(Not(zr), And(instr[13].Signal(), instr[15].Signal()))
	jle := And(And(instr[13].Signal(), instr[14].Signal()), Or(zr, ng))
	jmp := And(instr[15].Signal(), And(instr[13].Signal(), instr[14].Signal()))
	inc := Not(Or(jgt, Or(jeq, Or(jge, Or(jlt, Or(jne, Or(jle, jmp)))))))
	pc16 := c.pc.Out(NewPin(Not(inc)), NewPin(inc), rst, NewPin16(a))
	copy(pc[:], pc16[1:])
	return
}

type Memory struct {
	ram      RAM16K
	screen   Screen
	keyboard Register
}

func (m *Memory) Out(load Pin, addr [15]Pin, in [16]Pin) [16]Signal {
	rla, rlb, sl, kl := DMux4Way1(
		[2]Signal{addr[0].Signal(), addr[1].Signal()},
		load.Signal(),
	)

	addr14 := [14]Pin{}
	copy(addr14[:], addr[1:])
	addr13 := [13]Pin{}
	copy(addr13[:], addr14[1:])
	return Mux4Way16(
		[2]Signal{addr[0].Signal(), addr[1].Signal()},
		m.ram.Out(NewPin(rla), addr14, in),
		m.ram.Out(NewPin(rlb), addr14, in),
		m.screen.Out(NewPin(sl), addr13, in),
		m.keyboard.Out(NewPin(kl), in),
	)
}

// ROM32K provides a read-only addressable memory (15-bit address space) of 32K size.
type ROM32K struct {
	chips [2]RAM16K
}

func (r *ROM32K) Out(addr [15]Pin) [16]Signal {
	nxt := [14]Pin{}
	copy(nxt[:], addr[1:])
	return Mux2Way16(addr[0].Signal(), r.chips[0].Out(NewPin(Inactive), nxt, [16]Pin{}), r.chips[1].Out(NewPin(Inactive), nxt, [16]Pin{}))
}

// Screen provides volatile storage of 4096 words (16-bit values) that can be addressed with 12 pins.
type Screen struct {
	chips [2]RAM4K
}

// Out either sets and returns or just returns the value for the provided address. When the load pin is active the in
// value is set on the provided address and then returned. When the load pin is inactive the value on the given
// address is just returned.
func (s *Screen) Out(load Pin, addr [13]Pin, in [16]Pin) [16]Signal {
	al, bl := DMux2Way1(
		addr[0].Signal(),
		load.Signal(),
	)
	nxt := [12]Pin{addr[1], addr[2], addr[3], addr[4], addr[5], addr[6], addr[7], addr[8], addr[9], addr[10], addr[11], addr[12]}
	return Mux2Way16(
		addr[0].Signal(),
		s.chips[0].Out(NewPin(al), nxt, in),
		s.chips[1].Out(NewPin(bl), nxt, in),
	)
}
