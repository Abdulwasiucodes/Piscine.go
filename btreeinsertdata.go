package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	current := root
	var parent *TreeNode

	for current != nil {
		parent = current
		if data < current.Data {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	newNode := &TreeNode{Data: data, Parent: parent}
	if data < parent.Data {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}

	return root
}
