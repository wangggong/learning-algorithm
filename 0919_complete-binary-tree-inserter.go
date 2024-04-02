/*
 * @lc app=leetcode.cn id=complete-binary-tree-inserter lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [919] 完全二叉树插入器
 *
 * https://leetcode.cn/problems/complete-binary-tree-inserter/description/
 *
 * algorithms
 * Medium (65.91%)
 * Total Accepted:    24.2K
 * Total Submissions: 36.2K
 * Testcase Example:  '["CBTInserter","insert","get_root"]\n[[[1]],[2],[]]'
 *
 * 完全二叉树 是每一层（除最后一层外）都是完全填充（即，节点数达到最大）的，并且所有的节点都尽可能地集中在左侧。
 *
 * 设计一种算法，将一个新节点插入到一个完整的二叉树中，并在插入后保持其完整。
 *
 * 实现 CBTInserter 类:
 *
 *
 * CBTInserter(TreeNode root) 使用头节点为 root 的给定树初始化该数据结构；
 * CBTInserter.insert(int v)  向树中插入一个值为 Node.val == val的新节点
 * TreeNode。使树保持完全二叉树的状态，并返回插入节点 TreeNode 的父节点的值；
 * CBTInserter.get_root() 将返回树的头节点。
 *
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入
 * ["CBTInserter", "insert", "insert", "get_root"]
 * [[[1, 2]], [3], [4], []]
 * 输出
 * [null, 1, 2, [1, 2, 3, 4]]
 *
 * 解释
 * CBTInserter cBTInserter = new CBTInserter([1, 2]);
 * cBTInserter.insert(3);  // 返回 1
 * cBTInserter.insert(4);  // 返回 2
 * cBTInserter.get_root(); // 返回 [1, 2, 3, 4]
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数量范围为 [1, 1000]
 * 0 <= Node.val <= 5000
 * root 是完全二叉树
 * 0 <= val <= 5000
 * 每个测试用例最多调用 insert 和 get_root 操作 10^4 次
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
type CBTInserter struct {
	N    int
	Root *TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	n := 0
	var Q []*TreeNode
	Q = append(Q, root)
	for len(Q) > 0 {
		q := Q[0]
		Q = Q[1:]
		n++
		if q.Left != nil {
			Q = append(Q, q.Left)
		}
		if q.Right != nil {
			Q = append(Q, q.Right)
		}
	}
	return CBTInserter{
		Root: root,
		N:    n,
	}
}

func (c *CBTInserter) Insert(val int) int {
	c.N++
	cur, node := c.Root, &TreeNode{Val: val}
	for k := bitCnt(c.N) - 2; k > 0; k-- {
		switch c.N & (1 << k) {
		case 0:
			cur = cur.Left
		default:
			cur = cur.Right
		}
	}
	switch c.N & 1 {
	case 0:
		cur.Left = node
	default:
		cur.Right = node
	}
	return cur.Val
}

func (c *CBTInserter) Get_root() *TreeNode {
	return c.Root
}

func bitCnt(n int) int {
	ans := 0
	for ; n > 0; n >>= 1 {
		ans++
	}
	return ans
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
