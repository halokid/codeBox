#coding=utf-8

'''
python 因为 GIL 的关系，不能实现多线程的并行， 目前只可以用多进程的并行

性能会比 多线程的并行要低一些， 因为要切换上下文空间和更多的内存占用

每一个进程都要有属于自己的上下文空间 和一些独占的内存（保存一些属于自己的数据等用途）

而线程却不用有上面的资源占用，所以多线程的并发性能更优 
'''

import time
from multiprocessing import Pool

import numpy as np


#omit code...

pCOs = np.linspace(1e-5, 0.5, 10)
pO2s = np.linspace(1e-5, 0,5, 10)

#并行一个关键的地方是把 要并行的逻辑代码段封装成一个函数
def task(pO2):
    "接受一个O2分压， 根据当前的CO分压进行动力学求解"
    #omit code...

if "__main__" == __name__:
  '''
  try:
    start = time.time()

    pool = Pool()      #创建进程池对象， 进程数与 multiprocessing.cpu_count() 相同
    tofs = pool.map(task, pCOs)   #并行执行函数

    end =  time.time()
    t = end - start
  finally:
    #收集计算的结果并进行处理绘图
    '''


