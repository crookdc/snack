package chip

// Bit represents a digital snack.Pin that has been stored in a 1-bit register.
type Bit struct {
	dff DFF
}

func (b *Bit) Out(clk Pin, in Pin) Signal {
	b.dff.In.Set(Mux2Way1(clk.Signal(), b.dff.Out(clk.Signal()), in.Signal()))
	return b.dff.Out(clk.Signal())
}

// Register represents a simple array of 16 Bit coupled together to store a single 16 bit value.
type Register struct {
	bits [16]Bit
}

// Out reads the currently stored 16 bit value
func (r *Register) Out(clk Pin, in [16]Pin) [16]Signal {
	return [16]Signal{
		r.bits[0].Out(clk, in[0]),
		r.bits[1].Out(clk, in[1]),
		r.bits[2].Out(clk, in[2]),
		r.bits[3].Out(clk, in[3]),
		r.bits[4].Out(clk, in[4]),
		r.bits[5].Out(clk, in[5]),
		r.bits[6].Out(clk, in[6]),
		r.bits[7].Out(clk, in[7]),
		r.bits[8].Out(clk, in[8]),
		r.bits[9].Out(clk, in[9]),
		r.bits[10].Out(clk, in[10]),
		r.bits[11].Out(clk, in[11]),
		r.bits[12].Out(clk, in[12]),
		r.bits[13].Out(clk, in[13]),
		r.bits[14].Out(clk, in[14]),
		r.bits[15].Out(clk, in[15]),
	}
}

// RAM8 provides volatile storage of 8 words (16-bit values) that can be addressed with 3 pins.
type RAM8 struct {
	Registers [8]Register
}

// Out either sets and returns or just returns the value for the provided address. When the clk pin is active the in
// value is set on the provided address and then returned. When the clk pin is inactive the value on the given
// address is just returned.
func (r *RAM8) Out(clk Pin, addr [3]Pin, in [16]Pin) [16]Signal {
	al, bl, cl, dl, el, fl, gl, hl := DMux8Way1(
		[3]Signal{addr[0].Signal(), addr[1].Signal(), addr[2].Signal()},
		clk.Signal(),
	)
	return Mux8Way16(
		[3]Signal{addr[0].Signal(), addr[1].Signal(), addr[2].Signal()},
		r.Registers[0].Out(NewPin(al), in),
		r.Registers[1].Out(NewPin(bl), in),
		r.Registers[2].Out(NewPin(cl), in),
		r.Registers[3].Out(NewPin(dl), in),
		r.Registers[4].Out(NewPin(el), in),
		r.Registers[5].Out(NewPin(fl), in),
		r.Registers[6].Out(NewPin(gl), in),
		r.Registers[7].Out(NewPin(hl), in),
	)
}

// RAM64 provides volatile storage of 64 words (16-bit values) that can be addressed with 6 pins.
type RAM64 struct {
	Chips [8]RAM8
}

// Out either sets and returns or just returns the value for the provided address. When the clk pin is active the in
// value is set on the provided address and then returned. When the clk pin is inactive the value on the given
// address is just returned.
func (r *RAM64) Out(clk Pin, addr [6]Pin, in [16]Pin) [16]Signal {
	al, bl, cl, dl, el, fl, gl, hl := DMux8Way1(
		[3]Signal{addr[0].Signal(), addr[1].Signal(), addr[2].Signal()},
		clk.Signal(),
	)
	nxt := [3]Pin{addr[3], addr[4], addr[5]}
	return Mux8Way16(
		[3]Signal{addr[0].Signal(), addr[1].Signal(), addr[2].Signal()},
		r.Chips[0].Out(NewPin(al), nxt, in),
		r.Chips[1].Out(NewPin(bl), nxt, in),
		r.Chips[2].Out(NewPin(cl), nxt, in),
		r.Chips[3].Out(NewPin(dl), nxt, in),
		r.Chips[4].Out(NewPin(el), nxt, in),
		r.Chips[5].Out(NewPin(fl), nxt, in),
		r.Chips[6].Out(NewPin(gl), nxt, in),
		r.Chips[7].Out(NewPin(hl), nxt, in),
	)
}

// RAM512 provides volatile storage of 512 words (16-bit values) that can be addressed with 9 pins.
type RAM512 struct {
	Chips [8]RAM64
}

// Out either sets and returns or just returns the value for the provided address. When the clk pin is active the in
// value is set on the provided address and then returned. When the clk pin is inactive the value on the given
// address is just returned.
func (r *RAM512) Out(clk Pin, addr [9]Pin, in [16]Pin) [16]Signal {
	al, bl, cl, dl, el, fl, gl, hl := DMux8Way1(
		[3]Signal{addr[0].Signal(), addr[1].Signal(), addr[2].Signal()},
		clk.Signal(),
	)
	nxt := [6]Pin{addr[3], addr[4], addr[5], addr[6], addr[7], addr[8]}
	return Mux8Way16(
		[3]Signal{addr[0].Signal(), addr[1].Signal(), addr[2].Signal()},
		r.Chips[0].Out(NewPin(al), nxt, in),
		r.Chips[1].Out(NewPin(bl), nxt, in),
		r.Chips[2].Out(NewPin(cl), nxt, in),
		r.Chips[3].Out(NewPin(dl), nxt, in),
		r.Chips[4].Out(NewPin(el), nxt, in),
		r.Chips[5].Out(NewPin(fl), nxt, in),
		r.Chips[6].Out(NewPin(gl), nxt, in),
		r.Chips[7].Out(NewPin(hl), nxt, in),
	)
}

// RAM4K provides volatile storage of 4096 words (16-bit values) that can be addressed with 12 pins.
type RAM4K struct {
	Chips [8]RAM512
}

// Out either sets and returns or just returns the value for the provided address. When the clk pin is active the in
// value is set on the provided address and then returned. When the clk pin is inactive the value on the given
// address is just returned.
func (r *RAM4K) Out(clk Pin, addr [12]Pin, in [16]Pin) [16]Signal {
	al, bl, cl, dl, el, fl, gl, hl := DMux8Way1(
		[3]Signal{addr[0].Signal(), addr[1].Signal(), addr[2].Signal()},
		clk.Signal(),
	)
	nxt := [9]Pin{addr[3], addr[4], addr[5], addr[6], addr[7], addr[8], addr[9], addr[10], addr[11]}
	return Mux8Way16(
		[3]Signal{addr[0].Signal(), addr[1].Signal(), addr[2].Signal()},
		r.Chips[0].Out(NewPin(al), nxt, in),
		r.Chips[1].Out(NewPin(bl), nxt, in),
		r.Chips[2].Out(NewPin(cl), nxt, in),
		r.Chips[3].Out(NewPin(dl), nxt, in),
		r.Chips[4].Out(NewPin(el), nxt, in),
		r.Chips[5].Out(NewPin(fl), nxt, in),
		r.Chips[6].Out(NewPin(gl), nxt, in),
		r.Chips[7].Out(NewPin(hl), nxt, in),
	)
}

// RAM16K provides volatile storage of 16 384 words (16-bit values) that can be addressed with 14 pins.
type RAM16K struct {
	Chips [4]RAM4K
}

// Out either sets and returns or just returns the value for the provided address. When the clk pin is active the in
// value is set on the provided address and then returned. When the clk pin is inactive the value on the given
// address is just returned.
func (r *RAM16K) Out(clk Pin, addr [14]Pin, in [16]Pin) [16]Signal {
	al, bl, cl, dl := DMux4Way1(
		[2]Signal{addr[0].Signal(), addr[1].Signal()},
		clk.Signal(),
	)
	nxt := [12]Pin{addr[2], addr[3], addr[4], addr[5], addr[6], addr[7], addr[8], addr[9], addr[10], addr[11], addr[12], addr[13]}
	return Mux4Way16(
		[2]Signal{addr[0].Signal(), addr[1].Signal()},
		r.Chips[0].Out(NewPin(al), nxt, in),
		r.Chips[1].Out(NewPin(bl), nxt, in),
		r.Chips[2].Out(NewPin(cl), nxt, in),
		r.Chips[3].Out(NewPin(dl), nxt, in),
	)
}

// Counter provides a chip with the ability to store a single word as well as increment its value and reset it to 0.
type Counter struct {
	register Register
}

// Out allows setting of the counters current value by providing a value in the 16-pin parameter `in` and setting the
// clk to an active pin. To increment the stored value the inc pin must only be set. Finally, to reset the value the rst
// pin must be active.
func (c *Counter) Out(clk Pin, inc Pin, rst Pin, in [16]Pin) [16]Signal {
	out := c.register.Out(clk, in)
	out = Adder16(out, [16]Signal{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, And(Not(clk.Signal()), inc.Signal())})
	out = And16(out, expand16(Not(And(Not(clk.Signal()), rst.Signal()))))
	return c.register.Out(NewPin(Active), NewPin16(out))
}
