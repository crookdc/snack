package chip

import (
	"fmt"
	"testing"
)

func TestALU_Call(t *testing.T) {
	type assertion struct {
		x uint16
		y uint16

		r  uint16
		zr Signal
		ng Signal
	}
	runner := func(t *testing.T, alu *ALU, assertions []assertion) {
		for _, a := range assertions {
			t.Run(fmt.Sprintf("given x: %v, y: %v ", a.x, a.y), func(t *testing.T) {
				opx := split16(a.x)
				opy := split16(a.y)
				r, zr, ng := alu.Out(Wrap(&opx), Wrap(&opy))
				if r.Copy() != split16(a.r) {
					t.Errorf("expected r %v but got %v", a.r, r)
				}
				if zr != a.zr {
					t.Errorf("expected zr %v but got %v", a.zr, zr)
				}
				if ng != a.ng {
					t.Errorf("expected ng %v but got %v", a.ng, ng)
				}
			})
		}
	}
	t.Run("x + y", func(t *testing.T) {
		alu := ALU{
			ZX: Inactive,
			NX: Inactive,
			ZY: Inactive,
			NY: Inactive,
			F:  Active,
			NO: Inactive,
		}
		runner(t, &alu, []assertion{
			{
				x: 512,
				y: 512,
				r: 1024,
			},
			{
				x: 256,
				y: 512,
				r: 768,
			},
			{
				x: 255,
				y: 0,
				r: 255,
			},
			{
				x: 0,
				y: 255,
				r: 255,
			},
		})
	})

	t.Run("x + 1", func(t *testing.T) {
		alu := ALU{
			ZX: Inactive,
			NX: Active,
			ZY: Active,
			NY: Active,
			F:  Active,
			NO: Active,
		}
		runner(t, &alu, []assertion{
			{
				x: 256,
				y: 1024,
				r: 257,
			},
			{
				x: 1111,
				r: 1112,
			},
			{
				x:  65535,
				r:  0,
				zr: 1,
			},
		})
	})

	t.Run("x - 1", func(t *testing.T) {
		alu := ALU{
			ZX: Inactive,
			NX: Inactive,
			ZY: Active,
			NY: Active,
			F:  Active,
			NO: Inactive,
		}
		runner(t, &alu, []assertion{
			{
				x: 256,
				y: 1024,
				r: 255,
			},
			{
				x: 1111,
				r: 1110,
			},
			{
				x:  0,
				r:  65535,
				ng: 1,
			},
		})
	})

	t.Run("y - 1", func(t *testing.T) {
		alu := ALU{
			ZX: Active,
			NX: Active,
			ZY: Inactive,
			NY: Inactive,
			F:  Active,
			NO: Inactive,
		}
		runner(t, &alu, []assertion{
			{
				x: 256,
				y: 1024,
				r: 1023,
			},
			{
				y: 1111,
				r: 1110,
			},
			{
				y:  0,
				r:  65535,
				ng: 1,
			},
		})
	})

	t.Run("y + 1", func(t *testing.T) {
		alu := ALU{
			ZX: Active,
			NX: Active,
			ZY: Inactive,
			NY: Active,
			F:  Active,
			NO: Active,
		}
		runner(t, &alu, []assertion{
			{
				x: 256,
				y: 1024,
				r: 1025,
			},
			{
				y: 1111,
				r: 1112,
			},
			{
				y:  65535,
				r:  0,
				zr: 1,
			},
		})
	})

	t.Run("x - y", func(t *testing.T) {
		alu := ALU{
			ZX: Inactive,
			NX: Active,
			ZY: Inactive,
			NY: Inactive,
			F:  Active,
			NO: Active,
		}
		runner(t, &alu, []assertion{
			{
				x: 256,
				y: 255,
				r: 1,
			},
			{
				x: 1111,
				r: 1111,
			},
			{
				x:  1234,
				y:  1234,
				r:  0,
				zr: 1,
			},
			{
				x: 1024,
				y: 512,
				r: 512,
			},
		})
	})

	t.Run("y - x", func(t *testing.T) {
		alu := ALU{
			ZX: Inactive,
			NX: Inactive,
			ZY: Inactive,
			NY: Active,
			F:  Active,
			NO: Active,
		}
		runner(t, &alu, []assertion{
			{
				x: 255,
				y: 256,
				r: 1,
			},
			{
				y: 1111,
				r: 1111,
			},
			{
				x:  1234,
				y:  1234,
				r:  0,
				zr: 1,
			},
			{
				x: 512,
				y: 1024,
				r: 512,
			},
		})
	})

	t.Run("-x", func(t *testing.T) {
		alu := ALU{
			ZX: Inactive,
			NX: Inactive,
			ZY: Active,
			NY: Active,
			F:  Active,
			NO: Active,
		}
		runner(t, &alu, []assertion{
			{
				x:  0b0000_0000_1100_1010,
				r:  0b1111_1111_0011_0110,
				ng: 1,
			},
			{
				x:  0b0011_1011_1101_1011,
				y:  512,
				r:  0b1100_0100_0010_0101,
				ng: 1,
			},
			{
				x: 0b1100_0100_0010_0101,
				y: 512,
				r: 0b0011_1011_1101_1011,
			},
		})
	})

	t.Run("-y", func(t *testing.T) {
		alu := ALU{
			ZX: Active,
			NX: Active,
			ZY: Inactive,
			NY: Inactive,
			F:  Active,
			NO: Active,
		}
		runner(t, &alu, []assertion{
			{
				y:  0b0000_0000_1100_1010,
				r:  0b1111_1111_0011_0110,
				ng: 1,
			},
			{
				y:  0b0011_1011_1101_1011,
				x:  512,
				r:  0b1100_0100_0010_0101,
				ng: 1,
			},
			{
				y: 0b1100_0100_0010_0101,
				x: 512,
				r: 0b0011_1011_1101_1011,
			},
		})
	})

	t.Run("x & Y", func(t *testing.T) {
		alu := ALU{
			ZX: Inactive,
			NX: Inactive,
			ZY: Inactive,
			NY: Inactive,
			F:  Inactive,
			NO: Inactive,
		}
		runner(t, &alu, []assertion{
			{
				y: 0b0000_0000_1100_1010,
				x: 0b0111_0000_1000_1010,
				r: 0b0000_0000_1000_1010,
			},
		})
	})

	t.Run("x | y", func(t *testing.T) {
		alu := ALU{
			ZX: Inactive,
			NX: Active,
			ZY: Inactive,
			NY: Active,
			F:  Inactive,
			NO: Active,
		}
		runner(t, &alu, []assertion{
			{
				y: 0b0000_0000_1100_1010,
				x: 0b0111_0000_1000_1010,
				r: 0b0111_0000_1100_1010,
			},
		})
	})

	t.Run("!x", func(t *testing.T) {
		alu := ALU{
			ZX: Inactive,
			NX: Inactive,
			ZY: Active,
			NY: Active,
			F:  Inactive,
			NO: Active,
		}
		runner(t, &alu, []assertion{
			{
				x:  0b0111_0000_1000_1010,
				y:  0b0000_0000_1100_1010,
				r:  0b1000_1111_0111_0101,
				ng: 1,
			},
		})
	})

	t.Run("!y", func(t *testing.T) {
		alu := ALU{
			ZX: Active,
			NX: Active,
			ZY: Inactive,
			NY: Inactive,
			F:  Inactive,
			NO: Active,
		}
		runner(t, &alu, []assertion{
			{
				x:  0b0000_0000_1100_1010,
				y:  0b0111_0000_1000_1010,
				r:  0b1000_1111_0111_0101,
				ng: 1,
			},
		})
	})

	t.Run("x", func(t *testing.T) {
		alu := ALU{
			ZX: Inactive,
			NX: Inactive,
			ZY: Active,
			NY: Active,
			F:  Inactive,
			NO: Inactive,
		}
		runner(t, &alu, []assertion{
			{
				x:  45643,
				y:  3124,
				r:  45643,
				ng: 1,
			},
			{
				x: 21234,
				y: 0,
				r: 21234,
			},
		})
	})

	t.Run("y", func(t *testing.T) {
		alu := ALU{
			ZX: Active,
			NX: Active,
			ZY: Inactive,
			NY: Inactive,
			F:  Inactive,
			NO: Inactive,
		}
		runner(t, &alu, []assertion{
			{
				x:  3124,
				y:  45643,
				r:  45643,
				ng: 1,
			},
			{
				x: 0,
				y: 21234,
				r: 21234,
			},
		})
	})

	t.Run("-1", func(t *testing.T) {
		alu := ALU{
			ZX: Active,
			NX: Active,
			ZY: Active,
			NY: Inactive,
			F:  Active,
			NO: Inactive,
		}
		runner(t, &alu, []assertion{
			{
				x:  3124,
				y:  45643,
				r:  0b1111_1111_1111_1111,
				ng: 1,
			},
			{
				x:  0,
				y:  21234,
				r:  0b1111_1111_1111_1111,
				ng: 1,
			},
		})
	})

	t.Run("1", func(t *testing.T) {
		alu := ALU{
			ZX: Active,
			NX: Active,
			ZY: Active,
			NY: Active,
			F:  Active,
			NO: Active,
		}
		runner(t, &alu, []assertion{
			{
				x: 3124,
				y: 45643,
				r: 1,
			},
			{
				x: 0,
				y: 21234,
				r: 1,
			},
		})
	})

	t.Run("0", func(t *testing.T) {
		alu := ALU{
			ZX: Active,
			NX: Inactive,
			ZY: Active,
			NY: Inactive,
			F:  Active,
			NO: Inactive,
		}
		runner(t, &alu, []assertion{
			{
				x:  3124,
				y:  45643,
				r:  0,
				zr: 1,
			},
			{
				x:  0,
				y:  21234,
				r:  0,
				zr: 1,
			},
		})
	})
}
