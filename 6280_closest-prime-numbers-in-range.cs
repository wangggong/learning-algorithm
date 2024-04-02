/*
 * @lc app=leetcode.cn id=closest-prime-numbers-in-range lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6280] 范围内最接近的两个质数
 *
 * https://leetcode.cn/problems/closest-prime-numbers-in-range/description/
 *
 * algorithms
 * Medium (35.94%)
 * Total Accepted:    2.6K
 * Total Submissions: 7.3K
 * Testcase Example:  '10\n19'
 *
 * 给你两个正整数 left 和 right ，请你找到两个整数 num1 和 num2 ，它们满足：
 * 
 * 
 * left <= nums1 < nums2 <= right  。
 * nums1 和 nums2 都是 质数 。
 * nums2 - nums1 是满足上述条件的质数对中的 最小值 。
 * 
 * 
 * 请你返回正整数数组 ans = [nums1, nums2] 。如果有多个整数对满足上述条件，请你返回 nums1
 * 最小的质数对。如果不存在符合题意的质数对，请你返回 [-1, -1] 。
 * 
 * 如果一个整数大于 1 ，且只能被 1 和它自己整除，那么它是一个质数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：left = 10, right = 19
 * 输出：[11,13]
 * 解释：10 到 19 之间的质数为 11 ，13 ，17 和 19 。
 * 质数对的最小差值是 2 ，[11,13] 和 [17,19] 都可以得到最小差值。
 * 由于 11 比 17 小，我们返回第一个质数对。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：left = 4, right = 6
 * 输出：[-1,-1]
 * 解释：给定范围内只有一个质数，所以题目条件无法被满足。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= left <= right <= 10^6
 * 
 * 
 */
public class Solution
{
    private const int N = (int) 1e6;

    public int[] ClosestPrimes(int left, int right)
    {
        var isPrime = new bool[N + 1];
        var primes = new List<int>();
        for (var i = 0; i <= N; i++)
        {
            isPrime[i] = true;
        }
        for (var i = 2; i <= N; i++)
        {
            if (isPrime[i])
            {
                primes.Add(i);
                for (var j = i; j <= N / i; j++)
                {
                    isPrime[i * j] = false;
                }
            }
        }
        var n = primes.Count();
        int p = 0, q = n;
        while (p < q)
        {
            var mid = (p + q) >> 1;
            if (primes[mid] >= left)
            {
                q = mid;
            }
            else
            {
                p = mid + 1;
            }
        }
        var ans = new int[] { -1, -1, };
        for (var i = p; i + 1 < n && primes[i + 1] <= right; i++)
        {
            if (ans[0] == -1 || ans[1] - ans[0] > primes[i + 1] - primes[i])
            {
                ans[0] = primes[i];
                ans[1] = primes[i + 1];
            }
        }
        return ans;
    }
}
