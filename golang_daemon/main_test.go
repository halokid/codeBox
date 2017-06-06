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



