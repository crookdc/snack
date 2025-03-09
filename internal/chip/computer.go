package chip

// NewComputer creates a new Computer chip with the provided program preloaded into its ROM.
func NewComputer(rom Memory) Computer {
	return Computer{
		rom: rom,
		mem: &RAM{},
	}
}

type Memory interface {
	Out(load Signal, addr [15]Signal, in ReadonlyWord) *Word
}

type ROM [][16]Signal

func (r ROM) Out(_ Signal, addr [15]Signal, _ ReadonlyWord) *Word {
	return Wrap(&r[Join15(addr)])
}

type RAM struct {
	mem [32768][16]Signal
}

func (b *RAM) Out(load Signal, addr [15]Signal, in ReadonlyWord) *Word {
	idx := Join15(addr)
	if load == Inactive {
		return Wrap(&b.mem[idx])
	}
	b.mem[idx] = in.Copy()
	return Wrap(&b.mem[idx])
}

type Computer struct {
	rom Memory
	cpu CPU
	mem Memory
}

func (c *Computer) RAM() Memory {
	return c.mem
}

func (c *Computer) Tick(rst Signal) {
	addr := c.cpu.pc.Out(Inactive, Inactive, rst, NullWord)
	instr := c.rom.Out(Inactive, addr.Address(), NullWord)
	areg := c.cpu.a.Out(Inactive, NullWord)
	imem := c.mem.Out(Inactive, areg.Address(), NullWord)
	omem, wmem, maddr := c.cpu.Out(instr, imem, rst)
	c.mem.Out(wmem, maddr, omem)
	return
}

// CPU represents the central processing unit of the Computer. It is responsible for executing instructions coming from
// the ROM, which in turn has side effects on the RAM and internal state of the CPU.
type CPU struct {
	a   Register
	d   Register
	alu ALU
	pc  PC
}

func (c *CPU) Out(instr ReadonlyWord, imem ReadonlyWord, rst Signal) (omem *Word, wmem Signal, addr [15]Signal) {
	a := Mux2Way16(instr.Get(0), instr, c.a.Out(Inactive, NullWord))
	a = c.a.Out(Not(instr.Get(0)), a)
	instr = And16To1(instr.Get(0), instr)

	c.alu.ZX = instr.Get(4)
	c.alu.NX = instr.Get(5)
	c.alu.ZY = instr.Get(6)
	c.alu.NY = instr.Get(7)
	c.alu.F = instr.Get(8)
	c.alu.NO = instr.Get(9)
	opa := Mux2Way16(instr.Get(3), a, imem)
	opb := c.d.Out(Inactive, NullWord)
	omem, zr, ng := c.alu.Out(opb, opa)

	c.a.Out(instr.Get(10), omem)
	c.d.Out(instr.Get(11), omem)
	wmem = instr.Get(12)
	addr = a.Address()

	jgt := And(instr.Get(15), And(Not(zr), Not(ng)))
	jeq := And(instr.Get(14), zr)
	jge := And(Not(ng), And(instr.Get(15), instr.Get(14)))
	jlt := And(instr.Get(13), ng)
	jne := And(Not(zr), And(instr.Get(13), instr.Get(15)))
	jle := And(And(instr.Get(13), instr.Get(14)), Or(zr, ng))
	jmp := And(instr.Get(15), And(instr.Get(13), instr.Get(14)))
	inc := Not(Or(jgt, Or(jeq, Or(jge, Or(jlt, Or(jne, Or(jle, jmp)))))))
	c.pc.Out(Not(inc), inc, rst, a)
	return
}
