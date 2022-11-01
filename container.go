package gui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Container struct {
	X               int32
	Y               int32
	Width           int32
	Height          int32
	BackgroundColor sdl.Color
	OutlineColor    sdl.Color
	Elements        []interface{}
	ElementOffsetX  int32
	Alignment       int8
}

func NewContainer(x int32, y int32, w int32, h int32) *Container {
	return &Container{
		X:               x,
		Y:               y,
		Width:           w,
		Height:          h,
		BackgroundColor: COLOR_CLEAR,
		OutlineColor:    COLOR_CLEAR,
		Alignment:       ALIGN_FREE,
	}
}

func (c *Container) Draw() {
	rect := sdl.Rect{X: c.X, Y: c.Y, W: c.Width, H: c.Height}

	// Background color
	SetColor(c.BackgroundColor)
	DrawFilledRect(&rect)

	// Outline color
	SetColor(c.OutlineColor)
	DrawLineRect(&rect)

	// Draw and clip elements
	R.SetClipRect(&rect)
	for _, e := range c.Elements {
		e.(Element).Draw()
	}
	ResetClipRect()
}

func (c *Container) SetBackgroundColor(newColor sdl.Color) {
	c.BackgroundColor = newColor
}

func (c *Container) SetOutlineColor(newColor sdl.Color) {
	c.OutlineColor = newColor
}

func (c *Container) AddElement(e interface{}) {
	c.Elements = append(c.Elements, e)

	// (re)align elements
	//centerX := (c.X + c.Width) / 2
	//centerY := (c.Y + c.Height) / 2
}

func (c *Container) ResetHoverStates() {
	for _, e := range c.Elements {
		switch e.(type) {
		case *TextButton:
			btn := e.(*TextButton)
			btn.SetHover(false)
		case *ImgButton:
			btn := e.(*ImgButton)
			btn.SetHover(false)
		case *OneLineInput:
			i := e.(*OneLineInput)
			i.SetHover(false)
		}
	}
}
