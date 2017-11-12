#coding=utf-8

import time
import threading


value = 0

'''
下面的程序的 对value 的更改并不加上锁， 所以很多线程并发再更改value的值， 100个线程跑下来， 就看哪个在更改了
大家都改来改去， 所以value的值不固定。。。可能有20个线程，同时把 value 改成 1
'''
def getlock():
    global  value
    new = value + 1
    time.sleep(0.001)
    value = new

for i in range(100):
    t = threading.Thread(target=getlock)
    t.start()

main_thread = threading.current_thread()

for t in threading.enumerate():
    if t == main_thread:
        continue
    t.join()

print(value)



