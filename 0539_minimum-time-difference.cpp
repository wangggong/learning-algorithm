/*
 * @lc app=leetcode.cn id=minimum-time-difference lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [539] 最小时间差
 *
 * https://leetcode-cn.com/problems/minimum-time-difference/description/
 *
 * algorithms
 * Medium (66.46%)
 * Total Accepted:    49.5K
 * Total Submissions: 74.3K
 * Testcase Example:  '["23:59","00:00"]'
 *
 * 给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：timePoints = ["23:59","00:00"]
 * 输出：1
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：timePoints = ["00:00","23:59","00:00"]
 * 输出：0
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= timePoints.length <= 2 * 10^4
 * timePoints[i] 格式为 "HH:MM"
 * 
 * 
 */
const int MAXN = 24 * 60;

class Solution {
  int atoi(string s) {
    int ans = 0;
    for (char c : s) {
      ans *= 10;
      ans += c - '0';
    }
    return ans;
  }

  int tp2min(string t) {
    auto h = t.substr(0, 2), m = t.substr(3);
    return atoi(h) * 60 + atoi(m);
  }
public:
  int findMinDifference(vector<string>& timePoints) {
    vector<int> count(MAXN, 0);
    for (auto t : timePoints) {
      count[tp2min(t)]++;
    }
    int ans = MAXN;
    int first = -1, last = -1;
    for (auto i = 0; i < MAXN; i++) {
      if (count[i] == 0) continue;
      if (count[i] > 1) return 0;
      if (first == -1) first = i;
      if (last >= 0) {
        ans = min(ans, i - last);
      }
      last = i;
    }
    ans = min(ans, first + MAXN - last);
    return ans;
  }
};
