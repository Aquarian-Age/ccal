#!/bin/bash

# rsync  version 3.2.5  protocol version 31

/usr/bin/mount /dev/sdd1 /mnt/sdd1
sleep 3

debian(){
	rm /mnt/sdd1/rsync-debian.log

	rsync -aptgovrlHAX --partial --delete-excluded / /mnt/sdd1/rsync-backup-debian/ --log-file=/mnt/sdd1/rsync-debian.log --exclude={"/dev/*","/proc/*","/sys/*","/tmp/*","/run/*","/mnt/*","/media/*","/var/log/*","/lost+found","/home/user/.android/cache/*","/home/user/.android/avd","/home/user/.cache/*","/home/user/.config/Slack/Cache/*","/home/user/.config/Slack/Service\ Worker/CacheStorage","/home/user/.config/Code/Cache/*","/home/user/.dartServer/.analysis-driver","/home/user/.gradle/wrapper/dists/gradle-6.1.1-all/cfmwm155h49vnt3hynmlrsdst/gradle-6.1.1/src/internal-integ-testing/org/gradle/integtests/fixtures/cache/*","/home/user/.gradle/wrapper/dists/gradle-6.1.1-all/cfmwm155h49vnt3hynmlrsdst/gradle-6.1.1/src/dependency-management/org/gradle/api/internal/artifacts/cache/*","/home/user/.gradle/wrapper/dists/gradle-6.1.1-all/cfmwm155h49vnt3hynmlrsdst/gradle-6.1.1/src/model-core/org/gradle/model/internal/manage/schema/cache/*","/home/user/.gradle/wrapper/dists/gradle-6.1.1-all/cfmwm155h49vnt3hynmlrsdst/gradle-6.1.1/src/language-java/org/gradle/api/internal/tasks/compile/incremental/cache/*","/home/user/.gradle/wrapper/dists/gradle-6.1.1-all/cfmwm155h49vnt3hynmlrsdst/gradle-6.1.1/src/core/org/gradle/cache/*","/home/user/.gradle/wrapper/dists/gradle-6.1.1-all/cfmwm155h49vnt3hynmlrsdst/gradle-6.1.1/src/base-services/org/gradle/api/internal/cache/*","/home/user/.gradle/wrapper/dists/gradle-6.1.1-all/cfmwm155h49vnt3hynmlrsdst/gradle-6.1.1/src/kotlin-dsl/org/gradle/kotlin/dsl/cache/*","/home/user/.gradle/wrapper/dists/gradle-6.1.1-all/cfmwm155h49vnt3hynmlrsdst/gradle-6.1.1/src/persistent-cache/org/gradle/cache/*","/home/user/.gradle/wrapper/dists/gradle-5.6.2-all/9st6wgf78h16so49nn74lgtbb/gradle-5.6.2/src/internal-integ-testing/org/gradle/integtests/fixtures/cache/*","/home/user/.gradle/wrapper/dists/gradle-5.6.2-all/9st6wgf78h16so49nn74lgtbb/gradle-5.6.2/src/dependency-management/org/gradle/api/internal/artifacts/cache/*","/home/user/.gradle/wrapper/dists/gradle-5.6.2-all/9st6wgf78h16so49nn74lgtbb/gradle-5.6.2/src/model-core/org/gradle/model/internal/manage/schema/cache/*","/home/user/.gradle/wrapper/dists/gradle-5.6.2-all/9st6wgf78h16so49nn74lgtbb/gradle-5.6.2/src/language-java/org/gradle/api/internal/tasks/compile/incremental/cache/*","/home/user/.gradle/wrapper/dists/gradle-5.6.2-all/9st6wgf78h16so49nn74lgtbb/gradle-5.6.2/src/core/org/gradle/cache/*","/home/user/.gradle/wrapper/dists/gradle-5.6.2-all/9st6wgf78h16so49nn74lgtbb/gradle-5.6.2/src/base-services/org/gradle/api/internal/cache/*","/home/user/.gradle/wrapper/dists/gradle-5.6.2-all/9st6wgf78h16so49nn74lgtbb/gradle-5.6.2/src/kotlin-dsl/org/gradle/kotlin/dsl/cache/*","/home/user/.gradle/wrapper/dists/gradle-5.6.2-all/9st6wgf78h16so49nn74lgtbb/gradle-5.6.2/src/persistent-cache/org/gradle/cache/*","/home/user/.gradle/wrapper/dists/gradle-5.6.2-all/9st6wgf78h16so49nn74lgtbb/gradle-5.6.2/samples/userguide/dependencyManagement/troubleshooting/cache/*","/home/user/.gradle/wrapper/dists/gradle-6.7-all/cuy9mc7upwgwgeb72wkcrupxe/gradle-6.7/src/dependency-management/org/gradle/api/internal/artifacts/cache/*","/home/user/.gradle/wrapper/dists/gradle-6.7-all/cuy9mc7upwgwgeb72wkcrupxe/gradle-6.7/src/model-core/org/gradle/model/internal/manage/schema/cache/*","/home/user/.gradle/wrapper/dists/gradle-6.7-all/cuy9mc7upwgwgeb72wkcrupxe/gradle-6.7/src/language-java/org/gradle/api/internal/tasks/compile/incremental/cache/*","/home/user/.gradle/wrapper/dists/gradle-6.7-all/cuy9mc7upwgwgeb72wkcrupxe/gradle-6.7/src/core/org/gradle/cache/*","/home/user/.gradle/wrapper/dists/gradle-6.7-all/cuy9mc7upwgwgeb72wkcrupxe/gradle-6.7/src/base-services/org/gradle/api/internal/cache/*","/home/user/.gradle/wrapper/dists/gradle-6.7-all/cuy9mc7upwgwgeb72wkcrupxe/gradle-6.7/src/kotlin-dsl/org/gradle/kotlin/dsl/cache/*","/home/user/.gradle/wrapper/dists/gradle-6.7-all/cuy9mc7upwgwgeb72wkcrupxe/gradle-6.7/src/persistent-cache/org/gradle/cache/*","/home/user/.local/share/TelegramDesktop/tdata/user_data/cache/*","/home/user/.mozilla/firefox/fpiv7k1m.default-esr/storage/default/https+++cn.bing.com/cache/*","/home/user/.mozilla/firefox/fpiv7k1m.default-esr/storage/default/https+++forum.manjaro.org/cache/*","/home/user/.mozilla/firefox/fpiv7k1m.default-esr/storage/default/https+++blog.csdn.net/cache/*","/home/user/.vim/bundle/YouCompleteMe/third_party/ycmd/third_party/go/pkg/mod/golang.org/x/tools@v0.1.13-0.20220812184215-3f9b119300de/internal/lsp/cache/*","/home/user/.vim/bundle/YouCompleteMe/third_party/ycmd/third_party/go/pkg/mod/cache/*","/home/user/.vim/bundle/YouCompleteMe/third_party/ycmd/third_party/go/pkg/mod/honnef.co/go/tools@v0.3.2/lintcmd/cache/*","/home/user/.vim/bundle/YouCompleteMe/third_party/ycmd/third_party/jedi_deps/jedi/jedi/third_party/django-stubs/django-stubs/core/cache/*","/home/user/Blog/.deploy_git","/home/user/Android/Sdk/platforms/*
","/home/user/flutter/.pub-cache/*","/home/user/.pub-cache/*","/home/user/.npm/_cacache/*"}
}

remove_log(){
	cd /home/user/
	/usr/bin/find .config/ -type f -iname "log" -exec rm {} \;
}
####################################################

dir="/mnt/sdd1/rsync-backup-debian"
t=$(date)

if [ -d "$dir" ];then
		echo "backup debian 11....."
		debian
else
		echo "$t 磁盘没有挂载" >> /home/user/Pub/rsync.log
fi

# https://wiki.archlinux.org/title/Rsync
# rsync -aAXHv --exclude={"/dev/*","/proc/*","/sys/*","/tmp/*","/run/*","/mnt/*","/media/*","/lost+found"} / /path/to/backup
