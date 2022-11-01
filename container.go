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
	Alignment       uint8
	AlignMargin     int32
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
		AlignMargin:     10, // 10 pixels margin between elements
	}
}

func (c *Container) Draw() {
	rect := sdl.Rect{X: c.X, Y: c.Y, W: c.Width, H: c.Height}

	// Background color
	SetColor(c.BackgroundColor)
	DrawFilledRect(&rect)

	// Outline color
	if c.OutlineColor != COLOR_CLEAR {
		SetColor(c.OutlineColor)
		DrawLineRect(&rect)
	}

	// Draw and clip elements
	for _, e := range c.Elements {
		R.SetClipRect(&rect)
		e.(Element).Draw()
		ResetClipRect()
	}
}

func (c *Container) SetBackgroundColor(newColor sdl.Color) {
	c.BackgroundColor = newColor
}

func (c *Container) SetOutlineColor(newColor sdl.Color) {
	c.OutlineColor = newColor
}

func (c *Container) AddElement(e interface{}) {
	c.Elements = append(c.Elements, e)

	if c.Alignment == ALIGN_CENTER {
		c.AlignElementsCenter()
	}
}

func (c *Container) AlignElementsCenter() {
	centerX := c.X + c.Width/2
	centerY := c.Y + c.Height/2
	elemCount := int32(len(c.Elements))

	totalWidth := elemCount * c.AlignMargin

	// TODO: figure out how to do this in one pass :)

	for _, e := range c.Elements {
		w, _ := e.(Element).GetSize()
		totalWidth += w
	}

	currentX := centerX - (totalWidth / 2)
	for _, e := range c.Elements {
		eX := currentX
		elemW, elemH := e.(Element).GetSize()
		newX := eX
		newY := centerY - elemH/2
		currentX += elemW + c.AlignMargin
		e.(Element).SetPos(newX, newY)
	}
}

func (c *Container) ResetHoverStates() {
	for _, e := range c.Elements {
		e.(Element).SetHover(false)
	}
}

func (c *Container) SetAlignment(a uint8) {
	c.Alignment = a
}
