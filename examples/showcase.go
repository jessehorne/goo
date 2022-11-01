package main

import gui "github.com/jessehorne/goo"

func main() {
	app, err := gui.NewApp("Goo Showcase", 1000, 1000)
	if err != nil {
		panic(err)
	}

	titleBar := gui.NewContainer(0, 0, app.ScreenWidth, 50)
	titleBar.SetBackgroundColor(gui.COLOR_RED)
	titleBar.SetAlignment(gui.ALIGN_CENTER)

	helloBtn := gui.NewTextButton("Hello!", 0, 0)
	titleBar.AddElement(helloBtn)

	helloBtn2 := gui.NewTextButton("Hello world this is a long button!", 0, 0)
	titleBar.AddElement(helloBtn2)

	app.AddContainer(titleBar)
	app.RunLoop()
}
