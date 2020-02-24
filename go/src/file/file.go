/*
* @File Name: file.go
* @Author: idongzw
* @Date:   2020-02-21 13:30:36
* @Last Modified by:   idongzw
* @Last Modified time: 2020-02-21 15:32:37
*/
package main

import (
    "os"
    "fmt"
    "path/filepath"
    "path"
)

func main() {
    /*
    FileInfo 文件信息
     */
    
    // 权限
    /*
    -     ---    ---    ---
    type  owner  group  others
    - 文件
    d 目录
    l 链接符号
     */
    
    {
        fileInfo, err := os.Stat("./test.txt")
        if err != nil {
            fmt.Println("err:", err)

            return 
        }
        fmt.Printf("fileInfo type: %T\n", fileInfo) // *os.fileStat
        //Size() 字节为单位
        fmt.Println(fileInfo.Name(), fileInfo.Size(), fileInfo.Mode())
    }

    {
        /*
        路径：
            相对路径：relative
                相对于当前工程
            绝对路径：
                /home/dzw/opt
         */
        fileName1 := "/home/dzw/opt/study/go/src/file/test.txt"
        fileName2 := "test.txt"

        // 是不是绝对路径
        fmt.Println(filepath.IsAbs(fileName1)) // true
        fmt.Println(filepath.IsAbs(fileName2)) // false

        fmt.Println(filepath.Abs(fileName1)) // /home/dzw/opt/study/go/src/file/test.txt <nil>
        fmt.Println(filepath.Abs(fileName2)) // /home/dzw/opt/study/go/src/file/test.txt <nil>

        fmt.Println("获取父目录:", path.Join(fileName1, "..")) // 获取父目录: /home/dzw/opt/study/go/src/file
    }

    {
        /*
        创建文件夹
         */
        
        // 文件存在，创建失败
        // os.Mkdir     创建一层
        // os.MkdirAll  创建多层

        // 单层文件夹
        err := os.Mkdir("./test_mkdir", os.ModePerm)
        if err != nil {
            fmt.Println("err:", err)
        } else {
            fmt.Println("create dir success")
        }

        // 多层文件夹
        /*err2 := os.Mkdir("./aa/bb/cc", os.ModePerm)
        if err2 != nil {
            fmt.Println("err:", err2) // ./aa/bb/cc: no such file or directory
        } else {
            fmt.Println("create more dir success")
        }*/

        err3 := os.MkdirAll("./aa/bb/cc", os.ModePerm)
        if err3 != nil {
            fmt.Println("err:", err3) 
        } else {
            fmt.Println("create more dir success")
        }
    }

    // file
    {
        // 文件存在会截断（为空）
        file1, err1 := os.Create("./aa.txt")
        if err1 != nil {
            fmt.Println("err:", err1)
        } else {
            fmt.Println(file1)
        }
    }

    // open file
    {
        file, err := os.Open("./aa.txt") // 只读
        if err != nil {
            fmt.Println("err:", err)
        } else {
            fmt.Println(file)
            file.Close()
        }
    }

    {
        file, err := os.OpenFile("./aa.txt", os.O_RDONLY | os.O_WRONLY, os.ModePerm) 
        if err != nil {
            fmt.Println("err:", err)
        } else {
            fmt.Println(file)
            file.Close() // 关闭文件
        }
    }

    // remove file / dir
    {
        err := os.Remove("./aa.txt")
        if err != nil {
            fmt.Println("err:", err)
        } else {
            fmt.Println("remove file success")
        }

        err = os.Remove("./test_mkdir")
        if err != nil {
            fmt.Println("err:", err) 
        } else {
            fmt.Println("remove dir success")
        }

        err = os.RemoveAll("./aa")
        if err != nil {
            fmt.Println("err:", err) 
        } else {
            fmt.Println("remove more dir success")
        }
    }
}