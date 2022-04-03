/*
 * @lc app=leetcode.cn id=sum-root-to-leaf-numbers lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [129] 求根节点到叶节点数字之和
 *
 * https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/description/
 *
 * algorithms
 * Medium (68.87%)
 * Total Accepted:    139.5K
 * Total Submissions: 202.5K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
 *
 *
 * 每条从根节点到叶节点的路径都代表一个数字：
 *
 *
 * 例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
 *
 *
 * 计算从根节点到叶节点生成的 所有数字之和 。
 *
 * 叶节点 是指没有子节点的节点。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,3]
 * 输出：25
 * 解释：
 * 从根到叶子节点路径 1->2 代表数字 12
 * 从根到叶子节点路径 1->3 代表数字 13
 * 因此，数字总和 = 12 + 13 = 25
 *
 * 示例 2：
 *
 *
 * 输入：root = [4,9,0,5,1]
 * 输出：1026
 * 解释：
 * 从根到叶子节点路径 4->9->5 代表数字 495
 * 从根到叶子节点路径 4->9->1 代表数字 491
 * 从根到叶子节点路径 4->0 代表数字 40
 * 因此，数字总和 = 495 + 491 + 40 = 1026
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目在范围 [1, 1000] 内
 * 0 <= Node.val <= 9
 * 树的深度不超过 10
 *
 *
 *
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// import "strconv"

func sumNumbers(root *TreeNode) int {
	// strs := dfs(root)
	// ans := 0
	// for _, s := range strs {
	// 	n, _ := strconv.Atoi(s)
	// 	ans += n
	// }
	// return ans

	return dfs(root, 0)
}

func dfs(root *TreeNode, n int) int {
	if root == nil {
		return 0
	}
	ans := n*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return ans
	}

	return dfs(root.Left, ans) + dfs(root.Right, ans)
}

/*
 * func dfs(root *TreeNode) []string {
 * 	if root == nil {
 * 		return nil
 * 	}
 * 	var ans []string
 * 	left := dfs(root.Left)
 * 	right := dfs(root.Right)
 * 	s := string([]byte{byte(root.Val) + '0'})
 * 	if len(left) == 0 && len(right) == 0 {
 * 		ans = append(ans, s)
 * 	}
 * 	for _, l := range left {
 * 		ans = append(ans, s+l)
 * 	}
 * 	for _, r := range right {
 * 		ans = append(ans, s+r)
 * 	}
 * 	return ans
 * }
 */
