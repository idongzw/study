/*
 * @Author: dzw
 * @Date: 2020-03-09 11:08:55
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-15 14:01:24
 */

package service

import (
	"chatserver/chatlog"
	"chatserver/conf"
	"chatserver/mysql"
	"chatserver/protos"
	"fmt"
	"net"
	"strconv"
)

// Server ...
type Server struct {
	Conf   conf.Config
	Listen net.Listener
}

// defaultServerWork ...
func defaultServer() *Server {
	s := &Server{
		Conf:   conf.DefaultConf(),
		Listen: nil,
	}

	return s
}

// OptionServer ...
// 函数式选项模式
type OptionServer func(*Server)

// NewServer ...
func NewServer(opts ...OptionServer) *Server {
	s := defaultServer()

	for _, o := range opts {
		o(s)
	}

	return s
}

// Conf ...
func Conf(c conf.Config) OptionServer {
	return func(s *Server) {
		s.Conf = c
	}
}

// InitServer ...
func InitServer(c conf.Config) (*Server, error) {
	s := NewServer(Conf(c))
	// init db
	if err := s.initDB(); err != nil {
		return nil, fmt.Errorf("init db failed, %v", err)
	}

	// load info from db
	if err := s.loadInfoFromDB(); err != nil {
		db.Close()
		return nil, fmt.Errorf("load info from db failed, %v", err)
	}

	// tcp listen
	if err := s.tcpListen(); err != nil {
		db.Close()
		return nil, fmt.Errorf("tcp listen failed, %v", err)
	}

	return s, nil
}

// Run ...
func (s *Server) Run() {
	defer s.tcpClose()
	s.tcpAccept()
}

func (s *Server) initDB() (err error) {
	username := s.Conf.MysqlConf.UserName
	password := s.Conf.MysqlConf.Password
	ip := s.Conf.MysqlConf.IP
	port := strconv.FormatInt(int64(s.Conf.MysqlConf.Port), 10)
	dsn := username + ":" + password + "@tcp(" + ip + ":" + port + ")" + "/chat"

	// db 全局变量
	db, err = mysql.InitDB(dsn)
	if err != nil {
		return err
	}

	return nil
}

// load userInfo from db to memory
func (s *Server) loadInfoFromDB() error {
	var userinfo []*protos.UserInfo
	sqlStr := "select id, password, name, age, gender from userInfo"
	err := mysql.QueryRows(db, &userinfo, sqlStr)
	if err != nil {
		return err
	}

	for _, v := range userinfo {
		userInfoMap.Store(v.GetId(), v)
	}

	chatlog.Trace("load user info: %v", userinfo)
	return nil
}

// tcpListen ...
func (s *Server) tcpListen() error {
	port := strconv.FormatInt(int64(s.Conf.ServerConf.Port), 10)
	address := "0.0.0.0:" + port
	chatlog.Trace(address)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	s.Listen = listen

	return nil
}

func (s *Server) tcpAccept() {
	for {
		conn, err := s.Listen.Accept()
		if err != nil {
			chatlog.Warning("accept failed, %v", err)
			continue
		}
		go processClient(conn, s.Conf.ClientConf) // 开启goroutine处理客户端
	}
}

func (s *Server) tcpClose() {
	s.Listen.Close()
}

func processClient(conn net.Conn, c conf.ClientConf) {
	InitClient(conn, c)
}
