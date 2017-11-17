#coding=utf-8

import random, time, Queue
from multiprocessing.managers import BaseManager

#queue for send data task
task_queue = Queue.Queue()

#queue for receive result task
result_queue = Queue.Queue()

#QueueManager class extends from BaseManager
class QueueManager(BaseManager):
  pass


#register two queue in the network, callable args link to Queue object
QueueManager.register('get_task_queue', callable=lambda : task_queue)
QueueManager.register('get_result_queue', callable=lambda : result_queue)


manager = QueueManager(address=('', 5000), authkey='abc')
manager.start()

task = manager.get_task_queue()
result = manager.get_result_queue()



for i in range(10):
  n = random.randint(0, 10000)
  print('put task %d...' % n)
  task.put(n)

print('try get results...')

for i in range(10):
  r = result.get(timeout=10)
  print('result: %s' % r)

manager.shutdown()











