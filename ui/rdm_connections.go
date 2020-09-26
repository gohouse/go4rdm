package ui

import (
	"errors"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"github.com/gohouse/go4rdm/config"
	"github.com/gohouse/go4rdm/event"
	"github.com/gohouse/t"
	"log"
	"sync"
)

func (rdmc *RdmConnections) init() {
	event.NewEvent().Register(event.ETconnectionAdd, "rdmconnection", rdmc)
	event.NewEvent().Register(event.ETconnectionDel, "rdmconnection", rdmc)
	event.NewEvent().Register(event.ETconnectionInit, "rdmconnection", rdmc)
}

type RdmConnections struct {
	conf *config.Config
	window                 fyne.Window
	currentSelectedAddress string
	currentArg             []string
	currentWidget          *widget.Select
	currentValue           string
}

var rdmcOnce sync.Once
var rdmcObj *RdmConnections

func NewRdmConnections(window fyne.Window, conf *config.Config) *RdmConnections {
	rdmcOnce.Do(func() {
		rdmcObj = &RdmConnections{currentArg: []string{}}
		rdmcObj.init()
	})
	rdmcObj.window = window
	rdmcObj.conf = conf


	return rdmcObj
}

func (rdmc *RdmConnections) Build() fyne.CanvasObject {
	newSelect := widget.NewSelect([]string{}, func(s string) {
		rdmc.currentSelectedAddress = s
		go event.Produce(event.ETconnectionSelect, s)
	})
	//newSelect.Options = []string{"xxx.cc", "b", "c"}

	rdmc.currentWidget = newSelect

	return newSelect
}
func (rdmc *RdmConnections) ReBuild() {
	rdmc.currentWidget.Options = rdmc.currentArg
	rdmc.currentWidget.Selected = rdmc.currentSelectedAddress
	rdmc.currentWidget.Refresh()
}

func (rdmc *RdmConnections) Clear() {
	rdmc.currentWidget.Options = []string{}
	rdmc.currentWidget.Refresh()
}
func (rdmc *RdmConnections) Notify(object event.EventObject) {
	log.Println("RdmConnections 收到事件: ",object)
	switch object.Et {
	case event.ETconnectionInit:
		if v, ok := object.Obj.([]config.RedisClient); ok {
			log.Println("333 -- 33 : ",v)
			var tmp []string
			for _, item := range v {
				tmp = append(tmp, item.Addr)
			}
			rdmc.currentArg = tmp
		}
	case event.ETconnectionAdd:
		var cur = t.New(object.Obj).MapStringT()
		var exists bool
		for _, v := range rdmc.currentArg {
			if v == cur["address"].String() {
				exists = true
				break
			}
		}
		if exists {
			dialog.ShowError(errors.New("address already exists"), rdmc.window)
			return
		}
		rdmc.currentArg = append(rdmc.currentArg, cur["address"].String())
		rdmc.conf.AppendRds(config.RedisClient{
			Addr:     cur["address"].String(),
			Password: cur["password"].String(),
		})
		dialog.ShowInformation("add success","", rdmc.window)

	case event.ETconnectionDel:
		var cur = rdmc.currentSelectedAddress
		if cur == "" {
			dialog.ShowError(errors.New("please select a connection"), rdmc.window)
			return
		}
		for k, v := range rdmc.currentArg {
			if v == cur {
				log.Println("RdmConnections 找到要删除的事件: ",cur)
				var tmp []string
				if k == len(rdmc.currentArg)-1 {
					tmp = rdmc.currentArg[:k]
				} else {
					tmp = append(rdmc.currentArg[:k], rdmc.currentArg[k+1:]...)
				}
				rdmc.currentArg = tmp[:]
				rdmc.conf.RemoveRds(v)
				if rdmc.currentSelectedAddress == cur {
					rdmc.currentSelectedAddress = ""
				}
				dialog.ShowInformation("del success","", rdmc.window)
				break
			}
		}
	}
	rdmc.ReBuild()
}
