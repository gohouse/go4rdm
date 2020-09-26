package data

import (
	"encoding/json"
	"fmt"
	"github.com/gohouse/go4rdm/config"
	"github.com/gohouse/t"
	"log"
)

type Api struct {
}

func NewApi() *Api {
	return &Api{}
}

func (*Api) GetQaList() []Qa {
	//return []Qa{
	//	{
	//		Id:       1,
	//		Content:  "问题标题啊",
	//		Email:    "xxx@xxx",
	//		Nickname: "赵日天",
	//		Time:     time.Now(),
	//	},
	//	{
	//		Id:       2,
	//		Content:  "问题标题啊2",
	//		Email:    "xxx@xxx",
	//		Nickname: "叶良辰",
	//		Time:     time.Now().Add(1 * time.Minute),
	//	},
	//	{
	//		Id:       3,
	//		Content:  "问题标题啊3",
	//		Email:    "xxx@xxx",
	//		Nickname: "叶良辰",
	//		Time:     time.Now().Add(2 * time.Minute),
	//	},
	//}

	bytes, err := httpget(config.ApiQaList, nil)
	if err != nil {
		log.Println("get qa list error: ", err.Error())
		return nil
	}
	log.Printf("[GetQaList] get msg: %s\n", bytes)
	var qa []Qa
	err = json.Unmarshal(bytes, &qa)
	if err != nil {
		log.Println("get qa list ,json unmarshal error: ", err.Error())
		return nil
	}

	return qa
}

func (*Api) GetQaReplyList(qa *Qa) []QaReply {
	//return []QaReply{
	//	{
	//		Id:       1,
	//		Content:  "问题回复标题啊",
	//		Email:    "xxx@xxx",
	//		Nickname: "赵日天",
	//		Time:     time.Now(),
	//		QaId:     1,
	//		Pid:      0,
	//	},
	//	{
	//		Id:       2,
	//		Content:  "问题回复标题啊2",
	//		Email:    "xxx@xxx",
	//		Nickname: "赵日天",
	//		Time:     time.Now().Add(2 * time.Minute),
	//		QaId:     1,
	//		Pid:      1,
	//	},
	//}

	bytes, err := httpget(fmt.Sprintf(config.ApiQaReplyList, t.New(qa.Id).String()), *qa)
	if err != nil {
		log.Println("get qa list error: ", err.Error())
		return nil
	}
	log.Printf("[GetQaReplyList] get msg: %s\n", bytes)
	var qr []QaReply
	err = json.Unmarshal(bytes, &qr)
	if err != nil {
		log.Println("get qa list ,json unmarshal error 2: ", err.Error())
		return nil
	}

	log.Printf("[GetQaReplyList] get msg: %+v\n", qa)

	return qr
}

func (*Api) QaAdd(qa *Qa) error {
	log.Printf("add question data: %#v\n", qa)
	res, err := httppost(config.ApiQaAdd, *qa)
	if err != nil {
		return err
	}

	log.Printf("[QaAdd] response: %s\n", res)

	return nil
}

// QaReply 回复问题
func (*Api) QaReply(qr *QaReply) error {
	log.Printf("reply question data: %#v\n", qr)
	_, err := httppost(config.ApiQaReply, *qr)
	if err != nil {
		return err
	}

	return nil
}

func (*Api) GetVersion() (v Version) {
	bytes, err := httpget(config.ApiVersion, nil)
	if err != nil {
		log.Println("GetVersion error: ", err.Error())
		return
	}
	log.Printf("[GetVersion] data: %s\n", bytes)
	err = json.Unmarshal(bytes, &v)
	if err != nil {
		log.Println("GetVersion ,json unmarshal error: ", err.Error())
	}

	return
}

func (*Api) Statistic() {
	_, err := httpget(config.ApiStatistic, nil)
	if err != nil {
		log.Println("GetVersion error: ", err.Error())
	}
}
