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

// RAM8 provides volatile storage of 8 words (16-bit values) that can be addressed with 3 pins.
type RAM8 struct {
	registers [8]Register
}

// Out either sets and returns or just returns the value for the provided address. When the load pin is active the in
// value is set on the provided address and then returned. When the load pin is inactive the value on the given
// address is just returned.
func (r *RAM8) Out(load Signal, addr [3]Signal, in ReadonlyWord) *Word {
	al, bl, cl, dl, el, fl, gl, hl := DMux8Way1(
		[3]Signal{addr[0], addr[1], addr[2]},
		load,
	)
	return Mux8Way16(
		[3]Signal{addr[0], addr[1], addr[2]},
		r.registers[0].Out(al, in),
		r.registers[1].Out(bl, in),
		r.registers[2].Out(cl, in),
		r.registers[3].Out(dl, in),
		r.registers[4].Out(el, in),
		r.registers[5].Out(fl, in),
		r.registers[6].Out(gl, in),
		r.registers[7].Out(hl, in),
	)
}

// RAM64 provides volatile storage of 64 words (16-bit values) that can be addressed with 6 pins.
type RAM64 struct {
	chips [8]RAM8
}

// Out either sets and returns or just returns the value for the provided address. When the load pin is active the in
// value is set on the provided address and then returned. When the load pin is inactive the value on the given
// address is just returned.
func (r *RAM64) Out(load Signal, addr [6]Signal, in ReadonlyWord) *Word {
	al, bl, cl, dl, el, fl, gl, hl := DMux8Way1(
		[3]Signal{addr[0], addr[1], addr[2]},
		load,
	)
	nxt := [3]Signal{addr[3], addr[4], addr[5]}
	return Mux8Way16(
		[3]Signal{addr[0], addr[1], addr[2]},
		r.chips[0].Out(al, nxt, in),
		r.chips[1].Out(bl, nxt, in),
		r.chips[2].Out(cl, nxt, in),
		r.chips[3].Out(dl, nxt, in),
		r.chips[4].Out(el, nxt, in),
		r.chips[5].Out(fl, nxt, in),
		r.chips[6].Out(gl, nxt, in),
		r.chips[7].Out(hl, nxt, in),
	)
}

// RAM512 provides volatile storage of 512 words (16-bit values) that can be addressed with 9 pins.
type RAM512 struct {
	chips [8]RAM64
}

// Out either sets and returns or just returns the value for the provided address. When the load pin is active the in
// value is set on the provided address and then returned. When the load pin is inactive the value on the given
// address is just returned.
func (r *RAM512) Out(load Signal, addr [9]Signal, in ReadonlyWord) *Word {
	al, bl, cl, dl, el, fl, gl, hl := DMux8Way1(
		[3]Signal{addr[0], addr[1], addr[2]},
		load,
	)
	nxt := [6]Signal{addr[3], addr[4], addr[5], addr[6], addr[7], addr[8]}
	return Mux8Way16(
		[3]Signal{addr[0], addr[1], addr[2]},
		r.chips[0].Out(al, nxt, in),
		r.chips[1].Out(bl, nxt, in),
		r.chips[2].Out(cl, nxt, in),
		r.chips[3].Out(dl, nxt, in),
		r.chips[4].Out(el, nxt, in),
		r.chips[5].Out(fl, nxt, in),
		r.chips[6].Out(gl, nxt, in),
		r.chips[7].Out(hl, nxt, in),
	)
}

// RAM4K provides volatile storage of 4096 words (16-bit values) that can be addressed with 12 pins.
type RAM4K struct {
	chips [8]RAM512
}

// Out either sets and returns or just returns the value for the provided address. When the load pin is active the in
// value is set on the provided address and then returned. When the load pin is inactive the value on the given
// address is just returned.
func (r *RAM4K) Out(load Signal, addr [12]Signal, in ReadonlyWord) *Word {
	al, bl, cl, dl, el, fl, gl, hl := DMux8Way1(
		[3]Signal{addr[0], addr[1], addr[2]},
		load,
	)
	nxt := [9]Signal{addr[3], addr[4], addr[5], addr[6], addr[7], addr[8], addr[9], addr[10], addr[11]}
	return Mux8Way16(
		[3]Signal{addr[0], addr[1], addr[2]},
		r.chips[0].Out(al, nxt, in),
		r.chips[1].Out(bl, nxt, in),
		r.chips[2].Out(cl, nxt, in),
		r.chips[3].Out(dl, nxt, in),
		r.chips[4].Out(el, nxt, in),
		r.chips[5].Out(fl, nxt, in),
		r.chips[6].Out(gl, nxt, in),
		r.chips[7].Out(hl, nxt, in),
	)
}

// RAM16K provides volatile storage of 16 384 words (16-bit values) that can be addressed with 14 pins.
type RAM16K struct {
	chips [4]RAM4K
}

// Out either sets and returns or just returns the value for the provided address. When the load pin is active the in
// value is set on the provided address and then returned. When the load pin is inactive the value on the given
// address is just returned.
func (r *RAM16K) Out(load Signal, addr [14]Signal, in ReadonlyWord) *Word {
	al, bl, cl, dl := DMux4Way1(
		[2]Signal{addr[0], addr[1]},
		load,
	)
	nxt := [12]Signal{addr[2], addr[3], addr[4], addr[5], addr[6], addr[7], addr[8], addr[9], addr[10], addr[11], addr[12], addr[13]}
	return Mux4Way16(
		[2]Signal{addr[0], addr[1]},
		r.chips[0].Out(al, nxt, in),
		r.chips[1].Out(bl, nxt, in),
		r.chips[2].Out(cl, nxt, in),
		r.chips[3].Out(dl, nxt, in),
	)
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
