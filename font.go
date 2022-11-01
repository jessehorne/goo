package gui

import (
	"errors"
	"github.com/veandco/go-sdl2/ttf"
)

type Font struct {
	Font     *ttf.Font
	FontSize int32
}

var FontSizeRange = []int32{8, 14, 20, 24, 32, 36, 40, 48, 52}
var Fonts = make(map[string]map[int32]*Font)
var FontStyle = "normal" // can be regular or light
var FontSize int32 = 14

func CurrentFont() *Font {
	fontStyle, fontStyleExists := Fonts[FontStyle]
	if !fontStyleExists {
		return nil
	}

	f, exists := fontStyle[FontSize]
	if !exists {
		return nil
	}

	return f
}

func InitFonts() error {
	err := ttf.Init()
	if err != nil {
		return err
	}

	err = LoadFontRanges("normal", FontSizeRange)
	if err != nil {
		return err
	}

	err = LoadFontRanges("light", FontSizeRange)
	if err != nil {
		return err
	}

	return nil
}

func LoadFontRanges(fontStyle string, r []int32) error {
	for _, fontSize := range r {
		err := LoadFont(fontStyle, "./resources/fonts/Trueno.otf", fontSize)
		if err != nil {
			return err
		}
	}

	return nil
}

func LoadFont(fontStyle string, path string, fontSize int32) error {
	font, err := ttf.OpenFont(path, int(fontSize))
	if err != nil {
		return err
	}

	_, fontStyleExists := Fonts[fontStyle]
	if !fontStyleExists {
		Fonts[fontStyle] = make(map[int32]*Font)
	}

	Fonts[fontStyle][fontSize] = &Font{
		Font:     font,
		FontSize: fontSize,
	}

	return nil
}

func SetFont(fontStyle string, fontSize int32) error {
	fs, fontStyleExists := Fonts[fontStyle]
	if !fontStyleExists {
		return errors.New("no font exists with that fontStyle")
	}

	_, exists := fs[fontSize]
	if !exists {
		return errors.New("no font exists with that fontStyle and fontSize")
	}

	FontStyle = fontStyle
	FontSize = fontSize

	return nil
}

func GetTextSize(s string) (int, int) {
	fs, fontStyleExists := Fonts[FontStyle]
	if !fontStyleExists {
		return -1, -1
	}

	_, exists := fs[FontSize]
	if !exists {
		return -1, -1
	}

	w, h, err := fs[FontSize].Font.SizeUTF8(s)
	if err != nil {
		return -1, -1
	}

	return w, h
}
