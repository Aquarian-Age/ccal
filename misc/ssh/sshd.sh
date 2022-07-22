#!/bin/bash

id=$(id -u)

if [ $id -eq 0 ]
then
cd /etc/ssh/
mv sshd_config sshd_config.bak
cat >>/etc/ssh/sshd_config<<EOF
Port 22
AddressFamily any
ListenAddress 0.0.0.0
ListenAddress ::
PubkeyAuthentication yes
UsePAM yes
X11Forwarding yes
Subsystem       sftp    /usr/lib/openssh/sftp-server
EOF

systemctl stop sshd
systemctl start sshd
systemctl status sshd
else
echo "need root"
fi
