#!/bin/bash

:<<!
TYPE=Ethernet
BOOTPROTO=static
DEFROUTE=yes
PEERDNS=yes
PEERROUTES=yes
NAME=eth0
UUID=dfe6a427-ceeb-4681-8760-22ce0284addf
DEVICE=ens33
ONBOOT=yes
HWADDR=f6:98:1e:cb:4d:78


IPADDR=172.20.70.239
NETMASK=255.255.255.0
GATEWAY=172.20.70.1
!


echo -e "DEVICE=$1\nONBOOT=yes\nTYPE=$2\nBOOTPROTO=static\nIPADDR=$3\nNETMASK=$4\nGATEWAY=$5" > /etc/sysconfig/network-scripts/ifcfg-$1
#sudo service network restart
sleep 10


