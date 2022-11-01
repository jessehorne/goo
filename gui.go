package gui

import (
	"github.com/veandco/go-sdl2/sdl"
)

var A *App
var W *sdl.Window
var S *sdl.Surface
var R *sdl.Renderer

var Color = COLOR_WHITE // current draw color for many things such as primitives and font text

const ALIGN_CENTER = 0 // margin, padding and center alignment
const ALIGN_FREE = 1   // no automated position adjustment

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

	err = InitFonts()
	if err != nil {
		panic(err)
	}

	err = SetFont("normal", 14)
	if err != nil {
		panic(err)
	}

	A = app // gui now has access to the application
	W = window
	S = surface
	R = r

	return nil
}
