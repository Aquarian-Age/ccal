@echo off
echo set sh=WScript.CreateObject("WScript.Shell") >telnet_tmp.vbs

echo WScript.Sleep 1000 >>telnet_tmp.vbs
rem ----------------UNIX IPAddress
echo sh.SendKeys "open 192.168.1.100{ENTER}" >>telnet_tmp.vbs

echo WScript.Sleep 1000 >>telnet_tmp.vbs
rem ----------------userID
echo sh.SendKeys "root{ENTER}" >>telnet_tmp.vbs

echo WScript.Sleep 1000 >>telnet_tmp.vbs
rem ----------------password
echo sh.SendKeys "登录密码{ENTER}" >>telnet_tmp.vbs

echo WScript.Sleep 1000 >>telnet_tmp.vbs

echo sh.SendKeys "/usr/sbin/你要执行的bin文件 {ENTER}">>telnet_tmp.vbs

start telnet

cscript //nologo telnet_tmp.vbs

rem del telnet_tmp.vbs
