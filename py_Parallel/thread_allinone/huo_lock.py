#coding=utf-8

import time
import threading

mutex = threading.RLock()

class MyThread(threading.Thread):

  def __init__(self):
    threading.Thread.__init__(self)

  def run(self):
    self.fun1()
    self.fun2()

  def fun1(self):
    mutex.acquire()
    print("i am %s, get res: %s --- %s" % (self.name, "res_a", time.time()))

    mutex.acquire()
    print("i am %s, get res: %s --- %s" % (self.name, "res_b", time.time()))

    mutex.release()
    mutex.release()

  def fun2(self):
    mutex.acquire()
    print("i am %s, get res: %s --- %s" % (self.name, "res_b", time.time()))

    mutex.acquire()
    print("i am %s, get res: %s --- %s" % (self.name, "res_b", time.time()))

    mutex.release()
    mutex.release()

if __name__ == '__main__':
  print('start .......................... %s' % time.time())

  for i in range(0, 10):
    my_thread = MyThread()
    my_thread.start()


