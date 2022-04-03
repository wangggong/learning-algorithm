/*
 * @lc app=leetcode.cn id=populating-next-right-pointers-in-each-node-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 *
 * https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/description/
 *
 * algorithms
 * Medium (62.73%)
 * Total Accepted:    113.7K
 * Total Submissions: 180.9K
 * Testcase Example:  '[1,2,3,4,5,null,7]'
 *
 * 给定一个二叉树
 *
 *
 * struct Node {
 * ⁠ int val;
 * ⁠ Node *left;
 * ⁠ Node *right;
 * ⁠ Node *next;
 * }
 *
 * 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
 *
 * 初始状态下，所有 next 指针都被设置为 NULL。
 *
 *
 *
 * 进阶：
 *
 *
 * 你只能使用常量级额外空间。
 * 使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
 *
 *
 *
 *
 * 示例：
 *
 *
 *
 *
 * 输入：root = [1,2,3,4,5,null,7]
 * 输出：[1,#,2,3,#,4,5,7,#]
 * 解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化输出按层序遍历顺序（由 next
 * 指针连接），'#' 表示每层的末尾。
 *
 *
 *
 * 提示：
 *
 *
 * 树中的节点数小于 6000
 * -100 <= node.val <= 100
 *
 *
 *
 *
 *
 *
 *
 */
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	dfs(root, nil)
	return root
}

func dfs(root, next *Node) {
	if root == nil {
		return
	}
	root.Next = next
	var childNext *Node = nil
	for curr := next; curr != nil; curr = curr.Next {
		if curr.Left != nil {
			childNext = curr.Left
			break
		}
		if curr.Right != nil {
			childNext = curr.Right
			break
		}
	}
	if root.Left == nil && root.Right == nil {
		return
	}
	if root.Left != nil && root.Right != nil {
		dfs(root.Right, childNext)
		dfs(root.Left, root.Right)
		return
	}
	if root.Left != nil {
		dfs(root.Left, childNext)
		return
	}
	dfs(root.Right, childNext)
	return
}
