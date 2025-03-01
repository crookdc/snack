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
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	program    = flag.String("program", "", "file containing program to be written to rom")
)

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
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
	if err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_EVENTS); err != nil {
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
	computer := chip.NewComputer(prog)
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}
		wmem, maddr, omem := computer.Tick(chip.Inactive)
		mem := chip.Join15(maddr)
		if wmem == chip.Active && mem >= ScreenMemoryMapBegin && mem < ScreenMemoryMapBegin+ScreenMemoryMapLength {
			if err := screen.Draw(mem-ScreenMemoryMapBegin, omem); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func NewSDLScreen() (SDLScreen, error) {
	window, err := sdl.CreateWindow("Hack", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 256, 512, sdl.WINDOW_SHOWN)
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
	window   *sdl.Window
	renderer *sdl.Renderer
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

func (s *SDLScreen) Draw(position uint16, val [16]chip.Signal) error {
	row := position / 16
	for i, px := range val {
		var err error
		if px == chip.Active {
			err = s.renderer.SetDrawColor(255, 255, 255, 255)
		} else {
			err = s.renderer.SetDrawColor(0, 0, 0, 255)
		}
		if err != nil {
			return err
		}
		col := ((int(position) * 16) % 256) + i
		err = s.renderer.DrawPoint(int32(col), int32(row))
		if err != nil {
			return err
		}
	}
	s.renderer.Present()
	return nil
}

func (s *SDLScreen) Close() {
	_ = s.window.Destroy()
}

func loadProgram(file string) ([][16]chip.Signal, error) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var program [][16]chip.Signal
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
		program = append(program, instruction)
	}
	return program, nil
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
