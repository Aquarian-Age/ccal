#!/bin/sh

echo "进入转化的目录..."
pwd
sleep 2

ls * | while read CMD
do
echo $CMD

iconv -f GBK -t UTF-8 $CMD -o $CMD
done
