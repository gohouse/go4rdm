package uitheme

import (
"image/color"

"fyne.io/fyne"
"fyne.io/fyne/theme"
)

type DarkBlue struct{}

func NewDarkBlue() *DarkBlue {
	return &DarkBlue{}
}

func (DarkBlue) BackgroundColor() color.Color {
	return color.RGBA{R: 0x1e, G: 0x1e, B: 0x1e, A: 0xff}
}
//func (DarkBlue) ButtonColor() color.Color { return color.RGBA{R: 0x6f, G: 0x42, B: 0xc1, A: 0xff} }	// 紫色
//func (DarkBlue) ButtonColor() color.Color { return color.RGBA{R: 0x2e, G: 0x22, B: 0x8b, A: 0xff} }	// 2E228BFF
func (DarkBlue) ButtonColor() color.Color { return color.RGBA{R: 0x14, G: 0x14, B: 0x14, A: 0xff} }	// 141414FF
func (DarkBlue) DisabledButtonColor() color.Color {
	return color.RGBA{R: 0xf, G: 0xf, B: 0x11, A: 0xff}
}
func (DarkBlue) TextColor() color.Color { return color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff} }
func (DarkBlue) DisabledTextColor() color.Color {
	return color.RGBA{R: 0xc8, G: 0xc8, B: 0xc8, A: 0xff}
}
func (DarkBlue) IconColor() color.Color { return color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff} }
func (DarkBlue) DisabledIconColor() color.Color {
	return color.RGBA{R: 0xc8, G: 0xc8, B: 0xc8, A: 0xff}
}
func (DarkBlue) HyperlinkColor() color.Color { return color.RGBA{R: 0x0, G: 0x7b, B: 0xff, A: 0xff} }
func (DarkBlue) PlaceHolderColor() color.Color {
	return color.RGBA{R: 0x6c, G: 0x75, B: 0x7d, A: 0xff}
}
func (DarkBlue) PrimaryColor() color.Color         { return color.RGBA{R: 0x0, G: 0x7b, B: 0xff, A: 0xff} }
func (DarkBlue) HoverColor() color.Color           { return color.RGBA{R: 0x66, G: 0x10, B: 0xf2, A: 0xff} }
func (DarkBlue) FocusColor() color.Color           { return color.RGBA{R: 0x0, G: 0x7b, B: 0xff, A: 0xff} }
func (DarkBlue) ScrollBarColor() color.Color       { return color.RGBA{R: 0x23, G: 0x23, B: 0x23, A: 0x8} }
func (DarkBlue) ShadowColor() color.Color          { return color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x40} }
func (DarkBlue) TextSize() int                     { return 14 }
func (DarkBlue) TextFont() fyne.Resource           { return theme.LightTheme().TextFont() }
func (DarkBlue) TextBoldFont() fyne.Resource       { return theme.LightTheme().TextBoldFont() }
func (DarkBlue) TextItalicFont() fyne.Resource     { return theme.LightTheme().TextItalicFont() }
func (DarkBlue) TextBoldItalicFont() fyne.Resource { return theme.LightTheme().TextBoldItalicFont() }
func (DarkBlue) TextMonospaceFont() fyne.Resource  { return theme.LightTheme().TextMonospaceFont() }
func (DarkBlue) Padding() int                      { return 4 }
func (DarkBlue) IconInlineSize() int               { return 12 }
func (DarkBlue) ScrollBarSize() int                { return 12 }
func (DarkBlue) ScrollBarSmallSize() int           { return 3 }

