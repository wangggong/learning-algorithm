/*
 * @lc app=leetcode.cn id=root-equals-sum-of-children lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2236] 判断根结点是否等于子结点之和
 *
 * https://leetcode.cn/problems/root-equals-sum-of-children/description/
 *
 * algorithms
 * Easy (84.05%)
 * Total Accepted:    26.8K
 * Total Submissions: 32K
 * Testcase Example:  '[10,4,6]'
 *
 * 给你一个 二叉树 的根结点 root，该二叉树由恰好 3 个结点组成：根结点、左子结点和右子结点。
 * 
 * 如果根结点值等于两个子结点值之和，返回 true ，否则返回 false 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：root = [10,4,6]
 * 输出：true
 * 解释：根结点、左子结点和右子结点的值分别是 10 、4 和 6 。
 * 由于 10 等于 4 + 6 ，因此返回 true 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：root = [5,3,1]
 * 输出：false
 * 解释：根结点、左子结点和右子结点的值分别是 5 、3 和 1 。
 * 由于 5 不等于 3 + 1 ，因此返回 false 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树只包含根结点、左子结点和右子结点
 * -100 <= Node.val <= 100
 * 
 * 
 */
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     public int val;
 *     public TreeNode left;
 *     public TreeNode right;
 *     public TreeNode(int val=0, TreeNode left=null, TreeNode right=null) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
public class Solution
{
    public bool CheckTree(TreeNode root)
        => root.val == root.left.val + root.right.val;
}
