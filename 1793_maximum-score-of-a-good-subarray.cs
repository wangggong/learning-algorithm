/*
 * @lc app=leetcode.cn id=maximum-score-of-a-good-subarray lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1793] 好子数组的最大分数
 *
 * https://leetcode.cn/problems/maximum-score-of-a-good-subarray/description/
 *
 * algorithms
 * Hard (46.83%)
 * Total Accepted:    6.7K
 * Total Submissions: 14.2K
 * Testcase Example:  '[1,4,3,7,4,5]\n3'
 *
 * 给你一个整数数组 nums （下标从 0 开始）和一个整数 k 。
 * 
 * 一个子数组 (i, j) 的 分数 定义为 min(nums[i], nums[i+1], ..., nums[j]) * (j - i + 1)
 * 。一个 好 子数组的两个端点下标需要满足 i <= k <= j 。
 * 
 * 请你返回 好 子数组的最大可能 分数 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [1,4,3,7,4,5], k = 3
 * 输出：15
 * 解释：最优子数组的左右端点下标是 (1, 5) ，分数为 min(4,3,7,4,5) * (5-1+1) = 3 * 5 = 15 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [5,5,4,5,4,1,1,1], k = 0
 * 输出：20
 * 解释：最优子数组的左右端点下标是 (0, 4) ，分数为 min(5,5,4,5,4) * (4-0+1) = 4 * 5 = 20 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 2 * 10^4
 * 0 <= k < nums.length
 * 
 * 
 */
public class Solution
{
    public int MaximumScore(int[] nums, int k)
    {
        var n = nums.Length;
        var lefts = new int[n];
        var rights = new int[n];
        var S = new Stack<int>();
        S.Push(-1);
        for (var i = 0; i < n; i++)
        {
            for (; S.Peek() != -1 && nums[S.Peek()] > nums[i]; S.Pop()) { }
            lefts[i] = S.Peek();
            S.Push(i);
        }
        S = new Stack<int>();
        S.Push(n);
        for (var i = n - 1; i >= 0; i--)
        {
            for (; S.Peek() != n && nums[S.Peek()] >= nums[i]; S.Pop()) { }
            rights[i] = S.Peek();
            S.Push(i);
        }
        var ans = 0;
        for (var i = 0; i < n; i++)
        {
            if (lefts[i] < k && k < rights[i])
            {
                ans = Math.Max(ans, nums[i] * (rights[i] - lefts[i] - 1));
            }
        }
        return ans;
    }
}
