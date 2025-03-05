package chip

func NotAnd(a, b Signal) Signal {
	if a == b && a == Active {
		return Inactive
	}
	return Active
}

func NotAnd16(a, b ReadonlyWord) *Word {
	r := NewWord()
	for i := range 16 {
		r.Set(i, NotAnd(a.Get(i), b.Get(i)))
	}
	return r
}

func Not(a Signal) Signal {
	return NotAnd(a, Active)
}

func Not16(a ReadonlyWord) *Word {
	r := NewWord()
	for i := range 16 {
		r.Set(i, Not(a.Get(i)))
	}
	return r
}

func And(a, b Signal) Signal {
	return Not(NotAnd(a, b))
}

func And16(a, b ReadonlyWord) *Word {
	r := NewWord()
	for i := range 16 {
		r.Set(i, And(a.Get(i), b.Get(i)))
	}
	return r
}

func And16To1(a Signal, b ReadonlyWord) *Word {
	r := NewWord()
	for i := range 16 {
		r.Set(i, And(a, b.Get(i)))
	}
	return r
}

func Or(a, b Signal) Signal {
	return NotAnd(Not(a), Not(b))
}

func Or16(a, b ReadonlyWord) *Word {
	r := NewWord()
	for i := range 16 {
		r.Set(i, Or(a.Get(i), b.Get(i)))
	}
	return r
}

func Or16To1(a Signal, b ReadonlyWord) *Word {
	r := NewWord()
	for i := range 16 {
		r.Set(i, Or(a, b.Get(i)))
	}
	return r
}

func Xor(a, b Signal) Signal {
	return Or(And(a, Not(b)), And(Not(a), b))
}

func Xor16(a, b ReadonlyWord) *Word {
	r := NewWord()
	for i := range 16 {
		r.Set(i, Xor(a.Get(i), b.Get(i)))
	}
	return r
}

func Xor16To1(a Signal, b ReadonlyWord) *Word {
	r := NewWord()
	for i := range 16 {
		r.Set(i, Xor(a, b.Get(i)))
	}
	return r
}

// Mux2Way16 provides a multiplexer for 2 inputs.
func Mux2Way16(s Signal, a, b ReadonlyWord) *Word {
	return Or16(And16To1(Not(s), a), And16To1(s, b))
}

// Mux2Way1 provides a multiplexer to two single bit inputs. More information on the multiplexer is given in the
// Mux2Way16 function comment.
func Mux2Way1(s Signal, a, b Signal) Signal {
	return Or(And(Not(s), a), And(s, b))
}

// Mux4Way16 provides a multiplexer for 4 inputs and a selector consisting of 2 bytes. Non-zero values on
// selector bytes are considered as set and only zero is considered unset.
func Mux4Way16(s [2]Signal, a, b, c, d ReadonlyWord) *Word {
	ab := Mux2Way16(s[1], a, b)
	cd := Mux2Way16(s[1], c, d)
	return Mux2Way16(s[0], ab, cd)
}

// Mux8Way16 provides a multiplexer for 8 inputs and a selector consisting of 3 bytes.
func Mux8Way16(s [3]Signal, a, b, c, d, e, f, g, h ReadonlyWord) *Word {
	abcd := Mux4Way16([2]Signal{s[1], s[2]}, a, b, c, d)
	efgh := Mux4Way16([2]Signal{s[1], s[2]}, e, f, g, h)
	return Mux2Way16(s[0], abcd, efgh)
}

// DMux2Way16 provides a demultiplexer for an [16]pin.Signal and a binary selector represented by a single byte.
func DMux2Way16(s Signal, in ReadonlyWord) (a, b *Word) {
	a = And16To1(Not(s), in)
	b = And16To1(s, in)
	return a, b
}

// DMux2Way1 provides a demultiplexer for a pin.Signal and a binary selector represented by a single byte.
func DMux2Way1(s Signal, in Signal) (a, b Signal) {
	a = And(Not(s), in)
	b = And(s, in)
	return a, b
}

// DMux4Way16 provides a demultiplexer for 4 outputs based on a 2-byte selector.
func DMux4Way16(s [2]Signal, in ReadonlyWord) (a, b, c, d *Word) {
	a, b = DMux2Way16(s[1], in)
	a = And16To1(Not(s[0]), a)
	b = And16To1(Not(s[0]), b)
	c, d = DMux2Way16(s[1], in)
	c = And16To1(s[0], c)
	d = And16To1(s[0], d)
	return a, b, c, d
}

// DMux4Way1 provides a de-multiplexer for 4 outputs based on a 2-byte selector.
func DMux4Way1(s [2]Signal, in Signal) (a, b, c, d Signal) {
	a, b = DMux2Way1(s[1], in)
	a = And(Not(s[0]), a)
	b = And(Not(s[0]), b)
	c, d = DMux2Way1(s[1], in)
	c = And(s[0], c)
	d = And(s[0], d)
	return a, b, c, d
}

// DMux8Way16 provides a de-multiplexer for 8 outputs based on a 3-byte selector.
func DMux8Way16(s [3]Signal, in ReadonlyWord) (a, b, c, d, e, f, g, h *Word) {
	a, b, c, d = DMux4Way16([2]Signal{s[1], s[2]}, in)
	a = And16To1(Not(s[0]), a)
	b = And16To1(Not(s[0]), b)
	c = And16To1(Not(s[0]), c)
	d = And16To1(Not(s[0]), d)
	e, f, g, h = DMux4Way16([2]Signal{s[1], s[2]}, in)
	e = And16To1(s[0], e)
	f = And16To1(s[0], f)
	g = And16To1(s[0], g)
	h = And16To1(s[0], h)
	return a, b, c, d, e, f, g, h
}

// DMux8Way1 provides a de-multiplexer for 8 outputs based on a 3-byte selector.
func DMux8Way1(s [3]Signal, in Signal) (a, b, c, d, e, f, g, h Signal) {
	a, b, c, d = DMux4Way1([2]Signal{s[1], s[2]}, in)
	a = And(Not(s[0]), a)
	b = And(Not(s[0]), b)
	c = And(Not(s[0]), c)
	d = And(Not(s[0]), d)
	e, f, g, h = DMux4Way1([2]Signal{s[1], s[2]}, in)
	e = And(s[0], e)
	f = And(s[0], f)
	g = And(s[0], g)
	h = And(s[0], h)
	return a, b, c, d, e, f, g, h
}

// DFF represents a data flip-flop capable of holding a single bit of information across CPU cycles.
type DFF struct {
	In  Signal
	out Signal
}

func (d *DFF) Out(load Signal) Signal {
	if load == Active {
		d.out = d.In
	}
	return d.out
}
