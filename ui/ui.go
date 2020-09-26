package ui

import (
	"errors"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/gohouse/go4rdm/config"
	"github.com/gohouse/go4rdm/data"
	"github.com/gohouse/go4rdm/event"
	"github.com/gohouse/go4rdm/uitheme"
	"os"
)
func init() {
	//设置中文字体
	dir, err := os.UserHomeDir()
	if err!=nil {
		panic(err.Error())
	}
	os.Setenv("FYNE_FONT",fmt.Sprintf("%s/%s", dir,"Documents/YaHeiMonacoHybrid.ttf"))
}
type UI struct {
	conf   *config.Config
	data   *data.Redis
	App    fyne.App
	Window fyne.Window
	DefaultSize fyne.Size
}

func NewUI(conf *config.Config, data *data.Redis) *UI {
	return &UI{conf: conf, data: data, DefaultSize: fyne.NewSize(1000, 600)}
}

func (ui *UI) Build() {
	a := app.NewWithID("github.com/gohouse")
	w := a.NewWindow("go for redis desktop manager")

	defer func() {
		if err := recover(); err != nil {
			dialog.ShowError(errors.New(fmt.Sprint(err)), w)
		}
	}()

	w.Resize(ui.DefaultSize)
	ui.App = a
	ui.Window = w
	a.Settings().SetTheme(uitheme.NewDarkBlueNormal())

	// build body
	container := widget.NewTabContainer(
		widget.NewTabItemWithIcon("", theme.HomeIcon(), NewHome(ui).Build()),
		widget.NewTabItemWithIcon("", theme.StorageIcon(), NewRdm(ui).Build()),
		widget.NewTabItemWithIcon("", theme.HelpIcon(), NewQA(ui).Build()),
		//widget.NewTabItemWithIcon("", theme.NewThemedResource(resource.IconOfChat,nil), NewChat(ui).Build()),
		widget.NewTabItemWithIcon("", theme.SettingsIcon(), NewSetting(ui).Build()),
		widget.NewTabItemWithIcon("", theme.FileTextIcon(), NewDocument(ui).Build()),
	)

	container.SetTabLocation(widget.TabLocationTrailing)

	w.SetContent(container)
	// 初始化数据
	if len(ui.conf.Connections) > 0 {
		go event.Produce(event.ETconnectionInit, ui.conf.Connections)
	}
	w.ShowAndRun()
}

func (ui *UI) makeHomeCanvas() fyne.CanvasObject {
	return widget.NewLabel("test")

	//cav := fc.NewCanvas("test",200,200)
	//cav.CText(10,20,22,"text for test", color.RGBA{
	//	R: 0,
	//	G: 66,
	//	B: 99,
	//	A: 255,
	//})
	//
	//return cav.Container
}
