import gevent 
import time 
from gevent import select 

start = time.time()
tic = lambda:  'at %1.1f seconds' % (time.time() - start)


def gr1():
  # 程序在这里处理需要点时间， 但是我们并不希望程序停在这里
  print("gr1 开始轮询: %s" % tic())
  select.select([], [], [], 2)
  print('gr1 结束轮询: %s' % tic())

def gr2():
  # 程序在这里处理需要点时间， 但是我们并不希望程序停在这里
  print('gr2 开始轮询: %s' % tic())
  select.select([], [], [], 2)
  print('gr2 结束轮询: %s' % tic())

def gr3():
  print('轮询需要2秒，这个时候gr3 可以做一些其他的逻辑处理， 当程序在轮询的时候, 等你妹啊, %s' % tic())
  gevent.sleep(1)


gevent.joinall([
  gevent.spawn(gr1),
  gevent.spawn(gr2),
  gevent.spawn(gr3),
])


