#!/bin/bash

id=$(id -u)

if [ $id -eq 0 ];then
                echo "Don't use root!"
                exit
else 
                echo "$id"
sudo apt update
# sudo apt install xfce4 xfce4-GOODIEs
sudo apt-get -y install xfonts-base
sudo apt install -y tightvncserver
mkdir /home/linaro/.vnc
vncserver -kill :1
/bin/mv /home/linaro/.vnc/xstartup /home/linaro/.vnc/xstartup.bak

cd
/bin/cat >> /home/linaro/.vnc/xstartup<<EOF
cat >>/home/linaro/.vnc/xstartup<<EOF
#!/bin/bash
lxterminal &
/usr/bin/lxsession -s LXDE &
EOF
sudo /bin/chmod +x ~/.vnc/xstartup
fi

exit
#xrdb \$HOME/.Xresources
#startxfce4 &
# ssh -p port -L 5901:127.0.0.1:5901 -C -N -l ubuntu serverip
# sudo vi /etc/systemd/system/vncserver@.service
# sudo systemctl daemon-reload
# sudosystemctl enable vncserver@1.service
#vncserver -kill :1
# sudo systemctl start vncserver@1
# sudo systemctl status vncserver@1
