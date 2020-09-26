package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/gohouse/go4rdm/data"
	"github.com/gohouse/go4rdm/event"
	"github.com/gohouse/t"
	"log"
	"sync"
)

type RdmKeylist struct {
	lock *sync.Mutex
	window            fyne.Window
	currentWidget     *fyne.Container
	currentEntry      *widget.Entry
	currentSelected   string
	cursor            uint64
	currentButtonList []*widget.Button
	bottom fyne.CanvasObject
}

var rdmklOnce sync.Once
var rdmklObj *RdmKeylist

func NewRdmKeylist(window fyne.Window) *RdmKeylist {
	rdmklOnce.Do(func() {
		rdmklObj = &RdmKeylist{lock: &sync.Mutex{}}
		rdmklObj.init()
	})
	rdmklObj.window = window
	return rdmklObj
}
func (rdmkl *RdmKeylist) init() {
	event.NewEvent().Register(event.ETconnectionSelect, "keylist", rdmklObj)
	event.NewEvent().Register(event.ETredisKeyDelete, "keylist", rdmklObj)
}
func (rdmkl *RdmKeylist) Build() fyne.CanvasObject {
	withLayout := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil,nil,nil,nil), widget.NewLabel("no data"))
	rdmkl.currentWidget = withLayout
	return withLayout
}

func (rdmkl *RdmKeylist) BuildBottom() fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("keywords")
	entryContainer := widget.NewHScrollContainer(entry)
	entryContainer.SetMinSize(fyne.NewSize(100,20))
	rdmkl.currentEntry = entry

	newSelect := widget.NewSelect([]string{"exact", "vague"}, func(s string) {

	})
	newSelect.Selected = "exact"
	newSelect.PlaceHolder = "type"

	// topbar
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.SearchIcon(), func() {
			var drkl = data.DataRedisKeyList{
				Addr:      rdmkl.currentSelected,
				Match:     entry.Text,
				MatchType: newSelect.Selected,
				Cursor:    0,
				//Result: nil,
			}
			rdmkl.rebuildList(&drkl)
		}),
		widget.NewToolbarAction(theme.CancelIcon(), func() {
			entry.SetText("")
		}),
		widget.NewToolbarAction(theme.NavigateNextIcon(), func() {
			var drkl = data.DataRedisKeyList{
				Addr:      rdmkl.currentSelected,
				Match:     entry.Text,
				MatchType: newSelect.Selected,
				Cursor:    rdmkl.cursor,
				//Result: nil,
			}
			rdmkl.rebuildList(&drkl)
		}),
	)
	bottom := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, newSelect, toolbar), newSelect, entryContainer, toolbar)

	rdmkl.bottom = bottom
	bottom.Hide()

	return bottom
}

func (rdmkl *RdmKeylist) Notify(evt event.EventObject) {
	log.Println(evt)
	switch evt.Et {
	case event.ETconnectionSelect:
		rdmkl.bottom.Show()
		key := t.New(evt.Obj).String()
		var drkl = data.DataRedisKeyList{
			Addr: key,
		}
		rdmkl.rebuildList(&drkl)
	case event.ETredisKeyDelete:
		if v, ok := evt.Obj.(*data.DataRedisContent); ok {
			err := data.NewData().DeleteRedisKey(v)
			if err != nil {
				dialog.ShowError(err, rdmkl.window)
				return
			}
			dialog.ShowInformation("success", "delete success", rdmkl.window)
			rdmkl.rebuildList(&data.DataRedisKeyList{
				Addr: rdmkl.currentSelected,
			})
		}
	}
}
func (rdmkl *RdmKeylist) rebuildList(drkl *data.DataRedisKeyList) {
	rdmkl.lock.Lock()
	defer rdmkl.lock.Unlock()
	defer func() {
		if err := recover(); err != nil {
			log.Println("[panic recover]: ", err)
		}
	}()
	err := data.NewData().GetRedisKeyList(drkl)
	log.Println("errror 232332:", err)

	if err != nil {
		dialog.ShowError(err, rdmkl.window)
		return
	}
	//rdmkl.currentWidget.Children = []fyne.CanvasObject{}
	//rdmkl.buildKeyListWidgetWithButton(drkl)
	rdmkl.buildKeyListWidgetWithList(drkl.Result[:])
	//for _, v := range rdmkl.currentButtonList {
	//	rdmkl.currentWidget.Children = append(rdmkl.currentWidget.Children, v)
	//}
	rdmkl.cursor = drkl.Cursor
	rdmkl.currentSelected = drkl.Addr
	rdmkl.currentWidget.Refresh()
}

func (rdmkl *RdmKeylist) buildKeyListWidgetWithButton(drkl *data.DataRedisKeyList) {
	rdmkl.currentButtonList = []*widget.Button{}
	for _, v := range drkl.Result {
		tmp := v
		tmp2 := v
		if len(v) > 26 {
			tmp = tmp[:26-3] + "..."
		}
		var btnKey = &widget.Button{}
		rdmkl.currentButtonList = append(rdmkl.currentButtonList, btnKey)
		btnKey.Text = tmp
		btnKey.Alignment = widget.ButtonAlignLeading
		btnKey.OnTapped = func() {
			for _, v := range rdmkl.currentButtonList {
				v.Style = widget.DefaultButton
			}
			btnKey.Style = widget.PrimaryButton
			rdmkl.currentWidget.Refresh()
			go event.Produce(event.ETredisKeyClick, tmp2)
		}
	}
}

func (rdmkl *RdmKeylist) buildKeyListWidgetWithList(data []string) {
	log.Println("构建keys列表:",data)
	list := widget.NewList(func() int {
		return len(data)
	}, func() fyne.CanvasObject {
		//return fyne.NewContainerWithLayout(layout.NewHBoxLayout(), widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("nil a"))
		return widget.NewLabel("nil")
	}, func(index int, item fyne.CanvasObject) {
		//item.(*fyne.Container).Objects[1].(*widget.Label).SetText(drkl.Result[index])
		item.(*widget.Label).SetText(data[index])
	})
	list.OnItemSelected = func(index int) {
		//rdmkl.currentWidget.Refresh()
		go event.Produce(event.ETredisKeyClick, data[index])
	}
	//rdmkl.currentWidget.Length = list.Length
	//rdmkl.currentWidget.CreateItem = list.CreateItem
	//rdmkl.currentWidget.UpdateItem = list.UpdateItem
	//rdmkl.currentWidget.OnItemSelected = list.OnItemSelected
	rdmkl.currentWidget.Objects = []fyne.CanvasObject{
		list,
	}


	//rdmkl.currentWidget.Length = func() int {
	//	return len(data)
	//}
	//rdmkl.currentWidget.UpdateItem = func(index int, item fyne.CanvasObject) {
	//	//item.(*fyne.Container).Objects[1].(*widget.Label).SetText(drkl.Result[index])
	//	item.(*widget.Label).SetText(data[index])
	//}
	//rdmkl.currentWidget.CreateItem = func() fyne.CanvasObject {
	//	return widget.NewLabel("nil")
	//	//return fyne.NewContainerWithLayout(layout.NewHBoxLayout(), widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("nil a"))
	//}
	//rdmkl.currentWidget.OnItemSelected = func(index int) {
	//	//rdmkl.currentWidget.Refresh()
	//	go event.Produce(event.ETredisKeyClick, data[index])
	//}
}

func (rdmkl *RdmKeylist) Clear() {
	rdmkl.currentEntry.SetText("")
	//rdmkl.currentWidget.Children = []fyne.CanvasObject{}
	rdmkl.currentWidget.Refresh()
}
