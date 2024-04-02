/*
 * @lc app=leetcode.cn id=number-of-subarrays-with-bounded-maximum lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [795] 区间子数组个数
 *
 * https://leetcode.cn/problems/number-of-subarrays-with-bounded-maximum/description/
 *
 * algorithms
 * Medium (53.08%)
 * Total Accepted:    32.3K
 * Total Submissions: 55.9K
 * Testcase Example:  '[2,1,4,3]\n2\n3'
 *
 * 给你一个整数数组 nums 和两个整数：left 及 right 。找出 nums 中连续、非空且其中最大元素在范围 [left, right]
 * 内的子数组，并返回满足条件的子数组的个数。
 * 
 * 生成的测试用例保证结果符合 32-bit 整数范围。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [2,1,4,3], left = 2, right = 3
 * 输出：3
 * 解释：满足条件的三个子数组：[2], [2, 1], [3]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [2,9,2,5,6], left = 2, right = 8
 * 输出：7
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * 0 <= nums[i] <= 10^9
 * 0 <= left <= right <= 10^9
 * 
 * 
 */
public class Solution
{
    public int NumSubarrayBoundedMax(int[] nums, int l, int r)
    {
        var n = nums.Length;
        var left = new int[n];
        var right = new int[n];
        for (int i = 0; i < n; i++)
        {
            left[i] = -1;
            right[i] = n;
        }
        var S = new Stack<int>();
        for (int i = 0; i < n; i++)
        {
            while (S.Count > 0 && nums[S.Peek()] <= nums[i]) { S.Pop(); }
            if (S.Count > 0) { left[i] = S.Peek(); }
            S.Push(i);
        }
        S = new Stack<int>();
        for (int i = n - 1; i >= 0; i--)
        {
            while (S.Count > 0 && nums[S.Peek()] < nums[i]) { S.Pop(); }
            if (S.Count > 0) { right[i] = S.Peek(); }
            S.Push(i);
        }
        int ans = 0;
        for (int i = 0; i < n; i++)
        {
            if (l <= nums[i] && nums[i] <= r) { ans += (i - left[i]) * (right[i] - i); }
        }
        return ans;
    }
}
