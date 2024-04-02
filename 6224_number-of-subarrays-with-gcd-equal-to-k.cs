/*
 * @lc app=leetcode.cn id=number-of-subarrays-with-gcd-equal-to-k lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6224] 最大公因数等于 K 的子数组数目
 *
 * https://leetcode.cn/problems/number-of-subarrays-with-gcd-equal-to-k/description/
 *
 * algorithms
 * Medium (32.40%)
 * Total Accepted:    4K
 * Total Submissions: 12.4K
 * Testcase Example:  '[9,3,1,2,6,3]\n3'
 *
 * 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 nums 的子数组中元素的最大公因数等于 k 的子数组数目。
 * 
 * 子数组 是数组中一个连续的非空序列。
 * 
 * 数组的最大公因数 是能整除数组中所有元素的最大整数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [9,3,1,2,6,3], k = 3
 * 输出：4
 * 解释：nums 的子数组中，以 3 作为最大公因数的子数组如下：
 * - [9,3,1,2,6,3]
 * - [9,3,1,2,6,3]
 * - [9,3,1,2,6,3]
 * - [9,3,1,2,6,3]
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [4], k = 7
 * 输出：0
 * 解释：不存在以 7 作为最大公因数的子数组。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 1000
 * 1 <= nums[i], k <= 10^9
 * 
 * 
 */
public class Solution
{
    public int SubarrayGCD(int[] nums, int k)
    {
        int ans = 0;
        for (int p = 0, n = nums.Length; p < n; p++)
        {
            int cur = nums[p];
            for (int q = p; q < n; q++)
            {
                cur = Gcd(cur, nums[q]);
                if (cur == k)
                {
                    ans++;
                }
            }
        }
        return ans;
    }

    private int Gcd(int x, int y) { return y == 0 ? x : Gcd(y, x % y); }
}
