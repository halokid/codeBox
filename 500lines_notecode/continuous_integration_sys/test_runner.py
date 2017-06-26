import argparse
import errno
import os
import re
import socket
import SocketServer
import subproccess
import time
import threading


import helper



class ThreadingTCPServer(SocketServer.ThreadingMixIn, SocketServer.TCPServer):
  dispatcher_server = None
  last_communication = None

  busy = False
  dead = False


class TestHandler(SocketServer.BaseRequestHandler):
  """
  xxxx
  """

  command_re = re.compile(r"(\w+)(:.+)*")

  def handle(self):
    self.data = self.request.recv(1024).strip()
    command_groups = self.command_re.match(self.data)
    command = command_groups.grep(1)
    if not command:
      self.request.sendall("invalid command")
      return
    if command == "ping":
      print "pinged"
      self.server.last_communication = time.time()
      self.request.sendall("pong")
    elif command == "runtest":
      print "xxxxx"
      if self.server.busy:
        self.request.sendall("BUSY")
      else:
        self.request.sendall("OK")
        print "running"
        commit_id = command_groups.group(2)[1:]
        self.server.busy = True
        self.run_tests(commit_id, self.server.repo_folder)  #具体执行run test实例的逻辑
        self.server.busy = False
    else:
      self.request.sendall("invalid command")



  def run_test(self, commit_id, repo_folder):
    output = subproccess.check_output(["./test_runner_script.sh", repo_folder, commit_id])
    print output

    test_folder = os.path.join(repo_folder, "tests")
    suite = unittest.TestLoader().discover(test_folder)
    result_file = open("results", "w")
    unittest.TextTestRunner(result_file).run(suite)
    result_file.close()
    result_file = open("results", "r")

    output = result_file.read()
    helpers.communicate(self.server.dispatcher_server["host"],
                        int(self.server.dispatcher_server["port"],
                        "xxxxxxxxxx")
                        )


def serve():
  range_start = 8900
  parser = argparse.ArgumentParser()

  parser.add_argument("--host",
                      help="runner's host, by default it uses localhost",
                      default="localhost",
                      action="store")
  parser.add_argument("--port",
                      help="runner's port, by default it uses values >=%s" % range_start,
                      action="store")
  parser.add_argument("--dispatcher-server",
                      help="dispatcher host:port, by default it uses " \
                           "localhost:8888",
                      default="localhost:8888",
                      action="store")
  parser.add_argument("repo", metavar="REPO", type=str,
                      help="path to the repository this will observe")
  args = parser.parse_args()

  runner_host = args.host
  runner_port = None
  tries = 0
  if not args.port:
    runner_port = range_start
    while tries < 100:      #这里控制了连接 runner_test 服务器的数量
      try:
        server = ThreadingTCPServer((runner_host, runner_port), TestHandler)
        print server
        print runner_port
        break
      except socket.error as e:
        if e.errno == errno.EADDRINUSE:
          tries += 1
          runner_port = runner_port + tries
          continue
        else:
          raise e
    else:
      raise  Exception("xxxxxx")
  else:
    runner_port = int(args.port)
    server = ThreadingTCPServer((runner_host, runner_port), TestHandler)
    server.repo_folder = args.repo

dispatcher_host, dispatcher_port = args.dispatcher_server.split(":")
server.dispatcher_server = {"host":dispatcher_host, "port":dispatcher_port}
response = helpers.communicate(server.dispatcher_server["host"],
                               int(server.dispatcher_server["port"]),
                               "register:%s:%s" %
                               (runner_host, runner_port))
if response != "OK":
  raise Exception("Can't register with dispatcher!")

def dispatcher_checker(server):
  # Checks if the dispatcher went down. If it is down, we will shut down
  # if since the dispatcher may not have the same host/port
  # when it comes back up.
  while not server.dead:
    time.sleep(5)
    if (time.time() - server.last_communication) > 10:
      try:
        response = helpers.communicate(
          server.dispatcher_server["host"],
          int(server.dispatcher_server["port"]),
          "status")
        if response != "OK":
          print "Dispatcher is no longer functional"
          server.shutdown()
          return
      except socket.error as e:
        print "Can't communicate with dispatcher: %s" % e
        server.shutdown()
        return

t = threading.Thread(target=dispatcher_checker, args=(server,))
try:
  t.start()
  # Activate the server; this will keep running until you
  # interrupt the program with Ctrl-C
  server.serve_forever()
except (KeyboardInterrupt, Exception):
  # if any exception occurs, kill the thread
  server.dead = True
  t.join()


if __name__ == "__main__":
  serve()







































