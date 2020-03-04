/*
 * @Author: dzw
 * @Date: 2020-02-25 17:42:53
 * @Last Modified by: dzw
 * @Last Modified time: 2020-02-25 17:51:45
 */

package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "dzw", age: 12},
		{name: "asd", age: 13},
		{name: "weq", age: 14},
	}

	for _, v := range stus {
		// v 是结构体变量
		m[v.name] = &v
		fmt.Printf("&v =%p\n", &v)
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name) // weq => weq dzw => weq asd => weq
		fmt.Printf("&v =%p\n", v)
	}
}
