#!/usr/bin/env bash
# @File Name: stream_and_redirect.sh
# @Author:    dzw
# @Date:      2019-10-21 20:35:54
# @Last Modified by:   dzw
# @Last Modified time: 2019-10-21 20:42:06

# 流和重定向

:<<EOF
Bash 有很强大的工具来处理程序之间的协同工作。使用流，我们能将一个程序的输出发送到另一个程序或文件，因此，我们能方便地记录日志或做一些其它我们想做的事。

管道给了我们创建传送带的机会，控制程序的执行成为可能。
EOF

## 输入、输出流

:<<EOF
# Bash 接收输入，并以 字符序列 或 字符流 的形式产生输出。这些流能被重定向到文件或另一个流中。
有三个文件描述符：

代码  描述符 描述
0   stdin   标准输入
1   stdout  标准输出
2   stderr  标准错误输出

EOF

## 重定向

:<<EOF
重定向让我们可以控制一个命令的输入来自哪里，输出结果到什么地方。这些运算符在控制流的重定向时会被用到：

>   重定向输出
&>  重定向输出和错误输出
&>> 以附加的形式重定向输出和错误输出
<   重定向输入
<<  Here 文档语法
<<< Here 字符串

EOF

:<<EOF
### ls的结果将会被写到list.txt中
ls -l > list.txt

### 将输出附加到list.txt中
ls -a >> list.txt

### 所有的错误信息会被写到errors.txt中
grep da * 2> errors.txt

### 从errors.txt中读取输入
less < errors.txt
EOF

## /dev/null 文件
:<<EOF
# 如果希望执行某个命令，但又不希望在屏幕上显示输出结果，那么可以将输出重定向到 /dev/null：

command > /dev/null

/dev/null 是一个特殊的文件，写入到它的内容都会被丢弃；如果尝试从该文件读取内容，那么什么也读不到。
但是 /dev/null 文件非常有用，将命令的输出重定向到它，会起到"禁止输出"的效果。

如果希望屏蔽 stdout 和 stderr，可以这样写：

command > /dev/null 2>&1
EOF