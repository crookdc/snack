package mem

import (
	"github.com/crookdc/snack"
	"testing"
)

func TestBit_Out(t *testing.T) {
	bit := Bit{
		In: snack.InactiveSignal(),
	}
	bit.In.Activate()
	if b := bit.Out(snack.InactiveSignal()); b.IsActive() {
		t.Errorf("expected inactive signal but got active")
	}
	if b := bit.Out(snack.ActiveSignal()); !b.IsActive() {
		t.Errorf("expected active signal but got inactive")
	}
	if b := bit.Out(snack.InactiveSignal()); !b.IsActive() {
		t.Errorf("expected active signal but got inactive")
	}
	if b := bit.Out(snack.ActiveSignal()); !b.IsActive() {
		t.Errorf("expected active signal but got inactive")
	}
	bit.In.Deactivate()
	if b := bit.Out(snack.InactiveSignal()); !b.IsActive() {
		t.Errorf("expected active signal but got inactive")
	}
	if b := bit.Out(snack.ActiveSignal()); b.IsActive() {
		t.Errorf("expected inactive signal but got active")
	}
}
