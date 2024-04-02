/*
 * @lc app=leetcode.cn id=maximum-difference-between-node-and-ancestor lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1026] 节点与其祖先之间的最大差值
 *
 * https://leetcode.cn/problems/maximum-difference-between-node-and-ancestor/description/
 *
 * algorithms
 * Medium (69.34%)
 * Total Accepted:    16.6K
 * Total Submissions: 23.6K
 * Testcase Example:  '[8,3,10,1,6,null,14,null,null,4,7,13]'
 *
 * 给定二叉树的根节点 root，找出存在于 不同 节点 A 和 B 之间的最大值 V，其中 V = |A.val - B.val|，且 A 是 B
 * 的祖先。
 * 
 * （如果 A 的任何子节点之一为 B，或者 A 的任何子节点是 B 的祖先，那么我们认为 A 是 B 的祖先）
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：root = [8,3,10,1,6,null,14,null,null,4,7,13]
 * 输出：7
 * 解释： 
 * 我们有大量的节点与其祖先的差值，其中一些如下：
 * |8 - 3| = 5
 * |3 - 7| = 4
 * |8 - 1| = 7
 * |10 - 13| = 3
 * 在所有可能的差值中，最大值 7 由 |8 - 1| = 7 得出。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：root = [1,null,2,null,0,3]
 * 输出：3
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树中的节点数在 2 到 5000 之间。
 * 0 
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

/*
 * public class Solution
 * {
 *     public int MaxAncestorDiff(TreeNode root)
 *     {
 *         var ans = 0;
 *         var ancestorValues = new List<int>();
 *         void traversal(TreeNode node)
 *         {
 *             if (node is null)
 *             {
 *                 return;
 *             }
 *             foreach (var v in ancestorValues)
 *             {
 *                 ans = Math.Max(ans, Math.Abs(v - node.val));
 *             }
 *             ancestorValues.Add(node.val);
 *             traversal(node.left);
 *             traversal(node.right);
 *             ancestorValues.RemoveAt(ancestorValues.Count() - 1);
 *         }
 *         traversal(root);
 *         return ans;
 *     }
 * }
 */

/*
 * public class Solution
 * {
 *     public int MaxAncestorDiff(TreeNode root)
 *     {
 *         var ans = 0;
 *         void traversal(TreeNode node, int min, int max)
 *         {
 *             if (node is null)
 *             {
 *                 return;
 *             }
 *             min = Math.Min(min, node.val);
 *             max = Math.Max(max, node.val);
 *             ans = Math.Max(ans, Math.Abs(node.val - min));
 *             ans = Math.Max(ans, Math.Abs(node.val - max));
 *             traversal(node.left, min, max);
 *             traversal(node.right, min, max);
 *         }
 *         traversal(root, int.MaxValue, 0);
 *         return ans;
 *     }
 * }
 */

public class Solution
{
    public int MaxAncestorDiff(TreeNode root)
    {
        var ans = 0;
        (int, int) traversal(TreeNode node)
        {
            if (node is null)
            {
                return (int.MaxValue, 0);
            }
            var (minLeft, maxLeft) = traversal(node.left);
            var (minRight, maxRight) = traversal(node.right);
            var min = Math.Min(node.val, Math.Min(minLeft, minRight));
            var max = Math.Max(node.val, Math.Max(maxLeft, maxRight));
            ans = Math.Max(ans, Math.Abs(node.val - min));
            ans = Math.Max(ans, Math.Abs(node.val - max));
            return (min, max);
        }
        traversal(root);
        return ans;
    }
}
