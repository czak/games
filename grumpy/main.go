package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

var running = true
var flying = false
var paused = false

var bird []*sdl.Texture
var bg *sdl.Texture
var wood *sdl.Texture

var a float64

func run() int {
	win, renderer, err := initGraphics()
	if err != nil {
		log.Fatalf("Init graphics failed: %s", err)
	}
	defer win.Destroy()
	defer renderer.Destroy()

	// TODO: Check for errors
	bird, _ = loadSprite(renderer, "assets/bird%d.png", 4)
	bg, _ = img.LoadTexture(renderer, "assets/background.png")
	wood, _ = img.LoadTexture(renderer, "assets/wood.png")

	var x int32
	y := 50.0
	v := 0.0
	g := 0.2

	obstacles := []int32{}
	for i := 0; i < 100; i++ {
		obstacles = append(obstacles, rand.Int31n(200)+300)
	}

	frame := 0
	for running {
		handleEvents()
		if paused {
			continue
		}

		render(renderer, frame, x, int32(y), obstacles)

		// Physics
		if flying {
			v = -5.5
		} else {
			v += a + g
		}
		y += v

		if v < 0 {
			frame++
		}

		// Screen boundaries
		if y < 0 {
			y = 0
			v = 0
		} else if y >= 600-54 {
			y = 600 - 54
			v = 0
		}

		// Collisions
		index := (x + 90) / 150
		// fmt.Printf("Current obstacle: %d: %d\n", index, obstacles[index])

		if (x-90)%150 > 0 && (x-90)%150 < 56 {
			if int32(y)+54 >= obstacles[index] || int32(y) <= obstacles[index]-200 {
				// Crash
				renderer.SetDrawColor(255, 0, 0, 255)
				paused = true
			} else {
				renderer.SetDrawColor(255, 255, 0, 255)
			}
		} else {
			renderer.SetDrawColor(0, 255, 0, 255)
		}

		x++
		sdl.Delay(16)
	}

	sdl.Quit()
	return 0
}

func render(renderer *sdl.Renderer, frame int, x, y int32, obstacles []int32) {
	renderer.Copy(bg, nil, &sdl.Rect{X: 0 - (x % 1200), Y: 0, W: 1200, H: 600})
	renderer.Copy(bg, nil, &sdl.Rect{X: 1199 - (x % 1200), Y: 0, W: 1200, H: 600})

	// if (x-90)%150 > 0 && (x-90)%150 < 56 {
	// 	renderer.SetDrawColor(255, 255, 0, 255)
	// } else {
	// 	renderer.SetDrawColor(0, 0, 0, 255)
	// }

	for i, o := range obstacles {
		// renderer.FillRect(&sdl.Rect{X: int32(i)*300 - x*2, Y: o, W: 48, H: 600})
		renderer.Copy(wood, nil, &sdl.Rect{X: int32(i)*300 - x*2, Y: o, W: 48, H: 600})
		renderer.CopyEx(wood, nil, &sdl.Rect{X: int32(i)*300 - x*2, Y: o - 800, W: 48, H: 600}, 0, nil, sdl.FLIP_VERTICAL)
	}

	// renderer.FillRect(&sdl.Rect{X: 50, Y: int32(y), W: 64, H: 54})
	renderer.Copy(bird[frame/2%len(bird)], nil, &sdl.Rect{X: 50, Y: int32(y), W: 64, H: 54})

	renderer.Present()
}

func loadSprite(renderer *sdl.Renderer, pattern string, numFrames int) ([]*sdl.Texture, error) {
	var frames []*sdl.Texture

	for i := 1; i <= numFrames; i++ {
		path := fmt.Sprintf(pattern, i)
		frame, err := img.LoadTexture(renderer, path)
		if err != nil {
			return nil, fmt.Errorf("Unable to load frame: %s", err)
		}
		frames = append(frames, frame)
	}

	return frames, nil
}

func handleEvents() {
	keystate := sdl.GetKeyboardState()

	if keystate[sdl.SCANCODE_SPACE] != 0 || keystate[sdl.SCANCODE_UP] != 0 || keystate[sdl.SCANCODE_K] != 0 {
		flying = true
	} else {
		flying = false
	}

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			running = false
		case *sdl.KeyDownEvent:
			switch t.Keysym.Sym {
			case sdl.K_ESCAPE:
				running = false
			case sdl.K_p:
				paused = !paused
			}
		}
	}
}

func main() {
	os.Exit(run())
}
