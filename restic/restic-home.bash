#!/bin/bash

echo "restic backup home...."
sleep 5

sudo restic -r /mnt/sddx/restic-devices/home/ backup --tag noToolsVideo /home/ --exclude={"/path/Download/*","/path/.cache/*","/path/Android/Sdk/system-images/*","/path/Android/Sdk/platforms/*","/path/Sync/zysj/*","/path/.android/avd/*","/path/Sync/other/*","/path/go/pkg/mod/*","/path/Blog/.deploy_git/*","/path/nfs_share/*","/path/idevices/*","/path/tools/bilibili-down/*","/path/tools/ximalaya/*"} -v

sudo /sddx/restic-devices/home/ snapshots
sleep 3
sudo restic -r /mnt/sddx/restic-devices/home/ check
