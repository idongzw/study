/*
 * @Author: dzw
 * @Date: 2020-03-02 22:16:38
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-03 20:52:35
 */

package main

import (
	"homework/mylog"
	"time"
)

func testFunc() {
	i := 0
	log := mylog.New("./log.ini")
	for {
		log.Trace("Trace %d", i)
		log.Debug("Debug %d", i)
		log.Info("Info %d", i)
		log.Warning("Warning %d", i)
		log.Error("Error %d", i)
		log.Fatal("Fatal %d", i)
		time.Sleep(10 * time.Millisecond)
		i++
	}
}

func main() {
	// c := mylog.Config{}
	// err := mylog.LoadConfig("../mylog/log.ini", &c)
	// if err != nil {
	// 	fmt.Println("read config failed,", err)
	// 	return
	// }
	// fmt.Println(c)
	testFunc()
}
