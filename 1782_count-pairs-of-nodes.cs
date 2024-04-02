/*
 * @lc app=leetcode.cn id=count-pairs-of-nodes lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1782] 统计点对的数目
 *
 * https://leetcode.cn/problems/count-pairs-of-nodes/description/
 *
 * algorithms
 * Hard (36.55%)
 * Total Accepted:    2.4K
 * Total Submissions: 6.2K
 * Testcase Example:  '4\n[[1,2],[2,4],[1,3],[2,3],[2,1]]\n[2,3]'
 *
 * 给你一个无向图，无向图由整数 n  ，表示图中节点的数目，和 edges 组成，其中 edges[i] = [ui, vi] 表示 ui 和 vi
 * 之间有一条无向边。同时给你一个代表查询的整数数组 queries 。
 * 
 * 第 j 个查询的答案是满足如下条件的点对 (a, b) 的数目：
 * 
 * 
 * a < b
 * cnt 是与 a 或者 b 相连的边的数目，且 cnt 严格大于 queries[j] 。
 * 
 * 
 * 请你返回一个数组 answers ，其中 answers.length == queries.length 且 answers[j] 是第 j
 * 个查询的答案。
 * 
 * 请注意，图中可能会有 重复边 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 4, edges = [[1,2],[2,4],[1,3],[2,3],[2,1]], queries = [2,3]
 * 输出：[6,5]
 * 解释：每个点对中，与至少一个点相连的边的数目如上图所示。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 5, edges = [[1,5],[1,5],[3,4],[2,5],[1,3],[5,1],[2,3],[2,5]], queries
 * = [1,2,3,4,5]
 * 输出：[10,10,9,8,6]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 
 * 1 
 * 1 i, vi 
 * ui != vi
 * 1 
 * 0 
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/count-pairs-of-nodes/solutions/2398562/tong-ji-dian-dui-de-shu-mu-by-leetcode-s-yxlv/
//
// 想到容斥了. 甚至想到了枚举一个端点. 但 "先二分计算上限, 再减例外" 这思路还是挺厉害的.
// 看到这里有提到 CF 到了 1900, 也可以理解了.
// https://leetcode.cn/problems/count-pairs-of-nodes/solutions/1393385/by-v5qyy4q65w-pr2o/
public class Solution
{
    public int[] CountPairs(int n, int[][] edges, int[] queries)
    {
        var degrees = new int[n];
        var d = new Dictionary<(int, int), int>();
        foreach (var e in edges)
        {
            var (u, v) = (e[0] - 1, e[1] - 1);
            if (u > v) { (u, v) = (v, u); }
            degrees[u]++;
            degrees[v]++;
            d.TryAdd((u, v), 0);
            d[(u, v)]++;
        }
        var ordered = degrees.OrderBy(x => x).ToArray();
        return queries.Select(query =>
        {
            var total = 0;
            for (var u = 0; u < n; u++)
            {
                var (p, q) = (u + 1, n);
                while (p < q)
                {
                    var mid = (p + q) >> 1;
                    if (ordered[u] + ordered[mid] > query) { q = mid; }
                    else { p = mid + 1; }
                }
                total += n - p;
            }
            foreach (var kv in d)
            {
                var (u, v) = kv.Key;
                if (degrees[u] + degrees[v] > query
                    && degrees[u] + degrees[v] - d[(u, v)] <= query)
                {
                    total--;
                }
            }
            return total;
        }).ToArray();
    }
}
