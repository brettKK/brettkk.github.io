package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
)

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func build_tree_from_pre(arr []int) *TreeNode {
	arrChan := make(chan int, len(arr))
	for _, item := range arr {
		arrChan <- item
	}
	return build_pre(arrChan)
}

func build_pre(data chan int) *TreeNode {
	select {
	case val := <-data:
		if val == -1 {
			return nil
		}
		node := &TreeNode{val: val}
		node.left = build_pre(data)
		node.right = build_pre(data)
		return node
	default:
		return nil
	}
}

func pre_order(root *TreeNode) {
	if root == nil { return}
	fmt.Printf("%d \t", root.val)
	pre_order(root.left)
	pre_order(root.right)
}

func in_order(root *TreeNode) {
	if root == nil { return}
	pre_order(root.left)
	fmt.Printf("%d \t", root.val)
	pre_order(root.right)
}

func post_order(root *TreeNode) {
	if root == nil { return}
	post_order(root.left)
	post_order(root.right)
	fmt.Printf("%d \t", root.val)
}

func main() {
	// Ctrl+C  can exit the program...
	c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    go func() {
        <-c
        fmt.Printf("\nExiting\n")
        os.Exit(0)
    }()

REPL:for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>")
		inputStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("invalid input, try again")
			continue
		}
		strArr := strings.Split(strings.TrimSpace(inputStr), " ")
		var arrInt []int
		for _, str := range strArr {
			itemInt, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("invalid input=%s, try again", str)
				goto REPL
			}
			arrInt = append(arrInt, itemInt)
		}
		root := build_tree_from_pre(arrInt)
		fmt.Print("\npre_order:")
		pre_order(root)
		fmt.Print("\nin_order:")
		in_order(root)
		fmt.Print("\npost_order:")
		post_order(root)
		fmt.Println("")
	}
}