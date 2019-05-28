from multiprocessing import Process

def loop():
  while True:
    print("hello process")

if __name__ == "__main__":
  for i in range(3):
    t = Process(target=loop)
    t.start()