#!/bin/bash

id=$(id -u)

if [ $id -eq 0 ]
then

apt update
apt install vsftpd
cd /etc
mv vsftpd.conf vsftpd.conf.bak

cat >> /etc/vsftpd.conf<<EOF
listen=YES
connect_from_port_20=YES
anonymous_enable=YES
local_enable=YES
write_enable=YES
local_umask=022
anon_upload_enable=YES
anon_mkdir_write_enable=YES
dirmessage_enable=YES
use_localtime=YES
xferlog_enable=YES
utf8_filesystem=YES
EOF

/etc/init.d/vsftpd restart
/etc/init.d/vsftpd status
lsof -i:21
else 
echo "need root...."
fi

