package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func newGame() *game {
	game := game{
		fontSize:     80,
		fontColor:    sdl.Color{R: 255, B: 255, G: 255},
		textVelocity: 3,
		spriteVel:    5,
	}
	game.textXVelocity = game.textVelocity
	game.textYVelocity = game.textVelocity
	return &game
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

	iconSurface, err := img.Load("./images/Go-logo.png")
	if err != nil {
		return fmt.Errorf("error loading the surface: %v", err)
	}
	defer iconSurface.Free()
	g.window.SetIcon(iconSurface)

	g.keystate = sdl.GetKeyboardState()

	return err
}

func (g *game) close() {
	if g != nil {
		mix.HaltMusic()
		mix.HaltChannel(-1)

		g.music.Free()
		g.music = nil

		g.sdlSound.Free()
		g.sdlSound = nil

		g.goSound.Free()
		g.goSound = nil

		g.spriteImage.Destroy()
		g.spriteImage = nil

		g.textImage.Destroy()
		g.textImage = nil

		g.backgroundImage.Destroy()
		g.backgroundImage = nil

		g.renderer.Destroy()
		g.renderer = nil

		g.window.Destroy()
		g.window = nil
	}
}

func (g *game) run() {
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.KeyboardEvent:
				if e.Type == sdl.KEYDOWN {
					switch e.Keysym.Scancode {
					case sdl.SCANCODE_ESCAPE:
						return
					case sdl.SCANCODE_SPACE:
						g.randColor()
					case sdl.SCANCODE_P:
						g.pauseMusic()
					}
				}
			}
		}

		g.updateText()
		g.updateSprite()

		g.renderer.Clear()

		g.renderer.Copy(g.backgroundImage, nil, nil)
		g.renderer.Copy(g.textImage, nil, &g.textRectangle)
		g.renderer.Copy(g.spriteImage, nil, &g.spriteRectangle)

		g.renderer.Present()
		sdl.Delay(uint32(1000 / 60))
	}
}

func (g *game) loadMedia() error {
	var err error

	g.backgroundImage, err = img.LoadTexture(g.renderer, "./images/background.png")
	if err != nil {
		return fmt.Errorf("error leading background texture: %v", err)
	}

	font, err := ttf.OpenFont("./fonts/freesansbold.ttf", g.fontSize)
	if err != nil {
		return fmt.Errorf("error opening font: %v", err)
	}
	defer font.Close()

	fontSurf, err := font.RenderUTF8Blended("SDL", g.fontColor)
	if err != nil {
		return fmt.Errorf("error creating text surface: %v", err)
	}
	defer fontSurf.Free()

	g.textRectangle.W = fontSurf.W
	g.textRectangle.H = fontSurf.H

	g.textImage, err = g.renderer.CreateTextureFromSurface(fontSurf)
	if err != nil {
		return fmt.Errorf("error creating texture from surface: %v", err)
	}

	g.spriteImage, err = img.LoadTexture(g.renderer, "./images/Go-logo.png")
	if err != nil {
		return fmt.Errorf("error loading texture: %v", err)
	}
	_, _, g.spriteRectangle.W, g.spriteRectangle.H, err = g.spriteImage.Query()
	if err != nil {
		return fmt.Errorf("error querying texture: %v", err)
	}

	g.goSound, err = mix.LoadWAV("./sounds/Go.ogg")
	if err != nil {
		return fmt.Errorf("error loading goSound: %v", err)
	}

	g.sdlSound, err = mix.LoadWAV("./sounds/SDL.ogg")
	if err != nil {
		return fmt.Errorf("error loading sdlSound: %v", err)
	}

	g.music, err = mix.LoadMUS("./music/freesoftwaresong-8bit.ogg")
	if err != nil {
		return fmt.Errorf("error loading music: %v", err)
	}

	if err = g.music.Play(-1); err != nil {
		return fmt.Errorf("error playing music: %v", err)
	}

	return err
}

func (g *game) randColor() {
	g.renderer.SetDrawColor(uint8(rand.IntN(256)), uint8(rand.IntN(256)), uint8(rand.IntN(256)), 255)
	g.goSound.Play(-1, 0)
}

func (g *game) updateText() {
	g.textRectangle.X += g.textXVelocity
	g.textRectangle.Y += g.textYVelocity

	if g.textRectangle.X < 0 {
		g.textXVelocity = g.textVelocity
		g.sdlSound.Play(-1, 0)
	} else if (g.textRectangle.X + g.textRectangle.W) > WindowWidth {
		g.textXVelocity = -g.textVelocity
		g.sdlSound.Play(-1, 0)
	}

	if g.textRectangle.Y < 0 {
		g.textYVelocity = g.textVelocity
		g.sdlSound.Play(-1, 0)
	} else if (g.textRectangle.Y + g.textRectangle.H) > WindowHeight {
		g.textYVelocity = -g.textVelocity
		g.sdlSound.Play(-1, 0)
	}
}

func (g *game) updateSprite() {
	if g.keystate[sdl.SCANCODE_LEFT] == 1 || g.keystate[sdl.SCANCODE_A] == 1 {
		g.spriteRectangle.X -= g.spriteVel
	}
	if g.keystate[sdl.SCANCODE_RIGHT] == 1 || g.keystate[sdl.SCANCODE_D] == 1 {
		g.spriteRectangle.X += g.spriteVel
	}
	if g.keystate[sdl.SCANCODE_UP] == 1 || g.keystate[sdl.SCANCODE_W] == 1 {
		g.spriteRectangle.Y -= g.spriteVel
	}
	if g.keystate[sdl.SCANCODE_DOWN] == 1 || g.keystate[sdl.SCANCODE_S] == 1 {
		g.spriteRectangle.Y += g.spriteVel
	}
}

func (g *game) pauseMusic() {
	if mix.PausedMusic() {
		mix.ResumeMusic()
	} else {
		mix.PauseMusic()
	}
}
