#!/bin/bash

##18080
iptables -A INPUT -p tcp --dport 18080 -j ACCEPT
iptables -A OUTPUT -p tcp --sport 18080 -j ACCEPT

iptables -A INPUT -p udp --dport 18080 -j ACCEPT
iptables -A OUTPUT -p udp --sport 18080 -j ACCEPT

##18081
iptables -A INPUT -p tcp --dport 18081 -j ACCEPT
iptables -A OUTPUT -p tcp --sport 18081 -j ACCEPT

iptables -A INPUT -p udp --dport 18081 -j ACCEPT
iptables -A OUTPUT -p udp --sport 18081 -j ACCEPT

##19000
iptables -A INPUT -p tcp --dport 19000 -j ACCEPT
iptables -A OUTPUT -p tcp --sport 19000 -j ACCEPT

iptables -A INPUT -p udp --dport 19000 -j ACCEPT
iptables -A OUTPUT -p udp --sport 19000 -j ACCEPT

## 11180
iptables -A INPUT -p tcp --dport 11180 -j ACCEPT
iptables -A OUTPUT -p tcp --sport 11180 -j ACCEPT

iptables -A INPUT -p udp --dport 11180 -j ACCEPT
iptables -A OUTPUT -p udp --sport 11180 -j ACCEPT

##5206
iptables -A INPUT -p tcp --dport 5206 -j ACCEPT
iptables -A OUTPUT -p tcp --sport 5206  -j ACCEPT

iptables -A INPUT -p udp --dport 5206 -j ACCEPT
iptables -A OUTPUT -p udp --sport 5206 -j ACCEPT

## 18087,18088
iptables -A INPUT -p tcp --dport 18087:18088 -j ACCEPT
iptables -A OUTPUT -p tcp --sport 18087:18088 -j ACCEPT

iptables -A INPUT -p udp --dport 18087:18088 -j ACCEPT
iptables -A OUTPUT -p udp --sport 18087:18088 -j ACCEPT

#DROP##

##18089
iptables -A INPUT -p tcp --dport 18089 -j DROP
iptables -A INPUT -p udp --dport 18089 -j DROP

## 17000，17001
iptables -A INPUT -p tcp --dport 17000:17001 -j DROP
iptables -A INPUT -p udp --dport 17000:17001 -j DROP

##12181
iptables -A INPUT -p tcp --dport 12181 -j DROP
iptables -A INPUT -p udp --dport 12181 -j DROP

##16379
iptables -A INPUT -p tcp --dport 16379 -j DROP
iptables -A INPUT -p udp --dport 16379 -j DROP

##19999
iptables -A INPUT -p tcp --dport 19999 -j DROP
iptables -A INPUT -p udp --dport 19999 -j DROP

##30000
iptables -A INPUT -p tcp --dport 30000 -j DROP
iptables -A INPUT -p udp --dport 30000 -j DROP


service iptables save 
iptables -L -n
