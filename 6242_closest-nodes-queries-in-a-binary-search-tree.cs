/*
 * @lc app=leetcode.cn id=closest-nodes-queries-in-a-binary-search-tree lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6242] 二叉搜索树最近节点查询
 *
 * https://leetcode.cn/problems/closest-nodes-queries-in-a-binary-search-tree/description/
 *
 * algorithms
 * Medium (35.86%)
 * Total Accepted:    3.9K
 * Total Submissions: 10.9K
 * Testcase Example:  '[6,2,13,1,4,9,15,null,null,null,null,null,null,14]\n[2,5,16]'
 *
 * 给你一个 二叉搜索树 的根节点 root ，和一个由正整数组成、长度为 n 的数组 queries 。
 * 
 * 请你找出一个长度为 n 的 二维 答案数组 answer ，其中 answer[i] = [mini, maxi] ：
 * 
 * 
 * mini 是树中小于等于 queries[i] 的 最大值 。如果不存在这样的值，则使用 -1 代替。
 * maxi 是树中大于等于 queries[i] 的 最小值 。如果不存在这样的值，则使用 -1 代替。
 * 
 * 
 * 返回数组 answer 。
 * 
 * 
 * 
 * 示例 1 ：
 * 
 * 
 * 
 * 
 * 输入：root = [6,2,13,1,4,9,15,null,null,null,null,null,null,14], queries =
 * [2,5,16]
 * 输出：[[2,2],[4,6],[15,-1]]
 * 解释：按下面的描述找出并返回查询的答案：
 * - 树中小于等于 2 的最大值是 2 ，且大于等于 2 的最小值也是 2 。所以第一个查询的答案是 [2,2] 。
 * - 树中小于等于 5 的最大值是 4 ，且大于等于 5 的最小值是 6 。所以第二个查询的答案是 [4,6] 。
 * - 树中小于等于 16 的最大值是 15 ，且大于等于 16 的最小值不存在。所以第三个查询的答案是 [15,-1] 。
 * 
 * 
 * 示例 2 ：
 * 
 * 
 * 
 * 
 * 输入：root = [4,null,9], queries = [3]
 * 输出：[[-1,4]]
 * 解释：树中不存在小于等于 3 的最大值，且大于等于 3 的最小值是 4 。所以查询的答案是 [-1,4] 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树中节点的数目在范围 [2, 10^5] 内
 * 1 <= Node.val <= 10^6
 * n == queries.length
 * 1 <= n <= 10^5
 * 1 <= queries[i] <= 10^6
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
 *     private List<int> Inorder(TreeNode root)
 *     {
 *         if (root == null) { return new List<int>(); }
 *         List<int> ans = Inorder(root.left);
 *         ans.Add(root.val);
 *         foreach (var v in Inorder(root.right)) { ans.Add(v); }
 *         return ans;
 *     }
 * 
 *     private int LowerBound(List<int> list, int k)
 *     {
 *         var p = 0;
 *         var q = list.Count - 1;
 *         while (p <= q)
 *         {
 *             var mid = (p + q) >> 1;
 *             if (list[mid] >= k) { q = mid - 1; }
 *             else { p = mid + 1; }
 *         }
 *         return p;
 *     }
 * 
 *     public IList<IList<int>> ClosestNodes(TreeNode root, IList<int> queries)
 *     {
 *         var list = Inorder(root);
 *         var ans = new List<IList<int>>();
 *         foreach (var q in queries)
 *         {
 *             var minInd = LowerBound(list, q + 1) - 1;
 *             var min = minInd < 0 ? -1 : list[minInd];
 *             var maxInd = LowerBound(list, q);
 *             var max = maxInd >= list.Count ? -1 : list[maxInd];
 *             ans.Add(new List<int>{ min, max });
 *         }
 *         return ans;
 *     }
 * }
 */

// 2024.02.24: case 变强了, 更新一波
public class Solution
{
    public IList<IList<int>> ClosestNodes(TreeNode root, IList<int> queries)
    {
        var arr = new List<int>();
        void inorder(TreeNode node)
        {
            if (node is not null)
            {
                inorder(node.left);
                arr.Add(node.val);
                inorder(node.right);
            }
        }
        inorder(root);
        var n = arr.Count();
        int lowerBound(int k)
        {
            var (p, q) = (0, n);
            while (p < q)
            {
                var mid = (p + q) >> 1;
                if (arr[mid] >= k) { q = mid; }
                else { p = mid + 1; }
            }
            return p;
        }
        var k = -1;
        return queries.Select(q => new List<int>
        {
            (k = lowerBound(q + 1) - 1) < 0 ? -1 : arr[k],
            (k = lowerBound(q)) >= n ? -1 : arr[k],
        } as IList<int>).ToList();
    }
}
