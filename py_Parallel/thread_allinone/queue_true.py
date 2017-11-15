#coding=utf-8

from multiprocessing import Process, Queue
import os, time, random

#写数据进程执行的代码
def write(q):
  for value in ['A', 'B', 'C']:
    print('put %s to queue....' % value)
    q.put(value)
    time.sleep(random.random())


#读数据进程执行的代码
def read(q):
  while True:
    value = q.get(True)
    print('get %s from queue....' % value)


if __name__ == '__main__':
  q = Queue()
  pw = Process(target=write, args=(q, ))

  pr = Process(target=read, args=(q, ))

  pw.start()

  pw.join()


