package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	WindowWidth  = 800
	WindowHeight = 500
	WindowTitle  = "Game"
)

type game struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

func main() {
	defer sdlClose()
	if err := sdlInit(); err != nil {
		log.Fatalf(error.Error(err))
	}

	game := newGame()
	defer game.close()
	if err := game.init(); err != nil {
		log.Fatalf(error.Error(err))
	}

	game.run()
}

func sdlInit() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("error initializing sdl2: %v", err)
	}
	return nil
}

func sdlClose() {
	sdl.Quit()
}
