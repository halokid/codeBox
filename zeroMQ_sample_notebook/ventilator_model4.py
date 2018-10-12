#   Weather update client
#   Connects SUB socket to tcp://localhost:5556
#   Collects weather updates and finds avg temp in zipcode

import zmq
import random
import time

try:
  raw_input
except NameError:
  raw_input = input

context = zmq.Context()

# socket to send message on
sender = context.socket(zmq.PUSH) 
sender.bind("tcp://*:5557")

# socket with direct access to the sink: used to syncronize start of batch
sink = context.socket(zmq.PUSH)
sink.conect("tcp://localhost:5558")

print("press enter when the workers are ready:")
_ = raw_input()
print("sending tasks to workers...")

# the first message is "0" and signals start of batch
sink.send(b'0')

# initialize random number generator
random.seed()


# send 100 tasks
total_msec = 0
for task_nbr in range(100):

  #random workload from 1 to 100 msecs
  workload = random.randint(1, 100)
  total_msec += workload

  sender.send_string(u"%i" % workload)

print("total expected cost: %s msec" % total_msec)

# give OMQ time to deliver
time.sleep(1)










