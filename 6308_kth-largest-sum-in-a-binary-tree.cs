/*
 * @lc app=leetcode.cn id=kth-largest-sum-in-a-binary-tree lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6308] 二叉树中的第 K 大层和
 *
 * https://leetcode.cn/problems/kth-largest-sum-in-a-binary-tree/description/
 *
 * algorithms
 * Medium (42.35%)
 * Total Accepted:    5.3K
 * Total Submissions: 12.6K
 * Testcase Example:  '[5,8,9,2,1,3,7,4,6]\n2'
 *
 * 给你一棵二叉树的根节点 root 和一个正整数 k 。
 * 
 * 树中的 层和 是指 同一层 上节点值的总和。
 * 
 * 返回树中第 k 大的层和（不一定不同）。如果树少于 k 层，则返回 -1 。
 * 
 * 注意，如果两个节点与根节点的距离相同，则认为它们在同一层。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：root = [5,8,9,2,1,3,7,4,6], k = 2
 * 输出：13
 * 解释：树中每一层的层和分别是：
 * - Level 1: 5
 * - Level 2: 8 + 9 = 17
 * - Level 3: 2 + 1 + 3 + 7 = 13
 * - Level 4: 4 + 6 = 10
 * 第 2 大的层和等于 13 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：root = [1,2,null,3], k = 1
 * 输出：3
 * 解释：最大的层和是 3 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树中的节点数为 n
 * 2 <= n <= 10^5
 * 1 <= Node.val <= 10^6
 * 1 <= k <= n
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
public class Solution
{
    public long KthLargestLevelSum(TreeNode root, int k)
    {
        var Q = new Queue<TreeNode>();
        Q.Enqueue(root);
        var sums = new List<long>();
        while (Q.Count > 0)
        {
            long tot = 0;
            for (var c = Q.Count; c > 0; c--)
            {
                var q = Q.Dequeue();
                tot += (long)q.val;
                if (q.left != null)
                {
                    Q.Enqueue(q.left);
                }
                if (q.right != null)
                {
                    Q.Enqueue(q.right);
                }
            }
            sums.Add(tot);
        }
        if (sums.Count() < k)
        {
            return -1;
        }
        return sums.OrderBy(s => s).ToList()[^k];
    }
}
