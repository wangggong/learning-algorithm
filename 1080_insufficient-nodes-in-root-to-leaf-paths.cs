/*
 * @lc app=leetcode.cn id=insufficient-nodes-in-root-to-leaf-paths lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1080] 根到叶路径上的不足节点
 *
 * https://leetcode.cn/problems/insufficient-nodes-in-root-to-leaf-paths/description/
 *
 * algorithms
 * Medium (52.73%)
 * Total Accepted:    9.2K
 * Total Submissions: 16.8K
 * Testcase Example:  '[1,2,3,4,-99,-99,7,8,9,-99,-99,12,13,-99,14]\n1'
 *
 * 给你二叉树的根节点 root 和一个整数 limit ，请你同时删除树中所有 不足节点 ，并返回最终二叉树的根节点。
 * 
 * 假如通过节点 node 的每种可能的 “根-叶” 路径上值的总和全都小于给定的 limit，则该节点被称之为 不足节点 ，需要被删除。
 * 
 * 叶子节点，就是没有子节点的节点。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：root = [1,2,3,4,-99,-99,7,8,9,-99,-99,12,13,-99,14], limit = 1
 * 输出：[1,2,3,4,null,null,7,8,9,null,14]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：root = [5,4,8,11,null,17,4,7,1,null,null,5,3], limit = 22
 * 输出：[5,4,8,11,null,17,4,7,null,null,null,5]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：root = [1,2,-3,-5,null,4,null], limit = -1
 * 输出：[1,null,-3,4]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树中节点数目在范围 [1, 5000] 内
 * -10^5 <= Node.val <= 10^5
 * -10^9 <= limit <= 10^9
 * 
 * 
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

/*
 * public class Solution
 * {
 *     private int Dfs(TreeNode node, int limit, int k)
 *     {
 *         // assert node is not null;
 *         var ans = int.MinValue;
 *         if (node.left is not null)
 *         {
 *             var val = Dfs(node.left, limit, k + node.val);
 *             ans = Math.Max(ans, val);
 *             if (k + node.val + val < limit)
 *             {
 *                 node.left = null;
 *             }
 *         }
 *         if (node.right is not null)
 *         {
 *             var val = Dfs(node.right, limit, k + node.val);
 *             ans = Math.Max(ans, val);
 *             if (k + node.val + val < limit)
 *             {
 *                 node.right = null;
 *             }
 *         }
 *         return (ans == int.MinValue ? 0 : ans) + node.val;
 *     }
 * 
 *     public TreeNode SufficientSubset(TreeNode root, int limit) =>
 *         Dfs(root, limit, 0) < limit ? null : root;
 * }
 */

// 参考: https://leetcode.cn/problems/insufficient-nodes-in-root-to-leaf-paths/solution/gen-dao-xie-lu-jing-shang-de-bu-zu-jie-d-f4vz/
//
// 正难则反.
public class Solution
{
    public TreeNode SufficientSubset(TreeNode root, int limit) => HasSufficientLeaf(root, limit, 0) ? root : null;

    private bool HasSufficientLeaf(TreeNode node, int limit, int k)
    {
        if (node is null)
        {
            return false;
        }
        if (node.left is null && node.right is null)
        {
            return k + node.val >= limit;
        }
        var left = HasSufficientLeaf(node.left, limit, k + node.val);
        var right = HasSufficientLeaf(node.right, limit, k + node.val);
        node.left = left ? node.left : null;
        node.right = right ? node.right : null;
        return left || right;
    }
}
