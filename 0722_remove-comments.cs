/*
 * @lc app=leetcode.cn id=remove-comments lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [722] 删除注释
 *
 * https://leetcode.cn/problems/remove-comments/description/
 *
 * algorithms
 * Medium (34.23%)
 * Total Accepted:    17.8K
 * Total Submissions: 43.7K
 * 
 * 
 * 提示:
 * 
 * 
 * 1 <= source.length <= 100
 * 0 <= source[i].length <= 80
 * source[i] 由可打印的 ASCII 字符组成。
 * 每个块注释都会被闭合。
 * 给定的源码中不会有单引号、双引号或其他控制字符。
 * 
 * ​​​​​​
 */
public class Solution
{
    private const string LineCommitStart = "//";
    private const string BlockCommitStart = "/*";
    private const string BlockCommitEnd = "*/";

    public IList<string> RemoveComments(string[] source)
    {
        bool isLineCommitStart(string s) => s == LineCommitStart;
        bool isBlockCommitStart(string s) => s == BlockCommitStart;
        bool isBlockCommitEnd(string s) => s == BlockCommitEnd;
        bool isCommitStart(string s) => isLineCommitStart(s)
            || isBlockCommitStart(s);
        var s = string.Join('\n', source);
        var sb = new StringBuilder();
        for (var (p, q, n) = (0, 0, s.Length); p < n; p = q + 1)
        {
            for (; p + 1 < n && !isCommitStart(s.Substring(p, 2)); p++)
            {
                sb.Append(s[p]);
            }
            if (p + 1 == n)
            {
                sb.Append(s[p]);
                break;
            }
            if (isLineCommitStart(s.Substring(p, 2)))
            {
                for (q = p + 2; q < n && s[q] != '\n'; q++) { }
                sb.Append('\n');
            }
            else
            {
                for (q = p + 2;
                    q + 1 < n && !isBlockCommitEnd(s.Substring(q, 2));
                    q++) { }
                q++;
            }
        }
        return sb.ToString()
            .Split('\n')
            .Where(s => s.Length > 0)
            .ToList();
    }
}
