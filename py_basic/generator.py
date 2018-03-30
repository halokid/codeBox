#coding=utf-8


def generator_fun():
  for i in range(10):
    yield i
    
    
for item in generator_fun():
  print(item)
  
 
 
my_str = "yasoob"
my_iter = iter(my_str)
next(my_iter)



