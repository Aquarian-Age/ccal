#!/bin/bash
echo "restic backup root ...."
sleep 5

sudo restic -r /mnt/sddx/restic-devices/root/ backup --tag "kernel-5.19.8" --one-file-system / --exclude={"/home/","/var/log/*","/var/adm/backup"} --verbose

sudo restic -r /mnt/sddx/restic-devices/efi/

echo "snapshots root...."
sleep 3
sudo restic -r /mnt/sddx/restic-devices/root/ snapshots

echo "check root..."
sleep 3
sudo restic -r /mnt/sddx/restic-devices/root/ check
