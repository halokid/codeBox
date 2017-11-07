# coding:utf-8

import time
import threading


def text(name):
  # 装饰器这玩意确实有点绕， 传进来这个name你当然可以当是一个字符串来用，
  # 其实这里就是当是一个字符串来用的, 但是假如当name当作一个字符串来用的时候
  # 应该如何去理解下面的 func 这个参数呢？？？？
  # 看下面这一段
  '''

  def makebold(fn):
    def wrapped():
        return "<b>" + fn() + "</b>"
    return wrapped

  def makeitalic(fn):
    def wrapped():
        return "<i>" + fn() + "</i>"
    return wrapped

  @makebold
  @makeitalic
  def hello():
    return "hello world"

  print hello()

  执行结果
  <b><i>hello world</i></b>

'''

  #这里的 func 是怎么来的？？？
  #其实这里就是闭包的原理,这里是直接继承了上一层函数的func这个传递， 而这个传递其实就是被装饰的函数
  #这是装饰器的特别， 自动传递函数， 自动传递函数的参数
  def profile(func):
    def wrapper(*args, **kwargs):
      start = time.time()
      res = func(*args, **kwargs)
      end = time.time()
      print('{} cost: {}'.format(name, end - start))
      return res

    return wrapper

  return profile  # 返回profile本身，即是返回函数自身


'''
def text2(name):
  def wrapper2(*args, **kwargs):
    start = time.time()
    res = name(*args, **kwargs)
    end = time.time()
    print(end - start)
    return res
  return wrapper2
'''


def fib(n):
  if n <= 2:
    return 1
  return fib(n - 1) + fib(n - 2)


@text('nothread')
def nothread():
  fib(20)
  fib(20)

@text('hasthread')
def hasthread():
  for i in range(2):
    t = threading.Thread(target=fib, args=(20, ))
    t.start()
  main_thread = threading.current_thread()
  for t in threading.enumerate():
    if t is main_thread:
      continue
    t.join()

nothread()
hasthread()




