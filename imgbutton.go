package gui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type ImgButton struct {
	X               int32
	Y               int32
	Width           int32
	Height          int32
	BackgroundColor sdl.Color
	OutlineColor    sdl.Color
	MarginX         int32
	MarginY         int32
	PaddingX        int32
	PaddingY        int32
	IsHovering      bool
	IsHolding       bool
	Image           *Image
	Callback        func(b *ImgButton)
}

func NewImgButton(path string, x int32, y int32) (*ImgButton, error) {
	img, err := NewImageFromSVG(path, 0, 0, 24, 24)
	if err != nil {
		return nil, err
	}

	return &ImgButton{
		X:               x,
		Y:               y,
		Width:           36,
		Height:          36,
		BackgroundColor: THEME_BLUE_90,
		OutlineColor:    THEME_BLUE_90,
		MarginX:         0,
		MarginY:         0,
		PaddingX:        int32((36 - 24) / 2),
		PaddingY:        int32((36 - 24) / 2),
		Image:           img,
	}, nil
}

func (b *ImgButton) MouseButtonEvent(event *sdl.MouseButtonEvent) {
	if b.IsHovering {
		// 0 here means UP. idk what it is set to in the library
		if event.Button == sdl.BUTTON_LEFT {
			if event.State == 0 {
				// MOUSE UP
				b.Trigger()
				b.SetHolding(false)
			} else if event.State == 1 {
				// MOUSE DOWN
				b.SetHolding(true)
			}
		}
	}
}

func (b *ImgButton) SetHolding(which bool) {
	if which {
		b.BackgroundColor = BUTTON_LIGHTER
	} else {
		b.BackgroundColor = BUTTON_PRIMARY
	}

	b.IsHolding = which
}

func (b *ImgButton) SetCallback(f func(i *ImgButton)) {
	b.Callback = f
}

func (b *ImgButton) Trigger() {
	if b.Callback != nil {
		b.Callback(b)
	}
}

func (b *ImgButton) SetMargin(x int32, y int32) {
	b.MarginX = x
	b.MarginY = y
}

func (b *ImgButton) SetPadding(x int32, y int32) {
	b.PaddingX = x
	b.PaddingY = y
}

func (b *ImgButton) SetHover(which bool) {
	if which {
		b.BackgroundColor = BUTTON_LIGHT
		b.OutlineColor = BUTTON_LIGHT
	} else {
		b.BackgroundColor = THEME_BLUE_90
		b.OutlineColor = THEME_BLUE_90
	}

	b.IsHovering = which
}

func (b *ImgButton) Draw() {
	originX := b.X + b.MarginX
	originY := b.Y + b.MarginY

	rect := sdl.Rect{X: originX, Y: originY, W: b.Width, H: b.Height}

	// Background
	SetColor(b.BackgroundColor)
	DrawFilledRect(&rect)

	// Outline
	SetColor(b.OutlineColor)
	DrawLineRect(&rect)

	// Image
	SetColor(COLOR_WHITE)
	destX := originX + b.PaddingX
	destY := originY + b.PaddingY
	DrawImage(b.Image, destX, destY)

}

func (b *ImgButton) GetPos() (int32, int32) {
	return b.X, b.Y
}

func (b *ImgButton) SetPos(x int32, y int32) {
	b.X = x
	b.Y = y
}

func (b *ImgButton) GetSize() (int32, int32) {
	return b.Width, b.Height
}

func (b *ImgButton) SetSize(w int32, h int32) {
	b.Width = w
	b.Height = h
}

func (b *ImgButton) KeyboardEvent(keyType uint32, key sdl.Keysym) {
	return
}
