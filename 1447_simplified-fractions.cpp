/*
 * @lc app=leetcode.cn id=simplified-fractions lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1447] 最简分数
 *
 * https://leetcode-cn.com/problems/simplified-fractions/description/
 *
 * algorithms
 * Medium (61.29%)
 * Total Accepted:    14.5K
 * Total Submissions: 21.9K
 * Testcase Example:  '2'
 *
 * 给你一个整数 n ，请你返回所有 0 到 1 之间（不包括 0 和 1）满足分母小于等于  n 的 最简 分数 。分数可以以 任意 顺序返回。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：n = 2
 * 输出：["1/2"]
 * 解释："1/2" 是唯一一个分母小于等于 2 的最简分数。
 * 
 * 示例 2：
 * 
 * 输入：n = 3
 * 输出：["1/2","1/3","2/3"]
 * 
 * 
 * 示例 3：
 * 
 * 输入：n = 4
 * 输出：["1/2","1/3","1/4","2/3","3/4"]
 * 解释："2/4" 不是最简分数，因为它可以化简为 "1/2" 。
 * 
 * 示例 4：
 * 
 * 输入：n = 1
 * 输出：[]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 100
 * 
 * 
 */

class Solution {
  int gcd(int x, int y) {
    return y == 0 ? x : gcd(y, x%y);
  }
public:
  vector<string> simplifiedFractions(int n) {
    // assert n >= 1 && n <= 100;
    vector<string> ans;
    for (auto den = 2; den <= n; den++) {
      for (auto num = 1; num < den; num++) {
        if (num > 1 && gcd(num, den) > 1) continue;
        ans.push_back(to_string(num) + "/" + to_string(den));
      }
    }
    return ans;
  }
};
