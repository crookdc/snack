package chip

import (
	"fmt"
	"testing"
)

func TestBit_Out(t *testing.T) {
	bit := Bit{}
	if b := bit.Out(NewPin(Inactive), NewPin(Active)); b == Active {
		t.Errorf("expected inactive pin but got active")
	}
	if b := bit.Out(NewPin(Active), NewPin(Active)); b == Inactive {
		t.Errorf("expected active pin but got inactive")
	}
	if b := bit.Out(NewPin(Inactive), NewPin(Active)); b == Inactive {
		t.Errorf("expected active pin but got inactive")
	}
	if b := bit.Out(NewPin(Active), NewPin(Active)); b == Inactive {
		t.Errorf("expected active pin but got inactive")
	}
	if b := bit.Out(NewPin(Inactive), NewPin(Inactive)); b == Inactive {
		t.Errorf("expected active pin but got inactive")
	}
	if b := bit.Out(NewPin(Active), NewPin(Inactive)); b == Active {
		t.Errorf("expected inactive pin but got active")
	}
}

func TestRegister(t *testing.T) {
	reg := new(Register)
	if out := reg.Out(NewPin(Inactive), NewPin16(split16(65234))); out != split16(0) {
		t.Errorf("expected 0 but got %v", out)
	}
	// The register should still yield the initialized value since the clock is inactive
	if out := reg.Out(NewPin(Inactive), NewPin16(split16(65234))); out != split16(0) {
		t.Errorf("expected 0 but got %v", out)
	}
	// Once the clock becomes active we should be receiving the newly set value
	if out := reg.Out(NewPin(Active), NewPin16(split16(65234))); out != split16(65234) {
		t.Errorf("expected 65234 but got %v", out)
	}
	// Subsequent outputs should remain the same regardless of the clocks value
	if out := reg.Out(NewPin(Inactive), NewPin16(split16(65234))); out != split16(65234) {
		t.Errorf("expected 65234 but got %v", out)
	}
	if out := reg.Out(NewPin(Active), NewPin16(split16(65234))); out != split16(65234) {
		t.Errorf("expected 65234 but got %v", out)
	}
	if out := reg.Out(NewPin(Inactive), NewPin16(split16(40923))); out != split16(65234) {
		t.Errorf("expected 65234 but got %v", out)
	}
	if out := reg.Out(NewPin(Active), NewPin16(split16(40923))); out != split16(40923) {
		t.Errorf("expected 40923 but got %v", out)
	}
}

func TestRAM8_Out(t *testing.T) {
	equals := func(a [16]Pin, b [16]Signal) bool {
		converted := [16]Signal{}
		for i := range a {
			converted[i] = a[i].Signal()
		}
		return converted == b
	}
	address := func(n int) [3]Pin {
		return [3]Pin{
			NewPin(Signal(n >> 0 & 1)),
			NewPin(Signal(n >> 1 & 1)),
			NewPin(Signal(n >> 2 & 1)),
		}
	}
	ram := RAM8{}
	clk := NewPin(Inactive)
	for i := range 8 {
		addr := address(i)
		t.Run(fmt.Sprintf("setting address %v", addr), func(t *testing.T) {
			clk.Deactivate()
			n := ram.Out(clk, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(0)), n) {
				t.Errorf("expected %v but got %v", 0, n)
			}
			clk.Activate()
			n = ram.Out(clk, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Deactivate()
			n = ram.Out(clk, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}

	for i := range 8 {
		addr := address(i)
		t.Run(fmt.Sprintf("unsetting address %v", addr), func(t *testing.T) {
			clk.Deactivate()
			n := ram.Out(clk, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Activate()
			n = ram.Out(clk, addr, NewPin16(split16(0)))
			if !equals(NewPin16(split16(0)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Deactivate()
		})
	}
}

func TestRAM64_Out(t *testing.T) {
	equals := func(a [16]Pin, b [16]Signal) bool {
		converted := [16]Signal{}
		for i := range a {
			converted[i] = a[i].Signal()
		}
		return converted == b
	}
	address := func(n int) [6]Pin {
		n = n >> 3
		return [6]Pin{
			NewPin(Signal(n >> 0 & 1)),
			NewPin(Signal(n >> 1 & 1)),
			NewPin(Signal(n >> 2 & 1)),
		}
	}
	ram := RAM64{}
	clk := NewPin(Inactive)
	for i := 0; i < 64; i += 8 {
		addr := address(i)
		t.Run(fmt.Sprintf("setting address %v", addr), func(t *testing.T) {
			clk.Deactivate()
			n := ram.Out(clk, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(0)), n) {
				t.Errorf("expected %v but got %v", 0, n)
			}
			clk.Activate()
			n = ram.Out(clk, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Deactivate()
			n = ram.Out(clk, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}

	for i := 0; i < 64; i += 8 {
		addr := address(i)
		t.Run(fmt.Sprintf("unsetting address %v", addr), func(t *testing.T) {
			clk.Deactivate()
			n := ram.Out(clk, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Activate()
			n = ram.Out(clk, addr, NewPin16(split16(0)))
			if !equals(NewPin16(split16(0)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Deactivate()
		})
	}
}

func TestRAM512_Out(t *testing.T) {
	equals := func(a [16]Pin, b [16]Signal) bool {
		converted := [16]Signal{}
		for i := range a {
			converted[i] = a[i].Signal()
		}
		return converted == b
	}
	address := func(n int) [9]Pin {
		n = n >> 6
		return [9]Pin{
			NewPin(Signal(n >> 0 & 1)),
			NewPin(Signal(n >> 1 & 1)),
			NewPin(Signal(n >> 2 & 1)),
		}
	}
	ram := RAM512{}
	clk := NewPin(Inactive)
	for i := 0; i < 512; i += 64 {
		addr := address(i)
		t.Run(fmt.Sprintf("setting address %v", addr), func(t *testing.T) {
			clk.Deactivate()
			n := ram.Out(clk, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(0)), n) {
				t.Errorf("expected %v but got %v", 0, n)
			}
			clk.Activate()
			n = ram.Out(clk, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Deactivate()
			n = ram.Out(clk, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}

	for i := 0; i < 512; i += 64 {
		addr := address(i)
		t.Run(fmt.Sprintf("unsetting address %v", addr), func(t *testing.T) {
			clk.Deactivate()
			n := ram.Out(clk, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Activate()
			n = ram.Out(clk, addr, NewPin16(split16(0)))
			if !equals(NewPin16(split16(0)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Deactivate()
		})
	}
}

func TestRAM4K_Out(t *testing.T) {
	equals := func(a [16]Pin, b [16]Signal) bool {
		converted := [16]Signal{}
		for i := range a {
			converted[i] = a[i].Signal()
		}
		return converted == b
	}
	address := func(n int) [12]Pin {
		n = n >> 9
		return [12]Pin{
			NewPin(Signal(n >> 0 & 1)),
			NewPin(Signal(n >> 1 & 1)),
			NewPin(Signal(n >> 2 & 1)),
		}
	}
	ram := RAM4K{}
	clk := NewPin(Inactive)
	for i := 0; i < 4096; i += 512 {
		addr := address(i)
		t.Run(fmt.Sprintf("setting address %v", addr), func(t *testing.T) {
			clk.Deactivate()
			n := ram.Out(clk, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(0)), n) {
				t.Errorf("expected %v but got %v", 0, n)
			}
			clk.Activate()
			n = ram.Out(clk, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Deactivate()
			n = ram.Out(clk, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}

	for i := 0; i < 4096; i += 512 {
		addr := address(i)
		t.Run(fmt.Sprintf("unsetting address %v", addr), func(t *testing.T) {
			clk.Deactivate()
			n := ram.Out(clk, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Activate()
			n = ram.Out(clk, addr, NewPin16(split16(0)))
			if !equals(NewPin16(split16(0)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Deactivate()
		})
	}
}

func TestRAM16K_Out(t *testing.T) {
	equals := func(a [16]Pin, b [16]Signal) bool {
		converted := [16]Signal{}
		for i := range a {
			converted[i] = a[i].Signal()
		}
		return converted == b
	}
	address := func(n int) [14]Pin {
		n = n >> 12
		return [14]Pin{
			NewPin(Signal(n >> 0 & 1)),
			NewPin(Signal(n >> 1 & 1)),
		}
	}
	ram := RAM16K{}
	clk := NewPin(Inactive)
	for i := 0; i < 16384; i += 4096 {
		addr := address(i)
		t.Run(fmt.Sprintf("setting address %v", addr), func(t *testing.T) {
			clk.Deactivate()
			n := ram.Out(clk, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(0)), n) {
				t.Errorf("expected %v but got %v", 0, n)
			}
			clk.Activate()
			n = ram.Out(clk, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Deactivate()
			n = ram.Out(clk, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}

	for i := 0; i < 16384; i += 4096 {
		addr := address(i)
		t.Run(fmt.Sprintf("unsetting address %v", addr), func(t *testing.T) {
			clk.Deactivate()
			n := ram.Out(clk, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Activate()
			n = ram.Out(clk, addr, NewPin16(split16(0)))
			if !equals(NewPin16(split16(0)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			clk.Deactivate()
		})
	}
}

func TestCounter_Out(t *testing.T) {
	t.Run("when load is set then sets value", func(t *testing.T) {
		ctr := Counter{}
		clk, inc, rst := NewPin(Inactive), NewPin(Inactive), NewPin(Inactive)
		clk.Activate()
		out := ctr.Out(clk, inc, rst, NewPin16(split16(55467)))
		if out != split16(55467) {
			t.Errorf("expected %v but got %v", split16(55467), out)
		}
		clk.Deactivate()
		out = ctr.Out(clk, inc, rst, NewPin16(split16(0)))
		if out != split16(55467) {
			t.Errorf("expected %v but got %v", split16(55467), out)
		}
		clk.Activate()
		out = ctr.Out(clk, inc, rst, NewPin16(split16(33467)))
		if out != split16(33467) {
			t.Errorf("expected %v but got %v", split16(33467), out)
		}
		clk.Deactivate()
		out = ctr.Out(clk, inc, rst, NewPin16(split16(0)))
		if out != split16(33467) {
			t.Errorf("expected %v but got %v", split16(33467), out)
		}
		clk.Activate()
	})
	t.Run("when inc is set then increments value", func(t *testing.T) {
		ctr := Counter{}
		clk, inc, rst := NewPin(Inactive), NewPin(Inactive), NewPin(Inactive)
		out := ctr.Out(clk, inc, rst, NewPin16(split16(0)))
		if out != split16(0) {
			t.Errorf("expected %v but got %v", split16(0), out)
		}
		inc.Activate()
		out = ctr.Out(clk, inc, rst, NewPin16(split16(0)))
		if out != split16(1) {
			t.Errorf("expected %v but got %v", split16(1), out)
		}
		out = ctr.Out(clk, inc, rst, NewPin16(split16(0)))
		if out != split16(2) {
			t.Errorf("expected %v but got %v", split16(2), out)
		}
	})
	t.Run("when rst is set then resets value", func(t *testing.T) {
		ctr := Counter{}
		clk, inc, rst := NewPin(Inactive), NewPin(Inactive), NewPin(Inactive)
		out := ctr.Out(clk, inc, rst, NewPin16(split16(0)))
		if out != split16(0) {
			t.Errorf("expected %v but got %v", split16(0), out)
		}
		clk.Activate()
		out = ctr.Out(clk, inc, rst, NewPin16(split16(5123)))
		if out != split16(5123) {
			t.Errorf("expected %v but got %v", split16(5123), out)
		}
		clk.Deactivate()
		rst.Activate()
		out = ctr.Out(clk, inc, rst, NewPin16(split16(0)))
		if out != split16(0) {
			t.Errorf("expected %v but got %v", split16(1), out)
		}
	})
}
