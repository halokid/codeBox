#coding:utf-8

a = 3
b = 5

a, b = b, a+b

print a
print b


class Person():
  def getName(self):
    print "xx"


p = Person()
[p.getName() for i in range(4)]
