#include "csapp.h"
#define N 2
void *thread(void *vargp);

char **ptr;   //global variable   全局变量, 全局共享

int main() 
{
  int i;
  pthread_t tid;
  char **msgs[N] = {
    "Hello from foo",
    "Hello from bar"
  };
  
  ptr = msgs;
  for (i = 0; i < N; i++) {
    Pthread_create(&tid, NULL, thread, (void *)i);
  }
    pthread_exit(NULL);
}

void *thread(void *vargp)
{
  int myid = (int)vargp;    //本地自动变量
  static int cnt = 0;       //本地静态变量， 每个函数独享
  printf("[%d]:   %s (cnt=%d)\n", myid, ptr[myid], ++cnt);
  return NULL;
}