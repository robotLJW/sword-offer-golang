package _4_er_cha_shu_zhong_he_wei_mou_yi_zhi_de_lu_jing_lcof

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	ans := make([][]int, 0)
	tempData := make([]int, 0)
	var dfs func(node *TreeNode, currentValue int)
	dfs = func(node *TreeNode, currentValue int) {
		if node == nil {
			return
		}
		tempData = append(tempData, node.Val)
		if currentValue-node.Val == 0 && node.Left == nil && node.Right == nil {
			ans = append(ans, append([]int{}, tempData...))
		}
		if node.Left != nil {
			dfs(node.Left, currentValue-node.Val)
		}
		if node.Right != nil {
			dfs(node.Right, currentValue-node.Val)
		}
		tempData = tempData[:len(tempData)-1]
	}
	dfs(root, sum)
	return ans
}
