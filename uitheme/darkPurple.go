package uitheme

import (
"image/color"

"fyne.io/fyne"
"fyne.io/fyne/theme"
)

type darkPurple struct{}

func NewDarkPurple() *darkPurple {
	return &darkPurple{}
}

func (darkPurple) BackgroundColor() color.Color {
	return color.RGBA{R: 0x1e, G: 0x1e, B: 0x1e, A: 0xff}
}
//func (darkPurple) ButtonColor() color.Color { return color.RGBA{R: 0x6f, G: 0x42, B: 0xc1, A: 0xff} }	// 紫色
func (darkPurple) ButtonColor() color.Color { return color.RGBA{R: 0x2e, G: 0x22, B: 0x8b, A: 0xff} }	// 2E228BFF
func (darkPurple) DisabledButtonColor() color.Color {
	return color.RGBA{R: 0xf, G: 0xf, B: 0x11, A: 0xff}
}
func (darkPurple) TextColor() color.Color { return color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff} }
func (darkPurple) DisabledTextColor() color.Color {
	return color.RGBA{R: 0xc8, G: 0xc8, B: 0xc8, A: 0xff}
}
func (darkPurple) IconColor() color.Color { return color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff} }
func (darkPurple) DisabledIconColor() color.Color {
	return color.RGBA{R: 0xc8, G: 0xc8, B: 0xc8, A: 0xff}
}
func (darkPurple) HyperlinkColor() color.Color { return color.RGBA{R: 0x0, G: 0x7b, B: 0xff, A: 0xff} }
func (darkPurple) PlaceHolderColor() color.Color {
	return color.RGBA{R: 0x6c, G: 0x75, B: 0x7d, A: 0xff}
}
func (darkPurple) PrimaryColor() color.Color         { return color.RGBA{R: 0x0, G: 0x7b, B: 0xff, A: 0xff} }
func (darkPurple) HoverColor() color.Color           { return color.RGBA{R: 0x66, G: 0x10, B: 0xf2, A: 0xff} }
func (darkPurple) FocusColor() color.Color           { return color.RGBA{R: 0x0, G: 0x7b, B: 0xff, A: 0xff} }
func (darkPurple) ScrollBarColor() color.Color       { return color.RGBA{R: 0x23, G: 0x23, B: 0x23, A: 0x8} }
func (darkPurple) ShadowColor() color.Color          { return color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x40} }
func (darkPurple) TextSize() int                     { return 14 }
func (darkPurple) TextFont() fyne.Resource           { return theme.LightTheme().TextFont() }
func (darkPurple) TextBoldFont() fyne.Resource       { return theme.LightTheme().TextBoldFont() }
func (darkPurple) TextItalicFont() fyne.Resource     { return theme.LightTheme().TextItalicFont() }
func (darkPurple) TextBoldItalicFont() fyne.Resource { return theme.LightTheme().TextBoldItalicFont() }
func (darkPurple) TextMonospaceFont() fyne.Resource  { return theme.LightTheme().TextMonospaceFont() }
func (darkPurple) Padding() int                      { return 4 }
func (darkPurple) IconInlineSize() int               { return 20 }
func (darkPurple) ScrollBarSize() int                { return 16 }
func (darkPurple) ScrollBarSmallSize() int           { return 3 }

