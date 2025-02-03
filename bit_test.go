package snack

import "testing"

func TestNewBit(t *testing.T) {
	t.Run("given unset value", func(t *testing.T) {
		bit := NewBit(0)
		if bit.IsSet() {
			t.Error("got set bit")
		}
	})

	t.Run("given set value", func(t *testing.T) {
		bit := NewBit(1)
		if !bit.IsSet() {
			t.Error("got unset bit")
		}
	})

	t.Run("given invalid value", func(t *testing.T) {
		defer func() {
			if err := recover(); err == nil {
				t.Errorf("expected error but got nil")
			}
		}()
		NewBit(2)
	})
}
