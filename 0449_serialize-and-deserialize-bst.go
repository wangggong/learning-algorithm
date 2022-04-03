/*
 * @lc app=leetcode.cn id=serialize-and-deserialize-bst lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [449] 序列化和反序列化二叉搜索树
 *
 * https://leetcode-cn.com/problems/serialize-and-deserialize-bst/description/
 *
 * algorithms
 * Medium (56.20%)
 * Total Accepted:    19K
 * Total Submissions: 33.7K
 * Testcase Example:  '[2,1,3]'
 *
 * 序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。
 *
 * 设计一个算法来序列化和反序列化 二叉搜索树 。 对序列化/反序列化算法的工作方式没有限制。
 * 您只需确保二叉搜索树可以序列化为字符串，并且可以将该字符串反序列化为最初的二叉搜索树。
 *
 * 编码的字符串应尽可能紧凑。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [2,1,3]
 * 输出：[2,1,3]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数范围是 [0, 10^4]
 * 0 <= Node.val <= 10^4
 * 题目数据 保证 输入的树是一棵二叉搜索树。
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

import "fmt"

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	return fmt.Sprintf("[%s]%d[%s]", c.serialize(root.Left), root.Val, c.serialize(root.Right))
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	return c.inorder([]byte(data), 0, len(data))
}

func (c *Codec) inorder(arr []byte, l, r int) *TreeNode {
	if l >= r {
		return nil
	}
	// assert arr[l] == '[';
	p := 0
	ind := l
	for ; ind < r; ind++ {
		if arr[ind] == '[' {
			p++
		} else if arr[ind] == ']' {
			p--
		}
		if p == 0 {
			break
		}
	}
	ind++
	nInd := ind
	val := 0
	// 傻逼! 判定逻辑得是当前 index 的字符是否为 '['!
	for ; nInd < r && arr[nInd] != '['; nInd++ {
		val = val*10 + int(arr[nInd]-'0')
	}
	// 傻逼! 这里得把壳去了!
	return &TreeNode{
		Val:   val,
		Left:  c.inorder(arr, l+1, ind-1),
		Right: c.inorder(arr, nInd+1, r-1),
	}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
