#!/usr/bin/env python
#!coding=utf-8


'''
���ܽ���:
����jenkins�Զ�����֮�󣬰ѹ����ɹ�����k8s�ϵ��߼����̣���jenkins��job������ִ�е�shell�е���

ʹ�÷�����
./jenkins_k8s_deploy.py  $app_name $image_name  $deploy_name

1. $app_name ������app������
2. $image_name ���ɵ�docker image ������
3. $deploy_name ������ k8s ��deploy ������


��ŵ�shell��������:
------------------------------------------------------

cd ./xxshop/srv/auth

#�������
go build -o xxshop-srv-auth

#��������,ÿһ���µĹ�����Ҫ����һ���µ�image
docker build -t 10.86.20.57:5000/micro-xxshop-srv-auth:v$BUILD_NUMBER .

#push image
docker push 10.86.20.57:5000/micro-xxshop-srv-auth:v$BUILD_NUMBER

#���µ�image����k8s��deploy
/root/local/bin/kubectl set image deploy/micro-xxshop-srv-auth micro-xxshop-srv-auth=10.86.20.57:5000/micro-xxshop-srv-auth:v$BUILD_NUMBER -s http://10.86.20.57:8080

-------------------------------------------------------

'''


import sys, os

APP_NAME      =   sys.argv[1]
IMAGE_NAME    =   sys.argv[2]
DEPLOY_NAME   =   sys.argv[3]


GOPATH        =   "/usr/local/go/bin/go"
KUBECTL       =   "/root/local/bin/kubectl"

























