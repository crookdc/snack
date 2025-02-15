package mem

import (
	"github.com/crookdc/snack/internal/pin"
	"testing"
)

func TestBit_Out(t *testing.T) {
	bit := Bit{}
	if b := bit.Out(pin.New(pin.Inactive), pin.New(pin.Active)); b == pin.Active {
		t.Errorf("expected inactive pin but got active")
	}
	if b := bit.Out(pin.New(pin.Active), pin.New(pin.Active)); b == pin.Inactive {
		t.Errorf("expected active pin but got inactive")
	}
	if b := bit.Out(pin.New(pin.Inactive), pin.New(pin.Active)); b == pin.Inactive {
		t.Errorf("expected active pin but got inactive")
	}
	if b := bit.Out(pin.New(pin.Active), pin.New(pin.Active)); b == pin.Inactive {
		t.Errorf("expected active pin but got inactive")
	}
	if b := bit.Out(pin.New(pin.Inactive), pin.New(pin.Inactive)); b == pin.Inactive {
		t.Errorf("expected active pin but got inactive")
	}
	if b := bit.Out(pin.New(pin.Active), pin.New(pin.Inactive)); b == pin.Active {
		t.Errorf("expected inactive pin but got active")
	}
}

func TestRegister(t *testing.T) {
	reg := new(Register)
	if out := reg.Out(pin.New(pin.Inactive), pin.New16(pin.Split16(65234))); out != pin.Split16(0) {
		t.Errorf("expected 0 but got %v", out)
	}
	// The register should still yield the initialized value since the clock is inactive
	if out := reg.Out(pin.New(pin.Inactive), pin.New16(pin.Split16(65234))); out != pin.Split16(0) {
		t.Errorf("expected 0 but got %v", out)
	}
	// Once the clock becomes active we should be receiving the newly set value
	if out := reg.Out(pin.New(pin.Active), pin.New16(pin.Split16(65234))); out != pin.Split16(65234) {
		t.Errorf("expected 65234 but got %v", out)
	}
	// Subsequent outputs should remain the same regardless of the clocks value
	if out := reg.Out(pin.New(pin.Inactive), pin.New16(pin.Split16(65234))); out != pin.Split16(65234) {
		t.Errorf("expected 65234 but got %v", out)
	}
	if out := reg.Out(pin.New(pin.Active), pin.New16(pin.Split16(65234))); out != pin.Split16(65234) {
		t.Errorf("expected 65234 but got %v", out)
	}
	if out := reg.Out(pin.New(pin.Inactive), pin.New16(pin.Split16(40923))); out != pin.Split16(65234) {
		t.Errorf("expected 65234 but got %v", out)
	}
	if out := reg.Out(pin.New(pin.Active), pin.New16(pin.Split16(40923))); out != pin.Split16(40923) {
		t.Errorf("expected 40923 but got %v", out)
	}
}

func TestRAM8_Out(t *testing.T) {
	equals := func(a [16]pin.Pin, b [16]pin.Signal) bool {
		converted := [16]pin.Signal{}
		for i := range a {
			converted[i] = a[i].Signal()
		}
		return converted == b
	}
	addr := func(a, b, c pin.Signal) [3]pin.Pin {
		return [3]pin.Pin{pin.New(a), pin.New(b), pin.New(c)}
	}

	ram := RAM8{}
	clk := pin.New(pin.Inactive)
	n := ram.Out(clk, addr(0, 0, 1), pin.New16(pin.Split16(65343)))
	if !equals(pin.New16(pin.Split16(0)), n) {
		t.Errorf("expected %v but got %v", pin.New16(pin.Split16(0)), n)
	}
	clk.Activate()
	n = ram.Out(clk, addr(0, 0, 1), pin.New16(pin.Split16(65343)))
	if !equals(pin.New16(pin.Split16(65343)), n) {
		t.Errorf("expected %v but got %v", pin.New16(pin.Split16(65343)), n)
	}
	clk.Deactivate()
	n = ram.Out(clk, addr(0, 0, 0), pin.New16(pin.Split16(12345)))
	if !equals(pin.New16(pin.Split16(0)), n) {
		t.Errorf("expected %v but got %v", pin.New16(pin.Split16(0)), n)
	}
	clk.Activate()
	n = ram.Out(clk, addr(0, 0, 0), pin.New16(pin.Split16(12345)))
	if !equals(pin.New16(pin.Split16(12345)), n) {
		t.Errorf("expected %v but got %v", pin.New16(pin.Split16(12345)), n)
	}
	clk.Deactivate()
	n = ram.Out(clk, addr(0, 0, 1), pin.New16(pin.Split16(65343)))
	if !equals(pin.New16(pin.Split16(65343)), n) {
		t.Errorf("expected %v but got %v", pin.New16(pin.Split16(65343)), n)
	}
}

func TestRAM64_Out(t *testing.T) {
	equals := func(a [16]pin.Pin, b [16]pin.Signal) bool {
		converted := [16]pin.Signal{}
		for i := range a {
			converted[i] = a[i].Signal()
		}
		return converted == b
	}
	addr := func(a, b, c, d, e, f pin.Signal) [6]pin.Pin {
		return [6]pin.Pin{pin.New(a), pin.New(b), pin.New(c), pin.New(d), pin.New(e), pin.New(f)}
	}

	ram := RAM64{}
	clk := pin.New(pin.Inactive)
	n := ram.Out(clk, addr(0, 0, 0, 0, 0, 0), pin.New16(pin.Split16(50456)))
	if !equals(pin.New16(pin.Split16(0)), n) {
		t.Errorf("expected %v but got %v", pin.New16(pin.Split16(0)), n)
	}
	clk.Activate()
	n = ram.Out(clk, addr(0, 0, 0, 0, 0, 0), pin.New16(pin.Split16(50456)))
	if !equals(pin.New16(pin.Split16(50456)), n) {
		t.Errorf("expected %v but got %v", pin.New16(pin.Split16(50456)), n)
	}
	clk.Deactivate()
	n = ram.Out(clk, addr(0, 0, 0, 0, 1, 1), pin.New16(pin.Split16(12345)))
	if !equals(pin.New16(pin.Split16(0)), n) {
		t.Errorf("expected %v but got %v", pin.New16(pin.Split16(0)), n)
	}
	clk.Activate()
	n = ram.Out(clk, addr(0, 0, 0, 0, 1, 1), pin.New16(pin.Split16(12345)))
	if !equals(pin.New16(pin.Split16(12345)), n) {
		t.Errorf("expected %v but got %v", pin.New16(pin.Split16(0)), n)
	}
	clk.Deactivate()
	n = ram.Out(clk, addr(0, 0, 0, 0, 0, 0), pin.New16(pin.Split16(12345)))
	if !equals(pin.New16(pin.Split16(50456)), n) {
		t.Errorf("expected %v but got %v", pin.New16(pin.Split16(0)), n)
	}
	clk.Activate()
	n = ram.Out(clk, addr(0, 0, 1, 0, 1, 0), pin.New16(pin.Split16(54312)))
	if !equals(pin.New16(pin.Split16(54312)), n) {
		t.Errorf("expected %v but got %v", pin.New16(pin.Split16(0)), n)
	}
	clk.Deactivate()
	n = ram.Out(clk, addr(0, 0, 1, 0, 1, 0), pin.New16(pin.Split16(54312)))
	if !equals(pin.New16(pin.Split16(54312)), n) {
		t.Errorf("expected %v but got %v", pin.New16(pin.Split16(0)), n)
	}
	n = ram.Out(clk, addr(0, 0, 0, 0, 1, 1), pin.New16(pin.Split16(12345)))
	if !equals(pin.New16(pin.Split16(12345)), n) {
		t.Errorf("expected %v but got %v", pin.New16(pin.Split16(0)), n)
	}
}
