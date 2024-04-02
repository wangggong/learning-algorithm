/*
 * @lc app=leetcode.cn id=maximum-earnings-from-taxi lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2008] 出租车的最大盈利
 *
 * https://leetcode.cn/problems/maximum-earnings-from-taxi/description/
 *
 * algorithms
 * Medium (47.78%)
 * Total Accepted:    8K
 * Total Submissions: 16.2K
 * Testcase Example:  '5\n[[2,5,4],[1,5,1]]'
 *
 * 你驾驶出租车行驶在一条有 n 个地点的路上。这 n 个地点从近到远编号为 1 到 n ，你想要从 1 开到 n
 * ，通过接乘客订单盈利。你只能沿着编号递增的方向前进，不能改变方向。
 * 
 * 乘客信息用一个下标从 0 开始的二维数组 rides 表示，其中 rides[i] = [starti, endi, tipi] 表示第 i
 * 位乘客需要从地点 starti 前往 endi ，愿意支付 tipi 元的小费。
 * 
 * 每一位 你选择接单的乘客 i ，你可以 盈利 endi - starti + tipi 元。你同时 最多 只能接一个订单。
 * 
 * 给你 n 和 rides ，请你返回在最优接单方案下，你能盈利 最多 多少元。
 * 
 * 注意：你可以在一个地点放下一位乘客，并在同一个地点接上另一位乘客。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：n = 5, rides = [[2,5,4],[1,5,1]]
 * 输出：7
 * 解释：我们可以接乘客 0 的订单，获得 5 - 2 + 4 = 7 元。
 * 
 * 
 * 示例 2：
 * 
 * 输入：n = 20, rides =
 * [[1,6,1],[3,10,2],[10,12,3],[11,12,2],[12,15,2],[13,18,1]]
 * 输出：20
 * 解释：我们可以接以下乘客的订单：
 * - 将乘客 1 从地点 3 送往地点 10 ，获得 10 - 3 + 2 = 9 元。
 * - 将乘客 2 从地点 10 送往地点 12 ，获得 12 - 10 + 3 = 5 元。
 * - 将乘客 5 从地点 13 送往地点 18 ，获得 18 - 13 + 1 = 6 元。
 * 我们总共获得 9 + 5 + 6 = 20 元。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 10^5
 * 1 <= rides.length <= 3 * 10^4
 * rides[i].length == 3
 * 1 <= starti < endi <= n
 * 1 <= tipi <= 10^5
 * 
 * 
 */

/*
 * // 树状数组
 * public class Solution
 * {
 *     public long MaxTaxiEarnings(int n, int[][] rides)
 *     {
 *         var tr = new long[n + 1];
 *         int lowBit(int k) => k & -k;
 *         long query(int k)
 *         {
 *             var ans = 0l;
 *             for (; k > 0; k -= lowBit(k)) { ans = Math.Max(ans, tr[k]); }
 *             return ans;
 *         }
 *         void update(int k, long v)
 *         {
 *             for (; k <= n; k += lowBit(k)) { tr[k] = Math.Max(tr[k], v); }
 *         }
 *         var ans = 0l;
 *         foreach (var ride in rides
 *             .OrderBy(r => r[0])
 *             .ThenBy(r => r[1]))
 *         {
 *             var (s, e, t) = (ride[0], ride[1], ride[2]);
 *             var cur = query(s) + (e - s + t);
 *             update(e, cur);
 *             ans = Math.Max(ans, cur);
 *         }
 *         return ans;
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/maximum-earnings-from-taxi/solutions/1005392/dong-tai-gui-hua-by-endlesscheng-fawr/?envType=daily-question&envId=2023-12-08
// 选或不选的动态规划.
public class Solution
{
    public long MaxTaxiEarnings(int n, int[][] rides)
    {
        var D = rides
            .GroupBy(x => x[1])
            .ToDictionary(g => g.Key, g => g);
        var dp = new long[n + 1];
        for (var i = 1; i <= n; i++)
        {
            dp[i] = dp[i - 1];
            if (!D.ContainsKey(i)) { continue; }
            foreach (var ride in D[i])
            {
                var (s, e, t) = (ride[0], ride[1], ride[2]);
                dp[i] = Math.Max(dp[i], dp[s] + (e - s + t));
            }
        }
        return dp.Max();
    }
}
