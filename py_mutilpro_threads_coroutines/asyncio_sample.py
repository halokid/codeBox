'''
适合 py 3.6.x
'''

import asyncio
import time
start = time.time()
s = time.perf_counter()

async def do(x):
  print("Waitting: ", x)
  print("Runtime is: ", time.time())
  await asyncio.sleep(x)
  return "Finish after {}s".format(x)


task1 = do(1)
task2 = do(2)
task3 = do(4)

tasks = [
  asyncio.ensure_future(task1),
  asyncio.ensure_future(task2),
  asyncio.ensure_future(task3),
]

loop = asyncio.get_event_loop()
loop.run_until_complete(asyncio.wait(tasks))


for task in tasks:
  print("Task result: ", task.result())

end = time.time()
print("TIME: ", end - start)

# 另外一种计算耗时的方式
elapsed = time.perf_counter() - s
print(f"{__file__} executed in {elapsed:0.2f} seconds.")



