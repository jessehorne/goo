package gui

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type Image struct {
	X       int32
	Y       int32
	Width   int32
	Height  int32
	Texture *sdl.Texture
}

func NewImageFromSVG(path string, x int32, y int32, w int32, h int32) (*Image, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	rw, err := sdl.RWFromMem(data)
	if err != nil {
		return nil, err
	}

	surface, err := img.LoadRW(rw, true)
	if err != nil {
		return nil, err
	}

	texture, err := R.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, err
	}

	return &Image{
		X:       x,
		Y:       y,
		Width:   w,
		Height:  h,
		Texture: texture,
	}, nil
}

func DrawImage(img *Image, x int32, y int32) {
	rect := &sdl.Rect{X: x, Y: y, W: img.Width, H: img.Height}
	R.Copy(img.Texture, nil, rect)
}
