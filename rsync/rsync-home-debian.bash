#!/bin/bash

# rsync  version 3.2.5  protocol version 31

/usr/bin/mount /dev/sdd1 /mnt/sdd1
sleep 3

debian_home(){
		rsync -aAXHv --delete-excluded /home/user/Android /home/user/android-studio /home/user/OneDrive /home/user/Blog /home/user/ccal /home/user/flutter /home/user/Documents /home/user/Music /home/user/mysql /home/user/Picture /home/user/Sync /mnt/sdd1/rsync-backup-debian/
}

remove_log(){
		cd /home/user/
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
		echo "$t 磁盘没有挂载" >> /home/user/Pub/rsync.log
fi

# https://wiki.archlinux.org/title/Rsync
# rsync -aAXHv --exclude={"/dev/*","/proc/*","/sys/*","/tmp/*","/run/*","/mnt/*","/media/*","/lost+found"} / /path/to/backup
