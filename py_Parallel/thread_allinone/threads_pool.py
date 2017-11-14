# coding=utf-8

import time
import threading
from random import random
import queue


def double(n):
  return n * 2


class Worker(threading.Thread):
  def __init__(self, queue):
    super(Worker, self).__init__()
    self._q = queue
    self.daemon = True
    self.start()

  def run(self):
    while 1:
      f, args, kwargs = self._q.get()
      try:
        print('USE: {}'.format(self.name))
        print(f(*args, **kwargs))
      except Exception as e:
        print(e)
      self._q.task_done()



class ThreadPool(object):
  def __init__(self, max_num = 5):
    self._q = queue.Queue(max_num)
    for _ in range(max_num):
      Worker(self._q)     #创建 worker 线程

  def add_task(self, f, *args, **kwargs):
    self._q.put(f, args, kwargs)

  def wait_complete(self):
    self._q.join()



pool = ThreadPool()
for _ in range(8):
  wt = random()
  pool.add_task(double, wt)
  time.sleep(wt)

pool.wait_complete()
















