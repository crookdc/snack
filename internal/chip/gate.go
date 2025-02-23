package chip

func NotAnd(a, b Signal) Signal {
	if a == b && a == Inactive {
		return Active
	}
	return Inactive
}

func NotAnd16(a, b [16]Signal) [16]Signal {
	return [16]Signal{
		NotAnd(a[0], b[0]),
		NotAnd(a[1], b[1]),
		NotAnd(a[2], b[2]),
		NotAnd(a[3], b[3]),
		NotAnd(a[4], b[4]),
		NotAnd(a[5], b[5]),
		NotAnd(a[6], b[6]),
		NotAnd(a[7], b[7]),
		NotAnd(a[8], b[8]),
		NotAnd(a[9], b[9]),
		NotAnd(a[10], b[10]),
		NotAnd(a[11], b[11]),
		NotAnd(a[12], b[12]),
		NotAnd(a[13], b[13]),
		NotAnd(a[14], b[14]),
		NotAnd(a[15], b[15]),
	}
}

func Not(a Signal) Signal {
	return NotAnd(a, Inactive)
}

func Not16(a [16]Signal) [16]Signal {
	return [16]Signal{
		Not(a[0]),
		Not(a[1]),
		Not(a[2]),
		Not(a[3]),
		Not(a[4]),
		Not(a[5]),
		Not(a[6]),
		Not(a[7]),
		Not(a[8]),
		Not(a[9]),
		Not(a[10]),
		Not(a[11]),
		Not(a[12]),
		Not(a[13]),
		Not(a[14]),
		Not(a[15]),
	}
}

func And(a, b Signal) Signal {
	return NotAnd(Not(a), Not(b))
}

func And16(a, b [16]Signal) [16]Signal {
	return [16]Signal{
		And(a[0], b[0]),
		And(a[1], b[1]),
		And(a[2], b[2]),
		And(a[3], b[3]),
		And(a[4], b[4]),
		And(a[5], b[5]),
		And(a[6], b[6]),
		And(a[7], b[7]),
		And(a[8], b[8]),
		And(a[9], b[9]),
		And(a[10], b[10]),
		And(a[11], b[11]),
		And(a[12], b[12]),
		And(a[13], b[13]),
		And(a[14], b[14]),
		And(a[15], b[15]),
	}
}

func Or(a, b Signal) Signal {
	return Not(NotAnd(a, b))
}

func Or16(a, b [16]Signal) [16]Signal {
	return [16]Signal{
		Or(a[0], b[0]),
		Or(a[1], b[1]),
		Or(a[2], b[2]),
		Or(a[3], b[3]),
		Or(a[4], b[4]),
		Or(a[5], b[5]),
		Or(a[6], b[6]),
		Or(a[7], b[7]),
		Or(a[8], b[8]),
		Or(a[9], b[9]),
		Or(a[10], b[10]),
		Or(a[11], b[11]),
		Or(a[12], b[12]),
		Or(a[13], b[13]),
		Or(a[14], b[14]),
		Or(a[15], b[15]),
	}
}

func Xor(a, b Signal) Signal {
	return Or(And(a, Not(b)), And(Not(a), b))
}

func Xor16(a, b [16]Signal) [16]Signal {
	return [16]Signal{
		Xor(a[0], b[0]),
		Xor(a[1], b[1]),
		Xor(a[2], b[2]),
		Xor(a[3], b[3]),
		Xor(a[4], b[4]),
		Xor(a[5], b[5]),
		Xor(a[6], b[6]),
		Xor(a[7], b[7]),
		Xor(a[8], b[8]),
		Xor(a[9], b[9]),
		Xor(a[10], b[10]),
		Xor(a[11], b[11]),
		Xor(a[12], b[12]),
		Xor(a[13], b[13]),
		Xor(a[14], b[14]),
		Xor(a[15], b[15]),
	}
}

// Mux2Way16 provides a multiplexer for 2 inputs.
func Mux2Way16(s Signal, a, b [16]Signal) [16]Signal {
	return Or16(And16(Not16(expand16(s)), a), And16(expand16(s), b))
}

// Mux2Way1 provides a multiplexer to two single bit inputs. More information on the multiplexer is given in the
// Mux2Way16 function comment.
func Mux2Way1(s Signal, a, b Signal) Signal {
	return Or(And(Not(s), a), And(s, b))
}

// Mux4Way16 provides a multiplexer for 4 inputs and a selector consisting of 2 bytes. Non-zero values on
// selector bytes are considered as set and only zero is considered unset.
func Mux4Way16(s [2]Signal, a, b, c, d [16]Signal) [16]Signal {
	ab := Mux2Way16(s[1], a, b)
	cd := Mux2Way16(s[1], c, d)
	return Mux2Way16(s[0], ab, cd)
}

// Mux8Way16 provides a multiplexer for 8 inputs and a selector consisting of 3 bytes.
func Mux8Way16(s [3]Signal, a, b, c, d, e, f, g, h [16]Signal) [16]Signal {
	abcd := Mux4Way16([2]Signal{s[1], s[2]}, a, b, c, d)
	efgh := Mux4Way16([2]Signal{s[1], s[2]}, e, f, g, h)
	return Mux2Way16(s[0], abcd, efgh)
}

// DMux2Way16 provides a demultiplexer for an [16]pin.Signal and a binary selector represented by a single byte.
func DMux2Way16(s Signal, in [16]Signal) (a, b [16]Signal) {
	a = And16(Not16(expand16(s)), in)
	b = And16(expand16(s), in)
	return a, b
}

// DMux2Way1 provides a demultiplexer for a pin.Signal and a binary selector represented by a single byte.
func DMux2Way1(s Signal, in Signal) (a, b Signal) {
	a = And(Not(s), in)
	b = And(s, in)
	return a, b
}

// DMux4Way16 provides a demultiplexer for 4 outputs based on a 2-byte selector.
func DMux4Way16(s [2]Signal, in [16]Signal) (a, b, c, d [16]Signal) {
	sh := expand16(s[0])
	a, b = DMux2Way16(s[1], in)
	a = And16(Not16(sh), a)
	b = And16(Not16(sh), b)
	c, d = DMux2Way16(s[1], in)
	c = And16(sh, c)
	d = And16(sh, d)
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
func DMux8Way16(s [3]Signal, in [16]Signal) (a, b, c, d, e, f, g, h [16]Signal) {
	sh := expand16(s[0])
	a, b, c, d = DMux4Way16([2]Signal{s[1], s[2]}, in)
	a = And16(Not16(sh), a)
	b = And16(Not16(sh), b)
	c = And16(Not16(sh), c)
	d = And16(Not16(sh), d)
	e, f, g, h = DMux4Way16([2]Signal{s[1], s[2]}, in)
	e = And16(sh, e)
	f = And16(sh, f)
	g = And16(sh, g)
	h = And16(sh, h)
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
	In  Pin
	out Pin
}

func (d *DFF) Out(load Signal) Signal {
	if load == Active {
		d.out = d.In
	}
	return d.out.Signal()
}
