#!/bin/bash

#Zookeeper
cd /home/path/dir/
source /etc/profile
mkdir -p zookeeper
cd zookeeper

tar xvf ../apache-zookeeper-1.2.3-bin.tar.gz
cd apache-zookeeper-1.2.3-bin/
cp conf/zoo_sample.cfg conf/zoo.cfg
sed -i 's/2181/52181/g' conf/zoo.cfg
./bin/zkServer.sh start
echo "=================="
sleep 5
netstat -nalpt |grep 52181
