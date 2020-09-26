package data

import (
	"encoding/json"
	"errors"
	"github.com/gohouse/e"
	"github.com/gohouse/go4rdm/config"
	"github.com/gohouse/t"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func httppost(urladdr string, data interface{}) ([]byte, e.Error) {
	post, err := http.Post(urladdr, "application/json", strings.NewReader(t.New(data).String()))
	if err != nil {
		return nil, e.New(err.Error())
	}
	defer post.Body.Close()

	bytes,err := ioutil.ReadAll(post.Body)
	if err != nil {
		log.Println("[httppost] 1 error:",urladdr)
		return nil, e.New(err.Error())
	}

	logrus.Infof("httppost data response: %s", bytes)

	var js config.ApiReturn
	err = json.Unmarshal(bytes, &js)
	if err != nil {
		log.Println("[httppost] 2 error:",urladdr)
		return nil, e.New(err.Error())
	}
	if js.Code != 200 {
		log.Println("[httppost] 3 error:",urladdr)
		return nil, e.New(t.New(js.Result).String())
	}
	return t.New(js.Result).Bytes(), nil
}

func httpget(urladdr string, data interface{}) ([]byte, error) {
	post, err := http.Get(urladdr)
	if err != nil {
		log.Println("[httpget] 1 error:",urladdr)
		return nil, err
	}
	defer post.Body.Close()

	bytes,err := ioutil.ReadAll(post.Body)
	if err != nil {
		log.Println("[httpget] 2 error:",urladdr)
		return nil, err
	}

	var js config.ApiReturn
	err = json.Unmarshal(bytes, &js)
	if err != nil {
		log.Println("[httpget] 3 error:",urladdr)
		return nil, err
	}
	if js.Code != 200 {
		log.Println("[httpget] 4 error:",urladdr)
		return nil, errors.New(js.Msg)
	}
	return t.New(js.Result).Bytes(), nil
}
