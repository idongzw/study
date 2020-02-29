/*
 * @Author: dzw
 * @Date: 2020-02-26 11:16:11
 * @Last Modified by: dzw
 * @Last Modified time: 2020-02-26 12:27:54
 */

package main

import (
	"fmt"
	"homework/sims/control"
	"homework/sims/model"
	"homework/sims/view"
)

func main() {
	// var age int
	// var name string

	m := &control.Master{}
	m.SetAllUserInfo()
	if err := m.ReadInfoFromFile("info.dat"); err != nil {
		fmt.Println("read file failed,", err)
	}
	defer m.WriteAllInfoToFile("info.dat")

	for {
		view.MainUI()
		s, err := control.GetInput()
		if err != nil {
			fmt.Println("Input error")
			return
		}

		switch s {
		case "1":
			if m.Login(model.AllUserType[0]) {
				m.TeacherControl()
			}
		case "2":
			if m.Login(model.AllUserType[1]) {
				m.StudentControl()
			}
		case "3":
			fmt.Println("Exit system")
			return
			//os.Exit(0)
		default:
			fmt.Println("Input error")
		}
	}
}
