#!/usr/bin/env bash
# @File Name: function.sh
# @Author:    dzw
# @Date:      2019-10-21 20:32:51
# @Last Modified by:   dzw
# @Last Modified time: 2019-10-21 21:02:27

# 函数

:<<EOF
bash 函数定义语法如下:

[ function ] funname [()] {
    action;
    [return int;]
}

说明：

函数定义时，function 关键字可有可无。
函数返回值 - return 返回函数返回值，返回值类型只能为整数（0-255）。如果不加 return 语句，shell 默认将以最后一条命令的运行结果，作为函数返回值。
函数返回值在调用该函数后通过 $? 来获得。
所有函数在使用前必须定义。这意味着必须将函数放在脚本开始部分，直至 shell 解释器首次发现它时，才可以使用。调用函数仅使用其函数名即可。
EOF

calc() {
    PS3="Choose the oper: "
    select oper in + - \* /
    do
        echo -n "enter first num: " && read x
        echo -n "enter second num: " && read y
        case ${oper} in
            "+" )
                return $(( x + y ))
                ;;
            "-" )
                return $(( x - y ))
                ;;
            "*" )
                return $(( x * y ))
                ;;
            "/" )
                return $(( x / y ))
                ;;
            * )
                echo "${oper} is not support"
                return 0
                ;;
        esac
        break
    done
}

calc
echo "The result is: $?"

# 位置参数
:<<EOF
位置参数是在调用一个函数并传给它参数时创建的变量。

位置参数变量表：
$0              脚本名称
$1 … $9         第 1 个到第 9 个参数列表
${10} … ${N}    第 10 个到 N 个参数列表
$* or $@        除了 $0 外的所有位置参数，"$*" 和 "$@" 不一样 
$#              不包括$0在内的位置参数的个数
$FUNCNAME       函数名称（仅在函数内部有值）
EOF

x=0
if [[ -n $1 ]]; then
  echo "第一个参数为：$1"
  x=$1
else
  echo "第一个参数为空"
fi

y=0
if [[ -n $2 ]]; then
  echo "第二个参数为：$2"
  y=$2
else
  echo "第二个参数为空"
fi

paramsFunction(){
  echo "函数第一个入参：$1"
  echo "函数第二个入参：$2"
}
paramsFunction ${x} ${y}

# 函数处理参数
:<<EOF
还有几个特殊字符用来处理参数:
$#  返回参数个数
$*  返回所有参数
$$  脚本运行的当前进程 ID 号
$!  后台运行的最后一个进程的 ID 号
$@  返回所有参数
$-  返回 Shell 使用的当前选项，与 set 命令功能相同。
$?  函数返回值
EOF

runner() {
  return 0
}

name=zp
paramsFunction(){
  echo "函数第一个入参：$1"
  echo "函数第二个入参：$2"
  echo "传递到脚本的参数个数：$#"
  echo "所有参数：("\$\*")"
  printf "+ %s\n" "$*"
  echo "脚本运行的当前进程 ID 号：$$"
  echo "后台运行的最后一个进程的 ID 号：$!"
  echo "所有参数：("\$\@")"
  printf "+ %s\n" "$@"
  echo "Shell 使用的当前选项：$-"
  runner
  echo "runner 函数的返回值：$?"
}
paramsFunction 1 "abc" "hello, \"zp\""