/*
 * @lc app=leetcode.cn id=length-of-the-longest-valid-substring lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6924] 最长合法子字符串的长度
 *
 * https://leetcode.cn/problems/length-of-the-longest-valid-substring/description/
 *
 * algorithms
 * Hard (25.15%)
 * Total Accepted:    1.1K
 * Total Submissions: 4.4K
 * Testcase Example:  '"cbaaaabc"\n["aaa","cb"]'
 *
 * 给你一个字符串 word 和一个字符串数组 forbidden 。
 * 
 * 如果一个字符串不包含 forbidden 中的任何字符串，我们称这个字符串是 合法 的。
 * 
 * 请你返回字符串 word 的一个 最长合法子字符串 的长度。
 * 
 * 子字符串 指的是一个字符串中一段连续的字符，它可以为空。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：word = "cbaaaabc", forbidden = ["aaa","cb"]
 * 输出：4
 * 解释：总共有 9 个合法子字符串："c" ，"b" ，"a" ，"ba" ，"aa" ，"bc" ，"baa" ，"aab" 和 "aabc"
 * 。最长合法子字符串的长度为 4 。
 * 其他子字符串都要么包含 "aaa" ，要么包含 "cb" 。
 * 
 * 示例 2：
 * 
 * 输入：word = "leetcode", forbidden = ["de","le","e"]
 * 输出：4
 * 解释：总共有 11 个合法子字符串："l" ，"t" ，"c" ，"o" ，"d" ，"tc" ，"co" ，"od" ，"tco" ，"cod" 和
 * "tcod" 。最长合法子字符串的长度为 4 。
 * 所有其他子字符串都至少包含 "de" ，"le" 和 "e" 之一。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= word.length <= 10^5
 * word 只包含小写英文字母。
 * 1 <= forbidden.length <= 10^5
 * 1 <= forbidden[i].length <= 10
 * forbidden[i] 只包含小写英文字母。
 * 
 * 
 */
public class Solution
{
    private const long P = 13331;
    private const int N = 10;

    public int LongestValidSubstring(string word, IList<string> forbidden)
    {
        var n = word.Length;
        var hashs = new long[n + 1];
        var muls = new long[n + 1];
        muls[0] = 1;
        for (var i = 0; i < n; i++)
        {
            hashs[i + 1] = hashs[i] * P + (long)word[i];
            muls[i + 1] = muls[i] * P;
        }
        long get(int p, int q) => hashs[q] - hashs[p] * muls[q - p];
        var ds = new Dictionary<long, HashSet<string>>[N + 1];
        for (var i = 0; i <= N; i++) { ds[i] = new(); }
        foreach (var s in forbidden)
        {
            var h = s.Aggregate(0L, (x, y) => x * P + (long)y);
            if (!ds[s.Length].ContainsKey(h)) { ds[s.Length][h] = new(); }
            ds[s.Length][h].Add(s);
        }
        bool valid(int p, int q) => Enumerable.Range(1, N).All(l =>
        {
            if (q - p + 1 < l) { return true; }
            var h = get(q - l + 1, q + 1);
            if (!ds[l].ContainsKey(h)) { return true; }
            return !ds[l][h].Contains(word.Substring(q - l + 1, l));
        });
        var ans = 0;
        for (var (p, q) = (0, 0); p < n; p++)
        {
            for (; q < n && valid(p, q); q++) { }
            ans = Math.Max(ans, q - p);
        }
        return ans;
    }
}
