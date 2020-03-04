/*
 * @Author: dzw
 * @Date: 2020-03-03 10:51:33
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-03 12:08:10
 */

package main

import (
	"fmt"
	"path"
	"runtime"
)

func f1() {
	pc, file, line, ok := runtime.Caller(1) // 获取调用f1()处的信息，需向上一层

	if !ok {
		fmt.Println("f1 runtime.Caller() failed.")
		return
	}
	fmt.Println(pc, file, line, ok) // 获取文件名(全路径) 行号
	fmt.Println(path.Base(file))    // 只获取文件名

	funcName := runtime.FuncForPC(pc).Name() // 获取函数名
	// fmt.Printf("func f1 %T, %v\n", funcName, funcName)
	fmt.Println("func f1:", funcName)
}

func main() {
	pc, file, line, ok := runtime.Caller(0)

	if !ok {
		fmt.Println("runtime.Caller() failed.")
		return
	}
	fmt.Println(pc, file, line, ok) // 获取文件名(全路径) 行号

	funcName := runtime.FuncForPC(pc).Name() // 获取函数名
	fmt.Println("func name:", funcName)

	f1()
}
