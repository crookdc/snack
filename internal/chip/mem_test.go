package chip

import (
	"fmt"
	"testing"
)

func TestBit_Out(t *testing.T) {
	bit := Bit{}
	if b := bit.Out(Inactive, Active); b == Active {
		t.Errorf("expected inactive pin but got active")
	}
	if b := bit.Out(Active, Active); b == Inactive {
		t.Errorf("expected active pin but got inactive")
	}
	if b := bit.Out(Inactive, Active); b == Inactive {
		t.Errorf("expected active pin but got inactive")
	}
	if b := bit.Out(Active, Active); b == Inactive {
		t.Errorf("expected active pin but got inactive")
	}
	if b := bit.Out(Inactive, Inactive); b == Inactive {
		t.Errorf("expected active pin but got inactive")
	}
	if b := bit.Out(Active, Inactive); b == Active {
		t.Errorf("expected inactive pin but got active")
	}
}

func TestRegister(t *testing.T) {
	reg := new(Register)
	if out := reg.Out(Inactive, split16(65234)); out != split16(0) {
		t.Errorf("expected 0 but got %v", out)
	}
	// The register should still yield the initialized value since the clock is inactive
	if out := reg.Out(Inactive, split16(65234)); out != split16(0) {
		t.Errorf("expected 0 but got %v", out)
	}
	// Once the clock becomes active we should be receiving the newly set value
	if out := reg.Out(Active, split16(65234)); out != split16(65234) {
		t.Errorf("expected 65234 but got %v", out)
	}
	// Subsequent outputs should remain the same regardless of the clocks value
	if out := reg.Out(Inactive, split16(65234)); out != split16(65234) {
		t.Errorf("expected 65234 but got %v", out)
	}
	if out := reg.Out(Active, split16(65234)); out != split16(65234) {
		t.Errorf("expected 65234 but got %v", out)
	}
	if out := reg.Out(Inactive, split16(40923)); out != split16(65234) {
		t.Errorf("expected 65234 but got %v", out)
	}
	if out := reg.Out(Active, split16(40923)); out != split16(40923) {
		t.Errorf("expected 40923 but got %v", out)
	}
}

func TestRAM8_Out(t *testing.T) {
	equals := func(a [16]Signal, b [16]Signal) bool {
		converted := [16]Signal{}
		for i := range a {
			converted[i] = a[i]
		}
		return converted == b
	}
	address := func(n int) [3]Signal {
		return [3]Signal{
			Signal(n >> 0 & 1),
			Signal(n >> 1 & 1),
			Signal(n >> 2 & 1),
		}
	}
	ram := RAM8{}
	load := Inactive
	for i := range 8 {
		addr := address(i)
		t.Run(fmt.Sprintf("setting address %v", addr), func(t *testing.T) {
			load = Inactive
			n := ram.Out(load, addr, split16(uint16(i)))
			if !equals(split16(0), n) {
				t.Errorf("expected %v but got %v", 0, n)
			}
			load = Active
			n = ram.Out(load, addr, split16(uint16(i)))
			if !equals(split16(uint16(i)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load = Inactive
			n = ram.Out(load, addr, split16(uint16(i)))
			if !equals(split16(uint16(i)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}

	for i := range 8 {
		addr := address(i)
		t.Run(fmt.Sprintf("unsetting address %v", addr), func(t *testing.T) {
			load = Inactive
			n := ram.Out(load, addr, split16(uint16(i)))
			if !equals(split16(uint16(i)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load = Active
			n = ram.Out(load, addr, split16(0))
			if !equals(split16(0), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load = Inactive
		})
	}
}

func TestRAM64_Out(t *testing.T) {
	equals := func(a [16]Signal, b [16]Signal) bool {
		converted := [16]Signal{}
		for i := range a {
			converted[i] = a[i]
		}
		return converted == b
	}
	address := func(n int) [6]Signal {
		n = n >> 3
		return [6]Signal{
			Signal(n >> 0 & 1),
			Signal(n >> 1 & 1),
			Signal(n >> 2 & 1),
		}
	}
	ram := RAM64{}
	load := Inactive
	for i := 0; i < 64; i += 8 {
		addr := address(i)
		t.Run(fmt.Sprintf("setting address %v", addr), func(t *testing.T) {
			load = Inactive
			n := ram.Out(load, addr, split16(uint16(i)))
			if !equals(split16(0), n) {
				t.Errorf("expected %v but got %v", 0, n)
			}
			load = Active
			n = ram.Out(load, addr, split16(uint16(i)))
			if !equals(split16(uint16(i)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load = Inactive
			n = ram.Out(load, addr, split16(uint16(i)))
			if !equals(split16(uint16(i)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}

	for i := 0; i < 64; i += 8 {
		addr := address(i)
		t.Run(fmt.Sprintf("unsetting address %v", addr), func(t *testing.T) {
			load = Inactive
			n := ram.Out(load, addr, split16(uint16(i)))
			if !equals(split16(uint16(i)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load = Active
			n = ram.Out(load, addr, split16(0))
			if !equals(split16(0), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load = Inactive
		})
	}
}

func TestRAM512_Out(t *testing.T) {
	equals := func(a [16]Signal, b [16]Signal) bool {
		converted := [16]Signal{}
		for i := range a {
			converted[i] = a[i]
		}
		return converted == b
	}
	address := func(n int) [9]Signal {
		n = n >> 6
		return [9]Signal{
			Signal(n >> 0 & 1),
			Signal(n >> 1 & 1),
			Signal(n >> 2 & 1),
		}
	}
	ram := RAM512{}
	load := Inactive
	for i := 0; i < 512; i += 64 {
		addr := address(i)
		t.Run(fmt.Sprintf("setting address %v", addr), func(t *testing.T) {
			load = Inactive
			n := ram.Out(load, addr, split16(uint16(i)))
			if !equals(split16(0), n) {
				t.Errorf("expected %v but got %v", 0, n)
			}
			load = Active
			n = ram.Out(load, addr, split16(uint16(i)))
			if !equals(split16(uint16(i)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load = Inactive
			n = ram.Out(load, addr, split16(uint16(i)))
			if !equals(split16(uint16(i)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}

	for i := 0; i < 512; i += 64 {
		addr := address(i)
		t.Run(fmt.Sprintf("unsetting address %v", addr), func(t *testing.T) {
			load = Inactive
			n := ram.Out(load, addr, split16(uint16(i)))
			if !equals(split16(uint16(i)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load = Active
			n = ram.Out(load, addr, split16(0))
			if !equals(split16(0), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load = Inactive
		})
	}
}

func TestRAM4K_Out(t *testing.T) {
	equals := func(a [16]Signal, b [16]Signal) bool {
		converted := [16]Signal{}
		for i := range a {
			converted[i] = a[i]
		}
		return converted == b
	}
	address := func(n int) [12]Signal {
		n = n >> 9
		return [12]Signal{
			Signal(n >> 0 & 1),
			Signal(n >> 1 & 1),
			Signal(n >> 2 & 1),
		}
	}
	ram := RAM4K{}
	load := Inactive
	for i := 0; i < 4096; i += 512 {
		addr := address(i)
		t.Run(fmt.Sprintf("setting address %v", addr), func(t *testing.T) {
			load = Inactive
			n := ram.Out(load, addr, split16(uint16(i)))
			if !equals(split16(0), n) {
				t.Errorf("expected %v but got %v", 0, n)
			}
			load = Active
			n = ram.Out(load, addr, split16(uint16(i)))
			if !equals(split16(uint16(i)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load = Inactive
			n = ram.Out(load, addr, split16(uint16(i)))
			if !equals(split16(uint16(i)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}

	for i := 0; i < 4096; i += 512 {
		addr := address(i)
		t.Run(fmt.Sprintf("unsetting address %v", addr), func(t *testing.T) {
			load = Inactive
			n := ram.Out(load, addr, split16(uint16(i)))
			if !equals(split16(uint16(i)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load = Active
			n = ram.Out(load, addr, split16(0))
			if !equals(split16(0), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load = Inactive
		})
	}
}

func TestRAM16K_Out(t *testing.T) {
	equals := func(a [16]Signal, b [16]Signal) bool {
		converted := [16]Signal{}
		for i := range a {
			converted[i] = a[i]
		}
		return converted == b
	}
	address := func(n int) [14]Signal {
		n = n >> 12
		return [14]Signal{
			Signal(n >> 0 & 1),
			Signal(n >> 1 & 1),
		}
	}
	ram := RAM16K{}
	load := Inactive
	for i := 0; i < 16384; i += 4096 {
		addr := address(i)
		t.Run(fmt.Sprintf("setting address %v", addr), func(t *testing.T) {
			load = Inactive
			n := ram.Out(load, addr, split16(uint16(i)))
			if !equals(split16(0), n) {
				t.Errorf("expected %v but got %v", 0, n)
			}
			load = Active
			n = ram.Out(load, addr, split16(uint16(i)))
			if !equals(split16(uint16(i)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load = Inactive
			n = ram.Out(load, addr, split16(uint16(i)))
			if !equals(split16(uint16(i)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}

	for i := 0; i < 16384; i += 4096 {
		addr := address(i)
		t.Run(fmt.Sprintf("unsetting address %v", addr), func(t *testing.T) {
			load = Inactive
			n := ram.Out(load, addr, split16(uint16(i)))
			if !equals(split16(uint16(i)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load = Active
			n = ram.Out(load, addr, split16(0))
			if !equals(split16(0), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load = Inactive
		})
	}
}

func TestCounter_Out(t *testing.T) {
	t.Run("when load is set then sets value", func(t *testing.T) {
		ctr := ProgramCounter{}
		load, inc, rst := Inactive, Inactive, Inactive
		load = Active
		out := ctr.Out(load, inc, rst, split16(55467))
		if out != split16(55467) {
			t.Errorf("expected %v but got %v", split16(55467), out)
		}
		load = Inactive
		out = ctr.Out(load, inc, rst, split16(0))
		if out != split16(55467) {
			t.Errorf("expected %v but got %v", split16(55467), out)
		}
		load = Active
		out = ctr.Out(load, inc, rst, split16(33467))
		if out != split16(33467) {
			t.Errorf("expected %v but got %v", split16(33467), out)
		}
		load = Inactive
		out = ctr.Out(load, inc, rst, split16(0))
		if out != split16(33467) {
			t.Errorf("expected %v but got %v", split16(33467), out)
		}
		load = Active
	})
	t.Run("when inc is set then increments value", func(t *testing.T) {
		ctr := ProgramCounter{}
		load, inc, rst := Inactive, Inactive, Inactive
		out := ctr.Out(load, inc, rst, split16(0))
		if out != split16(0) {
			t.Errorf("expected %v but got %v", split16(0), out)
		}
		inc = Active
		out = ctr.Out(load, inc, rst, split16(0))
		if out != split16(1) {
			t.Errorf("expected %v but got %v", split16(1), out)
		}
		out = ctr.Out(load, inc, rst, split16(0))
		if out != split16(2) {
			t.Errorf("expected %v but got %v", split16(2), out)
		}
	})
	t.Run("when rst is set then resets value", func(t *testing.T) {
		ctr := ProgramCounter{}
		load, inc, rst := Inactive, Inactive, Inactive
		out := ctr.Out(load, inc, rst, split16(0))
		if out != split16(0) {
			t.Errorf("expected %v but got %v", split16(0), out)
		}
		load = Active
		out = ctr.Out(load, inc, rst, split16(5123))
		if out != split16(5123) {
			t.Errorf("expected %v but got %v", split16(5123), out)
		}
		load = Inactive
		rst = Active
		out = ctr.Out(load, inc, rst, split16(0))
		if out != split16(0) {
			t.Errorf("expected %v but got %v", split16(1), out)
		}
	})
}
