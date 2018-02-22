#coding=utf-8
'''
当  MyPackage 的 __init__.py 文件不为空时
'''
from MyPackage import *

if __name__ == "__main__":
  c1 = ClassOne()
  c1.printInfo()

  c2 = ClassTwo()
  c2.printInfo()



