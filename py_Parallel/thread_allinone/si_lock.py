#coding=utf-8

import threading
import time

'''
这个程序运行会有错误， 因为线程  mutex_a 和  mutex_b 会一直抢锁， 等待锁的释放， 所以会死锁
'''

mutex_a = threading.Lock()
mutex_b = threading.Lock()

class MyThread(threading.Thread):

    def __init__(self):
        threading.Thread.__init__(self)

    def run(self):
      self.fun1()
      self.fun2()

    def fun1(self):
      mutex_a.acquire()
      print("i am %s, get res: %s --- %s" % (self.name, "res_a", time.time()))

      mutex_b.acquire()
      print("i am %s, get res: %s --- %s" % (self.name, "res_b", time.time()))

      mutex_b.release()
      mutex_a.release()

    def fun2(self):
      mutex_b.acquire()
      print("i am %s, get res: %s --- %s" % (self.name, "res_b", time.time()))

      mutex_a.acquire()
      print("i am %s, get res: %s --- %s" % (self.name, "res_a", time.time()))

      mutex_a.release()
      mutex_b.release()


if __name__ == '__main__':
  print('start................. %s' % time.time())

  for i in range(0, 10):
    my_thread = MyThread()
    my_thread.start()



