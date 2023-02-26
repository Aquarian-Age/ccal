#!/bin/bash
# A sample mysql backup script
# Must be run as the root user
# Written by Vivek Gite
# Last updated on : 23/Aug/2003
# ---------------------------------
# MySQL Login Info
MUSER="root" 			# MySQL user
MHOST="localhost"	# MySQL server ip
MPASS="abcdefg" 			# MySQL password

# format dd-mm-yyyy 
NOW=$(date +"%d-%m-%Y")

# Backupfile path
BPATH=/path/path/mysql/$NOW

# if backup path does not exists, create it 
[ ! -d $BPATH ] && mkdir -p $BPATH

# get database name lists
DBS="$(/usr/bin/mysql -u $MUSER -h $MHOST -p$MPASS -Bse 'show databases')"
echo $DBS	#测试链接是否成功

for db in $DBS
do
	# Bakcup file name
	FILE="${BPATH}/${db}.gz"
	
	# skip database backup if database name is adserverstats or mint
	[ "$db" == "adserverstats"  ] && continue
	[ "$db" == "mint"  ] && continue
	
	# okay lets dump a database backup 
    /usr/bin/mysqldump --single-transaction -u $MUSER -h $MHOST -p$MPASS $db | /bin/gzip -9 > $FILE
done

sleep 2
ls -lR ~/tftp/mysql
echo "数据库已备份到$BPATH"

