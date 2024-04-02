/*
 * @lc app=leetcode.cn id=maximum-subarray-sum-with-one-deletion lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1186] 删除一次得到子数组最大和
 *
 * https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/description/
 *
 * algorithms
 * Medium (42.67%)
 * Total Accepted:    14.3K
 * Total Submissions: 32.2K
 * Testcase Example:  '[1,-2,0,3]'
 *
 * 给你一个整数数组，返回它的某个 非空
 * 子数组（连续元素）在执行一次可选的删除操作后，所能得到的最大元素总和。换句话说，你可以从原数组中选出一个子数组，并可以决定要不要从中删除一个元素（只能删一次哦），（删除后）子数组中至少应当有一个元素，然后该子数组（剩下）的元素总和是所有子数组之中最大的。
 * 
 * 注意，删除一个元素后，子数组 不能为空。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：arr = [1,-2,0,3]
 * 输出：4
 * 解释：我们可以选出 [1, -2, 0, 3]，然后删掉 -2，这样得到 [1, 0, 3]，和最大。
 * 
 * 示例 2：
 * 
 * 
 * 输入：arr = [1,-2,-2,3]
 * 输出：3
 * 解释：我们直接选出 [3]，这就是最大和。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：arr = [-1,-1,-1,-1]
 * 输出：-1
 * 解释：最后得到的子数组不能为空，所以我们不能选择 [-1] 并从中删去 -1 来得到 0。
 * ⁠    我们应该直接选择 [-1]，或者选择 [-1, -1] 再从中删去一个 -1。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 
 * 1 <= arr.length <= 10^5
 * -10^4 <= arr[i] <= 10^4
 * 
 * 
 */


/*
 * public class Solution
 * {
 *     public int MaximumSum(int[] arr)
 *     {
 *         var n = arr.Length;
 *         var dp = Enumerable.Range(0, n)
 *             .Select(_ => new int[2])
 *             .ToArray();
 *         (dp[0][0], dp[0][1]) = (arr[0], int.MinValue);
 *         for (var i = 1; i < n; i++)
 *         {
 *             dp[i][0] = Math.Max(dp[i - 1][0], 0) + arr[i];
 *             dp[i][1] = Math.Max(dp[i - 1][1], (i >= 2 ? dp[i - 2][0] : 0)) + arr[i];
 *         }
 *         return dp.SelectMany(row => row)
 *             .Max();
 *     }
 * }
 */

// 官解提供了更合理的状态定义.
public class Solution
{
    public int MaximumSum(int[] arr)
    {
        var dp = new int[] { arr[0], 0, };
        var ans = arr[0];
        foreach (var a in arr.Skip(1))
        {
            dp[1] = Math.Max(dp[1] + a, dp[0]);
            dp[0] = Math.Max(dp[0], 0) + a;
            ans = Math.Max(ans, dp.Max());
        }
        return ans;
    }
}
