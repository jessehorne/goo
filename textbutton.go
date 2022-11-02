package gui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type TextButton struct {
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
	Text            string
	IsHovering      bool
	IsHolding       bool
	Callback        func(i *TextButton)
	FontStyle       string
	FontSize        int32
}

func NewTextButton(t string, x int32, y int32) *TextButton {
	SetFont("normal", 14)
	w, h := GetTextSize(t)

	paddingX := int32(5)
	paddingY := int32(0)

	width := int32(w) + (paddingX * 2)
	height := int32(h) + (paddingY * 2)

	return &TextButton{
		X:               x,
		Y:               y,
		Width:           width,
		Height:          height,
		BackgroundColor: BUTTON_PRIMARY,
		OutlineColor:    BUTTON_PRIMARY,
		MarginX:         0,
		MarginY:         0,
		PaddingX:        paddingX,
		PaddingY:        paddingY,
		Text:            t,
		FontStyle:       "normal",
		FontSize:        14,
	}
}

func (b *TextButton) SetFont(fontStyle string, fontSize int32) {
	b.FontStyle = fontStyle
	b.FontSize = fontSize

	// update width/height
	SetFont(fontStyle, fontSize)
	w, h := GetTextSize(b.Text)
	b.Width = int32(w) + b.PaddingX*2
	b.Height = int32(h) + b.PaddingY*2
}

func (b *TextButton) MouseButtonEvent(event *sdl.MouseButtonEvent) {
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

func (b *TextButton) SetHolding(which bool) {
	if which {
		b.BackgroundColor = BUTTON_LIGHTER
		b.OutlineColor = BUTTON_LIGHTER
	} else {
		b.BackgroundColor = BUTTON_PRIMARY
		b.OutlineColor = BUTTON_PRIMARY
	}

	b.IsHolding = which
}

func (b *TextButton) SetCallback(f func(i *TextButton)) {
	b.Callback = f
}

func (b *TextButton) Trigger() {
	if b.Callback != nil {
		b.Callback(b)
	}
}

func (b *TextButton) SetMargin(x int32, y int32) {
	b.MarginX = x
	b.MarginY = y
}

func (b *TextButton) SetPadding(x int32, y int32) {
	b.PaddingX = x
	b.PaddingY = y
}

func (b *TextButton) SetHover(which bool) {
	if which {
		b.BackgroundColor = BUTTON_LIGHT
		b.OutlineColor = BUTTON_LIGHT
	} else {
		b.BackgroundColor = BUTTON_PRIMARY
		b.OutlineColor = BUTTON_PRIMARY
	}

	b.IsHovering = which
}

func (b *TextButton) Draw() {
	originX := b.X + b.MarginX
	originY := b.Y + b.MarginY

	rect := sdl.Rect{X: originX, Y: originY, W: b.Width, H: b.Height}

	// Background
	SetColor(b.BackgroundColor)
	DrawFilledRect(&rect)

	// Outline
	SetColor(b.OutlineColor)
	DrawLineRect(&rect)

	// Text
	SetColor(COLOR_WHITE)
	SetFont(b.FontStyle, b.FontSize)
	textW, textH := GetTextSize(b.Text)
	textX := b.X + (b.Width / 2) - (int32(textW) / 2)
	textY := b.Y + (b.Height / 2) - (int32(textH) / 2)
	Print(b.Text, textX, textY)

}

func (b *TextButton) GetPos() (int32, int32) {
	return b.X, b.Y
}

func (b *TextButton) SetPos(x int32, y int32) {
	b.X = x
	b.Y = y
}

func (b *TextButton) GetSize() (int32, int32) {
	return b.Width, b.Height
}

func (b *TextButton) SetSize(w int32, h int32) {
	b.Width = w
	b.Height = h
}

func (b *TextButton) KeyboardEvent(keyType uint32, key sdl.Keysym) {
	return
}
