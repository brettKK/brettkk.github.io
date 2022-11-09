package main

type (
	node struct {
		val int
		pre *node
	}
	stack struct {
		top *node
		len int
	}
)

func new_stack() *stack { return &stack{nil, 0} }

func push(val int, stk *stack) {
	new_item := &node{val: val, pre: stk.top}
	stk.top = new_item
	// one line is ok: stk.top = &StackItem{val: val, pre: stk.top}
	stk.len++
}
func peek(stk *stack) int {
	if stk == nil || stk.len == 0 {
		return -1
	}
	return stk.top.val
}
func pop(stk *stack) int {
	if stk == nil || stk.top == nil {
		return -1
	}
	res := stk.top.val
	stk.top = stk.top.pre
	stk.len--
	return res
}
