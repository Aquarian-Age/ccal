#!/usr/bin/expect -f

#######rk3288-Debian双屏异显########
######默认HDMI为主显示#######


#HDMI+LVDS异显 主显示覆盖副显示。
# spawn su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --above LVDS-1" 

# HDMI+LVDS 同显 HDMI+LVDS显示相同内容（主显示器内容）
spawn su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --same-as  LVDS-1" 


# HDMI+LVDS 异显   主显示器在副显示器左边。
# spawn su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --left-of LVDS-1"

#HDMI+LVDS 异显  主显示器在副显示器右边。 
# spawn su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --right-of LVDS-1" 

# HDMI+LVDS 异显 主显示器在副显示器下面。
# spawn  su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --below LVDS-1"

expect "Password:"
send "linaro\r"
interact

##################################################################
# spawn su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --above LVDS-1" 
# spawn su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --above LVDS-1" 
# spawn su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --same-as  LVDS-1" 
# spawn  su linaro -c "DISPLAY=:0 xrandr --output HDMI-1 --left-of LVDS-1"