/*
 * @lc app=leetcode.cn id=the-number-of-beautiful-subsets lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6352] 美丽子集的数目
 *
 * https://leetcode.cn/problems/the-number-of-beautiful-subsets/description/
 *
 * algorithms
 * Medium (25.46%)
 * Total Accepted:    2.6K
 * Total Submissions: 10K
 * Testcase Example:  '[2,4,6]\n2'
 *
 * 给你一个由正整数组成的数组 nums 和一个 正 整数 k 。
 * 
 * 如果 nums 的子集中，任意两个整数的绝对差均不等于 k ，则认为该子数组是一个 美丽 子集。
 * 
 * 返回数组 nums 中 非空 且 美丽 的子集数目。
 * 
 * nums 的子集定义为：可以经由 nums
 * 删除某些元素（也可能不删除）得到的一个数组。只有在删除元素时选择的索引不同的情况下，两个子集才会被视作是不同的子集。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [2,4,6], k = 2
 * 输出：4
 * 解释：数组 nums 中的美丽子集有：[2], [4], [6], [2, 6] 。
 * 可以证明数组 [2,4,6] 中只存在 4 个美丽子集。
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [1], k = 1
 * 输出：1
 * 解释：数组 nums 中的美丽数组有：[1] 。
 * 可以证明数组 [1] 中只存在 1 个美丽子集。 
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 20
 * 1 <= nums[i], k <= 1000
 * 
 * 
 */
public class Solution
{
    public int BeautifulSubsets(int[] nums, int k)
    {
        var n = nums.Length;
        var ans = 0;
        var edges = new int[n];
        for (var i = 0; i < n; i++)
        {
            for (var j = i + 1; j < n; j++)
            {
                if (Math.Abs(nums[i] - nums[j]) == k)
                {
                    edges[i] |= 1 << j;
                    edges[j] |= 1 << i;
                }
            }
        }
        void traversal(int k, int mask)
        {
            if (k == n)
            {
                ans++;
                return;
            }
            if ((mask & edges[k]) == 0)
            {
                traversal(k + 1, mask | (1 << k));
            }
            traversal(k + 1, mask);
        }
        traversal(0, 0);
        return ans - 1;
    }
}
