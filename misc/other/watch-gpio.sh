#!/bin/sh

id=$(id -u)

if [ $id -eq 0 ]
then
cd /sys/class/gpio/
for i in gpiochip* ; do echo `cat $i/label`: `cat $i/base` ; done  
else
	echo "need root login"
fi
