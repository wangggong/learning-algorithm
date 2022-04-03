/*
 * @lc app=leetcode.cn id=escape-a-large-maze lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1036] 逃离大迷宫
 *
 * https://leetcode-cn.com/problems/escape-a-large-maze/description/
 *
 * algorithms
 * Hard (47.06%)
 * Total Accepted:    18.6K
 * Total Submissions: 39.5K
 * Testcase Example:  '[[0,1],[1,0]]\n[0,0]\n[0,2]'
 *
 * 在一个 10^6 x 10^6 的网格中，每个网格上方格的坐标为 (x, y) 。
 * 
 * 现在从源方格 source = [sx, sy] 开始出发，意图赶往目标方格 target = [tx, ty] 。数组 blocked
 * 是封锁的方格列表，其中每个 blocked[i] = [xi, yi] 表示坐标为 (xi, yi) 的方格是禁止通行的。
 * 
 * 每次移动，都可以走到网格中在四个方向上相邻的方格，只要该方格 不 在给出的封锁列表 blocked 上。同时，不允许走出网格。
 * 
 * 只有在可以通过一系列的移动从源方格 source 到达目标方格 target 时才返回 true。否则，返回 false。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：blocked = [[0,1],[1,0]], source = [0,0], target = [0,2]
 * 输出：false
 * 解释：
 * 从源方格无法到达目标方格，因为我们无法在网格中移动。
 * 无法向北或者向东移动是因为方格禁止通行。
 * 无法向南或者向西移动是因为不能走出网格。
 * 
 * 示例 2：
 * 
 * 
 * 输入：blocked = [], source = [0,0], target = [999999,999999]
 * 输出：true
 * 解释：
 * 因为没有方格被封锁，所以一定可以到达目标方格。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 <= blocked.length <= 200
 * blocked[i].length == 2
 * 0 i, yi < 10^6
 * source.length == target.length == 2
 * 0 x, sy, tx, ty < 10^6
 * source != target
 * 题目数据保证 source 和 target 不在封锁列表内
 * 
 * 
 */

const int MAXN = 1e6;

class Solution {
  int directions[4][2] = {
    {0, 1},
    {0, -1},
    {1, 0},
    {-1, 0}
  };
  string stringify(int x, int y) {
    return to_string(x) + " " + to_string(y);
  }

  bool check(unordered_map<string, bool>& M, int area, pair<int, int> s, pair<int, int> t) {
    unordered_map<string, bool> vis;
    queue<pair<int, int>> Q;
    Q.push(s);
    vis[stringify(s.first, s.second)] = true;
    while (!Q.empty() && vis.size() <= area) {
      auto q = Q.front(); Q.pop();
      if (q == t) return true;
      for (auto d : directions) {
        auto np = make_pair(q.first + d[0], q.second + d[1]);
        if (np.first < 0 || np.first >= MAXN || np.second < 0 || np.second >= MAXN) continue;
        auto str = stringify(np.first, np.second);
        if (M[str] || vis[str]) continue;
        Q.push(np);
        vis[str] = true;
      }
    }
    return vis.size() > area;
  }
public:
  bool isEscapePossible(vector<vector<int>>& blocked, vector<int>& s, vector<int>& t) {
    auto sp = make_pair(s[0], s[1]), tp = make_pair(t[0], t[1]);
    auto n = blocked.size();
    auto maxArea = n * (n-1) / 2;
    unordered_map<string, bool> M;
    for (auto b : blocked) {
      M[stringify(b[0], b[1])] = true;
    }
    return check(M, maxArea, sp, tp) && check(M, maxArea, tp, sp);
  }
};
