package alu

import (
	"fmt"
	"github.com/crookdc/snack/internal/pin"
	"testing"
)

func TestHalfAdder(t *testing.T) {
	var assertions = []struct {
		a     pin.Signal
		b     pin.Signal
		carry pin.Signal
		sum   pin.Signal
	}{
		{
			a:     pin.Inactive,
			b:     pin.Inactive,
			carry: pin.Inactive,
			sum:   pin.Inactive,
		},
		{
			a:     pin.Inactive,
			b:     pin.Active,
			carry: pin.Inactive,
			sum:   pin.Active,
		},
		{
			a:     pin.Active,
			b:     pin.Inactive,
			carry: pin.Inactive,
			sum:   pin.Active,
		},

		{
			a:     pin.Active,
			b:     pin.Active,
			carry: pin.Active,
			sum:   pin.Inactive,
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
		a     pin.Signal
		b     pin.Signal
		c     pin.Signal
		sum   pin.Signal
		carry pin.Signal
	}{
		{
			a:     pin.Inactive,
			b:     pin.Inactive,
			c:     pin.Inactive,
			sum:   pin.Inactive,
			carry: pin.Inactive,
		},
		{
			a:     pin.Inactive,
			b:     pin.Inactive,
			c:     pin.Active,
			sum:   pin.Active,
			carry: pin.Inactive,
		},
		{
			a:     pin.Inactive,
			b:     pin.Active,
			c:     pin.Inactive,
			sum:   pin.Active,
			carry: pin.Inactive,
		},
		{
			a:     pin.Active,
			b:     pin.Inactive,
			c:     pin.Inactive,
			sum:   pin.Active,
			carry: pin.Inactive,
		},
		{
			a:     pin.Inactive,
			b:     pin.Active,
			c:     pin.Active,
			sum:   pin.Inactive,
			carry: pin.Active,
		},
		{
			a:     pin.Active,
			b:     pin.Active,
			c:     pin.Inactive,
			sum:   pin.Inactive,
			carry: pin.Active,
		},
		{
			a:     pin.Active,
			b:     pin.Inactive,
			c:     pin.Active,
			sum:   pin.Inactive,
			carry: pin.Active,
		},
		{
			a:     pin.Active,
			b:     pin.Active,
			c:     pin.Active,
			sum:   pin.Active,
			carry: pin.Active,
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
		r := Adder16(pin.Split16(a.a), pin.Split16(a.b))
		if r != pin.Split16(a.r) {
			t.Errorf("expected %v but got %v", a.r, r)
		}
	}
}
