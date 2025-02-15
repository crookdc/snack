package gate

import (
	"fmt"
	"github.com/crookdc/snack/internal/pin"
	"testing"
)

func TestNotAnd(t *testing.T) {
	var assertions = []struct {
		a pin.Signal
		b pin.Signal
		r pin.Signal
	}{
		{
			a: pin.Inactive,
			b: pin.Inactive,
			r: pin.Active,
		},
		{
			a: pin.Inactive,
			b: pin.Active,
			r: pin.Inactive,
		},
		{
			a: pin.Active,
			b: pin.Inactive,
			r: pin.Inactive,
		},
		{
			a: pin.Active,
			b: pin.Active,
			r: pin.Inactive,
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
			r: 0x0000,
		},
		{
			a: 0x0F00,
			b: 0xF000,
			r: 0x00FF,
		},
		{
			a: 0xFFFF,
			b: 0xFFFF,
			r: 0x0000,
		},
		{
			a: 0xFFFF,
			b: 0x0000,
			r: 0x0000,
		},
		{
			a: 0x0000,
			b: 0xFFFF,
			r: 0x0000,
		},
		{
			a: 0b1001_0001_0011_1111,
			b: 0b0000_0000_0000_0000,
			r: 0b0110_1110_1100_0000,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given a is %v and b is %v", a.a, a.b), func(t *testing.T) {
			r := NotAnd16(pin.Split16(a.a), pin.Split16(a.b))
			if r != pin.Split16(a.r) {
				t.Errorf("expected %v with a: %v and b: %v but got %v", a.r, a.a, a.b, r)
			}
		})
	}
}

func TestNot(t *testing.T) {
	var assertions = []struct {
		a pin.Signal
		r pin.Signal
	}{
		{
			a: pin.Inactive,
			r: pin.Active,
		},
		{
			a: pin.Active,
			r: pin.Inactive,
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
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given a is %v", a.a), func(t *testing.T) {
			r := Not16(pin.Split16(a.a))
			if r != pin.Split16(a.r) {
				t.Errorf("expected %v with a: %v but got %v", a.r, a.a, r)
			}
		})
	}
}

func TestAnd(t *testing.T) {
	var assertions = []struct {
		a pin.Signal
		b pin.Signal
		r pin.Signal
	}{
		{
			a: pin.Inactive,
			b: pin.Inactive,
			r: pin.Inactive,
		},
		{
			a: pin.Inactive,
			b: pin.Active,
			r: pin.Inactive,
		},
		{
			a: pin.Active,
			b: pin.Inactive,
			r: pin.Inactive,
		},
		{
			a: pin.Active,
			b: pin.Active,
			r: pin.Active,
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
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given a is %v and b is %v", a.a, a.b), func(t *testing.T) {
			r := And16(pin.Split16(a.a), pin.Split16(a.b))
			if r != pin.Split16(a.r) {
				t.Errorf("expected %v with a: %v and b: %v but got %v", r, a.r, a.a, a.b)
			}
		})
	}
}

func TestOr(t *testing.T) {
	var assertions = []struct {
		a pin.Signal
		b pin.Signal
		r pin.Signal
	}{
		{
			a: pin.Inactive,
			b: pin.Inactive,
			r: pin.Inactive,
		},
		{
			a: pin.Inactive,
			b: pin.Active,
			r: pin.Active,
		},
		{
			a: pin.Active,
			b: pin.Inactive,
			r: pin.Active,
		},
		{
			a: pin.Active,
			b: pin.Active,
			r: pin.Active,
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
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given a is %v and b is %v", a.a, a.b), func(t *testing.T) {
			r := Or16(pin.Split16(a.a), pin.Split16(a.b))
			if r != pin.Split16(a.r) {
				t.Errorf("expected %v with a: %v and b: %v but got %v", a.r, a.a, a.b, r)
			}
		})
	}
}

func TestXor(t *testing.T) {
	var assertions = []struct {
		a pin.Signal
		b pin.Signal
		r pin.Signal
	}{
		{
			a: pin.Inactive,
			b: pin.Inactive,
			r: pin.Inactive,
		},
		{
			a: pin.Inactive,
			b: pin.Active,
			r: pin.Active,
		},
		{
			a: pin.Active,
			b: pin.Inactive,
			r: pin.Active,
		},
		{
			a: pin.Active,
			b: pin.Active,
			r: pin.Inactive,
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
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given a is %v and b is %v", a.a, a.b), func(t *testing.T) {
			r := Xor16(pin.Split16(a.a), pin.Split16(a.b))
			if r != pin.Split16(a.r) {
				t.Errorf("expected %v with a: %v and b: %v but got %v", a.r, a.a, a.b, r)
			}
		})
	}
}

func TestMux2Way16(t *testing.T) {
	var assertions = []struct {
		a   uint16
		b   uint16
		sel pin.Signal
		r   uint16
	}{
		{
			a:   55,
			b:   96,
			sel: pin.Inactive,
			r:   55,
		},
		{
			a:   0,
			b:   53,
			sel: pin.Inactive,
			r:   0,
		},
		{
			a:   255,
			b:   12,
			sel: pin.Inactive,
			r:   255,
		},
		{
			a:   0,
			b:   0,
			sel: pin.Inactive,
			r:   0,
		},
		{
			a:   12,
			b:   0,
			sel: pin.Active,
			r:   0,
		},
		{
			a:   123,
			b:   99,
			sel: pin.Active,
			r:   99,
		},
		{
			a:   0,
			b:   123,
			sel: pin.Active,
			r:   123,
		},
		{
			a:   0,
			b:   0,
			sel: pin.Active,
			r:   0,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given a is %v, b is %v and sel is %v", a.a, a.b, a.sel), func(t *testing.T) {
			r := Mux2Way16(a.sel, pin.Split16(a.a), pin.Split16(a.b))
			if r != pin.Split16(a.r) {
				t.Errorf("expected %v with a: %v, b: %v and sel: %v but got %v", a.r, a.a, a.b, a.sel, r)
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
		s [2]pin.Signal
		r uint16
	}
	var assertions = []assertion{
		{
			a: 10,
			b: 20,
			c: 30,
			d: 40,
			s: [2]pin.Signal{pin.Inactive, pin.Inactive},
			r: 10,
		},
		{
			a: 10,
			b: 20,
			c: 30,
			d: 40,
			s: [2]pin.Signal{pin.Inactive, pin.Active},
			r: 20,
		},
		{
			a: 10,
			b: 20,
			c: 30,
			d: 40,
			s: [2]pin.Signal{pin.Active, pin.Inactive},
			r: 30,
		},
		{
			a: 10,
			b: 20,
			c: 30,
			d: 40,
			s: [2]pin.Signal{pin.Active, pin.Active},
			r: 40,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given s: %v", a.s), func(t *testing.T) {
			r := Mux4Way16(a.s, pin.Split16(a.a), pin.Split16(a.b), pin.Split16(a.c), pin.Split16(a.d))
			if r != pin.Split16(a.r) {
				t.Errorf("expected %v given s: %v but got %v", a.r, a.s, r)
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

		s [3]pin.Signal
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

			s: [3]pin.Signal{pin.Inactive, pin.Inactive, pin.Inactive},
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

			s: [3]pin.Signal{pin.Inactive, pin.Inactive, pin.Active},
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

			s: [3]pin.Signal{pin.Inactive, pin.Active, pin.Inactive},
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

			s: [3]pin.Signal{pin.Inactive, pin.Active, pin.Active},
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

			s: [3]pin.Signal{pin.Active, pin.Inactive, pin.Inactive},
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

			s: [3]pin.Signal{pin.Active, pin.Inactive, pin.Active},
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

			s: [3]pin.Signal{pin.Active, pin.Active, pin.Inactive},
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

			s: [3]pin.Signal{pin.Active, pin.Active, pin.Active},
			r: 80,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("with s: %v", a.s), func(t *testing.T) {
			r := Mux8Way16(a.s, pin.Split16(a.a), pin.Split16(a.b), pin.Split16(a.c), pin.Split16(a.d), pin.Split16(a.e), pin.Split16(a.f), pin.Split16(a.g), pin.Split16(a.h))
			if r != pin.Split16(a.r) {
				t.Errorf("expected %v with s: %v got got %v", a.r, a.s, r)
			}
		})
	}
}

func TestDemux2Way(t *testing.T) {
	type assertion struct {
		in uint16
		s  pin.Signal
		a  uint16
		b  uint16
	}
	var assertions = []assertion{
		{
			in: 0,
			s:  pin.Inactive,
			a:  0,
			b:  0,
		},
		{
			in: 65_535,
			s:  pin.Inactive,
			a:  65_535,
			b:  0,
		},
		{
			in: 256,
			s:  pin.Active,
			a:  0,
			b:  256,
		},
		{
			in: 0,
			s:  pin.Active,
			a:  0,
			b:  0,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given in is %v and s is %v", a.in, a.s), func(t *testing.T) {
			ar, br := DMux2Way16(a.s, pin.Split16(a.in))
			if ar != pin.Split16(a.a) {
				t.Errorf("expected ar %v but got %v", a.a, ar)
			}
			if br != pin.Split16(a.b) {
				t.Errorf("expected br %v but got %v", a.b, br)
			}
		})
	}
}

func TestDemux4Way(t *testing.T) {
	type assertion struct {
		in uint16
		s  [2]pin.Signal
		a  uint16
		b  uint16
		c  uint16
		d  uint16
	}
	var assertions = []assertion{
		{
			in: 65_313,
			s:  [2]pin.Signal{pin.Inactive, pin.Inactive},
			a:  65_313,
			b:  0,
			c:  0,
			d:  0,
		},
		{
			in: 23_230,
			s:  [2]pin.Signal{pin.Inactive, pin.Active},
			a:  0,
			b:  23_230,
			c:  0,
			d:  0,
		},
		{
			in: 9012,
			s:  [2]pin.Signal{pin.Active, pin.Inactive},
			a:  0,
			b:  0,
			c:  9012,
			d:  0,
		},
		{
			in: 1234,
			s:  [2]pin.Signal{pin.Active, pin.Active},
			a:  0,
			b:  0,
			c:  0,
			d:  1234,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given in %v and s %v", a.in, a.s), func(t *testing.T) {
			ar, br, cr, dr := DMux4Way16(a.s, pin.Split16(a.in))
			if ar != pin.Split16(a.a) {
				t.Errorf("expected a %v but got %v", a.a, ar)
			}
			if br != pin.Split16(a.b) {
				t.Errorf("expected b %v but got %v", a.b, br)
			}
			if cr != pin.Split16(a.c) {
				t.Errorf("expected c %v but got %v", a.c, cr)
			}
			if dr != pin.Split16(a.d) {
				t.Errorf("expected d %v but got %v", a.d, dr)
			}
		})
	}
}

func TestDemux8Way(t *testing.T) {
	type assertion struct {
		in uint16
		s  [3]pin.Signal
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
			s:  [3]pin.Signal{pin.Inactive, pin.Inactive, pin.Inactive},
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
			s:  [3]pin.Signal{pin.Inactive, pin.Inactive, pin.Active},
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
			s:  [3]pin.Signal{pin.Inactive, pin.Active, pin.Inactive},
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
			s:  [3]pin.Signal{pin.Inactive, pin.Active, pin.Active},
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
			s:  [3]pin.Signal{pin.Active, pin.Inactive, pin.Inactive},
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
			s:  [3]pin.Signal{pin.Active, pin.Inactive, pin.Active},
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
			s:  [3]pin.Signal{pin.Active, pin.Active, pin.Inactive},
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
			s:  [3]pin.Signal{pin.Active, pin.Active, pin.Active},
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
			ar, br, cr, dr, er, fr, gr, hr := DMux8Way16(a.s, pin.Split16(a.in))
			if ar != pin.Split16(a.a) {
				t.Errorf("expected a %v but got %v", a.a, ar)
			}
			if br != pin.Split16(a.b) {
				t.Errorf("expected b %v but got %v", a.b, br)
			}
			if cr != pin.Split16(a.c) {
				t.Errorf("expected c %v but got %v", a.c, cr)
			}
			if dr != pin.Split16(a.d) {
				t.Errorf("expected d %v but got %v", a.d, dr)
			}
			if er != pin.Split16(a.e) {
				t.Errorf("expected e %v but got %v", a.e, er)
			}
			if fr != pin.Split16(a.f) {
				t.Errorf("expected f %v but got %v", a.f, fr)
			}
			if gr != pin.Split16(a.g) {
				t.Errorf("expected g %v but got %v", a.g, gr)
			}
			if hr != pin.Split16(a.h) {
				t.Errorf("expected h %v but got %v", a.h, hr)
			}
		})
	}
}

func TestDFF(t *testing.T) {
	dff := &DFF{}
	dff.In.Activate()
	if bit := dff.Out(pin.Inactive); bit == pin.Active {
		t.Errorf("expected dff to be unset before tick")
	}
	if bit := dff.Out(pin.Active); bit == pin.Inactive {
		t.Errorf("expected dff to be set after tick")
	}
	dff.In.Deactivate()
	if bit := dff.Out(pin.Inactive); bit == pin.Inactive {
		t.Errorf("expected dff to be set before tick")
	}
	if bit := dff.Out(pin.Active); bit == pin.Active {
		t.Errorf("expected dff to be unset after tick")
	}
	if bit := dff.Out(pin.Active); bit == pin.Active {
		t.Errorf("expected dff to be unset after tick")
	}
	if bit := dff.Out(pin.Active); bit == pin.Active {
		t.Errorf("expected dff to be unset after tick")
	}
}
