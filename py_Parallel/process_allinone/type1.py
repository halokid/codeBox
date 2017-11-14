# -*- coding: UTF-8 -*-

import time
import random
from multiprocessing import Process

def piao(name):
  print('%s piaoing' % name)
  time.sleep(random.randrange(1, 5))
  print('%s piao end' % name)


if __name__ == '__main__':
  for i in ('aaa', 'bbb', 'ccc'):
    p = Process(name=i, target=piao, args=(i, ))
    p.start()

  print('主进程')


