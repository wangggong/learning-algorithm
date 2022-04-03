/*
 * @lc app=leetcode.cn id=serialize-and-deserialize-binary-tree lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [297] 二叉树的序列化与反序列化
 *
 * https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/description/
 *
 * algorithms
 * Hard (56.72%)
 * Total Accepted:    136.2K
 * Total Submissions: 239.1K
 * Testcase Example:  '[1,2,3,null,null,4,5]'
 *
 *
 * 序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
 *
 * 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 /
 * 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
 *
 * 提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode
 * 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,3,null,null,4,5]
 * 输出：[1,2,3,null,null,4,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1]
 * 输出：[1]
 *
 *
 * 示例 4：
 *
 *
 * 输入：root = [1,2]
 * 输出：[1,2]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中结点数在范围 [0, 104] 内
 * -1000 <= Node.val <= 1000
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

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "nil"
	}
	return fmt.Sprintf("%d|%s|%s", root.Val, c.serialize(root.Left), c.serialize(root.Right))
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(s string) *TreeNode {
	vals := strings.Split(s, "|")
	k := 0
	return preorder(vals, &k, len(vals))
}

func preorder(arr []string, p *int, q int) *TreeNode {
	if *p >= q {
		return nil
	}
	// 1|2|n|n|3|n|n => [1,2,n,n,3,n,n]
	if arr[*p] == "nil" {
		*p = *p + 1
		return nil
	}
	v, _ := strconv.Atoi(arr[*p])
	root := &TreeNode{Val: v}
	*p = *p + 1
	root.Left = preorder(arr, p, q)
	root.Right = preorder(arr, p, q)
	return root

}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
