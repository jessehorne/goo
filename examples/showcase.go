package main

import gui "github.com/jessehorne/goo"

func main() {
	app, err := gui.NewApp("Goo Showcase", 1000, 1000)
	if err != nil {
		panic(err)
	}
	app.SetBackgroundColor(gui.COLOR_WHITE)

	// Top bar
	titleBar := gui.NewContainer(0, 0, app.ScreenWidth, 50)
	titleBar.SetBackgroundColor(gui.COLOR_RED)
	titleBar.SetAlignment(gui.ALIGN_CENTER)

	titleLabel := gui.NewLabel("Goo Showcase", 32, 0, 0)
	titleBar.AddElement(titleLabel)

	// Buttons
	buttonsContainer := gui.NewContainer(0, 50, app.ScreenWidth, 50)
	buttonsContainer.SetBackgroundColor(gui.COLOR_GREEN)
	buttonsContainer.SetAlignment(gui.ALIGN_CENTER)

	helloBtn := gui.NewTextButton("Hello!", 0, 0)
	buttonsContainer.AddElement(helloBtn)

	helloBtn2 := gui.NewTextButton("Hello world this is a long button!", 0, 0)
	buttonsContainer.AddElement(helloBtn2)

	// Inputs
	inputsContainer := gui.NewContainer(0, 100, app.ScreenWidth, 50)
	inputsContainer.SetBackgroundColor(gui.COLOR_BLUE)
	inputsContainer.SetAlignment(gui.ALIGN_CENTER)

	smallInput := gui.NewOneLineInput("text goes here", 0, 0, 400, 32)
	smallInput.SetFont("normal", 24)
	inputsContainer.AddElement(smallInput)

	app.AddContainer(titleBar)
	app.AddContainer(buttonsContainer)
	app.AddContainer(inputsContainer)
	app.RunLoop()
}
