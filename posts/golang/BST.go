package main

type TreeNode struct {
	left, right *TreeNode
	val int
}

type BST struct {
	root *TreeNode
}

// insert_BST return root of the tree
func insert_BST(tree *BST, val int) *TreeNode{
	if node := search(tree.root, val); node != nil {
		return tree.root
	}
	node := &TreeNode{left: nil, right: nil, val: val}
	parent, cur := tree.root, tree.root
	for cur.val != val {
		parent = cur
		if cur.val > val {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	// assert parent  not nil
	if parent.val > val {
		parent.left = node
	} else {
		parent.right = node
	}
	return tree.root
}
//insert_recursion return The root of tree
func insert_recursion(root *TreeNode, val int) *TreeNode{
	if root == nil {
		return &TreeNode{left: nil, right: nil, val: val}
	}
	if root.val == val {
		return root
	} else if root.val > val {
		root.left = insert_recursion(root.left, val)
		return root
	} else if root.val < val {
		root.right = insert_recursion(root.right, val)
		return root
	}
	// unreachable
	return nil
}

// return the target node with the `val` value
func search_BST(tree *BST, val int) *TreeNode {
	return search(tree.root, val)
}
func search(root *TreeNode, val int) *TreeNode {
	if (root == nil) {return nil}
	if (root.val == val) {return root}
	if (root.val > val) { return search(root.left, val) }
	if (root.val < val) { return search(root.right, val) }
	return nil
}

// delete_node_BST return the root of the tree
func delete_node_BST(tree *BST, val int) *TreeNode {
	return delete_node(tree.root, val)
}

// delete_node return the root of bst after delete node 
func delete_node(root *TreeNode, val int) *TreeNode {
	if root == nil {return nil}
	if root.val > val { 
		root.left = delete_node(root.left, val) 
		return root
	}
	if root.val < val { 
		root.right =  delete_node(root.right, val) 
		return root
	}
	// root.val == val , get the target root == delete_node
	// case 1
	if root.left == nil { return root.right }
	// case 2
	if root.right == nil { return root.left }
	// case 3  target node 's left and right is not nil
	//get right side smallest node which little bigger than val
	minnode := get_min(root.right)
	root.val = minnode.val
	root.right = delete_min(root.right)
	// or get left side biggest node which little smaller than val also ok...
	return root
}

// return the leftest node
func get_min(root *TreeNode) *TreeNode {
	cur := root
	for (cur.left != nil) {
		cur = cur.left
	}
	return cur
}

// delete_min return the root of the tree
func delete_min(root *TreeNode) *TreeNode {
	if (nil == root.left) {
		pRight := root.right
		root.right = nil
		return pRight
	}
	root.left = delete_min(root.left)
	return root
}

func main() {
}