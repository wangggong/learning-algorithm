/*
 * @lc app=leetcode.cn id=greatest-sum-divisible-by-three lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1262] 可被三整除的最大和
 *
 * https://leetcode.cn/problems/greatest-sum-divisible-by-three/description/
 *
 * algorithms
 * Medium (52.76%)
 * Total Accepted:    24.1K
 * Total Submissions: 44.8K
 * Testcase Example:  '[3,6,5,1,8]'
 *
 * 给你一个整数数组 nums，请你找出并返回能被三整除的元素最大和。
 * 
 * 
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [3,6,5,1,8]
 * 输出：18
 * 解释：选出数字 3, 6, 1 和 8，它们的和是 18（可被 3 整除的最大和）。
 * 
 * 示例 2：
 * 
 * 输入：nums = [4]
 * 输出：0
 * 解释：4 不能被 3 整除，所以无法选出数字，返回 0。
 * 
 * 
 * 示例 3：
 * 
 * 输入：nums = [1,2,3,4,4]
 * 输出：12
 * 解释：选出数字 1, 3, 4 以及 4，它们的和是 12（可被 3 整除的最大和）。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 4 * 10^4
 * 1 <= nums[i] <= 10^4
 * 
 * 
 */
public class Solution
{
    private const int Maxn = 0x3f3f3f3f;

    public int MaxSumDivThree(int[] nums)
    {
        var sum = nums.Sum();
        var (p, q, s, t) = (Maxn, Maxn, Maxn, Maxn);
        foreach (var n in nums)
        {
            switch (n % 3)
            {
                case 1:
                    if (n < p)
                    {
                        (p, q) = (n, p);
                    }
                    else if (n < q)
                    {
                        q = n;
                    }
                    break;
                case 2:
                    if (n < s)
                    {
                        (s, t) = (n, s);
                    }
                    else if (n < t)
                    {
                        t = n;
                    }
                    break;
                default:
                    break;
            }
        }
        return sum - (sum % 3 is 0
            ? 0
            : (sum % 3 is 1
                ? Math.Min(s + t, p)
                : Math.Min(p + q, s)));

    }
}
