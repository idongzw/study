/*
 * @Author: dzw
 * @Date: 2020-03-08 18:14:13
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-12 10:09:05
 */

package main

import (
	"chatserver/chatlog"
	"chatserver/conf"
	_ "chatserver/protos"
	"chatserver/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	sc, err := conf.LoadConf("./conf/conf.ini")
	if err != nil {
		chatlog.Warning("load config failed, [%v] use default config", err)
	}
	chatlog.Info("config info:%v", sc)
	s, err := service.InitServer(sc)
	if err != nil {
		chatlog.Fatal("Init Server failed, %v", err)
		return
	}
	s.Run() // start tcp server
}
