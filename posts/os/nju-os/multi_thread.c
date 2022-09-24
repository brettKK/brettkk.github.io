#include "thread.h"
#include <stdio.h>


// gcc multi_thread.c -lpthread && ./a.out

// also ok gcc multi_thread.c && ./a.out


void Ta() {
	while(1)  { printf("a");
}}

void Tb() {
	while(1) { printf("b");  }
}

int main() {
	create(Ta);
	create(Tb);
}
