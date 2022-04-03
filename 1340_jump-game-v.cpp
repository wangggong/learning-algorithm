/*
 * @lc app=leetcode.cn id=jump-game-v lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1340] 跳跃游戏 V
 *
 * https://leetcode-cn.com/problems/jump-game-v/description/
 *
 * algorithms
 * Hard (57.54%)
 * Total Accepted:    4.9K
 * Total Submissions: 8.4K
 * Testcase Example:  '[6,4,14,6,8,13,9,7,10,6,12]\n2'
 *
 * 给你一个整数数组 arr 和一个整数 d 。每一步你可以从下标 i 跳到：
 * 
 * 
 * i + x ，其中 i + x < arr.length 且 0 < x <= d 。
 * i - x ，其中 i - x >= 0 且 0 < x <= d 。
 * 
 * 
 * 除此以外，你从下标 i 跳到下标 j 需要满足：arr[i] > arr[j] 且 arr[i] > arr[k] ，其中下标 k 是所有 i 到 j
 * 之间的数字（更正式的，min(i, j) < k < max(i, j)）。
 * 
 * 你可以选择数组的任意下标开始跳跃。请你返回你 最多 可以访问多少个下标。
 * 
 * 请注意，任何时刻你都不能跳到数组的外面。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：arr = [6,4,14,6,8,13,9,7,10,6,12], d = 2
 * 输出：4
 * 解释：你可以从下标 10 出发，然后如上图依次经过 10 --> 8 --> 6 --> 7 。
 * 注意，如果你从下标 6 开始，你只能跳到下标 7 处。你不能跳到下标 5 处因为 13 > 9 。你也不能跳到下标 4 处，因为下标 5 在下标 4 和
 * 6 之间且 13 > 9 。
 * 类似的，你不能从下标 3 处跳到下标 2 或者下标 1 处。
 * 
 * 
 * 示例 2：
 * 
 * 输入：arr = [3,3,3,3,3], d = 3
 * 输出：1
 * 解释：你可以从任意下标处开始且你永远无法跳到任何其他坐标。
 * 
 * 
 * 示例 3：
 * 
 * 输入：arr = [7,6,5,4,3,2,1], d = 1
 * 输出：7
 * 解释：从下标 0 处开始，你可以按照数值从大到小，访问所有的下标。
 * 
 * 
 * 示例 4：
 * 
 * 输入：arr = [7,1,7,1,7,1], d = 2
 * 输出：2
 * 
 * 
 * 示例 5：
 * 
 * 输入：arr = [66], d = 1
 * 输出：1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= arr.length <= 1000
 * 1 <= arr[i] <= 10^5
 * 1 <= d <= arr.length
 * 
 * 
 */
class Solution {
  int dfs(const vector<int>& arr, vector<int>& dp, int d, int index, int n) {
    // assert n > 1 && n < 1e3;
    if (dp[index] >= 0) return dp[index];
    if (
        (index == 0 && arr[index] <= arr[index+1]) ||
        (index == n-1 && arr[index] <= arr[index-1]) ||
        (index > 0 && index < n-1 && arr[index] <= arr[index+1] && arr[index] <= arr[index-1])
       ) {
      dp[index] = 1;
      return dp[index];
    }
    for (auto i = 1; i <= d; i++) {
      if (index+i >= n || arr[index] <= arr[index+i]) break;
      dp[index] = max(dp[index], dfs(arr, dp, d, index+i, n) + 1);
    }
    for (auto i = 1; i <= d; i++) {
      if (index-i < 0 || arr[index] <= arr[index-i]) break;
      dp[index] = max(dp[index], dfs(arr, dp, d, index-i, n) + 1);
    }
    return dp[index];
  }
public:
  int maxJumps(vector<int>& arr, int d) {
    auto n = arr.size();
    if (n == 1) return n;
    // assert n > 0 && n < 1e3;
    // assert d > 0 && d <= n;
    vector<int> dp(n, -1);
    auto ans = -1;
    for (auto i = 0; i < n; i++) ans = max(ans, dfs(arr, dp, d, i, n));
    return ans;
  }
};
