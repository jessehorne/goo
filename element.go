package gui

import "github.com/veandco/go-sdl2/sdl"

// Element represents all drawable objects. Other libraries call these "widgets"
type Element interface {
	Draw()
	Trigger()
	GetPos() (int32, int32)   // get position inside container
	SetPos(x int32, y int32)  // set position inside container
	GetSize() (int32, int32)  // get size of element
	SetSize(w int32, h int32) // set size of element
	SetHover(which bool)      // sets whether mouse is hovering over an element
	KeyboardEvent(keyType uint32, key sdl.Keysym)
	MouseButtonEvent(event *sdl.MouseButtonEvent)
}
