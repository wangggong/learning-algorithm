/*
 * @lc app=leetcode.cn id=minimum-cost-tree-from-leaf-values lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1130] 叶值的最小代价生成树
 *
 * https://leetcode.cn/problems/minimum-cost-tree-from-leaf-values/description/
 *
 * algorithms
 * Medium (64.95%)
 * Total Accepted:    13.1K
 * Total Submissions: 19.2K
 * Testcase Example:  '[6,2,4]'
 *
 * 给你一个正整数数组 arr，考虑所有满足以下条件的二叉树：
 * 
 * 
 * 每个节点都有 0 个或是 2 个子节点。
 * 数组 arr 中的值与树的中序遍历中每个叶节点的值一一对应。
 * 每个非叶节点的值等于其左子树和右子树中叶节点的最大值的乘积。
 * 
 * 
 * 在所有这样的二叉树中，返回每个非叶节点的值的最小可能总和。这个和的值是一个 32 位整数。
 * 
 * 如果一个节点有 0 个子节点，那么该节点为叶节点。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：arr = [6,2,4]
 * 输出：32
 * 解释：有两种可能的树，第一种的非叶节点的总和为 36 ，第二种非叶节点的总和为 32 。 
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：arr = [4,11]
 * 输出：44
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= arr.length <= 40
 * 1 <= arr[i] <= 15
 * 答案保证是一个 32 位带符号整数，即小于 2^31 。
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     public int MctFromLeafValues(int[] arr)
 *     {
 *         var memo = new Dictionary<(int, int), (int, int)>();
 *         (int, int) dfs(int p, int q)
 *         {
 *             var key = (p, q);
 *             if (memo.ContainsKey(key))
 *             {
 *                 return memo[key];
 *             }
 *             var (max, sum) = (arr[p], int.MaxValue);
 *             for (var i = p; i < q; i++)
 *             {
 *                 var (leftMax, leftSum) = dfs(p, i);
 *                 var (rightMax, rightSum) = dfs(i + 1, q);
 *                 max = Math.Max(max, Math.Max(leftMax, rightMax));
 *                 sum = Math.Min(sum, leftMax * rightMax + leftSum + rightSum);
 *             }
 *             return memo[key] = (max, sum == int.MaxValue ? arr[p] : sum);
 *         }
 *         return dfs(0, arr.Length - 1).Item2 - arr.Sum();
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/minimum-cost-tree-from-leaf-values/solution/xie-zhi-de-zui-xiao-dai-jie-sheng-cheng-26ozf/
//
// 单调栈
public class Solution
{
    public int MctFromLeafValues(int[] arr)
    {
        var S = new Stack<int>();
        var ans = 0;
        foreach (var a in arr)
        {
            while (S.Count > 0 && S.Peek() <= a)
            {
                var v = S.Pop();
                ans += (S.Count == 0 || S.Peek() >= a ? a : S.Peek()) * v;
            }
            S.Push(a);
        }
        while (S.Count > 1)
        {
            var v = S.Pop();
            ans += v * S.Peek();
        }
        return ans;
    }
}
