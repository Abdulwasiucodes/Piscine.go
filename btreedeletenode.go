package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node.Left == nil {
		return BTreeTransplant(root, node, node.Right)
	} else if node.Right == nil {
		return BTreeTransplant(root, node, node.Left)
	} else {
		successor := BTreeMin(node.Right)
		if successor.Parent != node {
			root = BTreeTransplant(root, successor, successor.Right)
			successor.Right = node.Right
			if successor.Right != nil {
				successor.Right.Parent = successor
			}
		}
		root = BTreeTransplant(root, node, successor)
		successor.Left = node.Left
		if successor.Left != nil {
			successor.Left.Parent = successor
		}
		return root
	}
}
