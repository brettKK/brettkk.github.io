void hello();

int main() {
    hello();
}


/*
gcc -c main.c
gcc -c hello.c

gcc -static main.o hello.o 
//Static linking does not work on mac os x with the llvm/clang gcc
//ld: library not found for -lcrt0.o
//clang: error: linker command failed with exit code 1 (use -v to see invocation)

./a.out

nm main.o
编译成汇编正常完成。

gcc -S -c main.c
vim main.s

objdump -d main.o

静态链接时。

0xa: callq 00 00 00 00 // xxx = S + A - P
0xf: 31 c0

P=main + oxb （当前需要改写为重定向的地址）

offset 与next_pc的地址相差4个字节。

next_pc = main + oxf 
offset = S + A - P = hello + (-4) - P = hello - 4 - （main + 0xb）
next_pc + offset = next_pc + (S + A - P) 
                 =  main +0xf + hello - 4 - main - 0xb
                 = hello函数的地址

所以 callq offset 就是这样计算出来的。 callq offset + next_pc = callq hello.

静态链接完成后a.out里 offset的值已经填好了。

*/
