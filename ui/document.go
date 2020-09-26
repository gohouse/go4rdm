package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/gohouse/go4rdm/documents"
)
type Document struct {
	ui *UI
}

func NewDocument(ui *UI) *Document {
	return &Document{ui: ui}
}
func (home *Document) Build() fyne.CanvasObject {
	var tabs = map[string]func() string {
		"string": documents.GetDocString,
		"hash": documents.GetDocHash,
		"set": documents.GetDocSet,
		"zset": documents.GetDocZset,
		"list": documents.GetDocList,
	}

	//cmds := widget.NewLabel(documents.GetAllCommands())
	cmds := widget.NewMultiLineEntry()
	cmds.SetText(documents.GetAllCommands())
	cmds.Disable()
	cmds.Wrapping = fyne.TextWrapWord
	scrollContainer := widget.NewVScrollContainer(cmds)
	container := widget.NewTabContainer(
		widget.NewTabItem("redis document", widget.NewLabel("this is redis document")),
		widget.NewTabItem("commands", scrollContainer),
		)

	for k,v := range tabs {
		cmds := widget.NewMultiLineEntry()
		cmds.SetText(v())
		cmds.Disable()
		//cmds := widget.NewLabel(v())
		cmds.Wrapping = fyne.TextWrapWord
		scrollContainer := widget.NewVScrollContainer(cmds)

		container.Append(
			widget.NewTabItem(k, scrollContainer))
	}

	//container.SetTabLocation(widget.TabLocationTrailing)
	return container
}

