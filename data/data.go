package data

import (
	"net"
	"time"
)

type Qa struct {
	Id       int64
	Content  string
	Email    string
	Nickname string
	Time     time.Time
}
type QaReply struct {
	Id       int64
	Content  string
	Email    string
	Nickname string
	Time     time.Time
	QaId     int64
	Pid      int64
}
type UserApi struct {
	Email          string `json:"Email,omitempty"`
	Password       string `json:"Password,omitempty"`
	PasswordNew    string `json:"PasswordNew,omitempty"`
	PasswordRepeat string `json:"PasswordRepeat,omitempty"`
	Nickname       string `json:"Nickname,omitempty"`
	Token          string `json:"Token,omitempty"`
}

type MessageType int

const (
	MTchatText MessageType = iota
	MTauth
	MTsystem
	MTuserList
	MTchatImage
	MTchatFile
	MTheartBeat
)

type Message struct {
	User        User        `json:"User,omitempty"`
	GroupId     string      `json:"GroupId,omitempty"`
	Content     string      `json:"Content,omitempty"`
	MessageType MessageType `json:"MessageType,omitempty"`
	Time        time.Time   `json:"Time,omitempty"`
	ExtraInfo   interface{} `json:"ExtraInfo,omitempty"`
}
type User struct {
	Id       string
	Email    string
	Nickname string
	Password string
	conn     net.Conn
}

type Version struct {
	Url     string // 下载地址
	Num     int64  // 升级序号, 10
	NumText string // 版本号, v0.0.1
	Notes   string // 更新内容
}
