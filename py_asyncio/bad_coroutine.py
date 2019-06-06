import asyncio
import os
import urllib.request
import time

import ssl
ssl._create_default_https_context = ssl._create_unverified_context

'''
这个协程程序根本就没有达到协程的效果，为什么呢？
因为仅仅只是 调用 downloadCorontine 函数本身用了协程调用，但是使用urlib去download文件的时候，并没有真正协程执行
因为urlib不是异步的， 所以这是一个失败的协程程序
'''

async def downloadCorontine(url):
  print("start download time ----- ", time.time())
  req = urllib.request.urlopen(url)
  fileName = os.path.basename(url)

  with open(fileName, "wb") as fileHlder:
    while True:
      chunk = req.read(1024)
      if not chunk:
        break
      fileHlder.write(chunk)

  msg = "Finished downloading " + fileName
  print("finished download time ----- ", time.time(), "\n\n")
  return msg

async def main(urls):
  coroutines = [downloadCorontine(url) for url in urls]
  completed, pending = await asyncio.wait(coroutines)
  for item in completed:
    print(item.result())

if __name__ == "__main__":
  urls = ["http://www.irs.gov/pub/irs-pdf/f1040.pdf",
          "http://www.irs.gov/pub/irs-pdf/f1040a.pdf",
          "http://www.irs.gov/pub/irs-pdf/f1040ez.pdf",
          "http://www.irs.gov/pub/irs-pdf/f1040es.pdf",
          "http://www.irs.gov/pub/irs-pdf/f1040sb.pdf"]
  eventLoop = asyncio.get_event_loop()
  try:
    eventLoop.run_until_complete(main(urls))
  finally:
    eventLoop.close()


