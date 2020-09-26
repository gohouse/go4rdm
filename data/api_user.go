package data

import (
	"encoding/json"
	"github.com/gohouse/e"
	"github.com/gohouse/go4rdm/config"
	"github.com/sirupsen/logrus"
	"log"
)

type ApiUser struct {

}

func NewApiUser() *ApiUser {
	return &ApiUser{}
}

func (*ApiUser) Login(ua *UserApi) e.Error {
	logrus.Infof("[ApiUser] Login data: %#v\n", ua)
	res, err := httppost(config.ApiUserLogin, *ua)
	if err!=nil {
		logrus.Errorf("[ApiUser] Login http error: %#v\n", err.Error())
		log.Println(err.ErrorWithStack())
		return err
	}

	logrus.Infof("[ApiUser] Login response: %s\n",res)

	err2 := json.Unmarshal(res, ua)
	if err2!=nil {
		err = e.New(err2.Error())
		logrus.Errorf("[ApiUser] Login http response Unmarshal error: %#v\n", err.Error())
		return err
	}

	return nil
}

func (*ApiUser) Register(ua *UserApi) e.Error {
	logrus.Infof("[ApiUser] Register data: %#v\n", ua)
	res, err := httppost(config.ApiUserRegister, *ua)
	if err!=nil {
		logrus.Errorf("[ApiUser] Register http error: %#v\n", err.Error())
		log.Println(err.ErrorWithStack())
		return err
	}

	logrus.Infof("[ApiUser] Register response: %s\n",res)

	err2 := json.Unmarshal(res, ua)
	if err2!=nil {
		err = e.New(err2.Error())
		logrus.Errorf("[ApiUser] Register http response Unmarshal error: %#v\n", err.Error())
		return err
	}

	return nil
}

func (*ApiUser) PasswordReset(ua *UserApi) e.Error {
	logrus.Infof("[ApiUser] PasswordReset data: %#v\n", ua)
	res, err := httppost(config.ApiUserPasswordReset, *ua)
	if err!=nil {
		logrus.Errorf("[ApiUser] PasswordReset http error: %#v\n", err.Error())
		log.Println(err.ErrorWithStack())
		return err
	}

	logrus.Infof("[ApiUser] PasswordReset response: %s\n",res)

	err2 := json.Unmarshal(res, ua)
	if err2!=nil {
		err = e.New(err2.Error())
		logrus.Errorf("[ApiUser] PasswordReset http response Unmarshal error: %#v\n", err.Error())
		return err
	}

	return nil
}