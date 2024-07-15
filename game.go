package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func newGame() *game {
	return &game{}
}

func (g *game) init() error {
	var err error

	g.window, err = sdl.CreateWindow(WindowTitle, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, WindowWidth, WindowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("error creating window: %v", err)
	}

	g.renderer, err = sdl.CreateRenderer(g.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return fmt.Errorf("error creating renderer: %v", err)
	}

	return err
}

func (g *game) close() {
	if g != nil {
		g.renderer.Destroy()
		g.renderer = nil

		g.window.Destroy()
		g.window = nil
	}
}

func (g *game) run() {
	g.renderer.Clear()
	g.renderer.Present()
	sdl.Delay(5000)
}
