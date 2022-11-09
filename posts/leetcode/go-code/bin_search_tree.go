package main

type compareFunc func(c1 interface{}, c2 interface{}) bool

type TreeNode struct {
	left, right *TreeNode
	val interface{}
	lessFunc compareFunc
}

func new(compareF compareFunc) *TreeNode {
	return &TreeNode{
		val: nil, left: nil, right: nil,
		lessFunc: compareF,
	}
}

func insert(val interface{}, root *TreeNode) {
	if root.val == nil {
		root.val = val
		root.left = new(root.lessFunc)
		root.right = new(root.lessFunc)
		return
	}else {
		if root.lessFunc(val, root.val) {
			insert(val, root.left)
		} else {
			insert(val, root.right)
		}
	}
}

func search(val interface{}, root *TreeNode) *TreeNode{
	if root == nil || root.val == nil {
		return nil
	}
	if root.val == val {
		return root
	} else if root.lessFunc(root.val, val) {
		return search(val, root.right)
	} else {
		return search(val, root.left)
	}
}

func min(root *TreeNode) *TreeNode{
	if root == nil || root.left == nil {
		return root
	}
	return min(root.left)
}

func max(root *TreeNode) *TreeNode{
	if root == nil || root.right == nil {
		return root
	}
	return max(root.right)
}

func main() {

}
