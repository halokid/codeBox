import time
import numpy as np

pCOs = np.linspace(1e-5, 0.5, 10)
pO2s = np.linspace(1e-5, 0.5, 10)

if "__main__" == __name__:
  try:
    start = time.time()

    for i, pO2 in enumerate(pO2s):
      #...
      for j, pCO in enumerate(pCOs):
        #...
        #针对当前的分压值 pCO, pO2 进行动力学求解
        #具体的逻辑代码
    end = time.time()
    t = end - start
  finally:
    #收集计算的结果并进行处理绘图


      