#!/usr/bin/env bash
# @File Name: control_fan.sh
# @Author:    dzw
# @Date:      2020-01-02 16:36:54
# @Last Modified by:   dzw
# @Last Modified time: 2020-01-05 17:28:18

# default
clear_temp=25
high_temp=35
log_file="/tmp/control_fan.log"
log_time="[`date +%Y-%m-%d\ %H:%M:%S`]: "

if [[ $# -eq 2 ]]; then
    clear_temp=$1
    high_temp=$2
elif [[ $# -ne 0 ]]; then
    echo "Usage: $0 [clear_temp] [high_temp]"
    exit
fi

echo ${log_time}"--------------------------------------" >> ${log_file}
echo ${log_time}"start control cpu temperature..." >> ${log_file}

gpio18="/sys/class/gpio/gpio18"
gpio_export="/sys/class/gpio/export"
gpio_unexport="/sys/class/gpio/unexport"

if [[ -e ${gpio18} ]]; then
    echo ${log_time}"${gpio18} is exist." >> ${log_file}
else
    echo ${log_time}"gpio export 18, direction --- out" >> ${log_file}
    `echo 18 > ${gpio_export}`
    `echo out > ${gpio18}/direction`
fi

cpu_temp="/sys/class/thermal/thermal_zone0/temp"

while true; do

temp=`cat ${cpu_temp}`
temp=$((temp / 1000))
log_time="[`date +%Y-%m-%d\ %H:%M:%S`]: "

if [[ ${temp} -ge ${high_temp} ]]; then
    echo ${log_time}"Current cpu temperature: ${temp}" >> ${log_file}
    `echo 1 > ${gpio18}/value`
    echo ${log_time}"gpio18 value --- 1" >> ${log_file}
	sleep 5
elif [[ ${temp} -eq ${clear_temp} ]]; then
    echo ${log_time}"Current cpu temperature: ${temp}" >> ${log_file}
    `echo 0 > ${gpio18}/value`
    echo ${log_time}"gpio18 value --- 0" >> ${log_file}
	sleep 5
fi

sleep 1

done
