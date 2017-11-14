#coding=utf-8

import time
import threading
from random import randint

TIMEOUT = 2

def consumer(event, l):
  t = threading.current_thread()
  while l:
    event_is_set = event.wait(TIMEOUT)
    if event_is_set:
      try:
        integer = l.pop()
        print('{} popped from list by {}' . format(integer, t.name))
        event.clear()  #重置状态
      except IndexError:
        pass

def producer(event, l):
  t = threading.current_thread()
  while l:
    integer = randint(10, 100)
    l.append(integer)
    print('{} append to list by {}' . format(integer, t.name))
    event.set()
    time.sleep(1)

event = threading.Event()
l = []    #l is a list
threads = []

p = threading.Thread(name='producer1', target=producer, args=(event, l))
p.start()
threads.append(p)

for name in ('consumer1', 'consumer2'):
  t = threading.Thread(target=consumer, name=name, args=(event, l))
  t.start()
  threads.append(t)


for t in threads:
  t.join()
print('ending')















