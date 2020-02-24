/*
* @File Name: input.go
* @Author: idongzw
* @Date:   2020-02-21 15:51:05
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-21 16:01:06
*/
package main

import (
    "fmt"
    "os"
    "io"
)

func main() {
    // 打开文件
    file, err := os.Open("test.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Open file success")

    // 关闭文件
    defer file.Close()

    // 读取文件
    buf := make([]byte, 4, 4)
    for {
        n, err := file.Read(buf)
        if n == 0 || err == io.EOF {
            fmt.Println("Read done")
            break
        }
        fmt.Print(string(buf[:n]))
    }
}