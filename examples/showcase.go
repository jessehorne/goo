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
	titleLabel.SetFont("normal", 36)
	titleBar.AddElement(titleLabel)

	// Text Buttons
	buttonsContainer := gui.NewContainer(0, 50, app.ScreenWidth, 50)
	buttonsContainer.SetBackgroundColor(gui.COLOR_GREEN)
	buttonsContainer.SetAlignment(gui.ALIGN_CENTER)

	helloBtn := gui.NewTextButton("Hello!", 0, 0)
	helloBtn.SetFont("normal", 24)
	buttonsContainer.AddElement(helloBtn)

	helloBtn2 := gui.NewTextButton("Hello world this is a long button!", 0, 0)
	helloBtn2.SetFont("normal", 24)
	buttonsContainer.AddElement(helloBtn2)

	// Image Buttons
	imgButtonsContainer := gui.NewContainer(0, 100, app.ScreenWidth, 50)
	imgButtonsContainer.SetBackgroundColor(gui.COLOR_BLUE)
	imgButtonsContainer.SetAlignment(gui.ALIGN_CENTER)

	imgButton1, err := gui.NewImgButton("./resources/images/refresh.svg", 0, 0)
	if err != nil {
		panic(err)
	}
	imgButton1.SetSize(40, 40)
	imgButtonsContainer.AddElement(imgButton1)

	// Inputs
	inputsContainer := gui.NewContainer(0, 150, app.ScreenWidth, 50)
	inputsContainer.SetBackgroundColor(gui.COLOR_GREY)
	inputsContainer.SetAlignment(gui.ALIGN_CENTER)

	smallInput := gui.NewOneLineInput("text goes here", 0, 0, 400, 32)
	smallInput.SetFont("normal", 24)
	inputsContainer.AddElement(smallInput)

	app.AddContainer(titleBar)
	app.AddContainer(buttonsContainer)
	app.AddContainer(imgButtonsContainer)
	app.AddContainer(inputsContainer)
	app.RunLoop()
}
