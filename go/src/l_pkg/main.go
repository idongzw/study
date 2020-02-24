/*
* @File Name: main.go
* @Author: idongzw
* @Date:   2020-02-21 12:25:42
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-21 13:04:56
*/
package main

import (
    "fmt"
    "l_pkg/utils" // 会调用 utils 包下的 init() 函数
    _ "l_pkg/pk1" // 只调用 pk1 包下的 init() 函数
)

/*
init() main()
相同点：
    两个函数在定义的时候不能有任何的参数和返回值
    该函数只能由go程序自动调用，不可以被引用
不同点：
    init 可以应用于任意包中，且可重复定义多个
    main 只能用于main包中，且只能定义一个
两个函数的执行顺序：
    在main包中的go文件默认总会被执行
    对同一个go文件的init()调用顺序是从上到下的
    对同一package中的不同文件，将文件名按字符串进行“从大到小”排序，之后顺序调用各文件中的init()函数
    对于不同的package,如果不相互依赖的话，按照main包中import的顺序调用其包中的init()函数
    如果 package 存在依赖，调用顺序为最后被依赖的最先被初始化
        eg. 依赖顺序 main -> A -> B -> C，初始化顺序 C -> B -> A -> main，依次执行 init()函数
        main 包总是被最后一个初始化，因为它总依赖别的包

        避免循环import 依赖 A->B->C->A
        一个包被其他多个包import，但只能被初始化一次
 */

func main() {
    // package utils func init() 1...
    // package utils func init() 2...
    // main init() 1
    // main init() 2
    fmt.Println("main...")
    utils.Count() // package utils func Count()...
    //utils.init() // cannot refer to unexported name utils.init
}

func init() {
    fmt.Println("main init() 1")
}

func init() {
    fmt.Println("main init() 2")
}