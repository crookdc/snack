package chip

// Bit represents a digital snack.Signal that has been stored in a 1-bit register.
type Bit struct {
	dff DFF
}

func (b *Bit) Out(load Signal, in Signal) Signal {
	b.dff.In = Mux2Way1(load, b.dff.Out(load), in)
	return b.dff.Out(load)
}

// Register represents a simple array of 16 Bit coupled together to store a single 16 bit value.
type Register struct {
	bits [16]Bit
}

// Out reads the currently stored 16 bit value
func (r *Register) Out(load Signal, in ReadonlyWord) *Word {
	w := NewWord()
	for i := range 16 {
		w.Set(i, r.bits[i].Out(load, in.Get(i)))
	}
	return w
}

// PC provides a chip with the ability to store a single word as well as increment its value and reset it to 0.
type PC struct {
	register Register
}

// Out allows setting of the counters current value by providing a value in the 16-pin parameter `in` and setting the
// load to an active pin. To increment the stored value the inc pin must only be set. Finally, to reset the value the rst
// pin must be active.
func (c *PC) Out(load Signal, inc Signal, rst Signal, in ReadonlyWord) *Word {
	out := c.register.Out(load, And16To1(Not(inc), in))
	out = Adder16(out, Wrap(&[16]Signal{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, And(Not(load), inc)}))
	out = And16To1(Not(And(Not(load), rst)), out)
	return c.register.Out(Active, out)
}
