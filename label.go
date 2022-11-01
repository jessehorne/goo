package gui

import "github.com/veandco/go-sdl2/sdl"

type Label struct {
	Text       string
	X          int32
	Y          int32
	Width      int32
	Height     int32
	FontSize   uint16
	Color      sdl.Color
	IsHovering bool
}

func NewLabel(text string, fontSize uint16, x int32, y int32) *Label {
	SetFont("normal", 14)
	w, h := GetTextSize(text)

	return &Label{
		Text:       text,
		X:          x,
		Y:          y,
		Width:      int32(w),
		Height:     int32(h),
		FontSize:   fontSize,
		Color:      COLOR_WHITE,
		IsHovering: false,
	}
}

func (l *Label) Draw() {
	SetFont("normal", 14)
	SetColor(l.Color)
	Print(l.Text, l.X, l.Y)
}

func (l *Label) Trigger() {
	return
}

func (l *Label) GetPos() (int32, int32) {
	return l.X, l.Y
}

func (l *Label) SetPos(x int32, y int32) {
	l.X = x
	l.Y = y
}

func (l *Label) GetSize() (int32, int32) {
	return l.Width, l.Height
}

func (l *Label) SetSize(w int32, h int32) {
	l.Width = w
	l.Height = h
}

func (l *Label) SetHover(which bool) {
	l.IsHovering = which
}

func (l *Label) KeyboardEvent(keyType uint32, key sdl.Keysym) {
	return
}

func (l *Label) MouseButtonEvent(event *sdl.MouseButtonEvent) {
	return
}
