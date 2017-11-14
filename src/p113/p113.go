package p113

func deepCopyInts(ns []int) []int {
	ns1 := make([]int, len(ns))
	copy(ns1, ns)
	return ns1
}

func pathSum(root *TreeNode, sum int) [][]int {
	pathes := [][]int {}

	var findPath func (root *TreeNode, sum int, path []int)
	findPath = func (root *TreeNode, sum int, path []int) {
		if root == nil {
			return
		}

		if root.Val == sum && root.Left == nil && root.Right == nil {
			pathes = append(pathes, append(path, root.Val))
		}

		path1 := deepCopyInts(path)
		findPath(root.Left, sum - root.Val, append(path1, root.Val))

		path2 := deepCopyInts(path)
		findPath(root.Right, sum - root.Val, append(path2, root.Val))
	}

	findPath(root, sum, []int {})

	return pathes
}
