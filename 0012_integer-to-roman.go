/*
 * @lc app=leetcode.cn id=integer-to-roman lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [12] 整数转罗马数字
 *
 * https://leetcode-cn.com/problems/integer-to-roman/description/
 *
 * algorithms
 * Medium (66.39%)
 * Total Accepted:    264.1K
 * Total Submissions: 397.9K
 * Testcase Example:  '3'
 *
 * 罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。
 *
 *
 * 字符          数值
 * I             1
 * V             5
 * X             10
 * L             50
 * C             100
 * D             500
 * M             1000
 *
 * 例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V +
 * II 。
 *
 * 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数
 * 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
 *
 *
 * I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
 * X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
 * C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
 *
 *
 * 给你一个整数，将其转为罗马数字。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: num = 3
 * 输出: "III"
 *
 * 示例 2:
 *
 *
 * 输入: num = 4
 * 输出: "IV"
 *
 * 示例 3:
 *
 *
 * 输入: num = 9
 * 输出: "IX"
 *
 * 示例 4:
 *
 *
 * 输入: num = 58
 * 输出: "LVIII"
 * 解释: L = 50, V = 5, III = 3.
 *
 *
 * 示例 5:
 *
 *
 * 输入: num = 1994
 * 输出: "MCMXCIV"
 * 解释: M = 1000, CM = 900, XC = 90, IV = 4.
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= num <= 3999
 *
 *
 */

var wordMap = map[int]string{
	1:    "I",
	2:    "II",
	3:    "III",
	4:    "IV",
	5:    "V",
	6:    "VI",
	7:    "VII",
	8:    "VIII",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

type RomanValue struct {
	Val int
	Rom string
}

var romanValues = []RomanValue{
	{Val: 1000, Rom: "M"},
	{Val: 900, Rom: "CM"},
	{Val: 500, Rom: "D"},
	{Val: 400, Rom: "CD"},
	{Val: 100, Rom: "C"},
	{Val: 90, Rom: "XC"},
	{Val: 50, Rom: "L"},
	{Val: 40, Rom: "XL"},
	{Val: 10, Rom: "X"},
	{Val: 9, Rom: "IX"},
	{Val: 8, Rom: "VIII"},
	{Val: 7, Rom: "VII"},
	{Val: 6, Rom: "VI"},
	{Val: 5, Rom: "V"},
	{Val: 4, Rom: "IV"},
	{Val: 3, Rom: "III"},
	{Val: 2, Rom: "II"},
	{Val: 1, Rom: "I"},
}

/*
 * func intToRoman(num int) string {
 * 	var ans []byte
 * 	for factor := 1000; factor > 0; factor /= 10 {
 * 		p := num / factor
 * 		if p == 0 {
 * 			continue
 * 		}
 * 		ans = append(ans, getRomanValue(p, factor)...)
 * 		num %= factor
 * 		// fmt.Printf("%v %s %v\n", factor, ans, num)
 * 	}
 * 	return string(ans)
 * }
 */

func intToRoman(num int) string {
	ans := ""
	for _, v := range romanValues {
		for num >= v.Val {
			ans = ans + v.Rom
			num -= v.Val
		}
	}
	return ans
}

func getRomanValue(p, factor int) []byte {
	// assert 1 <= p <= 9;
	// assert factor in [1, 10, 100, 1000];
	var ans []byte
	switch p {
	case 1:
		factorBytes := []byte(wordMap[factor])
		ans = append(ans, factorBytes...)
	case 2, 3, 6, 7, 8:
		// pBytes := []byte(wordMap[p])
		// factorBytes := []byte(wordMap[factor])
		// ans = append(ans, pBytes...)
		// if factor > 1 {
		// 	ans = append(ans, factorBytes...)
		// }
		if p > 5 {
			bytes := []byte(wordMap[5*factor])
			ans = append(ans, bytes...)
			p -= 5
		}
		factorBytes := []byte(wordMap[factor])
		for i := 0; i < p; i++ {
			ans = append(ans, factorBytes...)
		}
	case 4, 5, 9:
		bytes := []byte(wordMap[p*factor])
		ans = append(ans, bytes...)
	default:
		// unreachable
	}
	return ans
}
