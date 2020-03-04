/*
 * @Author: dzw
 * @Date: 2020-02-26 11:21:02
 * @Last Modified by: dzw
 * @Last Modified time: 2020-02-29 21:18:47
 */

package model

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	// InterError ...
	InterError = iota
	// NotExistError ...
	NotExistError
	// TypeError ...
	TypeError
	// PasswordError ...
	PasswordError
	// FormatError ...
	FormatError
	// ExistError ...
	ExistError
	// RstParamError ...
	RstParamError
)

// SortMethod ...
type SortMethod uint8

const (
	// SortInvalid ...
	SortInvalid SortMethod = iota
	// SortInAscend ...
	SortInAscend
	// SortInDescend ...
	SortInDescend
)

// UserType ...
// for file operator
type UserType byte

const (
	// TypeInvalid ...
	TypeInvalid UserType = iota
	// TypeTeacher ...
	TypeTeacher
	// TypeStudent ...
	TypeStudent
)

// AllUserType ...
var AllUserType = []string{"Teacher", "Student"}

// AllUserInfo ...
var AllUserInfo = make(map[string]map[int64]interface{}, len(AllUserType))

// SimsError ...
type SimsError struct {
	ErrCode int
	ErrMsg  string
}

// Student ...
type Student struct {
	ID       int64
	Name     string
	Age      uint8
	Score    float32
	Password string // default: ID
}

// IDScore ...
type IDScore struct {
	ID    int64
	Score float32
}

// IDScoreSlice ...
// 按成绩排序
type IDScoreSlice []IDScore

// Len ...
// 重写 Len() 方法
func (is IDScoreSlice) Len() int {
	return len(is)
}

// Swap ...
// 重写 Swap() 方法
func (is IDScoreSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

// Less ...
// 重写 Less() 方法， 从小到大排序
func (is IDScoreSlice) Less(i, j int) bool {
	return is[j].Score > is[i].Score
}

// Int64Slice ...
type Int64Slice []int64

// Len ...
// 重写 Len() 方法
func (is Int64Slice) Len() int {
	return len(is)
}

// Swap ...
// 重写 Swap() 方法
func (is Int64Slice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

// Less ...
// 重写 Less() 方法， 从小到大排序
func (is Int64Slice) Less(i, j int) bool {
	return is[j] > is[i]
}

// Teacher ...
type Teacher struct {
	ID       int64
	Name     string
	Age      uint8
	Password string
}

// LoginInfo ...
type LoginInfo struct {
	Type     string
	ID       int64
	Password string
}

func (s *SimsError) Error() string {
	return fmt.Sprintf("error code: %d, error msg: %s", s.ErrCode, s.ErrMsg)
}

func init() {
	for _, v := range AllUserType {
		if _, ok := AllUserInfo[v]; !ok {
			AllUserInfo[v] = make(map[int64]interface{}, 100)
		}
	}

	// AllUserInfo[AllUserType[1]][1001] = &Student{1001, "dzw", 26, 98.5, "dzw"}
	// AllUserInfo[AllUserType[1]][1002] = &Student{1002, "qsq", 23, 99.0, "1"}
	AllUserInfo[AllUserType[0]][1] = &Teacher{1, "dzw", 26, "1"}
}

// DisplayInfo ...
func (s Student) DisplayInfo() {
	fmt.Println("ID:", s.ID, "Name:", s.Name, "Age:", s.Age, "Score:", s.Score)
}

// SetName ...
func (s *Student) SetName(name string) {
	s.Name = name
}

// SetAge ...
func (s *Student) SetAge(age uint8) {
	s.Age = age
}

// SetPassword ...
func (s *Student) SetPassword(psw string) {
	s.Password = psw
}

// SetScore ...
func (s *Student) SetScore(score float32) {
	s.Score = score
}

// GetScore ...
func (s *Student) GetScore() float32 {
	return s.Score
}

// Serialize ...
func (s *Student) Serialize() ([]byte, error) {
	lenName := int64(len(s.Name))
	lenPsw := int64(len(s.Password))
	//info = make([]byte, 0, )

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, TypeStudent) // 1 byte, user type
	binary.Write(bytesBuffer, binary.LittleEndian, s.ID)
	binary.Write(bytesBuffer, binary.LittleEndian, lenName)
	// binary.Write(bytesBuffer, binary.LittleEndian, []byte(s.Name))
	if n, err := bytesBuffer.WriteString(s.Name); err != nil || n != len(s.Name) {
		return []byte{}, err
	}
	binary.Write(bytesBuffer, binary.LittleEndian, s.Age)
	binary.Write(bytesBuffer, binary.LittleEndian, s.Score)
	binary.Write(bytesBuffer, binary.LittleEndian, lenPsw)
	// binary.Write(bytesBuffer, binary.LittleEndian, []byte(s.Password))
	if n, err := bytesBuffer.WriteString(s.Password); err != nil || n != len(s.Password) {
		return []byte{}, err
	}

	// fmt.Println(bytesBuffer.Bytes(), bytesBuffer.Len())

	return bytesBuffer.Bytes(), nil
}

// UnSerialize ...
func (s *Student) UnSerialize(b []byte) error {
	bytesBuffer := bytes.NewBuffer(b)
	var usertype UserType
	var id int64
	var name string
	var age uint8
	var score float32
	var password string
	var len int64
	binary.Read(bytesBuffer, binary.LittleEndian, &usertype) // user type
	binary.Read(bytesBuffer, binary.LittleEndian, &id)
	binary.Read(bytesBuffer, binary.LittleEndian, &len)
	// fmt.Println("len =", len)
	nameslice := make([]byte, len, len)
	if n, err := bytesBuffer.Read(nameslice); int64(n) != len || err != nil {
		return err
	}
	// binary.Read(bytesBuffer, binary.LittleEndian, &name)
	name = string(nameslice)
	binary.Read(bytesBuffer, binary.LittleEndian, &age)
	binary.Read(bytesBuffer, binary.LittleEndian, &score)
	binary.Read(bytesBuffer, binary.LittleEndian, &len)
	pswslice := make([]byte, len, len)
	// binary.Read(bytesBuffer, binary.LittleEndian, &password)
	if n, err := bytesBuffer.Read(pswslice); int64(n) != len || err != nil {
		return err
	}
	password = string(pswslice)
	// fmt.Println("len =", len)
	// fmt.Println("Student info=", id, name, age, score, password)
	s.ID = id
	s.Name = name
	s.Age = age
	s.Score = score
	s.Password = password

	return nil
}

// DisplayInfo ...
func (t Teacher) DisplayInfo() {
	fmt.Println("ID:", t.ID, "Name:", t.Name, "Age:", t.Age)
}

// SetName ...
func (t *Teacher) SetName(name string) {
	t.Name = name
}

// SetAge ...
func (t *Teacher) SetAge(age uint8) {
	t.Age = age
}

// SetPassword ...
func (t *Teacher) SetPassword(psw string) {
	t.Password = psw
}

// Serialize ...
func (t *Teacher) Serialize() ([]byte, error) {
	lenName := int64(len(t.Name))
	lenPsw := int64(len(t.Password))
	//info = make([]byte, 0, )

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, TypeTeacher) // 1 byte, user type
	binary.Write(bytesBuffer, binary.LittleEndian, t.ID)
	binary.Write(bytesBuffer, binary.LittleEndian, lenName)
	// binary.Write(bytesBuffer, binary.LittleEndian, []byte(s.Name))
	if n, err := bytesBuffer.WriteString(t.Name); err != nil || n != len(t.Name) {
		return []byte{}, err
	}
	binary.Write(bytesBuffer, binary.LittleEndian, t.Age)
	binary.Write(bytesBuffer, binary.LittleEndian, lenPsw)
	// binary.Write(bytesBuffer, binary.LittleEndian, []byte(s.Password))
	if n, err := bytesBuffer.WriteString(t.Password); err != nil || n != len(t.Password) {
		return []byte{}, err
	}

	// fmt.Println(bytesBuffer.Bytes(), bytesBuffer.Len())

	return bytesBuffer.Bytes(), nil
}

// UnSerialize ...
func (t *Teacher) UnSerialize(b []byte) error {
	bytesBuffer := bytes.NewBuffer(b)
	var usertype UserType
	var id int64
	var name string
	var age uint8
	var password string
	var len int64
	binary.Read(bytesBuffer, binary.LittleEndian, &usertype) // user type
	binary.Read(bytesBuffer, binary.LittleEndian, &id)
	binary.Read(bytesBuffer, binary.LittleEndian, &len)
	// fmt.Println("len =", len)
	nameslice := make([]byte, len, len)
	if n, err := bytesBuffer.Read(nameslice); int64(n) != len || err != nil {
		return err
	}
	// binary.Read(bytesBuffer, binary.LittleEndian, &name)
	name = string(nameslice)
	binary.Read(bytesBuffer, binary.LittleEndian, &age)
	binary.Read(bytesBuffer, binary.LittleEndian, &len)
	pswslice := make([]byte, len, len)
	// binary.Read(bytesBuffer, binary.LittleEndian, &password)
	if n, err := bytesBuffer.Read(pswslice); int64(n) != len || err != nil {
		return err
	}
	password = string(pswslice)
	// fmt.Println("len =", len)
	// fmt.Println("Teacher info=", id, name, age, password)
	t.ID = id
	t.Name = name
	t.Age = age
	t.Password = password

	return nil
}
