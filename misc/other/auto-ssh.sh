#!/bin/bash

LANG=zh_CN.UTF-8
export LANG

ssh_copy_id_auto() {
    user=$name
	pass=$passwd
    host=$host
    port=$port
    key=$(echo ~/.ssh/id_rsa.pub)
expect << EOF
    set timeout 8;
    spawn ssh-copy-id -i $key -p $port $user@$host;
    expect {
        *(yes/no)?* { send yes\r; exp_continue; }
        *password:* { send $pass\r; exp_continue; }
        eof         { exit 0; }
    };
EOF
}

while :
	do
	echo "目前自动拷贝公钥只支持root用户"
	echo -n "请选择　1自动拷贝公钥到远程主机; 2SSH链接远程主机"
	echo " "
	read choice
	case $choice in
		1) #Copy ID
			echo "输入远程主机用户名:"
			read -p "" name
			echo "输入远程主机密码"
			read -p "" passwd
			echo "输入远程主机ip地址"
			read -p "" host
			echo "输入远程主机端口"
			read -p "" port
			
			echo "远程主机用户名 $name"
			echo "远程主机密码 $passwd"
			echo "远程主机ip地址 $host"
			echo "远程主机端口 $port"
			echo "Now start copy public_key to remote host..."
			sleep 3
			ssh_copy_id_auto
			break
			;;

		2) #SSH Connect Remote Host
			echo "输入远程主机用户名"
			read -p "" name
			echo "输入远程主机IP"
			read -p "" host
			echo "输入远程主机端口"
			read -p "" port
			echo "远程主机端口 $port"
			echo "开始SSH链接远程主机..."
			ssh -p $port $name@$host
			break
			;;
		*) # error
			echo "ERROR Try again?"
			echo "强制结束按Ctrl+C"
	esac	
done
