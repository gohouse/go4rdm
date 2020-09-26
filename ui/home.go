package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/gohouse/go4rdm/resource"
	"github.com/gohouse/go4rdm/uitheme"
	"math/rand"
	"net/url"
)

type Home struct {
	ui       *UI
	loginBox *widget.Box
}

func NewHome(ui *UI) *Home {
	return &Home{ui: ui}
}
func (home *Home) Build() fyne.CanvasObject {
	var leftUse = home.buildLeft2
	intn := rand.Intn(1000)
	if intn%5 == 0 {
		leftUse = home.buildLeft
	}
	// 左边
	left := leftUse()

	// 中间
	mid := home.buildMid()

	// 右边
	right := home.buildRight()

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, left, right), left, mid, right)
}

func (home *Home) buildMid() fyne.CanvasObject {
	label := widget.NewLabel("GO4RDM")
	label1_0 := widget.NewLabel("----------------------------")
	label2 := widget.NewLabel("Go For Redis Desktop Manager")
	label3 := widget.NewLabel("Written By Golang")
	label4 := widget.NewLabel("Cross-Platform")
	label4_0 := widget.NewLabel("----------------------------")
	label5 := widget.NewLabel("Windows")
	label6 := widget.NewLabel("Linux")
	label7 := widget.NewLabel("MaxOS")
	label8 := widget.NewLabel("Android")
	label9 := widget.NewLabel("IOS")
	label9_0 := widget.NewLabel("----------------------------")
	parse, _ := url.Parse("https://github.com/gohouse/go4rdm")
	parsefyne, _ := url.Parse("https://github.com/fyne-io/fyne")
	link := widget.NewHyperlink("Go4RDM", parse)
	link2 := widget.NewHyperlink("Fyne", parsefyne)
	box := widget.NewVBox(
		layout.NewSpacer(),
		label,
		label1_0,
		label2,
		label3,
		label4,
		label4_0,
		label5,
		label6,
		label7,
		label8,
		label9,
		label9_0,
		widget.NewHBox(link, widget.NewLabel("|"), link2),
		layout.NewSpacer(),
	)

	return widget.NewHBox(layout.NewSpacer(), box, layout.NewSpacer())
}

func (home *Home) buildLeft() fyne.CanvasObject {
	// 天下风云出我辈，一入江湖岁月催。
	//皇图霸业谈笑中，不胜人生一场醉。
	//提剑跨骑挥鬼雨，白骨如山鸟惊飞。
	//尘事如潮人如水，只叹江湖几人回。

	logo := widget.NewVScrollContainer(
		canvas.NewImageFromFile("assets/sds3.png"),
		//canvas.NewImageFromResource(resource.PngOfSword),
	)
	logo.SetMinSize(fyne.NewSize(200, 250))
	logoBox := widget.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer())

	//c2 := canvas.NewImageFromFile("knife2.png")
	c2 := canvas.NewImageFromResource(resource.PngOfKnife)
	//c2 := canvas.NewImageFromResource(resource.PngOfSword)
	lable := widget.NewLabelWithStyle("\n天下风云出我辈\n\n一入江湖岁月催\n\n皇图霸业谈笑中\n\n不胜人生一场醉\n\n提剑跨骑挥鬼雨\n\n白骨如山鸟惊飞\n\n尘事如潮人如水\n\n只叹江湖几人回", fyne.TextAlignCenter, fyne.TextStyle{})
	//lable := widget.NewTextGridFromString("G\nO\n4\nR\nD\nM")
	withLayout := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil),
		c2,
		lable,
	)
	container := widget.NewVScrollContainer(withLayout)
	container.SetMinSize(fyne.NewSize(250, 350))

	login := widget.NewVBox(logoBox, layout.NewSpacer(), container)
	//login := fyne.NewContainerWithLayout(layout.NewBorderLayout(logoBox,container,nil,nil),logoBox,layout.NewSpacer(),container)
	return login
}

func (home *Home) buildLeft2() fyne.CanvasObject {
	var box = []fyne.CanvasObject{
		layout.NewSpacer(),
	}
	var str = "GO4RDM"
	for i := 0; i < 6; i++ {
		gradient := canvas.NewHorizontalGradient(uitheme.ColorWhite, uitheme.ColorDark)
		gradient.StartColor = uitheme.ColorBlack
		gradient.EndColor = uitheme.ColorWhite
		//lable := widget.NewTextGridFromString("G\nO\n4\nR\nD\nM")
		label := canvas.NewText(str[i:i+1], uitheme.ColorBlack)
		label.TextSize = 50
		label.Alignment = fyne.TextAlignCenter
		label.TextStyle = fyne.TextStyle{
			Bold:      true,
			Italic:    true,
			Monospace: true,
		}
		withLayout := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil),
			gradient,
			label,
		)
		container := widget.NewVScrollContainer(withLayout)
		container.SetMinSize(fyne.NewSize(100, 95))

		box = append(box, container)
	}
	box = append(box, layout.NewSpacer())

	login := widget.NewVBox(box...)
	//login := fyne.NewContainerWithLayout(layout.NewBorderLayout(logoBox,container,nil,nil),logoBox,layout.NewSpacer(),container)
	return login
}

func (home *Home) buildRight() fyne.CanvasObject {
	logo := widget.NewVScrollContainer(
		//canvas.NewImageFromFile("logo.png"),
		canvas.NewImageFromResource(resource.PngOfLogo),
	)
	logo.SetMinSize(fyne.NewSize(300, 300))

	var loginBox *widget.Box
	if home.ui.conf.Token == "" {
		loginBox = widget.NewVBox(NewHomeLogin(home).Build())
	} else {
		loginBox = widget.NewVBox(home.buildLoggedIn())
	}
	home.loginBox = loginBox

	return widget.NewVBox(logo, layout.NewSpacer(), loginBox)
}

func (home *Home) buildLoggedIn() fyne.CanvasObject {
	nickname := widget.NewLabel(home.ui.conf.Nickname)
	nickname.Alignment = fyne.TextAlignTrailing
	email := widget.NewLabel(home.ui.conf.Email)
	email.Alignment = fyne.TextAlignTrailing
	btn := widget.NewButton("Logout", func() {
		//go event.Produce(event.ETapiLogout, nil)
		home.loginBox.Children = []fyne.CanvasObject{NewHomeLogin(home).Build()}
		home.loginBox.Refresh()
		home.ui.conf.Token = ""
		home.ui.conf.Email = ""
		home.ui.conf.Nickname = ""
		home.ui.conf.Save()
	})
	return widget.NewVBox(nickname, email, widget.NewHBox(
		layout.NewSpacer(), btn,
	))
}
