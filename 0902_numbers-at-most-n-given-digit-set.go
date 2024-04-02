/*
 * @lc app=leetcode.cn id=numbers-at-most-n-given-digit-set lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [902] 最大为 N 的数字组合
 *
 * https://leetcode-cn.com/problems/numbers-at-most-n-given-digit-set/description/
 *
 * algorithms
 * Hard (35.09%)
 * Total Accepted:    4.2K
 * Total Submissions: 11.9K
 * Testcase Example:  '["1","3","5","7"]\n100'
 *
 * 给定一个按 非递减顺序 排列的数字数组 digits 。你可以用任意次数 digits[i] 来写的数字。例如，如果 digits =
 * ['1','3','5']，我们可以写数字，如 '13', '551', 和 '1351315'。
 *
 * 返回 可以生成的小于或等于给定整数 n 的正整数的个数 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：digits = ["1","3","5","7"], n = 100
 * 输出：20
 * 解释：
 * 可写出的 20 个数字是：
 * 1, 3, 5, 7, 11, 13, 15, 17, 31, 33, 35, 37, 51, 53, 55, 57, 71, 73, 75,
 * 77.
 *
 *
 * 示例 2：
 *
 *
 * 输入：digits = ["1","4","9"], n = 1000000000
 * 输出：29523
 * 解释：
 * 我们可以写 3 个一位数字，9 个两位数字，27 个三位数字，
 * 81 个四位数字，243 个五位数字，729 个六位数字，
 * 2187 个七位数字，6561 个八位数字和 19683 个九位数字。
 * 总共，可以使用D中的数字写出 29523 个整数。
 *
 * 示例 3:
 *
 *
 * 输入：digits = ["7"], n = 8
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 * 1 <= digits.length <= 9
 * digits[i].length == 1
 * digits[i] 是从 '1' 到 '9' 的数
 * digits 中的所有值都 不同
 * digits 按 非递减顺序 排列
 * 1 <= n <= 10^9
 *
 *
 */

/*
 * // 参考: https://mp.weixin.qq.com/s?__biz=MzU4NDE3MTEyMA==&mid=2247490779&idx=1&sn=9a07bef5a856ca34f5c18a4541a50e9c
 * //
 * // "数位 DP + 二分", 我既没看出 DP 也没看出二分来, 还有救吗...
 * func atMostNGivenDigitSet(digits []string, n int) int {
 * 	var nDigits []int
 * 	for _, d := range digits {
 * 		nDigits = append(nDigits, int(d[0]-'0'))
 * 	}
 * 	return _atMostNGivenDigitSet(nDigits, n)
 * }
 *
 * // 三种情况:
 * // - case 1: 位数与 `n` 相同, 前缀小于 `n`
 * // - case 2: 位数与 `n` 相同, 前缀等于 `n`
 * // - case 3: 位数比 `n` 小
 * func _atMostNGivenDigitSet(arr []int, n int) int {
 * 	if n == 0 {
 * 		return 0
 * 	}
 * 	var digits []int
 * 	for k := n; k > 0; k /= 10 {
 * 		digits = append(digits, k%10)
 * 	}
 * 	na, nd := len(arr), len(digits)
 * 	ans := 0
 * 	// case 3: 位数比 `n` 小, 乘法原理.
 * 	for i := 1; i < nd; i++ {
 * 		ans += pow(na, i)
 * 	}
 * 	// 位数从高到低遍历.
 * 	for i := nd - 1; i >= 0; i-- {
 * 		j := na - 1
 * 		for j >= 0 {
 * 			if arr[j] <= digits[i] {
 * 				break
 * 			}
 * 			j--
 * 		}
 * 		if j < 0 {
 * 			break
 * 		}
 * 		// assert arr[j] <= digits[i]
 * 		if arr[j] == digits[i] {
 * 			// 如果 `arr[j] == digits[i]`, 考虑两种情况:
 * 			// case 1: 到当前位的前缀小于 `digits` 对应前缀, 乘法原理 (当前位置只能选 `arr[0..j-1]` 共计 `j` 个);
 * 			// case 2: 继续往下找, 每一个结果加上当前位的 `digits[i]` 都一一对应一个结果;
 * 			ans += j * pow(na, i)
 * 			if i == 0 {
 * 				ans++
 * 			}
 * 		} else {
 * 			// 如果 `arr[j] < digits[i]`, 仅有 case 1 一种情况:
 * 			// case 1: 到当前位的前缀小于 `digits` 对应前缀, 乘法原理 (当前位置能选 `arr[0..j]` 共计 `j + 1` 个);
 * 			ans += (j + 1) * pow(na, i)
 * 			break
 * 		}
 * 	}
 * 	return ans
 * }
 *
 * func pow(n, k int) int {
 * 	ans := 1
 * 	cur := n
 * 	for ; k > 0; k >>= 1 {
 * 		if k&1 != 0 {
 * 			ans *= cur
 * 		}
 * 		cur *= cur
 * 	}
 * 	return ans
 * }
 */

// 参考: https://leetcode.cn/problems/numbers-at-most-n-given-digit-set/solution/shu-wei-dp-tong-yong-mo-ban-xiang-xi-zhu-e5dg/
// 更新了灵神的数位 DP 模版.
import "strconv"

type Info struct {
	D       int
	Mask    int
	IsLimit bool
	IsNum   bool
}

func atMostNGivenDigitSet(digits []string, n int) int {
	s := strconv.Itoa(n)
	mask := 0
	for _, d := range digits {
		mask |= 1 << int(d[0]-'0')
	}
	memo := make(map[Info]int)
	return dfs(s, 0, mask, true, false, memo)
}

func dfs(s string, d, mask int, isLimit, isNum bool, memo map[Info]int) int {
	info := Info{d, mask, isLimit, isNum}
	if ans, ok := memo[info]; ok {
		return ans
	}
	ans := 0
	defer func() { memo[info] = ans }()
	if d == len(s) {
		if isNum {
			ans = 1
		}
		return ans
	}
	if !isNum {
		ans += dfs(s, d+1, mask, false, false, memo)
	}
	l, r := 0, 9
	if !isNum {
		l = 1
	}
	if isLimit {
		r = int(s[d] - '0')
	}
	for i := l; i <= r; i++ {
		if (mask>>i)&1 == 0 {
			continue
		}
		ans += dfs(s, d+1, mask, isLimit && i == r, true, memo)
	}
	return ans
}
