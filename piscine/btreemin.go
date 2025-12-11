package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	present := root
	for present.Left != nil {
		present = present.Left
	}
	return present
}
