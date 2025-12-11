package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	present := root
	for present.Right != nil {
		present = present.Right
	}
	return present
}
