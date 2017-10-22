


void memcpy(void *dst, void *src, int count)
{
  while (count--) {
    *dst = *src;
    dst++;
    src++;
  }
}


//问题是 void* 不能直接累加  *dst = * src 也是不对的 
void memcpy(void *dst, void *src, int count)
{
  unsigned char *pdst = (unsigned char *)dst;
  unsigned char *psrc = (unsigned char *)src;

  while (count--) {
    *pdst = *psrc;
    pdst++;
    psrc++;
  }
}



//原始内存地址是不能修改的， 还要判断空指针的情况
void memcpy(void *dst, const void *src, size_t count) {
  assert(dst != NULL);
  assert(src != NULL);
  unsigned char *pdst =  (unsigned char *)dst;
  const unsigned char *psrc = (const unsigned char *)src;

  while (count--) {
    *pdst = *psrc;
    pdst++;
    psrc++;
  }
}















