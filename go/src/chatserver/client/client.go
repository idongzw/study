/*
 * @Author: dzw
 * @Date: 2020-02-24 11:15:42
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-12 17:02:25
 */

package main

import (
	"bufio"
	"chatserver/protos"
	"chatserver/utils"
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
)

var wg = sync.WaitGroup{}

func main() {
	// 请求连接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("connect tcp server failed, err:", err)
		return
	}

	// 关闭连接
	defer conn.Close()

	chanMsg := make(chan proto.Message, 100)

	ctx, cancel := context.WithCancel(context.Background())

	// signupReq := &protos.SignupRequest{
	// 	Name:   "dzw",
	// 	Age:    26,
	// 	Gender: protos.GenderType_FEMALE,
	// }

	wg.Add(2)
	go processRecv(ctx, cancel, conn, chanMsg)
	go processSend(ctx, cancel, conn, chanMsg)

	// chanMsg <- signupReq

	time.Sleep(time.Second)

	userLogin := &protos.LoginRequest{
		Id:       1584265159,
		Password: "1584265159",
	}
	chanMsg <- userLogin

	time.Sleep(time.Second)

	// userLogout := &protos.LogoutRequest{
	// 	Id: 12345,
	// }

	// chanMsg <- userLogout

	// time.Sleep(time.Second)

	// chanMsg <- userLogout

	// chatMsg := &protos.ChatMsg{
	// 	MsgId:   1,
	// 	MsgType: protos.ChatMsgType_PRIVATE_CHAT,
	// 	MsgFrom: 12346,
	// 	MsgTo:   12345,
	// 	Data:    "hello world",
	// }

	// chanMsg <- chatMsg

	// for i := 0; i < 20; i++ {
	// 	//msg := `hello,How are you.`
	// 	m := &msg.UserInfo{
	// 		Id:       12,
	// 		Password: "qwertyuiop",
	// 		Name:     "dzw",
	// 		Age:      26,
	// 		Gender:   msg.GenderType_FEMALE,
	// 	}
	// 	data, err := proto.EncodeMessage(m)
	// 	if err != nil {
	// 		fmt.Println("encode failed,", err)
	// 		continue
	// 	}
	// 	fmt.Println("data len,", len(data))
	// 	conn.Write(data)
	// 	time.Sleep(5 * time.Second)
	// }

	// inputReader := bufio.NewReader(os.Stdin)
	// for {
	// 	input, err := inputReader.ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println("read string failed, err:", err)
	// 		continue
	// 	}
	// 	inputInfo := strings.Trim(input, "\r\n")
	// 	if strings.ToUpper(inputInfo) == "Q" {
	// 		fmt.Println("client exit")
	// 		return
	// 	}

	// 	_, err = conn.Write([]byte(inputInfo))
	// 	if err != nil {
	// 		fmt.Println("send msg failed, err", err)
	// 	}

	// 	buf := [1024]byte{}
	// 	n, err := conn.Read(buf[:])
	// 	if err != nil {
	// 		fmt.Println("recv msg failed, err", err)
	// 		return
	// 	}

	// 	fmt.Println("recv msg:", string(buf[:n]))
	// }

	// time.Sleep(2 * time.Second)
	// cancel()
	wg.Wait()
}

func processRecv(ctx context.Context, cancel context.CancelFunc, conn net.Conn, cpm chan<- proto.Message) {
	reader := bufio.NewReader(conn)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("process recv stop...")
			return
		default:
		}

		_, pm, err := utils.DecodeMessage(reader)
		if err == io.EOF {
			cancel()
			continue
		}
		if err != nil {
			fmt.Println("decode msg failed,", err)
			continue
		}

		fmt.Println("recv msg:", pm.String())
		signupRsp, ok := pm.(*protos.SignupResponse)
		if ok {
			userLogin := &protos.LoginRequest{
				Id:       signupRsp.Id,
				Password: signupRsp.Password,
			}
			cpm <- userLogin
		}

		loginResponse, ok := pm.(*protos.LoginResponse)
		if ok {
			fmt.Println("loginResponse.GetResult() =", loginResponse.GetResult())
			if loginResponse.GetResult() == protos.LoginResult_SUCCESS {
				go processHb(conn, cpm)
			}
		}
	}
}

func processSend(ctx context.Context, cancel context.CancelFunc, conn net.Conn, cpm <-chan proto.Message) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("process send stop...")
			return
		case pm := <-cpm:
			fmt.Println(pm)
			sendMsg(conn, pm)
		default:
			time.Sleep(time.Millisecond * 20)
		}
	}
}

func processHb(conn net.Conn, cpm chan<- proto.Message) {
	t := time.NewTicker(time.Second * 5) // heartbeat 5s
	defer t.Stop()

	for {
		<-t.C
		hb := &protos.HeartbeatMsg{
			Msg: "hb",
		}
		cpm <- hb
	}
}

func sendMsg(conn net.Conn, pm proto.Message) {
	data, err := utils.EncodeMessage(pm)
	if err != nil {
		fmt.Println("encode failed,", err)
		return
	}
	// fmt.Println("data len,", len(data))
	conn.Write(data)
}
