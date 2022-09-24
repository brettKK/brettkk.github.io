#include <stdio.h>
extern int done;

// void join() {
// 	while (!done) {
// 		;
// 	}
// }

//clang a_example.c && ./a.out
//objdump -d a.out| less
//100003f54: ff c3 00 d1
int main() {
	unsigned *p;
	p = (void *)main;
	printf("%x\n", *p);
	// sizeof(int);
	return 0;
}