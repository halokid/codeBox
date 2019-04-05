import gevent 

def foo():
  print("Running in foo")
  gevent.sleep(0)
  print("再次显性地切换上下文到foo")

def bar():
  print("显性地切换上下文到bar")
  gevent.sleep(0)
  print("程序暗地自动切换回上下文到bar")

gevent.joinall([
  gevent.spawn(foo),
  gevent.spawn(bar),
])
