package chat

import (
	"net"
	"time"
)

type User struct {
	Addr     *net.Addr
	Id       string
	Email    string
	Password string
}

type Group struct {
	Name string
	Id string
	Users []*User
	MsgHitory []*Message
	MsgSend chan *Message
	MsgRecv chan *Message
}

type MessageType int
type Message struct {
	User *User
	Group *Group
	Message *Message
	MessageType MessageType
	Time time.Time
}
