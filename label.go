package gui

import "github.com/veandco/go-sdl2/sdl"

type Label struct {
	Text     string
	X        int32
	Y        int32
	Width    int32
	Height   int32
	FontSize uint16
	Color    sdl.Color
}

func NewLabel(text string, fontSize uint16, x int32, y int32) *Label {
	SetFont("normal")
	w, h := GetTextSize(text)

	return &Label{
		Text:     text,
		X:        x,
		Y:        y,
		Width:    int32(w),
		Height:   int32(h),
		FontSize: fontSize,
		Color:    COLOR_WHITE,
	}
}

func (l *Label) Draw() {
	SetFont("normal")
	SetColor(l.Color)
	Print(l.Text, l.X, l.Y)
}

func (l *Label) Trigger() {
	return
}
