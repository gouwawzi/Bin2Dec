#include <stdio.h>
#include <stdlib.h>
#include "convert.h"


int main(int argc, char **argv) {
	unsigned long int result = 0;	
	int err = convert(&result, argv[1]);
	
	if (err == -1) {
		fprintf(stderr, "'%s' is too long\n", argv[1]);
		exit(1);
	}

	if (err == -2) {
		fprintf(stderr, "'%s' is a valid binary number\n", argv[1]);
		exit(2);
	}

	fprintf(stdout, "%lu\n", result);
}

