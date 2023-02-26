#!/bin/bash
# 进程监听并重启web服务
pgrep ccal-web &> /dev/null
if [ $? -gt 0 ]; #$?代表上面的返回值
then
echo "`date` restart" >> ./web.log #`date`是当前时间，>>输出到log文件中作为新的一行，log文件也可以与脚本放在同一个目录下
cd /home/liangzi/ccal && ./ccal-web >> ./web.log 2>&1  #进程启动命令
fi

# * * * * * /bin/listen.sh
