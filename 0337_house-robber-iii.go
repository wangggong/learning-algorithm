/*
 * @lc app=leetcode.cn id=house-robber-iii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [337] 打家劫舍 III
 *
 * https://leetcode-cn.com/problems/house-robber-iii/description/
 *
 * algorithms
 * Medium (60.73%)
 * Total Accepted:    166.1K
 * Total Submissions: 273.6K
 * Testcase Example:  '[3,2,3,null,3,null,1]'
 *
 * 小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。
 *
 * 除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果
 * 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。
 *
 * 给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入: root = [3,2,3,null,3,null,1]
 * 输出: 7
 * 解释: 小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7
 *
 * 示例 2:
 *
 *
 *
 *
 * 输入: root = [3,4,5,1,3,null,1]
 * 输出: 9
 * 解释: 小偷一晚能够盗取的最高金额 4 + 5 = 9
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 *
 * 树的节点数在 [1, 10^4] 范围内
 * 0 <= Node.val <= 10^4
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

/*
 * var m map[*TreeNode][2]int
 *
 * func rob(root *TreeNode) int {
 * 	m = make(map[*TreeNode][2]int)
 * 	ans := traversal(root)
 * 	return max(ans[0], ans[1])
 * }
 *
 * func traversal(root *TreeNode) [2]int {
 * 	if root == nil {
 * 		return [2]int{0, 0}
 * 	}
 * 	if _, ok := m[root]; !ok {
 * 		left, right := traversal(root.Left), traversal(root.Right)
 * 		m[root] = [2]int{
 * 			root.Val + left[1] + right[1],
 * 			max(left[0], left[1]) + max(right[0], right[1]),
 * 		}
 * 	}
 * 	return m[root]
 * }
 */

var m map[*TreeNode]int

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	m = make(map[*TreeNode]int)
	return max(_rob(root), _rob(root.Left)+_rob(root.Right))
}

func _rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if v, ok := m[root]; ok {
		return v
	}
	ans := root.Val
	defer func() { m[root] = ans }()
	if root.Left != nil {
		ans += _rob(root.Left.Left) + _rob(root.Left.Right)
	}
	if root.Right != nil {
		ans += _rob(root.Right.Left) + _rob(root.Right.Right)
	}
	ans = max(ans, _rob(root.Left)+_rob(root.Right))
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}


