/*
 * @lc app=leetcode.cn id=find-xor-beauty-of-array lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2527] 查询数组 Xor 美丽值
 *
 * https://leetcode.cn/problems/find-xor-beauty-of-array/description/
 *
 * algorithms
 * Medium (62.40%)
 * Total Accepted:    2.7K
 * Total Submissions: 4K
 * Testcase Example:  '[1,4]'
 *
 * 给你一个下标从 0 开始的整数数组 nums 。
 * 
 * 三个下标 i ，j 和 k 的 有效值 定义为 ((nums[i] | nums[j]) & nums[k]) 。
 * 
 * 一个数组的 xor 美丽值 是数组中所有满足 0 <= i, j, k < n  的三元组 (i, j, k) 的 有效值 的异或结果。
 * 
 * 请你返回 nums 的 xor 美丽值。
 * 
 * 注意：
 * 
 * 
 * val1 | val2 是 val1 和 val2 的按位或。
 * val1 & val2 是 val1 和 val2 的按位与。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,4]
 * 输出：5
 * 解释：
 * 三元组和它们对应的有效值如下：
 * - (0,0,0) 有效值为 ((1 | 1) & 1) = 1
 * - (0,0,1) 有效值为 ((1 | 1) & 4) = 0
 * - (0,1,0) 有效值为 ((1 | 4) & 1) = 1
 * - (0,1,1) 有效值为 ((1 | 4) & 4) = 4
 * - (1,0,0) 有效值为 ((4 | 1) & 1) = 1
 * - (1,0,1) 有效值为 ((4 | 1) & 4) = 4
 * - (1,1,0) 有效值为 ((4 | 4) & 1) = 0
 * - (1,1,1) 有效值为 ((4 | 4) & 4) = 4 
 * 数组的 xor 美丽值为所有有效值的按位异或 1 ^ 0 ^ 1 ^ 4 ^ 1 ^ 4 ^ 0 ^ 4 = 5 。
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [15,45,20,2,34,35,5,44,32,30]
 * 输出：34
 * 解释：数组的 xor 美丽值为 34 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^9
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     public int XorBeauty(int[] nums)
 *     {
 *         long n = (long)nums.Length;
 *         var ans = 0;
 *         for (var d = 32; d >= 0; d--)
 *         {
 *             long zeroCnt = (long)nums.Where(x => (x & (1 << d)) == 0).Count();
 *             var tot = (n - zeroCnt) * (n * n - zeroCnt * zeroCnt);
 *             if ((tot & 1) != 0)
 *             {
 *                 ans |= 1 << d;
 *             }
 *         }
 *         return ans;
 *     }
 * }
 */

// 从上面观察可知:
// 要求的是每一位的以下指标的奇偶性:
// `(n - zeroCnt) * (n * n - zeroCnt * zeroCnt) = oneCount * oneCount * (n + zeroCnt)`
// 奇偶性取决于 `n - zeroCnt` 的奇偶性, 也即 `oneCount` 的奇偶性.
// 所以直接把所有数字异或一遍即可 (异或结果也即 `oneCount` 的奇偶性, 恰好就是以上指标, 也即所求).
public class Solution
{
    public int XorBeauty(int[] nums)
        => nums.Aggregate(0, (x, y) => x ^ y);
}
