#coding:utf-8

import logging, threading

from queue import Queue

logger = logging.getLogger()
logger.setLevel(logging.DEBUG)
formatter = logging.Formatter('%(asctime)s - %(message)s')

ch = logging.StreamHandler()
ch.setLevel(logging.DEBUG)
ch.setFormatter(formatter)
logger.addHandler(ch)


fibo_dict = {}
shared_queue = Queue()
input_list = [3, 10, 5, 7]

queue_condition = threading.Condition()


lock = threading.Lock()

def fibonacci_task(condition):
  with condition:
    while shared_queue.empty():
      lock.acquire()
      logger.info("[%s] - waitting for elements in queue ..." % threading.current_thread().name)
      condition.wait()
      lock.release()
    else:
      value = shared_queue.get()
      a, b = 0, 1

      for item in range(value):
        a, b = b, a+b
        fibo_dict[value] = a

      shared_queue.task_done()
      logger.debug("[%s] fibonzcci of key [%d] with result [%d]" %
                   (threading.current_thread().name, value, fibo_dict[value]))



def queue_task(condition):
  logging.debug('starting queue_task ...')
  while condition:
    for item in input_list:
      shared_queue.put(item)    #put data into share_queue
    logging.debug("notifying fibonacci_task threads that the queue is ready to consume...")
    lock.acquire()
    condition.notifyAll()
    lock.release()


#开四条线程来处理斐波拉计算
threads = [threading.Thread(daemon=True, target=fibonacci_task, args=(queue_condition, )) for i in range(4)]  #create 4 threads for process

[thread.start() for thread in threads]


#这里是单线程执行队列任务的，主要是把要计算斐波拉的数字的列表传进去任务队列里面
prod = threading.Thread(name='queue_task_thread', daemon=True, target=queue_task, args=(queue_condition, ))
prod.start()

[thread.join() for thread in threads]

# prod.join()

logger.info("[%s] - Result: %s" % (threading.current_thread().name, fibo_dict))













