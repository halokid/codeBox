#! /bin/bash


# 直接停止启动
#docker-compose stop webappxx
#docker-compose start webappxx


# 停止之后，用最新的image来启动
docker-compose stop webappxx
docker-compose up -d --build 


