/*
 * @lc app=leetcode.cn id=minimize-the-maximum-difference-of-pairs lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6359] 最小化数对的最大差值
 *
 * https://leetcode.cn/problems/minimize-the-maximum-difference-of-pairs/description/
 *
 * algorithms
 * Medium (32.15%)
 * Total Accepted:    2K
 * Total Submissions: 6.4K
 * Testcase Example:  '[10,1,2,7,1,3]\n2'
 *
 * 给你一个下标从 0 开始的整数数组 nums 和一个整数 p 。请你从 nums 中找到 p 个下标对，每个下标对对应数值取差值，你需要使得这 p
 * 个差值的 最大值 最小。同时，你需要确保每个下标在这 p 个下标对中最多出现一次。
 * 
 * 对于一个下标对 i 和 j ，这一对的差值为 |nums[i] - nums[j]| ，其中 |x| 表示 x 的 绝对值 。
 * 
 * 请你返回 p 个下标对对应数值 最大差值 的 最小值 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [10,1,2,7,1,3], p = 2
 * 输出：1
 * 解释：第一个下标对选择 1 和 4 ，第二个下标对选择 2 和 5 。
 * 最大差值为 max(|nums[1] - nums[4]|, |nums[2] - nums[5]|) = max(0, 1) = 1 。所以我们返回
 * 1 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [4,2,1,2], p = 1
 * 输出：0
 * 解释：选择下标 1 和 3 构成下标对。差值为 |2 - 2| = 0 ，这是最大差值的最小值。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * 0 <= nums[i] <= 10^9
 * 0 <= p <= (nums.length)/2
 * 
 * 
 */
public class Solution
{
    public int MinimizeMax(int[] nums, int m)
    {
        Array.Sort(nums);
        var n = nums.Length;
        if (m == 0)
        {
            return 0;
        }
        var D = new int[n - 1];
        for (var i = 1; i < n; i++)
        {
            D[i - 1] = nums[i] - nums[i - 1];
        }
        // 1 1 2 3 7 10
        //   0 1 1 4 3
        bool check(int k)
        {
            var left = m;
            for (var i = 0; i < n - 1; i++)
            {
                if (D[i] <= k)
                {
                    left--;
                    i++;
                }
            }
            return left <= 0;
        }
        var (p, q) = (0, D.Max());
        while (p < q)
        {
            var mid = (p + q) >> 1;
            if (check(mid))
            {
                q = mid;
            }
            else
            {
                p = mid + 1;
            }
        }
        return p;
    }
}
