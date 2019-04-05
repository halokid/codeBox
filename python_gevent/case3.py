import gevent
import random

def task(pid):
  """
  处理结果不确定的任务
  """
  gevent.sleep(random.randint(0, 2) * 0.001)
  print("任务 %s done" % pid)


def syncDo():
  for i in range(1, 10):
    task(i)

def asyncDo():
  threads = [gevent.spawn(task, i) for i in range(10)] 
  gevent.joinall(threads)

print("同步处理...")
syncDo()

print("异步处理...")
asyncDo()

