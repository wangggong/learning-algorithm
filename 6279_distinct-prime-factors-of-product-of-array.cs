/*
 * @lc app=leetcode.cn id=distinct-prime-factors-of-product-of-array lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6279] 数组乘积中的不同质因数数目
 *
 * https://leetcode.cn/problems/distinct-prime-factors-of-product-of-array/description/
 *
 * algorithms
 * Medium (57.61%)
 * Total Accepted:    2.9K
 * Total Submissions: 5K
 * Testcase Example:  '[2,4,3,7,10,6]'
 *
 * 给你一个正整数数组 nums ，对 nums 所有元素求积之后，找出并返回乘积中 不同质因数 的数目。
 * 
 * 注意：
 * 
 * 
 * 质数 是指大于 1 且仅能被 1 及自身整除的数字。
 * 如果 val2 / val1 是一个整数，则整数 val1 是另一个整数 val2 的一个因数。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [2,4,3,7,10,6]
 * 输出：4
 * 解释：
 * nums 中所有元素的乘积是：2 * 4 * 3 * 7 * 10 * 6 = 10080 = 2^5 * 3^2 * 5 * 7 。
 * 共有 4 个不同的质因数，所以返回 4 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [2,4,8,16]
 * 输出：1
 * 解释：
 * nums 中所有元素的乘积是：2 * 4 * 8 * 16 = 1024 = 2^10 。
 * 共有 1 个不同的质因数，所以返回 1 。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^4
 * 2 <= nums[i] <= 1000
 * 
 * 
 */
public class Solution
{
    public int DistinctPrimeFactors(int[] nums)
    {
        var S = new HashSet<int>();
        foreach (var num in nums)
        {
            var n = num;
            for (var k = 2; k <= n / k; k++)
            {
                while (n % k == 0)
                {
                    S.Add(k);
                    n /= k;
                }
            }
            if (n > 1)
            {
                S.Add(n);
            }
        }
        return S.Count();
    }
}
