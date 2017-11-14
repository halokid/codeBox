#coding=utf-8

import time
import threading
import random
import queue

q = queue.Queue()

def double(n):
  return n*2

def producer():
  while 1:
    wt = random.randint(1, 10)
    time.sleep(random.random())
    q.put(double, wt)

def consumer():
  while 1:
    task, arg = q.get()
    print(arg, task(arg))

    q.task_done()

for target in (producer, consumer):
  t = threading.Thread(target=target)
  t.start()






