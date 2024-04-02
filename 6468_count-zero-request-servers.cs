/*
 * @lc app=leetcode.cn id=count-zero-request-servers lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6468] 统计没有收到请求的服务器数目
 *
 * https://leetcode.cn/problems/count-zero-request-servers/description/
 *
 * algorithms
 * Medium (30.29%)
 * Total Accepted:    600
 * Total Submissions: 2K
 * Testcase Example:  '3\n[[1,3],[2,6],[1,5]]\n5\n[10,11]'
 *
 * 给你一个整数 n ，表示服务器的总数目，再给你一个下标从 0 开始的 二维 整数数组 logs ，其中 logs[i] = [server_id,
 * time] 表示 id 为 server_id 的服务器在 time 时收到了一个请求。
 * 
 * 同时给你一个整数 x 和一个下标从 0 开始的整数数组 queries  。
 * 
 * 请你返回一个长度等于 queries.length 的数组 arr ，其中 arr[i] 表示在时间区间 [queries[i] - x,
 * queries[i]] 内没有收到请求的服务器数目。
 * 
 * 注意时间区间是个闭区间。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 3, logs = [[1,3],[2,6],[1,5]], x = 5, queries = [10,11]
 * 输出：[1,2]
 * 解释：
 * 对于 queries[0]：id 为 1 和 2 的服务器在区间 [5, 10] 内收到了请求，所以只有服务器 3 没有收到请求。
 * 对于 queries[1]：id 为 2 的服务器在区间 [6,11] 内收到了请求，所以 id 为 1 和 3
 * 的服务器在这个时间段内没有收到请求。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 3, logs = [[2,4],[2,1],[1,2],[3,1]], x = 2, queries = [3,4]
 * 输出：[0,1]
 * 解释：
 * 对于 queries[0]：区间 [1, 3] 内所有服务器都收到了请求。
 * 对于 queries[1]：只有 id 为 3 的服务器在区间 [2,4] 内没有收到请求。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 10^5
 * 1 <= logs.length <= 10^5
 * 1 <= queries.length <= 10^5
 * logs[i].length == 2
 * 1 <= logs[i][0] <= n
 * 1 <= logs[i][1] <= 10^6
 * 1 <= x <= 10^5
 * x < queries[i] <= 10^6
 * 
 * 
 */
public class Solution
{
    public int[] CountServers(int n, int[][] logs, int x, int[] queries)
    {
        var ans = new int[queries.Length];
        logs = logs
            .OrderBy(x => x[1])
            .ToArray();
        var (p, q, m) = (0, 0, logs.Length);
        var counts = new Dictionary<int, int>();
        foreach (var (t, i) in queries
            .Select((t, i) => (t, i))
            .OrderBy(x => x.t))
        {
            for (; p < m && logs[p][1] <= t; p++)
            {
                var id = logs[p][0];
                counts.TryGetValue(id, out var c);
                counts[id] = c + 1;
            }
            for (; q < m && logs[q][1] < t - x; q++)
            {
                var id = logs[q][0];
                counts[id]--;
                if (counts[id] is 0)
                {
                    counts.Remove(id);
                }
            }
            ans[i] = n - counts.Count();
        }
        return ans;
    }
}
