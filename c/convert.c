#include <string.h>

int convert(unsigned long int *r, char *input) {
	*r = 0;
	int w = 0;
	char *i = input;

	if (strlen(input) > 32) {
		return -1;
	}	


	while (*i != '\0') {
		if (*i < '0' || *i > '1') {
			return -2;
		}	

		*r += (*i-48) * (1<<w);
		i++;
		w++;
	}
	return 0;
}
