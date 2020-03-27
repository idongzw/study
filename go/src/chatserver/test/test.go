/*
 * @Author: dzw
 * @Date: 2020-03-08 18:21:15
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-11 10:42:52
 */

package main

import (
	"chatserver/protos"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
)

// ServerConf test ini
type ServerConf struct {
	Server string `ini:"server"` // 字段名字需大写
	Port   uint16 `ini:"port"`
}

// Configure ...
type Configure struct {
	ServerConf
}

func main() {
	// file, err := ini.Load("../conf/client.ini")
	// if err != nil {
	// 	fmt.Println("load ini failed,", err)
	// 	return
	// }

	// conf := &Configure{}
	// err = file.MapTo(conf)
	// if err != nil {
	// 	fmt.Println("map to struct failed,", err)
	// 	return
	// }

	// fmt.Println(conf)

	// u := msg.UserInfo{}

	// fmt.Println("user =", u)
	// in, err := utils.GetInput()
	// if err == nil {
	// 	fmt.Println(in)
	// }
	m := &protos.ChatMsg{
		MsgId:   1,
		MsgType: protos.ChatMsgType_GROUP_CHAT,
		MsgFrom: 12,
		MsgTo:   13,
		Data:    "dzw",
	}
	//	fmt.Println(m)
	f(m)

	fmt.Println(time.Now().Format("20060102150405"))
	fmt.Println(time.Now().UnixNano())
}

func f(pm proto.Message) {
	s := pm.String()
	p := proto.MessageName(pm)
	v := proto.MessageType(p)
	fmt.Println(s, p, v)
}
