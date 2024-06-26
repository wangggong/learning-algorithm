/*
 * @lc app=leetcode.cn id=cousins-in-binary-tree-ii lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6335] 二叉树的堂兄弟节点 II
 *
 * https://leetcode.cn/problems/cousins-in-binary-tree-ii/description/
 *
 * algorithms
 * Medium (65.89%)
 * Total Accepted:    2K
 * Total Submissions: 3.1K
 * Testcase Example:  '[5,4,9,1,10,null,7]'
 *
 * 给你一棵二叉树的根 root ，请你将每个节点的值替换成该节点的所有 堂兄弟节点值的和 。
 * 
 * 如果两个节点在树中有相同的深度且它们的父节点不同，那么它们互为 堂兄弟 。
 * 
 * 请你返回修改值之后，树的根 root 。
 * 
 * 注意，一个节点的深度指的是从树根节点到这个节点经过的边数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：root = [5,4,9,1,10,null,7]
 * 输出：[0,0,0,7,7,null,11]
 * 解释：上图展示了初始的二叉树和修改每个节点的值之后的二叉树。
 * - 值为 5 的节点没有堂兄弟，所以值修改为 0 。
 * - 值为 4 的节点没有堂兄弟，所以值修改为 0 。
 * - 值为 9 的节点没有堂兄弟，所以值修改为 0 。
 * - 值为 1 的节点有一个堂兄弟，值为 7 ，所以值修改为 7 。
 * - 值为 10 的节点有一个堂兄弟，值为 7 ，所以值修改为 7 。
 * - 值为 7 的节点有两个堂兄弟，值分别为 1 和 10 ，所以值修改为 11 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：root = [3,1,2]
 * 输出：[0,0,0]
 * 解释：上图展示了初始的二叉树和修改每个节点的值之后的二叉树。
 * - 值为 3 的节点没有堂兄弟，所以值修改为 0 。
 * - 值为 1 的节点没有堂兄弟，所以值修改为 0 。
 * - 值为 2 的节点没有堂兄弟，所以值修改为 0 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树中节点数目的范围是 [1, 10^5] 。
 * 1 <= Node.val <= 10^4
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
    public TreeNode ReplaceValueInTree(TreeNode root)
    {
        var dSum = new Dictionary<TreeNode, int>();
        var levelSums = new List<int>();
        var dummy = new TreeNode();
        void bfs(Action<TreeNode, TreeNode, int> f)
        {
            
            var Q = new Queue<(TreeNode, TreeNode)>();
            Q.Enqueue((root, dummy));
            for (var i = 0; Q.Count > 0; i++)
            {
                if (levelSums.Count() <= i)
                {
                    levelSums.Add(0);
                }
                for (var c = Q.Count; c > 0; c--)
                {
                    var (node, pa) = Q.Dequeue();
                    f(node, pa, i);
                    if (node.left is not null)
                    {
                        Q.Enqueue((node.left, node));
                    }
                    if (node.right is not null)
                    {
                        Q.Enqueue((node.right, node));
                    }
                }
            }
        }
        bfs((node, pa, k) =>
        {
            levelSums[k] += node.val;
            dSum[pa] = (dSum.ContainsKey(pa) ? dSum[pa] : 0) + node.val;
        });
        bfs((node, pa, k) => node.val = levelSums[k] - dSum[pa]);
        return root;
    }
}

// 两次 DFS, 更简洁些.

/*
 * public class Solution
 * {
 *     public TreeNode ReplaceValueInTree(TreeNode root)
 *     {
 *         var sums = new List<int>();
 *         var d = new Dictionary<TreeNode, int>();
 *         var dummy = new TreeNode();
 *         void dfs(TreeNode n, TreeNode p, int depth, Action<TreeNode, TreeNode, int> f)
 *         {
 *             if (n is null) { return; }
 *             f(n, p, depth);
 *             dfs(n.left, n, depth + 1, f);
 *             dfs(n.right, n, depth + 1, f);
 *         }
 *         dfs(root, dummy, 0, (n, p, depth) =>
 *         {
 *             if (sums.Count <= depth) { sums.Add(0); }
 *             sums[depth] += n.val;
 *             d.TryGetValue(p, out var s);
 *             d[p] = s + n.val;
 *         });
 *         dfs(root, dummy, 0, (n, p, depth) => n.val = sums[depth] - d[p]);
 *         return root;
 *     }
 * }
 */
