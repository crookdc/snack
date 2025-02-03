package snack

import "fmt"

type Bit struct {
	n uint8
}

func NewBit(n uint8) Bit {
	if n == 0 {
		return UnsetBit()
	} else if n == 1 {
		return SetBit()
	}
	panic(fmt.Errorf("invalid bit value %d", n))
}

func SetBit() Bit {
	return Bit{n: 1}
}

func UnsetBit() Bit {
	return Bit{n: 0}
}

func (b *Bit) Mask() uint8 {
	if b.IsSet() {
		return 0xFF
	}
	return 0
}

func (b *Bit) Set() {
	b.n = 1
}

func (b *Bit) Unset() {
	b.n = 0
}

func (b *Bit) IsSet() bool {
	return b.n == 1
}

func (b *Bit) Bin() uint8 {
	return b.n
}
