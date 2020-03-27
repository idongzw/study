/*
 * @Author: dzw
 * @Date: 2020-03-09 10:44:23
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-09 11:02:43
 */

package ui

import (
	"fmt"
	"time"
)

// MainUI software main ui
func MainUI() {
	fmt.Println("1. Sign up")
	fmt.Println("2. Log  in")
	fmt.Println("3. exit")
	selectUI()
}

func selectUI() {
	fmt.Println("\033[34;1mCurrent time:", time.Now(), "\033[0m")
	fmt.Print("\033[33mPlease select the following options: \033[0m")
}

func inputIDUI() {
	fmt.Print("Please input your id:")
}

func inputPasswordUI() {
	fmt.Print("Please input your password:")
}

func loginSuccess() {
	fmt.Println("Login success")
}
