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
}

func NewContainer(x int32, y int32, w int32, h int32, bgColor sdl.Color) *Container {
	return &Container{
		X:               x,
		Y:               y,
		Width:           w,
		Height:          h,
		BackgroundColor: bgColor,
		OutlineColor:    bgColor,
	}
}

func (c *Container) Draw() {
	rect := sdl.Rect{X: c.X, Y: c.Y, W: c.Width, H: c.Height}

	// Background Color
	SetColor(c.BackgroundColor)
	DrawFilledRect(&rect)

	// Outline
	SetColor(c.OutlineColor)
	DrawLineRect(&rect)

	// Draw Elements
	for _, e := range c.Elements {
		e.(Element).Draw()
	}
}

func (c *Container) SetOutlineColor(newColor sdl.Color) {
	c.OutlineColor = newColor
}

func (c *Container) AddElement(e interface{}) {
	c.Elements = append(c.Elements, e)
}
