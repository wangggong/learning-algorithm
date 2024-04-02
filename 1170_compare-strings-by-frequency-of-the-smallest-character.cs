/*
 * @lc app=leetcode.cn id=compare-strings-by-frequency-of-the-smallest-character lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1170] 比较字符串最小字母出现频次
 *
 * https://leetcode.cn/problems/compare-strings-by-frequency-of-the-smallest-character/description/
 *
 * algorithms
 * Medium (61.72%)
 * Total Accepted:    17.2K
 * Total Submissions: 27.7K
 * Testcase Example:  '["cbd"]\n["zaaaz"]'
 *
 * 定义一个函数 f(s)，统计 s  中（按字典序比较）最小字母的出现频次 ，其中 s 是一个非空字符串。
 * 
 * 例如，若 s = "dcce"，那么 f(s) = 2，因为字典序最小字母是 "c"，它出现了 2 次。
 * 
 * 现在，给你两个字符串数组待查表 queries 和词汇表 words 。对于每次查询 queries[i] ，需统计 words 中满足
 * f(queries[i]) < f(W) 的 词的数目 ，W 表示词汇表 words 中的每个词。
 * 
 * 请你返回一个整数数组 answer 作为答案，其中每个 answer[i] 是第 i 次查询的结果。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：queries = ["cbd"], words = ["zaaaz"]
 * 输出：[1]
 * 解释：查询 f("cbd") = 1，而 f("zaaaz") = 3 所以 f("cbd") < f("zaaaz")。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：queries = ["bbb","cc"], words = ["a","aa","aaa","aaaa"]
 * 输出：[1,2]
 * 解释：第一个查询 f("bbb") < f("aaaa")，第二个查询 f("aaa") 和 f("aaaa") 都 > f("cc")。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 1 
 * 1 
 * queries[i][j]、words[i][j] 都由小写英文字母组成
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     public int[] NumSmallerByFrequency(string[] queries, string[] words)
 *     {
 * 	    int f(string s) => s
 *             .GroupBy(c => c)
 *             .Select(g => (g.Key, g.Count()))
 *             .OrderBy(x => x.Item1)
 *             .First()
 *             .Item2;
 *         var fs = words
 *             .Select(w => f(w))
 *             .OrderBy(x => x)
 *             .ToArray();
 *         var (j, n) = (0, fs.Length);
 *         var ans = new List<(int, int)>();
 *         foreach (var (q, i) in queries
 *             .Select((q, i) => (f(q), i))
 *             .OrderBy(x => x.Item1))
 *         {
 *             for (; j < n && fs[j] <= q; j++) { }
 *             ans.Add((i, n - j));
 *         }
 *         return ans
 *             .OrderBy(x => x.Item1)
 *             .Select(x => x.Item2)
 *             .ToArray();
 *     }
 * }
 */

// 官解多少沾点没事找事了 (
public class Solution
{
    private const int N = 10;

    public int[] NumSmallerByFrequency(string[] queries, string[] words)
    {
        int f(string s)
        {
            var (ch, k) = ('z', 0);
            foreach (var c in s)
            {
                if (c < ch)
                {
                    (ch, k) = (c, 1);
                }
                else if (c == ch)
                {
                    k++;
                }
            }
            return k;
        }
        var counts = new int[N + 1];
        foreach (var word in words)
        {
            counts[f(word)]++;
        }
        return queries.Select(query =>
        {
            var ans = 0;
            for (var i = f(query) + 1; i <= N; i++)
            {
                ans += counts[i];
            }
            return ans;
        }).ToArray();
    }
}
