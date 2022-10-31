package gui

import "github.com/veandco/go-sdl2/ttf"

type Font struct {
	Font     *ttf.Font
	FontSize int32
}

var Fonts = make(map[string]*Font)
var CurrentFont = ""
