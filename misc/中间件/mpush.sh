#!/bin/bash


#alloc
cd /home/path/dir/
mkdir mpush
tar xvf alloc-release-x.y.z.tar.gz
cd mpush-alloc-x.y.z/
mv conf/mpush.conf conf/mpush.conf.bak
cp /home/path/dir/script/allloc-mpush.conf conf/mpush.conf
cd conf/
chmod +x ../bin/mp.sh 
source /etc/profile #need
sleep 3
../bin/mp.sh start
echo "=================================="
ps -ef | grep alloc
sleep 3

#mpush
cd /home/path/dir/mpush
tar xvf ../mpush-release-x.y.z.tar.gz 
cd mpush-x.y.z
mv conf/mpush.conf conf/mpush.conf.bak
cp /home/path/dir/script/mpush.conf conf/mpush.conf
cd conf/
source /etc/profile #need
sleep 2
./bin/mp.sh start
echo "===================================="
sleep netstat -nalpt | grep 38088

