#!/bin/bash

# rsync  version 3.2.5  protocol version 31

/usr/bin/mount /dev/sdd1 /mnt/sdd1
sleep 3

debian_home(){
		rsync -aAXHv --delete-excluded /home/xuan/Android /home/xuan/android-studio /home/xuan/OneDrive /home/xuan/Blog /home/xuan/ccal /home/xuan/flutter /home/xuan/Documents /home/xuan/Music /home/xuan/mysql /home/xuan/Picture /home/xuan/Sync /mnt/sdd1/rsync-backup-debian/
}

remove_log(){
		cd /home/xuan/
		/usr/bin/find .config/ -type f -iname "log" -exec rm {} \;
}
####################################################

dir="/mnt/sdd1/rsync-backup-debian"
t=$(date)

if [ -d "$dir" ];then
		echo "backup debian home....."
		sleep 5
		remove_log
		debian_home
else
		echo "$t 磁盘没有挂载" >> /home/xuan/Pub/rsync.log
fi

# https://wiki.archlinux.org/title/Rsync
# rsync -aAXHv --exclude={"/dev/*","/proc/*","/sys/*","/tmp/*","/run/*","/mnt/*","/media/*","/lost+found"} / /path/to/backup
