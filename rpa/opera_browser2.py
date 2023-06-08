# coding = utf-8
import time
from selenium import webdriver


def refresh():
  # driver = webdriver.Chrome()
  driver = webdriver.Edge()
  url = "http://baidu.com"  # 网页地址
  driver.get(url)
  try:
    for i in range(10):
      time.sleep(1)
      driver.refresh()  # 调用webdriver中刷新页面的方法
      print(i)  # 记录刷新次数
  except Exception as e:
    print("Exception found", format(e))
  driver.close()


if __name__ == "__main__":
  refresh()

