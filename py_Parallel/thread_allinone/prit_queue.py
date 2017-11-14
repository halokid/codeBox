#coding=utf-8

import time
import threading
from random import randint
import queue


q = queue.PriorityQueue()


def double(n):
  return n*2


def producer():
  count = 0
  while 1:
    if count > 2:
      break
    prit = randint(0, 100)
    print("put: {}" . format(prit))
    q.put(prit, double, prit)   #优先级， 函数， 参数
    count += 1


def consumer():
  while 1:
    if q.empty():
      break
    pri, task, arg = q.get()    #照输入 put 的时候，原样获取这些数值
    print('PRI: {} {} * 2 = {}' . format(pri, arg, task(arg)))
    q.task_done()
    time.sleep(0.1)   #给时间系统切换线程

t = threading.Thread(target=producer)
t.start()

time.sleep(1)

t = threading.Thread(target=consumer)
t.start()



















