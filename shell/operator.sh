#!/usr/bin/env bash
# @File Name: operator.sh
# @Author:    dzw
# @Date:      2019-10-19 22:11:57
# @Last Modified by:   dzw
# @Last Modified time: 2019-10-20 20:54:51

# 算术运算符

:<<EOF
假定变量 x 为 10，变量 y 为 20:

+   加法  expr $x + $y 结果为 30。
-   减法  expr $x - $y 结果为 -10。
*   乘法  expr $x * $y 结果为 200。
/   除法  expr $y / $x 结果为 2。
%   取余  expr $y % $x 结果为 0。
=   赋值  x=$y 将把变量 y 的值赋给 x。
==  相等。用于比较两个数字，相同则返回 true。 [ $x == $y ] 返回 false。
!=  不相等。用于比较两个数字，不相同则返回 true。   [ $x != $y ] 返回 true。
EOF

x=10
y=20

echo "x = $x; y = $y"

echo "$x + $y = `expr $x + $y`"
echo "$x - $y = `expr $x - $y`"
echo "$x * $y = `expr $x \* $y`"
echo "$y / $x = `expr $y / $x`"
echo "$y % $x = `expr $y % $x`"
echo "$x % $y = `expr $x % $y`"

echo "x = $x"
x=$y
echo "x = $x"

if [[ $x == $y ]]; then
    #statements
    echo "x == y"
fi

x=10

if [[ $x != $y ]]; then
    #statements
    echo "x != y"
fi

# 关系运算符

:<<EOF
关系运算符只支持数字，不支持字符串，除非字符串的值是数字

假定变量 x 为 10，变量 y 为 20:

-eq 检测两个数是否相等，相等返回 true。    [ $a -eq $b ]返回 false。
-ne 检测两个数是否相等，不相等返回 true。   [ $a -ne $b ] 返回 true。
-gt 检测左边的数是否大于右边的，如果是，则返回 true。 [ $a -gt $b ] 返回 false。
-lt 检测左边的数是否小于右边的，如果是，则返回 true。 [ $a -lt $b ] 返回 true。
-ge 检测左边的数是否大于等于右边的，如果是，则返回 true。   [ $a -ge $b ] 返回 false。
-le 检测左边的数是否小于等于右边的，如果是，则返回 true。   [ $a -le $b ]返回 true。
EOF

echo "----------------------------------------------"

x=10
y=20

echo "x = $x; y = $y"

if [[ $x -eq $y ]]; then
    #statements
    echo "$x -eq $y -> x == y"
else
    echo "$x -eq $y -> x != y"
fi

if [[ $x -ne $y ]]; then
    #statements
    echo "$x -ne $y -> x != y"
else
    echo "$x -ne $y -> x == y"
fi

if [[ $x -gt $y ]]; then
    #statements
    echo "$x -gt $y -> x > y"
else
    echo "$x -gt $y -> x < y"
fi

if [[ $x -lt $y ]]; then
    #statements
    echo "$x -lt $y -> x < y"
else
    echo "$x -lt $y -> x > y"
fi

if [[ $x -ge $y ]]; then
    #statements
    echo "$x -ge $y -> x >= y"
else
    echo "$x -ge $y -> x < y"
fi

if [[ $x -le $y ]]; then
    #statements
    echo "$x -le $y -> x <= y"
else
    echo "$x -le $y -> x > y"
fi


# 布尔运算符
# 逻辑运算符

:<<EOF
假定变量 x 为 10，变量 y 为 20:

布尔运算符：
!   非运算，表达式为 true 则返回 false，否则返回 true。  [ ! false ] 返回 true。
-o  或运算，有一个表达式为 true 则返回 true。  [ $a -lt 20 -o $b -gt 100 ] 返回 true。
-a  与运算，两个表达式都为 true 才返回 true。  [ $a -lt 20 -a $b -gt 100 ] 返回 false。

逻辑运算符：
&&  逻辑的 AND [[ ${x} -lt 100 && ${y} -gt 100 ]] 返回 false
||  逻辑的 OR  [[ ${x} -lt 100 || ${y} -gt 100 ]] 返回 true

EOF

echo "----------------------------------------------"

echo "x = $x; y = $y"

if [[ $x != $y ]]; then
    #statements
    echo "$x != $y -> x != y"
else
    echo "$x == $y -> x == y"
fi

:<<EOF
在使用 "[[]]" 时，不能使用 -o 或者 -a 对多个条件进行连接"
在使用 "[]" 时，如果使用 -o 或者 -a 对多个条件进行连接，-o 或者 -a 必须包含在 "[]" 之内
在使用 "[]" 时，如果使用 && 或者 || 对多个条件进行连接，&& 或者 || 必须在 "[]" 之外

# PS:   =~ 运算符判断变量是否满足某个正则表达式，只能用于 "[[]]" 中
EOF

if [ $x -lt 20 -a $x -gt 5 ]; then
    #statements
    echo "$x -lt 20 -a $x -gt 5 -> 5 < x < 20 "
else
    echo "x <= 5 || x >= 20"
fi

if [ $x -lt 20 ] && [ $x -gt 5 ]; then
    #statements
    echo "$x -lt 20 && $x -gt 5 -> 5 < x < 20 "
else
    echo "x <= 5 || x >= 20"
fi

if [[ $x -lt 20 ]] && [[ $x -gt 5 ]]; then
    #statements
    echo "$x -lt 20 && $x -gt 5 -> 5 < x < 20 "
else
    echo "x <= 5 || x >= 20"
fi

if [[ $x -lt 20 && $x -gt 5 ]]; then
    #statements
    echo "$x -lt 20 -a $x -gt 5 -> 5 < x < 20 "
else
    echo "x <= 5 || x >= 20"
fi

if [ $y -gt 50 -o $y -lt 30 ]; then
    echo "$y -gt 50 -o $y -lt 30 -> y > 50 || y < 30"
else
    echo " 30 <= x <= 50"
fi

# 字符串运算符

:<<EOF
假定变量 a 为 "abc"，变量 b 为 "efg":

=   检测两个字符串是否相等，相等返回 true。  [ $a = $b ] 返回 false。
!=  检测两个字符串是否相等，不相等返回 true。 [ $a != $b ] 返回 true。
-z  检测字符串长度是否为 0，为 0 返回 true。   [ -z $a ] 返回 false。
-n  检测字符串长度是否为 0，不为 0 返回 true。  [ -n $a ] 返回 true。
str 检测字符串是否为空，不为空返回 true。   [ $a ] 返回 true。
EOF

echo "----------------------------------------------"

x="abc"
y="efg"

echo "x = $x; y = $y"

if [[ $x = $y ]]; then
    #statements
    echo "$x = $y; x == y"
else
    echo "$x != $y; x != y"
fi

if [[ $x != $y ]]; then
    #statements
    echo "$x != $y; x != y"
else
    echo "$x = $y; x == y"
fi

if [[ -z $x ]]; then
    #statements
    echo "-z $x ; string length is 0"
else
    echo "-z $x ; string length is not 0"
fi

if [[ -n $x ]]; then
    #statements
    echo "-n $x ; string length is not 0"
else
    echo "-n $x ; string length is 0"
fi

if [[ $x ]]; then
    #statements
    echo "$x is not null"
else
    echo "$x is null"
fi

# 文件测试运算符

:<<EOF
文件测试运算符用于检测 Unix 文件的各种属性

-b file 检测文件是否是块设备文件，如果是，则返回 true。  [ -b $file ] 返回 false。
-c file 检测文件是否是字符设备文件，如果是，则返回 true。 [ -c $file ] 返回 false。
-d file 检测文件是否是目录，如果是，则返回 true。 [ -d $file ] 返回 false。
-f file 检测文件是否是普通文件（既不是目录，也不是设备文件），如果是，则返回 true。    [ -f $file ] 返回 true。
-g file 检测文件是否设置了 SGID 位，如果是，则返回 true。  [ -g $file ] 返回 false。
-k file 检测文件是否设置了粘着位(Sticky Bit)，如果是，则返回 true。  [ -k $file ]返回 false。
-p file 检测文件是否是有名管道，如果是，则返回 true。   [ -p $file ] 返回 false。
-u file 检测文件是否设置了 SUID 位，如果是，则返回 true。  [ -u $file ] 返回 false。
-r file 检测文件是否可读，如果是，则返回 true。  [ -r $file ] 返回 true。
-w file 检测文件是否可写，如果是，则返回 true。  [ -w $file ] 返回 true。
-x file 检测文件是否可执行，如果是，则返回 true。 [ -x $file ] 返回 true。
-s file 检测文件是否为空（文件大小是否大于 0），不为空返回 true。    [ -s $file ] 返回 true。
-e file 检测文件（包括目录）是否存在，如果是，则返回 true。    [ -e $file ] 返回 true。
EOF

echo "----------------------------------------------"

file="/etc/hosts"

echo "file: $file"

if [[ -b ${file} ]]; then
    #statements
    echo "${file} is a block device"
else
    echo "${file} is not a block device"
fi

if [[ -f ${file} ]]; then
    #statements
    echo "${file} is a normal file"
else
    echo "${file} is not a normal file"
fi
