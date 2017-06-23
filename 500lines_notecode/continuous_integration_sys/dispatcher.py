
import argparse
import os
import re
import socket
import SocketServer
import time
import threading

import helper


def dispatch_tests(server, commit_id):
  #FIXME: usually we dont run this forever
  while True:
    print "trying to dispatch to runners"
    for runner in server.runners:
      # notify the dispatch server, 通知调度器
      response = helpers.communicate(runner["host"],
                                     int(runner["port"]),
                                     "runtest:%s" % commit_id)
      if response == "OK":
        print "adding id %s" % commit_id
        # 更新已经通知了的调度器的列表
        server.dispatched_commits[commit_id] = runner
        if commit_id in server.pending_commits:
          # 既然已经通知了调度器，那么就在等待被通知的调度器列表里面删除已经通知了的调度器元素
          server.pending_commits.remove(commit_id)
        return

    time.sleep(2)


# 此class作用就相当于定义一个结构体
class ThreadingTCPServer(SocketServer.ThreadingMixIn, SocketServer.TCPServer):
  runners = []
  dead = False
  dispatched_commits = {}
  pending_commits = []


class DispatcherHandler(SocketServer.BaseRequestHandler):
  """
  doc here
  """
  command_re = re.compile(r"(\w+)(:.+)*")
  BUF_SIZE = 1024

  def handle(self):
    self.data = self.request.recv(self.BUF_SIZE).strip()
    command_groups = self.command_re.match(self.data)

    if not command_groups:
      self.request.sendall("invalid command")
      return

    command = command_groups.group(1)

    if command == "status":
      print "in status"
      self.request.sendall("OK")
    elif command == "register":     # 注册这一台调度器
      print "register"
      address = command_groups.group(2)
      host, port = re.findall(r":(\w*)", address)
      runner = {"host": host, "port": port}
      self.server.runners.append(runner)    # 一个储存map的slice
      self.request.sendall("OK")
    elif command == "dispatch":       # 调度这台调度器
      print "going to dispatch"
      commit_id = command_groups(2)[1:]
      if not self.server.runners:
        self.request.sendall("no runners are registered")
      else:
        self.request.sendall("OK")
        dispatch_tests(self.server, commit_id)
    elif command == "results":
      print "got test results"
      results = command_groups.group(2)[1:]
      results = results.split(":")
      commit_id = results[0]
      length_msg = int(results[1])
      remaining_buffer = self.BUF_SIZE - (len(command) + len(commit_id) + len(results[1]) + 3)
      if length_msg > remaining_buffer:
        self.data += self.request.recv(length_msg - remaining_buffer).strip()
      del self.server.dispatched_commits[commit_id]
      if not os.path.exists("test_results"):
        os.makedirs("test_results")
      with open("test_results/%s" % commit_id, "w") as f:
        data = self.data.split(":")[3:]
        data = "\n".join(data)
        f.write(data)
      self.request.sendall("OK")
    else:
      self.request.sendall("invalid command")


  def serve():
    parser = argparse.ArgumentParser()
    parser.add_argument("--host",
                        help = "dispatcher's host, by default it uses localhost",
                        default = "localhost",
                        action = "store"
                        )
    parser.add_argument("--port",
                        help = "dispatcher's port, by default is uses 8888",
                        default = 8888,
                        action = "store"
                        )

    args = parser.parse_args()

    # create the server
    server = ThreadingTCPServer((args.host, int(args.port)), DispatcherHandler)
    print "serving on %s:%s" % (args.host, int(args.port))
    # create a thread to check teh runner pool
    def runner_checker(server):
      def manage_commit_lists(runner):
        for commit, assigned_runner in server.dispatched_commits.iteritems():
          if assigned_runner == runner:
            del server.dispatched_commits[commit]
            server.pending_commits.append(commit)
            break
        server.runners.remove(runner)

      while not server.dead:
        time.sleep(1)
        for runner in server.runners:
          s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
          try:
            response = helper.communicate(runner["host"],
                                          int(runner["port"]),
                                          "ping")
            if response != "pong":
              print "removing runner %s" % runner
              manage_commit_lists(runner)
          except socket.error as e:
            manage_commit_lists(runner)

    # this will kick off tests that failed
    def redistribute(server):
      while not server.dead:
        for commit in server.pending_commits:
          print "running redistribute"
          print server.pending_commits
          dispatch_tests(server, commit)
          time.sleep(5)


    # multi thread process code model
    runner_hearbeat = threading.Thread(target = runner_checker,args = (server,))
    redistributor = threading.Thread(target = redistribute, args=(server,))

    try:
      runner_hearbeat.start()
      redistributor.start()
      # active the server; this will kepp running until you
      # intrrupt the program with CTRL+C or Cmd+C
      server.serve_forever()
    except (KeyboardInterrupt, Exception):
      # if any exception occurs, kill the thread
      server.dead = True
      runner_hearbeat.join()
      redistributor.join()



if __name__ == '__main__':
  serve()



































