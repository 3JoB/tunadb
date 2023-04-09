// Open Source: MIT License
// Author: Leon Ding <ding@ibyte.me>
// Date: 2022/2/28 - 10:14 下午 - UTC/GMT+08:00

//go:build ignore
// +build ignore

package main

import (
	"fmt"

	"github.com/3JoB/tunadb"
)

func init() {
	if err := tunadb.Open(tunadb.Option{
		Directory:       "./testdata",
		DataFileMaxSize: 10240,
	}); err != nil {
		panic(err)
	}
}

type UserInfo struct {
	Name  string
	Age   uint8
	Skill []string
}

func main() {
	// for i := 0; i < 1000; i++ {
	//	tunadb.Put([]byte("999"), tunadb.Bson(&UserInfo{
	//		Name:  fmt.Sprintf("user-%d", i),
	//		Age:   22,
	//		Skill: []string{"Java", "Go", "Rust"},
	//	}))
	// }

	var u UserInfo

	key := fmt.Sprintf("%d", 999)
	fmt.Println("FIND KEY IS:", key)

	// 通过Unwrap解析出结构体
	tunadb.Get([]byte(key)).Unwrap(&u)

	// 打印取值
	fmt.Println("FIND KEY VALUE IS:", u)

	tunadb.Close()
}
