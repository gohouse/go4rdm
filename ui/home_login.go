package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"github.com/gohouse/e"
	"github.com/gohouse/go4rdm/data"
	"github.com/gohouse/go4rdm/event"
	"github.com/sirupsen/logrus"
	"sync"
)

type HomeLogin struct {
	home *Home
}

var onceHomeLogin sync.Once
var loginObj *HomeLogin

func NewHomeLogin(home *Home) *HomeLogin {
	onceHomeLogin.Do(func() {
		loginObj = &HomeLogin{}
		loginObj.init()
	})
	loginObj.home = home

	return loginObj
}

func (l *HomeLogin) init() {
	event.NewEvent().Register(event.ETapiLogin, "HomeLogin", loginObj)
	event.NewEvent().Register(event.ETapiRegister, "HomeLogin", loginObj)
	event.NewEvent().Register(event.ETapiPasswordReset, "HomeLogin", loginObj)
}
func (l *HomeLogin) Notify(evt event.EventObject) {
	switch evt.Et {
	case event.ETapiLogin:
		logrus.Infof("login param: %#v", evt.Obj)
		if v, ok := evt.Obj.(data.UserApi); ok {
			l.doLogin(&v, data.NewApiUser().Login)
		}
	case event.ETapiRegister:
		logrus.Infof("login param: %#v", evt.Obj)
		if v, ok := evt.Obj.(data.UserApi); ok {
			l.doLogin(&v, data.NewApiUser().Register)
		}
	case event.ETapiPasswordReset:
		logrus.Infof("login param: %#v", evt.Obj)
		if v, ok := evt.Obj.(data.UserApi); ok {
			l.doLogin(&v, data.NewApiUser().PasswordReset)
		}
	}
}

func (l *HomeLogin) doLogin(v *data.UserApi, f func(*data.UserApi) e.Error) {
	err := f(v)
	if err != nil {
		logrus.Errorf("http login error: %#v", err.Error())
		dialog.ShowError(err, l.home.ui.Window)
		return
	}
	logrus.Infof("get login respose info: %#v", v)

	l.home.ui.conf.Email = v.Email
	l.home.ui.conf.Nickname = v.Nickname
	l.home.ui.conf.Token = v.Token
	l.home.ui.conf.Save()

	l.home.loginBox.Children = []fyne.CanvasObject{l.home.buildLoggedIn()}
	l.home.loginBox.Refresh()
}

func (l *HomeLogin) Build() fyne.CanvasObject {
	tab := widget.NewTabContainer(
		widget.NewTabItem("Login", l.buildLogin()),
		widget.NewTabItem("Register", l.buildRegister()),
		widget.NewTabItem("Reset", l.buildPasswordReset()),
	)
	//tab.SetTabLocation(widget.TabLocationLeading)

	return tab
}
func (l *HomeLogin) buildLogin() fyne.CanvasObject {
	addr := widget.NewEntry()
	addr.SetPlaceHolder("Email")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")

	nickname := widget.NewEntry()
	nickname.SetPlaceHolder("Nickname")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:   "Email*",
				Widget: addr,
			},
			{
				Text:   "Password*",
				Widget: password,
			},
			{
				Text:   "Nickname",
				Widget: nickname,
			},
		},
		OnSubmit: func() {
			if addr.Text == "" || password.Text == "" {
				return
			}
			// 保存数据
			go event.Produce(event.ETapiLogin, data.UserApi{
				Email:    addr.Text,
				Password: password.Text,
				Nickname: nickname.Text,
			})

			//// 清空form
			//addr.SetText("")
			//password.SetText("")
			//nickname.SetText("")
		},
		SubmitText: "Login",
	}

	return form
}
func (l *HomeLogin) buildRegister() fyne.CanvasObject {
	addr := widget.NewEntry()
	addr.SetPlaceHolder("Email")

	nickname := widget.NewEntry()
	nickname.SetPlaceHolder("Nickname")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")

	//passwordRepeat := widget.NewPasswordEntry()
	//passwordRepeat.SetPlaceHolder("Password Repeat")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:   "Email",
				Widget: addr,
			},
			{
				Text:   "Password",
				Widget: password,
			},
			//{
			//	Text:   "Password",
			//	Widget: passwordRepeat,
			//},
			{
				Text:   "Nickname",
				Widget: nickname,
			},
		},
		OnSubmit: func() {
			if addr.Text == "" || password.Text == "" || nickname.Text == "" {
				return
			}
			// 保存数据
			go event.Produce(event.ETapiRegister, data.UserApi{
				Email:    addr.Text,
				Password: password.Text,
				Nickname: nickname.Text,
			})

			//// 清空form
			//addr.SetText("")
			//password.SetText("")
			//nickname.SetText("")
			////passwordRepeat.SetText("")
		},
		SubmitText: "Register",
	}

	return form
}
func (l *HomeLogin) buildPasswordReset() fyne.CanvasObject {
	addr := widget.NewEntry()
	addr.SetPlaceHolder("Email")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")

	passwordNew := widget.NewPasswordEntry()
	passwordNew.SetPlaceHolder("New Password")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:   "Email",
				Widget: addr,
			},
			{
				Text:   "Password",
				Widget: password,
			},
			{
				Text:   "Password",
				Widget: passwordNew,
			},
		},
		OnSubmit: func() {
			if addr.Text == "" || password.Text == "" || passwordNew.Text == "" {
				return
			}
			// 保存数据
			go event.Produce(event.ETapiPasswordReset, data.UserApi{
				Email:       addr.Text,
				Password:    password.Text,
				PasswordNew: passwordNew.Text,
			})

			//// 清空form
			//addr.SetText("")
			//password.SetText("")
			//passwordNew.SetText("")
		},
		SubmitText: "Password Reset",
	}

	return form
}
