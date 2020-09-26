package uitheme

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

type DarkBlueNormal struct{}

func NewDarkBlueNormal() *DarkBlueNormal {
	return &DarkBlueNormal{}
}

func (DarkBlueNormal) BackgroundColor() color.Color {
	return ColorDark
}

//func (DarkBlueNormal) ButtonColor() color.Color { return color.RGBA{R: 0x6f, G: 0x42, B: 0xc1, A: 0xff} }	// 紫色
//func (DarkBlueNormal) ButtonColor() color.Color { return color.RGBA{R: 0x2e, G: 0x22, B: 0x8b, A: 0xff} }	// 2E228BFF
func (DarkBlueNormal) ButtonColor() color.Color { return color.RGBA{R: 0x14, G: 0x14, B: 0x14, A: 0xff} } // 141414FF
func (DarkBlueNormal) DisabledButtonColor() color.Color {
	return color.RGBA{R: 0xf, G: 0xf, B: 0x11, A: 0xff}
}
func (DarkBlueNormal) TextColor() color.Color { return ColorLight }
func (DarkBlueNormal) DisabledTextColor() color.Color {
	return color.RGBA{R: 0xc8, G: 0xc8, B: 0xc8, A: 0xff}
}
func (DarkBlueNormal) IconColor() color.Color { return color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff} }
func (DarkBlueNormal) DisabledIconColor() color.Color {
	return color.RGBA{R: 0xc8, G: 0xc8, B: 0xc8, A: 0xff}
}
func (DarkBlueNormal) HyperlinkColor() color.Color {
	return color.RGBA{R: 0x0, G: 0x7b, B: 0xff, A: 0xff}
}
func (DarkBlueNormal) PlaceHolderColor() color.Color {
	return color.RGBA{R: 0x6c, G: 0x75, B: 0x7d, A: 0xff}
}
func (DarkBlueNormal) PrimaryColor() color.Color        { return color.RGBA{R: 0x0, G: 0x7b, B: 0xff, A: 0xff} }
func (DarkBlueNormal) HoverColor() color.Color          { return ColorSecondary }
func (DarkBlueNormal) FocusColor() color.Color          { return color.RGBA{R: 0x0, G: 0x7b, B: 0xff, A: 0xff} }
func (DarkBlueNormal) ScrollBarColor() color.Color {
	return color.RGBA{R: 0x23, G: 0x23, B: 0x23, A: 0x8}
}
func (DarkBlueNormal) ShadowColor() color.Color         { return color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x40} }
func (DarkBlueNormal) TextSize() int                    { return 14 }
func (DarkBlueNormal) TextFont() fyne.Resource          { return theme.LightTheme().TextFont() }
//func (DarkBlueNormal) TextFont() fyne.Resource          { return resource.FontOfYaHeiMonacoHybrid }
func (DarkBlueNormal) TextBoldFont() fyne.Resource      { return theme.LightTheme().TextBoldFont() }
func (DarkBlueNormal) TextItalicFont() fyne.Resource    { return theme.LightTheme().TextItalicFont() }
func (DarkBlueNormal) TextBoldItalicFont() fyne.Resource {
	return theme.LightTheme().TextBoldItalicFont()
}
func (DarkBlueNormal) TextMonospaceFont() fyne.Resource { return theme.LightTheme().TextMonospaceFont() }
//func (DarkBlueNormal) TextMonospaceFont() fyne.Resource { return resource.FontOfYaHeiMonacoHybrid }
func (DarkBlueNormal) Padding() int                     { return 4 }
func (DarkBlueNormal) IconInlineSize() int              { return 12 }
func (DarkBlueNormal) ScrollBarSize() int               { return 12 }
func (DarkBlueNormal) ScrollBarSmallSize() int          { return 3 }
