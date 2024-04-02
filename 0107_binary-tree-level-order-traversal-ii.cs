/*
 * @lc app=leetcode.cn id=binary-tree-level-order-traversal-ii lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [107] 二叉树的层序遍历 II
 *
 * https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Medium (73.15%)
 * Total Accepted:    310.5K
 * Total Submissions: 423.8K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：[[15,7],[9,20],[3]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：root = [1]
 * 输出：[[1]]
 * 
 * 
 * 示例 3：
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
 * 树中节点数目在范围 [0, 2000] 内
 * -1000 <= Node.val <= 1000
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
    public IList<IList<int>> LevelOrderBottom(TreeNode root)
    {
        var ans = new List<IList<int>>();
        var Q = new Queue<TreeNode>();
        if (root is not null) { Q.Enqueue(root); }
        while (Q.Count > 0)
        {
            var cur = new List<int>();
            for (var c = Q.Count; c > 0; c--)
            {
                var node = Q.Dequeue();
                cur.Add(node.val);
                if (node.left is not null) { Q.Enqueue(node.left); }
                if (node.right is not null) { Q.Enqueue(node.right); }
            }
            ans.Add(cur);
        }
        ans.Reverse();
        return ans;
    }
}
