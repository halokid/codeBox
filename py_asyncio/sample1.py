import asyncio

async def count():
  print("One")
  await asyncio.sleep(1)
  print("Two")

async def main():
  await asyncio.gather(count(), count(), count())

if __name__ == "__main__":
  import time
  s = time.perf_counter()
  # asyncio.run(main())           # py 3.7之后才支持

  loop = asyncio.get_event_loop()
  f = [count(), count(), count()]
  loop.run_until_complete(asyncio.wait(f))
  loop.close()

  elapsed = time.perf_counter() - s
  print(f"{__file__} executed in {elapsed:0.2f} seconds.")


