/*
 * @lc app=leetcode.cn id=determine-if-two-strings-are-close lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1657] 确定两个字符串是否接近
 *
 * https://leetcode.cn/problems/determine-if-two-strings-are-close/description/
 *
 * algorithms
 * Medium (48.17%)
 * Total Accepted:    20.3K
 * Total Submissions: 41.6K
 * Testcase Example:  '"abc"\n"bca"'
 *
 * 如果可以使用以下操作从一个字符串得到另一个字符串，则认为两个字符串 接近 ：
 * 
 * 
 * 操作 1：交换任意两个 现有 字符。
 * 
 * 
 * 例如，abcde -> aecdb
 * 
 * 
 * 操作 2：将一个 现有 字符的每次出现转换为另一个 现有 字符，并对另一个字符执行相同的操作。
 * 
 * 例如，aacabb -> bbcbaa（所有 a 转化为 b ，而所有的 b 转换为 a ）
 * 
 * 
 * 
 * 
 * 你可以根据需要对任意一个字符串多次使用这两种操作。
 * 
 * 给你两个字符串，word1 和 word2 。如果 word1 和 word2 接近 ，就返回 true ；否则，返回 false 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：word1 = "abc", word2 = "bca"
 * 输出：true
 * 解释：2 次操作从 word1 获得 word2 。
 * 执行操作 1："abc" -> "acb"
 * 执行操作 1："acb" -> "bca"
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：word1 = "a", word2 = "aa"
 * 输出：false
 * 解释：不管执行多少次操作，都无法从 word1 得到 word2 ，反之亦然。
 * 
 * 示例 3：
 * 
 * 
 * 输入：word1 = "cabbba", word2 = "abbccc"
 * 输出：true
 * 解释：3 次操作从 word1 获得 word2 。
 * 执行操作 1："cabbba" -> "caabbb"
 * 执行操作 2："caabbb" -> "baaccc"
 * 执行操作 2："baaccc" -> "abbccc"
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：word1 = "cabbba", word2 = "aabbss"
 * 输出：false
 * 解释：不管执行多少次操作，都无法从 word1 得到 word2 ，反之亦然。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * word1 和 word2 仅包含小写英文字母
 * 
 * 
 */
public class Solution
{
    public bool CloseStrings(string word1, string word2) =>
        SameChars(word1, word2) && SameCounts(word1, word2);

    private bool SameChars(string word1, string word2)
    {
        var chars1 = word1.ToHashSet();
        var chars2 = word2.ToHashSet();
        return chars1.Count() == chars2.Count() &&
            chars1.Intersect(chars2).Count() == chars1.Count();
    }

    private bool SameCounts(string word1, string word2)
    {
        IList<int> getCounts(string s) => s
            .GroupBy(c => c)
            .Select(g => g.Count())
            .OrderBy(x => x)
            .ToList();
        var counts1 = getCounts(word1);
        var counts2 = getCounts(word2);
        return counts1.Count() == counts2.Count()
            && counts1
                .Zip(counts2, (c1, c2) => (c1, c2))
                .All(x => x.c1 == x.c2);
    }
}
