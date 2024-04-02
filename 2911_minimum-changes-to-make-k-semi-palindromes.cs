/*
 * Hard
 * 
 * 给你一个字符串 s 和一个整数 k ，请你将 s 分成 k 个 子字符串 ，使得每个 子字符串 变成 半回文串 需要修改的字符数目最少。
 * 
 * 请你返回一个整数，表示需要修改的 最少 字符数目。
 * 
 * 注意：
 * 
 * 如果一个字符串从左往右和从右往左读是一样的，那么它是一个 回文串 。
 * 如果长度为 len 的字符串存在一个满足 1 <= d < len 的正整数 d ，len % d == 0 成立且所有对 d 做除法余数相同的下标对应的字符连起来得到的字符串都是 回文串 ，那么我们说这个字符串是 半回文串 。比方说 "aa" ，"aba" ，"adbgad" 和 "abab" 都是 半回文串 ，而 "a" ，"ab" 和 "abca" 不是。
 * 子字符串 指的是一个字符串中一段连续的字符序列。
 *  
 * 
 * 示例 1：
 * 
 * 输入：s = "abcac", k = 2
 * 输出：1
 * 解释：我们可以将 s 分成子字符串 "ab" 和 "cac" 。子字符串 "cac" 已经是半回文串。如果我们将 "ab" 变成 "aa" ，它也会变成一个 d = 1 的半回文串。
 * 该方案是将 s 分成 2 个子字符串的前提下，得到 2 个半回文子字符串需要的最少修改次数。所以答案为 1 。
 * 示例 2:
 * 
 * 输入：s = "abcdef", k = 2
 * 输出：2
 * 解释：我们可以将 s 分成子字符串 "abc" 和 "def" 。子字符串 "abc" 和 "def" 都需要修改一个字符得到半回文串，所以我们总共需要 2 次字符修改使所有子字符串变成半回文串。
 * 该方案是将 s 分成 2 个子字符串的前提下，得到 2 个半回文子字符串需要的最少修改次数。所以答案为 2 。
 * 示例 3：
 * 
 * 输入：s = "aabbaa", k = 3
 * 输出：0
 * 解释：我们可以将 s 分成子字符串 "aa" ，"bb" 和 "aa" 。
 * 字符串 "aa" 和 "bb" 都已经是半回文串了。所以答案为 0 。
 *  
 * 
 * 提示：
 * 
 * 2 <= s.length <= 200
 * 1 <= k <= s.length / 2
 * s 只包含小写英文字母。
 * 
 * 通过次数 1.6K
 * 提交次数 3.5K
 * 通过率 46.2%
 * 
 */

// 参考: https://leetcode.cn/problems/minimum-changes-to-make-k-semi-palindromes/solutions/2493147/yu-chu-li-ji-yi-hua-sou-suo-by-endlessch-qp47/
//
// 预处理+划分型 DP.
// 预处理部分直接暴力搞就可以, 但这里有个结论:
//  `n` 的真因子数平均是 `O(logn)` 的, 最坏大致是 `O(n^(1/3))` 的.
public class Solution
{
    private const int N = 200;
    private static List<int>[] divisors = new List<int>[N + 1];

    static Solution()
    {
        for (var i = 1; i <= N; i++) { divisors[i] = new List<int>(); }
        for (var i = 1; i <= N; i++)
        {
            for (var j = 2; j <= N / i; j++) { divisors[i * j].Add(i); }
        }
    }

    private int GetModify(string s) => divisors[s.Length]
        .Select(d => Enumerable.Range(0, d)
            .Select(k =>
            {
                var sb = new StringBuilder();
                for (var (i, n) = (k, s.Length); i < n; i += d) { sb.Append(s[i]); }
                var t = sb.ToString();
                return Enumerable.Range(0, t.Length / 2)
                    .Count(i => t[i] != t[t.Length - 1 - i]);
            }).Sum())
        .Min();

    public int MinimumChanges(string s, int k)
    {
        var n = s.Length;
        var modifies = new int[n][];
        for (var i = 0; i < n; i++)
        {
            modifies[i] = new int[n + 1];
            for (var j = i + 2; j <= n; j++) { modifies[i][j] = GetModify(s[i..j]); }
        }
        var memo = new Dictionary<(int, int), int>();
        int dfs(int n, int k)
        {
            var key = (n, k);
            if (!memo.ContainsKey(key))
            {
                memo[key] = k is 0
                    ? modifies[0][n]
                    : Enumerable.Range(2 * k, (n - 1) - 2 * k)
                        .Select(i => dfs(i, k - 1) + modifies[i][n])
                        .Min();
            }
            return memo[key];
        }
        return dfs(n, k - 1);
    }
}
