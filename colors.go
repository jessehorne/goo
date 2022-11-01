package gui

import (
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

var COLOR_CLEAR = RGBAToColor(50, 0, 0, 0)
var COLOR_BLACK = RGBAToColor(0, 0, 0, 0)
var COLOR_WHITE = RGBAToColor(255, 255, 255, 0)
var COLOR_WHITE_ALPHA = RGBAToColor(255, 255, 255, 30)
var COLOR_GREY = RGBAToColor(155, 155, 155, 0)
var COLOR_RED = RGBAToColor(255, 0, 0, 50)
var COLOR_GREEN = RGBAToColor(0, 255, 0, 0)
var COLOR_BLUE = RGBAToColor(0, 0, 255, 0)

var COLOR_PRIMARY_DARK = RGBAToColor(27, 47, 51, 0)
var COLOR_PRIMARY_LIGHT = RGBAToColor(57, 77, 81, 0)

var THEME_BLUE_90 = HexToColor("1A1A2E")
var THEME_BLUE_85 = HexToColor("27324F")
var THEME_BLUE_80 = HexToColor("16213E")
var THEME_BLUE_70 = HexToColor("003eaa")
var THEME_BLUE_60 = HexToColor("0060df")
var THEME_BLUE_50 = HexToColor("0a84ff")

var BUTTON_PRIMARY = HexToColor("5837D0")
var BUTTON_SECONDARY = HexToColor("7DE5ED")
var BUTTON_LIGHT = HexToColor("5DA7DB")
var BUTTON_LIGHTER = HexToColor("81C6E8")

// RGBAToColor returns the correct sdl.Color because it's mixed up somehow
func RGBAToColor(r uint8, g uint8, b uint8, a uint8) sdl.Color {
	return sdl.Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

// ColorToRGBA returns the correct RGBA values because it's mixed up somehow
func ColorToRGBA(c sdl.Color) (uint8, uint8, uint8, uint8) {
	return c.R, c.G, c.B, c.A
}

func HexToColor(hex string) sdl.Color {
	values, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return COLOR_WHITE
	}

	return sdl.Color{
		R: uint8(values >> 16),
		G: uint8((values >> 8) & 0xFF),
		B: uint8(values & 0xFF),
		A: uint8(0),
	}
}
