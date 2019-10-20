#!/usr/bin/env bash
# @File Name: array.sh
# @Author:    dzw
# @Date:      2019-10-19 21:41:58
# @Last Modified by:   dzw
# @Last Modified time: 2019-10-19 22:17:11

# 数组
:<<EOF
bash 只支持一维数组。

数组下标从 0 开始，下标可以是整数或算术表达式，其值应大于或等于 0。
EOF
 
# 创建数组的不同方式
nums=([2]=2 [0]=0 [1]=1)
colors=(red yellow "dark blue")

# 访问数组的单个元素
echo ${nums} #echo ${nums[0]}
echo ${nums[1]}

# 访问数组的所有元素
echo ${nums[*]}
echo ${nums[@]}

echo ${colors[*]}
echo ${colors[@]}
# red yellow dark blue

printf "+ %s\n" ${colors[*]}
printf "+ %s\n" ${colors[@]}
:<<EOF
+ red
+ yellow
+ dark
+ blue
EOF

printf "+ %s\n" "${colors[*]}"
# + red yellow dark blue

printf "+ %s\n" "${colors[@]}"
:<<EOF
+ red
+ yellow
+ dark blue
EOF

# 在引号内，${colors[@]}将数组中的每个元素扩展为一个单独的参数；数组元素中的空格得以保留
for i in "${colors[@]}"; do
    echo $i
done
:<<EOF
red
yellow
dark blue
EOF

# 访问数组的部分元素
echo ${nums[@]:0:2}
# ${array[@]} 扩展为整个数组，:0:2取出了数组中从 0 开始，取 2 个元素

# 访问数组长度
echo ${#nums[*]}
# 3

# 向数组中添加元素
colors=(white "${colors[@]}" green block)
echo ${colors[@]}
# white red yellow dark blue green block
# ${colors[@]} 扩展为整个数组，并被置换到复合赋值语句中，接着，对数组colors的赋值覆盖了它原来的值

nums=(${nums[@]} 3 4 5 6 7)
echo ${nums[@]}
# 0 1 2 3 4 5 6 7

# 从数组中删除元素
unset nums[0]
echo ${nums[@]}
# 1 2 3 4 5 6 7