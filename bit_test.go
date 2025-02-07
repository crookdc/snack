package snack

import (
	"fmt"
	"testing"
)

func TestNewBit(t *testing.T) {
	t.Run("given unset value", func(t *testing.T) {
		bit := NewBit(0)
		if bit.IsSet() {
			t.Error("got set bit")
		}
	})

	t.Run("given set value", func(t *testing.T) {
		bit := NewBit(1)
		if !bit.IsSet() {
			t.Error("got unset bit")
		}
	})

	t.Run("given invalid value", func(t *testing.T) {
		defer func() {
			if err := recover(); err == nil {
				t.Errorf("expected error but got nil")
			}
		}()
		NewBit(2)
	})
}

func TestBitSplit16(t *testing.T) {
	type assertion struct {
		n uint16
		r []Bit
	}
	var assertions = []assertion{
		{
			n: 0b0000_0000_0000_0000,
			r: []Bit{
				{0}, {0}, {0}, {0}, {0}, {0}, {0}, {0},
				{0}, {0}, {0}, {0}, {0}, {0}, {0}, {0},
			},
		},
		{
			n: 0b0000_1001_0010_0001,
			r: []Bit{
				{0}, {0}, {0}, {0}, {1}, {0}, {0}, {1},
				{0}, {0}, {1}, {0}, {0}, {0}, {0}, {1},
			},
		},
		{
			n: 0b1010_1111_0010_1101,
			r: []Bit{
				{1}, {0}, {1}, {0}, {1}, {1}, {1}, {1},
				{0}, {0}, {1}, {0}, {1}, {1}, {0}, {1},
			},
		},
		{
			n: 0b1111_1111_1111_1111,
			r: []Bit{
				{1}, {1}, {1}, {1}, {1}, {1}, {1}, {1},
				{1}, {1}, {1}, {1}, {1}, {1}, {1}, {1},
			},
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given %b", a.n), func(t *testing.T) {
			r := BitSplit16(a.n)
			for i, n := range r {
				if n != a.r[i] {
					t.Errorf("expected %v on index %v but got %v", a.r[i], i, n)
				}
			}
		})
	}
}

func TestBitJoin16(t *testing.T) {
	t.Run("given too long slice", func(t *testing.T) {
		defer func() {
			if err := recover(); err == nil {
				t.Errorf("expected panic but none ocurred")
			}
		}()
		BitJoin16(make([]Bit, 17))
	})

	t.Run("given too short slice", func(t *testing.T) {
		defer func() {
			if err := recover(); err == nil {
				t.Errorf("expected panic but none ocurred")
			}
		}()
		BitJoin16(make([]Bit, 15))
	})

	type assertion struct {
		n []Bit
		r uint16
	}
	var assertions = []assertion{
		{
			n: []Bit{
				{0}, {0}, {0}, {0}, {0}, {0}, {0}, {0},
				{0}, {0}, {0}, {0}, {0}, {0}, {0}, {0},
			},
			r: 0b0000_0000_0000_0000,
		},
		{
			n: []Bit{
				{0}, {0}, {0}, {0}, {1}, {0}, {0}, {1},
				{0}, {0}, {1}, {0}, {0}, {0}, {0}, {1},
			},
			r: 0b0000_1001_0010_0001,
		},
		{
			n: []Bit{
				{1}, {0}, {1}, {0}, {1}, {1}, {1}, {1},
				{0}, {0}, {1}, {0}, {1}, {1}, {0}, {1},
			},
			r: 0b1010_1111_0010_1101,
		},
		{
			n: []Bit{
				{1}, {1}, {1}, {1}, {1}, {1}, {1}, {1},
				{1}, {1}, {1}, {1}, {1}, {1}, {1}, {1},
			},
			r: 0b1111_1111_1111_1111,
		},
	}
	for _, a := range assertions {
		r := BitJoin16(a.n)
		if r != a.r {
			t.Errorf("expected %v but got %v", a.r, r)
		}
	}
}

func TestExpand16(t *testing.T) {
	var assertions = []struct {
		n Bit
		r uint16
	}{
		{
			n: SetBit(),
			r: 0xFFFF,
		},
		{
			n: UnsetBit(),
			r: 0,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given n %v", a.n), func(t *testing.T) {
			r := Expand16(a.n)
			if r != a.r {
				t.Errorf("expected %v but got %v", a.r, r)
			}
		})
	}
}
