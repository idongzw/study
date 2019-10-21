#!/usr/bin/env bash
# @File Name: shell_expand.sh
# @Author:    dzw
# @Date:      2019-10-21 20:34:37
# @Last Modified by:   dzw
# @Last Modified time: 2019-10-21 20:57:18

# Shell 扩展

:<<EOF
扩展 发生在一行命令被分成一个个的 记号（tokens） 之后。
换言之，扩展是一种执行数学运算的机制，还可以用来保存命令的执行结果，等等。
EOF

## 大括号扩展
# 大括号扩展让生成任意的字符串成为可能。它跟 文件名扩展 很类似

echo beg{i,a,u}n

# 大括号扩展还可以用来创建一个可被循环迭代的区间

echo {0..5}

## 命令置换
:<<EOF
命令置换允许我们对一个命令求值，并将其值置换到另一个命令或者变量赋值表达式中。
当一个命令被 `` 或 $() 包围时，命令置换将会执行
EOF

echo `date +%T`
echo $(date +%T)

## 算数扩展
# 在 bash 中，执行算数运算是非常方便的。算数表达式必须包在 $(( )) 中。

echo $(( (10 + 12 * 2) / 4 ))

# 在算数表达式中，使用变量无需带上 $ 前缀
x=4
y=7

echo $(( x + y ))
echo $(( ++x + y++ ))
echo $(( x + y ))

## 单引号和双引号
:<<EOF
单引号和双引号之间有很重要的区别。
在双引号中，变量引用或者命令置换是会被展开的。在单引号中是不会的。
EOF

echo "home = ${HOME}"
echo 'home = ${HOME}'

# 当局部变量和环境变量包含空格时，它们在引号中的扩展要格外注意

INPUT="A string  with   strange    whitespace."

echo ${INPUT}   # A string with strange whitespace.
echo "${INPUT}" # A string  with   strange    whitespace.

# 调用第一个 echo 时给了它 5 个单独的参数 —— $INPUT 被分成了单独的词，echo 在每个词之间打印了一个空格。
for item in ${INPUT}; do
    #statements
    echo "${item}"
done

# 第二种情况，调用 echo 时只给了它一个参数（整个 $INPUT 的值，包括其中的空格）。

for item in "${INPUT}"; do
    #statements
    echo "${item}"
done

FILE="Favorite Things.txt"
cat $FILE   ### 尝试输出两个文件: `Favorite` 和 `Things.txt`
cat "$FILE" ### 输出一个文件: `Favorite Things.txt`

:<<EOF
尽管这个问题可以通过把 FILE 重命名成Favorite-Things.txt来解决，
但是，假如这个值来自某个环境变量，来自一个位置参数，或者来自其它命令（find, cat, 等等）呢。
因此，如果输入可能包含空格，务必要用 "" 把表达式包起来。
EOF

cat << EOF > test_eof.txt # 重定向信息到 test_eof.txt 
# message
# 1
# 2
# 3
EOF

# cat test_eof.txt
# message
# 1
# 2
# 3

# 追加写
cat << EOF >> test_eof.txt
# message
# 1
# 2
# etc.
EOF

# 换一种写法
cat > test_eof.txt << EOF
# message
# 1
# 2
# 3
# 4
# 5
# 
EOF

# EOF只是标识，不是固定的，HHH也可以
# << HHH message HHH