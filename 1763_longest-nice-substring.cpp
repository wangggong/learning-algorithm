/*
 * @lc app=leetcode.cn id=longest-nice-substring lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1763] 最长的美好子字符串
 *
 * https://leetcode-cn.com/problems/longest-nice-substring/description/
 *
 * algorithms
 * Easy (58.71%)
 * Total Accepted:    20.1K
 * Total Submissions: 28.7K
 * Testcase Example:  '"YazaAay"'
 *
 * 当一个字符串 s 包含的每一种字母的大写和小写形式 同时 出现在 s 中，就称这个字符串 s 是 美好 字符串。比方说，"abABB"
 * 是美好字符串，因为 'A' 和 'a' 同时出现了，且 'B' 和 'b' 也同时出现了。然而，"abA" 不是美好字符串因为 'b' 出现了，而
 * 'B' 没有出现。
 * 
 * 给你一个字符串 s ，请你返回 s 最长的 美好子字符串 。如果有多个答案，请你返回 最早
 * 出现的一个。如果不存在美好子字符串，请你返回一个空字符串。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "YazaAay"
 * 输出："aAa"
 * 解释："aAa" 是一个美好字符串，因为这个子串中仅含一种字母，其小写形式 'a' 和大写形式 'A' 也同时出现了。
 * "aAa" 是最长的美好子字符串。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "Bb"
 * 输出："Bb"
 * 解释："Bb" 是美好字符串，因为 'B' 和 'b' 都出现了。整个字符串也是原字符串的子字符串。
 * 
 * 示例 3：
 * 
 * 
 * 输入：s = "c"
 * 输出：""
 * 解释：没有美好子字符串。
 * 
 * 示例 4：
 * 
 * 
 * 输入：s = "dDzeE"
 * 输出："dD"
 * 解释："dD" 和 "eE" 都是最长美好子字符串。
 * 由于有多个美好子字符串，返回 "dD" ，因为它出现得最早。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 100
 * s 只包含大写和小写英文字母。
 * 
 * 
 */
class Solution {
  const int ALPHA_SIZE = 26;
public:
  /*
   * string longestNiceSubstring(string s) {
   *   string ans;
   *   auto n = s.size();
   *   bool visit[ALPHA_SIZE * 2];
   *   for (auto i = 0; i < n; i++) {
   *     memset(visit, false, sizeof(bool) * (ALPHA_SIZE * 2));
   *     for (auto j = i; j < n; j++) {
   *       auto k = ('a' <= s[j] && s[j] <= 'z') ? s[j] - 'a' : ALPHA_SIZE + s[j] - 'A';
   *       visit[k] = true;
   *       auto nice = true;
   *       for (auto i = 0; i < ALPHA_SIZE; i++) {
   *         if (visit[i] ^ visit[ALPHA_SIZE + i]) {
   *           nice = false;
   *           break;
   *         }
   *       }
   *       if (nice && j-i+1 > ans.size()) {
   *         ans = s.substr(i, j - i + 1);
   *       }
   *     }
   *   }
   *   return ans;
   * }
   */

  void dfs(string s, int p, int q, int &pos, int &len) {
    if (p >= q) {
      return;
    }
    auto lower = 0, upper = 0;
    for (auto i = p; i <= q; i++) {
      if (islower(s[i])) {
        lower |= 1 << (s[i] - 'a');
      } else {
        upper |= 1 << (s[i] - 'A');
      }
    }
    if (lower == upper) {
      if (q - p + 1 > len) {
        pos = p;
        len = q - p + 1;
      }
      return;
    }
    auto valid = lower & upper;
    auto idx = p;
    while (idx <= q) {
      p = idx;
      while (idx <= q && valid & (1 << (tolower(s[idx]) - 'a'))) {
        idx++;
      }
      dfs(s, p, idx - 1, pos, len);
      idx++;
    }
  }

  string longestNiceSubstring(string s) {
    auto mp = 0, ml = 0;
    dfs(s, 0, s.size() - 1, mp, ml);
    return s.substr(mp, ml);
  }
};
