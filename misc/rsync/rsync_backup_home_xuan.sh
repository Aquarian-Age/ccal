#!/bin/sh

echo "......删除回收站 清理缓存......"
sudo rm -rf $HOME/.local/share/Trash/files/*
/bin/sh /home/userName/bin/clean-cache-dir.sh

sudo rsync -avrz --delete-excluded /home/ /mnt/sdchome/ --exclude={"OSX-KVM/*","Downloads/*","VirtualBox\ VMs/*","桌面/*",".cache/*",".nuget/*","go_work/pkg/mod/cache/","/home/userName/.wine/*"} #> /dev/null 2>&1 &
# -aptgovrlHAXzP
