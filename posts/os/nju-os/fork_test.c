#include <unistd.h>
#include <stdio.h>

int main(){
    pid_t pid1 = fork();
    pid_t pid2 = fork();
    pid_t pid3 = fork();
    printf("hello world from (%d, %d, %d)\n", pid1, pid2, pid3);
    return 0;
}

/*
画状态机

pc  = 5 pid ~ 2

pc = 6 pid ~ 2 pid1 = 3
pc = 6 pid ~ 3 pid1 = 0

pc = 7 pid ~ 2 pid1 = 3,  pid2 = 4
pc = 7 pid ~ 4 pid1 = 3,  pid2 = 0
pc = 7 pid ~ 3 pid1 = 0,  pid2 = 5
pc = 7 pid ~ 5 pid1 = 0,  pid2 = 0

pc = 8 pid ~ 2 pid1 = 3, pid2 = 4, pid3 = 6 
pc = 8 pid ~ 6 pid1 = 3, pid2 = 4, pid3 = 0
pc = 8 pid ~ 4 pid1 = 3, pid2 = 0, pid3 = 7
pc = 8 pid ~ 7 pid1 = 3, pid2 = 0, pid3 = 0
pc = 8 pid ~ 3 pid1 = 0, pid2 = 5, pid3 = 8
pc = 8 pid ~ 8 pid1 = 0, pid2 = 5, pid3 = 0
pc = 8 pid ~ 5 pid1 = 0, pid2 = 0, pid3 = 9
pc = 8 pid ~ 9 pid1 = 0, pid2 = 0, pid3 = 0 

pc = 9 时输出8次。
*/