/*
 * @lc app=leetcode.cn id=count-the-number-of-infection-sequences lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100146] 统计感冒序列的数目
 *
 * https://leetcode.cn/problems/count-the-number-of-infection-sequences/description/
 *
 * algorithms
 * Hard (34.42%)
 * Total Accepted:    855
 * Total Submissions: 2.3K
 * Testcase Example:  '5\n[0,4]'
 *
 * 给你一个整数 n 和一个下标从 0 开始的整数数组 sick ，数组按 升序 排序。
 * 
 * 有 n 位小朋友站成一排，按顺序编号为 0 到 n - 1 。数组 sick 包含一开始得了感冒的小朋友的位置。如果位置为 i
 * 的小朋友得了感冒，他会传染给下标为 i - 1 或者 i + 1 的小朋友，前提 是被传染的小朋友存在且还没有得感冒。每一秒中， 至多一位
 * 还没感冒的小朋友会被传染。
 * 
 * 经过有限的秒数后，队列中所有小朋友都会感冒。感冒序列 指的是 所有 一开始没有感冒的小朋友最后得感冒的顺序序列。请你返回所有感冒序列的数目。
 * 
 * 由于答案可能很大，请你将答案对 10^9 + 7 取余后返回。
 * 
 * 注意，感冒序列 不 包含一开始就得了感冒的小朋友的下标。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 5, sick = [0,4]
 * 输出：4
 * 解释：一开始，下标为 1 ，2 和 3 的小朋友没有感冒。总共有 4 个可能的感冒序列：
 * - 一开始，下标为 1 和 3 的小朋友可以被传染，因为他们分别挨着有感冒的小朋友 0 和 4 ，令下标为 1 的小朋友先被传染。
 * 然后，下标为 2 的小朋友挨着感冒的小朋友 1 ，下标为 3 的小朋友挨着感冒的小朋友 4 ，两位小朋友都可以被传染，令下标为 2 的小朋友被传染。
 * 最后，下标为 3 的小朋友被传染，因为他挨着感冒的小朋友 2 和 4 ，感冒序列为 [1,2,3] 。
 * - 一开始，下标为 1 和 3 的小朋友可以被传染，因为他们分别挨着感冒的小朋友 0 和 4 ，令下标为 1 的小朋友先被传染。
 * 然后，下标为 2 的小朋友挨着感冒的小朋友 1 ，下标为 3 的小朋友挨着感冒的小朋友 4 ，两位小朋友都可以被传染，令下标为 3 的小朋友被传染。
 * 最后，下标为 2 的小朋友被传染，因为他挨着感冒的小朋友 1 和 3 ，感冒序列为  [1,3,2] 。
 * - 感冒序列 [3,1,2] ，被传染的顺序：[0,1,2,3,4] => [0,1,2,3,4] => [0,1,2,3,4] =>
 * [0,1,2,3,4] 。
 * - 感冒序列 [3,2,1] ，被传染的顺序：[0,1,2,3,4] => [0,1,2,3,4] => [0,1,2,3,4] =>
 * [0,1,2,3,4] 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 4, sick = [1]
 * 输出：3
 * 解释：一开始，下标为 0 ，2 和 3 的小朋友没有感冒。总共有 3 个可能的感冒序列：
 * - 感冒序列 [0,2,3] ，被传染的顺序：[0,1,2,3] => [0,1,2,3] => [0,1,2,3] => [0,1,2,3] 。
 * - 感冒序列 [2,0,3] ，被传染的顺序：[0,1,2,3] => [0,1,2,3] => [0,1,2,3] => [0,1,2,3] 。
 * - 感冒序列 [2,3,0] ，被传染的顺序：[0,1,2,3] => [0,1,2,3] => [0,1,2,3] => [0,1,2,3]
 * 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= n <= 10^5
 * 1 <= sick.length <= n - 1
 * 0 <= sick[i] <= n - 1
 * sick 按升序排列。
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/count-the-number-of-infection-sequences/solutions/2551734/zu-he-shu-xue-ti-by-endlesscheng-5fjp/
//
// 组合数学. 快速幂, 逆元, 乘法原理.
public class Solution
{
    private const long Mod = (long)1e9 + 7;
    private const int N = (int)1e5;
    private static long[] F;
    private static long[] If;

    static Solution()
    {
        F = new long[N + 1];
        F[0] = 1;
        for (var i = 1; i <= N; i++) { F[i] = F[i - 1] * i % Mod; }
        If = new long[N + 1];
        If[N] = Pow(F[N], Mod - 2);
        for (var i = N - 1; i >= 0; i--) { If[i] = If[i + 1] * (i + 1) % Mod; }
    }

    private static long Pow(long n, long k)
    {
        var ans = 1l;
        for (; k > 0; k >>= 1)
        {
            if ((k & 1) is not 0) { ans = ans * n % Mod; }
            n = n * n % Mod;
        }
        return ans;
    }

    private static long Comb(int n, int k) => F[n] * If[k] % Mod * If[n - k] % Mod;

    public int NumberOfSequence(int n, int[] sick)
    {
        var m = sick.Length;
        var total = n - m;
        var ans = Comb(total, sick.First())
            * Comb(total - sick.First(), n - 1 - sick.Last())
            % Mod;
        total -= sick.First() + (n - 1 - sick.Last());
        var e = 0;
        for (var i = 0; i + 1 < m; i++)
        {
            var cur = sick[i + 1] - sick[i] - 1;
            if (cur <= 0) { continue; }
            e += cur - 1;
            ans = ans * Comb(total, cur) % Mod;
            total -= cur;
        }
        return (int)(ans * Pow(2, e) % Mod);
    }
}
