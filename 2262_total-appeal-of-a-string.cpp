/*
 * @lc app=leetcode.cn id=total-appeal-of-a-string lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2262] 字符串的总引力
 *
 * https://leetcode-cn.com/problems/total-appeal-of-a-string/description/
 *
 * algorithms
 * Hard (50.98%)
 * Total Accepted:    4.5K
 * Total Submissions: 8.4K
 * Testcase Example:  '"abbca"'
 *
 * 字符串的 引力 定义为：字符串中 不同 字符的数量。
 * 
 * 
 * 例如，"abbca" 的引力为 3 ，因为其中有 3 个不同字符 'a'、'b' 和 'c' 。
 * 
 * 
 * 给你一个字符串 s ，返回 其所有子字符串的总引力 。
 * 
 * 子字符串 定义为：字符串中的一个连续字符序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：s = "abbca"
 * 输出：28
 * 解释："abbca" 的子字符串有：
 * - 长度为 1 的子字符串："a"、"b"、"b"、"c"、"a" 的引力分别为 1、1、1、1、1，总和为 5 。
 * - 长度为 2 的子字符串："ab"、"bb"、"bc"、"ca" 的引力分别为 2、1、2、2 ，总和为 7 。
 * - 长度为 3 的子字符串："abb"、"bbc"、"bca" 的引力分别为 2、2、3 ，总和为 7 。
 * - 长度为 4 的子字符串："abbc"、"bbca" 的引力分别为 3、3 ，总和为 6 。
 * - 长度为 5 的子字符串："abbca" 的引力为 3 ，总和为 3 。
 * 引力总和为 5 + 7 + 7 + 6 + 3 = 28 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：s = "code"
 * 输出：20
 * 解释："code" 的子字符串有：
 * - 长度为 1 的子字符串："c"、"o"、"d"、"e" 的引力分别为 1、1、1、1 ，总和为 4 。
 * - 长度为 2 的子字符串："co"、"od"、"de" 的引力分别为 2、2、2 ，总和为 6 。
 * - 长度为 3 的子字符串："cod"、"ode" 的引力分别为 3、3 ，总和为 6 。
 * - 长度为 4 的子字符串："code" 的引力为 4 ，总和为 4 。
 * 引力总和为 4 + 6 + 6 + 4 = 20 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 10^5
 * s 由小写英文字母组成
 * 
 * 
 */
class Solution {
public:
    typedef long long LL;

    // 参考: https://leetcode-cn.com/problems/total-appeal-of-a-string/solution/by-endlesscheng-g405/
    //
    // 是一个 DP. 但状态转移需要捋明白:
    // - 如果当前字符没遇到过: `s[0..i]` 每个位置加上 `s[i+1]` 都相当于引力++, 因此直接加上 `i+1` 即可;
    // - 如果当前字符遇到过, 出现在 `j` 位置: 这里没分析出来, 其实是 `s[0..j]` 每个位置加上 `s[i+1]` 引力不变, 但 `s[j+1..i]` 的每个位置加上 `s[i+1]` 引力++, 因此需要加上 `(i+1) - (j+1)`;
    // 综上, 维护 `right` 数组后只需要加上 `i - right[s[i]]` 即可.
    LL appealSum(string s) {
        LL ans = 0, cur = 0;
        vector<int> right(26, -1);
        for (int i = 0; i < s.size(); i++) {
            cur += i - right[s[i] - 'a'];
            ans += cur;
            right[s[i] - 'a'] = i;
        }
        return ans;
    }
};
