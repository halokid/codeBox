#coding=utf-8
'''
当  MyPackage 的 __init__.py 文件不为空时
'''
import MyPackage

if __name__ == "__main__":
  c1 = MyPackage.ClassOne()
  c1.printInfo()

  c2 = MyPackage.ClassTwo()
  c2.printInfo()



