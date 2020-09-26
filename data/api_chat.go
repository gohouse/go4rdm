package data

import (
	"encoding/json"
	"github.com/gohouse/go4rdm/config"
	"github.com/sirupsen/logrus"
	"log"
)

type ChatApi struct {

}

func NewChatApi() *ChatApi {
	return &ChatApi{}
}

func (*ChatApi) GetChatHistory() (msg []Message) {
	bytes, err2 := httpget(config.ApiChatHistory, nil)
	if err2!=nil {
		logrus.Println("[GetChatHistory] error: ", err2.Error())
		return
	}
	log.Printf("[GetChatHistory] get data: %s\n", bytes)
	err2 = json.Unmarshal(bytes, &msg)
	if err2!=nil {
		logrus.Println("[GetChatHistory] json unmarshal error: ", err2.Error())
		return
	}

	return
}

func (*ChatApi) GetChatUsers() (u []User) {
	bytes, err2 := httpget(config.ApiChatUsers, nil)
	if err2!=nil {
		logrus.Println("[GetChatUsers] error: ", err2.Error())
		return
	}
	log.Printf("[GetChatUsers] get data: %s\n", bytes)
	err2 = json.Unmarshal(bytes, &u)
	if err2!=nil {
		logrus.Println("[GetChatUsers] json unmarshal error: ", err2.Error())
		return
	}

	return
}
