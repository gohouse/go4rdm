package customewidget

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"sync"
)

type EntrySearch struct {
	*widget.Entry
	lock *sync.Mutex
	evts map[fyne.KeyName]func()
}

func NewEntrySearch() *EntrySearch {
	search := widget.NewEntry()
	search.SetPlaceHolder(" Search")
	return &EntrySearch{Entry: search, evts: make(map[fyne.KeyName]func()), lock: &sync.Mutex{}}
}

//func (s *EntrySearch) CreateRenderer() fyne.WidgetRenderer {
//	return widget.Renderer(s.Entry)
//}

func (s *EntrySearch) CreateRenderer() fyne.WidgetRenderer {
	return s.Entry.CreateRenderer()
}

func (s *EntrySearch) TypedKey(k *fyne.KeyEvent) {
	if v, ok := s.evts[k.Name]; ok {
		v()
	}
	s.Entry.TypedKey(k)
}

func (s *EntrySearch) FocusGained() {
	s.Entry.FocusGained()
}
func (s *EntrySearch) OnTypedKey(k *fyne.KeyEvent, f func()) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.evts[k.Name] = f
}
