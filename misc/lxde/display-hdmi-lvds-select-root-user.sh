#!/bin/bash

# export XAUTHORITY=/home/linaro/.Xauthority
# export DISPLAY=:0.0

echo $DISPLAY >> /home/linaro/ds.txt

W=$(whoami)
D=$(date)
WHO=$(id -u)
if [ WHO eq 0 ];then echo "$D $WHO $W is root ">> /home/linaro/isxx.txt
sudo -u linaro -s /bin/bash

# sudo su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --above LVDS-1" 
sudo su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --same-as  LVDS-1" 
# sudo su linaro -c "DISPLAY=:0.0 xrandr --output HDMI-1 --same-as  LVDS-1" 
# sudo su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --left-of LVDS-1" 
# sudo su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --right-of LVDS-1"
# sudo su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --below LVDS-1"

else
	echo "$D $WHO $W not root" >> /home/linaro/isxx.txt
sudo -u linaro -s /bin/bash
	echo "=========$D $WHO $Wi=======" >> /home/linaro/isxx.txt
# sudo su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --above LVDS-1" 
sudo su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --same-as  LVDS-1" 
# sudo su linaro -c "DISPLAY=:0.0 xrandr --output HDMI-1 --same-as  LVDS-1" 
# sudo su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --left-of LVDS-1" 
# sudo su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --right-of LVDS-1"
# sudo su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --below LVDS-1"
fi
