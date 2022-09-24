/*
$python3
>>> def numbers(init=0, step=1):
...     n = init
...     while True:
...         n += step
...         yield n
... 
>>> numbers
<function numbers at 0x102ae0af0>
>>> g = numbers()
>>> g
<generator object numbers at 0x102accf20>
>>> g.__next__()
1
>>> g.__next__()
2
>>> 
*/
/*

维护状态机。
python generator 

*/
def numbers(init=0, step=1):
    n = init
    while True:
        n += step
        yield n

