/*
 * @lc app=leetcode.cn id=maximum-value-at-a-given-index-in-a-bounded-array lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1802] 有界数组中指定下标处的最大值
 *
 * https://leetcode.cn/problems/maximum-value-at-a-given-index-in-a-bounded-array/description/
 *
 * algorithms
 * Medium (29.27%)
 * Total Accepted:    10K
 * Total Submissions: 30.1K
 * Testcase Example:  '4\n2\n6'
 *
 * 给你三个正整数 n、index 和 maxSum 。你需要构造一个同时满足下述所有条件的数组 nums（下标 从 0 开始 计数）：
 * 
 * 
 * nums.length == n
 * nums[i] 是 正整数 ，其中 0 <= i < n
 * abs(nums[i] - nums[i+1]) <= 1 ，其中 0 <= i < n-1
 * nums 中所有元素之和不超过 maxSum
 * nums[index] 的值被 最大化
 * 
 * 
 * 返回你所构造的数组中的 nums[index] 。
 * 
 * 注意：abs(x) 等于 x 的前提是 x >= 0 ；否则，abs(x) 等于 -x 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：n = 4, index = 2,  maxSum = 6
 * 输出：2
 * 解释：数组 [1,1,2,1] 和 [1,2,2,1] 满足所有条件。不存在其他在指定下标处具有更大值的有效数组。
 * 
 * 
 * 示例 2：
 * 
 * 输入：n = 6, index = 1,  maxSum = 10
 * 输出：3
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= maxSum <= 10^9
 * 0 <= index < n
 * 
 * 
 */
public class Solution
{
    public int MaxValue(int n, int index, int maxSum)
    {
        bool check(long k)
        {
            long tot = 0;
            if (index > 0)
            {
                var first = (long) Math.Max(0, k - index - 1);
                tot += (k + first - 2) * (k - first - 1) / 2;
            }
            if (index < n)
            {
                var last = (long) Math.Max(0, k - n + index);
                tot += (k + last - 1) * (k - last) / 2;
            }
            return tot <= maxSum - n;
        }
        int p = 0, q = maxSum + 10;
        while (p < q)
        {
            var mid = (p + q) >> 1;
            if (!check((long) mid))
            {
                q = mid;
            }
            else
            {
                p = mid + 1;
            }
        }
        return p - 1;
    }
}
