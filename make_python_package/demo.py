#coding=utf-8
'''
当  MyPackage 的 __init__.py 文件为空时
'''
from MyPackage.classOne import ClassOne
from MyPackage.classTwo import ClassTwo

if __name__ == "__main__":
  c1 = ClassOne()
  c1.printInfo()

  c2 = ClassTwo()
  c2.printInfo()

