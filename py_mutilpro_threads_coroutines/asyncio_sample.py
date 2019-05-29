import asyncio
import time
start = time.time()

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



