/*
 * @Author: dzw
 * @Date: 2020-02-25 12:42:57
 * @Last Modified by: dzw
 * @Last Modified time: 2020-02-25 13:09:39
 */

package main

import (
	"flag"
	"fmt"
	"os"
)

// Go语言内置的flag包实现了命令行参数的解析，flag包使得开发命令行工具更为简单
func main() {
	if len(os.Args) > 0 {
		for i, arg := range os.Args {
			fmt.Printf("args[%d] = %v\n", i, arg)
		}
	}

	//flag.Type(flag名, 默认值, 帮助信息) *Type
	// name、age、married、delay均为对应类型的指针
	// flag.Type()
	name := flag.String("name", "dzw", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否")
	delay := flag.Duration("d", 0, "时间间隔")

	// or
	// flag.TypeVar()
	/*
		var name string
		var age int
		var married bool
		var delay time.Duration

		flag.StringVar(&name, "name", "dzw", "姓名")
		flag.IntVar(&age, "age", 18, "年龄")
		flag.BoolVar(married, "married", false, "婚否")
		flag.DurationVar(delay, "d", 0, "时间间隔")
	*/

	// flag.Parse()
	/*
		支持的命令行参数格式有以下几种：

		-flag xxx （使用空格，一个-符号）
		--flag xxx （使用空格，两个-符号）
		-flag=xxx （使用等号，一个-符号）
		--flag=xxx （使用等号，两个-符号）
		其中，布尔类型的参数必须使用等号的方式指定。
	*/
	flag.Parse()
	fmt.Println(*name, *age, *married, *delay)
	// 返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	// 返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	// 返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
