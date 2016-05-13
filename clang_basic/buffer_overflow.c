char *gets(char *s)
{
  int c;
  char *dest = s;
  int gotchar = 0;   //has at least one character been read?
  while ( (c = getchar()) != '\n' && c != EOF) {
    *dest++ = c;    //no bounds checking
    gotchar = 1;
  }
  
  *dest++ = '\0';   //terminate string
  if (c == EOF && !gotchar)
    return NULL;    //end of file or error
  return s;
}

//read input line and write it back
void echo()
{
  char buf[8];    //way too small
  gets(buf);
  puts(buf);
}

