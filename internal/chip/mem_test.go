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
	load := NewPin(Inactive)
	for i := range 8 {
		addr := address(i)
		t.Run(fmt.Sprintf("setting address %v", addr), func(t *testing.T) {
			load.Deactivate()
			n := ram.Out(load, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(0)), n) {
				t.Errorf("expected %v but got %v", 0, n)
			}
			load.Activate()
			n = ram.Out(load, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load.Deactivate()
			n = ram.Out(load, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}

	for i := range 8 {
		addr := address(i)
		t.Run(fmt.Sprintf("unsetting address %v", addr), func(t *testing.T) {
			load.Deactivate()
			n := ram.Out(load, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load.Activate()
			n = ram.Out(load, addr, NewPin16(split16(0)))
			if !equals(NewPin16(split16(0)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load.Deactivate()
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
	load := NewPin(Inactive)
	for i := 0; i < 64; i += 8 {
		addr := address(i)
		t.Run(fmt.Sprintf("setting address %v", addr), func(t *testing.T) {
			load.Deactivate()
			n := ram.Out(load, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(0)), n) {
				t.Errorf("expected %v but got %v", 0, n)
			}
			load.Activate()
			n = ram.Out(load, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load.Deactivate()
			n = ram.Out(load, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}

	for i := 0; i < 64; i += 8 {
		addr := address(i)
		t.Run(fmt.Sprintf("unsetting address %v", addr), func(t *testing.T) {
			load.Deactivate()
			n := ram.Out(load, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load.Activate()
			n = ram.Out(load, addr, NewPin16(split16(0)))
			if !equals(NewPin16(split16(0)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load.Deactivate()
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
	load := NewPin(Inactive)
	for i := 0; i < 512; i += 64 {
		addr := address(i)
		t.Run(fmt.Sprintf("setting address %v", addr), func(t *testing.T) {
			load.Deactivate()
			n := ram.Out(load, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(0)), n) {
				t.Errorf("expected %v but got %v", 0, n)
			}
			load.Activate()
			n = ram.Out(load, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load.Deactivate()
			n = ram.Out(load, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}

	for i := 0; i < 512; i += 64 {
		addr := address(i)
		t.Run(fmt.Sprintf("unsetting address %v", addr), func(t *testing.T) {
			load.Deactivate()
			n := ram.Out(load, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load.Activate()
			n = ram.Out(load, addr, NewPin16(split16(0)))
			if !equals(NewPin16(split16(0)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load.Deactivate()
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
	load := NewPin(Inactive)
	for i := 0; i < 4096; i += 512 {
		addr := address(i)
		t.Run(fmt.Sprintf("setting address %v", addr), func(t *testing.T) {
			load.Deactivate()
			n := ram.Out(load, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(0)), n) {
				t.Errorf("expected %v but got %v", 0, n)
			}
			load.Activate()
			n = ram.Out(load, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load.Deactivate()
			n = ram.Out(load, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}

	for i := 0; i < 4096; i += 512 {
		addr := address(i)
		t.Run(fmt.Sprintf("unsetting address %v", addr), func(t *testing.T) {
			load.Deactivate()
			n := ram.Out(load, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load.Activate()
			n = ram.Out(load, addr, NewPin16(split16(0)))
			if !equals(NewPin16(split16(0)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load.Deactivate()
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
	load := NewPin(Inactive)
	for i := 0; i < 16384; i += 4096 {
		addr := address(i)
		t.Run(fmt.Sprintf("setting address %v", addr), func(t *testing.T) {
			load.Deactivate()
			n := ram.Out(load, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(0)), n) {
				t.Errorf("expected %v but got %v", 0, n)
			}
			load.Activate()
			n = ram.Out(load, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load.Deactivate()
			n = ram.Out(load, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}

	for i := 0; i < 16384; i += 4096 {
		addr := address(i)
		t.Run(fmt.Sprintf("unsetting address %v", addr), func(t *testing.T) {
			load.Deactivate()
			n := ram.Out(load, addr, NewPin16(split16(uint16(i))))
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load.Activate()
			n = ram.Out(load, addr, NewPin16(split16(0)))
			if !equals(NewPin16(split16(0)), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
			load.Deactivate()
		})
	}
}

func TestROM32K_Out(t *testing.T) {
	equals := func(a [16]Pin, b [16]Signal) bool {
		converted := [16]Signal{}
		for i := range a {
			converted[i] = a[i].Signal()
		}
		return converted == b
	}
	address := func(n int) [15]Pin {
		n = n >> 14
		return [15]Pin{
			NewPin(Signal(n >> 0 & 1)),
			NewPin(Signal(n >> 1 & 1)),
		}
	}
	rom := ROM32K{}
	for i := 0; i < 32768; i += 16384 {
		addr := address(i)
		t.Run(fmt.Sprintf("reading address %v", addr), func(t *testing.T) {
			// Reach in and set the ROM on the provided address, we cannot use the load bit for this like we have in the
			// RAM testing since the ROM does not allow writes
			nxt := [14]Pin{}
			copy(nxt[:], addr[1:])
			Mux2Way16(
				addr[0].Signal(),
				rom.chips[0].Out(NewPin(Active), nxt, NewPin16(split16(uint16(i)))),
				rom.chips[1].Out(NewPin(Active), nxt, NewPin16(split16(uint16(i)))),
			)
			n := rom.Out(addr)
			n = rom.Out(addr)
			if !equals(NewPin16(split16(uint16(i))), n) {
				t.Errorf("expected %v but got %v", i, n)
			}
		})
	}
}

func TestCounter_Out(t *testing.T) {
	t.Run("when load is set then sets value", func(t *testing.T) {
		ctr := ProgramCounter{}
		load, inc, rst := NewPin(Inactive), NewPin(Inactive), NewPin(Inactive)
		load.Activate()
		out := ctr.Out(load, inc, rst, NewPin16(split16(55467)))
		if out != split16(55467) {
			t.Errorf("expected %v but got %v", split16(55467), out)
		}
		load.Deactivate()
		out = ctr.Out(load, inc, rst, NewPin16(split16(0)))
		if out != split16(55467) {
			t.Errorf("expected %v but got %v", split16(55467), out)
		}
		load.Activate()
		out = ctr.Out(load, inc, rst, NewPin16(split16(33467)))
		if out != split16(33467) {
			t.Errorf("expected %v but got %v", split16(33467), out)
		}
		load.Deactivate()
		out = ctr.Out(load, inc, rst, NewPin16(split16(0)))
		if out != split16(33467) {
			t.Errorf("expected %v but got %v", split16(33467), out)
		}
		load.Activate()
	})
	t.Run("when inc is set then increments value", func(t *testing.T) {
		ctr := ProgramCounter{}
		load, inc, rst := NewPin(Inactive), NewPin(Inactive), NewPin(Inactive)
		out := ctr.Out(load, inc, rst, NewPin16(split16(0)))
		if out != split16(0) {
			t.Errorf("expected %v but got %v", split16(0), out)
		}
		inc.Activate()
		out = ctr.Out(load, inc, rst, NewPin16(split16(0)))
		if out != split16(1) {
			t.Errorf("expected %v but got %v", split16(1), out)
		}
		out = ctr.Out(load, inc, rst, NewPin16(split16(0)))
		if out != split16(2) {
			t.Errorf("expected %v but got %v", split16(2), out)
		}
	})
	t.Run("when rst is set then resets value", func(t *testing.T) {
		ctr := ProgramCounter{}
		load, inc, rst := NewPin(Inactive), NewPin(Inactive), NewPin(Inactive)
		out := ctr.Out(load, inc, rst, NewPin16(split16(0)))
		if out != split16(0) {
			t.Errorf("expected %v but got %v", split16(0), out)
		}
		load.Activate()
		out = ctr.Out(load, inc, rst, NewPin16(split16(5123)))
		if out != split16(5123) {
			t.Errorf("expected %v but got %v", split16(5123), out)
		}
		load.Deactivate()
		rst.Activate()
		out = ctr.Out(load, inc, rst, NewPin16(split16(0)))
		if out != split16(0) {
			t.Errorf("expected %v but got %v", split16(1), out)
		}
	})
}

func TestMemory_Out(t *testing.T) {
	ramAddress := func(n uint16) [14]Pin {
		res := [14]Signal{}
		for i := range 14 {
			res[i] = Signal(uint8(n>>(13-i)) & 1)
		}
		p := [14]Pin{}
		for i := range 14 {
			p[i] = NewPin(res[i])
		}
		return p
	}
	var ramw = []struct {
		address uint16
		value   uint16
		r       uint16
	}{
		{
			address: 0,
			value:   0xF0FF,
			r:       0xF0FF,
		},
		{
			address: 12345,
			value:   0x0FFF,
			r:       0x0FFF,
		},
		{
			address: 16383,
			value:   0xFFFF,
			r:       0xFFFF,
		},
		{
			address: 16384, // The address for this test is outside the RAM range
			value:   0xFFFF,
			r:       0,
		},
	}
	for _, w := range ramw {
		t.Run(fmt.Sprintf("writing %v to %v address", w.value, w.address), func(t *testing.T) {
			mem := Memory{}
			mem.Out(NewPin(Active), NewPin15(split15(w.address)), NewPin16(split16(w.value)))
			n := mem.ram.Out(NewPin(Inactive), ramAddress(w.address), [16]Pin{})
			if n != split16(w.r) {
				t.Errorf("expected RAM n to be %v but got %v", split16(w.r), n)
			}
		})
	}

	screenAddress := func(n uint16) [13]Pin {
		res := [13]Signal{}
		for i := range 13 {
			res[i] = Signal(uint8(n>>(12-i)) & 1)
		}
		p := [13]Pin{}
		for i := range 13 {
			p[i] = NewPin(res[i])
		}
		return p
	}
	var screenw = []struct {
		addr  uint16
		value uint16
		r     uint16
	}{
		{
			addr:  0b011_0100_0110_0011,
			value: 0xF0FF,
			r:     0,
		},
		{
			addr:  0b100_0100_0110_0011,
			value: 0x0FFF,
			r:     0x0FFF,
		},
		{
			addr:  0b101_1100_0110_0011,
			value: 0xFFFF,
			r:     0xFFFF,
		},
		{
			addr:  0b111_1100_0110_0011,
			value: 0xFFFF,
			r:     0,
		},
	}
	for _, w := range screenw {
		t.Run(fmt.Sprintf("writing %v to %v address", w.value, w.addr), func(t *testing.T) {
			mem := Memory{}
			mem.Out(NewPin(Active), NewPin15(split15(w.addr)), NewPin16(split16(w.value)))
			n := mem.screen.Out(NewPin(Inactive), screenAddress(w.addr), [16]Pin{})
			if n != split16(w.r) {
				t.Errorf("expected screen n to be %v but got %v", split16(w.r), n)
			}
		})
	}

	t.Run("reading and writing keyboard", func(t *testing.T) {
		addr := uint16(24576)
		mem := Memory{}
		mem.Out(NewPin(Active), NewPin15(split15(addr)), NewPin16(split16(1012)))
		n := mem.keyboard.Out(NewPin(Inactive), [16]Pin{})
		if n != split16(1012) {
			t.Errorf("expected keyboard to be %v but got %v", split16(1012), n)
		}
	})
}
