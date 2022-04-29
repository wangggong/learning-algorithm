/*
 * @lc app=leetcode.cn id=sum-of-scores-of-built-strings lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2223] 构造字符串的总得分和
 *
 * https://leetcode-cn.com/problems/sum-of-scores-of-built-strings/description/
 *
 * algorithms
 * Hard (32.66%)
 * Total Accepted:    2.4K
 * Total Submissions: 7.2K
 * Testcase Example:  '"babab"'
 *
 * 你需要从空字符串开始 构造 一个长度为 n 的字符串 s ，构造的过程为每次给当前字符串 前面 添加 一个 字符。构造过程中得到的所有字符串编号为 1
 * 到 n ，其中长度为 i 的字符串编号为 si 。
 * 
 * 
 * 比方说，s = "abaca" ，s1 == "a" ，s2 == "ca" ，s3 == "aca" 依次类推。
 * 
 * 
 * si 的 得分 为 si 和 sn 的 最长公共前缀 的长度（注意 s == sn ）。
 * 
 * 给你最终的字符串 s ，请你返回每一个 si 的 得分之和 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "babab"
 * 输出：9
 * 解释：
 * s1 == "b" ，最长公共前缀是 "b" ，得分为 1 。
 * s2 == "ab" ，没有公共前缀，得分为 0 。
 * s3 == "bab" ，最长公共前缀为 "bab" ，得分为 3 。
 * s4 == "abab" ，没有公共前缀，得分为 0 。
 * s5 == "babab" ，最长公共前缀为 "babab" ，得分为 5 。
 * 得分和为 1 + 0 + 3 + 0 + 5 = 9 ，所以我们返回 9 。
 * 
 * 示例 2 ：
 * 
 * 
 * 输入：s = "azbazbzaz"
 * 输出：14
 * 解释：
 * s2 == "az" ，最长公共前缀为 "az" ，得分为 2 。
 * s6 == "azbzaz" ，最长公共前缀为 "azb" ，得分为 3 。
 * s9 == "azbazbzaz" ，最长公共前缀为 "azbazbzaz" ，得分为 9 。
 * 其他 si 得分均为 0 。
 * 得分和为 2 + 3 + 9 = 14 ，所以我们返回 14 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 10^5
 * s 只包含小写英文字母。
 * 
 * 
 */
const int mod = 1e9 + 7;
const int base = 37;

class Solution {
    /*
     * // 暴力, `O(n^2)`.
     * vector<int> zfunc(string s) {
     *     auto n = s.size();
     *     vector<int> z(n, 0);
     *     for (auto i = 0; i < n; i++)
     *         while (i + z[i] < n && s[i + z[i]] == s[z[i]]) z[i]++;
     *     return z;
     * }
     */

    // 基于 DFA 的 `O(n)` 解.
    //
    // 参考:
    // - https://oi-wiki.org/string/z-func/
    // - https://personal.utdallas.edu/~besp/demo/John2010/z-algorithm.htm
    vector<int> zfunc(string s) {
        auto n = s.size();
        vector<int> z(n, 0);
        if (n == 0) return z;
        // 特殊情况: `0` 位置一定能做到完全匹配.
        z[0] = n;
        // 维护一个当前匹配的最大窗口.
        // 这个窗口始终满足以下特点:
        // - `l <= i`
        // - `s[l:r] == s[0:r - l]`
        int l = 0, r = 1;
        for (auto i = 1; i < n; i++) {
            // - 如果 `i >= r` (当前窗口不包含当前位置), 暴力匹配.
            // - 否则, **`s[i:r] == s[i - l:r - l]`**
            //   - 如果 `z[i - l] < r - i`, 说明 `s[i - l:r - l]` 与 `s[0:r - i]` 是失配的. 直接返回 `z[i - l]` 即可.
            //   - 否则, 需要从 `r - i` 开始暴力匹配 (但好像不存在? 反证的). `s[i:r]` 到 `s[i - l:r - l]` 是匹配的, 不用走了.
            if (i < r && z[i - l] < r - i)
                z[i] = z[i - l];
            else
                for (z[i] = max(0, r - i); i + z[i] < n && s[i + z[i]] == s[z[i]]; z[i]++);
            if (i + z[i] > r) l = i, r = i + z[i];
        }
        return z;
    }
public:
    /*
     * // 扩展 KMP (Z 函数)
     * long long sumScores(string s) {
     *     long long ans = 0;
     *     auto zf = zfunc(s);
     *     for (auto z : zf) ans += z;
     *     return ans;
     * }
     */

    long long sumScores(string s) {
        auto n = s.size();
        if (n == 0) return 0;
        // 预处理 prefix & mul
        vector<long long> prefix(n + 1, 0), mul(n + 1, 0);
        mul[0] = 1;
        for (auto i = 0; i < n; i++) {
            prefix[i + 1] = (prefix[i] * base + (s[i] - 'a' + 1)) % mod;
            mul[i + 1] = (mul[i] * base) % mod;
        }
        long long ans = n;
        // 枚举每个后缀串长度...
        for (auto i = 1; i < n; i++) {
            // 如果首字母都不一样直接 pass...
            if (s[n - i] != s[0]) continue;
            int p = 0, q = i;
            while (p < q) {
                auto mid = (p + q + 1) >> 1;
                // 这里:
                //
                // mid 枚举的是每个公共前缀的 **长度**, 我们有以下公式 (参考 Horner):
                // - `encode(arr[:q]) = encode(arr[:q-1]) * base + int(arr[q] - 'a' + 1)`
                // - `encode(arr[p:q]) = encode(arr[:q]) - encode(arr[:p]) * base^(q-p)`
                // 类似前缀和的结论.
                // 至于后面为啥又加了一个 mod, 大概是怕出负数吧 (
                auto h = ((prefix[n - i + mid] - prefix[n - i] * mul[mid] % mod) % mod + mod) % mod;
                if (h == prefix[mid]) p = mid;
                else q = mid - 1;
            }
            ans += p;
        }
        return ans;
    }
};
