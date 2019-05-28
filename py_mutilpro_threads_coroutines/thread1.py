def coroutine(func):
  def ret():
    f = func()
    f.next()
    return f
  return ret


@coroutine
def consumer():
  print "Wait to getting a task"
  while True:
    n = (yield)
    print "Got task %s" % n


import time
def producer():
  c = consumer()
  task_id = 0
  while True:
    time.sleep(1)
    print "Send a task to consumer %s" % task_id
    c.send("task %s" % task_id)

if __name__ == "__main__":
  producer()




