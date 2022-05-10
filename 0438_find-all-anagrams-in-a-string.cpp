/*
 * @lc app=leetcode.cn id=find-all-anagrams-in-a-string lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [438] 找到字符串中所有字母异位词
 *
 * https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/description/
 *
 * algorithms
 * Medium (54.42%)
 * Total Accepted:    178.2K
 * Total Submissions: 327.3K
 * Testcase Example:  '"cbaebabacd"\n"abc"'
 *
 * 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
 * 
 * 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入: s = "cbaebabacd", p = "abc"
 * 输出: [0,6]
 * 解释:
 * 起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
 * 起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: s = "abab", p = "ab"
 * 输出: [0,1,2]
 * 解释:
 * 起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
 * 起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
 * 起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
 * 
 * 
 * 
 * 
 * 提示:
 * 
 * 
 * 1 <= s.length, p.length <= 3 * 10^4
 * s 和 p 仅包含小写字母
 * 
 * 
 */
class Solution {
    bool check(vector<int> &count) {
        for (auto c : count)
            if (c) return false;
        return true;
    }
public:
    vector<int> findAnagrams(string s, string p) {
        auto n = s.size(), m = p.size();
        vector<int> ans;
        if (m > n) return ans;
        vector<int> count(26, 0);
        for (int i = 0; i < m; i++) {
            count[p[i] - 'a']--;
        }
        for (int i = 0; i < n; i++) {
            if (check(count)) ans.push_back(i - m);
            if (i >= m) count[s[i - m] - 'a']--;
            count[s[i] - 'a']++;
        }
        if (check(count)) ans.push_back(n - m);
        return ans;
    }
};
