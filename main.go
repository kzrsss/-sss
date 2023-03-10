package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// 原代码没有`json:"username"`等字段，会直接以字段名称为解析内容，修改后即可正常输出
type user struct {
	Username string    `json:"username"`
	Nickname string    `json:"nickname"`
	Sex      uint8     `json:"sex"`
	Birthday time.Time `json:"birthday"`
}

func main() {
	u := user{
		Username: "坤坤",
		Nickname: "阿坤",
		Sex:      20,
		Birthday: time.Now(),
	}
	bs, err := json.Marshal(&u)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(bs))
}
