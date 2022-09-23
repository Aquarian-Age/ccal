#!/bin/bash

echo "restic backup efi...."
sleep 5
 
sudo restic -r /mnt/sdd1/restic-devices/efi/ backup --tag EFI /boot/efi/ -v
