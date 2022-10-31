package gui

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type App struct {
	Title        string
	ScreenWidth  int32
	ScreenHeight int32
	Containers   []*Container
	Running      bool
	MouseX       int32
	MouseY       int32
}

func NewApp(title string, w int32, h int32) (*App, error) {
	newApp := App{
		Title:        title,
		ScreenWidth:  w,
		ScreenHeight: h,
		Containers:   nil,
		Running:      true,
		MouseX:       0,
		MouseY:       0,
	}

	err := Init(&newApp)
	if err != nil {
		return nil, err
	}

	return &newApp, nil
}

func (a *App) SetTitle(title string) {
	W.SetTitle(title)
}

func (a *App) AddContainer(c *Container) {
	a.Containers = append(a.Containers, c)

	// redraw containers
	RedrawAllContainers(a.Containers)
}

func (a *App) RunLoop() {
	for a.Running {
		// Events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("Exiting inet...")
				a.Running = false
				break
			case *sdl.MouseMotionEvent:
				e := event.(*sdl.MouseMotionEvent)
				// Update Mouse Position
				UpdateMousePos(e.X, e.Y)

				// Update Loop
				UpdateHoverStateForAllContainers(a.Containers)
			case *sdl.MouseButtonEvent:
				e := event.(*sdl.MouseButtonEvent)
				MouseButtonEvent(a.Containers, e)
			case *sdl.KeyboardEvent:
				e := event.(*sdl.KeyboardEvent)
				KeyboardEvent(a.Containers, e.Type, e.Keysym)
			}
		}
	}

	W.Destroy()
	sdl.Quit()
}

func (s *App) Close() {
	s.Running = false
}
