package alu

import (
	"fmt"
	"github.com/crookdc/snack"
	"testing"
)

func TestALU_Call(t *testing.T) {
	type assertion struct {
		x uint16
		y uint16
		r uint16
	}
	t.Run("addition", func(t *testing.T) {
		alu := ALU{
			ZX: snack.UnsetBit(),
			NX: snack.UnsetBit(),
			ZY: snack.UnsetBit(),
			NY: snack.UnsetBit(),
			F:  snack.SetBit(),
			NO: snack.UnsetBit(),
		}
		var assertions = []assertion{
			{
				x: 512,
				y: 512,
				r: 1024,
			},
			{
				x: 256,
				y: 512,
				r: 768,
			},
			{
				x: 255,
				y: 0,
				r: 255,
			},
			{
				x: 0,
				y: 255,
				r: 255,
			},
		}
		for _, a := range assertions {
			t.Run(fmt.Sprintf("given x: %v, y: %v ", a.x, a.y), func(t *testing.T) {
				r := alu.Call(a.x, a.y)
				if r != a.r {
					t.Errorf("expected %v but got %v", a.r, r)
				}
			})
		}
	})
}
