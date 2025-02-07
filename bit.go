package snack

import "fmt"

type Bit struct {
	n uint8
}

// BitSplit16 transforms a 16-bit integer to its bit representation in the Bit abstraction. The index of the slice is
// the position of the bit, meaning that in the 0th index you will find the 2^0 position bit, and on the 15th index you
// will find the 2^15 positioned bit.
func BitSplit16(n uint16) []Bit {
	res := make([]Bit, 16)
	for i := range 16 {
		res[i] = NewBit(uint8(n>>(15-i)) & 1)
	}
	return res
}

// BitJoin16 transforms a Bit slice of length 16 to a 16-bit integer
func BitJoin16(n []Bit) uint16 {
	if len(n) != 16 {
		panic(fmt.Errorf("invalid bit join length %d", len(n)))
	}
	res := uint16(0)
	for i := range 16 {
		res = res | (uint16(n[i].Bin()) << (15 - i))
	}
	return res
}

// Expand16 takes a single bit and expands its value to cover 16 bits. That is, if the bit value is 0 then a 16-bit
// unsigned integer containing all zeroes is returned. If the input bit value is 1 then an unsigned 16-bit integer
// containing all ones is returned.
func Expand16(n Bit) uint16 {
	if n.IsSet() {
		return 0xFFFF
	}
	return 0
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
