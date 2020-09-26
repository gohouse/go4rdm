package data

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gohouse/t"
	"log"
	"regexp"
	"strings"
)

type Data struct {
	*Redis
}

func NewData() *Data {
	return &Data{rds}
}

type DataRedisKeyList struct {
	Addr      string
	Match     string
	MatchType string
	Cursor    uint64
	Result    []string
}
type DataRedisContent struct {
	Addr      string
	Key       string
	Match     string
	MatchType string
	Cursor    uint64
	Result    interface{}
	Type      string
	Ttl       string
}

func (d *Data) DeleteRedisKey(drk *DataRedisContent) (err error) {
	var rds = d.GetRds(drk.Addr)
	if rds == nil {
		return fmt.Errorf("connection fail, please retry again: %s", drk.Addr)
	}
	_, err = rds.Del(drk.Key).Result()
	return
}

func (d *Data) DeleteRedisContet(drk *DataRedisContent) (err error) {
	var rds = d.GetRds(drk.Addr)
	if rds == nil {
		return fmt.Errorf("connection fail, please retry again: %s", drk.Addr)
	}
	switch drk.Type {
	case "hash":
		if v, ok := drk.Result.(map[string]string); ok {
			for field, _ := range v {
				_, err = rds.HDel(drk.Key, field).Result()
			}
		}
	case "zset":
		if v, ok := drk.Result.(map[string]redis.Z); ok {
			_, err = rds.ZRem(drk.Key, v["origin"].Member).Result()
		}
	case "set":
		if v, ok := drk.Result.(map[string]string); ok {
			_, err = rds.SRem(drk.Key, v["origin"]).Result()
		}
	}
	return
}

func (d *Data) UpdateRedisContet(drk *DataRedisContent) (err error) {
	log.Printf("收到update的内容:%#v\n",drk)
	var rds = d.GetRds(drk.Addr)
	if rds == nil {
		return fmt.Errorf("connection fail, please retry again: %s", drk.Addr)
	}
	switch drk.Type {
	case "string":
		_, err = rds.Set(drk.Key, drk.Result, rds.TTL(drk.Key).Val()).Result()
	case "hash":
		if v, ok := drk.Result.(map[string]string); ok {
			for field, val := range v {
				_, err = rds.HSet(drk.Key, field, val).Result()
			}
		}
	case "zset":
		if v, ok := drk.Result.(map[string]redis.Z); ok {
			result, err := rds.ZRem(drk.Key, v["origin"].Member).Result()
			if err == nil && result != 0 {
				_, err = rds.ZAdd(drk.Key, v["new"]).Result()
			}
		}
	case "set":
		if v, ok := drk.Result.(map[string]string); ok {
			result, err := rds.SRem(drk.Key, v["origin"]).Result()
			log.Println("del set :",result,err)
			if err == nil && result != 0 {
				result, err = rds.SAdd(drk.Key, v["new"]).Result()
				log.Println("add set :",result,drk.Key,err,v["new"])
			}
		}
	}
	return
}

func (d *Data) GetRedisKeyList(drk *DataRedisKeyList) error {
	if drk.Addr == "" {
		return errors.New("addr needed")
	}
	var match = drk.Match
	if drk.MatchType == "vague" {
		match = fmt.Sprintf("*%s*", drk.Match)
	}
	rds := d.Redis.GetRds(drk.Addr)
	if rds == nil {
		return errors.New("connect redis fail,please retry or check the password")
	}
	result, cursor, err := rds.Scan(drk.Cursor, match, d.Redis.Config.UiConf.LimitKey).Result()
	if err != nil {
		return err
	}
	drk.Cursor = cursor
	drk.Result = result
	return nil
}

func (d *Data) GetRedisContent(drk *DataRedisContent) error {
	rds := d.Redis.GetRds(drk.Addr)
	if rds == nil {
		return errors.New("connect redis fail,please retry or check the password")
	}
	val := rds.Type(drk.Key).Val()
	drk.Type = val
	log.Printf("key:%s 的type为 %s\n", drk.Key, val)
	var match = drk.Match
	if drk.MatchType == "vague" {
		match = fmt.Sprintf("*%s*", drk.Match)
	}
	switch val {
	case "string":
		result, err := rds.Get(drk.Key).Result()
		drk.Result = result
		if err != nil {
			return err
		}
	case "list":
		result, err := rds.LRange(drk.Key, int64(drk.Cursor), int64(drk.Cursor)+d.Redis.Config.UiConf.LimitContent).Result()
		drk.Cursor = (drk.Cursor) + uint64(d.Redis.Config.UiConf.LimitContent)
		drk.Result = result
		if err != nil {
			return err
		}
	case "hash":
		result, cursor, err := rds.HScan(drk.Key, drk.Cursor, match, d.Redis.Config.UiConf.LimitContent).Result()
		if err != nil {
			return err
		}
		var hr = map[string]string{}
		for i := 0; i < len(result); i += 2 {
			hr[result[i]] = result[i+1]
		}
		drk.Cursor = cursor
		drk.Result = hr
	case "set":
		result, cursor, err := rds.SScan(drk.Key, drk.Cursor, match, d.Redis.Config.UiConf.LimitContent).Result()
		if err != nil {
			return err
		}
		drk.Cursor = cursor
		drk.Result = result
	case "zset":
		result, cursor, err := rds.ZScan(drk.Key, drk.Cursor, match, d.Redis.Config.UiConf.LimitContent).Result()
		if err != nil {
			return err
		}
		var hr []redis.Z
		for i := 0; i < len(result); i += 2 {
			hr = append(hr, redis.Z{
				Score:  t.New(result[i+1]).Float64(),
				Member: result[i],
			})
		}
		drk.Cursor = cursor
		drk.Result = hr
	}
	ttl := rds.TTL(drk.Key).Val().String()
	drk.Ttl = ttl
	return nil
}

func (d *Data) GetRedisCmdResult2(addr, cmd string) (interface{}, error) {
	if addr == "" {
		return nil, errors.New("please select a connetion")
	}
	compile := regexp.MustCompile(`\s+`)
	all := compile.ReplaceAll([]byte(cmd), []byte(" "))
	var a []interface{}
	for _, v := range strings.Split(string(all), " ") {
		a = append(a, v)
	}
	rds := d.Redis.GetRds(addr)
	if rds == nil {
		return nil, errors.New("connect redis fail,please retry or check the password")
	}
	return rds.Do(a...).Result()
}

func (d *Data) GetRedisCmdResult(addr, cmd string) (interface{}, error) {
	log.Printf("cmd: %+v\n",cmd)
	if addr == "" {
		return nil, errors.New("please select a connetion")
	}
	var resStr []interface{}
	var indexQuote int
	if strings.Contains(cmd, "'") {
		indexQuote = strings.Index(cmd, "'")
	}
	if strings.Contains(cmd, `"`) {
		indexQuote2 := strings.Index(cmd, `"`)
		if indexQuote != 0 {
			if indexQuote2 < indexQuote {
				indexQuote  = indexQuote2
			}
		} else {
			indexQuote  = indexQuote2
		}
	}
	if indexQuote > 0 {
		cmd1 := cmd[:indexQuote]
		cmd2 := cmd[indexQuote:]
		split := strings.Split(cmd1, " ")
		for _,v := range split {
			if v != "" {
				resStr = append(resStr, v)
			}
		}
		space := strings.Trim(strings.TrimSpace(cmd2), "'")
		space2 := strings.Trim(space, `"`)
		log.Println("space2:",space2)
		resStr = append(resStr, space2)
	} else {
		split := strings.Split(cmd, " ")
		for _,v := range split {
			resStr = append(resStr, v)
		}
	}

	log.Printf("收到原生命令行,参数为:%#v\n",resStr)

	rds := d.Redis.GetRds(addr)
	if rds == nil {
		return nil, errors.New("connect redis fail,please retry or check the password")
	}
	return rds.Do(resStr...).Result()
}
