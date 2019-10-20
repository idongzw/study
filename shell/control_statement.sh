#!/usr/bin/env bash
# @File Name: control_statement.sh
# @Author:    dzw
# @Date:      2019-10-20 20:50:40
# @Last Modified by:   dzw
# @Last Modified time: 2019-10-20 21:29:36

# 条件语句

:<<EOF
跟其它程序设计语言一样，Bash 中的条件语句让我们可以决定一个操作是否被执行。结果取决于一个包在[[ ]]里的表达式。

由[[ ]]（sh中是[ ]）包起来的表达式被称作 检测命令 或 基元。这些表达式帮助我们检测一个条件的结果。

共有两个不同的条件表达式：if 和 case
EOF

# if

## if 语句
# if在使用上跟其它语言相同。如果中括号里的表达式为真，那么 then 和 fi 之间的代码会被执行。
# fi标志着条件代码块的结束

if [[ 1 -eq 1 ]]; then
    #statements
    echo "1 == 1"
fi

## if else 语句

if [[ 1 -ne 2 ]]; then
    #statements
    echo "1 != 2"
else
    echo "1 == 2"
fi

## if elif else 语句

x=10
y=20

if [[ $x > $y ]]; then
    #statements
    echo "$x > $y"
elif [[ $x < $y ]]; then
    #statements
    echo "$x < $y"
else
    echo "$x == $y"
fi

echo "-----------------------------------------"

# case
# 如果你需要面对很多情况，分别要采取不同的措施，那么使用case会比嵌套的if更有用。
# 使用case来解决复杂的条件判断

oper="+"
if [[ -n $1 ]]; then
    #statements
    oper=$1
fi

exec // 有啥用呢？？
case ${oper} in
    "+" )
        echo "${oper}"
        ;;
    "-" )
        echo "${oper}"
        ;;
    "*" )
        echo "${oper}"
        ;;
    "/" )
        echo "${oper}"
        ;;
    * )
        echo "unknown oper"
        ;;
esac

:<<EOF
每种情况都是匹配了某个模式的表达式。| 用来分割多个模式，) 用来结束一个模式序列。
第一个匹配上的模式对应的命令将会被执行。
* 代表任何不匹配以上给定模式的模式。命令块儿之间要用 ;; 分隔。
EOF

# 循环语句

# Bash 中有四种循环：for，while，until和select