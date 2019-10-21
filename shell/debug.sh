#!/usr/bin/env bash
# @File Name: debug.sh
# @Author:    dzw
# @Date:      2019-10-21 20:36:43
# @Last Modified by:   dzw
# @Last Modified time: 2019-10-21 20:43:26

# Debug
:<<EOF
shell 提供了用于 debug 脚本的工具。
如果想采用 debug 模式运行某脚本，可以在其 shebang 中使用一个特殊的选项：

#!/bin/bash options

options 是一些可以改变 shell 行为的选项。下表是一些可能对你有用的选项：

-f  noglob  禁止文件名展开（globbing）
-i  interactive 让脚本以交互模式运行
-n  noexec  读取命令，但不执行（语法检查）
-t  —   执行完第一条命令后退出
-v  verbose 在执行每条命令前，向 stderr 输出该命令
-x  xtrace  在执行每条命令前，向 stderr 输出该命令以及该命令的扩展参数
EOF

:<<EOF
如果我们在脚本中指定了-x例如：

#!/bin/bash -x

for (( i = 0; i < 3; i++ )); do
  echo $i
done

这会向 stdout 打印出变量的值和一些其它有用的信息：

+ (( i = 0 ))
+ (( i < 3 ))
+ echo 0
0
+ (( i++   ))
+ (( i < 3 ))
+ echo 1
1
+ (( i++   ))
+ (( i < 3 ))
+ echo 2
2
+ (( i++   ))
+ (( i < 3 ))

EOF

:<<EOF
有时我们值需要 debug 脚本的一部分。这种情况下，使用set命令会很方便。
这个命令可以启用或禁用选项。使用 - 启用选项，+ 禁用选项：
EOF

set -x # 开启 debug

for (( i = 0; i < 3; i++ )); do
  printf ${i}
done

set +x # 关闭 debug

for i in {1..5}; do printf ${i}; done
printf "\n"