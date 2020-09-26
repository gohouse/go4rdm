package config

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne"
	"github.com/go-redis/redis"
	"github.com/gohouse/golib/file"
	"io/ioutil"
	"os"
	"sync"
)

// {
//  "Username": "",
//  "Token": "",
//  "DefaultWindowSize": {
//    "Width": 1300,
//    "Height": 80
//  },
//  "Connections": [
//    {
//      "Addr": "",
//      "Password": "",
//      "DB": 0
//    }
//  ],
//  "Theme": "dark",
//  "Document": "",
//  "Notice": ""
//}
type RedisClient struct {
	Client   *redis.Client `json:"-"`
	Addr     string
	Password string
	DB       int
	Title    string
}
type Config struct {
	Email    string // 用户名
	Nickname string // 用户名
	Token    string // 加密key

	Connections []RedisClient // redis链接

	ApiUrl ApiUrl
	UiConf UiConf
}
type UiConf struct {
	LimitKey          int64
	LimitContent      int64
	ContentMode       string    // view / edit
	DefaultWindowSize fyne.Size // 窗口默认大小
	Theme             string    // dark,light
}
type ApiUrl struct {
	UrlForDocument string // 文档地址
	UrlForNotice   string // 公告地址
	UrlForJoke     string // 选项地址
}

var onceConf sync.Once
var onceConfObj *Config

func NewConfig() *Config {
	onceConf.Do(func() {
		onceConfObj = &Config{}
		onceConfObj.load()

		if onceConfObj.UiConf.LimitKey == 0 {
			onceConfObj.UiConf.LimitKey = 10
		}

		if onceConfObj.UiConf.LimitContent == 0 {
			onceConfObj.UiConf.LimitContent = 10
		}
	})

	return onceConfObj
}

func (conf *Config) AppendRds(rc ...RedisClient) {
	conf.Connections = append(conf.Connections, rc...)
	go conf.Save()
}

func (conf *Config) RemoveRds(addr string) {
	for k, v := range conf.Connections {
		if v.Addr == addr {
			var tmp []RedisClient
			tmp = append(conf.Connections[:k], conf.Connections[k+1:]...)
			conf.Connections = tmp[:]
			go conf.Save()
			break
		}
	}
}

func (conf *Config) getFilePath() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}
	filepath := fmt.Sprintf("%s/%s", dir, ".go4rdm")

	return filepath
}

func (conf *Config) load() {
	//var f *file.File
	if !file.FileExists(conf.getFilePath()) {
		f2, err := os.Create(conf.getFilePath())
		if err != nil {
			panic(err.Error())
		}
		f2.Close()
		return
	}
	//f = file.NewFile(conf.getFilePath())
	//readFile, err := f.ReadFile()
	readFile, err := ioutil.ReadFile(conf.getFilePath())
	if err != nil {
		panic(err.Error())
	}
	if len(readFile) > 0 {
		err = json.Unmarshal(readFile, &conf)
		if err != nil {
			panic(err.Error())
		}
	}
}

func (conf *Config) Save() {
	var f *file.File
	f = file.NewFile(conf.getFilePath())
	marshal, err := json.Marshal(conf)
	if err != nil {
		panic(err.Error())
	}
	_, err = f.Write(marshal, os.O_CREATE|os.O_WRONLY|os.O_TRUNC)
	if err != nil {
		panic(err.Error())
	}
}
