/*
 * @Author: dzw
 * @Date: 2020-03-11 09:30:16
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-12 16:42:17
 */

package service

import (
	"bufio"
	"chatserver/chatlog"
	"chatserver/conf"
	"chatserver/utils"
	"context"
	"io"
	"net"
	"time"
)

// Client ...
type Client struct {
	ID          uint64
	Conn        net.Conn
	Ctx         context.Context
	CancelFunc  context.CancelFunc
	chanRecvMsg chan *interMsg
	chanSendMsg chan *interMsg
	chanHbMsg   chan bool
}

func defaultClient() *Client {
	ctx, cancel := context.WithCancel(context.Background())
	c := &Client{
		ID:          0,
		Conn:        nil,
		Ctx:         ctx,
		CancelFunc:  cancel,
		chanRecvMsg: make(chan *interMsg, 100), // default bufsize 100
		chanSendMsg: make(chan *interMsg, 100), // default bufsize 100
		chanHbMsg:   make(chan bool),
	}
	return c
}

// func (tc *Client) WriteSendMsg

// OptionClient ...
// 函数式选项模式
type OptionClient func(*Client)

// NewClient ...
func NewClient(opts ...OptionClient) *Client {
	c := defaultClient()

	for _, o := range opts {
		o(c)
	}

	return c
}

// ID ...
func ID(id uint64) OptionClient {
	return func(c *Client) {
		c.ID = id
	}
}

// Conn ...
func Conn(conn net.Conn) OptionClient {
	return func(c *Client) {
		c.Conn = conn
	}
}

// Ctx ...
func Ctx(ctx context.Context) OptionClient {
	return func(c *Client) {
		c.Ctx = ctx
	}
}

// CancelFunc ...
func CancelFunc(cf context.CancelFunc) OptionClient {
	return func(c *Client) {
		c.CancelFunc = cf
	}
}

func chanRecvMsg(size uint32) OptionClient {
	return func(c *Client) {
		c.chanRecvMsg = make(chan *interMsg, size)
	}
}

func chanSendMsg(size uint32) OptionClient {
	return func(c *Client) {
		c.chanSendMsg = make(chan *interMsg, size)
	}
}

// InitClient ...
func InitClient(conn net.Conn, c conf.ClientConf) {
	nc := NewClient(Conn(conn), chanRecvMsg(c.RecvBufSize), chanSendMsg(c.SendBufSize))

	go nc.handleRecv()
	go nc.handleSend()
	go nc.handleMsg()
}

func (c *Client) close() {
	c.Conn.Close()
}

func (c *Client) clientExit() {
	c.CancelFunc()             // client goroutine exit
	c.close()                  // client link close
	userOnlineMap.Delete(c.ID) // delete user to online map
	chanMsgMap.Delete(c.ID)    // delete client send msg channel
}

func (c *Client) handleRecv() {
	reader := bufio.NewReader(c.Conn)
	for {
		select {
		case <-c.Ctx.Done():
			chatlog.Info("handleRecv stop")
			return
		default:
		}

		msgName, msg, err := utils.DecodeMessage(reader)
		if err == io.EOF {
			chatlog.Info("read from client done")
			c.clientExit()
			continue
		}

		if err != nil {
			chatlog.Warning("decode msg failed,%v", err)
			continue
		}

		im := &interMsg{
			msgName: msgName,
			msg:     msg,
		}
		c.chanRecvMsg <- im // 传入channel，其他goroutine处理
		chatlog.Trace("recv msg from client: msgName = %v,msg = %v", msgName, msg)
	}
}

func (c *Client) handleSend() {
	for {
		select {
		case <-c.Ctx.Done():
			chatlog.Info("handleSend stop")
			return
		case msg := <-c.chanSendMsg:
			data, err := utils.EncodeMessage(msg.msg)
			if err != nil {
				chatlog.Warning("encode msg failed,%v", err)
				continue
			}
			c.Conn.Write(data)
		default:
			time.Sleep(time.Millisecond * 20)
		}
	}
}

func (c *Client) handleMsg() {
	for {
		select {
		case <-c.Ctx.Done():
			chatlog.Info("handleMsg stop")
			return
		case msg := <-c.chanRecvMsg:
			chatlog.Trace("handleMsg %v...", msg.msgName)
			msg.handleMsg(c)
		default:
			time.Sleep(time.Millisecond * 20)
		}
	}
}

func (c *Client) handleHeartbeat() {
	count := 0
	t := time.NewTicker(time.Second * 5) // heartbeat 5s
	defer t.Stop()
	for {
		// chatlog.Trace("handleHeartbeat count = %v", count)
		if count >= 3 { // count 3
			c.clientExit()
			chatlog.Info("timeout heartbeat, handleHeartbeat stop")
			return
		}
		select {
		case <-c.Ctx.Done():
			chatlog.Info("handleHeartbeat stop")
			return
		// 5s timeout handle
		case <-t.C:
			count++
		// recv heartbeat msg
		case <-c.chanHbMsg:
			count = 0
		default:
			time.Sleep(time.Millisecond * 20)
		}
	}
}
