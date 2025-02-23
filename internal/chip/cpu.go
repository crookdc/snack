package chip

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
