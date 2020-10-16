package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/gohouse/go4rdm/uitheme"
	"github.com/gohouse/t"
	"image/color"
	"time"
)

type Setting struct {
	ui *UI
}

func NewSetting(ui *UI) *Setting {
	return &Setting{ui: ui}
}
func (s *Setting) Build() fyne.CanvasObject {
	return widget.NewHBox(
		widget.NewVBox(
			widget.NewGroup("redis limit setting", s.buildForm()),
			widget.NewGroup("default page setting", s.buildDefaultSettingForm()),
			),
		layout.NewSpacer(),
		widget.NewVBox(
			layout.NewSpacer(),
			s.buildMid(),
			),
		layout.NewSpacer(),
		widget.NewVBox(
			widget.NewGroup("tool setting",
				widget.NewButtonWithIcon("fullscreen", theme.ViewFullScreenIcon(), func() {
					s.ui.Window.SetFullScreen(!s.ui.Window.FullScreen())
				})),
			widget.NewGroup("theme setting", s.buildTheme()),
		),
	)
}
func (s *Setting) buildMid() fyne.CanvasObject {
	text := canvas.NewText("", &color.RGBA{222,222,007, 255})
	text.TextSize = 24
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-ticker.C:
				//text.Text = time.Now().Format("2006-01-02 15:04:05")
				text.Text = time.Now().Format(time.RFC1123Z)
				text.Refresh()
			}
		}
	}()
	return text
}
func (s *Setting) buildDefaultSettingForm() fyne.CanvasObject {
	var pageList = widget.NewSelect([]string{"Home", "RedisManager", "QA", "Setting", "Document"}, func(str string) {
		//s.ui.conf.UiConf.LimitKey = t.New(str).Int64()
		//s.ui.conf.Save()
	})
	pageList.SetSelected(t.New(s.ui.conf.UiConf.DefaultPage).String())

	var rdmList = widget.NewSelect([]string{"Connection", "Result", "Command"}, func(str string) {
		//s.ui.conf.UiConf.LimitKey = t.New(str).Int64()
		//s.ui.conf.Save()
	})
	rdmList.SetSelected(t.New(s.ui.conf.UiConf.DefaultRdm).String())

	var form = &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:   "default page",
				Widget: pageList,
			},
			{
				Text:   "default rdm",
				Widget: rdmList,
			},
		},
		OnSubmit: func() {
			if pageList.Selected != "" {
				s.ui.conf.UiConf.DefaultPage = pageList.Selected
			}
			if rdmList.Selected != "" {
				s.ui.conf.UiConf.DefaultRdm = rdmList.Selected
			}
			s.ui.conf.Save()
			dialog.ShowInformation("success", "set success", s.ui.Window)
		},
	}

	return form
}
func (s *Setting) buildForm() fyne.CanvasObject {
	var limitEntryRedisKeyList = widget.NewSelect([]string{"10", "20", "30", "50"}, func(str string) {
		//s.ui.conf.UiConf.LimitKey = t.New(str).Int64()
		//s.ui.conf.Save()
	})
	limitEntryRedisKeyList.SetSelected(t.New(s.ui.conf.UiConf.LimitKey).String())
	var limitEntryRedisResultList = widget.NewSelect([]string{"10", "20", "30", "50"}, func(str string) {
		//s.ui.conf.UiConf.LimitContent = t.New(str).Int64()
		//s.ui.conf.Save()
	})
	limitEntryRedisResultList.SetSelected(t.New(s.ui.conf.UiConf.LimitContent).String())

	var form = &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:   "keys limit",
				Widget: limitEntryRedisKeyList,
			},
			{
				Text:   "result limit",
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
	return fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		widget.NewButton("DarkBlueNormal", func() {
			s.ui.SetTheme(uitheme.NewDarkBlueNormal())
		}),
		widget.NewButton("dark", func() {
			s.ui.SetTheme(theme.DarkTheme())
		}),
		widget.NewButton("Light", func() {
			s.ui.SetTheme(theme.LightTheme())
		}),
		widget.NewButton("DarkPurple", func() {
			s.ui.SetTheme(uitheme.NewDarkPurple())
		}),
		widget.NewButton("DarkBlue", func() {
			s.ui.SetTheme(uitheme.NewDarkBlue())
		}),
		//widget.NewButton("Bootstrap", func() {
		//	s.ui.App.Settings().SetTheme(uitheme.Bootstrap())
		//}),
	)
}
