package chip

import (
	"fmt"
	"testing"
)

func TestHalfAdder(t *testing.T) {
	var assertions = []struct {
		a     Signal
		b     Signal
		carry Signal
		sum   Signal
	}{
		{
			a:     Inactive,
			b:     Inactive,
			carry: Inactive,
			sum:   Inactive,
		},
		{
			a:     Inactive,
			b:     Active,
			carry: Inactive,
			sum:   Active,
		},
		{
			a:     Active,
			b:     Inactive,
			carry: Inactive,
			sum:   Active,
		},

		{
			a:     Active,
			b:     Active,
			carry: Active,
			sum:   Inactive,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given a: %v and b: %v", a.a, a.b), func(t *testing.T) {
			carry, sum := HalfAdder(a.a, a.b)
			if carry != a.carry {
				t.Errorf("expected carry %v but got %v", a.carry, carry)
			}
			if sum != a.sum {
				t.Errorf("expected sum %v but got %v", a.sum, sum)
			}
		})
	}
}

func TestFullAdder(t *testing.T) {
	var assertions = []struct {
		a     Signal
		b     Signal
		c     Signal
		sum   Signal
		carry Signal
	}{
		{
			a:     Inactive,
			b:     Inactive,
			c:     Inactive,
			sum:   Inactive,
			carry: Inactive,
		},
		{
			a:     Inactive,
			b:     Inactive,
			c:     Active,
			sum:   Active,
			carry: Inactive,
		},
		{
			a:     Inactive,
			b:     Active,
			c:     Inactive,
			sum:   Active,
			carry: Inactive,
		},
		{
			a:     Active,
			b:     Inactive,
			c:     Inactive,
			sum:   Active,
			carry: Inactive,
		},
		{
			a:     Inactive,
			b:     Active,
			c:     Active,
			sum:   Inactive,
			carry: Active,
		},
		{
			a:     Active,
			b:     Active,
			c:     Inactive,
			sum:   Inactive,
			carry: Active,
		},
		{
			a:     Active,
			b:     Inactive,
			c:     Active,
			sum:   Inactive,
			carry: Active,
		},
		{
			a:     Active,
			b:     Active,
			c:     Active,
			sum:   Active,
			carry: Active,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given a: %v, b: %v, c: %v", a.a, a.b, a.c), func(t *testing.T) {
			carry, sum := FullAdder(a.a, a.b, a.c)
			if sum != a.sum {
				t.Errorf("expected sum %v but got %v", a.sum, sum)
			}
			if carry != a.carry {
				t.Errorf("expected carry %v but got %v", a.carry, carry)
			}
		})
	}
}

func TestAdder16(t *testing.T) {
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
			a: 256,
			b: 0,
			r: 256,
		},
		{
			a: 0,
			b: 256,
			r: 256,
		},
		{
			a: 256,
			b: 256,
			r: 512,
		},
		{
			a: 10000,
			b: 10000,
			r: 20000,
		},
		{
			a: 65535,
			b: 2,
			r: 1,
		},
	}
	for _, a := range assertions {
		opa := split16(a.a)
		opb := split16(a.b)
		r := Adder16(Wrap(&opa), Wrap(&opb))
		if r.Copy() != split16(a.r) {
			t.Errorf("expected %v but got %v", a.r, r)
		}
	}
}
