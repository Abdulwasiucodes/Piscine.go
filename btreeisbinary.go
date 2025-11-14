package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return isBSTHelper(root, "", "")
}

func isBSTHelper(node *TreeNode, min, max string) bool {
	if node == nil {
		return true
	}

	if (min != "" && node.Data <= min) || (max != "" && node.Data >= max) {
		return false
	}

	return isBSTHelper(node.Left, min, node.Data) &&
		isBSTHelper(node.Right, node.Data, max)
}
