
import argparse
import os
import re
import socket
import SockServer
import subprocess
import sys
import time

import helper


def poll():
  parser = argparse.ArgumentParser()
  parser.add_argument("--dispatcher-server", help="dispatcher host:port, by default it uses localhost:8888", action="store")
  parser.add_argument("repo", metavar="REPO", type=str, help="path to the repository this will observer")

  args = parser.parse_args()
  dispatcher_host, dispatcher_port  = args.dispatcher_server.split(":")

  while true:
    try:
      subprocess.check_output(["./update_repo.sh", args.repo])
    except subprocess.CallProcessError as e:
      raise Exception("could not update and check repository. reason: %s" % e.output)


    if os.path.isfile(".commit_id"):
      try:
        response = helper.communicate(dispatcher_host,
                                      int(dispatcher_port), "status")     #先检查看看 disaptcher 服务器存活不存活
      except socket.error as e:
        raise Exception("could not communicate with dispatcher server: %s " % e )

      if response == "OK":
        commit = ""
        with open(".commit_id", "r") as f:
          commit = f.readline()

        response = helper.communicate(dispatcher_host,
                                      int(dispatcher_port),
                                      "dispatch: %s" % commit)    #假如disaptcher是存活的，那么则往其发送commit_id

        if response !=  "OK":
          raise Exception("could not dispatch the test: %s" % response)

        print "dispatched!"
      else:
        raise Exception("could not dispatch the test: %s" % response)

    time.sleep(5)



if __name__ == "__main__":
  poll()



