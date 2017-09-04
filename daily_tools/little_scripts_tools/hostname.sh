#!/bin/bash

oldname="`hostname`"
sudo hostname $1 
HOSTNAME=`hostname`
Network="/etc/sysconfig/network"
Hosts="/etc/hosts"
sudo sed -r -i 's/^HOSTNAME='${oldname}'.*$/HOSTNAME='`hostname`'/g' ${Network}
sudo sed -r -i 's/'${oldname}'/\ '`hostname`'\ /g;t;s/$/&\ '`hostname`'/g' ${Hosts}
sudo echo -e "hostname:{`hostname`}\n
Setting up hosts file:\n{\n`cat ${Hosts}`\n}\n
Setting up network file:\n{\n`cat ${Network}`\n}\n
HOSTNAME:\n{\n`echo ${HOSTNAME}`\n}\n"
echo -e "Change Hostname finished!"
