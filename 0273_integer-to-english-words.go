import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=integer-to-english-words lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [273] 整数转换英文表示
 *
 * https://leetcode-cn.com/problems/integer-to-english-words/description/
 *
 * algorithms
 * Hard (36.58%)
 * Total Accepted:    30.7K
 * Total Submissions: 83.9K
 * Testcase Example:  '123'
 *
 * 将非负整数 num 转换为其对应的英文表示。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：num = 123
 * 输出："One Hundred Twenty Three"
 *
 *
 * 示例 2：
 *
 *
 * 输入：num = 12345
 * 输出："Twelve Thousand Three Hundred Forty Five"
 *
 *
 * 示例 3：
 *
 *
 * 输入：num = 1234567
 * 输出："One Million Two Hundred Thirty Four Thousand Five Hundred Sixty
 * Seven"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= num <= 2^31 - 1
 *
 *
 */

var factors = []int{1e9, 1e6, 1e3}
var quantities = []string{"Billion", "Million", "Thousand"}
var wordMap = map[int]string{
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
	11: "Eleven",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
	16: "Sixteen",
	17: "Seventeen",
	18: "Eighteen",
	19: "Nineteen",
	20: "Twenty",
	30: "Thirty",
	40: "Forty",
	50: "Fifty",
	60: "Sixty",
	70: "Seventy",
	80: "Eighty",
	90: "Ninety",
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	ans := ""
	for i := 0; i < 3; i++ {
		if num/factors[i] == 0 {
			continue
		}
		// p, q := num/factors[i], num%factors[i]
		p := num / factors[i]
		ans = fmt.Sprintf("%v %v %v", ans, numberUpToThousand(p), quantities[i])
		// if q != 0 {
		// 	ans = fmt.Sprintf("%v %v", ans, numberUpToThousand(q))
		// }
		num %= factors[i]
	}
	if num != 0 {
		ans = fmt.Sprintf("%v %v", ans, numberUpToThousand(num))
	}
	ans = strings.Trim(ans, " ")
	return ans
}

func numberUpToHundred(num int) string {
	// assert 0 < num < 100;
	if w, ok := wordMap[num]; ok {
		return w
	}
	p, q := num/10*10, num%10
	return fmt.Sprintf("%v %v", wordMap[p], wordMap[q])
}

func numberUpToThousand(num int) string {
	// assert 0 < num < 1000;
	p, q := num/100, num%100
	underHundred := numberUpToHundred(q)
	if p == 0 {
		return underHundred
	}
	hundred := fmt.Sprintf("%v Hundred", wordMap[p])
	if q == 0 {
		return hundred
	}
	return fmt.Sprintf("%v %v", hundred, underHundred)
}
