#!/bin/bash

iface=wlan0
ip_addr=$(ip -4 addr show $iface | grep -oP '(?<=inet\s)\d+(\.\d+){3}')

sysctl -w net.ipv4.ip_forward=1
sysctl -w net.ipv6.conf.all.forwarding=1
sysctl -w net.ipv4.conf.all.send_redirects=0

iptables -t nat -F PREROUTING

iptables -t nat -A PREROUTING -i $iface -p tcp --dport 80 -j REDIRECT --to-port 8080
iptables -t nat -A PREROUTING -i $iface -p tcp --dport 443 -j REDIRECT --to-port 8081
