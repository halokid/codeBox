#coding=utf-8

import time
import threading

#消费者
def consumer(cond):
  t = threading.current_thread()
  with cond:
    cond.wait() #创立一个锁，等待producer解锁
    print('{}: resource is availcble to consumer' . format(t.name))


def producer(cond):
  t = threading.current_thread()
  with cond:
    print('{}: making resource available' . format(t.name))
    cond.notifyAll()    #释放锁， 唤醒消费者

condition = threading.Condition()


c1 = threading.Thread(name='c1', target=consumer, args=(condition, ))
p = threading.Thread(name='p', target=producer, args=(condition, ))
c2 = threading.Thread(name='c2', target=consumer, args=(condition, ))


c1.start()
time.sleep(1)

c2.start()
time.sleep(1)

p.start()
