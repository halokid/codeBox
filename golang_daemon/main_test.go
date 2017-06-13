package main

import (
  "testing"
)


var runDaemon RunDaemon

func Test_runMysql(t *testing.T) {
  runDaemon.runMysql()
}

func Test_runRedis(t *testing.T) {
  runDaemon.runRedis()
}

func Test_runWeb(t *testing.T) {
  runDaemon.runWeb()
}


func Test_runWeb1(t *testing.T) {
  runDaemon.runWeb1()
}

func Test_runWeb2(t *testing.T) {
  runDaemon.runWeb2()
}







