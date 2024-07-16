package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	WindowWidth  = 800
	WindowHeight = 500
	WindowTitle  = "Go Sound, Sdl Sound, & Music"
)

type game struct {
	window          *sdl.Window
	renderer        *sdl.Renderer
	backgroundImage *sdl.Texture

	fontSize      int
	fontColor     sdl.Color
	textImage     *sdl.Texture
	textVelocity  int32
	textXVelocity int32
	textYVelocity int32
	textRectangle sdl.Rect

	spriteImage     *sdl.Texture
	spriteRectangle sdl.Rect
	spriteVel       int32

	keystate []uint8
	goSound  *mix.Chunk
	sdlSound *mix.Chunk
	music    *mix.Music
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
	// can pipe in other flags
	// var sdlFlags uint32 = sdl.INIT_EVENTS | sdl.INIT_something_else

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("error initializing sdl2: %v", err)
	}

	if err = img.Init(img.INIT_PNG); err != nil {
		return fmt.Errorf("error initializing sdl_image: %v", err)
	}

	if err = ttf.Init(); err != nil {
		return fmt.Errorf("error initializing ttf: %v", err)
	}

	if err = mix.Init(mix.INIT_OGG); err != nil {
		return fmt.Errorf("error initializing sdl mixer: %v", err)
	}

	if err = mix.OpenAudio(mix.DEFAULT_FREQUENCY, mix.DEFAULT_FORMAT, mix.DEFAULT_CHANNELS, mix.DEFAULT_CHUNKSIZE); err != nil {
		return fmt.Errorf("error opening audio: %v", err)
	}

	return err
}

func sdlClose() {
	mix.CloseAudio()
	mix.Quit()
	ttf.Quit()
	img.Quit()
	sdl.Quit()
}
