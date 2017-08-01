#!/usr/bin/env python
#coding=utf-8

'''

wget -c -O /work/jboss.tar.gz --no-check-certificate  http://172.21.28.168/jboss.tar.gz
#cd /work
sleep 2
tar -zxf /work/jboss.tar.gz -C /work/ 

'''

import os

WORKDIR = "/work/"
DOWNLINK = "http://172.21.28.168/jboss.tar.gz"


def deply_jboss():
    downres = os.popen("wget -c -O " + WORKDIR + "jboss.tar.gz --no-check-certificate " + DOWNLINK).read()
    print downres
    tmp = os.popen('sleep 2')
    tarres = os.popen("tar -zxf " + WORKDIR  + "jboss.tar.gz -C " + WORKDIR).read()
    print tarres


if __name__ == '__main__':
    deply_jboss()



