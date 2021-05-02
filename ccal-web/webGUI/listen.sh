#!/bin/bash
# 进程监听并重启web服务
pgrep -f ccal-web
if [ $? -eq 0 ]; #$?代表上面的返回值
then
        cd /home/xuan/ccal/web/
        echo " "> ./web.log #清空上次log
        ./ccal-web & >> ./web.log 2>&1  #进程启动命令
fi

