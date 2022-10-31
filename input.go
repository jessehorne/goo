package gui

import (
	"github.com/veandco/go-sdl2/sdl"
	"strings"
)

type OneLineInput struct {
	X                 int32
	Y                 int32
	Width             int32
	Height            int32
	BackgroundColor   sdl.Color
	OutlineColor      sdl.Color
	TextColor         sdl.Color
	HighlightColor    sdl.Color
	PlaceholderColor  sdl.Color
	CursorPos         int32
	CursorX           int32
	MaxChars          int32
	Text              string
	Placeholder       string
	HighlightStartPos int32
	HighlightEndPos   int32
	MarginX           int32
	MarginY           int32
	PaddingX          int32
	PaddingY          int32
	Active            bool
	IsHovering        bool
	IsHoldingShift    bool
	Callback          func(i *OneLineInput)
}

func NewOneLineInput(placeholder string, x int32, y int32, w int32, h int32) *OneLineInput {
	return &OneLineInput{
		X:                 x,
		Y:                 y,
		Width:             w,
		Height:            h,
		BackgroundColor:   THEME_BLUE_80,
		OutlineColor:      THEME_BLUE_90,
		TextColor:         COLOR_WHITE,
		HighlightColor:    THEME_BLUE_60,
		PlaceholderColor:  COLOR_WHITE_ALPHA,
		CursorPos:         0,
		MaxChars:          1000,
		Text:              "",
		Placeholder:       placeholder,
		HighlightStartPos: 0,
		HighlightEndPos:   0,
		MarginX:           0,
		MarginY:           0,
		PaddingX:          5,
		PaddingY:          2,
		Active:            false,
		IsHovering:        false,
		IsHoldingShift:    false,
	}
}

func (i *OneLineInput) MouseButtonEvent(event *sdl.MouseButtonEvent) {
	if i.IsHovering {
		if event.Button == sdl.BUTTON_LEFT {
			i.SetActive(true)
		}
	} else {
		i.SetActive(false)
	}
}

func (i *OneLineInput) SetCallback(f func(i *OneLineInput)) {
	i.Callback = f
}

func (i *OneLineInput) Trigger() {
	if i.Callback != nil {
		i.Callback(i)
	}
}

func (i *OneLineInput) KeyboardEvent(t uint32, key sdl.Keysym) {
	SetFont("light")

	if i.Active {
		k := sdl.GetKeyName(key.Sym)
		if t == sdl.KEYUP {
			if k == "Return" {
				i.Trigger()
				return
			}
		}

		// 768 below is kinda magic, but it's the type uint32 for keydown as well as holding a key, so it works well for
		// handling keyboard stuff and also letting users hold keys
		if t == sdl.KEYDOWN && t == 768 {
			if k == "Space" {
				i.AddText(" ")
				return
			}

			if k == "Backspace" {
				// delete highlighted or delete last
				if i.HighlightEndPos > 0 {
					i.ClearText()
				} else {
					i.Backspace()
				}
				return
			}

			// Copy/Pasta
			isCopying := k == "C" && key.Mod == uint16(64)
			isPasting := k == "V" && key.Mod == uint16(64)

			if isCopying {
				err := sdl.SetClipboardText(i.Text[i.HighlightStartPos:i.HighlightEndPos])
				if err != nil {
					return
				}
				return
			}

			if isPasting {
				copyText, err := sdl.GetClipboardText()
				if err != nil {
					return
				}
				i.AddText(copyText)
				return
			}

			if k == "Left" {
				i.MoveCursor(i.CursorPos - 1)
			} else if k == "Right" {
				i.MoveCursor(i.CursorPos + 1)
			} else if k == "Up" {
				i.MoveCursor(0)
			} else if k == "Down" {
				i.MoveCursor(int32(len(i.Text)))
			}

			// check if char is typeable
			charIndex := strings.Index(CharsTypeable, k)
			if charIndex == -1 {
				return
			}

			// if mod == 1 or 2 then user is holding left or right shift
			if key.Mod == uint16(1) || key.Mod == uint16(2) {
				i.AddText(string(CharsUpper[charIndex]))
			} else {
				i.AddText(string(CharsLower[charIndex]))
			}
		}
	}
}

func (i *OneLineInput) SetHover(which bool) {
	i.IsHovering = which
}

func (i *OneLineInput) SetActive(which bool) {
	if which {
		i.HighlightStartPos = 0
		i.HighlightEndPos = int32(len(i.Text))
	} else {
		i.HighlightStartPos = 0
		i.HighlightEndPos = 0
	}

	i.CursorPos = i.HighlightEndPos
	i.Active = which
}

func (i *OneLineInput) MoveCursor(pos int32) {
	i.HighlightStartPos = 0
	i.HighlightEndPos = 0

	if pos < 0 {
		i.CursorPos = 0
	} else if pos > int32(len(i.Text)) {
		i.CursorPos = int32(len(i.Text))
	} else {
		i.CursorPos = pos
	}

	SetFont("light")
	w, _ := GetTextSize(i.Text[:i.CursorPos])
	i.CursorX = int32(w)
}

func (i *OneLineInput) AddText(t string) {
	SetFont("light")

	if i.HighlightEndPos > int32(0) {
		i.ClearText()
		i.HighlightEndPos = 0
	}

	i.CursorPos += int32(len(t))
	var first, second string
	if i.CursorPos == int32(len(i.Text))+1 {
		first = i.Text
		second = t
	} else {
		if i.CursorPos > int32(len(i.Text)) {
			first = i.Text
			second = t
		} else {
			first = i.Text[:i.CursorPos-1]
			second = t + i.Text[i.CursorPos-1:]
		}
	}

	i.Text = first + second

	w, _ := GetTextSize(i.Text[:i.CursorPos])
	i.CursorX = int32(w)
}

func (i *OneLineInput) Backspace() {
	if len(i.Text) == 0 {
		return
	}

	if i.CursorPos == int32(0) {
		return
	}

	i.CursorPos -= int32(1)
	var first, second string
	first = i.Text[:i.CursorPos]
	second = i.Text[i.CursorPos+1:]

	i.Text = first + second

	SetFont("light")
	w, _ := GetTextSize(i.Text[:i.CursorPos])
	i.CursorX = int32(w)
}

func (i *OneLineInput) ClearText() {
	i.CursorPos = 0
	i.CursorX = int32(0)
	i.Text = ""
}

func (i *OneLineInput) Draw() {
	originX := i.X + i.MarginX
	originY := i.Y + i.MarginY

	rect := sdl.Rect{X: originX, Y: originY, W: i.Width, H: i.Height}

	// Background
	SetColor(i.BackgroundColor)
	DrawFilledRect(&rect)

	// Outline
	SetColor(i.OutlineColor)
	DrawLineRect(&rect)

	SetFont("light")

	// Draw Placeholder
	if i.Active {
		width, height, err := Fonts[CurrentFont].Font.SizeUTF8(i.Text)

		// draw highlight background
		if i.HighlightEndPos > 0 {
			if err != nil {
				return
			}

			boxX := originX + i.PaddingX + 2
			boxY := originY + i.PaddingY + 2
			boxW := int32(width)
			boxH := int32(height) - 2
			rect := sdl.Rect{X: boxX, Y: boxY, W: boxW, H: boxH}

			SetColor(i.HighlightColor)
			DrawFilledRect(&rect)
		}

		printX := originX + i.PaddingX + 2
		printY := originY + i.PaddingY
		printW := i.Width - i.PaddingX - 8
		printH := int32(height) - 2
		printRect := sdl.Rect{X: printX, Y: printY, W: printW, H: printH}

		R.SetClipRect(&printRect)

		// draw text if it's set
		if len(i.Text) > 0 {

			SetColor(COLOR_WHITE)
			Print(i.Text, printX, printY)
		}

		cursorX := originX + i.PaddingX + 2 + i.CursorX
		cursorY := originY + i.PaddingY + 2
		lineRect := sdl.Rect{X: cursorX, Y: cursorY, W: 2, H: int32(height) - 2}
		SetColor(COLOR_WHITE)
		DrawFilledRect(&lineRect)

		ResetClipRect()
	} else {
		if len(i.Text) == 0 {
			// draw placeholder text if there's no text and input isn't active
			SetColor(i.PlaceholderColor)
			Print(i.Placeholder, originX+i.PaddingX, originY+i.PaddingY)
		} else {
			// draw text since it's set and input isn't active
			SetColor(COLOR_WHITE)
			Print(i.Text, originX+i.PaddingX+2, originY+i.PaddingY)
		}
	}
}
