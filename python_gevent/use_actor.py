import gevent
from gevent.queue import Queue 
from gevent import Greenlet 

class Pinger(Actor):
  def recv(self, msg):
    print(msg)
    pong.inbox.put("ping")
    gevent.sleep(0)

class Ponger(Actor):
  def recv(self, msg):
    print(msg)
    ping.inbox.put("pong")
    gevent.sleep(0)

ping = Pinger()
pong = Ponger()

ping.start()
pong.start()

ping.inbox.put("start")
gevent.joinall([ping, pong])

