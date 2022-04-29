/*
 * @lc app=leetcode.cn id=pacific-atlantic-water-flow lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [417] 太平洋大西洋水流问题
 *
 * https://leetcode-cn.com/problems/pacific-atlantic-water-flow/description/
 *
 * algorithms
 * Medium (50.21%)
 * Total Accepted:    42.3K
 * Total Submissions: 81.6K
 * Testcase Example:  '[[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]'
 *
 * 有一个 m × n 的矩形岛屿，与 太平洋 和 大西洋 相邻。 “太平洋” 处于大陆的左边界和上边界，而 “大西洋” 处于大陆的右边界和下边界。
 * 
 * 这个岛被分割成一个由若干方形单元格组成的网格。给定一个 m x n 的整数矩阵 heights ， heights[r][c] 表示坐标 (r, c)
 * 上单元格 高于海平面的高度 。
 * 
 * 岛上雨水较多，如果相邻单元格的高度 小于或等于 当前单元格的高度，雨水可以直接向北、南、东、西流向相邻单元格。水可以从海洋附近的任何单元格流入海洋。
 * 
 * 返回 网格坐标 result 的 2D列表 ，其中 result[i] = [ri, ci] 表示雨水可以从单元格 (ri, ci) 流向
 * 太平洋和大西洋 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入: heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
 * 输出: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入: heights = [[2,1],[1,2]]
 * 输出: [[0,0],[0,1],[1,0],[1,1]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == heights.length
 * n == heights[r].length
 * 1 <= m, n <= 200
 * 0 <= heights[r][c] <= 10^5
 * 
 * 
 */
typedef pair<int, int> PII;

const int pacificMask = 1 << 0;
const int atlanticMask = 1 << 1;

class Solution {
    vector<PII> directions = {
        {0, 1},
        {0, -1},
        {1, 0},
        {-1, 0},
    };
    void bfs(queue<PII> &Q, const vector<vector<int>> &heights, vector<vector<int>> &visit, int mask) {
        auto n = heights.size(), m = heights[0].size();
        while (!Q.empty()) {
            auto q = Q.front(); Q.pop();
            auto x = q.second, y = q.first;
            for (auto d : directions) {
                auto dx = d.second, dy = d.first;
                auto nx = x + dx, ny = y + dy;
                if (nx < 0 || nx >= m || ny < 0 || ny >= n) continue;
                if (!(visit[ny][nx] & mask) && heights[y][x] <= heights[ny][nx])
                    Q.push({ny, nx}), visit[ny][nx] |= mask;
            }
        }
        return;
    }

public:
    vector<vector<int>> pacificAtlantic(vector<vector<int>>& heights) {
        auto n = heights.size(), m = heights[0].size();
        vector<vector<int>> visit(n, vector<int>(m, 0));
        queue<PII> Q;
        for (auto i = 0; i < n; i++) Q.push({i, 0}), visit[i][0] |= pacificMask;
        for (auto i = 0; i < m; i++) Q.push({0, i}), visit[0][i] |= pacificMask;
        bfs(Q, heights, visit, pacificMask);
        for (auto i = 0; i < n; i++) Q.push({i, m - 1}), visit[i][m - 1] |= atlanticMask;
        for (auto i = 0; i < m; i++) Q.push({n - 1, i}), visit[n - 1][i] |= atlanticMask;
        bfs(Q, heights, visit, atlanticMask);
        vector<vector<int>> ans;
        for (auto i = 0; i < n; i++)
            for (auto j = 0; j < m; j++)
                if ((visit[i][j] & pacificMask) && (visit[i][j] & atlanticMask)) ans.push_back({i, j});
        return ans;
    }
};
