package ui

import (
	"encoding/json"
	"errors"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/denisbrodbeck/machineid"
	"github.com/gohouse/go4rdm/data"
	"github.com/gohouse/t"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"time"
)

type Chat struct {
	lineEntry *widget.Entry
	chatEntry *widget.Entry
	ui        *UI
	conn      net.Conn
	listBox   *fyne.Container
}

func NewChat(ui *UI) *Chat {
	return &Chat{ui: ui}
}

func (c *Chat) Build() fyne.CanvasObject {
	userList := c.buildUserList()
	chat := c.buildChat()
	chatBox := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, userList, nil), userList, chat)
	container := widget.NewTabContainer(
		widget.NewTabItem("chat room", chatBox),
		widget.NewTabItem("information", c.buildInformation()),
	)

	return container
}
func (c *Chat) buildUserList() fyne.CanvasObject {
	var dataarr = []string{
		"no user",
	}
	log.Println("构建user列表:", dataarr)
	list := widget.NewList(func() int {
		return len(dataarr)
	}, func() fyne.CanvasObject {
		return widget.NewLabel("nil")
	}, func(index int, item fyne.CanvasObject) {
		//item.(*fyne.Container).Objects[1].(*widget.Label).SetText(drkl.Result[index])
		item.(*widget.Label).SetText(dataarr[index])
	})
	list.OnItemSelected = func(index int) {
		//rdmkl.currentWidget.Refresh()
		//go event.Produce(event.ETredisKeyClick, data[index])
		c.chatEntry.SetText(fmt.Sprintf("@%s %s", dataarr[index], c.chatEntry.Text))
	}
	listBox := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil), list)
	c.listBox = listBox

	// paginate
	page := widget.NewToolbar(
		widget.NewToolbarAction(theme.NavigateBackIcon(), func() {

		}),
		widget.NewToolbarAction(theme.NavigateNextIcon(), func() {

		}),
	)

	// oper
	btn := widget.NewButton("conn", func() {
		if c.conn != nil {
			logrus.Error("chat reconnect by hand, conn already connected")
			return
		}
		logrus.Info("chat reconnect by hand")
		c.lineEntry.SetText("")
		go c.buildTcpConnection()
	})
	btn2 := widget.NewButton("close", func() {
		logrus.Info("chat close connect by hand")
		c.lineEntry.SetText("")
		if c.conn != nil {
			c.conn.Close()
			c.conn = nil
		}
	})
	btn3 := widget.NewButton("msg", func() {
		logrus.Info("chat get history msg by hand")
		c.lineEntry.SetText("")
		c.getMsgHistory()
	})

	operBox := widget.NewVBox(btn3, btn, btn2, page)

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, operBox, nil, nil), listBox, operBox)
}
func (c *Chat) buildInformation() fyne.CanvasObject {
	return widget.NewLabel("infomation")
}
func (c *Chat) buildChat() fyne.CanvasObject {
	lineEntry := widget.NewMultiLineEntry()
	lineEntry.Wrapping = fyne.TextWrapWord
	lineEntry.SetPlaceHolder("chat history")
	lineEntry.Disable()
	c.lineEntry = lineEntry
	container := widget.NewVScrollContainer(lineEntry)
	lineEntry.OnChanged = func(s string) {
		container.Refresh()
		container.ScrollToBottom()
	}
	//container.SetMinSize(fyne.NewSize(300, 100))

	chatEntry := widget.NewEntry()
	chatEntry.SetPlaceHolder("chat content")
	c.chatEntry = chatEntry
	chatcontainer := widget.NewHScrollContainer(chatEntry)
	chatEntry.OnCursorChanged = func() {
		if len(chatEntry.Text) > 80 {
			chatEntry.MultiLine = true
			chatEntry.Wrapping = fyne.TextWrapWord
		} else {
			chatEntry.MultiLine = false
		}
	}

	searchicon := widget.NewButtonWithIcon("send", theme.MailSendIcon(), func() {
		if chatEntry.Text == "" {
			return
		}
		if c.conn == nil {
			dialog.ShowError(errors.New("please connect first"), c.ui.Window)
			return
		}
		c.sendMsg(data.Message{Content: chatEntry.Text})
		chatEntry.SetText("")
	})
	searchiconBox := widget.NewVBox(layout.NewSpacer(), searchicon)
	searchiconwithLayout := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, searchiconBox), chatcontainer, searchiconBox)

	// get history msg for once in init
	go func() {
		time.Sleep(3 * time.Second)
		//c.getMsgHistory()
		c.buildTcpConnection()
	}()

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, searchiconwithLayout, nil, nil), container, searchiconwithLayout)
}

func (c *Chat) buildTcpConnection() {
	c.conn = nil
	var tcpAddrStr = ":9394"
	// 向服务器拨号
	conn, err := net.Dial("tcp", tcpAddrStr)
	if err != nil {
		log.Printf("Dial to server failed: %v\n", err)
		return
	}
	defer conn.Close()
	c.conn = conn

	// 获取历史消息
	c.getMsgHistory()

	// 认证啊
	err = c.sendMsg(data.Message{MessageType: data.MTauth, Content: fmt.Sprintf("system: %s connected to the room", c.ui.conf.Nickname)})
	if err != nil {
		c.conn.Close()
		dialog.ShowError(err, c.ui.Window)
		return
	}

	// heart beat
	go c.heartBeat()

	// 循环接收数据
	buf := make([]byte, 1024)
	for {
		length, err := conn.Read(buf)
		if err != nil {
			logrus.Printf("recv server msg failed: %v\n", err)
			break
		}

		if length == 0 {
			continue
		}
		//fmt.Println(string(buf[0:length]))
		// 解析数据
		var msg data.Message
		err = json.Unmarshal(buf[0:length], &msg)
		if err != nil {
			logrus.Errorf("receive chat data Unmarshal error: %s", err.Error())
			continue
		}
		logrus.Infof("receive chat message type: %v", msg.MessageType)
		switch msg.MessageType {
		case data.MTuserList:
			// update user list
			c.updateUserList(msg)
		default:
			// 写入content
			c.updateMessage([]data.Message{msg})
		}
	}
}

func (c *Chat) heartBeat() {
	ticker := time.NewTicker(8 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.conn.Write(t.New(data.Message{
				MessageType: data.MTheartBeat,
			}).Bytes())
		}
	}
}

//func (c *Chat) getUserList() {
//	users := data.NewChatApi().GetChatUsers()
//	list := widget.NewList(func() int {
//		return len(users)
//	}, func() fyne.CanvasObject {
//		return widget.NewLabel("nil")
//	}, func(index int, item fyne.CanvasObject) {
//		//item.(*fyne.Container).Objects[1].(*widget.Label).SetText(drkl.Result[index])
//		item.(*widget.Label).SetText(users[index].Nickname)
//	})
//	list.OnItemSelected = func(index int) {
//		//rdmkl.currentWidget.Refresh()
//		//go event.Produce(event.ETredisKeyClick, data[index])
//		c.chatEntry.SetText(fmt.Sprintf("@%s %s", users[index].Nickname, c.chatEntry.Text))
//	}
//
//	c.listBox.Objects = []fyne.CanvasObject{list}
//	c.listBox.Refresh()
//}
func (c *Chat) sendMsg(msg data.Message) error {
	if c.conn != nil {
		id, _ := machineid.ID()
		msg.User = data.User{
			Id:       id,
			Email:    c.ui.conf.Email,
			Nickname: c.ui.conf.Nickname,
		}
		msg.GroupId = "1"
		_, err := c.conn.Write(t.New(msg).Bytes())
		if err != nil {
			return err
		}
	}
	return nil
}
func (c *Chat) updateUserList(msgs data.Message) {
	if msgs.ExtraInfo == nil {
		return
	}
	logrus.Infof("update user list data: %#v", msgs.ExtraInfo)
	var users []data.User
	err := json.Unmarshal(t.New(msgs.ExtraInfo).Bytes(), &users)
	if err != nil {
		logrus.Errorf("userlist Unmarshal error: %s", err.Error())
		return
	}
	if len(users) > 0 {
		logrus.Infof("update userlist: %v", users)
		list := widget.NewList(func() int {
			return len(users)
		}, func() fyne.CanvasObject {
			return widget.NewLabel("nil")
		}, func(index int, item fyne.CanvasObject) {
			//item.(*fyne.Container).Objects[1].(*widget.Label).SetText(drkl.Result[index])
			item.(*widget.Label).SetText(users[index].Nickname)
		})
		list.OnItemSelected = func(index int) {
			//rdmkl.currentWidget.Refresh()
			//go event.Produce(event.ETredisKeyClick, data[index])
			c.chatEntry.SetText(fmt.Sprintf("@%s %s", users[index].Nickname, c.chatEntry.Text))
		}

		c.listBox.Objects = []fyne.CanvasObject{list}
		c.listBox.Refresh()
	}
}
func (c *Chat) updateMessage(msgs []data.Message) {
	var datastr string
	for _, msg := range msgs {
		m := fmt.Sprintf("%s %s\n %s", msg.User.Nickname, msg.Time.Format("2006-01-02 15:04:05"), msg.Content)
		datastr = fmt.Sprintf("%s%s\n\n", datastr, m)
	}
	c.lineEntry.SetText(fmt.Sprintf("%s%s", c.lineEntry.Text, datastr))
}

func (c *Chat) getMsgHistory() {
	history := data.NewChatApi().GetChatHistory()

	c.updateMessage(history)
}

func (c *Chat) cmdAct(entry *widget.Entry, container *widget.ScrollContainer) {
	if entry.Text == "" {
		return
	}
	msg := fmt.Sprintf("%s %s\n %s", "张三", time.Now().Format("2006-01-02 15:04:05"), entry.Text)
	c.lineEntry.SetText(fmt.Sprintf("%s%s\n\n", c.lineEntry.Text, msg))
	container.Refresh()
	container.ScrollToBottom()
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
