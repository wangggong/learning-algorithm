/*
 * @lc app=leetcode.cn id=24-game lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [679] 24 点游戏
 *
 * https://leetcode.cn/problems/24-game/description/
 *
 * algorithms
 * Hard (53.88%)
 * Total Accepted:    32.4K
 * Total Submissions: 60.1K
 * Testcase Example:  '[4,1,8,7]'
 *
 * 给定一个长度为4的整数数组 cards 。你有 4 张卡片，每张卡片上都包含一个范围在 [1,9] 的数字。您应该使用运算符 ['+', '-',
 * '*', '/'] 和括号 '(' 和 ')' 将这些卡片上的数字排列成数学表达式，以获得值24。
 *
 * 你须遵守以下规则:
 *
 *
 * 除法运算符 '/' 表示实数除法，而不是整数除法。
 *
 *
 * 例如， 4 /(1 - 2 / 3)= 4 /(1 / 3)= 12 。
 *
 *
 * 每个运算都在两个数字之间。特别是，不能使用 “-” 作为一元运算符。
 *
 * 例如，如果 cards =[1,1,1,1] ，则表达式 “-1 -1 -1 -1” 是 不允许 的。
 *
 *
 * 你不能把数字串在一起
 *
 * 例如，如果 cards =[1,2,1,2] ，则表达式 “12 + 12” 无效。
 *
 *
 *
 *
 * 如果可以得到这样的表达式，其计算结果为 24 ，则返回 true ，否则返回 false 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: cards = [4, 1, 8, 7]
 * 输出: true
 * 解释: (8-4) * (7-1) = 24
 *
 *
 * 示例 2:
 *
 *
 * 输入: cards = [1, 2, 1, 2]
 * 输出: false
 *
 *
 *
 *
 * 提示:
 *
 *
 * cards.length == 4
 * 1 <= cards[i] <= 9
 *
 *
 */

const (
	add = "add"
	sub = "sub"
	mul = "mul"
	div = "div"
)

const (
	eps    = 1e-6
	target = 24.0
)

var ops = []string{
	add,
	sub,
	mul,
	div,
}

// 被神奇的回溯看傻了...
//
// 关键点有以下几个:
// 1. 转成 `float64`, 为了保证精度只要误差在 `1e-6` 内都算正确;
// 2. 回溯, 每次任取两个数字运算, 算完往里面放;
func judgePoint24(cards []int) bool {
	var arr []float64
	for _, c := range cards {
		arr = append(arr, float64(c))
	}
	return backtrack(arr)
}

func backtrack(arr []float64) bool {
	n := len(arr)
	if n == 0 {
		return false
	}
	if n == 1 {
		return abs(arr[0]-target) <= eps
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			newArr := make([]float64, 0, n-1)
			for k := 0; k < n; k++ {
				if k == i || k == j {
					continue
				}
				newArr = append(newArr, arr[k])
			}
			for _, op := range ops {
				c, ok := calc(arr[i], arr[j], op)
				if !ok {
					continue
				}
				newArr = append(newArr, c)
				if backtrack(newArr) {
					return true
				}
				newArr = newArr[:len(newArr)-1]
			}
		}
	}
	return false
}

func calc(p, q float64, op string) (float64, bool) {
	switch op {
	case add:
		return p + q, true
	case sub:
		return p - q, true
	case mul:
		return p * q, true
	case div:
		if abs(q) <= eps {
			return 0, false
		}
		return p / q, true
	default:
		return 0, false
	}
}

func abs(x float64) float64 {
	if x >= 0 {
		return x
	}
	return -x
}
