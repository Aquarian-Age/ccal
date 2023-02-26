#!/bin/bash

#检查目录
if [ -d /home/path/dir ]; then 

echo "ok "; 
cd /home/path/dir

#jdk
if [ -f jdk-8u281-linux-x64.tar.gz ];then echo "ok"; mkdir jdk && tar xvf jdk-8u281-linux-x64.tar.gz -C ./jdk ;else echo "jdk-8u281-linux-x64.tar.gz not fond";fi

cd jdk/jdk1.8.0_281/
pwd=$(pwd)
cat >> /etc/profile<<EOF
JAVA_HOME=$(pwd)
CLASSPATH=\$JAVA_HOME/lib
PATH=\$PATH:\$JAVA_HOME/bin
export JAVA_HOME CLASSPATH PATH
EOF

echo "======================================"
source /etc/profile
java -version
cd -
else  echo "err";
fi

