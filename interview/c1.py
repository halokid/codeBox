

class D(object):
  def __add__(self, other):
    return 1


a = D()
b = D()
print(a + b)

a = 2**3
print(a)

l = [1, 2]
x = map(lambda x: 2**x, l)
print(list(x))

a = list()
a.append(1)
a.append('x')

a = [1, 2]
b = a
# b[1] = 3
a.insert(1, 5)
print(a)


a = "HELLO WOE"
print(a.capitalize())

a = [2, 5, 6]
# a.sort()
# print(a)

b = [3, 4]
c = a.append(b)
print(c)

z = set('abc')
z.add('san')
z.update(set(['p', 'q']))
print(z)

print([i.lower() for i in "TURING"])