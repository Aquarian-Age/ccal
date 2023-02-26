#!/bin/bash

# redis
cd /home/path/dir/
tar xvf redis-x.y.z.tar
cd redis-x.y.z
make
sed -i '50s/6379/36379/g' redis.conf 
sed -i '42s/no/yes/g' redis.conf 
./src/redis-server redis.conf 
echo "====================="
sleep 5
netstat -nalpt |grep 16379
cd /home/path/dir/
