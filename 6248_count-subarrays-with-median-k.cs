/*
 * @lc app=leetcode.cn id=count-subarrays-with-median-k lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6248] 统计中位数为 K 的子数组
 *
 * https://leetcode.cn/problems/count-subarrays-with-median-k/description/
 *
 * algorithms
 * Hard (26.84%)
 * Total Accepted:    2.2K
 * Total Submissions: 6.9K
 * Testcase Example:  '[3,2,1,4,5]\n4'
 *
 * 给你一个长度为 n 的数组 nums ，该数组由从 1 到 n 的 不同 整数组成。另给你一个正整数 k 。
 * 
 * 统计并返回 num 中的 中位数 等于 k 的非空子数组的数目。
 * 
 * 注意：
 * 
 * 
 * 数组的中位数是按 递增 顺序排列后位于 中间 的那个元素，如果数组长度为偶数，则中位数是位于中间靠 左 的那个元素。
 * 
 * 
 * 例如，[2,3,1,4] 的中位数是 2 ，[8,4,3,5,1] 的中位数是 4 。
 * 
 * 
 * 子数组是数组中的一个连续部分。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [3,2,1,4,5], k = 4
 * 输出：3
 * 解释：中位数等于 4 的子数组有：[4]、[4,5] 和 [1,4,5] 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [2,3,1], k = 3
 * 输出：1
 * 解释：[3] 是唯一一个中位数等于 3 的子数组。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == nums.length
 * 1 <= n <= 10^5
 * 1 <= nums[i], k <= n
 * nums 中的整数互不相同
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/count-subarrays-with-median-k/solution/deng-jie-zhuan-huan-pythonjavacgo-by-end-5w11/
public class Solution
{
    public int CountSubarrays(int[] nums, int k)
    {
        var n = nums.Length;
        var index = Array.IndexOf(nums, k);
        var d = new Dictionary<int, int>();
        d[0] = 1;
        for (int i = index + 1, cur = 0; i < n; i++)
        {
            cur += nums[i] > k ? 1 : -1;
            d[cur] = (d.ContainsKey(cur) ? d[cur] : 0) + 1;
        }
        var ans = d[0] + (d.ContainsKey(1) ? d[1] : 0);
        for (int i = index - 1, cur = 0; i >= 0; i--)
        {
            cur += nums[i] > k ? 1 : -1;
            ans += d.ContainsKey(-cur) ? d[-cur] : 0;
            ans += d.ContainsKey(1 - cur) ? d[1 - cur] : 0;
        }
        return ans;
    }
}
