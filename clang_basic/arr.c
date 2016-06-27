#include <stdio.h>

int main(int argc, char **argv[]) 
{
	char *const xx[] = {
		['a'] = "hello",
		['b'] = "world",
		['c'] = "you",
	};

	printf("%s\n", xx['a']);
	printf("%x\n", xx);
}
