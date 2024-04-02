/*
 * @lc app=leetcode.cn id=all-possible-full-binary-trees lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [894] 所有可能的真二叉树
 *
 * https://leetcode.cn/problems/all-possible-full-binary-trees/description/
 *
 * algorithms
 * Medium (77.72%)
 * Total Accepted:    27.5K
 * Total Submissions: 34.6K
 * Testcase Example:  '7'
 *
 * 给你一个整数 n ，请你找出所有可能含 n 个节点的 真二叉树 ，并以列表形式返回。答案中每棵树的每个节点都必须符合 Node.val == 0 。
 * 
 * 答案的每个元素都是一棵真二叉树的根节点。你可以按 任意顺序 返回最终的真二叉树列表。
 * 
 * 真二叉树 是一类二叉树，树中每个节点恰好有 0 或 2 个子节点。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 7
 * 
 * 输出：[[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 3
 * 输出：[[0,0,0]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 20
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
    public IList<TreeNode> AllPossibleFBT(int n)
    {
        if (n % 2 == 0)
        {
            return new List<TreeNode>();
        }
        if (n is 1)
        {
            return new List<TreeNode> { new(), };
        }
        var ans = new List<TreeNode>();
        for (var i = 0; i + 1 < n; i++)
        {
            foreach (var l in AllPossibleFBT(i))
            {
                foreach (var r in AllPossibleFBT(n - 1 - i))
                {
                    ans.Add(new(0, l, r));
                }
            }
        }
        return ans;
    }
}
