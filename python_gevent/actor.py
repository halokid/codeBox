import gevent
from gevent.queue import Queue 

class Actor(gevent.Greenlet):

  def __init__(self):
    self.inbox = Queue()
    Greenlet.__init__(self)

  def recv(self, msg):
    raise NotImplemented()

  def _run(self):
    self.running = True

    while self.running:
      msg = self.inbox.get()
      self.recv(msg)


      