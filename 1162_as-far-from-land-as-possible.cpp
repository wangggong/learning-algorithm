/*
 * @lc app=leetcode.cn id=as-far-from-land-as-possible lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1162] 地图分析
 *
 * https://leetcode-cn.com/problems/as-far-from-land-as-possible/description/
 *
 * algorithms
 * Medium (47.17%)
 * Total Accepted:    40.8K
 * Total Submissions: 86.3K
 * Testcase Example:  '[[1,0,1],[0,0,0],[1,0,1]]'
 *
 * 你现在手里有一份大小为 n x n 的 网格 grid，上面的每个 单元格 都用 0 和 1 标记好了。其中 0 代表海洋，1
 * 代表陆地，请你找出一个海洋单元格，这个海洋单元格到离它最近的陆地单元格的距离是最大的。如果网格上只有陆地或者海洋，请返回 -1。
 * 
 * 我们这里说的距离是「曼哈顿距离」（ Manhattan Distance）：(x0, y0) 和 (x1, y1) 这两个单元格之间的距离是 |x0 -
 * x1| + |y0 - y1| 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：grid = [[1,0,1],[0,0,0],[1,0,1]]
 * 输出：2
 * 解释： 
 * 海洋单元格 (1, 1) 和所有陆地单元格之间的距离都达到最大，最大距离为 2。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：grid = [[1,0,0],[0,0,0],[0,0,0]]
 * 输出：4
 * 解释： 
 * 海洋单元格 (2, 2) 和所有陆地单元格之间的距离都达到最大，最大距离为 4。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 
 * 
 * n == grid.length
 * n == grid[i].length
 * 1 <= n <= 100
 * grid[i][j] 不是 0 就是 1
 * 
 * 
 */
class Solution {
  int pos[4][2] = {
    {0, 1},
    {0, -1},
    {1, 0},
    {-1, 0},
  };
public:
  int maxDistance(vector<vector<int>>& M) {
    auto ans = -1;
    auto n = M.size();
    vector<vector<int>> D(n, vector<int>(n, -1));
    queue<pair<int, int>> Q;
    for (auto i = 0; i < n; i++) {
      for (auto j = 0; j < n; j++) {
        if (M[i][j]) {
          Q.push(make_pair(i, j));
          D[i][j] = 0;
        }
      }
    }
    while (!Q.empty()) {
      auto q = Q.front(); Q.pop();
      auto x = q.first, y = q.second;
      for (auto p : pos) {
        auto nx = x + p[0], ny = y + p[1];
        if (nx < 0 || nx >= n || ny < 0 || ny >= n) { continue; }
        if (M[nx][ny] || D[nx][ny] != -1) { continue; }
        D[nx][ny] = D[x][y] + 1;
        Q.push(make_pair(nx, ny));
        ans = max(ans, D[nx][ny]);
      }
    }
    return ans;
  }
};
