/*
 * @lc app=leetcode.cn id=find-maximum-non-decreasing-array-length lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2945] 找到最大非递减数组的长度
 *
 * https://leetcode.cn/problems/find-maximum-non-decreasing-array-length/description/
 *
 * algorithms
 * Hard (23.10%)
 * Total Accepted:    851
 * Total Submissions: 3.6K
 * Testcase Example:  '[5,2,2]'
 *
 * 给你一个下标从 0 开始的整数数组 nums 。
 * 
 * 你可以执行任意次操作。每次操作中，你需要选择一个 子数组 ，并将这个子数组用它所包含元素的 和 替换。比方说，给定数组是 [1,3,5,6]
 * ，你可以选择子数组 [3,5] ，用子数组的和 8 替换掉子数组，然后数组会变为 [1,8,6] 。
 * 
 * 请你返回执行任意次操作以后，可以得到的 最长非递减 数组的长度。
 * 
 * 子数组 指的是一个数组中一段连续 非空 的元素序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [5,2,2]
 * 输出：1
 * 解释：这个长度为 3 的数组不是非递减的。
 * 我们有 2 种方案使数组长度为 2 。
 * 第一种，选择子数组 [2,2] ，对数组执行操作后得到 [5,4] 。
 * 第二种，选择子数组 [5,2] ，对数组执行操作后得到 [7,2] 。
 * 这两种方案中，数组最后都不是 非递减 的，所以不是可行的答案。
 * 如果我们选择子数组 [5,2,2] ，并将它替换为 [9] ，数组变成非递减的。
 * 所以答案为 1 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,2,3,4]
 * 输出：4
 * 解释：数组已经是非递减的。所以答案为 4 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [4,3,2,6]
 * 输出：3
 * 解释：将 [3,2] 替换为 [5] ，得到数组 [4,5,6] ，它是非递减的。
 * 最大可能的答案为 3 。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^5
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/find-maximum-non-decreasing-array-length/solutions/2542102/dan-diao-dui-lie-you-hua-dp-by-endlessch-j5qd/
//
// DP + 单调队列优化
public class Solution
{
    public int FindMaximumLength(int[] nums)
    {
        var n = nums.Length;
        var S = new long[n + 1];
        for (var i = 0; i < n; i++) { S[i + 1] = S[i] + (long)nums[i]; }
        var fs = new int[n + 1];
        var lasts = new long[n + 1];
        var (Q, h, t) = (new int[n + 1], 0, 1);
        for (var i = 1; i <= n; i++)
        {
            for (; t - h > 1 && S[Q[h + 1]] + lasts[Q[h + 1]] <= S[i]; h++) { }
            (fs[i], lasts[i]) = (fs[Q[h]] + 1, S[i] - S[Q[h]]);
            for (; t > h && S[Q[t - 1]] + lasts[Q[t - 1]] >= S[i] + lasts[i]; t--) { }
            Q[t++] = i;
        }
        return fs.Last();
    }
}
