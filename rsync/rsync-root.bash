#!/bin/bash

# rsync  version 3.2.5  protocol version 31

openSUSE_root(){
		sudo rm /mnt/sdd1/rsync-root-openSUSE.log

		sudo rsync -aAXHv --delete-excluded / /mnt/sdd1/rsync-backup-root-openSUSE/ --log-file=/mnt/sdd1/rsync-root-openSUSE.log --exclude={"/home/*","/dev/*","/proc/*","/sys/*","/tmp/*","/run/*","/mnt/*","/media/*","/lost+found","/var/log/*"}
}


dir="/mnt/sdd1/rsync-backup-root-openSUSE"
t=$(date)

if [ -d "$dir" ];then
		echo "backup openSUSE root....."
		echo " "
		openSUSE_root
else
		echo "$t 磁盘没有挂载" >> /home/userName/logs/rsync.log
fi
