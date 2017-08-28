#!/bin/bash

TOMCATDIR=/home/tomcat
TOMCATTAR=$TOMCATDIR/apache-tomcat-7.0.65.tar.gz
APACHETOMCAT=$TOMCATDIR/apache-tomcat-7.0.65
JAVADIR=/usr/java
JDKDIR=$JAVADIR/jdk1.8.0_65
JDKTAR=$JAVADIR/jdk-8u65-linux-x64.tar.gz
####8080####
T8080DIR=$TOMCATDIR/tomcat-rest
START8080=$T8080DIR/bin/startup.sh
####8081####
T8081DIR=$TOMCATDIR/tomcat-wap
CONF8081=$T8081DIR/conf/server.xml
CATA8081=$T8081DIR/bin/catalina.sh
START8081=$T8081DIR/bin/startup.sh
####8082####
T8082DIR=$TOMCATDIR/tomcat-activity
CONF8082=$T8082DIR/conf/server.xml
CATA8082=$T8082DIR/bin/catalina.sh
START8082=$T8082DIR/bin/startup.sh
. /etc/init.d/functions
panduan (){
  if [ $? -eq 0 ] 
  then
      action " $1 " /bin/true
    else
      action " $1 " /bin/false
    exit 1
  fi
}
CHECKNET(){
  if [ `netstat -lntup|grep java|grep -v grep|wc -l`  -ne 0 ]
  then 
    echo '########################'
    echo 'TOMCAT already existing'
    echo '########################'
    echo '===If you want to continue==='
    echo '===you can KILL tomcat first==='
    exit 1
  fi
}
CHECKDIR(){
  [ ! -d $JAVADIR ] && mkdir -p $JAVADIR
  rm -fr $APACHETOMCAT
  rm -fr $T8080DIR
}
CHECKTAR(){
  if [ ! -e $JDKTAR ]
  then
    echo "===>PLS DOWNLOAD 'JDKTAR' AT W8 D<===" 
    exit：
  fi
  if [ ! -e $TOMCATTAR ]
  then
    echo "===>PLS DOWNLOAD 'TOMCATTAR' AT 'W8 D'<===" 
    exit    
  fi 
  panduan Check_Tar
}
INSTALLTOM(){
  cd $JAVADIR
  tar xf $JDKTAR
  cd $TOMCATDIR
  tar xf $TOMCATTAR
  /bin/mv -f $APACHETOMCAT $T8080DIR
  panduan MV_TOMCAT
  /bin/cp -a $T8080DIR $T8081DIR
  /bin/cp -a $T8080DIR $T8082DIR
  panduan Install
}
PROFILE(){
  if [ `grep '===>TOMCAT<===bylc' /etc/profile|wc -l` -eq 0 ]
  then
cat >>/etc/profile<<AA
########===>TOMCAT<===bylc########
export JAVA_HOME=$JDKDIR
export CLASSPATH=.:\$JAVA_HOME/lib/dt.jar:\$JAVA_HOME/lib/tools.jar
export PATH=$JAVA_HOME/bin:\$PATH
export CATALINA_BASE=$T8080DIR
export CATALINA_HOME=$T8080DIR
export TOMCAT_HOME=$T8080DIR
export CATALINA_2_BASE=$T8081DIR
export CATALINA_2_HOME=$T8081DIR
export TOMCAT_2_HOME=$T8081DIR
export CATALINA_3_BASE=$T8082DIR
export CATALINA_3_HOME=$T8082DIR
export TOMCAT_3_HOME=$T8082DIR
##############################
AA
 fi
panduan profile
source /etc/profile
}
SERVERXML(){
#vi /home/tomcat/tomcat-rest/conf/server.xml
####8081####
  sed -i '22s#8005#8006#g' $CONF8081
  sed -i '71s#8080#8081#g' $CONF8081
  sed -i '93s#8009#8010#g' $CONF8081
####8082####
  sed -i '22s#8005#8007#g' $CONF8082
  sed -i '71s#8080#8082#g' $CONF8082
  sed -i '93s#8009#8011#g' $CONF8082
}
CATASH(){
#vi /home/tomcat/tomcat-rest/bin/catalina.sh
####8081####
  sed -i '2 iexport CATALINA_BASE=$CATALINA_2_BASE' $CATA8081
  sed -i '2 iexport CATALINA_HOME=$CATALINA_2_HOME' $CATA8081
####8082####
  sed -i '2 iexport CATALINA_BASE=$CATALINA_3_BASE' $CATA8082
  sed -i '2 iexport CATALINA_HOME=$CATALINA_3_HOME' $CATA8082
}
STARTALL(){
  #ps -ef|grep java|grep -v 'grep'|awk '{print $2}'|xargs kill 
  $START8080
  panduan start8080
  $START8081
  panduan start8081
  $START8082
  panduan start8082
  sleep 3
}
CHECKPROT(){
  if [ `netstat -lntup|grep java|grep -v grep|wc -l`  -ge 6 ]
  then 
    action 'PROT' /bin/true      
  else 
    action 'PROT' /bin/false
  fi
}
main (){
  CHECKNET
  CHECKDIR
  CHECKTAR
  INSTALLTOM
  PROFILE
  SERVERXML
  CATASH
  STARTALL
  CHECKPROT
}
main