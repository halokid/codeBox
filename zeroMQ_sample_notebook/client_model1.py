import zmq

context = zmq.Context()

# socket to talk to server
print("connectin to hello world server ....")
socket = context.socket(zmq.REQ)
socket.connect("tcp://localhost:5555")

# do 10 requests, waitting 
for request in range(10):
  print("sending request %s ------" % request)
  socket.send(b"hello")

  # get the reply
  message = socket.recv()
  print("received reply %s [%s]" % (request, message))

  