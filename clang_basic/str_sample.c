#include <stdio.h>

int main(int argc, char **argv[])
{
	char *query_string = NULL;
	char url[1024] = "/wenqian1991/article/details/46011357?xxx=aaa";
	
	query_string = url;
	//printf("%d\n", query_string);

	while ( (*query_string != '?') && (*query_string != '\0') ){
		//printf("%d\n", query_string);
		//printf("%c\n", *query_string);
		//printf("------------------------------\n");
		query_string++;
	}

	if (*query_string == '?') 
	{
		printf("------------- into ? -----------------\n");
		*query_string = '\0';
		query_string++;
	}

	printf("%d\n", query_string);
	printf("%c\n", *query_string);
	printf("%s\n", query_string);

}
