/*
 * @lc app=leetcode.cn id=count-anagrams lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6276] 统计同位异构字符串数目
 *
 * https://leetcode.cn/problems/count-anagrams/description/
 *
 * algorithms
 * Hard (35.16%)
 * Total Accepted:    813
 * Total Submissions: 2.3K
 * Testcase Example:  '"too hot"'
 *
 * 给你一个字符串 s ，它包含一个或者多个单词。单词之间用单个空格 ' ' 隔开。
 * 
 * 如果字符串 t 中第 i 个单词是 s 中第 i 个单词的一个 排列 ，那么我们称字符串 t 是字符串 s 的同位异构字符串。
 * 
 * 
 * 比方说，"acb dfe" 是 "abc def" 的同位异构字符串，但是 "def cab" 和 "adc bef" 不是。
 * 
 * 
 * 请你返回 s 的同位异构字符串的数目，由于答案可能很大，请你将它对 10^9 + 7 取余 后返回。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：s = "too hot"
 * 输出：18
 * 解释：输入字符串的一些同位异构字符串为 "too hot" ，"oot hot" ，"oto toh" ，"too toh" 以及 "too oht"
 * 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：s = "aa"
 * 输出：1
 * 解释：输入字符串只有一个同位异构字符串。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 10^5
 * s 只包含小写英文字母和空格 ' ' 。
 * 相邻单词之间由单个空格隔开。
 * 
 * 
 */

typedef long long ll;
const ll N = 1e5;
const ll MOD = 1e9 + 7;
const ll ALPHA = 26;
ll F[N + 10], IF[N + 10], cnt[ALPHA + 10];

ll POW(ll n, ll k) {
    ll ans = 1;
    for (; k; k >>= 1)
        ans = k & 1 ? (ans * n) % MOD : ans, n = (n * n) % MOD;
    return ans;
}

ll get(string s) {
    memset(cnt, 0, sizeof cnt);
    for (auto ch : s)
        cnt[ch - 'a']++;
    ll ans = F[s.size()];
    for (auto c : cnt)
        ans = (ans * IF[c]) % MOD;
    return ans;
}

class Solution {
public:
    int countAnagrams(string s) {
        F[0] = 1;
        for (ll i = 1; i <= N; i++)
            F[i] = (F[i - 1] * i) % MOD;
        IF[N] = POW(F[N], MOD - 2);
        for (ll i = N - 1; i >= 0; i--)
            IF[i] = (IF[i + 1] * (i + 1)) % MOD;
        ll tot = 1;
        stringstream ss;
        ss << s;
        while (ss >> s)
            tot = (tot * get(s)) % MOD;
        return tot;
    }
};
