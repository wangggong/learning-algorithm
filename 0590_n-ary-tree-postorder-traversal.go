/*
 * @lc app=leetcode.cn id=n-ary-tree-postorder-traversal lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [590] N 叉树的后序遍历
 *
 * https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/description/
 *
 * algorithms
 * Easy (76.55%)
 * Total Accepted:    82.4K
 * Total Submissions: 106.6K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 n 叉树的根节点 root ，返回 其节点值的 后序遍历 。
 *
 * n 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：root = [1,null,3,2,4,null,5,6]
 * 输出：[5,6,3,2,4,1]
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：root =
 * [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
 * 输出：[2,6,14,11,7,3,12,8,4,13,9,10,5,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 节点总数在范围 [0, 10^4] 内
 * 0 <= Node.val <= 10^4
 * n 叉树的高度小于或等于 1000
 *
 *
 *
 *
 * 进阶：递归法很简单，你可以使用迭代法完成此题吗?
 *
 */
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

var ans []int

/*
 * func postorder(root *Node) []int {
 * 	ans = nil
 * 	_postorder(root)
 * 	return ans
 * }
 */

/*
 * // 迭代 ver. 0
 * func postorder(root *Node) []int {
 * 	var S []*Node
 * 	ans = nil
 * 	if root == nil {
 * 		return ans
 * 	}
 * 	visit := make(map[*Node]struct{})
 * 	S = append(S, root)
 * 	for len(S) > 0 {
 * 		node := S[len(S)-1]
 * 		if _, ok := visit[node]; !ok {
 * 			for i := len(node.Children) - 1; i >= 0; i-- {
 * 				c := node.Children[i]
 * 				S = append(S, c)
 * 			}
 * 			visit[node] = struct{}{}
 * 			continue
 * 		}
 * 		ans = append(ans, node.Val)
 * 		S = S[:len(S)-1]
 * 	}
 * 	return ans
 * }
 */

// 迭代 ver. 1
func postorder(root *Node) []int {
	ans = nil
	if root == nil {
		return ans
	}
	var S []*Node
	var postS []*Node
	S = append(S, root)
	for len(S) > 0 {
		node := S[len(S)-1]
		S = S[:len(S)-1]
		// ans = append(ans, node.Val)
		postS = append(postS, node)
		for _, c := range node.Children {
			S = append(S, c)
		}
	}
	for len(postS) > 0 {
		l := len(postS)
		ans = append(ans, postS[l-1].Val)
		postS = postS[:l-1]
	}
	// for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
	// 	ans[i], ans[j] = ans[j], ans[i]
	// }
	return ans
}

func _postorder(root *Node) {
	if root == nil {
		return
	}
	for _, c := range root.Children {
		_postorder(c)
	}
	ans = append(ans, root.Val)
}
