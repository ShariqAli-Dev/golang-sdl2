package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	WindowWidth  = 800
	WindowHeight = 500
	WindowTitle  = "Changing Colors & Surface Icon"
)

type game struct {
	window          *sdl.Window
	renderer        *sdl.Renderer
	backgroundImage *sdl.Texture
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

	if err := game.loadMedia(); err != nil {
		log.Fatalf(error.Error(err))
	}

	game.run()
}

func sdlInit() error {
	var err error

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("error initializing sdl2: %v", err)
	}

	if err = img.Init(img.INIT_PNG); err != nil {
		return fmt.Errorf("error initializing sdl_image: %v", err)
	}
	return err
}

func sdlClose() {
	sdl.Quit()
}
