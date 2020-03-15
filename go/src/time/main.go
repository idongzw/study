/*
 * @Author: dzw
 * @Date: 2020-03-12 16:24:39
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-12 16:50:01
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		t := time.NewTicker(time.Second)
		defer wg.Done()
		defer t.Stop()
		for {
			<-t.C
			fmt.Println("Ticker:", time.Now().Unix())
		}
	}()

	go func() {
		t := time.NewTimer(time.Second * 2)
		defer wg.Done()
		defer t.Stop()
		for {
			<-t.C
			fmt.Println("Timer:", time.Now().Unix())
			t.Reset(time.Second * 3) // 需重置
		}
	}()
	wg.Wait()
}
