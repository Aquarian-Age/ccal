#!/bin/sh

### sshpass 面密码输入远程备份
### sshpass -p后面跟ssh的密码　ssh -p 后面跟远程端口

rsync -aptgovrlHAXzP --log-file=/path/wikijs_vps.log --delete-excluded --exclude={"/swap","/media/*","/sys/*","/proc/*","/mnt/*","/tmp/*","/dev/*","/home/other/*","/var/run/*","/var/tmp/*","/usr/share/doc/*"} -e "sshpass -p password ssh -p port" root@1.2.8.2:/ /path/wikijs_vps/
