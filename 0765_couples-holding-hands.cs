/*
 * @lc app=leetcode.cn id=couples-holding-hands lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [765] 情侣牵手
 *
 * https://leetcode.cn/problems/couples-holding-hands/description/
 *
 * algorithms
 * Hard (65.78%)
 * Total Accepted:    39.7K
 * Total Submissions: 60.3K
 * Testcase Example:  '[0,2,1,3]'
 *
 * n 对情侣坐在连续排列的 2n 个座位上，想要牵到对方的手。
 * 
 * 人和座位由一个整数数组 row 表示，其中 row[i] 是坐在第 i 个座位上的人的 ID。情侣们按顺序编号，第一对是 (0, 1)，第二对是 (2,
 * 3)，以此类推，最后一对是 (2n-2, 2n-1)。
 * 
 * 返回 最少交换座位的次数，以便每对情侣可以并肩坐在一起。 每次交换可选择任意两人，让他们站起来交换座位。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入: row = [0,2,1,3]
 * 输出: 1
 * 解释: 只需要交换row[1]和row[2]的位置即可。
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: row = [3,2,0,1]
 * 输出: 0
 * 解释: 无需交换座位，所有的情侣都已经可以手牵手了。
 * 
 * 
 * 
 * 
 * 提示:
 * 
 * 
 * 2n == row.length
 * 2 <= n <= 30
 * n 是偶数
 * 0 <= row[i] < 2n
 * row 中所有元素均无重复
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/couples-holding-hands/solutions/603505/liang-chong-100-de-jie-fa-bing-cha-ji-ta-26a6/?envType=daily-question&envId=2023-11-11
//

/*
 * // 并查集. 就是统计有几个环.
 * public class Solution
 * {
 *     public int MinSwapsCouples(int[] row)
 *     {
 *         var n = row.Length >> 1;
 *         var pa = new int[n];
 *         int query(int p) => pa[p] == p ? p : (pa[p] = query(pa[p]));
 *         void merge(int p, int q) => pa[query(p)] = query(q);
 *         for (var i = 0; i < n; i++) { pa[i] = i; }
 *         for (var i = 0; i < n; i++) { merge(row[i << 1] >> 1, row[i << 1 | 1] >> 1); }
 *         return n - Enumerable
 *             .Range(0, n)
 *             .Count(i => i == query(i));
 *     }
 * }
 */

// 官解是并查集. 贪心有点不好写.
public class Solution
{
    public int MinSwapsCouples(int[] row)
    {
        int other(int k) => (k >> 1 << 2 | 1) - k;
        var n = row.Length >> 1;
        var ans = 0;
        var visits = new bool[n];
        var indexes = new int[n << 1];
        for (var i = 0; i < n << 1; i++) { indexes[row[i]] = i; }
        for (var i = 0; i < n; i++)
        {
            if (visits[i]) { continue; }
            for (var (p, q) = (indexes[i << 1], indexes[i << 1 | 1]);
                p >> 1 != q >> 1; ans++)
            {
                var s = other(p);
                (row[s], row[q], indexes[row[s]], indexes[row[q]])
                    = (row[q], row[s], indexes[row[q]], indexes[row[s]]);
                visits[row[p] >> 1] = true;
                (p, q) = (indexes[row[q]], indexes[other(row[q])]);
            }
        }
        return ans;
    }
}
