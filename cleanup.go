package gui

func Cleanup() {
	// cleanup all fonts
	for _, f := range Fonts {
		f.Font.Close()
	}
}
