/*
 * @Author: dzw
 * @Date: 2020-02-26 10:42:53
 * @Last Modified by: dzw
 * @Last Modified time: 2020-02-26 11:18:02
 */

package view

import (
	"fmt"
	"time"
)

func selectUI() {
	fmt.Println("\033[34;1mCurrent time:", time.Now(), "\033[0m")
	fmt.Print("\033[33mPlease select the following options: \033[0m")
}

func inputSortUI() {
	fmt.Print("\033[33mPlease select a sort method: \033[0m")
}

// InputIDUI ...
func InputIDUI() {
	fmt.Print("\033[31mPlease input ID: \033[0m")
}

// InputPasswordUI ...
func InputPasswordUI() {
	fmt.Print("\033[31mPlease input password: \033[0m")
}

// LoginSuccessUI ...
func LoginSuccessUI(name string) {
	fmt.Println("\033[33mLogin Success. Welcome", name, "\033[0m")
}

// LoginFailedUI ...
func LoginFailedUI(name string) {
	fmt.Println("\033[31m ", name, ",Login Failed.\033[0m")
}

// MainUI ...
func MainUI() {
	fmt.Println(`
				---Student information management system---
				-------------1. Teacher Login--------------
				-------------2. Student Login--------------
				-------------3. Exit-----------------------`)
	selectUI()
}

// TeacherControlUI ...
func TeacherControlUI() {
	fmt.Println(`
				---------1. Show All Student Info----------
				---------2. Add Student Info---------------
				---------3. Delete Student Info------------
				---------4. Modify Student Info------------
				---------5. Find Student Info--------------
				---------6. Sort Student Info--------------
				---------7. Show My Info-------------------
				---------8. Change My Info-----------------
				---------9. Exit---------------------------`)
	selectUI()
}

// SortControlUI ...
func SortControlUI() {
	fmt.Println(`
				---------1. Display by ID in ascending order------------
				---------2. Display by ID in descending order-----------
				---------3. Display by Score in descending order--------
				---------4. Display by Score in descending order--------
				---------5. Exit-----------------------------------------`)
	inputSortUI()
}

// StudentControlUI ...
func StudentControlUI() {
	fmt.Println(`
				-------------1. Show My Info---------------
				-------------2. Change My Info-------------
				-------------3. Show My Class ranking------
				-------------4. Exit-----------------------`)
	selectUI()
}

// ChangeInfoUI ...
func ChangeInfoUI() {
	fmt.Println(`
				---------1. Change Name---------
				---------2. Change Age----------
				---------3. Change Password-----
				---------4. Exit----------------`)
	selectUI()
}

// ChangeStudentUI ...
func ChangeStudentUI() {
	fmt.Println(`
				---------1. Change Student Score--------
				---------2. Change Student Password-----
				---------3. Exit------------------------`)
	selectUI()
}

// ChangeNameUI ...
func ChangeNameUI() {
	fmt.Printf("Please input your new name:")
}

// ChangeAgeUI ...
func ChangeAgeUI() {
	fmt.Printf("Please input your new age:")
}

// ChangePasswordUI ...
func ChangePasswordUI() {
	fmt.Printf("Please input your new password:")
}

// ChangeScoreUI ...
func ChangeScoreUI() {
	fmt.Printf("Please input new socre:")
}

// ConfirmDeleteUI ...
func ConfirmDeleteUI() {
	fmt.Print("\033[31;1mEnter yes/no to confirm the operation(yes/no): \033[0m")
}

// AddStuInfoUI ...
func AddStuInfoUI() {
	fmt.Printf("\033[31;1mPlease enter in the following format():\tid=1001 name=dzw age=28 score=12.8\n\033[0mUser info:")
}
