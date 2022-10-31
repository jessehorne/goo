package gui

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var A *App
var W *sdl.Window
var S *sdl.Surface
var R *sdl.Renderer

var Color = COLOR_WHITE // current draw color for many things such as primitives and font text

// Init sets up pointers to everything that will be used for drawing and so on
func Init(app *App) error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}

	window, err := sdl.CreateWindow(app.Title, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, app.ScreenWidth, app.ScreenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}

	surface, err := window.GetSurface()
	if err != nil {
		return err
	}

	r, err := window.GetRenderer()
	if err != nil {
		return err
	}

	err = ttf.Init()
	if err != nil {
		return err
	}

	// load fonts
	err = LoadFont("normal", "./resources/fonts/Trueno.otf", 24)
	if err != nil {
		panic(err)
	}

	err = LoadFont("light", "./resources/fonts/TruenoLight.otf", 24)
	if err != nil {
		panic(err)
	}

	err = LoadFont("light-button", "./resources/fonts/TruenoLight.otf", 18)
	if err != nil {
		panic(err)
	}

	err = LoadFont("normal-button", "./resources/fonts/Trueno.otf", 18)
	if err != nil {
		panic(err)
	}

	// topbar container
	err = SetFont("normal")
	if err != nil {
		panic(err)
	}

	A = app // gui now has access to the application
	W = window
	S = surface
	R = r

	return nil
}
