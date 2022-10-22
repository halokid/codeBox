
def minpositive(a):
  count = 0
  b = list(set([i for i in a if i>0]))
  if min(b, default = 0)  > 1 or  min(b, default = 0)  ==  0 :
    min_val = 1
  else:
    min_val = min([b[i-1]+1 for i, x in enumerate(b) if x - b[i - 1] >1], default=b[-1]+1)

  return min_val