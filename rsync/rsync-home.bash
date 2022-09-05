#!/bin/bash

# rsync  version 3.2.5  protocol version 31

openSUSE_home(){
		sudo rm /mnt/sdd1/rsync-home-openSUSE.log

		sudo rsync -aAXHv --delete-before --delete-excluded /home/userName /mnt/sdd1/rsync-backup-home-openSUSE/ --log-file=/mnt/sdd1/rsync-home-openSUSE.log --exclude={"/home/userName/.cache/*","/home/userName/.android/avd/*","/home/userName/.android/cache/*","/home/userName/.bundle/cache/*","/home/userName/.cache/VNote/VNote/QtWebEngine/Default/Cache/*","/home/userName/.cache/chromium/Default/Cache/*","/home/userName/.cantact/dev/var/cache/*","/home/userName/.config/mob/Cache/*","/home/userName/.config/vysor/Cache/*","/home/userName/.config/Code/Cache/*","/home/userName/.config/Slack/Cache/*","/home/userName/.config/libreoffice/4/user/uno_packages/cache/*","/home/userName/.config/libreoffice/4-suse/user/uno_packages/cache/*","/home/userName/.config/Typora/Cache/*","/home/userName/.config/Element/Cache/*","/home/userName/.config/SlackIntegrationTest/Cache/*","/home/userName/.googleearth/Cache/*","/home/userName/.local/share/flatpak/repo/tmp/cache/*","/home/userName/.local/share/TelegramDesktop/tdata/user_data/cache/*","/home/userName/.local/share/marble/cache/*","/home/userName/.rdesktop/cache/*","/home/userName/.vmodules/*","/home/userName/Blog/.deploy_git/*","/home/userName/Download/*","/home/userName/VirtualBox\ VMs/*","/home/userName/Sync/*","/home/userName/Android/Sdk/system-images/*","/home/userName/Android/Sdk/platforms/*"}
}


dir="/mnt/sdd1/rsync-backup-home-openSUSE"
t=$(date)

if [ -d "$dir" ];then
		echo "backup openSUSE Home....."
		openSUSE_home
else
		echo "$t 磁盘没有挂载" >> /home/userName/logs/rsync.log
fi
