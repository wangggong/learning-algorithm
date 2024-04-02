/*
 * @lc app=leetcode.cn id=letter-tile-possibilities lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1079] 活字印刷
 *
 * https://leetcode.cn/problems/letter-tile-possibilities/description/
 *
 * algorithms
 * Medium (73.86%)
 * Total Accepted:    23.8K
 * Total Submissions: 30.8K
 * Testcase Example:  '"AAB"'
 *
 * 你有一套活字字模 tiles，其中每个字模上都刻有一个字母 tiles[i]。返回你可以印出的非空字母序列的数目。
 * 
 * 注意：本题中，每个活字字模只能使用一次。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入："AAB"
 * 输出：8
 * 解释：可能的序列为 "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA"。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入："AAABBC"
 * 输出：188
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入："V"
 * 输出：1
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= tiles.length <= 7
 * tiles 由大写英文字母组成
 * 
 * 
 */
public class Solution
{
    private const int N = 7;

    public int NumTilePossibilities(string tiles)
    {
        var ans = 0;
        var factorials = new int[N + 1];
        factorials[0] = 1;
        for (var i = 1; i <= N; i++)
        {
            factorials[i] = factorials[i - 1] * i;
        }
        int getPermutate(string s)
        {
            var ans = factorials[s.Length];
            foreach (var count in s.GroupBy(c => c).Select(g => g.Count()))
            {
                ans /= factorials[count];
            }
            return ans;
        }
        void backtrack(string curr, int k)
        {
            var n = tiles.Length;
            if (k == n)
            {
                ans += getPermutate(curr);
                return;
            }
            ans += getPermutate(curr);
            for (var i = k; i < n; i++)
            {
                if (i > k && tiles[i] == tiles[i - 1])
                {
                    continue;
                }
                backtrack(curr + tiles.Substring(i, 1), i + 1);
            }
        }
        tiles = string.Concat(tiles.OrderBy(c => c));
        backtrack("", 0);
        return ans - 1;
    }
}
