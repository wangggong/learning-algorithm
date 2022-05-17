/*
 * Hard
 *
 * 字符串的 波动 定义为子字符串中出现次数 最多 的字符次数与出现次数 最少 的字符次数之差。
 *
 * 给你一个字符串 s ，它只包含小写英文字母。请你返回 s 里所有 子字符串的 最大波动 值。
 *
 * 子字符串 是一个字符串的一段连续字符序列。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "aababbb"
 * 输出：3
 * 解释：
 * 所有可能的波动值和它们对应的子字符串如以下所示：
 * - 波动值为 0 的子字符串："a" ，"aa" ，"ab" ，"abab" ，"aababb" ，"ba" ，"b" ，"bb" 和 "bbb" 。
 * - 波动值为 1 的子字符串："aab" ，"aba" ，"abb" ，"aabab" ，"ababb" ，"aababbb" 和 "bab" 。
 * - 波动值为 2 的子字符串："aaba" ，"ababbb" ，"abbb" 和 "babb" 。
 * - 波动值为 3 的子字符串 "babbb" 。
 * 所以，最大可能波动值为 3 。
 * 示例 2：
 *
 * 输入：s = "abcde"
 * 输出：0
 * 解释：
 * s 中没有字母出现超过 1 次，所以 s 中每个子字符串的波动值都是 0 。
 *
 *
 * 提示：
 *
 * 1 <= s.length <= 104
 * s  只包含小写英文字母。
 * 通过次数676
 * 提交次数3,193
 */

// 参考: https://leetcode.cn/problems/substring-with-largest-variance/solution/by-endlesscheng-5775/
//
// PS: Amazon OA 题. 灵神的思路令人豁然开朗. 半暴力也未尝不可.
func largestVariance(s string) int {
	ans := 0
	// 枚举字符 `a` & `b`, 查看 `a` 比 `b` 多多少.
	for a := 'a'; a <= 'z'; a++ {
		for b := 'a'; b <= 'z'; b++ {
			if a == b {
				continue
			}
			// `diff`: `a` 比 `b` 多的字符数. 需要保证 `diff >= 0`;
			// `diffWithB`: (有 `b` 字符时) `a` 比 `b` 多的字符数. 需要保证如果没有的话 `diffWithB <= 0`. 所以初始化时为系统最小.
			diff, diffWithB := 0, -len(s)
			for _, c := range s {
				// 只考虑 `a` / `b` 的情况:
				switch c {
				case a:
					// 如果是 `a`: `diff++`, `diffWithB++`.
					// 但如果还没有 `b` 的话 `diffWithB <= 0`, 大丈夫.
					diff, diffWithB = diff+1, diffWithB+1
				case b:
					// 如果是 `b`: `diff--`, `diffWithB = diff`, 直接取当前值.
					// 如果 `diff < 0`, 说明当前前缀 `a` 比 `b` 少, 抛弃. `diff` 归零. 但 `diffWithB` 不归零.
					diff, diffWithB = max(diff-1, 0), diff-1
				default:
					// do nothing
				}
				// relax `ans`.
				ans = max(ans, diffWithB)
			}
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
 * // 现场另解.
 * const alphaCnt int = 26
 *
 * func largestVariance(s string) int {
 *     n := len(s)
 *     ans := 0
 *     var sum, right [alphaCnt]int
 *     var pos, minVal [alphaCnt][alphaCnt]int
 *     for i := 0; i < n; i++ {
 *         ind := int(s[i] - 'a')
 *         right[ind] = i
 *         sum[ind]++
 *         for j := 0; j < alphaCnt; j++ {
 *             if ind != j && sum[j] > 0 {
 *                 p, q := 0, 0
 *                 if right[j] == pos[ind][j] {
 *                     p = 1
 *                 }
 *                 if right[j] == pos[j][ind] {
 *                     q = 1
 *                 }
 *                 cur := max(
 *                     sum[ind] - sum[j] - minVal[ind][j] - p,
 *                     sum[j] - sum[ind] - minVal[j][ind] - q,
 *                 )
 *                 ans = max(ans, cur)
 *             }
 *         }
 *         for j := 0; j < alphaCnt; j++ {
 *             if sum[ind] - sum[j] < minVal[ind][j] {
 *                 minVal[ind][j] = sum[ind] - sum[j]
 *                 pos[ind][j] = i
 *             }
 *             if sum[j] - sum[ind] < minVal[j][ind] {
 *                 minVal[j][ind] = sum[j] - sum[ind]
 *                 pos[j][ind] = i
 *             }
 *         }
 *     }
 *     return ans
 * }
 *
 * func max(x, y int) int {
 *     if x > y {
 *         return x
 *     }
 *     return y
 * }
 */