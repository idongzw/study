/*
 * @Author: dzw
 * @Date: 2020-03-09 10:51:12
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-11 11:09:24
 */

package utils

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/golang/protobuf/proto"
)

// GetInput get console input
func GetInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)  // 从标准输入生成读对象
	text, err := reader.ReadString('\n') // 读到换行
	if err != nil {
		return "", err
	}

	text = strings.TrimSpace(text) // 去除字符串前后空格
	// fmt.Println("Input msg:", text)
	return text, err
}

// EncodeStrMsg ...
// length msgType msg
func EncodeStrMsg(msg string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	len := int32(len(msg))
	pkg := new(bytes.Buffer)

	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, len)
	if err != nil {
		return nil, err
	}

	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(msg))
	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}

// DecodeStrMsg ...
func DecodeStrMsg(reader *bufio.Reader) (string, error) {
	msgLenBytes, err := reader.Peek(4) // 读取前4个bytes，为数据的长度
	if err != nil {
		return "", err
	}

	// 取出消息体的长度
	buf := bytes.NewBuffer(msgLenBytes)
	var len int32
	err = binary.Read(buf, binary.LittleEndian, &len)
	if err != nil {
		return "", err
	}

	// 缓冲中现有的可读取的字节数
	if int32(reader.Buffered()) < len+4 {
		return "", err
	}

	// 读取消息体
	pkg := make([]byte, int(4+len))
	_, err = reader.Read(pkg)
	if err != nil {
		return "", err
	}

	return string(pkg[4:]), nil
}

// EncodeMessage encode protobuf message
func EncodeMessage(pm proto.Message) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	dataBytes, err := proto.Marshal(pm)
	msgName := proto.MessageName(pm)

	if err != nil {
		return nil, err
	}

	dataBytesLen := int32(len(dataBytes))
	msgNameLen := int32(len(msgName))
	totleLength := dataBytesLen + msgNameLen + 4
	pkg := new(bytes.Buffer)

	// 写入消息总长度
	err = binary.Write(pkg, binary.LittleEndian, totleLength)
	if err != nil {
		return nil, err
	}

	// 写入消息类型的长度
	err = binary.Write(pkg, binary.LittleEndian, msgNameLen)
	if err != nil {
		return nil, err
	}

	// 写入消息类型
	err = binary.Write(pkg, binary.LittleEndian, []byte(msgName))
	if err != nil {
		return nil, err
	}

	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, dataBytes)
	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}

// DecodeMessage decode protobuf message
// return msg name, msg, error
func DecodeMessage(reader *bufio.Reader) (string, proto.Message, error) {
	msgLenBytes, err := reader.Peek(8) // 读取前8个bytes，为数据的长度
	if err != nil {
		return "", nil, err
	}

	// 取出消息体的长度
	buf := bytes.NewBuffer(msgLenBytes)
	var totleLen int32
	var msgNameLen int32
	// 取出总长度
	err = binary.Read(buf, binary.LittleEndian, &totleLen)
	if err != nil {
		return "", nil, err
	}
	// 取出消息类型长度
	err = binary.Read(buf, binary.LittleEndian, &msgNameLen)
	if err != nil {
		return "", nil, err
	}

	// 缓冲中现有的可读取的字节数
	if int32(reader.Buffered()) < totleLen+4 {
		return "", nil, err
	}

	// 读取消息体
	pkg := make([]byte, int(4+totleLen))
	_, err = reader.Read(pkg)
	if err != nil {
		return "", nil, err
	}

	msgName := string(pkg[8 : 8+msgNameLen]) // 取消息类型
	// fmt.Println("msgName:", msgName)
	// 根据消息类型创建proto.Message
	t := proto.MessageType(msgName)
	if t == nil {
		return "", nil, fmt.Errorf("msg type [%s] find failed", msgName)
	}

	// Elem() Type
	// 返回map类型的键的类型。如非映射类型将panic
	pm := reflect.Indirect(reflect.New(t.Elem())).Addr().Interface().(proto.Message)
	err = proto.Unmarshal(pkg[8+msgNameLen:], pm)
	// fmt.Println("pm,", pm)
	if err != nil {
		return "", nil, err
	}
	return msgName, pm, nil
}
