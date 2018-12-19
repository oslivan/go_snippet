package main

import (
	"encoding/json"
	"log"
)

// 一般HTTP返回的结构都是小写了，所以在后面加上 json:xxx 注解
type HttpResponse struct {
	Code int    `json:code`
	Msg  string `json:msg`
}

func main() {
	var err error
	body := `{"code": 0, "msg": "success"}`
	mapRes := make(map[string]interface{})
	err = json.Unmarshal([]byte(body), &mapRes)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("== Map result ==")
	log.Println("Code:", mapRes["code"])
	log.Println("Msg:", mapRes["msg"])

	structRes := HttpResponse{}
	err = json.Unmarshal([]byte(body), &structRes)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("== Struct result ==")
	log.Println("Code:", structRes.Code)
	log.Println("Msg:", structRes.Msg)
}
