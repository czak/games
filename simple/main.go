package main

import (
	"fmt"
	"log"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

// Sprite is the main game object
type Sprite struct {
	sdl.Rect
	boxes []sdl.Rect
}

var running = true

var bird = Sprite{
	Rect: sdl.Rect{X: 100, Y: 100, W: 64, H: 54},
	boxes: []sdl.Rect{
		sdl.Rect{X: 37, Y: 1, W: 13, H: 5},
		sdl.Rect{X: 14, Y: 6, W: 44, H: 11},
		sdl.Rect{X: 2, Y: 17, W: 59, H: 18},
		sdl.Rect{X: 15, Y: 35, W: 47, H: 7},
		sdl.Rect{X: 23, Y: 42, W: 28, H: 4},
	},
}

func (s *Sprite) hasIntersection(rect *sdl.Rect) bool {
	// move rect to s coordinates
	rect = &sdl.Rect{
		X: rect.X - s.X,
		Y: rect.Y - s.Y,
		W: rect.W,
		H: rect.H,
	}

	for _, box := range s.boxes {
		if box.HasIntersection(rect) {
			return true
		}
	}
	return false
}

var obstacle = sdl.Rect{X: 180, Y: 150, W: 50, H: 150}

func run() int {
	win, renderer, err := initGraphics()
	if err != nil {
		log.Fatalf("Init graphics failed: %s", err)
	}
	defer win.Destroy()
	defer renderer.Destroy()

	texture, _ := img.LoadTexture(renderer, "assets/bird1.png")

	freq := float64(sdl.GetPerformanceFrequency())

	var ticks uint64

	for running {
		ticks = sdl.GetPerformanceCounter()

		handleEvents()
		// animate()

		// Render
		renderer.SetDrawColor(190, 200, 250, 255)
		renderer.SetDrawBlendMode(sdl.BLENDMODE_NONE)
		renderer.Clear()

		if bird.hasIntersection(&obstacle) {
			renderer.SetDrawColor(255, 0, 0, 255)
		} else {
			renderer.SetDrawColor(255, 255, 255, 255)
		}

		// Render obstacle
		renderer.FillRect(&obstacle)

		// Render player
		renderer.Copy(texture, nil, &bird.Rect)
		// ...with bounding boxes
		/*
			renderer.SetDrawColor(100, 255, 100, 255)
			renderer.SetDrawBlendMode(sdl.BLENDMODE_ADD)
			for _, box := range bird.boxes {
				box = sdl.Rect{X: box.X + bird.X, Y: box.Y + bird.Y, W: box.W, H: box.H}
				renderer.FillRect(&box)
			}
		*/

		renderer.Present()

		ticks = sdl.GetPerformanceCounter() - ticks
		fmt.Println(float64(ticks) / freq * 1000.0)
	}

	sdl.Quit()
	return 0
}

func handleEvents() {
	keystate := sdl.GetKeyboardState()

	if keystate[sdl.SCANCODE_LEFT] != 0 || keystate[sdl.SCANCODE_H] != 0 {
		bird.X--
	} else if keystate[sdl.SCANCODE_RIGHT] != 0 || keystate[sdl.SCANCODE_L] != 0 {
		bird.X++
	}

	if keystate[sdl.SCANCODE_UP] != 0 || keystate[sdl.SCANCODE_K] != 0 {
		bird.Y--
	} else if keystate[sdl.SCANCODE_DOWN] != 0 || keystate[sdl.SCANCODE_J] != 0 {
		bird.Y++
	}

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			running = false
		case *sdl.KeyDownEvent:
			switch t.Keysym.Sym {
			case sdl.K_ESCAPE:
				running = false
			}
		}
	}
}

func main() {
	os.Exit(run())
}
