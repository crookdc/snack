package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/crookdc/nand2tetris/internal/chip"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
)

const (
	ScreenMemoryMapBegin  = 16_384
	ScreenMemoryMapLength = 8192
	ScreenRefreshRateHz   = 33
)

var (
	profile = flag.String("profile", "", "write profiling data to files with this base name")
	program = flag.String("program", "", "file containing program to be written to rom")
)

func main() {
	flag.Parse()
	if *profile != "" {
		f, err := os.Create(*profile + ".cpu")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

	if *program == "" {
		log.Fatal("missing path to program")
	}
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal(err)
	}
	defer sdl.Quit()

	screen, err := NewSDLScreen()
	if err != nil {
		log.Fatal(err)
	}
	defer screen.Close()
	if err := screen.Clear(); err != nil {
		log.Fatal(err)
	}
	screen.renderer.Present()

	prog, err := loadProgram(*program)
	if err != nil {
		log.Fatal(err)
	}
	ram := &RAM{window: screen}
	computer := chip.NewComputer(
		prog,
		ram,
	)
	running := true
	renderTick := sdl.GetTicks64()
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}
		computer.Tick(chip.Inactive)
		if sdl.GetTicks64()-renderTick > 1000/ScreenRefreshRateHz {
			if err := screen.Draw(&ram.mem); err != nil {
				log.Fatal(err)
			}
			renderTick = sdl.GetTicks64()
		}
	}

	if *profile != "" {
		f, err := os.Create(*profile + ".heap")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal(err)
		}
	}
}

type RAM struct {
	mem    [32768][16]chip.Signal
	window SDLScreen
}

func (b *RAM) Out(load chip.Signal, addr [15]chip.Signal, in chip.ReadonlyWord) *chip.Word {
	idx := chip.Join15(addr)
	if load == chip.Inactive {
		return chip.Wrap(&b.mem[idx])
	}
	b.mem[idx] = in.Copy()
	return chip.Wrap(&b.mem[idx])
}

type ROM [][16]chip.Signal

func (r ROM) Out(_ chip.Signal, addr [15]chip.Signal, _ chip.ReadonlyWord) *chip.Word {
	return chip.Wrap(&r[chip.Join15(addr)])
}

func NewSDLScreen() (SDLScreen, error) {
	window, err := sdl.CreateWindow("Hack", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 512, 256, sdl.WINDOW_SHOWN)
	if err != nil {
		return SDLScreen{}, err
	}
	renderer, err := sdl.CreateRenderer(window, 0, 0)
	if err != nil {
		return SDLScreen{}, err
	}
	return SDLScreen{
		window:   window,
		renderer: renderer,
	}, nil
}

type SDLScreen struct {
	window    *sdl.Window
	renderer  *sdl.Renderer
	presented uint64
}

func (s *SDLScreen) Clear() error {
	if err := s.renderer.SetDrawColor(0, 0, 0, 255); err != nil {
		return err
	}
	if err := s.renderer.Clear(); err != nil {
		return err
	}
	return nil
}

func (s *SDLScreen) Draw(mem *[32768][16]chip.Signal) error {
	if err := s.Clear(); err != nil {
		return err
	}
	if err := s.renderer.SetDrawColor(255, 255, 255, 255); err != nil {
		return err
	}
	points := make([]sdl.Point, 0, ScreenMemoryMapLength)
	for i := range ScreenMemoryMapLength {
		points = append(points, s.points(i, &mem[ScreenMemoryMapBegin+i])...)
	}
	if len(points) == 0 {
		// If there are no points to render then the renderer will return an error in DrawPoints. Even if that was not
		// the case then it would just be wasteful to call the renderer if there is nothing to render.
		return nil
	}
	if err := s.renderer.DrawPoints(points); err != nil {
		return err
	}
	s.renderer.Present()
	return nil
}

func (s *SDLScreen) points(position int, val *[16]chip.Signal) []sdl.Point {
	points := make([]sdl.Point, 0, 16)
	row := position / 32
	for i, px := range val {
		col := ((position * 16) % 512) + i
		if px == chip.Inactive {
			continue
		}
		points = append(points, sdl.Point{
			X: int32(col),
			Y: int32(row),
		})
	}
	return points
}

func (s *SDLScreen) Close() {
	_ = s.window.Destroy()
}

func loadProgram(file string) (ROM, error) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var rom ROM
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		if len(line) != 16 {
			return nil, fmt.Errorf("invalid line length '%s'", line)
		}
		instruction, err := parseInstruction(line)
		if err != nil {
			return nil, err
		}
		rom = append(rom, instruction)
	}
	return rom, nil
}

func parseInstruction(line string) ([16]chip.Signal, error) {
	instruction := [16]chip.Signal{}
	for i := range 16 {
		bit, err := strconv.Atoi(string(line[i]))
		if err != nil {
			return [16]chip.Signal{}, err
		}
		instruction[i] = chip.Signal(bit)
	}
	return instruction, nil
}
