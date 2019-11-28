#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Desc  : 动态添加协程任务


import asyncio
from threading import Thread


async def create_task(event_loop):
    i = 0
    while True:
        # 每秒产生一个任务, 提交到线程里的循环中, event_loop作为参数

        '''
        假如有多个数据源的话， 比如是一个列表  [1, 2, 3] ， 在这里应该循环加入，循环添加以  每个元素为原子单位的 任务， 这样的做法应该比较好，粒度比较少， 好调试
        '''
        asyncio.run_coroutine_threadsafe(production(i), event_loop)
        asyncio.run_coroutine_threadsafe(production(i + 10), event_loop)
        await asyncio.sleep(1)
        i += 1


async def production(i):
    while True:
        print("第{}个coroutine任务".format(i))
        await asyncio.sleep(1)


def start_loop(loop):
    #  运行事件循环， loop作为参数
    asyncio.set_event_loop(loop)
    loop.run_forever()


thread_loop = asyncio.new_event_loop()  # 创建事件循环
run_loop_thread = Thread(target=start_loop, args=(thread_loop,))  # 新起线程运行事件循环, 防止阻塞主线程
run_loop_thread.start()  # 运行线程，即运行协程事件循环

main_loop = asyncio.new_event_loop()
main_loop.run_until_complete(create_task(thread_loop))  # 主线程负责create coroutine object