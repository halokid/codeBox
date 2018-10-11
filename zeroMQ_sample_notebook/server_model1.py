import time
import zmq


context = zmq.Conext()
socket = context.socket(zmp.REP)
socket.bind("tcp://*:5555")


while True:
  # wait for next request from client
  message = socket.recv()
  print("received request:  %s" % message)

  # do some work
  time.sleep(1) 

  # send reply back to client
  socket.send(b"world")



