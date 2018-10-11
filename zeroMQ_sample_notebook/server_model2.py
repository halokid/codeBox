import zmq
from random import randrange


context = zmq.Context()
socket = context.socket(zmq.PUB)
socket.bind("tcp://*:5556")

# 这个更新流没有开始和结束，像一个永远不会终止的广播。
while True:
  zipcode = randrange(1, 100000)
  temperature = randrange(-80, 135)  
  relhumidity = randrange(10, 60)

  socket.send_string("%i %i %i" % (zipcode, temperature, relhumidity))




