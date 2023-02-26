#!/bin/bash

######默认HDMI为主显示#######

#HDMI+LVDS异显 主显示覆盖副显示。
# sudo su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --above LVDS-1" 

# HDMI+LVDS 同显 HDMI+LVDS显示相同内容（主显示器内容）
# sudo su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --same-as  LVDS-1" 


# HDMI+LVDS 异显   主显示器在副显示器左边。
sudo su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --left-of LVDS-1"

#HDMI+LVDS 异显  主显示器在副显示器右边。 
# sudo su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --right-of LVDS-1" 

# HDMI+LVDS 异显 主显示器在副显示器下面。
# sudo su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --below LVDS-1"