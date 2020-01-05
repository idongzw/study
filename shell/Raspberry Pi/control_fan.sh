#!/usr/bin/env bash
# @File Name: control_fan.sh
# @Author:    dzw
# @Date:      2020-01-02 16:36:54
# @Last Modified by:   dzw
# @Last Modified time: 2020-01-05 17:28:18

# default
clear_temp=25
high_temp=35

if [[ $# -eq 2 ]]; then
    clear_temp=$1
    high_temp=$2
elif [[ $# -ne 0 ]]; then
    echo "Usage: $0 [clear_temp] [high_temp]"
    exit
fi

gpio18="/sys/class/gpio/gpio18"
gpio_export="/sys/class/gpio/export"
gpio_unexport="/sys/class/gpio/unexport"

if [[ -e ${gpio18} ]]; then
    echo "${gpio18} is exist."
else
    `echo 18 > ${gpio_export}`
    `echo out > ${gpio18}/direction`
fi

cpu_temp="/sys/class/thermal/thermal_zone0/temp"

while true; do

temp=`cat ${cpu_temp}`
temp=$((temp / 1000))

if [[ ${temp} -ge ${high_temp} ]]; then
    echo "[`date +%Y-%m-%d\ %H:%M:%S`]:Current cpu temperature: ${temp}"
    `echo 1 > ${gpio18}/value`
	sleep 5
elif [[ ${temp} -eq ${clear_temp} ]]; then
    echo "[`date +%Y-%m-%d\ %H:%M:%S`]:Current cpu temperature: ${temp}"
    `echo 0 > ${gpio18}/value`
	sleep 5
fi

sleep 1

done
