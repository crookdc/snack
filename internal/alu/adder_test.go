package alu

import (
	"fmt"
	"github.com/crookdc/snack"
	"testing"
)

func TestHalfAdder(t *testing.T) {
	type assertion struct {
		a     snack.Bit
		b     snack.Bit
		carry snack.Bit
		sum   snack.Bit
	}
	var assertions = []assertion{
		{
			a:     snack.UnsetBit(),
			b:     snack.UnsetBit(),
			carry: snack.UnsetBit(),
			sum:   snack.UnsetBit(),
		},
		{
			a:     snack.UnsetBit(),
			b:     snack.SetBit(),
			carry: snack.UnsetBit(),
			sum:   snack.SetBit(),
		},
		{
			a:     snack.SetBit(),
			b:     snack.UnsetBit(),
			carry: snack.UnsetBit(),
			sum:   snack.SetBit(),
		},

		{
			a:     snack.SetBit(),
			b:     snack.SetBit(),
			carry: snack.SetBit(),
			sum:   snack.UnsetBit(),
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
		a     snack.Bit
		b     snack.Bit
		c     snack.Bit
		sum   snack.Bit
		carry snack.Bit
	}
	var asssertions = []assertion{
		{
			a:     snack.UnsetBit(),
			b:     snack.UnsetBit(),
			c:     snack.UnsetBit(),
			sum:   snack.UnsetBit(),
			carry: snack.UnsetBit(),
		},
		{
			a:     snack.UnsetBit(),
			b:     snack.UnsetBit(),
			c:     snack.SetBit(),
			sum:   snack.SetBit(),
			carry: snack.UnsetBit(),
		},
		{
			a:     snack.UnsetBit(),
			b:     snack.SetBit(),
			c:     snack.UnsetBit(),
			sum:   snack.SetBit(),
			carry: snack.UnsetBit(),
		},
		{
			a:     snack.SetBit(),
			b:     snack.UnsetBit(),
			c:     snack.UnsetBit(),
			sum:   snack.SetBit(),
			carry: snack.UnsetBit(),
		},
		{
			a:     snack.UnsetBit(),
			b:     snack.SetBit(),
			c:     snack.SetBit(),
			sum:   snack.UnsetBit(),
			carry: snack.SetBit(),
		},
		{
			a:     snack.SetBit(),
			b:     snack.SetBit(),
			c:     snack.UnsetBit(),
			sum:   snack.UnsetBit(),
			carry: snack.SetBit(),
		},
		{
			a:     snack.SetBit(),
			b:     snack.UnsetBit(),
			c:     snack.SetBit(),
			sum:   snack.UnsetBit(),
			carry: snack.SetBit(),
		},
		{
			a:     snack.SetBit(),
			b:     snack.SetBit(),
			c:     snack.SetBit(),
			sum:   snack.SetBit(),
			carry: snack.SetBit(),
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
