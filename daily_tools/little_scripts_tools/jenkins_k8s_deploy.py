#!/usr/bin/env python
#!coding=utf-8


'''
功能介绍:
用于jenkins自动构建之后，把构建成果部署到k8s上的逻辑流程，在jenkins的job构建后执行的shell中调用

使用方法：
./jenkins_k8s_deploy.py  $app_name $image_name  $deploy_name

1. $app_name 构建的app的名称
2. $image_name 生成的docker image 的名称
3. $deploy_name 部署在 k8s 上deploy 的名称


大概的shell流程如下:
------------------------------------------------------

cd ./xxshop/srv/auth

#编译程序
go build -o xxshop-srv-auth

#生成容器,每一次新的构建都要生成一个新的image
docker build -t 10.86.20.57:5000/micro-xxshop-srv-auth:v$BUILD_NUMBER .

#push image
docker push 10.86.20.57:5000/micro-xxshop-srv-auth:v$BUILD_NUMBER

#用新的image更新k8s的deploy
/root/local/bin/kubectl set image deploy/micro-xxshop-srv-auth micro-xxshop-srv-auth=10.86.20.57:5000/micro-xxshop-srv-auth:v$BUILD_NUMBER -s http://10.86.20.57:8080

-------------------------------------------------------

'''


import sys, os

APP_NAME      =   sys.argv[1]
IMAGE_NAME    =   sys.argv[2]
DEPLOY_NAME   =   sys.argv[3]


GOBIN         =   "/usr/local/go/bin/go"
KUBECTL       =   "/root/local/bin/kubectl"
# DEPLOY_FILES   =   "./k8s_deploy/" + APP_NAME + "-srv"
K8S_MASTER    =   "http://10.86.20.57:8080"


class Jkd(object):
  def __init__(self, appName, imageName, deployName):
    self.appName = appName
    self.imageName = imageName
    self.deployName = deployName
    
    
  def buildApp(self):
    """
    编译go程序
    """
    r = os.popen(GOBIN + "build -o" + self.appName).read()
    print "go build result:   " + r
    
  
  
  def makeImage(self):
    """
    生成新的image
    """
    r = os.popen("docker build -t " + self.imageName + " .").read()
    print "make new image result: " + r
      
      
  def pushImage(self):
    """
    push image
    """
    r = os.popen("docker push " + self.imageName).read()
    print "push new image result: " + r
  
  
  def setDeploy(self):
    """
    deploy to k8s
    作为一个 go-micro 的服务来说， 要部署两个deploy
    
    deploy 1：  srv本身的image
    deploy 2：  srv的 api image
    
    其中两个deploy 都是注册到 consul 的
    """
    r = os.popen(KUBECTL + " get deploy -s " + K8S_MASTER + " grep " + self.appName + " | awk '{print $1}'").read()
    if r == "":
      print "deploy不存在， 创建中............"
      r = os.popen( KUBECTL + " create -f " + self.appName + ".yaml " + "-s" + K8S_MASTER).read()
    else:
      print "deploy存在， scale deploy的image.............."
      r = os.popen( KUBECTL + " set image deploy/" + self.appName + " " + self.appName + "=" + self.imageName + "-s" + K8S_MASTER).read()
























