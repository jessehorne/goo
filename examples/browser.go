package main

import (
	"fmt"
	gui "github.com/jessehorne/goo"
)

func main() {
	app, err := gui.NewApp("example", 900, 600)
	if err != nil {
		panic(err)
	}
	app.SetTitle("example 2")

	topBar := gui.NewContainer(0, 0, app.ScreenWidth, 36+(10*2), gui.THEME_BLUE_90)

	backImgButton, err := gui.NewImgButton("./resources/images/arrow-left.svg", 10, 10)
	if err != nil {
		panic(err)
	}
	topBar.AddElement(backImgButton)

	refreshImgButton, err := gui.NewImgButton("./resources/images/refresh.svg", 10+36, 10)
	if err != nil {
		panic(err)
	}
	topBar.AddElement(refreshImgButton)

	forwardImgButton, err := gui.NewImgButton("./resources/images/arrow-right.svg", 10+(36*2), 10)
	if err != nil {
		panic(err)
	}
	topBar.AddElement(forwardImgButton)

	tiX := int32(10 + (36 * 3) + 20)
	tiY := int32(10)
	tiW := int32(900 - tiX - (36 + 20))
	tiH := int32(36)
	textInput := gui.NewOneLineInput("http://example.com", tiX, tiY, tiW, tiH)
	textInput.SetCallback(func(i *gui.OneLineInput) {
		fmt.Println(i.Text)
	})
	topBar.AddElement(textInput)

	searchX := tiX + tiW + 10
	searchButton, err := gui.NewImgButton("./resources/images/eye-empty.svg", searchX, 10)
	if err != nil {
		panic(err)
	}
	searchButton.SetCallback(func(b *gui.ImgButton) {
		fmt.Println("pressed search button")
	})
	topBar.AddElement(searchButton)

	textButton := gui.NewTextButton("Testing 123", 200, 200)
	textButton.SetCallback(func(b *gui.TextButton) {
		fmt.Println("pressed text button")
	})
	topBar.AddElement(textButton)

	app.AddContainer(topBar)

	app.RunLoop()
}
