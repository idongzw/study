#!/usr/bin/env bash
# @File Name: control_statement.sh
# @Author:    dzw
# @Date:      2019-10-20 20:50:40
# @Last Modified by:   dzw
# @Last Modified time: 2019-10-21 20:43:47

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

exec # 有啥用呢？？
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

# for循环

for i in {1..5}; do # do 之前加上 ;
    #statements
    echo ${i}
done

for i in {1..5} 
do
    #statements
    echo ${i}
done

for (( i = 0; i < 10; i++ )); do
    #statements
    echo ${i}
done

for (( i = 0; i < 10; i++ ))
do
    #statements
    echo ${i}
done

DIR="."
for file in ${DIR}/*.sh; do
    #statements
    cp "${file}" "${DIR}/scripts/"
    chmod +x "${DIR}/scripts/${file}"
done

if [[ $? != 0 ]]; then
    #statements
    echo "cp failed."
else
    echo "cp done."
fi


# while循环

:<<EOF
while循环检测一个条件，只要这个条件为 真，就执行一段命令。
被检测的条件跟if..then中使用的基元并无二异。
EOF

x=0
while [[ ${x} -lt 10 ]]; do
    #statements
    echo $((x * x))
    x=${x}+1 # x+1
done

x=0
while [[ ${x} -lt 10 ]]
do
    #statements
    echo `expr ${x} \* ${x}`
    x=$(( x + 1 )) # x+1
done

# until循环

:<<EOF
until循环跟while循环正好相反。
它跟while一样也需要检测一个测试条件，但不同的是，只要该条件为 假 就一直执行循环
EOF

y=0
until [[ ${y} -ge 5 ]]; do
    #statements
    echo ${y}
    y=`expr ${y} + 1` #y+1
done

# select循环

:<<EOF
select循环帮助我们组织一个用户菜单。它的语法几乎跟for循环一致

select answer in elem1 elem2 ... elemN
do
  #statements
done

select会打印elem1..elemN以及它们的序列号到屏幕上，之后会提示用户输入。
通常看到的是 $?（PS3变量）。
用户的选择结果会被保存到answer中。如果answer是一个在1..N之间的数字，那么语句会被执行，紧接着会进行下一次迭代 —— 如果不想这样的话我们可以使用break语句。
EOF

PS3="Choose the package manager: "

select item in bower npm gem pip
do
    echo -n "Enter the package name: " && read package
    case ${item} in
        bower )
            bower install ${package}
            ;;
        npm )
            npm install ${package}
            ;;
        gem )
            gem install ${package}
            ;;
        pip )
            pip install ${package}
            ;;
        * )
            echo "select error"
            ;;
    esac
    break # 避免无限循环
done

# 循环控制

:<<EOF
如果想提前结束一个循环或跳过某次循环执行，可以使用 shell 的break和continue语句来实现。
它们可以在任何循环中使用。

break 语句用来提前结束当前循环。
continue 语句用来跳过某次迭代。
EOF

# 打印 10 以内第一个能整除 2 和 3 的数

echo "-------------------------------------"
x=1
while [[ ${x} -lt 10 ]]; do
    #statements
    if [[ $((x % 2)) -eq 0 ]] && [[ $((x % 3)) -eq 0 ]]; then
        #statements
        echo ${x}
        break
    fi
    x=$((x + 1))
done

# 打印 10 以内的奇数

echo "-------------------------------------"
for i in {0..10}; do
    #statements
    if [[ $((i % 2)) -eq 0 ]]; then
        #statements
        continue
    fi
    echo ${i}
done
