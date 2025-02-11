package mem

import (
	"github.com/crookdc/snack/internal/pin"
	"testing"
)

func TestBit_Out(t *testing.T) {
	bit := Bit{
		In: pin.New(pin.Inactive),
	}
	bit.In.Activate()
	if b := bit.Out(pin.New(pin.Inactive)); b == pin.Active {
		t.Errorf("expected inactive pin but got active")
	}
	if b := bit.Out(pin.New(pin.Active)); b == pin.Inactive {
		t.Errorf("expected active pin but got inactive")
	}
	if b := bit.Out(pin.New(pin.Inactive)); b == pin.Inactive {
		t.Errorf("expected active pin but got inactive")
	}
	if b := bit.Out(pin.New(pin.Active)); b == pin.Inactive {
		t.Errorf("expected active pin but got inactive")
	}
	bit.In.Deactivate()
	if b := bit.Out(pin.New(pin.Inactive)); b == pin.Inactive {
		t.Errorf("expected active pin but got inactive")
	}
	if b := bit.Out(pin.New(pin.Active)); b == pin.Active {
		t.Errorf("expected inactive pin but got active")
	}
}

func TestRegister(t *testing.T) {
	reg := new(Register)
	if out := reg.Out(pin.New(pin.Inactive)); out != 0 {
		t.Errorf("expected 0 but got %v", out)
	}
	reg.Set(65234)
	// The register should still yield the initialized value since the clock is inactive
	if out := reg.Out(pin.New(pin.Inactive)); out != 0 {
		t.Errorf("expected 0 but got %v", out)
	}
	// Once the clock becomes active we should be receiving the newly set value
	if out := reg.Out(pin.New(pin.Active)); out != 65234 {
		t.Errorf("expected 65234 but got %v", out)
	}
	// Subsequent outputs should remain the same regardless of the clocks value
	if out := reg.Out(pin.New(pin.Inactive)); out != 65234 {
		t.Errorf("expected 65234 but got %v", out)
	}
	if out := reg.Out(pin.New(pin.Active)); out != 65234 {
		t.Errorf("expected 65234 but got %v", out)
	}
	reg.Set(40923)
	if out := reg.Out(pin.New(pin.Inactive)); out != 65234 {
		t.Errorf("expected 65234 but got %v", out)
	}
	if out := reg.Out(pin.New(pin.Active)); out != 40923 {
		t.Errorf("expected 40923 but got %v", out)
	}
}
