package ui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/gohouse/go4rdm/data"
	"github.com/gohouse/go4rdm/event"
	"log"
	"net/url"
)

type QA struct {
	ui *UI
	tabContailner *widget.TabContainer
	qalistBox *fyne.Container
}

func NewQA(ui *UI) *QA {
	// id, err := machineid.ID()
	return &QA{ui: ui}
}

func (qa *QA) Build() fyne.CanvasObject {
	left := qa.buildLeft()

	//right := qa.buildRight()

	//return fyne.NewContainerWithLayout(layout.NewGridLayout(2),left,right)
	return left

	//return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil,nil,nil,right), left,right)
}

func (qa *QA) buildLeft() fyne.CanvasObject {
	parse, _ := url.Parse("https://github.com/maxsky/Yahei-Monaco-Hybrid-Font/raw/master/YaHeiMonacoHybrid.ttf")
	container := widget.NewTabContainer(
		widget.NewTabItem("Question", qa.buildQuestion()),
		widget.NewTabItem("Answer", NewQADetail(qa.ui.Window).Build()),
		widget.NewTabItem("Chinese Font Download", widget.NewHyperlink("中文字体下载(download YaHeiMonacoHybrid.ttf)", parse)),
	)

	qa.tabContailner = container

	return container
}

func (qa *QA) buildQuestion() fyne.CanvasObject {
	search := widget.NewEntry()
	search.SetPlaceHolder(" submit a new question")
	//scrollContainer := customewidget.NewHScrollContainer(search)
	//search.OnChanged = func(s string) {
	//	scrollContainer.ScrollToEnd()
	//}
	search.OnChanged = func(s string) {
		if len(s) > 80 {
			search.MultiLine = true
			search.Wrapping = fyne.TextWrapWord
		} else {
			search.MultiLine = false
		}
	}
	searchicon := widget.NewButtonWithIcon("Submit", theme.DocumentSaveIcon(), func() {
		if search.Text == "" {
			return
		}
		err := data.NewApi().QaAdd(&data.Qa{
			Content:  search.Text,
		})
		if err != nil {
			dialog.ShowError(err, qa.ui.Window)
			return
		}
		search.SetText("")

		qa.buildQuestionList()
	})
	box := widget.NewVBox(layout.NewSpacer(),searchicon)
	searchiconwithLayout := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, box), search, box)

	// question list
	//qalist := qa.buildQuestionList()
	qalistBox := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil,nil,nil,nil), widget.NewLabel("no data"))
	qa.qalistBox = qalistBox
	qa.buildQuestionList()

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(searchiconwithLayout, nil, nil, nil), searchiconwithLayout, qalistBox)
}
func (qa *QA) buildQuestionList() {
	log.Println("[buildQuestionList] 开始请求数据...")
	var datas = data.NewApi().GetQaList()
	list := widget.NewList(
		func() int {
			return len(datas)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("q&a")
		},
		func(index int, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(fmt.Sprintf("%v. %s",datas[index].Id, datas[index].Content))
		},
	)
	list.OnItemSelected = func(index int) {
		qa.tabContailner.SelectTabIndex(1)
		go event.Produce(event.ETqaClick, &datas[index])
	}

	qa.qalistBox.Objects = []fyne.CanvasObject{list}
	qa.qalistBox.Refresh()
}
func (qa *QA) cmdAct(entry *widget.Entry) {
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

//func (qa *QA) buildRight()fyne.CanvasObject {
//	return NewChat().Build()
//}
