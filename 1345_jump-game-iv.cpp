/*
 * @lc app=leetcode.cn id=jump-game-iv lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1345] 跳跃游戏 IV
 *
 * https://leetcode-cn.com/problems/jump-game-iv/description/
 *
 * algorithms
 * Hard (36.72%)
 * Total Accepted:    24.1K
 * Total Submissions: 51.3K
 * Testcase Example:  '[100,-23,-23,404,100,23,23,23,3,404]'
 *
 * 给你一个整数数组 arr ，你一开始在数组的第一个元素处（下标为 0）。
 * 
 * 每一步，你可以从下标 i 跳到下标：
 * 
 * 
 * i + 1 满足：i + 1 < arr.length
 * i - 1 满足：i - 1 >= 0
 * j 满足：arr[i] == arr[j] 且 i != j
 * 
 * 
 * 请你返回到达数组最后一个元素的下标处所需的 最少操作次数 。
 * 
 * 注意：任何时候你都不能跳到数组外面。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：arr = [100,-23,-23,404,100,23,23,23,3,404]
 * 输出：3
 * 解释：那你需要跳跃 3 次，下标依次为 0 --> 4 --> 3 --> 9 。下标 9 为数组的最后一个元素的下标。
 * 
 * 
 * 示例 2：
 * 
 * 输入：arr = [7]
 * 输出：0
 * 解释：一开始就在最后一个元素处，所以你不需要跳跃。
 * 
 * 
 * 示例 3：
 * 
 * 输入：arr = [7,6,9,6,9,6,9,7]
 * 输出：1
 * 解释：你可以直接从下标 0 处跳到下标 7 处，也就是数组的最后一个元素处。
 * 
 * 
 * 示例 4：
 * 
 * 输入：arr = [6,1,9]
 * 输出：2
 * 
 * 
 * 示例 5：
 * 
 * 输入：arr = [11,22,7,7,7,7,7,7,7,22,13]
 * 输出：3
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= arr.length <= 5 * 10^4
 * -10^8 <= arr[i] <= 10^8
 * 
 * 
 */
class Solution {
public:
  int minJumps(vector<int>& arr) {
    auto n = arr.size();
    // assert n > 0 && n <= 5e4;
    vector visit(n, false);
    unordered_map<int, vector<int>> valueMap;
    for (auto i = 0; i < n; i++) {
      valueMap[arr[i]].push_back(i);
    }
    queue<pair<int, int>> Q;
    Q.push(make_pair(0, 0)); visit[0] = true;
    while (!Q.empty()) {
      auto q = Q.front(); Q.pop();
      auto index = q.first, step = q.second;
      if (index == n-1) return step;
      auto next = index + 1;
      if (next < n && (!visit[next])) {
        if (next == n-1) return step + 1;
        Q.push(make_pair(next, step + 1)); visit[next] = true;
      }
      next = index - 1;
      if (next >= 0 && (!visit[next])) {
        Q.push(make_pair(next, step + 1)); visit[next] = true;
      }
      auto indexes = valueMap[arr[index]];
      for (auto next : indexes) {
        if (next == n-1) return step + 1;
        if (next == index || visit[next]) continue;
        Q.push(make_pair(next, step + 1)); visit[next] = true;
      }
      // NOTE: 存在大量重复 index 时, 子图构成了一个 **稠密的无向无权图**.
      // 常规的 BFS 会有一个 `O(n^2)` 的查询; 但实际上没必要.
      // 子图中有一个节点验证后其余节点直接跳过 (表现为把子图直接删了) 即可.
      /* indexes.clear(); 这个好像不太行 */ valueMap[arr[index]].clear();
    }
    // not reachable
    return -1;
  }
};
