package ui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/gohouse/go4rdm/fmcolor"
	"github.com/gohouse/go4rdm/config"
	"github.com/gohouse/go4rdm/data"
	"github.com/gohouse/go4rdm/event"
	"log"
	"sync"
)

type QADetail struct {
	window fyne.Window
	replyBox *widget.Box
	qaData *data.Qa
	title *widget.Label
	replyTitleBtn *widget.Button
}
var onceQADetail sync.Once
var onceQADetailObj *QADetail
func NewQADetail(window fyne.Window) *QADetail {
	// id, err := machineid.ID()
	onceQADetail.Do(func() {
		onceQADetailObj = &QADetail{}
		onceQADetailObj.init()
	})
	onceQADetailObj.window = window
	return onceQADetailObj
}

func (qa *QADetail) init() {
	event.NewEvent().Register(event.ETqaClick, "qaDetail", onceQADetailObj)
}
func (qa *QADetail) Build() fyne.CanvasObject {
	//title := widget.NewLabel("title")
	title := widget.NewLabel("no answer")
	title.Wrapping = fyne.TextWrapWord
	title.Alignment = fyne.TextAlignCenter
	qa.title = title
	icon := widget.NewButtonWithIcon("", theme.MailReplyIcon(), func() {
		log.Println(qa.qaData.Id)
		qa.buildReplyForm(&data.QaReply{
			QaId:     qa.qaData.Id,
		})
	})
	qa.replyTitleBtn = icon
	icon.Hide()
	replyBoxTitle := widget.NewVBox(layout.NewSpacer(), icon)
	withLayout := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, replyBoxTitle), title, replyBoxTitle)

	var replyBox = widget.NewVBox(widget.NewLabel("no reply"))
	qa.replyBox = replyBox

	replyGroup := widget.NewGroup("reply", replyBox)

	//pageNext := widget.NewButtonWithIcon("Next", theme.NavigateNextIcon(), func() {
	//})
	//pageNext.IconPlacement = widget.ButtonIconTrailingText
	//page := widget.NewHBox(
	//	widget.NewButtonWithIcon("Prev", theme.NavigateBackIcon(), func() {
	//
	//	}),
	//	widget.NewLabel("1/20"),
	//	pageNext,
	//	)

	vBox := widget.NewVBox(withLayout, replyGroup,
		//widget.NewHBox(layout.NewSpacer(),page,layout.NewSpacer()),
		)

	//replyEntryGroup := widget.NewGroup("reply", qa.buildEntry())


	return widget.NewVScrollContainer(vBox)
}
//func (qa *QADetail) buildEntry()fyne.CanvasObject {
//	addr := widget.NewMultiLineEntry()
//	addr.SetPlaceHolder("please type a new question for submit ~~~")
//	addr.Wrapping = fyne.TextWrapWord
//	entryScroller := widget.NewHScrollContainer(addr)
//
//	icon := widget.NewButtonWithIcon("submit", theme.DocumentSaveIcon(), func() {
//
//	})
//	box := widget.NewVBox(layout.NewSpacer(), icon)
//
//	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil,nil,nil,box),entryScroller,box)
//}

func (qa *QADetail) Notify(evt event.EventObject) {
	switch evt.Et {
	case event.ETqaClick:
		if v,ok := evt.Obj.(*data.Qa);ok {
			list := data.NewApi().GetQaReplyList(v)
			qa.title.SetText(v.Content)
			qa.qaData = v
			qa.buildReplyList(list)
			qa.replyTitleBtn.Show()
		}
	}
}

func (qa *QADetail) buildReplyForm(qr *data.QaReply) {
	title := widget.NewLabel("")
	title.Alignment = fyne.TextAlignCenter
	title.Wrapping = fyne.TextWrapBreak

	val := widget.NewMultiLineEntry()
	val.Wrapping = fyne.TextWrapWord
	val.SetPlaceHolder("please input reply content")
	container := widget.NewVScrollContainer(val)
	container.SetMinSize(fyne.NewSize(400, 200))

	dialog.ShowCustomConfirm("reply", "Submit", "Cancel", container, func(b bool) {
		if !b {
			return
		}

		if val.Text == "" {
			return
		}

		qr.Nickname = config.NewConfig().Nickname
		qr.Email = config.NewConfig().Email
		qr.Content = val.Text
		err := data.NewApi().QaReply(qr)
		if err!=nil {
			log.Printf("reply error: %s", err.Error())
		} else {
			qa.buildReplyList(data.NewApi().GetQaReplyList(qa.qaData))
		}
	}, qa.window)
}
func (qa *QADetail) buildReplyList(datas []data.QaReply) {
	var list []fyne.CanvasObject
	for _,item := range datas{
		v := item

		var labelHead *canvas.Text
		labelHead = canvas.NewText(fmt.Sprintf("#%v %s %s", v.Id, v.Nickname, v.Time.Format("2006-01-02 15:04")), fmcolor.Normal.Info)
		labelHead.TextSize = 12
		var label = widget.NewLabel("")
		label.Wrapping = fyne.TextWrapWord
		if v.Pid == 0 {
			//label.SetText(fmt.Sprintf("#%v %s %s\n  %s", v.Id, v.Nickname, v.Time.Format("2006-01-02 15:04"), v.Content))
			label.SetText(v.Content)
		} else {
			//label.SetText(fmt.Sprintf("#%v %s %s\n  @%v %s", v.Id, v.Nickname, v.Time.Format("2006-01-02 15:04"), v.Pid, v.Content))
			label.SetText(fmt.Sprintf("@%v %s", v.Pid, v.Content))
		}
		labelBox := widget.NewVBox(labelHead, widget.NewHBox(widget.NewLabel(" "),label))
		icon := widget.NewButtonWithIcon("", theme.MailReplyIcon(), func() {
			qa.buildReplyForm(&data.QaReply{
				QaId:     qa.qaData.Id,
				Pid:      v.Id,
			})
		})
		iconbox := widget.NewVBox(layout.NewSpacer(), icon)
		withLayout := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, iconbox), labelBox, iconbox)
		list = append(list, withLayout)

		//var list2 []fyne.CanvasObject
		//for j:=0;j<5; j++{
		//	label := widget.NewLabel("为啥主题这么少啊, 多来几款吧")
		//	icon := widget.NewButtonWithIcon("", theme.MailReplyIcon(), func() {
		//
		//	})
		//	box := widget.NewVBox(layout.NewSpacer(), icon)
		//	list2 = append(list2, fyne.NewContainerWithLayout(layout.NewBorderLayout(nil,nil,nil,box),label,box))
		//}
		//vBox := widget.NewVBox(list2...)
		//label2 := widget.NewLabel("    ")
		//list = append(list, fyne.NewContainerWithLayout(layout.NewBorderLayout(nil,nil,label2,nil),label2, vBox))
	}

	qa.replyBox.Children = list[:]
	qa.replyBox.Refresh()
}
func (qa *QADetail) cmdAct(entry *widget.Entry, lineEntry *widget.Entry) {
	if entry.Text == "" {
		return
	}
	//result,err := data.NewData().GetRedisCmdResult(c.currentAddr, entry.Text)
	//if err!=nil {
	//	//dialog.ShowError(err, c.window)
	//	return
	//}
	//if result != "" {
	//	var res = fmt.Sprintf("%s\n\n%s", lineEntry.Text, result)
	//	if lineEntry.Text == "" {
	//		res = fmt.Sprintf("%s", result)
	//	}
	//	lineEntry.SetText(res)
	//	entry.SetText("")
	//}
}

//func (qa *QADetail) buildRight()fyne.CanvasObject {
//	return NewChat().Build()
//}