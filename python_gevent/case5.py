import gevent
from gevent.event import Event 


evt = Event()

def setter():
  print("A: 请等待我，我正在做一些事情...")
  gevent.sleep(3)
  print("OK, 我已经完成了")
  evt.set()

def waiter():
  print("我在等待你")
  evt.wait()
  print("已经等待完成。。。")

def main():
  gevent.joinall([
    gevent.spawn(setter),
    gevent.spawn(waiter),
    gevent.spawn(waiter),
    gevent.spawn(waiter),
    gevent.spawn(waiter),
    gevent.spawn(waiter),
  ])


if __name__ == "__main__": main()


