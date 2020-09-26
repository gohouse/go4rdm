package ui

import (
	"errors"
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/gohouse/go4rdm/event"
	"github.com/gohouse/go4rdm/resource"
	"log"
	"sync"
)

type RdmConnManage struct {
	window fyne.Window
}

var rcmOnce sync.Once
var rcmObj *RdmConnManage

func NewRdmConnManage(window fyne.Window) *RdmConnManage {
	rcmOnce.Do(func() {
		rcmObj = &RdmConnManage{window: window}
	})
	return rcmObj
}

func (rcm *RdmConnManage) Build() fyne.CanvasObject {
	icon := widget.NewButtonWithIcon("delete selected address ?", theme.DeleteIcon(), func() {
		// 不要问我为啥不用注释掉, 因为不好使啊
		//dialog.NewConfirm("delete", "delete the selected address ?", func(b bool) {
		//	log.Println("RdmConnManage 发起删除事件选择: ",b)
		//	if b {
		//		log.Println("RdmConnManage 发起删除事件: ",event.ETconnectionDel)
		//		go event.Produce(event.ETconnectionDel, nil)
		//	} else {
		//		log.Println("RdmConnManage 取消删除事件: ",event.ETconnectionDel)
		//	}
		//},rcm.window)
		log.Println("RdmConnManage 发起删除事件: ",event.ETconnectionDel)
		go event.Produce(event.ETconnectionDel, nil)
		//dialog.ShowInformation("del success","", rcm.window)
	})
	//label := widget.NewLabel("delete the selected address ?")
	//box := widget.NewHBox(label, icon)
	//box := widget.NewGroup("delete connection", widget.NewHBox(layout.NewSpacer(), icon, layout.NewSpacer()))
	groupicon := widget.NewGroup("del connection", widget.NewHBox(layout.NewSpacer(),icon))

	addr := widget.NewMultiLineEntry()
	addr.SetPlaceHolder("localhost:6379")
	addr.Wrapping = fyne.TextWrapWord

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")
	form := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:   "Address",
				Widget: addr,
			},
			{
				Text:   "password",
				Widget: password,
			},
		},
		OnSubmit: func() {
			// 保存数据
			//go right.ui.Render.ProduceConnection(NewRedisClientForProduce(nil, &redis.Options{Addr: addr.Text,Password: password.Text}, title.Text),RTadd)
			if addr.Text == "" {
				dialog.ShowError(errors.New("Address needed"), rcm.window)
				return
			}
			go event.Produce(event.ETconnectionAdd, map[string]string{
				"address":addr.Text,
				"password":password.Text,
			})
			//dialog.ShowInformation("add success", addr.Text, rcm.window)

			// 清空form
			addr.SetText("")
			password.SetText("")
		},
		OnCancel: func() {
			// 清空form
			addr.SetText("")
			password.SetText("")
		},
		SubmitText: "",
		CancelText: "Clear",
	}

	group := widget.NewGroup("add connection", form)

	//logo := widget.NewVScrollContainer(canvas.NewImageFromFile("redis.png"))
	logo := widget.NewVScrollContainer(canvas.NewImageFromResource(resource.PngOfRedis))
	logo.SetMinSize(fyne.NewSize(300,300))
	//logowithLayout := fyne.NewContainerWithLayout(layout.NewGridLayout(3),layout.NewSpacer(), logo)
	logowithLayout := widget.NewHBox(layout.NewSpacer(),logo,layout.NewSpacer())

	//containerWithLayout := fyne.NewContainerWithLayout(layout.NewGridLayout(2), groupicon, group)
	containerWithLayout := widget.NewVBox(layout.NewSpacer(),groupicon, group)

	//return fyne.NewContainerWithLayout(layout.NewGridLayoutWithRows(2),logowithLayout,containerWithLayout)
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(logowithLayout,containerWithLayout,nil,nil),logowithLayout,layout.NewSpacer(),containerWithLayout)
}

func (rcm *RdmConnManage) Clear() {

}

func (rcm *RdmConnManage) Notify(obj event.EventObject) {

}
