##### 设置语言环境为en_US.utf8 需要root密码

#!/bin/bash

lang=$(echo $LANG)
echo $lang
sleep 5

en="en_US.utf8"
echo $en

if [ "$lang" = "$en" ]
then
	reture 1
elif [ "$lang" != "$en" ]
then
	localectl set-locale LANG=en_US.utf8 
fi
