#！/bin/bash

#tomcat
cd /home/path/dir
if [ -f apache-tomcat-x.y.z.tar.gz ];then
mkdir tomcat && tar xvf apache-tomcat-x.y.z.tar.gz -C ./tomcat;
### 配置
cd tomcat/apache-tomcat-x.y.z/
cp conf/server.xml conf/server.xml.bak
sed -i '69s/8080/98080/g' conf/server.xml ### 依据具体行号指定
source /etc/profile

./bin/startup.sh

sleep 5
netstat -nalpt |grep 98080
cd -
else
echo "apache-tomcat-x.y.z.tar.gz not found"
fi
