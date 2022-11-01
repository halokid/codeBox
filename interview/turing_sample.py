
class Developer(object):
  def __init__(self, skills):
    print('-->>> __init__')
    self.skills = skills

  def __add__(self, other):
    print('-->>> __add__')
    skills = self.skills + other.skills
    return Developer(skills)

  def __str__(self):
    print('-->>> __str__')
    return 'Skills'


if __name__ == '__main__':
  A = Developer('NodeJS')
  B = Developer('Python')
  print(A + B)

  # --------------------------------------
  try:
    print('Hello')
  except:
    print('An exception occured')
  finally:
    print('World')


  # -------------------------------------
  x = 'abcdef'
  i = 'a'
  print(x[:-1])
  # while i in x[:-1]:
  #   print(i, end=' ')


  # ---------------------------------
  f = None
  for i in range(5):
    with open('app.log', 'w') as f:
      if i > 2:
        break

  print(f.close())




