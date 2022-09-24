#include <stdio.h>
#include <unistd.h>
// #include <sys/types.h>
#include <sys/wait.h>

// ./a.out  6个Hello
// clang fork_test_printf.c| ./a.out | wc -l  8 个hello
int main() {
    // setbuf(stdout, NULL);
    for (int i = 0; i < 2; i++) {
        fork(); // 无情的复制，库函数的内部状态也复制一份。
        printf("Hello\n");
    }
    for (int i = 0; i < 2; i++) {
        wait(NULL);
    }
    return 0;
}
// 库函数的实现。
// stdout
//        tty           line buffer（写满一行后，看到\n后调用 系统调用写出来）
//        file, pipe    full buffer (写满4096B后，才会调用 系统调用写出来)

void my_printf() {
    printf("Hello"); fflush(stdout);
    printf("World\n");
    int *p;
    p = NULL;
    *p = 1;
}