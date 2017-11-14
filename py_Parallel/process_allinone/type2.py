#coding=utf-8
import random
import time
from multiprocessing import Process



class Piao(Process):
  def __init__(self,name):
    super().__init__()
    self.name = name

  def run(self):
    print('%s piaoing' %self.name)
    time.sleep(random.randrange(1,5))
    print('%s piao end' %self.name)

if __name__ == '__main__':  # 进程开启必须放在main()下
  for i in ('aaa', 'bbb', 'ccc'):
    p = Process(name=i, target=Piao, args=(i,))
    p.start()
  print('主进程')


