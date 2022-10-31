package gui

import "github.com/veandco/go-sdl2/sdl"

func BBox(r1 sdl.Rect, r2 sdl.Rect) bool {
	return r1.X < r2.X+r2.W && r2.X < r1.X+r1.W && r1.Y < r2.Y+r2.H && r2.Y < r1.Y+r1.H
}
