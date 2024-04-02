/*
 * @lc app=leetcode.cn id=fraction-to-recurring-decimal lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [166] 分数到小数
 *
 * https://leetcode-cn.com/problems/fraction-to-recurring-decimal/description/
 *
 * algorithms
 * Medium (33.26%)
 * Total Accepted:    49.4K
 * Total Submissions: 148.6K
 * Testcase Example:  '1\n2'
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
 * 输入：numerator = 4, denominator = 333
 * 输出："0.(012)"
 *
 *
 *
 *
 * 提示：
 *
 *
 * -2^31 <= numerator, denominator <= 2^31 - 1
 * denominator != 0
 *
 *
 */

/*
 * func fractionToDecimal(num int, den int) string {
 * 	if num == 0 {
 * 		return "0"
 * 	}
 * 	// assert den != 0
 * 	neg := (num > 0) != (den > 0)
 * 	num, den = abs(num), abs(den)
 * 	intPart := num / den
 * 	intStr := strconv.Itoa(intPart)
 * 	if neg {
 * 		intStr = "-" + intStr
 * 	}
 * 	num %= den
 * 	if num == 0 {
 * 		return intStr
 * 	}
 * 	floatStr := ""
 * 	memo := make(map[int]string)
 * 	for num > 0 {
 * 		if _, ok := memo[num]; ok {
 * 			loop := findLoop(num, den, memo)
 * 			floatStr = floatStr[:len(floatStr)-len(loop)] + "(" + loop + ")"
 * 			break
 * 		}
 * 		cur := ""
 * 		newNum := num * 10
 * 		for newNum < den {
 * 			newNum *= 10
 * 			cur = cur + "0"
 * 		}
 * 		cur = cur + strconv.Itoa(newNum/den)
 * 		floatStr += cur
 * 		memo[num] = cur
 * 		num = newNum % den
 * 		// fmt.Println(memo, cur, num, newNum)
 * 	}
 * 	return intStr + "." + floatStr
 * }
 *
 * func findLoop(num, den int, memo map[int]string) string {
 * 	next := func(num, den int) int {
 * 		for ; num < den; num *= 10 {
 * 		}
 * 		ans := num % den
 * 		return ans
 * 	}
 * 	ans := memo[num]
 * 	for cur := next(num, den); cur != num; cur = next(cur, den) {
 * 		if cur != num {
 * 			ans = ans + memo[cur]
 * 		}
 * 	}
 * 	return ans
 * }
 *
 * func abs(x int) int {
 * 	if x >= 0 {
 * 		return x
 * 	}
 * 	return -x
 * }
 */

func fractionToDecimal(num int, den int) string {
	if num%den == 0 {
		return strconv.Itoa(num / den)
	}

	s := ""
	if (num < 0) != (den < 0) {
		s = s + "-"
		num, den = abs(num), abs(den)
	}

	s = s + strconv.Itoa(num/den)
	fp := "."
	r := num % den
	memo := make(map[int]int)
	for r != 0 && memo[r] == 0 {
		memo[r] = len(fp)
		r *= 10
		fp = fp + strconv.Itoa(r/den)
		r %= den
	}

	if r != 0 {
		index := memo[r]
		fp = fp[:index] + "(" + fp[index:] + ")"
	}

	return s + fp
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}