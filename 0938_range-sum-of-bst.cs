/*
 * @lc app=leetcode.cn id=range-sum-of-bst lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [938] 二叉搜索树的范围和
 *
 * https://leetcode.cn/problems/range-sum-of-bst/description/
 *
 * algorithms
 * Easy (82.11%)
 * Total Accepted:    147.7K
 * Total Submissions: 177.8K
 * Testcase Example:  '[10,5,15,3,7,null,18]\n7\n15'
 *
 * 给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：root = [10,5,15,3,7,null,18], low = 7, high = 15
 * 输出：32
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
 * 输出：23
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树中节点数目在范围 [1, 2 * 10^4] 内
 * 1 
 * 1 
 * 所有 Node.val 互不相同
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
    public int RangeSumBST(TreeNode root, int low, int high)
    {
        var ans = 0;
        void traversal(TreeNode node)
        {
            if (node is not null)
            {
                if (low <= node.val && node.val <= high) { ans += node.val; }
                traversal(node.left);
                traversal(node.right);
            }
        }
        traversal(root);
        return ans;
    }
}
