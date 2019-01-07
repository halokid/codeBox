#coding=utf-8


class Person():
  name = None
  
  def getName(self):
    return self.name


class SetPerson(Person):
  # super 指定的属性是不可以改变的， 这里会抛异常
  super.name = "r0x"
  def printStr(self):
    print("printStr .....")



if __name__ == "__main__":     
  print("test super")
  p = Person()
  print(p.getName())




