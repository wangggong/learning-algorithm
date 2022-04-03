/*
 * [166] 分数到小数  
 * 
 * @lc app=leetcode.cn id=fraction-to-recurring-decimal lang=cpp
 *
 * https://leetcode-cn.com/problems/fraction-to-recurring-decimal/description/
 * 
 * * algorithms
 * * Medium (33.17%)
 * * Source Code:       0166_fraction-to-recurring-decimal.cpp
 * * Total Accepted:    45.4K
 * * Total Submissions: 136.8K
 * * Testcase Example:  '1\n2'
 * 
 * 给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以 字符串形式返回小数 。
 * 
 * 如果小数部分为循环小数，则将循环的部分括在括号内。
 * 
 * 如果存在多个答案，只需返回 任意一个 。
 * 
 * 对于所有给定的输入，保证 答案字符串的长度小于 10^4 。
 * 
 *  
 * 
 * 示例 1：
 * 
 * 
 * 输入：numerator = 1, denominator = 2
 * 输出："0.5"
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：numerator = 2, denominator = 1
 * 输出："2"
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：numerator = 2, denominator = 3
 * 输出："0.(6)"
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：numerator = 4, denominator = 333
 * 输出："0.(012)"
 * 
 * 
 * 示例 5：
 * 
 * 
 * 输入：numerator = 1, denominator = 5
 * 输出："0.2"
 * 
 * 
 *  
 * 
 * 提示：
 * 
 * 
 *         -2^31 <= numerator, denominator <= 2^31 - 1
 *         denominator != 0
 *
 */
class Solution {
public:
  string fractionToDecimal(int num, int den) {
    long n = num;
    long d = den;
    if (n % d == 0) {
      return to_string(n / d);
    }

    string s;
    if ((n > 0) != (d > 0)) {
      s.push_back('-');
    }
    n = abs(n);
    d = abs(d);
    auto q = n / d;
    s += to_string(q);
    s.push_back('.');
    string fp;
    unordered_map<long, int> IM;
    long r = n % d;
    int idx = 0;
    while (r != 0 && !IM.count(r)) {
      IM[r] = idx;
      r *= 10;
      fp += to_string(r / d);
      r %= d;
      idx++;
    }
    if (r != 0) {
      auto i = IM[r];
      fp = fp.substr(0, i) + "(" + fp.substr(i) + ")";
    }
    s += fp;
    return s;
  }
};
