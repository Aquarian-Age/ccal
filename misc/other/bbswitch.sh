#!/bin/sh
set -v

echo "添加bbswitch源"
# zypper ar -f http://download.opensuse.org/repositories/X11:/Bumblebee/openSUSE_Leap_42.1/ blumblebee
echo "刷新并安装bbswitch相关软件"
# zypper ref
zypper in bbswitch dkms bbswitch-kmp-default
echo "禁用N卡"
echo "blacklist nouveau" >> /etc/modprobe.d/50-blacklist.conf
echo "设置开机加载模块"
echo "bbswitch" >> /etc/modules-load.d/modules.conf
echo "设置启动参数"
echo "options bbswitch load_state=0" > /etc/modprobe.d/50-bbswitch.conf
echo "重建引导"
mkinitrd
echo "输入reboot重启系统"
