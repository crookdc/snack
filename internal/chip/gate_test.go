package chip

import (
	"fmt"
	"testing"
)

func TestNotAnd(t *testing.T) {
	var assertions = []struct {
		a Signal
		b Signal
		r Signal
	}{
		{
			a: Inactive,
			b: Inactive,
			r: Active,
		},
		{
			a: Inactive,
			b: Active,
			r: Active,
		},
		{
			a: Active,
			b: Inactive,
			r: Active,
		},
		{
			a: Active,
			b: Active,
			r: Inactive,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given a is %v and b is %v", a.a, a.b), func(t *testing.T) {
			r := NotAnd(a.a, a.b)
			if r != a.r {
				t.Errorf("expected %v with a: %v and b: %v but got %v", a.r, a.a, a.b, r)
			}
		})
	}
}

func TestNotAndUint16(t *testing.T) {
	type assertion struct {
		a uint16
		b uint16
		r uint16
	}
	var assertions = []assertion{
		{
			a: 0x0000,
			b: 0x0000,
			r: 0xFFFF,
		},
		{
			a: 0xF0F0,
			b: 0x0F0F,
			r: 0xFFFF,
		},
		{
			a: 0x0F00,
			b: 0xF000,
			r: 0xFFFF,
		},
		{
			a: 0xFFFF,
			b: 0xFFFF,
			r: 0x0000,
		},
		{
			a: 0xFFFF,
			b: 0x0000,
			r: 0xFFFF,
		},
		{
			a: 0x0000,
			b: 0xFFFF,
			r: 0xFFFF,
		},
		{
			a: 0b1001_0001_0011_1111,
			b: 0b0000_0000_0000_0000,
			r: 0b1111_1111_1111_1111,
		},
	}
	for _, assert := range assertions {
		t.Run(fmt.Sprintf("given a is %v and b is %v", assert.a, assert.b), func(t *testing.T) {
			a := split16(assert.a)
			b := split16(assert.b)
			r := NotAnd16(Wrap(&a), Wrap(&b))
			if r.Copy() != split16(assert.r) {
				t.Errorf("expected %v but got %v", r, assert.r)
			}
		})
	}
}

func TestNot(t *testing.T) {
	var assertions = []struct {
		a Signal
		r Signal
	}{
		{
			a: Inactive,
			r: Active,
		},
		{
			a: Active,
			r: Inactive,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given a is %v", a.a), func(t *testing.T) {
			r := Not(a.a)
			if r != a.r {
				t.Errorf("expected %v with a: %v but got %v", a.r, a.a, r)
			}
		})
	}
}

func TestNotUint16(t *testing.T) {
	type assertion struct {
		a uint16
		r uint16
	}
	var assertions = []assertion{
		{
			a: 0x0000,
			r: 0xFFFF,
		},
		{
			a: 0xF0F0,
			r: 0x0F0F,
		},
		{
			a: 0x0F0F,
			r: 0xF0F0,
		},
		{
			a: 0xFFFF,
			r: 0x0000,
		},
	}
	for _, assert := range assertions {
		t.Run(fmt.Sprintf("given a is %v", assert.a), func(t *testing.T) {
			a := split16(assert.a)
			r := Not16(Wrap(&a))
			if r.Copy() != split16(assert.r) {
				t.Errorf("expected %v but got %v", assert.r, r)
			}
		})
	}
}

func TestAnd(t *testing.T) {
	var assertions = []struct {
		a Signal
		b Signal
		r Signal
	}{
		{
			a: Inactive,
			b: Inactive,
			r: Inactive,
		},
		{
			a: Inactive,
			b: Active,
			r: Inactive,
		},
		{
			a: Active,
			b: Inactive,
			r: Inactive,
		},
		{
			a: Active,
			b: Active,
			r: Active,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given a is %v and b is %v", a.a, a.b), func(t *testing.T) {
			r := And(a.a, a.b)
			if r != a.r {
				t.Errorf("expected %v with a: %v and b: %v but got %v", a.r, a.a, a.b, r)
			}
		})
	}
}

func TestAndUint16(t *testing.T) {
	type assertion struct {
		a uint16
		b uint16
		r uint16
	}
	var assertions = []assertion{
		{
			a: 0,
			b: 0,
			r: 0,
		},
		{
			a: 0xFFFF,
			b: 0,
			r: 0,
		},
		{
			a: 0xF0F0,
			b: 0xF000,
			r: 0xF000,
		},
		{
			a: 0xFFFF,
			b: 0xFFFF,
			r: 0xFFFF,
		},
	}
	for _, assert := range assertions {
		t.Run(fmt.Sprintf("given a is %v and b is %v", assert.a, assert.b), func(t *testing.T) {
			a := split16(assert.a)
			b := split16(assert.b)
			r := And16(Wrap(&a), Wrap(&b))
			if r.Copy() != split16(assert.r) {
				t.Errorf("expected %v but got %v", r, assert.r)
			}
		})
	}
}

func TestOr(t *testing.T) {
	var assertions = []struct {
		a Signal
		b Signal
		r Signal
	}{
		{
			a: Inactive,
			b: Inactive,
			r: Inactive,
		},
		{
			a: Inactive,
			b: Active,
			r: Active,
		},
		{
			a: Active,
			b: Inactive,
			r: Active,
		},
		{
			a: Active,
			b: Active,
			r: Active,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given a is %v and b is %v", a.a, a.b), func(t *testing.T) {
			r := Or(a.a, a.b)
			if r != a.r {
				t.Errorf("expected %v with a: %v and b: %v but got %v", a.r, a.a, a.b, r)
			}
		})
	}
}

func TestOrUint16(t *testing.T) {
	type assertion struct {
		a uint16
		b uint16
		r uint16
	}
	var assertions = []assertion{
		{
			a: 0,
			b: 0,
			r: 0,
		},
		{
			a: 0xFFFF,
			b: 0,
			r: 0xFFFF,
		},
		{
			a: 0x0F0F,
			b: 0xFFFF,
			r: 0xFFFF,
		},
		{
			a: 0xF000,
			b: 0xF00F,
			r: 0xF00F,
		},
	}
	for _, assert := range assertions {
		t.Run(fmt.Sprintf("given a is %v and b is %v", assert.a, assert.b), func(t *testing.T) {
			a := split16(assert.a)
			b := split16(assert.b)
			r := Or16(Wrap(&a), Wrap(&b))
			if r.Copy() != split16(assert.r) {
				t.Errorf("expected %v but got %v", assert.r, r)
			}
		})
	}
}

func TestXor(t *testing.T) {
	var assertions = []struct {
		a Signal
		b Signal
		r Signal
	}{
		{
			a: Inactive,
			b: Inactive,
			r: Inactive,
		},
		{
			a: Inactive,
			b: Active,
			r: Active,
		},
		{
			a: Active,
			b: Inactive,
			r: Active,
		},
		{
			a: Active,
			b: Active,
			r: Inactive,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given a is %v and b is %v", a.a, a.b), func(t *testing.T) {
			r := Xor(a.a, a.b)
			if r != a.r {
				t.Errorf("expected %v with a: %v and b: %v but got %v", a.r, a.a, a.b, r)
			}
		})
	}
}

func TestXorUint16(t *testing.T) {
	type assertion struct {
		a uint16
		b uint16
		r uint16
	}
	var assertions = []assertion{
		{
			a: 0xFFFF,
			b: 0xFFFF,
			r: 0,
		},
		{
			a: 0xF0F0,
			b: 0x0F0F,
			r: 0xFFFF,
		},
		{
			a: 0xF0F0,
			b: 0xF0F0,
			r: 0,
		},
	}
	for _, assert := range assertions {
		t.Run(fmt.Sprintf("given a is %v and b is %v", assert.a, assert.b), func(t *testing.T) {
			a := split16(assert.a)
			b := split16(assert.b)
			r := Xor16(Wrap(&a), Wrap(&b))
			if r.Copy() != split16(assert.r) {
				t.Errorf("expected %v but got %v", assert.r, r)
			}
		})
	}
}

func TestMux2Way16(t *testing.T) {
	var assertions = []struct {
		a   uint16
		b   uint16
		sel Signal
		r   uint16
	}{
		{
			a:   55,
			b:   96,
			sel: Inactive,
			r:   55,
		},
		{
			a:   0,
			b:   53,
			sel: Inactive,
			r:   0,
		},
		{
			a:   255,
			b:   12,
			sel: Inactive,
			r:   255,
		},
		{
			a:   0,
			b:   0,
			sel: Inactive,
			r:   0,
		},
		{
			a:   12,
			b:   0,
			sel: Active,
			r:   0,
		},
		{
			a:   123,
			b:   99,
			sel: Active,
			r:   99,
		},
		{
			a:   0,
			b:   123,
			sel: Active,
			r:   123,
		},
		{
			a:   0,
			b:   0,
			sel: Active,
			r:   0,
		},
	}
	for _, assert := range assertions {
		t.Run(fmt.Sprintf("given a is %v, b is %v and sel is %v", assert.a, assert.b, assert.sel), func(t *testing.T) {
			a := split16(assert.a)
			b := split16(assert.b)
			r := Mux2Way16(assert.sel, Wrap(&a), Wrap(&b))
			if r.Copy() != split16(assert.r) {
				t.Errorf("expected %v but got %v", assert.r, r)
			}
		})
	}
}

func TestMux4Way16(t *testing.T) {
	type assertion struct {
		a uint16
		b uint16
		c uint16
		d uint16
		s [2]Signal
		r uint16
	}
	var assertions = []assertion{
		{
			a: 10,
			b: 20,
			c: 30,
			d: 40,
			s: [2]Signal{Inactive, Inactive},
			r: 10,
		},
		{
			a: 10,
			b: 20,
			c: 30,
			d: 40,
			s: [2]Signal{Inactive, Active},
			r: 20,
		},
		{
			a: 10,
			b: 20,
			c: 30,
			d: 40,
			s: [2]Signal{Active, Inactive},
			r: 30,
		},
		{
			a: 10,
			b: 20,
			c: 30,
			d: 40,
			s: [2]Signal{Active, Active},
			r: 40,
		},
	}
	for _, assert := range assertions {
		t.Run(fmt.Sprintf("given s: %v", assert.s), func(t *testing.T) {
			a := split16(assert.a)
			b := split16(assert.b)
			c := split16(assert.c)
			d := split16(assert.d)

			r := Mux4Way16(
				assert.s,
				Wrap(&a),
				Wrap(&b),
				Wrap(&c),
				Wrap(&d),
			)
			if r.Copy() != split16(assert.r) {
				t.Errorf("expected %v given s: %v but got %v", assert.r, assert.s, r)
			}
		})
	}
}

func TestMux8Way16(t *testing.T) {
	type assertion struct {
		a uint16
		b uint16
		c uint16
		d uint16
		e uint16
		f uint16
		g uint16
		h uint16

		s [3]Signal
		r uint16
	}
	var assertions = []assertion{
		{
			a: 10,
			b: 20,
			c: 30,
			d: 40,
			e: 50,
			f: 60,
			g: 70,
			h: 80,

			s: [3]Signal{Inactive, Inactive, Inactive},
			r: 10,
		},
		{
			a: 10,
			b: 20,
			c: 30,
			d: 40,
			e: 50,
			f: 60,
			g: 70,
			h: 80,

			s: [3]Signal{Inactive, Inactive, Active},
			r: 20,
		},
		{
			a: 10,
			b: 20,
			c: 30,
			d: 40,
			e: 50,
			f: 60,
			g: 70,
			h: 80,

			s: [3]Signal{Inactive, Active, Inactive},
			r: 30,
		},
		{
			a: 10,
			b: 20,
			c: 30,
			d: 40,
			e: 50,
			f: 60,
			g: 70,
			h: 80,

			s: [3]Signal{Inactive, Active, Active},
			r: 40,
		},
		{
			a: 10,
			b: 20,
			c: 30,
			d: 40,
			e: 50,
			f: 60,
			g: 70,
			h: 80,

			s: [3]Signal{Active, Inactive, Inactive},
			r: 50,
		},
		{
			a: 10,
			b: 20,
			c: 30,
			d: 40,
			e: 50,
			f: 60,
			g: 70,
			h: 80,

			s: [3]Signal{Active, Inactive, Active},
			r: 60,
		},
		{
			a: 10,
			b: 20,
			c: 30,
			d: 40,
			e: 50,
			f: 60,
			g: 70,
			h: 80,

			s: [3]Signal{Active, Active, Inactive},
			r: 70,
		},
		{
			a: 10,
			b: 20,
			c: 30,
			d: 40,
			e: 50,
			f: 60,
			g: 70,
			h: 80,

			s: [3]Signal{Active, Active, Active},
			r: 80,
		},
	}
	for _, assert := range assertions {
		t.Run(fmt.Sprintf("with s: %v", assert.s), func(t *testing.T) {
			a := split16(assert.a)
			b := split16(assert.b)
			c := split16(assert.c)
			d := split16(assert.d)
			e := split16(assert.e)
			f := split16(assert.f)
			g := split16(assert.g)
			h := split16(assert.h)
			r := Mux8Way16(
				assert.s,
				Wrap(&a),
				Wrap(&b),
				Wrap(&c),
				Wrap(&d),
				Wrap(&e),
				Wrap(&f),
				Wrap(&g),
				Wrap(&h),
			)
			if r.Copy() != split16(assert.r) {
				t.Errorf("expected %v with s: %v got got %v", assert.r, assert.s, r)
			}
		})
	}
}

func TestDemux2Way(t *testing.T) {
	type assertion struct {
		in uint16
		s  Signal
		a  uint16
		b  uint16
	}
	var assertions = []assertion{
		{
			in: 0,
			s:  Inactive,
			a:  0,
			b:  0,
		},
		{
			in: 65_535,
			s:  Inactive,
			a:  65_535,
			b:  0,
		},
		{
			in: 256,
			s:  Active,
			a:  0,
			b:  256,
		},
		{
			in: 0,
			s:  Active,
			a:  0,
			b:  0,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given in is %v and s is %v", a.in, a.s), func(t *testing.T) {
			in := split16(a.in)
			ar, br := DMux2Way16(a.s, Wrap(&in))
			if ar.Copy() != split16(a.a) {
				t.Errorf("expected ar %v but got %v", a.a, ar)
			}
			if br.Copy() != split16(a.b) {
				t.Errorf("expected br %v but got %v", a.b, br)
			}
		})
	}
}

func TestDemux4Way(t *testing.T) {
	type assertion struct {
		in uint16
		s  [2]Signal
		a  uint16
		b  uint16
		c  uint16
		d  uint16
	}
	var assertions = []assertion{
		{
			in: 65_313,
			s:  [2]Signal{Inactive, Inactive},
			a:  65_313,
			b:  0,
			c:  0,
			d:  0,
		},
		{
			in: 23_230,
			s:  [2]Signal{Inactive, Active},
			a:  0,
			b:  23_230,
			c:  0,
			d:  0,
		},
		{
			in: 9012,
			s:  [2]Signal{Active, Inactive},
			a:  0,
			b:  0,
			c:  9012,
			d:  0,
		},
		{
			in: 1234,
			s:  [2]Signal{Active, Active},
			a:  0,
			b:  0,
			c:  0,
			d:  1234,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given in %v and s %v", a.in, a.s), func(t *testing.T) {
			in := split16(a.in)
			ar, br, cr, dr := DMux4Way16(a.s, Wrap(&in))
			if ar.Copy() != split16(a.a) {
				t.Errorf("expected a %v but got %v", a.a, ar)
			}
			if br.Copy() != split16(a.b) {
				t.Errorf("expected b %v but got %v", a.b, br)
			}
			if cr.Copy() != split16(a.c) {
				t.Errorf("expected c %v but got %v", a.c, cr)
			}
			if dr.Copy() != split16(a.d) {
				t.Errorf("expected dout %v but got %v", a.d, dr)
			}
		})
	}
}

func TestDemux8Way(t *testing.T) {
	type assertion struct {
		in uint16
		s  [3]Signal
		a  uint16
		b  uint16
		c  uint16
		d  uint16
		e  uint16
		f  uint16
		g  uint16
		h  uint16
	}
	var assertions = []assertion{
		{
			in: 65_313,
			s:  [3]Signal{Inactive, Inactive, Inactive},
			a:  65_313,
			b:  0,
			c:  0,
			d:  0,
			e:  0,
			f:  0,
			g:  0,
			h:  0,
		},
		{
			in: 56_555,
			s:  [3]Signal{Inactive, Inactive, Active},
			a:  0,
			b:  56_555,
			c:  0,
			d:  0,
			e:  0,
			f:  0,
			g:  0,
			h:  0,
		},
		{
			in: 1234,
			s:  [3]Signal{Inactive, Active, Inactive},
			a:  0,
			b:  0,
			c:  1234,
			d:  0,
			e:  0,
			f:  0,
			g:  0,
			h:  0,
		},
		{
			in: 9999,
			s:  [3]Signal{Inactive, Active, Active},
			a:  0,
			b:  0,
			c:  0,
			d:  9999,
			e:  0,
			f:  0,
			g:  0,
			h:  0,
		},
		{
			in: 8989,
			s:  [3]Signal{Active, Inactive, Inactive},
			a:  0,
			b:  0,
			c:  0,
			d:  0,
			e:  8989,
			f:  0,
			g:  0,
			h:  0,
		},
		{
			in: 13372,
			s:  [3]Signal{Active, Inactive, Active},
			a:  0,
			b:  0,
			c:  0,
			d:  0,
			e:  0,
			f:  13372,
			g:  0,
			h:  0,
		},
		{
			in: 12341,
			s:  [3]Signal{Active, Active, Inactive},
			a:  0,
			b:  0,
			c:  0,
			d:  0,
			e:  0,
			f:  0,
			g:  12341,
			h:  0,
		},
		{
			in: 4455,
			s:  [3]Signal{Active, Active, Active},
			a:  0,
			b:  0,
			c:  0,
			d:  0,
			e:  0,
			f:  0,
			g:  0,
			h:  4455,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given in %v and s %v", a.in, a.s), func(t *testing.T) {
			in := split16(a.in)
			ar, br, cr, dr, er, fr, gr, hr := DMux8Way16(a.s, Wrap(&in))
			if ar.Copy() != split16(a.a) {
				t.Errorf("expected a %v but got %v", a.a, ar)
			}
			if br.Copy() != split16(a.b) {
				t.Errorf("expected b %v but got %v", a.b, br)
			}
			if cr.Copy() != split16(a.c) {
				t.Errorf("expected c %v but got %v", a.c, cr)
			}
			if dr.Copy() != split16(a.d) {
				t.Errorf("expected dout %v but got %v", a.d, dr)
			}
			if er.Copy() != split16(a.e) {
				t.Errorf("expected e %v but got %v", a.e, er)
			}
			if fr.Copy() != split16(a.f) {
				t.Errorf("expected f %v but got %v", a.f, fr)
			}
			if gr.Copy() != split16(a.g) {
				t.Errorf("expected g %v but got %v", a.g, gr)
			}
			if hr.Copy() != split16(a.h) {
				t.Errorf("expected h %v but got %v", a.h, hr)
			}
		})
	}
}

func TestDFF(t *testing.T) {
	dff := &DFF{}
	dff.In = Active
	if bit := dff.Out(Inactive); bit == Active {
		t.Errorf("expected dff to be unset before tick")
	}
	if bit := dff.Out(Active); bit == Inactive {
		t.Errorf("expected dff to be set after tick")
	}
	dff.In = Inactive
	if bit := dff.Out(Inactive); bit == Inactive {
		t.Errorf("expected dff to be set before tick")
	}
	if bit := dff.Out(Active); bit == Active {
		t.Errorf("expected dff to be unset after tick")
	}
	if bit := dff.Out(Active); bit == Active {
		t.Errorf("expected dff to be unset after tick")
	}
	if bit := dff.Out(Active); bit == Active {
		t.Errorf("expected dff to be unset after tick")
	}
}
