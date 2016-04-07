#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(int argc, char *argv[])
{
	printf("int %d\n", 0xBA);
	//printf("string %s\n", 0xBA);
	printf("string %c\n", 0xBA);

	char *ptr;
	ptr = malloc(1000);
	memset(ptr, 0xBA, 500);
	printf("int ptr %d\n", ptr);
	printf("string ptr %s\n", ptr);

}
