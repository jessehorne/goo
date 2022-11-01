package gui

import "C"
import (
	"github.com/veandco/go-sdl2/sdl"
)

func UpdateMousePos(x int32, y int32) {
	A.MouseX = x
	A.MouseY = y
}

func KeyboardEvent(cs []*Container, t uint32, key sdl.Keysym) {
	for _, c := range cs {
		for _, e := range c.Elements {
			switch e.(type) {
			case *TextButton:
				//btn := e.(*TextButton)
				// TODO
			case *ImgButton:
				//btn := e.(*ImgButton)
				// TODO
			case *OneLineInput:
				i := e.(*OneLineInput)
				i.KeyboardEvent(t, key)
			}
			RedrawAllContainers(cs)
		}
	}
}

func MouseButtonEvent(cs []*Container, event *sdl.MouseButtonEvent) {
	// Mouse Button Event for all elements in all containers
	for _, c := range cs {
		for _, e := range c.Elements {
			switch e.(type) {
			case *TextButton:
				btn := e.(*TextButton)
				btn.MouseButtonEvent(event)
			case *ImgButton:
				btn := e.(*ImgButton)
				btn.MouseButtonEvent(event)
			case *OneLineInput:
				i := e.(*OneLineInput)
				i.MouseButtonEvent(event)
			}
			RedrawAllContainers(cs)
		}
	}
}

func UpdateHoverStateForAllContainers(cs []*Container) {
	// Update Element Hovers
	for _, c := range cs {
		// skip this step if the mouse isn't actually inside the container
		inside := A.MouseX >= c.X && A.MouseY <= c.X+c.Width && A.MouseY >= c.Y && A.MouseY <= c.Y+c.Height

		for _, e := range c.Elements {
			// if button
			switch e.(type) {
			case *TextButton:
				btn := e.(*TextButton)
				rect1 := sdl.Rect{X: A.MouseX, Y: A.MouseY, W: 1, H: 1}
				rect2 := sdl.Rect{X: btn.X, Y: btn.Y, W: btn.Width, H: btn.Height}
				if BBox(rect1, rect2) && inside {
					btn.SetHover(true)
				} else {
					btn.SetHover(false)
				}
			case *ImgButton:
				btn := e.(*ImgButton)
				rect1 := sdl.Rect{X: A.MouseX, Y: A.MouseY, W: 1, H: 1}
				rect2 := sdl.Rect{X: btn.X, Y: btn.Y, W: btn.Width, H: btn.Height}
				if BBox(rect1, rect2) && inside {
					btn.SetHover(true)
				} else {
					btn.SetHover(false)
				}
			case *OneLineInput:
				i := e.(*OneLineInput)
				rect1 := sdl.Rect{X: A.MouseX, Y: A.MouseY, W: 1, H: 1}
				rect2 := sdl.Rect{X: i.X, Y: i.Y, W: i.Width, H: i.Height}
				if BBox(rect1, rect2) && inside {
					i.SetHover(true)
				} else {
					i.SetHover(false)
				}
			}
			RedrawAllContainers(cs)
		}
	}
}
