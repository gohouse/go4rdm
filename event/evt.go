package event

import (
	"fmt"
	"sync"
)

type EventType int

const (
	ETconnectionAdd EventType = iota
	ETconnectionDel
	ETconnectionSelect
	ETconnectionInit

	ETredisKeyClick
	ETredisKeyDelete
	ETredisValDelete
	ETredisValUpdate

	ETqaClick
	ETqaReply

	ETapiLogin
	ETapiRegister
	ETapiPasswordReset
	ETapiLogout
)

type Event struct {
	lock   *sync.RWMutex
	Box    map[EventType]map[string]IEvent
	Events chan EventObject
}

type EventObject struct {
	Et  EventType
	Obj interface{}
}

var once sync.Once
var evt *Event

func NewEvent() *Event {
	once.Do(func() {
		evt = &Event{Events: make(chan EventObject, 11), Box: make(map[EventType]map[string]IEvent), lock: &sync.RWMutex{}}
		go evt.Listen()
	})
	return evt
}
func (evt *Event) Produce(Nt EventType, Obj interface{}) {
	evt.Events <- EventObject{
		Et:  Nt,
		Obj: Obj,
	}
}
func (evt *Event) Listen() {
	for {
		select {
		case nt := <-evt.Events:
			fmt.Println("收到了:", nt)
			for _, v := range evt.Box[nt.Et] {
				v.Notify(nt)
			}
		}
	}
}

func (evt *Event) Register(et EventType, name string, obj IEvent) {
	evt.lock.Lock()
	defer evt.lock.Unlock()
	if _, ok := evt.Box[et]; !ok {
		//	if _, ok2 := v[name]; !ok2 {
		//		evt.Box[et] = make(map[string]IEvent)
		//	}
		//} else {
		evt.Box[et] = make(map[string]IEvent)
	}
	evt.Box[et][name] = obj
	fmt.Println("现有注册者:", evt.Box)
}

func (evt *Event) Getter(et EventType, name string) IEvent {
	if v, ok := evt.Box[et]; ok {
		if _, ok2 := v[name]; ok2 {
			return evt.Box[et][name]
		}
	}
	return nil
}

func Produce(et EventType, Obj interface{}) {
	NewEvent().Produce(et, Obj)
}
