/*
* @File Name: util.go
* @Author: idongzw
* @Date:   2020-02-21 12:21:46
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-21 12:37:04
*/
package utils

import "fmt"

func Count() {
    fmt.Println("package utils func Count()...")
}

func init() {
    fmt.Println("package utils func init() 1...")
}

func init() {
    fmt.Println("package utils func init() 2...")
}