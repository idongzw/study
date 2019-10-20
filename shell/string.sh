#!/usr/bin/env bash
# @File Name: string.sh
# @Author:    dzw
# @Date:      2019-10-19 18:50:13
# @Last Modified by:   dzw
# @Last Modified time: 2019-10-19 21:47:52

:<<EOF
shell 字符串可以用单引号 ''，也可以用双引号 ""，也可以不用引号。

单引号的特点:
    - 单引号里不识别变量
    - 单引号里不能出现单独的单引号（使用转义符也不行），但可成对出现，作为字符串拼接使用。
双引号的特点:
    - 双引号里识别变量
    - 双引号里可以出现转义字符

推荐使用双引号
EOF

# 使用单引号拼接
name1='dzw'
str1='hello, '${name1}''
str2='hello, ${name1}'

echo ${str1}_${str2}

# hello, dzw_hello, ${name1}

# 使用双引号拼接
name2="dzw"
str3="hello, "${name2}""
str4="hello, ${name2}"

echo ${str3}_${str4}

# hello, dzw_hello, dzw

# 获取字符串长度
text="12345"
echo "\"${text}\" length is ${#text}"

# 截取子字符串
echo ${text:2:2} # 从第 3 个字符开始，截取 2 个字符

# 查找子字符串
text1="hello"
echo `expr index "${text1}" ll` # 查找 ll 子字符在 hello 字符串中的起始位置

# 3

