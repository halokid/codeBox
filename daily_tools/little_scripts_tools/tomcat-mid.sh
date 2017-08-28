#!/bin/bash
#tomcat environment java1.7.0_67  
#get tomcat

WORKDIR=/work
TOMCATDIR=/work/apache-tomcat-7.0.73

cd WORKDIR
wget http://apache.fayea.com/tomcat/tomcat-7/v7.0.73/bin/apache-tomcat-7.0.73.tar.gz

[ -d TOMCATDIR ] || tar -zxvf apache-tomcat-7.0.73.tar.gz
cd TOMCATDIR/bin
./startup.sh

