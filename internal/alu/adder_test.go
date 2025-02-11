package alu

import (
	"fmt"
	"github.com/crookdc/snack"
	"testing"
)

func TestHalfAdder(t *testing.T) {
	type assertion struct {
		a     snack.Signal
		b     snack.Signal
		carry snack.Signal
		sum   snack.Signal
	}
	var assertions = []assertion{
		{
			a:     snack.InactiveSignal(),
			b:     snack.InactiveSignal(),
			carry: snack.InactiveSignal(),
			sum:   snack.InactiveSignal(),
		},
		{
			a:     snack.InactiveSignal(),
			b:     snack.ActiveSignal(),
			carry: snack.InactiveSignal(),
			sum:   snack.ActiveSignal(),
		},
		{
			a:     snack.ActiveSignal(),
			b:     snack.InactiveSignal(),
			carry: snack.InactiveSignal(),
			sum:   snack.ActiveSignal(),
		},

		{
			a:     snack.ActiveSignal(),
			b:     snack.ActiveSignal(),
			carry: snack.ActiveSignal(),
			sum:   snack.InactiveSignal(),
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
	type assertion struct {
		a     snack.Signal
		b     snack.Signal
		c     snack.Signal
		sum   snack.Signal
		carry snack.Signal
	}
	var asssertions = []assertion{
		{
			a:     snack.InactiveSignal(),
			b:     snack.InactiveSignal(),
			c:     snack.InactiveSignal(),
			sum:   snack.InactiveSignal(),
			carry: snack.InactiveSignal(),
		},
		{
			a:     snack.InactiveSignal(),
			b:     snack.InactiveSignal(),
			c:     snack.ActiveSignal(),
			sum:   snack.ActiveSignal(),
			carry: snack.InactiveSignal(),
		},
		{
			a:     snack.InactiveSignal(),
			b:     snack.ActiveSignal(),
			c:     snack.InactiveSignal(),
			sum:   snack.ActiveSignal(),
			carry: snack.InactiveSignal(),
		},
		{
			a:     snack.ActiveSignal(),
			b:     snack.InactiveSignal(),
			c:     snack.InactiveSignal(),
			sum:   snack.ActiveSignal(),
			carry: snack.InactiveSignal(),
		},
		{
			a:     snack.InactiveSignal(),
			b:     snack.ActiveSignal(),
			c:     snack.ActiveSignal(),
			sum:   snack.InactiveSignal(),
			carry: snack.ActiveSignal(),
		},
		{
			a:     snack.ActiveSignal(),
			b:     snack.ActiveSignal(),
			c:     snack.InactiveSignal(),
			sum:   snack.InactiveSignal(),
			carry: snack.ActiveSignal(),
		},
		{
			a:     snack.ActiveSignal(),
			b:     snack.InactiveSignal(),
			c:     snack.ActiveSignal(),
			sum:   snack.InactiveSignal(),
			carry: snack.ActiveSignal(),
		},
		{
			a:     snack.ActiveSignal(),
			b:     snack.ActiveSignal(),
			c:     snack.ActiveSignal(),
			sum:   snack.ActiveSignal(),
			carry: snack.ActiveSignal(),
		},
	}
	for _, a := range asssertions {
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
		r := Adder16(a.a, a.b)
		if r != a.r {
			t.Errorf("expected %v but got %v", a.r, r)
		}
	}
}
