package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBST(root, nil, nil)
}

func isBST(node *TreeNode, min *string, max *string) bool {
	if node == nil {
		return true
	}

	if min != nil && node.Data <= *min {
		return false
	}
	if max != nil && node.Data >= *max {
		return false
	}

	return isBST(node.Left, min, &node.Data) && isBST(node.Right, &node.Data, max)
}
