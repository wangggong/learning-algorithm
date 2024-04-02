/*
 * @lc app=leetcode.cn id=rank-transform-of-a-matrix lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1632] 矩阵转换后的秩
 *
 * https://leetcode.cn/problems/rank-transform-of-a-matrix/description/
 *
 * algorithms
 * Hard (35.13%)
 * Total Accepted:    5.1K
 * Total Submissions: 10.3K
 * Testcase Example:  '[[1,2],[3,4]]'
 *
 * 给你一个 m x n 的矩阵 matrix ，请你返回一个新的矩阵 answer ，其中 answer[row][col] 是
 * matrix[row][col] 的秩。
 * 
 * 每个元素的 秩 是一个整数，表示这个元素相对于其他元素的大小关系，它按照如下规则计算：
 * 
 * 
 * 秩是从 1 开始的一个整数。
 * 如果两个元素 p 和 q 在 同一行 或者 同一列 ，那么：
 * 
 * 如果 p < q ，那么 rank(p) < rank(q)
 * 如果 p == q ，那么 rank(p) == rank(q)
 * 如果 p > q ，那么 rank(p) > rank(q)
 * 
 * 
 * 秩 需要越 小 越好。
 * 
 * 
 * 题目保证按照上面规则 answer 数组是唯一的。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：matrix = [[1,2],[3,4]]
 * 输出：[[1,2],[2,3]]
 * 解释：
 * matrix[0][0] 的秩为 1 ，因为它是所在行和列的最小整数。
 * matrix[0][1] 的秩为 2 ，因为 matrix[0][1] > matrix[0][0] 且 matrix[0][0] 的秩为 1 。
 * matrix[1][0] 的秩为 2 ，因为 matrix[1][0] > matrix[0][0] 且 matrix[0][0] 的秩为 1 。
 * matrix[1][1] 的秩为 3 ，因为 matrix[1][1] > matrix[0][1]， matrix[1][1] >
 * matrix[1][0] 且 matrix[0][1] 和 matrix[1][0] 的秩都为 2 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：matrix = [[7,7],[7,7]]
 * 输出：[[1,1],[1,1]]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：matrix = [[20,-21,14],[-19,4,19],[22,-47,24],[-19,4,19]]
 * 输出：[[4,2,3],[1,3,4],[5,1,6],[1,3,4]]
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：matrix = [[7,3,6],[1,4,5],[9,8,2]]
 * 输出：[[5,1,4],[1,2,3],[6,3,1]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == matrix.length
 * n == matrix[i].length
 * 1 
 * -10^9 
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/rank-transform-of-a-matrix/solution/by-lcbin-ll29/
// 从去重后结果开始分析 (贪心), 把所有相同值统计在一起, 再用并查集把其中秩相同的位置连起来统一赋值.
public class Solution
{
    public int[][] MatrixRankTransform(int[][] M)
    {
        var n = M.Length;
        var m = M[0].Length;
        var rowRate = new int[n];
        var colRate = new int[m];
        var d = new Dictionary<int, List<(int, int)>>();
        for (var i = 0; i < n; i++)
        {
            for (var j = 0; j < m; j++)
            {
                if (!d.ContainsKey(M[i][j]))
                {
                    d.Add(M[i][j], new());
                }
                d[M[i][j]].Add((i, j));
            }
        }
        foreach (var (k, list) in d
                .Select(kv => (kv.Key, kv.Value))
                .OrderBy(x => x.Item1))
        {
            var pa = new int[n + m];
            var rank = new Dictionary<int, int>();
            for (var i = 0; i < n + m; i++)
            {
                pa[i] = i;
            }
            int query(int k) => k == pa[k] ? pa[k] : (pa[k] = query(pa[k]));
            void merge(int p, int q) => pa[query(p)] = query(q);
            foreach (var (i, j) in list)
            {
                merge(i, j + n);
            }
            foreach (var (i, j) in list)
            {
                rank.TryGetValue(query(i), out var v);
                rank[query(i)] = Math.Max(v, Math.Max(rowRate[i], colRate[j]));
            }
            foreach (var (i, j) in list)
            {
                rank.TryGetValue(query(i), out var v);
                rowRate[i] = v + 1;
                colRate[j] = v + 1;
                M[i][j] = v + 1;
            }
        }
        return M;
    }
}
