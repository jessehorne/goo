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

	if c.Alignment == ALIGN_CENTER {
		c.AlignElementsCenter()
	}
}

func (c *Container) AlignElementsCenter() {
	centerX := (c.X + c.Width) / 2
	centerY := (c.Y + c.Height) / 2
	elemCount := int32(len(c.Elements))

	totalWidth := elemCount * c.AlignMargin

	// TODO: figure out how to do this in one pass :)

	for _, e := range c.Elements {
		switch e.(type) {
		case *TextButton:
			btn := e.(*TextButton)
			totalWidth += btn.Width
		case *ImgButton:
			btn := e.(*ImgButton)
			totalWidth += btn.Width
		case *OneLineInput:
			i := e.(*OneLineInput)
			totalWidth += i.Width
		}
	}

	currentX := centerX - (totalWidth / 2)
	for _, e := range c.Elements {
		eX := currentX

		switch e.(type) {
		case *TextButton:
			btn := e.(*TextButton)
			btn.X = eX
			btn.Y = centerY - btn.Height/2
			currentX += btn.Width + c.AlignMargin
		case *ImgButton:
			btn := e.(*ImgButton)
			btn.X = eX
			btn.Y = centerY - btn.Height/2
			currentX += btn.Width + c.AlignMargin
		case *OneLineInput:
			i := e.(*OneLineInput)
			i.X = eX
			i.Y = centerY - i.Height/2
			currentX += i.Width + c.AlignMargin
		}
	}
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

func (c *Container) SetAlignment(a uint8) {
	c.Alignment = a
}
