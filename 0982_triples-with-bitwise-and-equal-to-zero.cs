/*
 * @lc app=leetcode.cn id=triples-with-bitwise-and-equal-to-zero lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [982] 按位与为零的三元组
 *
 * https://leetcode.cn/problems/triples-with-bitwise-and-equal-to-zero/description/
 *
 * algorithms
 * Hard (57.69%)
 * Total Accepted:    4.6K
 * Total Submissions: 7.6K
 * Testcase Example:  '[2,1,3]'
 *
 * 给你一个整数数组 nums ，返回其中 按位与三元组 的数目。
 * 
 * 按位与三元组 是由下标 (i, j, k) 组成的三元组，并满足下述全部条件：
 * 
 * 
 * 0 <= i < nums.length
 * 0 <= j < nums.length
 * 0 <= k < nums.length
 * nums[i] & nums[j] & nums[k] == 0 ，其中 & 表示按位与运算符。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [2,1,3]
 * 输出：12
 * 解释：可以选出如下 i, j, k 三元组：
 * (i=0, j=0, k=1) : 2 & 2 & 1
 * (i=0, j=1, k=0) : 2 & 1 & 2
 * (i=0, j=1, k=1) : 2 & 1 & 1
 * (i=0, j=1, k=2) : 2 & 1 & 3
 * (i=0, j=2, k=1) : 2 & 3 & 1
 * (i=1, j=0, k=0) : 1 & 2 & 2
 * (i=1, j=0, k=1) : 1 & 2 & 1
 * (i=1, j=0, k=2) : 1 & 2 & 3
 * (i=1, j=1, k=0) : 1 & 1 & 2
 * (i=1, j=2, k=0) : 1 & 3 & 2
 * (i=2, j=0, k=1) : 3 & 2 & 1
 * (i=2, j=1, k=0) : 3 & 1 & 2
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [0,0,0]
 * 输出：27
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 1000
 * 0 <= nums[i] < 2^16
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/triples-with-bitwise-and-equal-to-zero/solution/you-ji-qiao-de-mei-ju-chang-shu-you-hua-daxit/
public class Solution
{
    private const int Maxn = 0xffff;

    public int CountTriplets(int[] nums)
    {
        var count = new Dictionary<int, int>();
        foreach (var x in nums)
        {
            foreach (var y in nums)
            {
                count.TryGetValue(x & y, out var c);
                count[x & y] = c + 1;
            }
        }
        var ans = 0;
        foreach (var n in nums)
        {
            var (U, s) = (n ^ Maxn, n ^ Maxn);
            do
            {
                count.TryGetValue(s, out var c);
                ans += c;
                s = (s - 1) & U;
            }
            while (s != U);
        }
        return ans;
    }

    /*
     * // 想复杂了还是... (虽然也能过)
     * public int CountTriplets(int[] nums)
     * {
     *     var len = nums.Length;
     *     long total = (long)len * (long)len * (long)len;
     *     var count = new Dictionary<int, int>();
     *     foreach (var n in nums)
     *     {
     *         if (n == 0)
     *         {
     *             continue;
     *         }
     *         count.TryGetValue(n, out var c);
     *         count[n] = c + 1;
     *     }
     *     for (var i = 0; i < 2; i++)
     *     {
     *         var newCount = new Dictionary<int, int>();
     *         foreach (var n in nums)
     *         {
     *             foreach (var (k, v) in count)
     *             {
     *                 var next = n & k;
     *                 if (next == 0)
     *                 {
     *                     continue;
     *                 }
     *                 newCount.TryGetValue(next, out var c);
     *                 newCount[next] = c + v;
     *             }
     *         }
     *         count = newCount;
     *     }
     *     return (int)(total - count.Select(kv => (long)kv.Value).Sum());
     * }
     */
}
