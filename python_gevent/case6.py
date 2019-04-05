import gevent
from gevent.queue import Queue, Empty

tasks = Queue(maxsize = 3)

def worker(name):
  try:
    while True:
      task = tasks.get(timeout = 1)    # 
      print("worker %s 正在处理任务 %s" % (name, task))
      gevent.sleep(0)
  except Empty:
    print("队列处理完成， 退出")


def boss():
  for i in range(1, 10):
    tasks.put(i)
  print("已经分配完所有任务给 worker1, work1会迭代处理...")

  for i in range(10, 20):
    tasks.put(i)  
  print("已经分配完所有任务给 worker2, work1会迭代处理...")

gevent.joinall([
  gevent.spawn(boss),
  gevent.spawn(worker, "steven"),
  gevent.spawn(worker, "john"),
  gevent.spawn(worker, "bob"),
])
