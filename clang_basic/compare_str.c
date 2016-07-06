#include <string.h>
#include <stdio.h>

int main(void)
{
	char *buf1 = "aaa";
	char *buf2 = "bbb";
	char *buf3 = "ccc";
	int ptr;

	ptr = strcmp(buf2, buf1);
	
	int xx;
	char buf[1024];
	//char buf[1024] = "abccjsdkfjsdiurerrwerwerwer";
	xx = strcmp("\n", buf);

	int yy;
	char tmp[1024];
	tmp[0] = 'A';
	tmp[1] = '\0';
	yy = strcmp("\n", tmp);



	printf("%d\n", ptr);
	printf("%d\n", xx);
	printf("%d\n", yy);


	char k[1024];
	k[0] = 'A';
	k[1] = '\0';
	printf("strcmp trturn:  %d\n", strcmp("\n", k));
	if (strcmp("\n", k) ) 
	{
	  printf("1111\n");
	}
	else 
	{
	  printf("2222\n");
	}

  if ( -55  ) 
	{
	  printf("1111\n");
	}
	else 
	{
	  printf("2222\n");
	}
 
}

