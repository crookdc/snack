package chip

import (
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
	in := split16(65234)
	if out := reg.Out(Inactive, Wrap(&in)); out.Copy() != split16(0) {
		t.Errorf("expected 0 but got %v", out)
	}
	// The register should still yield the initialized value since the clock is inactive
	if out := reg.Out(Inactive, Wrap(&in)); out.Copy() != split16(0) {
		t.Errorf("expected 0 but got %v", out)
	}
	// Once the clock becomes active we should be receiving the newly set value
	if out := reg.Out(Active, Wrap(&in)); out.Copy() != split16(65234) {
		t.Errorf("expected 65234 but got %v", out)
	}
	// Subsequent outputs should remain the same regardless of the clocks value
	if out := reg.Out(Inactive, Wrap(&in)); out.Copy() != split16(65234) {
		t.Errorf("expected 65234 but got %v", out)
	}
	if out := reg.Out(Active, Wrap(&in)); out.Copy() != split16(65234) {
		t.Errorf("expected 65234 but got %v", out)
	}
	in = split16(40923)
	if out := reg.Out(Inactive, Wrap(&in)); out.Copy() != split16(65234) {
		t.Errorf("expected 65234 but got %v", out)
	}
	if out := reg.Out(Active, Wrap(&in)); out.Copy() != split16(40923) {
		t.Errorf("expected 40923 but got %v", out)
	}
}

func TestCounter_Out(t *testing.T) {
	t.Run("when load is set then sets value", func(t *testing.T) {
		ctr := PC{}
		load, inc, rst := Inactive, Inactive, Inactive
		load = Active
		in := split16(55467)
		out := ctr.Out(load, inc, rst, Wrap(&in))
		if out.Copy() != split16(55467) {
			t.Errorf("expected %v but got %v", split16(55467), out)
		}
		load = Inactive
		out = ctr.Out(load, inc, rst, NullWord)
		if out.Copy() != split16(55467) {
			t.Errorf("expected %v but got %v", split16(55467), out)
		}
		load = Active
		in = split16(33467)
		out = ctr.Out(load, inc, rst, Wrap(&in))
		if out.Copy() != split16(33467) {
			t.Errorf("expected %v but got %v", split16(33467), out)
		}
		load = Inactive
		out = ctr.Out(load, inc, rst, NullWord)
		if out.Copy() != split16(33467) {
			t.Errorf("expected %v but got %v", split16(33467), out)
		}
		load = Active
	})
	t.Run("when inc is set then increments value", func(t *testing.T) {
		ctr := PC{}
		load, inc, rst := Inactive, Inactive, Inactive
		out := ctr.Out(load, inc, rst, NullWord)
		if out.Copy() != split16(0) {
			t.Errorf("expected %v but got %v", split16(0), out)
		}
		inc = Active
		out = ctr.Out(load, inc, rst, NullWord)
		if out.Copy() != split16(1) {
			t.Errorf("expected %v but got %v", split16(1), out)
		}
		out = ctr.Out(load, inc, rst, NullWord)
		if out.Copy() != split16(2) {
			t.Errorf("expected %v but got %v", split16(2), out)
		}
	})
	t.Run("when rst is set then resets value", func(t *testing.T) {
		ctr := PC{}
		load, inc, rst := Inactive, Inactive, Inactive
		out := ctr.Out(load, inc, rst, NullWord)
		if out.Copy() != split16(0) {
			t.Errorf("expected %v but got %v", split16(0), out)
		}
		load = Active
		in := split16(5123)
		out = ctr.Out(load, inc, rst, Wrap(&in))
		if out.Copy() != split16(5123) {
			t.Errorf("expected %v but got %v", split16(5123), out)
		}
		load = Inactive
		rst = Active
		out = ctr.Out(load, inc, rst, NullWord)
		if out.Copy() != split16(0) {
			t.Errorf("expected %v but got %v", split16(1), out)
		}
	})
}
