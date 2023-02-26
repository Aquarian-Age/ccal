## 如何重装GRUB？

使用UOS系统启动盘引导电脑启动，待进入安装界面后，按下Ctrl+Alt+F1，执行以下命令，稍等片刻，进入Live CD模式。
```bash
$ sudo service lightdm stop   
$ startx 
```

进入Live CD系统后打开终端，挂载需要修复系统的 / 挂载到/mnt，可以利用Gparted或者 sudo fdisk -l 命令查看，例如需要修复系统的/分区为/dev/sda1。
 执行以下命令：
```bash
$ sudo mount /dev/sda1 /mnt 
```
如果需要修复系统的/boot单独分了出来（假设为/dev/sda2），也要挂上，终端执行：
```bash
$ sudo mount /dev/sda2 /mnt/boot 
```
另外,将Live CD系统的/dev目录同时挂在/mnt下，终端执行：
```bash
$ sudo mount --bind /dev /mnt/dev 
```
然后使用chroot命令，将Live CD的 / 设为以前的/，终端执行：
```bash
$ sudo mount --bind /proc /mnt/proc 
$ sudo mount --bind /sys /mnt/sys 
$ sudo chroot /mnt 
```

安装并刷新GRUB设置(主板为BIOS引导)，请终端执行：
```bash
$ grub-probe -t device /boot/grub 
$ sudo grub-install /dev/sda 
$ sudo grub-install --recheck /dev/sda 
$ sudo update-grub
```

安装并刷新GRUB设置(主板为UEFI引导)： 启动root shell后，检查您的EFI系统分区（最可能 /dev/sda1）是否安装在 /boot /efi 上：
```bash
mount /dev/sda1 /boot/efi
```
重新安装grub-efi包：
```bash
apt-get install --reinstall grub-efi
```

将debian引导加载程序放在 /boot/efi 中，并在计算机NVRAM中创建一个适当的条目：
```bash
grub-install /dev/sda
```
重新创建一个基于磁盘分区模式的grub配置文件：
```bash
update-grub
```
挂载EFI分区到 /boot/efi，安装 grub-efi 包：
```bash
# grub-install --target=x86_64-efi --efi-directory=/boot/efi --bootloader-id=Deepin 
# grub-mkconfig -o /boot/grub/grub.cfg 
```
修复完成，重启电脑生效。 


[原文链接](https://ecology.chinauos.com/adaptidentification/knowledge/)

[knowledge](https://ecology.chinauos.com/adaptidentification/knowledge/)
