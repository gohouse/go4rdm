package uitheme

import (
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"github.com/blockpane/prettyfyne"
	"github.com/gohouse/go4rdm/fmcolor"
)

func Bootstrap() fyne.Theme {
	myTheme := prettyfyne.PrettyTheme{
		BackgroundColor:     fmcolor.Normal.Light,
		ButtonColor:         fmcolor.Light.Warning,
		DisabledButtonColor: fmcolor.Normal.Gray,
		HyperlinkColor:      fmcolor.Normal.Blue,
		TextColor:           fmcolor.Normal.Dark,
		DisabledTextColor:   fmcolor.Light.Light,
		IconColor:           fmcolor.Normal.Primary,
		DisabledIconColor:   fmcolor.Normal.Gray,
		PlaceHolderColor:    fmcolor.Normal.Gray,
		PrimaryColor:        fmcolor.Dark.Warning,
		HoverColor:          fmcolor.Dark.Warning,
		FocusColor:          fmcolor.Light.Warning,
		ScrollBarColor:      fmcolor.Normal.Gray,
		ShadowColor:         fmcolor.Normal.Gray,
		TextSize:            13,
		TextFont:            theme.DarkTheme().TextFont(),
		TextBoldFont:        theme.DarkTheme().TextBoldFont(),
		TextItalicFont:      theme.DarkTheme().TextItalicFont(),
		TextBoldItalicFont:  theme.DarkTheme().TextBoldItalicFont(),
		TextMonospaceFont:   theme.DarkTheme().TextMonospaceFont(),
		Padding:             4,
		IconInlineSize:      24,
		ScrollBarSize:       10,
		ScrollBarSmallSize:  4,
	}

	return myTheme.ToFyneTheme()
}
