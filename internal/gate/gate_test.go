package gate

import (
	"fmt"
	"testing"
)

func TestNotAnd(t *testing.T) {
	type assertion struct {
		a uint8
		b uint8
		r uint8
	}
	var assertions = []assertion{
		{
			a: 0x00,
			b: 0x00,
			r: 0xFF,
		},
		{
			a: 0xF0,
			b: 0x0F,
			r: 0x00,
		},
		{
			a: 0x0F,
			b: 0xF0,
			r: 0x00,
		},
		{
			a: 0xFF,
			b: 0xFF,
			r: 0x00,
		},
		{
			a: 0xFF,
			b: 0x00,
			r: 0x00,
		},
		{
			a: 0x00,
			b: 0xFF,
			r: 0x00,
		},
		{
			a: 0b10010001,
			b: 0x00,
			r: 0b01101110,
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
			r := NotAndUint16(a.a, a.b)
			if r != a.r {
				t.Errorf("expected %v with a: %v and b: %v but got %v", a.r, a.a, a.b, r)
			}
		})
	}
}

func TestNot(t *testing.T) {
	type assertion struct {
		a uint8
		r uint8
	}
	var assertions = []assertion{
		{
			a: 0x00,
			r: 0xFF,
		},
		{
			a: 0xF0,
			r: 0x0F,
		},
		{
			a: 0x0F,
			r: 0xF0,
		},
		{
			a: 0xFF,
			r: 0x00,
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
			r := NotUint16(a.a)
			if r != a.r {
				t.Errorf("expected %v with a: %v but got %v", a.r, a.a, r)
			}
		})
	}
}

func TestAnd(t *testing.T) {
	type assertion struct {
		a uint8
		b uint8
		r uint8
	}
	var assertions = []assertion{
		{
			a: 0x00,
			b: 0x00,
			r: 0x00,
		},
		{
			a: 0xF0,
			b: 0x0F,
			r: 0x00,
		},
		{
			a: 0x0F,
			b: 0xF0,
			r: 0x00,
		},
		{
			a: 0xFF,
			b: 0xFF,
			r: 0xFF,
		},
		{
			a: 0xFF,
			b: 0x00,
			r: 0x00,
		},
		{
			a: 0x00,
			b: 0xFF,
			r: 0x00,
		},
		{
			a: 0b10010001,
			b: 0b11000001,
			r: 0b10000001,
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
			r := AndUint16(a.a, a.b)
			if a.r != r {
				t.Errorf("expected %v with a: %v and b: %v but got %v", r, a.r, a.a, a.b)
			}
		})
	}
}

func TestOr(t *testing.T) {
	type assertion struct {
		a uint8
		b uint8
		r uint8
	}
	var assertions = []assertion{
		{
			a: 0x00,
			b: 0x00,
			r: 0x00,
		},
		{
			a: 0xF0,
			b: 0x0F,
			r: 0xFF,
		},
		{
			a: 0x0F,
			b: 0xF0,
			r: 0xFF,
		},
		{
			a: 0xFF,
			b: 0xFF,
			r: 0xFF,
		},
		{
			a: 0xFF,
			b: 0x00,
			r: 0xFF,
		},
		{
			a: 0x00,
			b: 0xFF,
			r: 0xFF,
		},
		{
			a: 0b10010001,
			b: 0b11000001,
			r: 0b11010001,
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
			r := OrUint16(a.a, a.b)
			if a.r != r {
				t.Errorf("expected %v with a: %v and b: %v but got %v", a.r, a.a, a.b, r)
			}
		})
	}
}

func TestXor(t *testing.T) {
	type assertion struct {
		a uint8
		b uint8
		r uint8
	}
	var assertions = []assertion{
		{
			a: 0x00,
			b: 0x00,
			r: 0x00,
		},
		{
			a: 0xF0,
			b: 0x0F,
			r: 0xFF,
		},
		{
			a: 0x0F,
			b: 0xF0,
			r: 0xFF,
		},
		{
			a: 0xFF,
			b: 0x00,
			r: 0xFF,
		},
		{
			a: 0x00,
			b: 0xFF,
			r: 0xFF,
		},
		{
			a: 0xFF,
			b: 0xFF,
			r: 0x00,
		},
		{
			a: 0b10010001,
			b: 0b11000001,
			r: 0b01010000,
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
			r := XorUint16(a.a, a.b)
			if a.r != r {
				t.Errorf("expected %v with a: %v and b: %v but got %v", a.r, a.a, a.b, r)
			}
		})
	}
}

func TestMux2Way(t *testing.T) {
	type assertion struct {
		a   uint16
		b   uint16
		sel uint8
		r   uint16
	}
	var assertions = []assertion{
		{
			a:   55,
			b:   96,
			sel: 0,
			r:   55,
		},
		{
			a:   0,
			b:   53,
			sel: 0,
			r:   0,
		},
		{
			a:   255,
			b:   12,
			sel: 0,
			r:   255,
		},
		{
			a:   0,
			b:   0,
			sel: 0,
			r:   0,
		},
		{
			a:   12,
			b:   0,
			sel: 1,
			r:   0,
		},
		{
			a:   123,
			b:   99,
			sel: 1,
			r:   99,
		},
		{
			a:   0,
			b:   123,
			sel: 1,
			r:   123,
		},
		{
			a:   0,
			b:   0,
			sel: 1,
			r:   0,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given a is %v, b is %v and sel is %v", a.a, a.b, a.sel), func(t *testing.T) {
			r := Mux2Way(a.sel, a.a, a.b)
			if a.r != r {
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
		s [2]uint8
		r uint16
	}
	var assertions = []assertion{
		{
			a: 10,
			b: 20,
			c: 30,
			d: 40,
			s: [2]uint8{0, 0},
			r: 10,
		},
		{
			a: 10,
			b: 20,
			c: 30,
			d: 40,
			s: [2]uint8{0, 1},
			r: 20,
		},
		{
			a: 10,
			b: 20,
			c: 30,
			d: 40,
			s: [2]uint8{1, 0},
			r: 30,
		},
		{
			a: 10,
			b: 20,
			c: 30,
			d: 40,
			s: [2]uint8{1, 1},
			r: 40,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given s: %v", a.s), func(t *testing.T) {
			r := Mux4Way(a.s, a.a, a.b, a.c, a.d)
			if a.r != r {
				t.Errorf("expected %v given s: %v but got %v", a.r, a.s, r)
			}
		})
	}
}

func TestMux8Way(t *testing.T) {
	type assertion struct {
		a uint16
		b uint16
		c uint16
		d uint16
		e uint16
		f uint16
		g uint16
		h uint16

		s [3]uint8
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

			s: [3]uint8{0, 0, 0},
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

			s: [3]uint8{0, 0, 1},
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

			s: [3]uint8{0, 1, 0},
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

			s: [3]uint8{0, 1, 1},
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

			s: [3]uint8{1, 0, 0},
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

			s: [3]uint8{1, 0, 1},
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

			s: [3]uint8{1, 1, 0},
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

			s: [3]uint8{1, 1, 1},
			r: 80,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("with s: %v", a.s), func(t *testing.T) {
			r := Mux8Way(a.s, a.a, a.b, a.c, a.d, a.e, a.f, a.g, a.h)
			if a.r != r {
				t.Errorf("expected %v with s: %v got got %v", a.r, a.s, r)
			}
		})
	}
}
