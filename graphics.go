package gui

import (
	"github.com/veandco/go-sdl2/sdl"
)

var CharsTypeable = "ABCDEFGHIJKLMNOPQRSTUVWXYZ`1234567890-=[]\\;',./"
var CharsLower = "abcdefghijklmnopqrstuvwxyz`1234567890-=[]\\;',./"
var CharsUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+{}|:\"<>?"

func RedrawAllContainers(cs []*Container) {
	r, g, b, a := ColorToRGBA(A.BackgroundColor)
	R.SetDrawColor(r, g, b, a)
	R.Clear()
	for _, c := range cs {
		c.Draw()
	}
	R.Present()
}

func RedrawContainer(c *Container) {
	r, g, b, a := ColorToRGBA(A.BackgroundColor)
	R.SetDrawColor(r, g, b, a)
	R.Clear()
	c.Draw()
	R.Present()
}

func ClearScreen() {
	r, g, b, a := ColorToRGBA(THEME_BLUE_90)
	R.SetDrawColor(r, g, b, a)
	R.FillRect(nil)
}

func SetColor(color sdl.Color) {
	r, g, b, a := ColorToRGBA(color)
	R.SetDrawColor(r, g, b, a)
	Color = color
}

func DrawFilledRect(rect *sdl.Rect) {
	R.FillRect(rect)
}

func DrawLineRect(rect *sdl.Rect) {
	R.DrawRect(rect)
}

func Print(s string, x int32, y int32) error {
	f := CurrentFont()
	msg, err := f.Font.RenderUTF8Blended(s, Color)
	if err != nil {
		return err
	}

	texture, err := R.CreateTextureFromSurface(msg)
	if err != nil {
		return err
	}

	width, height, err := f.Font.SizeUTF8(s)
	if err != nil {
		return err
	}
	textWidth := int32(width)
	textHeight := int32(height)

	rect := &sdl.Rect{X: x, Y: y, W: textWidth, H: textHeight}
	R.Copy(texture, nil, rect)
	defer msg.Free()
	defer texture.Destroy()

	return nil
}

func PrintInsideRect(s string, x int32, y int32, rect *sdl.Rect) error {
	f := CurrentFont()
	msg, err := f.Font.RenderUTF8Blended(s, Color)
	if err != nil {
		return err
	}

	texture, err := R.CreateTextureFromSurface(msg)
	if err != nil {
		return err
	}

	R.Copy(texture, nil, rect)
	defer msg.Free()
	defer texture.Destroy()

	return nil
}

func ResetClipRect() {
	R.SetClipRect(&sdl.Rect{X: 0, Y: 0, W: A.ScreenWidth, H: A.ScreenHeight})
}
