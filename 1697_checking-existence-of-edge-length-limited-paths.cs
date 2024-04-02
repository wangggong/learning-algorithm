/*
 * @lc app=leetcode.cn id=checking-existence-of-edge-length-limited-paths lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1697] 检查边长度限制的路径是否存在
 *
 * https://leetcode.cn/problems/checking-existence-of-edge-length-limited-paths/description/
 *
 * algorithms
 * Hard (55.09%)
 * Total Accepted:    5.2K
 * Total Submissions: 8.5K
 * Testcase Example:  '3\n[[0,1,2],[1,2,4],[2,0,8],[1,0,16]]\n[[0,1,2],[0,2,5]]'
 *
 * 给你一个 n 个点组成的无向图边集 edgeList ，其中 edgeList[i] = [ui, vi, disi] 表示点 ui 和点 vi
 * 之间有一条长度为 disi 的边。请注意，两个点之间可能有 超过一条边 。
 * 
 * 给你一个查询数组queries ，其中 queries[j] = [pj, qj, limitj] ，你的任务是对于每个查询 queries[j]
 * ，判断是否存在从 pj 到 qj 的路径，且这条路径上的每一条边都 严格小于 limitj 。
 * 
 * 请你返回一个 布尔数组 answer ，其中 answer.length == queries.length ，当 queries[j] 的查询结果为
 * true 时， answer 第 j 个值为 true ，否则为 false 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 3, edgeList = [[0,1,2],[1,2,4],[2,0,8],[1,0,16]], queries =
 * [[0,1,2],[0,2,5]]
 * 输出：[false,true]
 * 解释：上图为给定的输入数据。注意到 0 和 1 之间有两条重边，分别为 2 和 16 。
 * 对于第一个查询，0 和 1 之间没有小于 2 的边，所以我们返回 false 。
 * 对于第二个查询，有一条路径（0 -> 1 -> 2）两条边都小于 5 ，所以这个查询我们返回 true 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 5, edgeList = [[0,1,10],[1,2,5],[2,3,9],[3,4,13]], queries =
 * [[0,4,14],[1,4,13]]
 * 输出：[true,false]
 * 解释：上图为给定数据。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 
 * 1 
 * edgeList[i].length == 3
 * queries[j].length == 3
 * 0 i, vi, pj, qj 
 * ui != vi
 * pj != qj
 * 1 i, limitj 
 * 两个点之间可能有 多条 边。
 * 
 * 
 */
public class Solution
{
    public bool[] DistanceLimitedPathsExist(int n, int[][] edgeList, int[][] queries)
    {
        var pa = new int[n];
        for (var i = 0; i < n; i++) { pa[i] = i; }
        int query(int k) => k != pa[k] ? (pa[k] = query(pa[k])) : pa[k];
        void merge(int p, int q) => pa[query(p)] = query(q);
        Array.Sort(edgeList, (x, y) => x[2].CompareTo(y[2]));
        var qn = queries.Length;
        var queryInfos = new int[qn][];
        var ans = new bool[qn];
        for (var i = 0; i < qn; i++)
        {
            queryInfos[i] = new int[] {i, queries[i][0], queries[i][1], queries[i][2]};
        }
        Array.Sort(queryInfos, (x, y) => x[3].CompareTo(y[3]));
        var cur = 0;
        foreach (var q in queryInfos)
        {
            for (var m = edgeList.Length; cur < m && edgeList[cur][2] < q[3]; cur++)
            {
                merge(edgeList[cur][0], edgeList[cur][1]);
            }
            ans[q[0]] = (query(q[1]) == query(q[2]));
        }
        return ans;
    }
}
