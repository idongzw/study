/*
* @File Name: operator.go
* @Author: idongzw
* @Date:   2020-02-15 17:30:44
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-19 15:29:52
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

/*
位运算符：
    将数值，转为二进制之后，按位操作
按位与 &：
    对应位的值都为 1 才为 1，否则为 0
按位或 |：
    对应位的值都为 0 才为 0，否则为 1
异或 ^：
    二元：a ^ b
        对应位的值不同为 1，相同为 0
    一元：^a
        按位取反：  1 ----> 0
                    0 ----> 1
位清空：&^
    对于：a &^ b
        对于 b 上的每个数值
        如果为 0，则取 a 对应位上的数值
        如果位 1，则结果位就取 0        
 */

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