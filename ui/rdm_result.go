package ui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/atotto/clipboard"
	"github.com/go-redis/redis"
	"github.com/gohouse/go4rdm/config"
	"github.com/gohouse/go4rdm/data"
	"github.com/gohouse/go4rdm/event"
	"github.com/gohouse/t"
	"log"
	"sync"
)

func (rdmct *RdmContent) init() {
	event.NewEvent().Register(event.ETredisKeyClick, "rdmcontent", rdmct)
	event.NewEvent().Register(event.ETconnectionSelect, "rdmcontent", rdmct)
	event.NewEvent().Register(event.ETredisKeyDelete, "rdmcontent", rdmct)
	event.NewEvent().Register(event.ETredisValUpdate, "rdmcontent", rdmct)
	event.NewEvent().Register(event.ETredisValDelete, "rdmcontent", rdmct)
}

type RdmContent struct {
	window            fyne.Window
	actDataType       *widget.Label
	actTtl            *widget.Label
	currentWidget     *widget.Box
	currentEntry      *widget.Entry
	currentSelected   string
	currentKey        string
	currentForNewItem *currentForNewItem
	currentCursor     uint64
	cursor            uint64
	bottom            fyne.CanvasObject
	drkl *data.DataRedisContent
}
type currentForNewItem struct {
	currentKeyType *widget.Label
	fieldBox       *widget.Box
}

var rdmctOnce sync.Once
var rdmctObj *RdmContent

func NewRdmContent(window fyne.Window) *RdmContent {
	rdmctOnce.Do(func() {
		rdmctObj = &RdmContent{window: window}
		rdmctObj.init()
	})
	return rdmctObj
}

func (rdmct *RdmContent) Build() fyne.CanvasObject {
	top := rdmct.buildTop()

	var contentvBox = widget.NewVBox(
		widget.NewLabel("no data"),
	)
	rdmct.currentWidget = contentvBox
	contentBox := widget.NewVScrollContainer(
		contentvBox,
	)

	bottom := rdmct.buildBottom()

	menuWithLayout := fyne.NewContainerWithLayout(layout.NewBorderLayout(top, bottom, nil, nil), top, contentBox, bottom)
	return menuWithLayout
}

func (rdmct *RdmContent) buildBottom() fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("member(set/zset)/field(hash) for search")
	rdmct.currentEntry = entry
	entry.Disable()

	newSelect := widget.NewSelect([]string{"exact", "vague"}, func(s string) {

	})
	newSelect.Selected = "exact"
	newSelect.PlaceHolder = "type"

	// topbar
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.SearchIcon(), func() {
			var drkl = data.DataRedisContent{
				Addr:      rdmct.currentSelected,
				Key:       rdmct.currentKey,
				Match:     entry.Text,
				MatchType: newSelect.Selected,
				Cursor:    0,
				//Result: nil,
			}
			rdmct.rebuildList(&drkl)
		}),
		widget.NewToolbarAction(theme.CancelIcon(), func() {
			entry.SetText("")
		}),
		widget.NewToolbarAction(theme.NavigateNextIcon(), func() {
			var drkl = data.DataRedisContent{
				Addr:      rdmct.currentSelected,
				Key:       rdmct.currentKey,
				Match:     entry.Text,
				MatchType: newSelect.Selected,
				Cursor:    rdmct.cursor,
				//Result: nil,
			}
			rdmct.rebuildList(&drkl)
		}),
	)

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, newSelect, toolbar),
		newSelect, widget.NewHScrollContainer(entry), toolbar)
}

func (rdmct *RdmContent) buildTop() fyne.CanvasObject {
	var viewMod *widget.Button
	var editMod *widget.Button
	var modLabel *widget.Label
	modLabel = widget.NewLabel("")
	viewMod = widget.NewButtonWithIcon("", theme.VisibilityIcon(), func() {
		viewMod.Style = widget.PrimaryButton
		editMod.Style = widget.DefaultButton
		modLabel.SetText("view mod")
		config.NewConfig().UiConf.ContentMode = "view"
		rdmct.rebuildList(nil)
		viewMod.Refresh()
		editMod.Refresh()
	})
	editMod = widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), func() {
		editMod.Style = widget.PrimaryButton
		viewMod.Style = widget.DefaultButton
		modLabel.SetText("edit mod")
		config.NewConfig().UiConf.ContentMode = "edit"
		rdmct.rebuildList(nil)
		viewMod.Refresh()
		editMod.Refresh()
	})
	if config.NewConfig().UiConf.ContentMode == "edit" {
		modLabel.SetText("edit mod")
		editMod.Style = widget.PrimaryButton
		editMod.Refresh()
	} else {
		modLabel.SetText("view mod")
		viewMod.Style = widget.PrimaryButton
		viewMod.Refresh()
	}
	modBox := widget.NewHBox(viewMod, editMod, modLabel)

	dataType := widget.NewLabel("[data type]: [keyname]")
	ttl := widget.NewLabel("-1")
	rdmct.actDataType = dataType
	rdmct.actTtl = ttl

	btnDel := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		go event.Produce(event.ETredisKeyDelete, &data.DataRedisContent{
			Addr: rdmct.currentSelected,
			Key:  rdmct.currentKey,
		})
	})
	btnCopy := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		clipboard.WriteAll(rdmct.currentKey)
	})
	// confirm box
	newBox := rdmct.buildNewBox()
	btnNew := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		dialog.ShowCustomConfirm("add a new item", "Submit", "Cancel", newBox, func(b bool) {
			if !b {
				return
			}
		}, rdmct.window)
	})

	return widget.NewHBox(modBox, layout.NewSpacer(), dataType, layout.NewSpacer(), widget.NewLabel("ttl:"), ttl, btnDel, btnNew, btnCopy)
}
func (rdmct *RdmContent) buildNewBox() fyne.CanvasObject {
	title := widget.NewLabel("")
	title.Alignment = fyne.TextAlignCenter
	title.Wrapping = fyne.TextWrapBreak

	key := widget.NewMultiLineEntry()
	key.Wrapping = fyne.TextWrapWord
	key.SetPlaceHolder("please input field(hash)/score(zset)")
	vBox := widget.NewVBox(widget.NewLabel("field(hash)/score(zset)"), key)
	rdmct.currentForNewItem = &currentForNewItem{
		fieldBox:       vBox,
		currentKeyType: title,
	}

	val := widget.NewMultiLineEntry()
	val.Wrapping = fyne.TextWrapWord
	val.SetPlaceHolder("please input value/member(zset/set)")
	container := widget.NewVScrollContainer(val)
	container.SetMinSize(fyne.NewSize(400, 200))

	box := widget.NewVBox(title, vBox, widget.NewLabel("value"), container)

	return box
}

func (rdmct *RdmContent) Notify(evt event.EventObject) {
	log.Println("收到了notify: ", evt)
	switch evt.Et {
	case event.ETconnectionSelect:
		rdmct.Clear()
		rdmct.currentSelected = t.New(evt.Obj).String()
	case event.ETredisKeyClick:
		key := t.New(evt.Obj).String()
		var drkl = data.DataRedisContent{
			Addr: rdmct.currentSelected,
			Key:  key,
			//Match:  "",
			Cursor: rdmct.currentCursor,
			//Result: nil,
		}
		rdmct.currentKey = key
		rdmct.rebuildList(&drkl)
	case event.ETredisValUpdate:
		if v, ok := evt.Obj.(*data.DataRedisContent); ok {
			err := data.NewData().UpdateRedisContet(v)
			if err != nil {
				dialog.ShowError(err, rdmct.window)
				return
			}
			dialog.ShowInformation("success", "update success", rdmct.window)
			rdmct.rebuildList(&data.DataRedisContent{
				Addr: rdmct.currentSelected,
				Key:  rdmct.currentKey,
			})
		}
	case event.ETredisValDelete:
		if v, ok := evt.Obj.(*data.DataRedisContent); ok {
			err := data.NewData().DeleteRedisContet(v)
			if err != nil {
				dialog.ShowError(err, rdmct.window)
				return
			}
			dialog.ShowInformation("success", "delete success", rdmct.window)
			rdmct.rebuildList(&data.DataRedisContent{
				Addr: rdmct.currentSelected,
				Key:  rdmct.currentKey,
			})
		}
	case event.ETredisKeyDelete:
		rdmct.currentWidget.Children = []fyne.CanvasObject{}

		rdmct.currentWidget.Refresh()
	}
}
func (rdmct *RdmContent) rebuildList(drkl *data.DataRedisContent) {
	if drkl == nil {
		drkl = rdmct.drkl
	}
	rdmct.drkl = drkl
	if config.NewConfig().UiConf.ContentMode == "edit" {
		rdmct.rebuildListForEit(drkl)
	} else {
		rdmct.rebuildListForView(drkl)
	}
}
func (rdmct *RdmContent) rebuildListForEit(drkl *data.DataRedisContent) {
	err := data.NewData().GetRedisContent(drkl)
	if err != nil {
		dialog.ShowError(err, rdmct.window)
		return
	}
	log.Printf("获取到content值: %#v\n", drkl)
	//rdmct.currentEntry.SetText("")
	rdmct.currentWidget.Children = []fyne.CanvasObject{}
	rdmct.currentForNewItem.fieldBox.Hide()
	if drkl.Type != "string" {
		rdmct.currentForNewItem.currentKeyType.SetText(fmt.Sprintf("%s(%s)", drkl.Type, drkl.Key))
	} else {
		rdmct.currentForNewItem.currentKeyType.SetText(drkl.Type)
	}
	switch drkl.Type {
	case "string":
		field := widget.NewLabel(drkl.Type)
		value := widget.NewMultiLineEntry()
		value.SetText(t.New(drkl.Result).String())
		value.Wrapping = fyne.TextWrapWord

		var drklForOper = data.DataRedisContent{
			Addr: drkl.Addr,
			Key:  drkl.Key,
			Type: drkl.Type,
		}
		btn := widget.NewHBox(widget.NewButtonWithIcon("update", theme.DocumentSaveIcon(), func() {
			drklForOper.Result = value.Text
			go event.Produce(event.ETredisValUpdate, &drklForOper)
		}))
		rdmct.currentWidget.Children = append(rdmct.currentWidget.Children,
			widget.NewVBox(fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, btn), field, btn), value))
		rdmct.currentEntry.Disable()
		rdmct.currentEntry.SetText("")
	case "hash":
		rdmct.currentForNewItem.fieldBox.Show()
		if v, ok := drkl.Result.(map[string]string); ok {
			for k, v2 := range v {
				var vo = k
				label1 := widget.NewLabel("field:")
				label2 := widget.NewEntry()
				label2.SetText(vo)
				label2.Disable()
				field := widget.NewHBox(label1, label2)
				value := widget.NewMultiLineEntry()
				value.SetText(v2)
				value.Wrapping = fyne.TextWrapWord

				var drklForOper = data.DataRedisContent{
					Addr: drkl.Addr,
					Key:  drkl.Key,
					Type: drkl.Type,
				}
				btn := widget.NewHBox(
					widget.NewButtonWithIcon("del", theme.DeleteIcon(), func() {
						drklForOper.Result = map[string]string{vo: ""}
						go event.Produce(event.ETredisValDelete, &drklForOper)
					}), widget.NewButtonWithIcon("update", theme.DocumentSaveIcon(), func() {
						drklForOper.Result = map[string]string{vo: value.Text}
						go event.Produce(event.ETredisValUpdate, &drklForOper)
					}))
				rdmct.currentWidget.Children = append(rdmct.currentWidget.Children,
					widget.NewVBox(fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, btn), field, btn), value))
			}
			//rdmct.currentWidget.Refresh()
			if rdmct.currentEntry.Disabled() {
				rdmct.currentEntry.Enable()
			}
		}
	case "zset":
		rdmct.currentForNewItem.fieldBox.Show()
		if v, ok := drkl.Result.([]redis.Z); ok {
			for _, v2 := range v {
				var vo = v2
				//field := widget.NewLabel("score:" + t.New(v2.Score).String())
				//field.Wrapping = fyne.TextWrapWord
				entry := widget.NewEntry()
				entry.SetText(t.New(vo.Score).String())
				field := widget.NewHBox(widget.NewLabel("score"), entry)
				value := widget.NewMultiLineEntry()
				value.SetText(t.New(vo.Member).String())
				value.Wrapping = fyne.TextWrapWord

				var drklForOper = data.DataRedisContent{
					Addr: drkl.Addr,
					Key:  drkl.Key,
					Type: drkl.Type,
				}
				btn := widget.NewHBox(
					widget.NewButtonWithIcon("del", theme.DeleteIcon(), func() {
						drklForOper.Result = map[string]redis.Z{
							"origin": {
								Score:  vo.Score,
								Member: vo.Member,
							},
						}
						go event.Produce(event.ETredisValDelete, &drklForOper)
					}), widget.NewButtonWithIcon("update", theme.DocumentSaveIcon(), func() {
						drklForOper.Result = map[string]redis.Z{
							"origin": {
								Score:  vo.Score,
								Member: vo.Member,
							},
							"new": {
								Score:  t.New(entry.Text).Float64(),
								Member: value.Text,
							},
						}
						go event.Produce(event.ETredisValUpdate, &drklForOper)
					}))
				rdmct.currentWidget.Children = append(rdmct.currentWidget.Children,
					widget.NewVBox(fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, btn), field, btn), value))
			}
			//rdmct.currentWidget.Refresh()
			if rdmct.currentEntry.Disabled() {
				rdmct.currentEntry.Enable()
			}
		}
	case "set":
		if v, ok := drkl.Result.([]string); ok {
			for _, v2 := range v {
				var vo = v2
				field := widget.NewLabel(drkl.Type)
				value := widget.NewMultiLineEntry()
				value.SetText(vo)
				value.Wrapping = fyne.TextWrapWord

				//rdmct.currentWidget.Children = append(rdmct.currentWidget.Children, value)
				var drklForOper = data.DataRedisContent{
					Addr: drkl.Addr,
					Key:  drkl.Key,
					Type: drkl.Type,
				}
				btn := widget.NewHBox(
					widget.NewButtonWithIcon("del", theme.DeleteIcon(), func() {
						drklForOper.Result = map[string]string{
							"origin": vo,
						}
						go event.Produce(event.ETredisValDelete, &drklForOper)
					}), widget.NewButtonWithIcon("update", theme.DocumentSaveIcon(), func() {
						drklForOper.Result = map[string]string{
							"origin": vo,
							"new":    value.Text,
						}
						go event.Produce(event.ETredisValUpdate, &drklForOper)
					}))
				rdmct.currentWidget.Children = append(rdmct.currentWidget.Children,
					widget.NewVBox(fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, btn), field, btn), value))
			}
			if rdmct.currentEntry.Disabled() {
				rdmct.currentEntry.Enable()
			}
		}
	case "list":
		if v, ok := drkl.Result.([]string); ok {
			for _, v2 := range v {
				value := widget.NewMultiLineEntry()
				value.SetText(v2)
				value.Wrapping = fyne.TextWrapWord
				value.Disable()

				rdmct.currentWidget.Children = append(rdmct.currentWidget.Children, value)
			}
			rdmct.currentEntry.Disable()
			rdmct.currentEntry.SetText("")
		}
	}

	rdmct.actDataType.SetText(fmt.Sprintf("%s:%s", drkl.Type, drkl.Key))
	rdmct.actTtl.SetText(drkl.Ttl)
	rdmct.cursor = drkl.Cursor
	rdmct.currentWidget.Refresh()
}
func (rdmct *RdmContent) rebuildListForView(drkl *data.DataRedisContent) {
	err := data.NewData().GetRedisContent(drkl)
	if err != nil {
		dialog.ShowError(err, rdmct.window)
		return
	}
	log.Printf("获取到content值: %#v\n", drkl)
	//rdmct.currentEntry.SetText("")
	rdmct.currentWidget.Children = []fyne.CanvasObject{}
	rdmct.currentForNewItem.fieldBox.Hide()
	if drkl.Type != "string" {
		rdmct.currentForNewItem.currentKeyType.SetText(fmt.Sprintf("%s(%s)", drkl.Type, drkl.Key))
	} else {
		rdmct.currentForNewItem.currentKeyType.SetText(drkl.Type)
	}
	switch drkl.Type {
	case "string":
		res := widget.NewLabel(t.New(drkl.Result).String())
		res.Wrapping = fyne.TextWrapWord
		rdmct.currentWidget.Children = append(rdmct.currentWidget.Children, res)
	case "hash":
		if v, ok := drkl.Result.(map[string]string); ok {
			var showEntry = widget.NewMultiLineEntry()
			showEntry.Wrapping = fyne.TextWrapWord
			showEntry.Disable()
			var text string
			for k, v2 := range v {
				var ko = k
				var vo = v2
				text = fmt.Sprintf("%sfield: %s\n%s\n\n", text,ko, vo)
			}
			showEntry.SetText(text)
			rdmct.currentWidget.Children = append(rdmct.currentWidget.Children, showEntry)
			if rdmct.currentEntry.Disabled() {
				rdmct.currentEntry.Enable()
			}
		}
	case "zset":
		rdmct.currentForNewItem.fieldBox.Show()
		if v, ok := drkl.Result.([]redis.Z); ok {
			var showEntry = widget.NewMultiLineEntry()
			showEntry.Wrapping = fyne.TextWrapWord
			showEntry.Disable()
			var text string
			for _, v2 := range v {
				var vo = v2
				text = fmt.Sprintf("%sscore: %s\n%s\n\n", text,t.New(vo.Score).String(), t.New(vo.Member).String())
			}
			showEntry.SetText(text)
			rdmct.currentWidget.Children = append(rdmct.currentWidget.Children, showEntry)
			if rdmct.currentEntry.Disabled() {
				rdmct.currentEntry.Enable()
			}
		}
	case "set":
		if v, ok := drkl.Result.([]string); ok {
			var showEntry = widget.NewMultiLineEntry()
			showEntry.Wrapping = fyne.TextWrapWord
			showEntry.Disable()
			var text string
			for _, v2 := range v {
				var vo = v2
				text = fmt.Sprintf("%sscore: %s\n\n", text, vo)
			}
			showEntry.SetText(text)
			rdmct.currentWidget.Children = append(rdmct.currentWidget.Children, showEntry)
			if rdmct.currentEntry.Disabled() {
				rdmct.currentEntry.Enable()
			}
		}
	case "list":
		var showEntry = widget.NewMultiLineEntry()
		showEntry.Wrapping = fyne.TextWrapWord
		showEntry.Disable()
		var text string
		if v, ok := drkl.Result.([]string); ok {
			for _, v2 := range v {
				var vo = v2
				text = fmt.Sprintf("%sscore: %s\n\n", text, vo)
			}
			showEntry.SetText(text)
			rdmct.currentWidget.Children = append(rdmct.currentWidget.Children, showEntry)
			rdmct.currentEntry.Disable()
			rdmct.currentEntry.SetText("")
		}
	}

	rdmct.actDataType.SetText(fmt.Sprintf("%s:%s", drkl.Type, drkl.Key))
	rdmct.actTtl.SetText(drkl.Ttl)
	rdmct.cursor = drkl.Cursor
	rdmct.currentWidget.Refresh()
}

func (rdmct *RdmContent) Clear() {
	rdmct.currentEntry.SetText("")
	rdmct.currentWidget.Children = []fyne.CanvasObject{}
	rdmct.currentWidget.Refresh()
}
