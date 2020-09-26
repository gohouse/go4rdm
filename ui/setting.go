package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/gohouse/go4rdm/uitheme"
	"github.com/gohouse/t"
)
type Setting struct {
	ui *UI
}

func NewSetting(ui *UI) *Setting {
	return &Setting{ui: ui}
}
func (s *Setting) Build() fyne.CanvasObject {
	return widget.NewHBox(
		widget.NewGroup("theme setting", s.buildTheme()),
		widget.NewGroup("redis limit setting", s.buildForm()),
		widget.NewGroup("tool setting",
			widget.NewButtonWithIcon("fullscreen", theme.ViewFullScreenIcon(), func() {
				s.ui.Window.SetFullScreen(!s.ui.Window.FullScreen())
			})),
		)
}
func (s *Setting) buildForm() fyne.CanvasObject {
	var limitEntryRedisKeyList = widget.NewSelect([]string{"10","20","30","50"}, func(str string) {
		//s.ui.conf.UiConf.LimitKey = t.New(str).Int64()
		//s.ui.conf.Save()
	})
	limitEntryRedisKeyList.SetSelected(t.New(s.ui.conf.UiConf.LimitKey).String())
	var limitEntryRedisResultList = widget.NewSelect([]string{"10","20","30","50"}, func(str string) {
		//s.ui.conf.UiConf.LimitContent = t.New(str).Int64()
		//s.ui.conf.Save()
	})
	limitEntryRedisResultList.SetSelected(t.New(s.ui.conf.UiConf.LimitContent).String())

	var form = &widget.Form{
		Items: []*widget.FormItem{
			{
				Text: "keys limit",
				Widget: limitEntryRedisKeyList,
			},
			{
				Text: "result limit",
				Widget: limitEntryRedisResultList,
			},
		},
		OnSubmit: func() {
			if limitEntryRedisKeyList.Selected != "" {
				s.ui.conf.UiConf.LimitKey = t.New(limitEntryRedisKeyList.Selected).Int64()
			}
			if limitEntryRedisResultList.Selected != "" {
				s.ui.conf.UiConf.LimitContent = t.New(limitEntryRedisResultList.Selected).Int64()
			}
			s.ui.conf.Save()
			dialog.ShowInformation("success", "set success", s.ui.Window)
		},
	}

	return form
}
func (s *Setting) buildTheme() fyne.CanvasObject {
	return fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		widget.NewButton("dark", func() {
			s.ui.App.Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewButton("Light", func() {
			s.ui.App.Settings().SetTheme(theme.LightTheme())
		}),
		widget.NewButton("DarkPurple", func() {
			s.ui.App.Settings().SetTheme(uitheme.NewDarkPurple())
		}),
		widget.NewButton("DarkBlue", func() {
			s.ui.App.Settings().SetTheme(uitheme.NewDarkBlue())
		}),
		widget.NewButton("DarkBlueNormal", func() {
			s.ui.App.Settings().SetTheme(uitheme.NewDarkBlueNormal())
		}),
		//widget.NewButton("Bootstrap", func() {
		//	s.ui.App.Settings().SetTheme(uitheme.Bootstrap())
		//}),
	)
}
