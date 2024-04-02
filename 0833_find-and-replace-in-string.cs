/*
 * @lc app=leetcode.cn id=find-and-replace-in-string lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [833] 字符串中的查找与替换
 *
 * https://leetcode.cn/problems/find-and-replace-in-string/description/
 *
 * algorithms
 * Medium (44.18%)
 * Total Accepted:    8.6K
 * Total Submissions: 19.3K
 * Testcase Example:  '"abcd"\n[0, 2]\n["a", "cd"]\n["eee", "ffff"]'
 *
 * 你会得到一个字符串 s (索引从 0 开始)，你必须对它执行 k 个替换操作。替换操作以三个长度均为 k 的并行数组给出：indices,
 * sources,  targets。
 * 
 * 要完成第 i 个替换操作:
 * 
 * 
 * 检查 子字符串  sources[i] 是否出现在 原字符串 s 的索引 indices[i] 处。
 * 如果没有出现， 什么也不做 。
 * 如果出现，则用 targets[i] 替换 该子字符串。
 * 
 * 
 * 例如，如果 s = "abcd" ， indices[i] = 0 , sources[i] = "ab"， targets[i] = "eee"
 * ，那么替换的结果将是 "eeecd" 。
 * 
 * 所有替换操作必须 同时 发生，这意味着替换操作不应该影响彼此的索引。测试用例保证元素间不会重叠 。
 * 
 * 
 * 例如，一个 s = "abc" ，  indices = [0,1] ， sources = ["ab"，"bc"] 的测试用例将不会生成，因为
 * "ab" 和 "bc" 替换重叠。
 * 
 * 
 * 在对 s 执行所有替换操作后返回 结果字符串 。
 * 
 * 子字符串 是字符串中连续的字符序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：s = "abcd", indexes = [0,2], sources = ["a","cd"], targets =
 * ["eee","ffff"]
 * 输出："eeebffff"
 * 解释：
 * "a" 从 s 中的索引 0 开始，所以它被替换为 "eee"。
 * "cd" 从 s 中的索引 2 开始，所以它被替换为 "ffff"。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "abcd", indexes = [0,2], sources = ["ab","ec"], targets =
 * ["eee","ffff"]
 * 输出："eeecd"
 * 解释：
 * "ab" 从 s 中的索引 0 开始，所以它被替换为 "eee"。
 * "ec" 没有从原始的 S 中的索引 2 开始，所以它没有被替换。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 1000
 * k == indices.length == sources.length == targets.length
 * 1 <= k <= 100
 * 0 <= indexes[i] < s.length
 * 1 <= sources[i].length, targets[i].length <= 50
 * s 仅由小写英文字母组成
 * sources[i] 和 targets[i] 仅由小写英文字母组成
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     public string FindReplaceString(string s, int[] indices, string[] sources, string[] targets)
 *     {
 *         var n = s.Length;
 *         var sb = new StringBuilder();
 *         var d = sources.Zip(targets, (s, t) => (s, t))
 *             .Zip(indices, (st, i) => (i: i, s: st.s, t: st.t))
 *             .GroupBy(x => x.i)
 *             .ToDictionary(g => g.Key, g => g);
 *         (bool, int) tryToReplace(int i)
 *         {
 *             d.TryGetValue(i, out var pairs);
 *             if (pairs is null) { return (false, -1); }
 *             foreach (var (_, _s, _t) in pairs)
 *             {
 *                 if (i + _s.Length <= n && s.Substring(i, _s.Length) == _s)
 *                 {
 *                     sb.Append(_t);
 *                     return (true, _s.Length);
 *                 }
 *             }
 *             return (false, -1);
 *         }
 *         for (var i = 0; i < n; i++)
 *         {
 *             var (replaced, length) = tryToReplace(i);
 *             if (!replaced) { sb.Append(s[i]); }
 *             else { i += length - 1; }
 *         }
 *         return sb.ToString();
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/find-and-replace-in-string/solutions/2387388/zi-fu-chuan-zhong-de-cha-zhao-yu-ti-huan-9ns4/comments/2094858
public class Solution
{
    public string FindReplaceString(string s, int[] indices, string[] sources, string[] targets)
    {
        var n = s.Length;
        var replaces = s.Select(c => (new string(c, 1), 1)).ToArray();
        foreach (var (i, _s, _t) in sources.Zip(targets, (s, t) => (s, t))
            .Zip(indices, (st, i) => (i: i, s: st.s, t: st.t)))
        {
            if (i + _s.Length <= n && s.Substring(i, _s.Length) == _s)
            {
                replaces[i] = (_t, _s.Length);
            }
        }
        var sb = new StringBuilder();
        for (var i = 0; i < n; i++)
        {
            var (r, j) = replaces[i];
            sb.Append(r);
            i += j - 1;
        }
        return sb.ToString();
    }
}
