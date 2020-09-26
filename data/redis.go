package data

import (
	"github.com/go-redis/redis"
	"github.com/gohouse/go4rdm/config"
	"log"
	"sync"
)

type Redis struct {
	*config.Config
}

var once sync.Once
var rds *Redis

func NewRedis(conf *config.Config) *Redis {
	once.Do(func() {
		rds = &Redis{
			Config: conf,
		}
	})
	return rds
}

func (rds *Redis) DeleteRds(addr string) {
	for k, v := range rds.Connections {
		if v.Addr == addr {
			tmp := append(rds.Connections[:k], rds.Connections[k+1:]...)
			rds.Connections = tmp[:]
			go rds.Config.Save()
			return
		}
	}
}

func (rds *Redis) GetRds(addr string) *redis.Client {
	defer func() {
		if err := recover(); err != nil {
			log.Println("[panic recover]: ", err)
		}
	}()
	for _, v := range rds.Connections {
		if v.Addr == addr {
			if v.Client == nil {
				log.Println("开始连接redis了111: ",rds.Connections,addr)
				v.Client = redis.NewClient(&redis.Options{Addr: v.Addr, Password: v.Password})
			}
			return v.Client
		}
	}
	return nil
}
