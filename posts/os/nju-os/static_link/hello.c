#include <stdio.h>
#include <stdint.h>
#include <assert.h>

int main();

void hello() {
    printf("Hello, world!\n");

    char *p = (char*)main + 0xa + 1;
    int32_t offset = *(int32_t *)p;
    // next pc + offset = function address
    assert((char *)main + 0xf + offset == (char *)hello) ;
}