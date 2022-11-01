package gui

func Cleanup() {
	// cleanup all fonts
	for _, fs := range Fonts {
		for _, f := range fs {
			f.Font.Close()
		}
	}
}
