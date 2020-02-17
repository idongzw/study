/*
* @File Name: operator.go
* @Author: idongzw
* @Date:   2020-02-15 17:30:44
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-15 18:21:57
*/

package main

import (
    "fmt"
)

// 运算符优先级（从高到低）
/*
^ ! (一元运算符)
* / % << >> & &^
+ - | ^ (二元运算符)
== != < <= >= >
<- (专门用于channel)
&&
||

*/


// 存储单位
const (
    B float64 = 1 << (iota * 10)
    KB
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

func main() {
    fmt.Println(^2)     //按位取反
    fmt.Println(1 ^ 2)  //异或

    fmt.Println(!true)

    fmt.Println(1 << 10 << 10 >> 10)

    /*
    6:  0110
    11: 1011
    ---------------------
    &   0010 = 2
    |   1111 = 15
    ^   1101 = 13
    &^  0100 = 4
     */
    
    fmt.Println(6 & 11)
    fmt.Println(6 | 11)
    fmt.Println(6 ^ 11)
    fmt.Println(6 &^ 11)

    fmt.Println(B)
    fmt.Println(KB)
    fmt.Println(MB)
    fmt.Println(GB)
    fmt.Println(TB)
    fmt.Println(EB)
    fmt.Println(ZB)
    fmt.Println(YB)

    //++ 和 -- 只能作为语句，不能作为表达式
    //只能放在变量右边
    a := 1
    a++
    //b := a++ //syntax error: unexpected ++, expecting semicolon or newline or }
    fmt.Println(a)
}