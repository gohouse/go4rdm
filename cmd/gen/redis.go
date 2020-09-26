package main

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"strings"
)

func main() {
	//var a = "asdfasdf阿斯顿"
	//b := []rune(a)
	//fmt.Printf("%s",b[:2])
	//return

	//rds := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "123456"})
	//var cmd = `set c "a b"`
	//GetRedisCmdResult(rds,cmd)

	gen()
}

func GetRedisCmdResult(rds *redis.Client, cmd string) (interface{}, error) {
	log.Printf("cmd: %+v\n",cmd)
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

	if rds == nil {
		return nil, errors.New("connect redis fail,please retry or check the password")
	}
	return rds.Do(resStr...).Result()
}
func gen() {
	var words = []string {
		"hello","world","how","newbie","foo","bar","baz",
	}
	rds := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "123456"})
	rds.FlushAll().Val()
	// gen string
	for i:=0;i<50;i++ {
		result, err := rds.Set(fmt.Sprintf("string_redis_test_key_%d", i), fmt.Sprintf("%d -- some string value -- %s", i, strings.Join(words, " ")), 0).Result()
		fmt.Println(result,err)
		//return
	}
	// gen hash
	for i:=0;i<10;i++ {
		var field = fmt.Sprintf("hash_redis_test_key_%d",i)
		for j:=0;j<50;j++ {
			rds.HSet(field, fmt.Sprintf("field_%d",j), fmt.Sprintf("%s -- %d -- some hash value -- %s","hash value", j,strings.Join(words," "))).Val()
		}
	}
	// gen set
	for i:=0;i<10;i++ {
		var field = fmt.Sprintf("set_redis_test_key_%d",i)
		for j:=0;j<50;j++ {
			rds.SAdd(field, fmt.Sprintf("%s%d -- some zset value -- %s", "set value", j,strings.Join(words," "))).Val()
		}
	}
	// gen zset
	for i:=0;i<10;i++ {
		var field = fmt.Sprintf("zset_redis_test_key_%d",i)
		for j:=0;j<50;j++ {
			rds.ZAdd(field, redis.Z{
				Score:  float64(j)+1,
				Member: fmt.Sprintf("%s%d -- some zset value -- %s", "zset value", j,strings.Join(words," ")),
			})
		}
	}
	// gen list
	for i:=0;i<10;i++ {
		var field = fmt.Sprintf("list_redis_test_key_%d",i)
		for j:=0;j<50;j++ {
			rds.RPush(field, fmt.Sprintf("%s -- %d -- some zset value -- %s", "list value", j, strings.Join(words, " "))).Val()
		}
	}
}
