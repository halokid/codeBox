#coding=utf-8

def foo(*params):
  print params
  for p in params:
    print p



if __name__ == "__main__":
  foo('a', 'b', 'c')


