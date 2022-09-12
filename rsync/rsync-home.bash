#!/bin/bash

# rsync  version 3.2.5  protocol version 31

openSUSE_home(){
		sudo rm /mnt/sdd1/rsync-home-openSUSE.log

		sudo rsync -aAXHv --delete-before --delete-excluded /home/userName /mnt/sdd1/rsync-backup-home-openSUSE/ --log-file=/mnt/sdd1/rsync-home-openSUSE.log --exclude={"userName/.cache/*","userName/.android/avd/*","userName/.android/cache/*","userName/.bundle/cache/*","userName/.cache/VNote/VNote/QtWebEngine/Default/Cache/*","userName/.cache/chromium/Default/Cache/*","userName/.cantact/dev/var/cache/*","userName/.config/mob/Cache/*","userName/.config/vysor/Cache/*","userName/.config/Code/Cache/*","userName/.config/Slack/Cache/*","userName/.config/libreoffice/4/user/uno_packages/cache/*","userName/.config/libreoffice/4-suse/user/uno_packages/cache/*","userName/.config/Typora/Cache/*","userName/.config/Element/Cache/*","userName/.config/SlackIntegrationTest/Cache/*","userName/.googleearth/Cache/*","userName/.local/share/flatpak/repo/tmp/cache/*","userName/.local/share/TelegramDesktop/tdata/user_data/cache/*","userName/.local/share/marble/cache/*","userName/.rdesktop/cache/*","userName/.vmodules/*","userName/Blog/.deploy_git/*","userName/Download/*","userName/VirtualBox\ VMs/*","userName/Android/Sdk/system-images/*","userName/Android/Sdk/platforms/*"}
}


dir="/mnt/sdd1/rsync-backup-home-openSUSE"
t=$(date)

if [ -d "$dir" ];then
		echo "backup openSUSE Home....."
		openSUSE_home
else
		echo "$t 磁盘没有挂载" >> userName/logs/rsync.log
fi
