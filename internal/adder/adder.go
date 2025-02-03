// Package adder provides implementations for adder chips. Included are half adders, full adders and 'just' adders.
package adder

import (
	"github.com/crookdc/snack"
	"github.com/crookdc/snack/internal/gate"
)

// Half is a so-called half adder that accepts two bits as input and produces an output consisting of a sum and a carry.
func Half(a, b snack.Bit) (carry snack.Bit, sum snack.Bit) {
	return gate.AndBit(a, b), gate.XorBit(a, b)
}
