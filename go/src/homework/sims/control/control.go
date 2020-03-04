/*
 * @Author: dzw
 * @Date: 2020-02-26 11:18:38
 * @Last Modified by: dzw
 * @Last Modified time: 2020-02-29 21:14:01
 */

package control

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"homework/sims/model"
	"homework/sims/view"
	"io"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// Master ...
type Master struct {
	loginInfo   model.LoginInfo
	allUserInfo map[string]map[int64]interface{}
}

// GetInput ...
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

// Login ...
func (m *Master) Login(t string) bool {
	err := m.getLoginInfo(t)
	if err != nil {
		fmt.Println("GetLoginInfo failed,", err)
		return false
	}

	err = m.login()
	if err != nil {
		fmt.Println("Login failed,", err)
		return false
	}

	return true
}

// SetAllUserInfo ...
func (m *Master) SetAllUserInfo() {
	m.allUserInfo = model.AllUserInfo
}

// TeacherControl ...
func (m *Master) TeacherControl() {
	for {
		view.TeacherControlUI()
		s, err := GetInput()
		if err != nil {
			fmt.Println("Input error")
			return
		}

		switch s {
		case "1":
			m.displayAllStuInfo()
		case "2":
			if err := m.addStuInfo(); err != nil {
				fmt.Println("add student info failed,", err)
				break
			}
			fmt.Println("add student info success")
		case "3":
			m.deleteStuInfo()
		case "4":
			m.changeStuInfo()
		case "5":
			if id, err := m.findStuInfo(); err != nil {
				fmt.Printf("find student %d info failed, %v\n", id, err)
			}
		case "6":
			m.sortStuInfo()
		case "7":
			m.displayUserInfo(m.loginInfo.ID, m.loginInfo.Type)
		case "8":
			m.changeSelfInfo()
		case "9":
			fmt.Println("Logout")
			return
		}
	}
}

// StudentControl ...
func (m *Master) StudentControl() {
	for {
		view.StudentControlUI()
		s, err := GetInput()
		if err != nil {
			fmt.Println("Input error")
			return
		}

		switch s {
		case "1":
			m.displayUserInfo(m.loginInfo.ID, m.loginInfo.Type)
		case "2":
			m.changeSelfInfo()
		case "3":
			rank := m.rankInClass()
			fmt.Println("rangking in class:", rank)
		case "4":
			fmt.Println("Logout")
			return
		}
	}
}

// getLoginInfo ...
func (m *Master) getLoginInfo(t string) error {
	view.InputIDUI()
	id, err := GetInput()
	if err != nil {
		return err
	}

	view.InputPasswordUI()
	psw, err := GetInput()
	if err != nil {
		return err
	}

	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	m.loginInfo = model.LoginInfo{Type: t, ID: ID, Password: psw}
	return nil
}

// login ...
func (m *Master) login() error {
	// t := reflect.TypeOf(m.loginInfo)
	// if k := t.Kind(); k != reflect.Struct {
	// 	return &model.SimsError{ErrCode: model.InterError, ErrMsg: "Login user type error"}
	// }

	// v := reflect.ValueOf(m.loginInfo)
	// fv := v.FieldByName("ID")
	// if !fv.IsValid() {
	// 	return &model.SimsError{ErrCode: model.NotExistError, ErrMsg: "Login user id is not valid"}
	// }
	// id := fv.Interface().(int64)
	id := m.loginInfo.ID

	// fv = v.FieldByName("Password")
	// if !fv.IsValid() {
	// 	return &model.SimsError{ErrCode: model.NotExistError, ErrMsg: "Login user password is not valid"}
	// }
	// password := fv.Interface().(string)
	password := m.loginInfo.Password

	// fv = v.FieldByName("Type")
	// if !fv.IsValid() {
	// 	return &model.SimsError{ErrCode: model.NotExistError, ErrMsg: "Login user type is not valid"}
	// }
	// usertype := fv.Interface().(string)
	usertype := m.loginInfo.Type

	_, vvmm, err := m.getValueFromIDAndType(id, usertype)
	if err != nil {
		return err
	}

	fv := vvmm.Elem().FieldByName("Password")
	if !fv.IsValid() {
		return &model.SimsError{ErrCode: model.NotExistError, ErrMsg: "Login user password is not valid"}
	}
	// need reflect
	passwordlocal, ok := fv.Interface().(string)
	if !ok {
		return &model.SimsError{ErrCode: model.TypeError, ErrMsg: "Password type error"}
	}

	if passwordlocal == password {
		view.LoginSuccessUI(strconv.FormatInt(id, 10))
	} else {
		view.LoginFailedUI(strconv.FormatInt(id, 10))
		return &model.SimsError{ErrCode: model.PasswordError, ErrMsg: "Password error"}
	}

	return nil
}

func (m *Master) displayAllStuInfo() {
	stuMap, ok := m.allUserInfo[model.AllUserType[1]]
	if !ok {
		fmt.Println("System no student info.")
		return
	}

	for id := range stuMap {
		// t := reflect.TypeOf(stuInfo)
		// if k := t.Kind(); k != reflect.Ptr {
		// 	fmt.Println("id =", id, "kind is not Struct")
		// 	continue
		// }

		// v := reflect.ValueOf(stuInfo)
		// dis := v.MethodByName("DisplayInfo")
		// c := reflect.Value{}
		// if dis != c {
		// 	dis.Call(nil)
		// }
		m.displayUserInfo(id, model.AllUserType[1])
	}
}

func (m *Master) addStuInfo() error {
	view.AddStuInfoUI()

	info, err := GetInput()
	if err != nil {
		// fmt.Println("change student info failed,", err)
		return err
	}

	// 解析字符串
	// eg. id=1001 name=dzw age=28 score=12.8
	infoSlice := strings.Split(info, " ") // 按空格分割
	afterSlice := make([]string, 0, len(infoSlice))
	// fmt.Println(infoSlice, afterSlice)
	for _, v := range infoSlice {
		// fmt.Println("i =", i, ",len(v) =", len(v), ",", v)
		if len(v) > 3 { // 长度大于字符串 "id="
			afterSlice = append(afterSlice, v)
		}
	}
	// fmt.Println(afterSlice)
	if len(afterSlice) != 4 { // 4个参数
		// fmt.Println("param error")
		return &model.SimsError{ErrCode: model.FormatError, ErrMsg: "The number of parameters is not equal to 4"}
	}

	var id int64 = -1
	var name string = ""
	var age uint8 = 0
	var score float32 = -1.0
	var password string = ""

	flag := 4
	for _, v := range afterSlice {
		kv := strings.Split(v, "=") // 按=分割
		if len(kv) != 2 {           // 不是一个=报错
			return &model.SimsError{ErrCode: model.FormatError, ErrMsg: "Parameter format error"}
		}

		switch kv[0] {
		case "id":
			i, err := strconv.ParseInt(kv[1], 10, 64)
			if err != nil {
				return err
			}
			id = i
			password = kv[1]
			flag--
		case "name":
			name = kv[1]
			flag--
		case "age":
			i, err := strconv.ParseUint(kv[1], 10, 8)
			if err != nil {
				return err
			}
			age = uint8(i)
			flag--
		case "score":
			i, err := strconv.ParseFloat(kv[1], 32)
			if err != nil {
				return err
			}
			score = float32(i)
			flag--
		}
	}
	// fmt.Println(id, name, score, age, password)

	if flag != 0 {
		return &model.SimsError{ErrCode: model.FormatError, ErrMsg: "Parameter format error"}
	}

	if _, ok := m.allUserInfo[model.AllUserType[1]][id]; ok {
		return &model.SimsError{ErrCode: model.ExistError, ErrMsg: "user is exist"}
	}

	m.allUserInfo[model.AllUserType[1]][id] = &model.Student{ID: id, Name: name, Age: age, Score: score, Password: password}
	return nil
}

func (m *Master) deleteStuInfo() {
	id, err := m.findStuInfo()
	if err != nil {
		fmt.Println("delete student info failed,", err)
		return
	}

	vm, ok := m.allUserInfo[model.AllUserType[1]]
	if !ok {
		fmt.Printf("user %d not exist\n", id)
		return
	}

	view.ConfirmDeleteUI()
	in, err := GetInput()
	if err != nil {
		fmt.Println("delete student info failed,", err)
		return
	}

	switch in {
	case "yes":
		delete(vm, id)
		fmt.Println("delete user success")
	default:
		fmt.Println("user cancel operation")
	}
}

func (m *Master) changeStuInfo() {
	id, err := m.findStuInfo()
	if err != nil {
		fmt.Println("change student info failed,", err)
		return
	}

	for {
		view.ChangeStudentUI()

		s, err := GetInput()
		if err != nil {
			fmt.Println("change student info failed,", err)
			return
		}

		switch s {
		case "1":
			if err := m.changeScore(id, model.AllUserType[1]); err != nil {
				fmt.Println("change student score failed,", err)
				break
			}
			fmt.Println("change student score success")
		case "2":
			if err := m.changePassword(id, model.AllUserType[1]); err != nil {
				fmt.Println("change student password failed,", err)
				break
			}
			fmt.Println("change student password success")
		case "3":
			fmt.Println("exit change info")
			return
		}
	}
}

func (m *Master) findStuInfo() (int64, error) {
	fmt.Print("Please input student id:")

	id, err := GetInput()
	if err != nil {
		// fmt.Println("findStuInfo failed,", err)
		return 0, err
	}

	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		// fmt.Println("findStuInfo failed,", err)
		return ID, err
	}

	// if err := m.displayUserInfo(ID, model.AllUserType[1]); err != nil {
	// 	return err
	// }
	_, v, err := m.getValueFromIDAndType(ID, model.AllUserType[1])
	if err != nil {
		// fmt.Println("displayUserInfo failed:", err)
		return ID, err
	}

	dis := v.MethodByName("DisplayInfo")
	c := reflect.Value{}
	if dis == c {
		return ID, &model.SimsError{ErrCode: model.NotExistError, ErrMsg: "Not exist method"}
	}

	dis.Call(nil)
	return ID, nil
}

func (m *Master) sortStuInfo() {
	for {
		view.SortControlUI()
		s, err := GetInput()
		if err != nil {
			fmt.Println("sort student info failed,", err)
			return
		}

		switch s {
		case "1":
			m.sortInfoByID(model.AllUserType[1], model.SortInAscend)
		case "2":
			m.sortInfoByID(model.AllUserType[1], model.SortInDescend)
		case "3":
			m.sortInfoByScore(model.AllUserType[1], model.SortInAscend)
		case "4":
			m.sortInfoByScore(model.AllUserType[1], model.SortInDescend)
		case "5":
			fmt.Println("exit sort student info")
			return
		}
	}
}

func (m *Master) sortInfoByID(usertype string, sortmethod model.SortMethod) {
	// usertype := model.AllUserType[1]
	stuMap := m.allUserInfo[usertype]

	// fmt.Println(len(stuMap))
	s := make([]int64, 0, len(stuMap))
	for k := range stuMap {
		s = append(s, k)
	}

	switch sortmethod {
	case model.SortInAscend:
		sort.Sort(model.Int64Slice(s)) // 默认从小到大
	case model.SortInDescend:
		sort.Sort(sort.Reverse(model.Int64Slice(s))) // 从大到小
	}
	// fmt.Println(s)
	for _, v := range s {
		m.displayUserInfo(v, usertype)
	}
}

func (m *Master) sortInfoByScore(usertype string, sortmethod model.SortMethod) {
	// usertype := model.AllUserType[1]
	stuMap := m.allUserInfo[usertype]

	sIDScore := make(model.IDScoreSlice, 0, len(stuMap))
	for k := range stuMap {
		score, err := m.getUserScore(k, usertype)
		if err != nil {
			continue
		}
		sIDScore = append(sIDScore, model.IDScore{ID: k, Score: score})
	}

	switch sortmethod {
	case model.SortInAscend:
		sort.Sort(sIDScore) // 默认从小到大
	case model.SortInDescend:
		sort.Sort(sort.Reverse(sIDScore)) // 从大到小
	}

	for _, v := range sIDScore {
		m.displayUserInfo(v.ID, usertype)
	}
}

func (m *Master) getUserScore(id int64, usertype string) (float32, error) {
	_, v, err := m.getValueFromIDAndType(id, usertype)
	if err != nil {
		return 0, err
	}

	dis := v.MethodByName("GetScore")
	c := reflect.Value{}
	if dis == c {
		return 0, &model.SimsError{ErrCode: model.NotExistError, ErrMsg: "Not exist method"}
	}
	rst := dis.Call(nil)
	if len(rst) != 1 { // 返回参数个数
		return 0, &model.SimsError{ErrCode: model.RstParamError, ErrMsg: "Return param error"}
	}
	// need reflect
	score, ok := rst[0].Interface().(float32)
	if !ok {
		return 0, &model.SimsError{ErrCode: model.TypeError, ErrMsg: "Score type error"}
	}

	return score, nil
}

func (m *Master) displayUserInfo(id int64, usertype string) {
	_, v, err := m.getValueFromIDAndType(id, usertype)
	if err != nil {
		fmt.Println("displayUserInfo failed:", err)
		return
		// return err
	}

	dis := v.MethodByName("DisplayInfo")
	c := reflect.Value{}
	if dis == c {
		return
		// return &model.SimsError{ErrCode: model.NotExistError, ErrMsg: "Not exist method"}
	}

	dis.Call(nil)
	// return nil
}

func (m *Master) changeSelfInfo() {
	for {
		view.ChangeInfoUI()

		s, err := GetInput()
		if err != nil {
			return
		}

		switch s {
		case "1":
			if err := m.changeName(m.loginInfo.ID, m.loginInfo.Type); err != nil {
				fmt.Println("cahnge name failed,", err)
				break
			}
			fmt.Println("change name success")
		case "2":
			if err := m.changeAge(m.loginInfo.ID, m.loginInfo.Type); err != nil {
				fmt.Println("change age failed", err)
				break
			}
			fmt.Println("change age success")
		case "3":
			if err := m.changePassword(m.loginInfo.ID, m.loginInfo.Type); err != nil {
				fmt.Println("change password failed,", err)
				break
			}
			fmt.Println("change password success")
		case "4":
			fmt.Println("exit change self info")
			return
		}

	}

}

func (m *Master) rankInClass() int {
	usertype := model.AllUserType[1]
	stuMap := m.allUserInfo[usertype]
	id := m.loginInfo.ID

	sIDScore := make(model.IDScoreSlice, 0, len(stuMap))
	for k := range stuMap {
		score, err := m.getUserScore(k, usertype)
		if err != nil {
			continue
		}
		sIDScore = append(sIDScore, model.IDScore{ID: k, Score: score})
	}

	sort.Sort(sort.Reverse(sIDScore)) // 从大到小

	for i, v := range sIDScore {
		if v.ID == id {
			return i + 1
		}
	}

	return 0
}

func (m *Master) changeName(id int64, usertype string) error {
	view.ChangeNameUI()

	name, err := GetInput()
	if err != nil {
		return err
	}

	// id := m.loginInfo.ID
	// usertype := m.loginInfo.Type
	_, v, err := m.getValueFromIDAndType(id, usertype)
	if err != nil {
		// fmt.Println("changeName failed:", err)
		return err
	}

	dis := v.MethodByName("SetName")
	c := reflect.Value{}
	if dis == c {
		return &model.SimsError{ErrCode: model.NotExistError, ErrMsg: "Not exist method"}
	}

	dis.Call([]reflect.Value{reflect.ValueOf(name)})
	return nil
}

func (m *Master) changeAge(id int64, usertype string) error {
	view.ChangeAgeUI()

	age, err := GetInput()
	if err != nil {
		return err
	}

	// id := m.loginInfo.ID
	// usertype := m.loginInfo.Type
	_, v, err := m.getValueFromIDAndType(id, usertype)
	if err != nil {
		// fmt.Println("changeAge failed:", err)
		return err
	}

	dis := v.MethodByName("SetAge")
	c := reflect.Value{}
	if dis == c {
		return &model.SimsError{ErrCode: model.NotExistError, ErrMsg: "Not exist method"}
	}

	Age, err := strconv.ParseUint(age, 10, 8)
	if err != nil {
		return err
	}
	dis.Call([]reflect.Value{reflect.ValueOf(uint8(Age))})
	return nil
}

func (m *Master) changePassword(id int64, usertype string) error {
	view.ChangePasswordUI()

	psw, err := GetInput()
	if err != nil {
		return err
	}

	// id := m.loginInfo.ID
	// usertype := m.loginInfo.Type
	_, v, err := m.getValueFromIDAndType(id, usertype)
	if err != nil {
		// fmt.Println("changeName failed:", err)
		return err
	}

	dis := v.MethodByName("SetPassword")
	c := reflect.Value{}
	if dis == c {
		return &model.SimsError{ErrCode: model.NotExistError, ErrMsg: "Not exist method"}
	}

	dis.Call([]reflect.Value{reflect.ValueOf(psw)})
	return nil
}

func (m *Master) changeScore(id int64, usertype string) error {
	view.ChangeScoreUI()

	score, err := GetInput()
	if err != nil {
		return err
	}

	// id := m.loginInfo.ID
	// usertype := m.loginInfo.Type
	_, v, err := m.getValueFromIDAndType(id, usertype)
	if err != nil {
		// fmt.Println("changeAge failed:", err)
		return err
	}

	dis := v.MethodByName("SetScore")
	c := reflect.Value{}
	if dis == c {
		return &model.SimsError{ErrCode: model.NotExistError, ErrMsg: "Not exist method"}
	}

	Score, err := strconv.ParseFloat(score, 32)
	if err != nil {
		return err
	}
	dis.Call([]reflect.Value{reflect.ValueOf(float32(Score))})
	return nil
}

func (m *Master) getValueFromIDAndType(id int64, usertype string) (reflect.Type, reflect.Value, error) {
	vm, ok := m.allUserInfo[usertype]
	if !ok {
		err := fmt.Sprintf("user type %s not exist", usertype)
		return nil, reflect.Value{}, &model.SimsError{ErrCode: model.NotExistError, ErrMsg: err}
	}

	vmm, ok := vm[id]
	if !ok {
		err := fmt.Sprintf("user %d not exist", id)
		return nil, reflect.Value{}, &model.SimsError{ErrCode: model.NotExistError, ErrMsg: err}
	}

	t := reflect.TypeOf(vmm)
	if k := t.Kind(); k != reflect.Ptr {
		return nil, reflect.Value{}, &model.SimsError{ErrCode: model.InterError, ErrMsg: "user type error"}
	}

	v := reflect.ValueOf(vmm)
	if k := v.Elem().Kind(); k != reflect.Struct {
		return nil, reflect.Value{}, &model.SimsError{ErrCode: model.InterError, ErrMsg: "user type error"}
	}

	return t, v, nil
}

// ReadInfoFromFile ...
// file operation
// read file
func (m *Master) ReadInfoFromFile(filepath string) error {
	// open file
	file, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	// close file
	defer file.Close()

	// check file size
	// stats, err := file.Stat()
	// if err != nil {
	// 	return err
	// }
	// fileSize := stats.Size()
	// buf := make([]byte, fileSize)

	// read file
	dataLenBytes := make([]byte, 8)
	userTypeBytes := make([]byte, 1)
	for {
		// data len
		var dataLen int64 = 0 // data len
		n, err := io.ReadFull(file, dataLenBytes)
		if err != nil || n != 8 {
			break
		}
		bytesBuffer := bytes.NewBuffer(dataLenBytes)
		binary.Read(bytesBuffer, binary.LittleEndian, &dataLen)
		// fmt.Println("data len =", dataLen)

		// user type
		n, err = io.ReadFull(file, userTypeBytes)
		if err != nil || n != 1 {
			break
		}
		// fmt.Println("user type =", userTypeBytes[0])

		// data
		file.Seek(-1, 1) // 回退一个字节，用户类型
		dataBytes := make([]byte, dataLen)
		n, err = io.ReadFull(file, dataBytes)
		if err != nil || int64(n) != dataLen {
			break
		}

		switch model.UserType(userTypeBytes[0]) {
		case model.TypeTeacher:
			s := &model.Teacher{}
			if err := s.UnSerialize(dataBytes); err != nil {
				continue
			}
			// s.DisplayInfo()
			m.allUserInfo[model.AllUserType[0]][s.ID] = s
		case model.TypeStudent:
			t := &model.Student{}
			if err := t.UnSerialize(dataBytes); err != nil {
				continue
			}
			// t.DisplayInfo()
			m.allUserInfo[model.AllUserType[1]][t.ID] = t
		}
	}
	// fmt.Println("file size =", fileSize, ",read size =", n)

	return nil
}

// WriteAllInfoToFile ...
func (m *Master) WriteAllInfoToFile(filepath string) error {
	// open file
	// fmt.Println("WriteAllInfoToFile begin")
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("WriteAllInfoToFile failed,", err)
		return nil
	}

	// close file
	defer file.Close()

	// 带缓存的写
	bufferWriter := bufio.NewWriter(file)

	for usertype, infomap := range m.allUserInfo {
		for id := range infomap {
			_, v, err := m.getValueFromIDAndType(id, usertype)
			if err != nil {
				// fmt.Println("changeAge failed:", err)
				// return err
				continue
			}

			f := v.MethodByName("Serialize")
			c := reflect.Value{}
			if f == c {
				// return &model.SimsError{ErrCode: model.NotExistError, ErrMsg: "Not exist method"}
				continue
			}

			rst := f.Call(nil)
			if len(rst) != 2 { // 返回参数个数
				continue
			}
			switch rst[1].Interface().(type) {
			case error:
				continue
			case nil:
				// fmt.Println("err is nil")
			}
			// fmt.Println("write info...")
			// user info
			userBytesInfo := rst[0].Interface().([]byte)
			// user info len
			userBytesInfoLen := int64(len(userBytesInfo))
			// total data len
			totalDataLen := 8 + userBytesInfoLen // 8 bytes for data len
			// create total data
			totalData := make([]byte, 0, totalDataLen)
			// user info len ---> bytes
			bytesBuffer := bytes.NewBuffer([]byte{})
			binary.Write(bytesBuffer, binary.LittleEndian, userBytesInfoLen)
			// total data that should be written to the file
			totalData = append(totalData, bytesBuffer.Bytes()...)
			totalData = append(totalData, userBytesInfo...)
			// if n, err := file.Write(totalData); int64(n) != totalDataLen || err != nil {
			// 	continue
			// }
			writeSize, _ := bufferWriter.Write(totalData)
			if int64(writeSize) != totalDataLen {
				bufferWriter.Flush() // 写到磁盘
				bufferWriter.Write(totalData[writeSize:])
			}

			if bufferWriter.Available() <= 0 {
				bufferWriter.Flush() // 写到磁盘
			}
		}
	}
	bufferWriter.Flush() // 写到磁盘

	return nil
}
