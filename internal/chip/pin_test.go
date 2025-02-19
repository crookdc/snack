package chip

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("given unset value", func(t *testing.T) {
		bit := NewPin(0)
		if bit.Active() {
			t.Error("got set signal")
		}
	})

	t.Run("given set value", func(t *testing.T) {
		bit := NewPin(1)
		if !bit.Active() {
			t.Error("got unset signal")
		}
	})
}

func TestBitSplit16(t *testing.T) {
	var assertions = []struct {
		n uint16
		r []Signal
	}{
		{
			n: 0b0000_0000_0000_0000,
			r: []Signal{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
		{
			n: 0b0000_1001_0010_0001,
			r: []Signal{
				0, 0, 0, 0, 1, 0, 0, 1,
				0, 0, 1, 0, 0, 0, 0, 1,
			},
		},
		{
			n: 0b1010_1111_0010_1101,
			r: []Signal{
				1, 0, 1, 0, 1, 1, 1, 1,
				0, 0, 1, 0, 1, 1, 0, 1,
			},
		},
		{
			n: 0b1111_1111_1111_1111,
			r: []Signal{
				1, 1, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 1,
			},
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given %b", a.n), func(t *testing.T) {
			r := split16(a.n)
			for i, n := range r {
				if n != a.r[i] {
					t.Errorf("expected %v on index %v but got %v", a.r[i], i, n)
				}
			}
		})
	}
}

func TestExpand16(t *testing.T) {
	var assertions = []struct {
		n Signal
		r uint16
	}{
		{
			n: 1,
			r: 0xFFFF,
		},
		{
			n: 0,
			r: 0,
		},
	}
	for _, a := range assertions {
		t.Run(fmt.Sprintf("given n %v", a.n), func(t *testing.T) {
			r := expand16(a.n)
			if r != split16(a.r) {
				t.Errorf("expected %v but got %v", a.r, r)
			}
		})
	}
}
