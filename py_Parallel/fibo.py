#coding:utf-8

import logging, threading
from queue import Queue 

logger = logging.getLogger()
logger.setLevel(logging.DEBUG)

formatter = logging.Formatter('%(asctime)s - %(message)s')
ch = logging.StreamHandler()
ch.setFormatter(formatter)
logger.addHandler(ch)

fibo_dict = {}
shared_queue = Queue()
input_list = [3, 10, 5, 7]


def fibonacci_task(condition):
  with condition:
    while shared_queue.empty():
      logger.info("[%s] - waiting for elements in queue ..." % threading.current_thread().name)
      condition.wait()
    else:
      value = shared_queue.get()
      a, b = 0, 1
      
      for item in range(value):
        a, b = b, a+b   #斐波拉算法
        fibo_dict[item] = a
    
    shared_queue.task_done()
    logger.debug("[%s] fibonacci of key [%d] with result [%d]" % (threading.current_thread().name, value, fibo_dict[value]) 


def queue_task(condition):
  logging.debug('starting queue_task ...')
  with condition:
    for item in input_list








