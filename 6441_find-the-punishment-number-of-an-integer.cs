/*
 * @lc app=leetcode.cn id=find-the-punishment-number-of-an-integer lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6441] 求一个整数的惩罚数
 *
 * https://leetcode.cn/problems/find-the-punishment-number-of-an-integer/description/
 *
 * algorithms
 * Medium (58.20%)
 * Total Accepted:    2.5K
 * Total Submissions: 4.3K
 * Testcase Example:  '10'
 *
 * 给你一个正整数 n ，请你返回 n 的 惩罚数 。
 * 
 * n 的 惩罚数 定义为所有满足以下条件 i 的数的平方和：
 * 
 * 
 * 1 <= i <= n
 * i * i 的十进制表示的字符串可以分割成若干连续子字符串，且这些子字符串对应的整数值之和等于 i 。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 10
 * 输出：182
 * 解释：总共有 3 个整数 i 满足要求：
 * - 1 ，因为 1 * 1 = 1
 * - 9 ，因为 9 * 9 = 81 ，且 81 可以分割成 8 + 1 。
 * - 10 ，因为 10 * 10 = 100 ，且 100 可以分割成 10 + 0 。
 * 因此，10 的惩罚数为 1 + 81 + 100 = 182
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 37
 * 输出：1478
 * 解释：总共有 4 个整数 i 满足要求：
 * - 1 ，因为 1 * 1 = 1
 * - 9 ，因为 9 * 9 = 81 ，且 81 可以分割成 8 + 1 。
 * - 10 ，因为 10 * 10 = 100 ，且 100 可以分割成 10 + 0 。
 * - 36 ，因为 36 * 36 = 1296 ，且 1296 可以分割成 1 + 29 + 6 。
 * 因此，37 的惩罚数为 1 + 81 + 100 + 1296 = 1478
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 1000
 * 
 * 
 */

/*
 * // 注意, 这种写法不会调用静态构造.
 * public class Solution
 * {
 *     private const int N = 1000;
 *     private bool[] Ok;
 * 
 *     private bool Backtrack(string s, int n, int k)
 *     {
 *         if (k == s.Length)
 *         {
 *             return n == 0;
 *         }
 *         for (var (i, cur) = (k, 0); i < s.Length; i++)
 *         {
 *             cur = cur * 10 + (int)(s[i] - '0');
 *             if (cur > n)
 *             {
 *                 break;
 *             }
 *             if (Backtrack(s, n - cur, i + 1))
 *             {
 *                 return true;
 *             }
 *         }
 *         return false;
 *     }
 * 
 *     public Solution() => Ok = Enumerable.Range(0, N + 1)
 *         .Select(k => Backtrack((k * k).ToString(), k, 0))
 *         .ToArray();
 * 
 *     public int PunishmentNumber(int n) => Enumerable.Range(1, n)
 *         .Where(k => Ok[k])
 *         .Select(k => k * k)
 *         .Sum();
 * }
 */

public class Solution
{
    private const int N = 1000;
    private static int[] punishmentNumber = new int[N + 1];

    static Solution()
    {
        bool dfs(string s, int n, int i, int k)
        {
            var m = s.Length;
            if (i == m) { return k == n; }
            for (var c = 0; i < m; i++)
            {
                c = c * 10 + (int)(s[i] - '0');
                if (dfs(s, n, i + 1, k + c)) { return true; }
            }
            return false;
        }
        bool check(int n) => dfs((n * n).ToString(), n, 0, 0);
        for (var i = 1; i <= N; i++)
        {
            punishmentNumber[i] = punishmentNumber[i - 1] + (check(i) ? i * i : 0);
        }
    }

    public int PunishmentNumber(int n) => punishmentNumber[n];
}
