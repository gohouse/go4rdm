package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/gohouse/go4rdm/event"
	"sync"
)


type Rdm struct {
	ui *UI
	tabgroup *widget.TabContainer
	rdmkl *RdmKeylist
	//rdmc *RdmConnections
	//rdmct *RdmContent
}

var rdmOnce sync.Once
var rdmObj *Rdm
func NewRdm(ui *UI) *Rdm {
	rdmOnce.Do(func() {
		rdmObj = &Rdm{}
		rdmObj.init()
	})
	rdmObj.ui = ui
	return rdmObj
}

func (rdm *Rdm) init() {
	event.NewEvent().Register(event.ETredisKeyClick, "rdm", rdmObj)
}

func (rdm *Rdm) Build() fyne.CanvasObject {
	rdmObj.rdmkl = NewRdmKeylist(rdm.ui.Window)
	// left
	left := rdm.buildLeft()

	// right
	right := rdm.buildRight()

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, left, nil), left, right)
}

func (rdm *Rdm) buildLeft() fyne.CanvasObject {
	topSelect := NewRdmConnections(rdm.ui.Window, rdm.ui.conf).Build()
	topContainer := widget.NewHScrollContainer(topSelect)
	topContainer.SetMinSize(fyne.NewSize(250,10))

	makeRdmKeyList := rdm.rdmkl.Build()

	bottom := rdm.rdmkl.BuildBottom()

	left := fyne.NewContainerWithLayout(layout.NewBorderLayout(topContainer, bottom, nil, nil), topContainer, makeRdmKeyList, bottom)
	return left
}

func (rdm *Rdm) buildRight() fyne.CanvasObject {
	// tabgroup
	var body = []*widget.TabItem{
		{Text: "Connection", Icon: theme.ContentAddIcon(), Content: NewRdmConnManage(rdm.ui.Window).Build()},
		{Text: "Result", Icon: theme.MenuIcon(), Content: NewRdmContent(rdm.ui.Window).Build()},
		{Text: "Command", Icon: theme.MediaPlayIcon(), Content: NewRdmCommand(rdm.ui.Window).Build()},
	}
	tabgroup := widget.NewTabContainer()
	var index = 0
	for k, v := range body {
		tabgroup.Append(widget.NewTabItemWithIcon(v.Text, v.Icon, v.Content))
		if rdm.ui.conf.UiConf.DefaultRdm == v.Text {
			index = k
		}
	}
	tabgroup.SelectTabIndex(index)

	rdm.tabgroup = tabgroup

	return tabgroup

}

func (rdm *Rdm) Clear() {

}

func (rdm *Rdm) Notify(obj event.EventObject) {
	switch obj.Et {
	case event.ETredisKeyClick:
		rdm.tabgroup.SelectTabIndex(1)
	}
}

//func (rdm *Rdm) makeRdmKeyList() fyne.CanvasObject {
//	rdsKeyList := widget.NewVScrollContainer(
//		widget.NewVBox(
//			widget.NewButton("redis_test_key_001", func() {
//
//			}),
//			widget.NewButton("redis_test_key_001", func() {
//
//			}),
//		),
//	)
//	rdsKeyList.SetMinSize(fyne.NewSize(250, 100))
//
//	// topbar
//	toolbar := widget.NewToolbar(
//		widget.NewToolbarAction(theme.SearchIcon(), func() {
//
//		}),
//		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
//
//		}),
//		widget.NewToolbarAction(theme.NavigateNextIcon(), func() {
//
//		}),
//	)
//	entry := widget.NewEntry()
//	menuBottomWithLayout := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, toolbar), entry, toolbar)
//	menuWithLayout := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, menuBottomWithLayout, nil, nil), rdsKeyList, menuBottomWithLayout)
//	return menuWithLayout
//}
func (rdm *Rdm) makeRdmContent() fyne.CanvasObject {
	return widget.NewLabel("content")
}
func (rdm *Rdm) makeRdmList() fyne.CanvasObject {
	// 列表
	list := widget.NewVScrollContainer(
		widget.NewHBox(
			widget.NewButton("test", func() {

			}),
			widget.NewButton("test222", func() {

			}),
		),
	)

	return list
}
