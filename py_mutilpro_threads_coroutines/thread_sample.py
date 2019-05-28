from threading import Thread

def loop():
  while True:
    print("hello threads")

if __name__ == "__main__":
  for i in range(3):
    t = Thread(target=loop)
    t.start()