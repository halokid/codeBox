#!/bin/bash
# From here: http://www.codingsteps.com/install-redis-2-6-on-amazon-ec2-linux-ami-or-centos/
# Based on: https://github.com/saxenap/install-redis-amazon-linux-centos
# Thanks to https://raw.github.com/gist/2776679/b4f5f5ff85bddfa9e07664de4e8ccf0e115e7b83/install-redis.sh
# Uses redis-server init script from https://raw.github.com/saxenap/install-redis-amazon-linux-centos/master/redis-server
###############################################
# To use: 
## wget https://raw.github.com/jorgerc/install-redis-amazon-linux-centos/master/redis-install-script.sh
## chmod 777 redis-install-script.sh
## ./redis-install-script.sh
###############################################
# Set up SO:
####
yum -y update
ln -sf /usr/share/zoneinfo/UTC /etc/localtime
yum -y install gcc gcc-c++ make 
####
# Download and install Redis:
####
wget -q http://download.redis.io/releases/redis-2.8.24.tar.gz 
tar xzf redis-2.8.24.tar.gz
rm -f redis-2.8.24.tar.gz
cd redis-2.8.24
make
make install
####
# Set up Redis
####
rm -rf /etc/redis /var/lib/redis
mkdir /etc/redis /var/lib/redis
cp src/redis-server src/redis-cli /usr/local/bin
cp redis.conf /etc/redis
sed -e "s/^daemonize no$/daemonize yes/" -e "s/^# bind 127.0.0.1$/bind 127.0.0.1/" -e "s/^dir \.\//dir \/var\/lib\/redis\//" -e "s/^loglevel verbose$/loglevel notice/" -e "s/^logfile stdout$/logfile \/var\/log\/redis.log/" redis.conf > /etc/redis/redis.conf
####
# Redis correctly installed.
# Download script for running Redis
####
wget -q https://raw.github.com/saxenap/install-redis-amazon-linux-centos/master/redis-server
mv redis-server /etc/init.d
chmod 755 /etc/init.d/redis-server
chkconfig --add redis-server
chkconfig --level 345 redis-server on
####
# To start Redis just uncomment this line
####
#service redis-server start


