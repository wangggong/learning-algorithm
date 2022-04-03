/*
 * @lc app=leetcode.cn id=map-of-highest-peak lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1765] 地图中的最高点
 *
 * https://leetcode-cn.com/problems/map-of-highest-peak/description/
 *
 * algorithms
 * Medium (54.91%)
 * Total Accepted:    20.5K
 * Total Submissions: 30.8K
 * Testcase Example:  '[[0,1],[0,0]]'
 *
 * 给你一个大小为 m x n 的整数矩阵 isWater ，它代表了一个由 陆地 和 水域 单元格组成的地图。
 * 
 * 
 * 如果 isWater[i][j] == 0 ，格子 (i, j) 是一个 陆地 格子。
 * 如果 isWater[i][j] == 1 ，格子 (i, j) 是一个 水域 格子。
 * 
 * 
 * 你需要按照如下规则给每个单元格安排高度：
 * 
 * 
 * 每个格子的高度都必须是非负的。
 * 如果一个格子是是 水域 ，那么它的高度必须为 0 。
 * 任意相邻的格子高度差 至多 为 1 。当两个格子在正东、南、西、北方向上相互紧挨着，就称它们为相邻的格子。（也就是说它们有一条公共边）
 * 
 * 
 * 找到一种安排高度的方案，使得矩阵中的最高高度值 最大 。
 * 
 * 请你返回一个大小为 m x n 的整数矩阵 height ，其中 height[i][j] 是格子 (i, j) 的高度。如果有多种解法，请返回
 * 任意一个 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：isWater = [[0,1],[0,0]]
 * 输出：[[1,0],[2,1]]
 * 解释：上图展示了给各个格子安排的高度。
 * 蓝色格子是水域格，绿色格子是陆地格。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：isWater = [[0,0,1],[1,0,0],[0,0,0]]
 * 输出：[[1,1,0],[0,1,1],[1,2,2]]
 * 解释：所有安排方案中，最高可行高度为 2 。
 * 任意安排方案中，只要最高高度为 2 且符合上述规则的，都为可行方案。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == isWater.length
 * n == isWater[i].length
 * 1 <= m, n <= 1000
 * isWater[i][j] 要么是 0 ，要么是 1 。
 * 至少有 1 个水域格子。
 * 
 * 
 */
class Solution {
  int pos[4][2] = {
    {0, 1},
    {0, -1},
    {1, 0},
    {-1, 0}
  };
public:
  /*
  vector<vector<int>> highestPeak(vector<vector<int>>& M) {
    auto n = M.size(), m = M[0].size();
    vector<vector<int>> ans(n, vector<int>(m, -1));
    queue<pair<int, int>> Q;
    for (auto i = 0; i < n; i++) {
      for (auto j = 0; j < m; j++) {
        if (M[i][j]) {
          Q.push(make_pair(i, j));
          ans[i][j] = 0;
        }
      }
    }
    while (!Q.empty()) {
      auto q = Q.front(); Q.pop();
      for (auto p : pos) {
        auto nx = q.first + p[0], ny = q.second + p[1];
        if (nx < 0 || nx >= n || ny < 0 || ny >= m) {
          continue;
        }
        if (ans[nx][ny] != -1) {
          continue;
        }
        ans[nx][ny] = ans[q.first][q.second] + 1;
        Q.push(make_pair(nx, ny));
      }
    }
    return ans;
  }
  */
  vector<vector<int>> highestPeak(vector<vector<int>>& M) {
    auto n = M.size(), m = M[0].size();
    vector<vector<int>> ans(n, vector<int>(m, -1));
    queue<pair<pair<int, int>, int>> Q;
    for (auto i = 0; i < n; i++) {
      for (auto j = 0; j < m; j++) {
        if (M[i][j]) {
          Q.push(make_pair(make_pair(i, j), 0));
        }
      }
    }
    while (!Q.empty()) {
      auto q = Q.front(); Q.pop();
      auto i = q.first.first, j = q.first.second, v = q.second;
      if (ans[i][j] != -1) { continue; }
      ans[i][j] = v;
      for (auto p : pos) {
        auto ni = i + p[0], nj = j + p[1];
        if (ni < 0 || ni >= n || nj < 0 || nj >= m) { continue; }
        // if (ans[ni][nj] != -1) { continue; }
        Q.push(make_pair(make_pair(ni, nj), v + 1));
      }
    }
    return ans;
  }
};
