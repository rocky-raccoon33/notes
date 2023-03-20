#include <stdlib.h>
#include <stdio.h>
int main()
{
	int *t = (int *)malloc(sizeof(int));

	free(t);
	return 0;
}