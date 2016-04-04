#include <stdio.h>

typedef unsigned char *byte_pointer;    // byte_pointer 就是等于 unsigned char

void show_bytes(byte_pointer start, int len )
{
  int i;
  for (i = 0; i < len; i++) {
    printf("%.2x\n", start[i]);
  }
  printf("\n");
}

void show_int(int x) {
  show_bytes((byte_pointer) &x, sizeof(int));
}

void show_float(float x) {
  show_bytes((byte_pointer) &x, sizeof(float));
}

void show_pointer(void *x) {
  show_bytes((byte_pointer)  &x, sizeof(void *));
}



int main(int argc, char **argv)
{
  const char *s = "abcdef";
  show_bytes( (byte_pointer) s, strlen(s) );
  return 0;
}
