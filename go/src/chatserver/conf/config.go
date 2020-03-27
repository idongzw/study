/*
 * @Author: dzw
 * @Date: 2020-03-12 09:36:05
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-12 10:15:29
 */

package conf

import (
	"gopkg.in/ini.v1"
)

// Config server config
type Config struct {
	// 字段名字需大写
	ServerConf `ini:"server"`
	ClientConf `ini:"client"`
	MysqlConf  `ini:"mysql"`
}

// ServerConf for tcp socket
type ServerConf struct {
	Port uint16 `ini:"port"`
}

// ClientConf for tcp client
type ClientConf struct {
	RecvBufSize          uint32 `ini:"recv_buf_size"`
	SendBufSize          uint32 `ini:"send_buf_size"`
	RecvMsgGoRoutineSize uint32 `ini:"process_recvmsg_thread_size"`
	SendMsgGoRoutineSize uint32 `ini:"process_sendmsg_thread_size"`
}

// MysqlConf config
type MysqlConf struct {
	IP       string `ini:"ip"`
	Port     uint16 `ini:"port"`
	UserName string `ini:"username"`
	Password string `ini:"password"`
}

// DefaultConf ...
func DefaultConf() Config {
	sc := Config{
		// default config
		ServerConf{
			Port: 5050,
		},
		ClientConf{
			RecvBufSize:          100,
			SendBufSize:          100,
			RecvMsgGoRoutineSize: 1,
			SendMsgGoRoutineSize: 1,
		},
		MysqlConf{
			IP:       "127.0.0.1",
			Port:     3306,
			UserName: "root",
			Password: "qaz",
		},
	}

	return sc
}

// LoadConf load config
func LoadConf(path string) (Config, error) {
	sc := DefaultConf()

	// load config from ini file
	file, err := ini.Load(path)
	if err != nil {
		// fmt.Println("Load config file failed,", err)
		return sc, err // return default config
	}
	// map config to struct
	err = file.MapTo(&sc)
	if err != nil {
		// fmt.Println("map config to struct failed,", err)
		return sc, err
	}

	return sc, nil
}
