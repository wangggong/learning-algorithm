/*
 * @lc app=leetcode.cn id=tiling-a-rectangle-with-the-fewest-squares lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1240] 铺瓷砖
 *
 * https://leetcode.cn/problems/tiling-a-rectangle-with-the-fewest-squares/description/
 *
 * algorithms
 * Hard (52.32%)
 * Total Accepted:    11.6K
 * Total Submissions: 18.1K
 * Testcase Example:  '2\n3'
 *
 * 你是一位施工队的工长，根据设计师的要求准备为一套设计风格独特的房子进行室内装修。
 * 
 * 房子的客厅大小为 n x m，为保持极简的风格，需要使用尽可能少的 正方形 瓷砖来铺盖地面。
 * 
 * 假设正方形瓷砖的规格不限，边长都是整数。
 * 
 * 请你帮设计师计算一下，最少需要用到多少块方形瓷砖？
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：n = 2, m = 3
 * 输出：3
 * 解释：3 块地砖就可以铺满卧室。
 * ⁠    2 块 1x1 地砖
 * ⁠    1 块 2x2 地砖
 * 
 * 示例 2：
 * 
 * 
 * 
 * 输入：n = 5, m = 8
 * 输出：5
 * 
 * 
 * 示例 3：
 * 
 * 
 * 
 * 输入：n = 11, m = 13
 * 输出：6
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 13
 * 1 <= m <= 13
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/tiling-a-rectangle-with-the-fewest-squares/solution/pu-ci-zhuan-by-leetcode-solution-r1bk/
public class Solution
{
    public int TilingRectangle(int n, int m)
    {
        var ans = m * n;
        void backtrack(int x, int y, int[] G, int count)
        {
            if (count > ans) { return; }
            if (x >= n)
            {
                ans = count;
                return;
            }
            if (y >= m)
            {
                backtrack(x + 1, 0, G, count);
                return;
            }
            if (((G[y] >> x) & 1) != 0)
            {
                backtrack(x, y + 1, G, count);
                return;
            }
            for (var k = Math.Min(n, m); k > 0; k--)
            {
                if (valid(G, x, y, k))
                {
                    backtrack(x, y, filled(G, x, y, k), count + 1);
                }
            }
        }
        bool valid(int[] G, int x, int y, int k) => x + k <= n
            && y + k <= m
            && Enumerable
                .Range(0, k)
                .All(i => Enumerable
                    .Range(0, k)
                    .All(j => ((G[y + j] >> (x + i)) & 1) == 0));
        int[] filled(int[] G, int x, int y, int k)
        {
            var ans = G.Select(x => x).ToArray();
            for (var i = 0; i < k; i++)
            {
                for (var j = 0; j < k; j++) { ans[y + j] |= 1 << (x + i); }
            }
            return ans;
        }
        backtrack(0, 0, new int[m], 0);
        return ans;
    }
}
