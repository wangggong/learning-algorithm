/*
 * @lc app=leetcode.cn id=divide-array-into-arrays-with-max-difference lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100161] 划分数组并满足最大差限制
 *
 * https://leetcode.cn/problems/divide-array-into-arrays-with-max-difference/description/
 *
 * algorithms
 * Medium (63.77%)
 * Total Accepted:    3.2K
 * Total Submissions: 5.1K
 * Testcase Example:  '[1,3,4,8,7,9,3,5,1]\n2'
 *
 * 给你一个长度为 n 的整数数组 nums，以及一个正整数 k 。
 * 
 * 将这个数组划分为一个或多个长度为 3 的子数组，并满足以下条件：
 * 
 * 
 * nums 中的 每个 元素都必须 恰好 存在于某个子数组中。
 * 子数组中 任意 两个元素的差必须小于或等于 k 。
 * 
 * 
 * 返回一个 二维数组 ，包含所有的子数组。如果不可能满足条件，就返回一个空数组。如果有多个答案，返回 任意一个 即可。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,3,4,8,7,9,3,5,1], k = 2
 * 输出：[[1,1,3],[3,4,5],[7,8,9]]
 * 解释：可以将数组划分为以下子数组：[1,1,3]，[3,4,5] 和 [7,8,9] 。
 * 每个子数组中任意两个元素的差都小于或等于 2 。
 * 注意，元素的顺序并不重要。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,3,3,2,7,3], k = 3
 * 输出：[]
 * 解释：无法划分数组满足所有条件。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == nums.length
 * 1 <= n <= 10^5
 * n 是 3 的倍数
 * 1 <= nums[i] <= 10^5
 * 1 <= k <= 10^5
 * 
 * 
 */
public class Solution
{
    public int[][] DivideArray(int[] nums, int k)
    {
        Array.Sort(nums);
        var ans = new List<int[]>();
        for (var (i, n) = (0, nums.Length); i < n; i += 3)
        {
            if (nums[i + 2] - nums[i] > k) { return new int[0][]; }
            ans.Add(nums[i..(i + 3)]);
        }
        return ans.ToArray();
    }
}
