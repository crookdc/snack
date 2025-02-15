package mem

import (
	"fmt"
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
	address := func(n int) [3]pin.Pin {
		return [3]pin.Pin{
			pin.New(pin.Signal(n >> 0 & 1)),
			pin.New(pin.Signal(n >> 1 & 1)),
			pin.New(pin.Signal(n >> 2 & 1)),
		}
	}
	ram := RAM8{}
	clk := pin.New(pin.Inactive)
	for i := range 8 {
		addr := address(i)
		t.Run(fmt.Sprintf("setting address %v", addr), func(t *testing.T) {
			clk.Deactivate()
			n := ram.Out(clk, addr, pin.New16(pin.Split16(uint16(i))))
			if !equals(pin.New16(pin.Split16(0)), n) {
				t.Errorf("expected %v but got %v", 0, n)
			}
			clk.Activate()
			n = ram.Out(clk, addr, pin.New16(pin.Split16(uint16(i))))
			if !equals(pin.New16(pin.Split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Deactivate()
			n = ram.Out(clk, addr, pin.New16(pin.Split16(uint16(i))))
			if !equals(pin.New16(pin.Split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}

	for i := range 8 {
		addr := address(i)
		t.Run(fmt.Sprintf("unsetting address %v", addr), func(t *testing.T) {
			clk.Deactivate()
			n := ram.Out(clk, addr, pin.New16(pin.Split16(uint16(i))))
			if !equals(pin.New16(pin.Split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Activate()
			n = ram.Out(clk, addr, pin.New16(pin.Split16(0)))
			if !equals(pin.New16(pin.Split16(0)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Deactivate()
		})
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
	address := func(n int) [6]pin.Pin {
		n = n >> 3
		return [6]pin.Pin{
			pin.New(pin.Signal(n >> 0 & 1)),
			pin.New(pin.Signal(n >> 1 & 1)),
			pin.New(pin.Signal(n >> 2 & 1)),
		}
	}
	ram := RAM64{}
	clk := pin.New(pin.Inactive)
	for i := 0; i < 64; i += 8 {
		addr := address(i)
		t.Run(fmt.Sprintf("setting address %v", addr), func(t *testing.T) {
			clk.Deactivate()
			n := ram.Out(clk, addr, pin.New16(pin.Split16(uint16(i))))
			if !equals(pin.New16(pin.Split16(0)), n) {
				t.Errorf("expected %v but got %v", 0, n)
			}
			clk.Activate()
			n = ram.Out(clk, addr, pin.New16(pin.Split16(uint16(i))))
			if !equals(pin.New16(pin.Split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Deactivate()
			n = ram.Out(clk, addr, pin.New16(pin.Split16(uint16(i))))
			if !equals(pin.New16(pin.Split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}

	for i := 0; i < 64; i += 8 {
		addr := address(i)
		t.Run(fmt.Sprintf("unsetting address %v", addr), func(t *testing.T) {
			clk.Deactivate()
			n := ram.Out(clk, addr, pin.New16(pin.Split16(uint16(i))))
			if !equals(pin.New16(pin.Split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Activate()
			n = ram.Out(clk, addr, pin.New16(pin.Split16(0)))
			if !equals(pin.New16(pin.Split16(0)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Deactivate()
		})
	}
}

func TestRAM512_Out(t *testing.T) {
	equals := func(a [16]pin.Pin, b [16]pin.Signal) bool {
		converted := [16]pin.Signal{}
		for i := range a {
			converted[i] = a[i].Signal()
		}
		return converted == b
	}
	address := func(n int) [9]pin.Pin {
		n = n >> 6
		return [9]pin.Pin{
			pin.New(pin.Signal(n >> 0 & 1)),
			pin.New(pin.Signal(n >> 1 & 1)),
			pin.New(pin.Signal(n >> 2 & 1)),
		}
	}
	ram := RAM512{}
	clk := pin.New(pin.Inactive)
	for i := 0; i < 512; i += 64 {
		addr := address(i)
		t.Run(fmt.Sprintf("setting address %v", addr), func(t *testing.T) {
			clk.Deactivate()
			n := ram.Out(clk, addr, pin.New16(pin.Split16(uint16(i))))
			if !equals(pin.New16(pin.Split16(0)), n) {
				t.Errorf("expected %v but got %v", 0, n)
			}
			clk.Activate()
			n = ram.Out(clk, addr, pin.New16(pin.Split16(uint16(i))))
			if !equals(pin.New16(pin.Split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Deactivate()
			n = ram.Out(clk, addr, pin.New16(pin.Split16(uint16(i))))
			if !equals(pin.New16(pin.Split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}

	for i := 0; i < 512; i += 64 {
		addr := address(i)
		t.Run(fmt.Sprintf("unsetting address %v", addr), func(t *testing.T) {
			clk.Deactivate()
			n := ram.Out(clk, addr, pin.New16(pin.Split16(uint16(i))))
			if !equals(pin.New16(pin.Split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Activate()
			n = ram.Out(clk, addr, pin.New16(pin.Split16(0)))
			if !equals(pin.New16(pin.Split16(0)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Deactivate()
		})
	}
}

func TestRAM4K_Out(t *testing.T) {
	equals := func(a [16]pin.Pin, b [16]pin.Signal) bool {
		converted := [16]pin.Signal{}
		for i := range a {
			converted[i] = a[i].Signal()
		}
		return converted == b
	}
	address := func(n int) [12]pin.Pin {
		n = n >> 9
		return [12]pin.Pin{
			pin.New(pin.Signal(n >> 0 & 1)),
			pin.New(pin.Signal(n >> 1 & 1)),
			pin.New(pin.Signal(n >> 2 & 1)),
		}
	}
	ram := RAM4K{}
	clk := pin.New(pin.Inactive)
	for i := 0; i < 4096; i += 512 {
		addr := address(i)
		t.Run(fmt.Sprintf("setting address %v", addr), func(t *testing.T) {
			clk.Deactivate()
			n := ram.Out(clk, addr, pin.New16(pin.Split16(uint16(i))))
			if !equals(pin.New16(pin.Split16(0)), n) {
				t.Errorf("expected %v but got %v", 0, n)
			}
			clk.Activate()
			n = ram.Out(clk, addr, pin.New16(pin.Split16(uint16(i))))
			if !equals(pin.New16(pin.Split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Deactivate()
			n = ram.Out(clk, addr, pin.New16(pin.Split16(uint16(i))))
			if !equals(pin.New16(pin.Split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}

	for i := 0; i < 4096; i += 512 {
		addr := address(i)
		t.Run(fmt.Sprintf("unsetting address %v", addr), func(t *testing.T) {
			clk.Deactivate()
			n := ram.Out(clk, addr, pin.New16(pin.Split16(uint16(i))))
			if !equals(pin.New16(pin.Split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Activate()
			n = ram.Out(clk, addr, pin.New16(pin.Split16(0)))
			if !equals(pin.New16(pin.Split16(0)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Deactivate()
		})
	}
}

func TestRAM16K_Out(t *testing.T) {
	equals := func(a [16]pin.Pin, b [16]pin.Signal) bool {
		converted := [16]pin.Signal{}
		for i := range a {
			converted[i] = a[i].Signal()
		}
		return converted == b
	}
	address := func(n int) [14]pin.Pin {
		n = n >> 12
		return [14]pin.Pin{
			pin.New(pin.Signal(n >> 0 & 1)),
			pin.New(pin.Signal(n >> 1 & 1)),
		}
	}
	ram := RAM16K{}
	clk := pin.New(pin.Inactive)
	for i := 0; i < 16384; i += 4096 {
		addr := address(i)
		t.Run(fmt.Sprintf("setting address %v", addr), func(t *testing.T) {
			clk.Deactivate()
			n := ram.Out(clk, addr, pin.New16(pin.Split16(uint16(i))))
			if !equals(pin.New16(pin.Split16(0)), n) {
				t.Errorf("expected %v but got %v", 0, n)
			}
			clk.Activate()
			n = ram.Out(clk, addr, pin.New16(pin.Split16(uint16(i))))
			if !equals(pin.New16(pin.Split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Deactivate()
			n = ram.Out(clk, addr, pin.New16(pin.Split16(uint16(i))))
			if !equals(pin.New16(pin.Split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}

	for i := 0; i < 16384; i += 4096 {
		addr := address(i)
		t.Run(fmt.Sprintf("unsetting address %v", addr), func(t *testing.T) {
			clk.Deactivate()
			n := ram.Out(clk, addr, pin.New16(pin.Split16(uint16(i))))
			if !equals(pin.New16(pin.Split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Activate()
			n = ram.Out(clk, addr, pin.New16(pin.Split16(0)))
			if !equals(pin.New16(pin.Split16(0)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Deactivate()
		})
	}
}
