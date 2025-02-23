package chip

// Bit represents a digital snack.Pin that has been stored in a 1-bit register.
type Bit struct {
	dff DFF
}

func (b *Bit) Out(load Pin, in Pin) Signal {
	b.dff.In.Set(Mux2Way1(load.Signal(), b.dff.Out(load.Signal()), in.Signal()))
	return b.dff.Out(load.Signal())
}

// Register represents a simple array of 16 Bit coupled together to store a single 16 bit value.
type Register struct {
	bits [16]Bit
}

// Out reads the currently stored 16 bit value
func (r *Register) Out(load Pin, in [16]Pin) [16]Signal {
	return [16]Signal{
		r.bits[0].Out(load, in[0]),
		r.bits[1].Out(load, in[1]),
		r.bits[2].Out(load, in[2]),
		r.bits[3].Out(load, in[3]),
		r.bits[4].Out(load, in[4]),
		r.bits[5].Out(load, in[5]),
		r.bits[6].Out(load, in[6]),
		r.bits[7].Out(load, in[7]),
		r.bits[8].Out(load, in[8]),
		r.bits[9].Out(load, in[9]),
		r.bits[10].Out(load, in[10]),
		r.bits[11].Out(load, in[11]),
		r.bits[12].Out(load, in[12]),
		r.bits[13].Out(load, in[13]),
		r.bits[14].Out(load, in[14]),
		r.bits[15].Out(load, in[15]),
	}
}

// RAM8 provides volatile storage of 8 words (16-bit values) that can be addressed with 3 pins.
type RAM8 struct {
	registers [8]Register
}

// Out either sets and returns or just returns the value for the provided address. When the load pin is active the in
// value is set on the provided address and then returned. When the load pin is inactive the value on the given
// address is just returned.
func (r *RAM8) Out(load Pin, addr [3]Pin, in [16]Pin) [16]Signal {
	al, bl, cl, dl, el, fl, gl, hl := DMux8Way1(
		[3]Signal{addr[0].Signal(), addr[1].Signal(), addr[2].Signal()},
		load.Signal(),
	)
	return Mux8Way16(
		[3]Signal{addr[0].Signal(), addr[1].Signal(), addr[2].Signal()},
		r.registers[0].Out(NewPin(al), in),
		r.registers[1].Out(NewPin(bl), in),
		r.registers[2].Out(NewPin(cl), in),
		r.registers[3].Out(NewPin(dl), in),
		r.registers[4].Out(NewPin(el), in),
		r.registers[5].Out(NewPin(fl), in),
		r.registers[6].Out(NewPin(gl), in),
		r.registers[7].Out(NewPin(hl), in),
	)
}

// RAM64 provides volatile storage of 64 words (16-bit values) that can be addressed with 6 pins.
type RAM64 struct {
	chips [8]RAM8
}

// Out either sets and returns or just returns the value for the provided address. When the load pin is active the in
// value is set on the provided address and then returned. When the load pin is inactive the value on the given
// address is just returned.
func (r *RAM64) Out(load Pin, addr [6]Pin, in [16]Pin) [16]Signal {
	al, bl, cl, dl, el, fl, gl, hl := DMux8Way1(
		[3]Signal{addr[0].Signal(), addr[1].Signal(), addr[2].Signal()},
		load.Signal(),
	)
	nxt := [3]Pin{addr[3], addr[4], addr[5]}
	return Mux8Way16(
		[3]Signal{addr[0].Signal(), addr[1].Signal(), addr[2].Signal()},
		r.chips[0].Out(NewPin(al), nxt, in),
		r.chips[1].Out(NewPin(bl), nxt, in),
		r.chips[2].Out(NewPin(cl), nxt, in),
		r.chips[3].Out(NewPin(dl), nxt, in),
		r.chips[4].Out(NewPin(el), nxt, in),
		r.chips[5].Out(NewPin(fl), nxt, in),
		r.chips[6].Out(NewPin(gl), nxt, in),
		r.chips[7].Out(NewPin(hl), nxt, in),
	)
}

// RAM512 provides volatile storage of 512 words (16-bit values) that can be addressed with 9 pins.
type RAM512 struct {
	chips [8]RAM64
}

// Out either sets and returns or just returns the value for the provided address. When the load pin is active the in
// value is set on the provided address and then returned. When the load pin is inactive the value on the given
// address is just returned.
func (r *RAM512) Out(load Pin, addr [9]Pin, in [16]Pin) [16]Signal {
	al, bl, cl, dl, el, fl, gl, hl := DMux8Way1(
		[3]Signal{addr[0].Signal(), addr[1].Signal(), addr[2].Signal()},
		load.Signal(),
	)
	nxt := [6]Pin{addr[3], addr[4], addr[5], addr[6], addr[7], addr[8]}
	return Mux8Way16(
		[3]Signal{addr[0].Signal(), addr[1].Signal(), addr[2].Signal()},
		r.chips[0].Out(NewPin(al), nxt, in),
		r.chips[1].Out(NewPin(bl), nxt, in),
		r.chips[2].Out(NewPin(cl), nxt, in),
		r.chips[3].Out(NewPin(dl), nxt, in),
		r.chips[4].Out(NewPin(el), nxt, in),
		r.chips[5].Out(NewPin(fl), nxt, in),
		r.chips[6].Out(NewPin(gl), nxt, in),
		r.chips[7].Out(NewPin(hl), nxt, in),
	)
}

// RAM4K provides volatile storage of 4096 words (16-bit values) that can be addressed with 12 pins.
type RAM4K struct {
	chips [8]RAM512
}

// Out either sets and returns or just returns the value for the provided address. When the load pin is active the in
// value is set on the provided address and then returned. When the load pin is inactive the value on the given
// address is just returned.
func (r *RAM4K) Out(load Pin, addr [12]Pin, in [16]Pin) [16]Signal {
	al, bl, cl, dl, el, fl, gl, hl := DMux8Way1(
		[3]Signal{addr[0].Signal(), addr[1].Signal(), addr[2].Signal()},
		load.Signal(),
	)
	nxt := [9]Pin{addr[3], addr[4], addr[5], addr[6], addr[7], addr[8], addr[9], addr[10], addr[11]}
	return Mux8Way16(
		[3]Signal{addr[0].Signal(), addr[1].Signal(), addr[2].Signal()},
		r.chips[0].Out(NewPin(al), nxt, in),
		r.chips[1].Out(NewPin(bl), nxt, in),
		r.chips[2].Out(NewPin(cl), nxt, in),
		r.chips[3].Out(NewPin(dl), nxt, in),
		r.chips[4].Out(NewPin(el), nxt, in),
		r.chips[5].Out(NewPin(fl), nxt, in),
		r.chips[6].Out(NewPin(gl), nxt, in),
		r.chips[7].Out(NewPin(hl), nxt, in),
	)
}

// RAM16K provides volatile storage of 16 384 words (16-bit values) that can be addressed with 14 pins.
type RAM16K struct {
	chips [4]RAM4K
}

// Out either sets and returns or just returns the value for the provided address. When the load pin is active the in
// value is set on the provided address and then returned. When the load pin is inactive the value on the given
// address is just returned.
func (r *RAM16K) Out(load Pin, addr [14]Pin, in [16]Pin) [16]Signal {
	al, bl, cl, dl := DMux4Way1(
		[2]Signal{addr[0].Signal(), addr[1].Signal()},
		load.Signal(),
	)
	nxt := [12]Pin{addr[2], addr[3], addr[4], addr[5], addr[6], addr[7], addr[8], addr[9], addr[10], addr[11], addr[12], addr[13]}
	return Mux4Way16(
		[2]Signal{addr[0].Signal(), addr[1].Signal()},
		r.chips[0].Out(NewPin(al), nxt, in),
		r.chips[1].Out(NewPin(bl), nxt, in),
		r.chips[2].Out(NewPin(cl), nxt, in),
		r.chips[3].Out(NewPin(dl), nxt, in),
	)
}

// ProgramCounter provides a chip with the ability to store a single word as well as increment its value and reset it to 0.
type ProgramCounter struct {
	register Register
}

// Out allows setting of the counters current value by providing a value in the 16-pin parameter `in` and setting the
// load to an active pin. To increment the stored value the inc pin must only be set. Finally, to reset the value the rst
// pin must be active.
func (c *ProgramCounter) Out(load Pin, inc Pin, rst Pin, in [16]Pin) [16]Signal {
	out := c.register.Out(load, NewPin16(And16(read16(in), Not16(expand16(inc.Signal())))))
	out = Adder16(out, [16]Signal{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, And(Not(load.Signal()), inc.Signal())})
	out = And16(out, expand16(Not(And(Not(load.Signal()), rst.Signal()))))
	return c.register.Out(NewPin(Active), NewPin16(out))
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
