package customewidget

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type ScrollContainerForH struct {
	*widget.ScrollContainer
}

func NewHScrollContainer(content fyne.CanvasObject) *ScrollContainerForH {
	return &ScrollContainerForH{ScrollContainer: widget.NewHScrollContainer(content)}
}

func (s *ScrollContainerForH) ScrollToBegin()  {
	s.Offset.X = 0
	s.Refresh()
}

func (s *ScrollContainerForH) ScrollToEnd()  {
	s.Offset.X = s.Content.Size().Width - s.Size().Width
	s.Refresh()
}