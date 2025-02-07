package alu

import (
	"github.com/crookdc/snack"
	"github.com/crookdc/snack/internal/gate"
)

type ALU struct {
	ZX snack.Bit
	NX snack.Bit
	ZY snack.Bit
	NY snack.Bit
	F  snack.Bit
	NO snack.Bit
}

func (a *ALU) Call(x, y uint16) uint16 {
	x = gate.AndUint16(x, snack.Expand16(gate.NotBit(a.ZX)))
	x = gate.XorUint16(x, snack.Expand16(a.NX))

	y = gate.AndUint16(y, snack.Expand16(gate.NotBit(a.ZY)))
	y = gate.XorUint16(y, snack.Expand16(a.NY))

	out := gate.Mux2Way(a.F.Bin(), gate.AndUint16(x, y), Adder16(x, y))
	return gate.XorUint16(out, snack.Expand16(a.NO))
}
