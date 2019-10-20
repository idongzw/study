#!/usr/bin/env bash
# @File Name: var.sh
# @Author:    dzw
# @Date:      2019-10-19 18:23:15
# @Last Modified by:   dzw
# @Last Modified time: 2019-10-19 18:55:27

:<<EOF
Bash 中没有数据类型，bash 中的变量可以保存一个数字、一个字符、一个字符串等等。同时无需提前声明变量，给变量赋值会直接创建变量。

// 变量命名原则
命名只能使用英文字母，数字和下划线，首个字符不能以数字开头。
中间不能有空格，可以使用下划线（_）。
不能使用标点符号。
不能使用 bash 里的关键字（可用 help 命令查看保留关键字）。

// 声明变量
访问变量的语法形式为：${var} 和 $var
变量名外面的花括号是可选的，加不加都行，加花括号是为了帮助解释器识别变量的边界，所以推荐加花括号
EOF

name="dzw"
echo ${name}

:<<EOF
// 只读变量
使用 readonly 命令可以将变量定义为只读变量，只读变量的值不能被改变。
EOF

r_var="readonly var"
echo ${r_var}
r_var="change var"
echo ${r_var}
readonly r_var
#r_var="change var" # r_var: readonly variable

readonly r_var2="readonly var2"
echo ${r_var2}
#r_var2="change var" # r_var2: readonly variable

:<<EOF
// 删除变量
使用 unset 命令可以删除变量。变量被删除后不能再次使用。unset 命令不能删除只读变量
EOF

unset_var="test unset"
echo ${unset_var}
unset unset_var
echo ${unset_var} # ouput null

:<<EOF
// 变量类型
局部变量 - 局部变量是仅在某个脚本内部有效的变量。它们不能被其他的程序和脚本访问。
环境变量 - 环境变量是对当前 shell 会话内所有的程序或脚本都可见的变量。创建它们跟创建局部变量类似，但使用的是 export 关键字，shell 脚本也可以定义环境变量。

常见的环境变量：
$HOME   当前用户的用户目录
$PATH   用分号分隔的目录列表，shell 会到这些目录中查找命令
$PWD    当前工作目录
$RANDOM 0 到 32767 之间的整数
$UID    数值类型，当前用户的用户 ID
$PS1    主要系统输入提示符
$PS2    次要系统输入提示符
EOF

echo $HOME
echo $PATH
echo $PWD
echo $RANDOM
echo $UID
echo $PS1
echo $PS2

