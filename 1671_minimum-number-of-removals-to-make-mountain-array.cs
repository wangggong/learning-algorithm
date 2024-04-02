/*
 * @lc app=leetcode.cn id=minimum-number-of-removals-to-make-mountain-array lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1671] 得到山形数组的最少删除次数
 *
 * https://leetcode.cn/problems/minimum-number-of-removals-to-make-mountain-array/description/
 *
 * algorithms
 * Hard (46.00%)
 * Total Accepted:    4.9K
 * Total Submissions: 10.6K
 * Testcase Example:  '[1,3,1]'
 *
 * 我们定义 arr 是 山形数组 当且仅当它满足：
 * 
 * 
 * arr.length >= 3
 * 存在某个下标 i （从 0 开始） 满足 0 < i < arr.length - 1 且：
 * 
 * arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
 * arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
 * 
 * 
 * 
 * 
 * 给你整数数组 nums​ ，请你返回将 nums 变成 山形状数组 的​ 最少 删除次数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,3,1]
 * 输出：0
 * 解释：数组本身就是山形数组，所以我们不需要删除任何元素。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [2,1,1,5,6,2,3,1]
 * 输出：3
 * 解释：一种方法是将下标为 0，1 和 5 的元素删除，剩余元素为 [1,5,6,3,1] ，是山形数组。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 3 <= nums.length <= 1000
 * 1 <= nums[i] <= 10^9
 * 题目保证 nums 删除一些元素后一定能得到山形数组。
 * 
 * 
 */

/*
 * // 状态机 DP
 * public class Solution
 * {
 *     public int MinimumMountainRemovals(int[] nums)
 *     {
 *         const int Maxn = 0x3f3f3f3f;
 *         var n = nums.Length;
 *         var dp = new int[n][];
 *         for (var i = 0; i < n; i++) { dp[i] = new int[] { 1, -Maxn, }; }
 *         for (var i = 1; i < n; i++)
 *         {
 *             for (var j = 0; j <= i; j++)
 *             {
 *                 if (nums[j] < nums[i])
 *                 { dp[i][0] = Math.Max(dp[i][0], dp[j][0] + 1); }
 *                 else if (nums[j] > nums[i])
 *                 {
 *                     dp[i][1] = Math.Max(dp[i][1], dp[j][1] + 1);
 *                     if (dp[j][0] > 1)
 *                     { dp[i][1] = Math.Max(dp[i][1], dp[j][0] + 1); }
 *                 }
 *             }
 *         }
 *         return dp[2..]
 *             .Select(d => n - d[1])
 *             .Min();
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/minimum-number-of-removals-to-make-mountain-array/solutions/2570598/zui-chang-di-zeng-zi-xu-lie-by-leetcode-2ipno/?envType=daily-question&envId=2023-12-22
//
// LIS 应用.
public class Solution
{
    private const int Maxn = 0x3f3f3f3f;

    private int[] Lis(int[] arr)
    {
        var n = arr.Length;
        var dp = new int[n];
        Array.Fill(dp, 1);
        for (var i = 0; i < n; i++)
        {
            for (var j = 0; j < i; j++)
            {
                if (arr[j] < arr[i]) { dp[i] = Math.Max(dp[i], dp[j] + 1); }
            }
        }
        return dp;
    }

    public int MinimumMountainRemovals(int[] nums) => Lis(nums)
        .Zip(Lis(nums.Reverse()
                    .ToArray())
                .Reverse(),
            (p, s) => p > 1 && s > 1
                ? nums.Length - (p + s - 1)
                : Maxn)
        .Min();
}
