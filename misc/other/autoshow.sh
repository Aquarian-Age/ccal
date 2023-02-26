#!/bin/sh 

PWD=$(pwd)
cd /usr/local/bin

echo $PQD
if [ "$PWD" = "/usr/local/bin" ];then 
echo "oK"
echo "................."
echo "above"
sed -i '3s/^\#\ //g' display-hdmi-lvds.sh ## 启用第3行
cat display-hdmi-lvds.sh
sleep 5
sudo systemctl start display-dual.service
sleep 5

echo "same-as"
sed -i '3s/^/\#\ /g' display-hdmi-lvds.sh ### 注释第3行 
sudo systemctl start display-dual.service
sed -i '4s/^\#\ //g' display-hdmi-lvds.sh
cat display-hdmi-lvds.sh
sleep 5

echo "left-of"
sed -i '4s/^/\#\ /g' display-hdmi-lvds.sh ### 注释第4行 
sudo systemctl start display-dual.service
sed -i '5s/^\#\ //g' display-hdmi-lvds.sh
cat display-hdmi-lvds.sh
sleep 5

echo "right-of"
sed -i '5s/^/\#\ /g' display-hdmi-lvds.sh ### 注释第5行 
sudo systemctl start display-dual.service
sed -i '6s/^\#\ //g' display-hdmi-lvds.sh
cat display-hdmi-lvds.sh
sleep 5

echo "below"
sed -i '6s/^/\#\ /g' display-hdmi-lvds.sh ### 注释第6行 
sudo systemctl start display-dual.service
sed -i '7s/^\#\ //g' display-hdmi-lvds.sh
cat display-hdmi-lvds.sh
	
sed -i '7s/^/\#\ /g' display-hdmi-lvds.sh ### 注释第7行 
else 
	echo "bad"
fi
