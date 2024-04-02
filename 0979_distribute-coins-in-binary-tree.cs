/*
 * @lc app=leetcode.cn id=distribute-coins-in-binary-tree lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [979] 在二叉树中分配硬币
 *
 * https://leetcode.cn/problems/distribute-coins-in-binary-tree/description/
 *
 * algorithms
 * Medium (72.42%)
 * Total Accepted:    14K
 * Total Submissions: 19.2K
 * Testcase Example:  '[3,0,0]'
 *
 * 给定一个有 N 个结点的二叉树的根结点 root，树中的每个结点上都对应有 node.val 枚硬币，并且总共有 N 枚硬币。
 * 
 * 在一次移动中，我们可以选择两个相邻的结点，然后将一枚硬币从其中一个结点移动到另一个结点。(移动可以是从父结点到子结点，或者从子结点移动到父结点。)。
 * 
 * 返回使每个结点上只有一枚硬币所需的移动次数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：[3,0,0]
 * 输出：2
 * 解释：从树的根结点开始，我们将一枚硬币移到它的左子结点上，一枚硬币移到它的右子结点上。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 输入：[0,3,0]
 * 输出：3
 * 解释：从根结点的左子结点开始，我们将两枚硬币移到根结点上 [移动两次]。然后，我们把一枚硬币从根结点移到右子结点上。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 
 * 输入：[1,0,2]
 * 输出：2
 * 
 * 
 * 示例 4：
 * 
 * 
 * 
 * 输入：[1,0,0,null,3]
 * 输出：4
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1<= N <= 100
 * 0 <= node.val <= N
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

// 参考: https://leetcode.cn/problems/distribute-coins-in-binary-tree/solution/zai-er-cha-shu-zhong-fen-pei-ying-bi-by-e4poq/
//
// 二叉树, 确实想不到. 学习了.
public class Solution
{
    public int DistributeCoins(TreeNode root)
    {
        var ans = 0;
        int dfs(TreeNode node)
        {
            if (node is null) { return 0; }
            var (left, right) = (dfs(node.left), dfs(node.right));
            ans += Math.Abs(left) + Math.Abs(right);
            return node.val + left + right - 1;
        }
        dfs(root);
        return ans;
    }
}
