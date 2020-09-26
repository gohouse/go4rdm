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
	"github.com/gohouse/t"
	"math/rand"
	"sync"
	"time"
)

type RdmCommand struct {
	window fyne.Window
	currentAddr string
	lock *sync.Mutex
}

var cmdOnce sync.Once
var cmdObj *RdmCommand

func NewRdmCommand(window fyne.Window) *RdmCommand {
	cmdOnce.Do(func() {
		cmdObj = &RdmCommand{lock: &sync.Mutex{}}
		cmdObj.init()
	})
	cmdObj.window = window
	return cmdObj
}

func (cmd *RdmCommand) init() {
	event.NewEvent().Register(event.ETconnectionSelect, "rdmcmd", cmdObj)
}
func (cmd *RdmCommand) Build() fyne.CanvasObject {
	var imgcode = []func() string {
		getImgCodePikachu,
		getImgCodeViraus,
	}
	rand.Seed(time.Now().UnixNano())
	intn := rand.Intn(len(imgcode))
	//intn = 2
	//label := widget.NewLabel("")
	//label.Wrapping = fyne.TextWrapWord
	//label.SetText(imgcode[intn]())

	lineEntry := widget.NewMultiLineEntry()
	lineEntry.Wrapping  = fyne.TextWrapWord
	lineEntry.SetPlaceHolder(imgcode[intn]())
	lineEntry.Disable()

	//entry := customewidget.NewEntrySearch()
	entry := widget.NewEntry()
	entry.SetPlaceHolder("redis command")
	container := widget.NewVScrollContainer(lineEntry)
	//entry.OnTypedKey(&fyne.KeyEvent{fyne.KeyReturn}, func() {
	//	cmd.cmdAct(entry.Entry,lineEntry)
	//})
	icon := widget.NewButtonWithIcon("", theme.MailSendIcon(), func() {
		cmd.cmdAct(entry,lineEntry)
		container.Refresh()
		container.ScrollToBottom()
	})
	withLayout := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, icon), widget.NewHScrollContainer(entry), icon)

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, withLayout, nil, nil), container, withLayout)
}

func (cmd *RdmCommand) cmdAct(entry *widget.Entry, lineEntry *widget.Entry)  {
	cmd.lock.Lock()
	defer cmd.lock.Unlock()
	if entry.Text == "" {
		return
	}
	result,err := data.NewData().GetRedisCmdResult(cmd.currentAddr, entry.Text)
	if err!=nil {
		dialog.ShowError(err, cmd.window)
		return
	}
	if result != "" {
		var resReal = t.New(result).String()
		var originText = []rune(lineEntry.Text)
		if len(originText) > 5000 {
			originText = originText[len(originText)-5000:]
		}
		var res = fmt.Sprintf("%s\n------------------------------------------------------------\n%s\n%s", string(originText), entry.Text, resReal)
		if lineEntry.Text == "" {	// 第一次运行
			res = fmt.Sprintf("%s\n%s", entry.Text, resReal)
		}
		lineEntry.SetText(res)
		//entry.SetText("")
	}
}
func (cmd *RdmCommand) Clear()  {

}

func (cmd *RdmCommand) Notify(obj event.EventObject)  {
	switch obj.Et {
	case event.ETconnectionSelect:
		cmd.currentAddr = t.New(obj.Obj).String()
	}
}


func getImgCodePikachu() string {
	return `Pikachu..__
         $$$b  '---.__
          "$$b        '--.                          ___.---uuudP
           '$$b           '.__.------.__     __.---'      $$$$"              .
             "$b          -'            '-.-'            $$$"              .'|
               ".                                       d$"             _.'  |
                 '.   /                              ..."             .'     |
                   './                           ..::-'            _.'       |
                    /                         .:::-'            .-'         .'
                   :                          ::''\          _.'            |
                  .' .-.             .-.           '.      .'               |
                  : /'$$|           .@"$\           '.   .'              _.-'
                 .'|$u$$|          |$$,$$|           |  <            _.-'
                 | ':$$:'          :$$$$$:           '.  '.       .-'
                 :                  '"--'             |    '-.     \
                :##.       ==             .###.       '.      '.    '\
                |##:                      :###:        |        >     >
                |#'     '..''..'          '###'        x:      /     /
                 \                                   xXX|     /    ./
                  \                                xXXX'|    /   ./
                  /'-.                                  '.  /   /
                 :    '-  ...........,                   | /  .'
                 |         ':::::::'       .            |<    '.
                 |             '          |           x| \ '.:'.
                 |                         .'    /'   xXX|  ':'M'M':.
                 |    |                    ;    /:' xXXX'|  -'MMMMM:'
                 '.  .'                   :    /:'       |-'MMMM.-'
                  |  |                   .'   /'        .'MMM.-'
                  ''''                   :  ,'          |MMM<
                    |                     ''            |tbap\
                     \                                  :MM.-'
                      \                 |              .''
                       \.               '.            /
                        /     .:::::::.. :           /
                       |     .:::::::::::'.         /
                       |   .:::------------\       /
                      /   .''               >::'  /
                      '',:                 :    .'
                                           ':.:'
`
}

func getImgCodeViraus() string {
	return `/**
 **************************************************************
 *                                                            *
 *   .=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-.       *
 *    |                     ______                     |      *
 *    |                  .-"      "-.                  |      *
 *    |                 /            \                 |      *
 *    |     _          |              |          _     |      *
 *    |    ( \         |,  .-.  .-.  ,|         / )    |      *
 *    |     > "=._     | )(__/  \__)( |     _.=" <     |      *
 *    |    (_/"=._"=._ |/     /\     \| _.="_.="\_)    |      *
 *    |           "=._"(_     ^^     _)"_.="           |      *
 *    |               "=\__|IIIIII|__/="               |      *
 *    |              _.="| \IIIIII/ |"=._              |      *
 *    |    _     _.="_.="\          /"=._"=._     _    |      *
 *    |   ( \_.="_.="     '--------'     "=._"=._/ )   |      *
 *    |    > _.="                            "=._ <    |      *
 *    |   (_/                                    \_)   |      *
 *    |                                                |      *
 *    '-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-='      *
 *                                                            *
 *           LASCIATE OGNI SPERANZA, VOI CH'ENTRATE           *
 **************************************************************
 */
`
}
