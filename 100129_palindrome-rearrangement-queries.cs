/*
 * @lc app=leetcode.cn id=palindrome-rearrangement-queries lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100129] 回文串重新排列查询
 *
 * https://leetcode.cn/problems/palindrome-rearrangement-queries/description/
 *
 * algorithms
 * Hard (15.21%)
 * Total Accepted:    821
 * Total Submissions: 2.9K
 * Testcase Example:  '"abcabc"\n[[1,1,3,5],[0,2,5,5]]'
 *
 * 给你一个长度为 偶数 n ，下标从 0 开始的字符串 s 。
 * 
 * 同时给你一个下标从 0 开始的二维整数数组 queries ，其中 queries[i] = [ai, bi, ci, di] 。
 * 
 * 对于每个查询 i ，你需要执行以下操作：
 * 
 * 
 * 将下标在范围 0 <= ai <= bi < n / 2 内的 子字符串 s[ai:bi] 中的字符重新排列。
 * 将下标在范围 n / 2 <= ci <= di < n 内的 子字符串 s[ci:di] 中的字符重新排列。
 * 
 * 
 * 对于每个查询，你的任务是判断执行操作后能否让 s 变成一个 回文串 。
 * 
 * 每个查询与其他查询都是 独立的 。
 * 
 * 请你返回一个下标从 0 开始的数组 answer ，如果第 i 个查询执行操作后，可以将 s 变为一个回文串，那么 answer[i] =
 * true，否则为 false 。
 * 
 * 
 * 子字符串 指的是一个字符串中一段连续的字符序列。
 * s[x:y] 表示 s 中从下标 x 到 y 且两个端点 都包含 的子字符串。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "abcabc", queries = [[1,1,3,5],[0,2,5,5]]
 * 输出：[true,true]
 * 解释：这个例子中，有 2 个查询：
 * 第一个查询：
 * - a0 = 1, b0 = 1, c0 = 3, d0 = 5
 * - 你可以重新排列 s[1:1] => abcabc 和 s[3:5] => abcabc 。
 * - 为了让 s 变为回文串，s[3:5] 可以重新排列得到 => abccba 。
 * - 现在 s 是一个回文串。所以 answer[0] = true 。
 * 第二个查询：
 * - a1 = 0, b1 = 2, c1 = 5, d1 = 5.
 * - 你可以重新排列 s[0:2] => abcabc 和 s[5:5] => abcabc 。
 * - 为了让 s 变为回文串，s[0:2] 可以重新排列得到 => cbaabc 。
 * - 现在 s 是一个回文串，所以 answer[1] = true 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "abbcdecbba", queries = [[0,2,7,9]]
 * 输出：[false]
 * 解释：这个示例中，只有一个查询。
 * a0 = 0, b0 = 2, c0 = 7, d0 = 9.
 * 你可以重新排列 s[0:2] => abbcdecbba 和 s[7:9] => abbcdecbba 。
 * 无法通过重新排列这些子字符串使 s 变为一个回文串，因为 s[3:6] 不是一个回文串。
 * 所以 answer[0] = false 。
 * 
 * 示例 3：
 * 
 * 
 * 输入：s = "acbcab", queries = [[1,2,4,5]]
 * 输出：[true]
 * 解释：这个示例中，只有一个查询。
 * a0 = 1, b0 = 2, c0 = 4, d0 = 5.
 * 你可以重新排列 s[1:2] => acbcab 和 s[4:5] => acbcab 。
 * 为了让 s 变为回文串，s[1:2] 可以重新排列得到 => abccab 。
 * 然后 s[4:5] 重新排列得到 abccba 。
 * 现在 s 是一个回文串，所以 answer[0] = true 。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= n == s.length <= 10^5
 * 1 <= queries.length <= 10^5
 * queries[i].length == 4
 * ai == queries[i][0], bi == queries[i][1]
 * ci == queries[i][2], di == queries[i][3]
 * 0 <= ai <= bi < n / 2
 * n / 2 <= ci <= di < n 
 * n 是一个偶数。
 * s 只包含小写英文字母。
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/palindrome-rearrangement-queries/solutions/2585862/fen-lei-tao-lun-by-endlesscheng-jxg0/
public class Solution
{
    public bool[] CanMakePalindromeQueries(string s, int[][] queries)
    {
        const int Alphas = 26;
        var n = s.Length >> 1;
        var (sf, st) = (s[..n], new string(s[n..].Reverse().ToArray()));
        int[][] getPrefixSum(string s) => Enumerable.Range(0, Alphas)
            .Select(c =>
            {
                var S = new int[n + 1];
                for (var i = 0; i < n; i++)
                { S[i + 1] = S[i] + ((int)(s[i] - 'a') == c ? 1 : 0); }
                return S;
            }).ToArray();
        var (Sf, St) = (getPrefixSum(sf), getPrefixSum(st));
        var Sd = new int[n + 1];
        for (var i = 0; i < n; i++)
        { Sd[i + 1] = Sd[i] + (sf[i] != st[i] ? 1 : 0); }
        IEnumerable<int> get(int[][] S, int p, int q) => Enumerable
            .Range(0, Alphas)
            .Select(i => S[i][q] - S[i][p]);
        bool equalTo(IEnumerable<int> s, IEnumerable<int> t) => s
            .Zip(t, (s, t) => s == t)
            .All(x => x is true);
        IEnumerable<int> substract(IEnumerable<int> s, IEnumerable<int> t) => s
            .Zip(t, (s, t) => s - t);
        return queries.Select(q =>
        {
            var (a, b, c, d) = (q[0], q[1] + 1, n * 2 - 1 - q[3], n * 2 - q[2]);
            var (f, t) = (Sf, St);
            if (a > c) { (a, b, c, d, f, t) = (c, d, a, b, t, f); }
            if ((Sd[a], Sd[n] - Sd[Math.Max(b, d)]) is not (0, 0))
            { return false; }
            if (b >= d) { return equalTo(get(f, a, b), get(t, a, b)); }
            if (b <= c)
            {
                return Sd[c] - Sd[b] is 0
                    && equalTo(get(f, a, b), get(t, a, b))
                    && equalTo(get(f, c, d), get(t, c, d));
            }
            var df = substract(get(f, a, b), get(t, a, c));
            var dt = substract(get(t, c, d), get(f, b, d));
            return !df.Any(x => x < 0)
                && !dt.Any(x => x < 0)
                && equalTo(df, dt);
        }).ToArray();
    }
}
