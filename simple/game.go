package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func initGraphics() (*sdl.Window, *sdl.Renderer, error) {
	sdl.Init(sdl.INIT_EVERYTHING)

	win, err := sdl.CreateWindow("Square Bird", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to create window: %s", err)
	}

	renderer, err := sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to create renderer: %s", err)
	}

	return win, renderer, nil
}
