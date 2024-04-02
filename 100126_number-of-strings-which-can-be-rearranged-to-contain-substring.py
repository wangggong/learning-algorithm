#
# @lc app=leetcode.cn id=number-of-strings-which-can-be-rearranged-to-contain-substring lang=python3
#
# NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
#
# [100126] 重新排列后包含指定子字符串的字符串数目
#
# https://leetcode.cn/problems/number-of-strings-which-can-be-rearranged-to-contain-substring/description/
#
# algorithms
# Medium (55.26%)
# Total Accepted:    927
# Total Submissions: 1.7K
# Testcase Example:  '4'
#
# 给你一个整数 n 。
# 
# 如果一个字符串 s 只包含小写英文字母，且 将 s 的字符重新排列后，新字符串包含 子字符串 "leet" ，那么我们称字符串 s 是一个 好 字符串。
# 
# 比方说：
# 
# 
# 字符串 "lteer" 是好字符串，因为重新排列后可以得到 "leetr" 。
# "letl" 不是好字符串，因为无法重新排列并得到子字符串 "leet" 。
# 
# 
# 请你返回长度为 n 的好字符串 总 数目。
# 
# 由于答案可能很大，将答案对 10^9 + 7 取余 后返回。
# 
# 子字符串 是一个字符串中一段连续的字符序列。
# 
# 
# 
# 示例 1：
# 
# 
# 输入：n = 4
# 输出：12
# 解释：总共有 12 个字符串重新排列后包含子字符串 "leet" ："eelt" ，"eetl" ，"elet" ，"elte" ，"etel"
# ，"etle" ，"leet" ，"lete" ，"ltee" ，"teel" ，"tele" 和 "tlee" 。
# 
# 
# 示例 2：
# 
# 
# 输入：n = 10
# 输出：83943898
# 解释：长度为 10 的字符串重新排列后包含子字符串 "leet" 的方案数为 526083947580 。所以答案为 526083947580 %
# (10^9 + 7) = 83943898 。
# 
# 
# 
# 
# 提示：
# 
# 
# 1 <= n <= 10^5
# 
# 
# 参考: https://leetcode.cn/problems/number-of-strings-which-can-be-rearranged-to-contain-substring/solutions/2522964/olog-n-rong-chi-yuan-li-fu-ji-yi-hua-sou-okjf/
# 正难则反
class Solution:
    def stringCount(self, n: int) -> int:
        return 0 if n < 4 else (26 ** n
            - (25 ** n) * 3 - (25 ** (n - 1)) * n
            + (24 ** n) * 3 + (24 ** (n - 1)) * n * 2
            - 23 ** n       - (23 ** (n - 1)) * n) % 1000000007
