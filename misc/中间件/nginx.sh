#!/bin/bash

cd /home/path/dir/
mkdir nginx && cd $_
mv ../pcre-8.43.tar.gz ../zlib-1.2.11.tar.gz ../nginx-1.18.0.tar.gz openssl-1.0.2t.tar.gz .
rpm -ivh ./gcc-c++/*.rpm #安装编译工具依赖

# pcre
cd /home/path/dir/nginx
tar xzvf pcre-8.43.tar.gz　　　　
cd /home/path/dir/nginx/pcre-8.43
./configure
make
make install
rpm -qa pcre
cd - 

# openssl
cd /home/path/dir/nginx
tar xvf openssl-1.0.2t.tar.gz 
cd openssl-1.0.2t
# make && make install
./config shared --prefix=/usr/local --openssldir=/usr/local/ssl
make clean
make && make install
openssl version -a
rpm -qa openssl
cd -

# zlib
cd /home/path/dir/nginx
tar xvf zlib-1.2.11.tar.gz 
cd zlib-1.2.11
CFLAGS="-O3 -fPIC" ./configure
make && make install
rpm -qa zlib
cd -

#NG
tar xvf nginx-1.18.0.tar.gz 
cd nginx-1.18.0

./configure --prefix=/home/path/dir/nginx/nginx-server \
--with-pcre=/home/path/dir/nginx/pcre-8.43 \
--with-zlib=/home/path/dir/nginx/zlib-1.2.11 \
--with-openssl=/home/path/dir/nginx/openssl-1.0.2t \
--with-http_ssl_module \
--with-http_flv_module \
--with-http_stub_status_module \
--with-http_gzip_static_module \
--http-client-body-temp-path=/home/path/dir/nginx/nginx-plug/client/ \
--http-proxy-temp-path=/home/path/dir/nginx/nginx-plug/proxy/ \
--http-fastcgi-temp-path=/home/path/dir/nginx/nginx-plug/fcgi/ \
--http-uwsgi-temp-path=/home/path/dir/nginx/nginx-plug/uwsgi \
--http-scgi-temp-path=/home/path/dir/nginx/nginx-plug/scgi \
--with-stream --with-stream_ssl_module

make && make install 
sleep 5
echo "======================================"
cd ../nginx-server/
mkdir -p ../nginx-plug/client
./sbin/nginx -t

#config
cd /home/path/dir/nginx/nginx-server
mv conf/nginx.conf conf/nginx.conf.bak
cp ../../script/nginx.conf conf/
./sbin/nginx -s reload
./sbin/nginx -c conf/nginx.conf


