/*
 * @lc app=leetcode.cn id=shortest-subarray-with-or-at-least-k-ii lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100271] 或值至少为 K 的最短子数组 II
 *
 * https://leetcode.cn/problems/shortest-subarray-with-or-at-least-k-ii/description/
 *
 * algorithms
 * Medium (31.34%)
 * Total Accepted:    1.6K
 * Total Submissions: 5K
 * Testcase Example:  '[1,2,3]\n2'
 *
 * 给你一个 非负 整数数组 nums 和一个整数 k 。
 * 
 * 如果一个数组中所有元素的按位或运算 OR 的值 至少 为 k ，那么我们称这个数组是 特别的 。
 * 
 * 请你返回 nums 中 最短特别非空 子数组的长度，如果特别子数组不存在，那么返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,3], k = 2
 * 
 * 输出：1
 * 
 * 解释：
 * 
 * 子数组 [3] 的按位 OR 值为 3 ，所以我们返回 1 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [2,1,8], k = 10
 * 
 * 输出：3
 * 
 * 解释：
 * 
 * 子数组 [2,1,8] 的按位 OR 值为 11 ，所以我们返回 3 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [1,2], k = 0
 * 
 * 输出：1
 * 
 * 解释：
 * 
 * 子数组 [1] 的按位 OR 值为 1 ，所以我们返回 1 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 2 * 10^5
 * 0 <= nums[i] <= 10^9
 * 0 <= k <= 10^9
 * 
 * 
 */

// 基础不牢, 地动山摇
public class Solution
{
    public int MinimumSubarrayLength(int[] nums, int k)
    {
        const int B = 30;
        if (k == 0)
        {
            return 1;
        }
        var n = nums.Length;
        var ans = n + 1;
        var counts = new int[B];
        int getValue() => Enumerable
            .Range(0, B)
            .Select(b => counts[b] > 0 ? (1 << b) : 0)
            .Sum();
        void update(int n, int v)
        {
            for (var b = 0; b < B; b++)
            {
                if (((n >> b) & 1) != 0)
                {
                    counts[b] += v;
                }
            }
        }
        for (var (p, q) = (0, 0); q < n; q++)
        {
            update(nums[q], 1);
            if (getValue() >= k)
            {
                for (; getValue() >= k; p++)
                {
                    update(nums[p], -1);
                }
                ans = Math.Min(ans, q - p + 2);
            }
        }
        return ans > n
            ? -1
            : ans;
    }
}
