/*
 * @lc app=leetcode.cn id=find-beautiful-indices-in-the-given-array-i lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100165] 找出数组中的美丽下标 I
 *
 * https://leetcode.cn/problems/find-beautiful-indices-in-the-given-array-i/description/
 *
 * algorithms
 * Medium (41.19%)
 * Total Accepted:    3K
 * Total Submissions: 7.3K
 * Testcase Example:  '"isawsquirrelnearmysquirrelhouseohmy"\n"my"\n"squirrel"\n15'
 *
 * 给你一个下标从 0 开始的字符串 s 、字符串 a 、字符串 b 和一个整数 k 。
 * 
 * 如果下标 i 满足以下条件，则认为它是一个 美丽下标：
 * 
 * 
 * 0 <= i <= s.length - a.length
 * s[i..(i + a.length - 1)] == a
 * 存在下标 j 使得：
 * 
 * 0 <= j <= s.length - b.length
 * s[j..(j + b.length - 1)] == b
 * |j - i| <= k
 * 
 * 
 * 
 * 
 * 以数组形式按 从小到大排序 返回美丽下标。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "isawsquirrelnearmysquirrelhouseohmy", a = "my", b = "squirrel", k =
 * 15
 * 输出：[16,33]
 * 解释：存在 2 个美丽下标：[16,33]。
 * - 下标 16 是美丽下标，因为 s[16..17] == "my" ，且存在下标 4 ，满足 s[4..11] == "squirrel" 且 |16
 * - 4| <= 15 。
 * - 下标 33 是美丽下标，因为 s[33..34] == "my" ，且存在下标 18 ，满足 s[18..25] == "squirrel" 且
 * |33 - 18| <= 15 。
 * 因此返回 [16,33] 作为结果。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "abcd", a = "a", b = "a", k = 4
 * 输出：[0]
 * 解释：存在 1 个美丽下标：[0]。
 * - 下标 0 是美丽下标，因为 s[0..0] == "a" ，且存在下标 0 ，满足 s[0..0] == "a" 且 |0 - 0| <= 4 。
 * 因此返回 [0] 作为结果。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= k <= s.length <= 10^5
 * 1 <= a.length, b.length <= 10
 * s、a、和 b 只包含小写英文字母。
 * 
 * 
 */
public class Solution
{
    public IList<int> BeautifulIndices(string s, string a, string b, int k)
    {
        var (n, na, nb) = (s.Length, a.Length, b.Length);
        if (n < na || n < nb) { return new List<int>(); }
        var indexes = Enumerable.Range(0, n - nb + 1)
            .Where(i => s[i..(i + nb)] == b)
            .ToList();
        var m = indexes.Count();
        return Enumerable.Range(0, n - na + 1)
            .Where(i => s[i..(i + na)] == a)
            .Where(i =>
            {
                var (p, q) = (0, m);
                while (p < q)
                {
                    var mid = (p + q) >> 1;
                    if (indexes[mid] >= i) { q = mid; }
                    else { p = mid + 1; }
                }
                return (p < m && Math.Abs(indexes[p] - i) <= k)
                    || (p - 1 >= 0 && Math.Abs(indexes[p - 1] - i) <= k);
            }).ToList();
    }
}
