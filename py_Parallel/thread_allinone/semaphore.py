#create by r00x {$DATE}
#coding=utf-8

import time
from random import random
from threading import Thread, Semaphore, current_thread, enumerate

sema = Semaphore(3)

def foo(tid):
    #以sema变量作为锁的标准，来锁住整个多线程流程
    with sema:
        print('{} acquire sema' . format(tid))
        wt  = random() * 2
        time.sleep(wt)

    print('{} release sema' . format(tid))



#loop, create thread for function foo
for i in range(5):
    t = Thread(target=foo,  args=(i, ))
    t.start()


main_thread = current_thread()
for t in enumerate():
    if t is main_thread:
        continue
    t.join()
    #join的意思是 主进程（其实就是第一个线程），一直在等待所有的线程结束,让所有的线程都join进去， 除了主线程不用join







