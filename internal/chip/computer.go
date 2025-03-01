package chip

// NewComputer creates a new Computer chip with the provided program preloaded into its ROM.
func NewComputer(program [][16]Signal) Computer {
	c := Computer{}
	c.rom.write(program)
	return c
}

type Computer struct {
	rom ROM32K
	cpu CPU
	mem Memory
}

func (c *Computer) Tick(rst Signal) {
	addr := c.cpu.pc.Out(Inactive, Inactive, rst, [16]Signal{})
	iaddr := [15]Signal{}
	copy(iaddr[:], addr[1:])
	instr := c.rom.Out(iaddr)
	aout := c.cpu.a.Out(Inactive, [16]Signal{})
	maddr := [15]Signal{}
	copy(maddr[:], aout[1:])
	imem := c.mem.Out(Inactive, maddr, [16]Signal{})
	omem, wmem, maddr, _ := c.cpu.Out(instr, imem, rst)
	c.mem.Out(wmem, maddr, omem)
}

// CPU represents the central processing unit of the Computer. It is responsible for executing instructions coming from
// the ROM, which in turn has side effects on the RAM and internal state of the CPU.
type CPU struct {
	a   Register
	d   Register
	alu ALU
	pc  ProgramCounter
}

func (c *CPU) Out(instr [16]Signal, imem [16]Signal, rst Signal) (omem [16]Signal, wmem Signal, addr [15]Signal, pc [15]Signal) {
	a := Mux2Way16(instr[0], instr, c.a.Out(Inactive, [16]Signal{}))
	a = c.a.Out(Not(instr[0]), a)
	instr = And16(expand16(instr[0]), instr)

	c.alu.ZX = instr[4]
	c.alu.NX = instr[5]
	c.alu.ZY = instr[6]
	c.alu.NY = instr[7]
	c.alu.F = instr[8]
	c.alu.NO = instr[9]
	opa := Mux2Way16(instr[3], a, imem)  // Choose between the A register and incoming memory data for the first operand
	opb := c.d.Out(Inactive, split16(0)) // Just read the D register to act as the second operand
	omem, zr, ng := c.alu.Out(opb, opa)

	c.a.Out(instr[10], omem)
	c.d.Out(instr[11], omem)
	wmem = instr[12]
	copy(addr[:], a[1:])

	jgt := And(instr[15], And(Not(zr), Not(ng)))
	jeq := And(instr[14], zr)
	jge := And(Not(ng), And(instr[15], instr[14]))
	jlt := And(instr[13], ng)
	jne := And(Not(zr), And(instr[13], instr[15]))
	jle := And(And(instr[13], instr[14]), Or(zr, ng))
	jmp := And(instr[15], And(instr[13], instr[14]))
	inc := Not(Or(jgt, Or(jeq, Or(jge, Or(jlt, Or(jne, Or(jle, jmp)))))))
	pc16 := c.pc.Out(Not(inc), inc, rst, a)
	copy(pc[:], pc16[1:])
	return
}

type Memory struct {
	ram      RAM16K
	screen   Screen
	keyboard Register
}

func (m *Memory) Out(load Signal, addr [15]Signal, in [16]Signal) [16]Signal {
	rla, rlb, sl, kl := DMux4Way1(
		[2]Signal{addr[0], addr[1]},
		load,
	)

	addr14 := [14]Signal{}
	copy(addr14[:], addr[1:])
	addr13 := [13]Signal{}
	copy(addr13[:], addr14[1:])
	return Mux4Way16(
		[2]Signal{addr[0], addr[1]},
		m.ram.Out(rla, addr14, in),
		m.ram.Out(rlb, addr14, in),
		m.screen.Out(sl, addr13, in),
		m.keyboard.Out(kl, in),
	)
}

// ROM32K provides a read-only addressable memory (15-bit address space) of 32K size.
type ROM32K struct {
	chips [2]RAM16K
}

func (r *ROM32K) Out(addr [15]Signal) [16]Signal {
	nxt := [14]Signal{}
	copy(nxt[:], addr[1:])
	return Mux2Way16(addr[0], r.chips[0].Out(Inactive, nxt, [16]Signal{}), r.chips[1].Out(Inactive, nxt, [16]Signal{}))
}

func (r *ROM32K) write(program [][16]Signal) {
	for i, instr := range program {
		addr := split16(uint16(i))
		al, bl := DMux2Way1(addr[0], Active)
		nxt := [14]Signal{}
		copy(nxt[:], addr[2:])
		Mux2Way16(
			addr[0],
			r.chips[0].Out(al, nxt, instr),
			r.chips[1].Out(bl, nxt, instr),
		)
	}
}

// Screen provides volatile storage of 4096 words (16-bit values) that can be addressed with 12 pins.
type Screen struct {
	chips [2]RAM4K
}

// Out either sets and returns or just returns the value for the provided address. When the load pin is active the in
// value is set on the provided address and then returned. When the load pin is inactive the value on the given
// address is just returned.
func (s *Screen) Out(load Signal, addr [13]Signal, in [16]Signal) [16]Signal {
	al, bl := DMux2Way1(
		addr[0],
		load,
	)
	nxt := [12]Signal{addr[1], addr[2], addr[3], addr[4], addr[5], addr[6], addr[7], addr[8], addr[9], addr[10], addr[11], addr[12]}
	return Mux2Way16(
		addr[0],
		s.chips[0].Out(al, nxt, in),
		s.chips[1].Out(bl, nxt, in),
	)
}
