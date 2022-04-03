/*
 * @lc app=leetcode.cn id=expression-add-operators lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [282] 给表达式添加运算符
 *
 * https://leetcode-cn.com/problems/expression-add-operators/description/
 *
 * algorithms
 * Hard (48.29%)
 * Total Accepted:    21.1K
 * Total Submissions: 43.6K
 * Testcase Example:  '"123"\n6'
 *
 * 给定一个仅包含数字 0-9 的字符串 num 和一个目标值整数 target ，在 num 的数字之间添加 二元 运算符（不是一元）+、- 或 *
 * ，返回所有能够得到目标值的表达式。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: num = "123", target = 6
 * 输出: ["1+2+3", "1*2*3"]
 *
 *
 * 示例 2:
 *
 *
 * 输入: num = "232", target = 8
 * 输出: ["2*3+2", "2+3*2"]
 *
 * 示例 3:
 *
 *
 * 输入: num = "105", target = 5
 * 输出: ["1*0+5","10-5"]
 *
 * 示例 4:
 *
 *
 * 输入: num = "00", target = 0
 * 输出: ["0+0", "0-0", "0*0"]
 *
 *
 * 示例 5:
 *
 *
 * 输入: num = "3456237490", target = 9191
 * 输出: []
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= num.length <= 10
 * num 仅含数字
 * -2^31 <= target <= 2^31 - 1
 *
 *
 */

/*
 * // 调了半宿发现超时了, 难受.
 * var ans []string
 * var cur []byte
 * var ops = []byte("+-*")
 * var Snum []int
 * var Sop []byte
 * var priority = map[byte]int{
 * 	'+': 1,
 * 	'-': 1,
 * 	'*': 2,
 * }
 *
 * func addOperators(num string, target int) []string {
 * 	bytes := []byte(num)
 * 	n := len(bytes)
 * 	// assert 0 < n && n <= 10;
 * 	ans, cur = nil, nil
 * 	dfs(bytes, 0, n, target)
 * 	return ans
 * }
 *
 * func dfs(arr []byte, start, end, target int) {
 * 	if start >= end {
 * 		v := eval(cur)
 * 		if v != target {
 * 			return
 * 		}
 * 		tmp := make([]byte, len(cur))
 * 		copy(tmp, cur)
 * 		ans = append(ans, string(tmp))
 * 		return
 * 	}
 * 	cur = append(cur, arr[start:end]...)
 * 	dfs(arr, end, end, target)
 * 	cur = cur[:len(cur)-(end-start)]
 * 	for i := start + 1; i < end; i++ {
 * 		cur = append(cur, arr[start:i]...)
 * 		for _, op := range ops {
 * 			cur = append(cur, op)
 * 			dfs(arr, i, end, target)
 * 			cur = cur[:len(cur)-1]
 * 		}
 * 		cur = cur[:len(cur)-(i-start)]
 * 	}
 * }
 *
 * func eval(arr []byte) int {
 * 	fmt.Printf("%s\n", arr)
 * 	n := len(arr)
 * 	Snum, Sop = nil, nil
 * 	Snum = append(Snum, 0)
 * 	for i := 0; i < n; i++ {
 * 		a := arr[i]
 * 		switch a {
 * 		case '0', '1', '2', '3', '4',
 * 			'5', '6', '7', '8', '9':
 * 			v, j := 0, i
 * 			for j < n && '0' <= arr[j] && arr[j] <= '9' {
 * 				if j > i && a == '0' {
 * 					return -10086
 * 				}
 * 				v *= 10
 * 				v += int(arr[j] - '0')
 * 				j++
 * 			}
 * 			Snum = append(Snum, v)
 * 			i = j - 1
 * 		case '+', '-', '*':
 * 			if i > 0 && (arr[i-1] == '+' || arr[i-1] == '-') {
 * 				Snum = append(Snum, 0)
 * 			}
 * 			for len(Sop) > 0 {
 * 				preOp := Sop[len(Sop)-1]
 * 				if priority[preOp] >= priority[a] {
 * 					_calculate()
 * 				} else {
 * 					break
 * 				}
 * 			}
 * 			Sop = append(Sop, a)
 * 		default:
 * 			// unreachable
 * 		}
 * 	}
 * 	for len(Sop) > 0 {
 * 		_calculate()
 * 	}
 * 	return Snum[len(Snum)-1]
 * }
 *
 * func _calculate() {
 * 	n, m := len(Snum), len(Sop)
 * 	if n < 2 || m < 1 {
 * 		return
 * 	}
 * 	a, b := Snum[n-2], Snum[n-1]
 * 	op := Sop[m-1]
 * 	Snum, Sop = Snum[:n-2], Sop[:m-1]
 * 	switch op {
 * 	case '+':
 * 		a += b
 * 	case '-':
 * 		a -= b
 * 	case '*':
 * 		a *= b
 * 	default:
 * 		// unreachable
 * 	}
 * 	Snum = append(Snum, a)
 * }
 */

var ans []string
var bytes []byte
var target int

func addOperators(num string, t int) []string {
	ans, bytes, target = nil, []byte(num), t
	n := len(bytes)
	cur := make([]byte, 0, 2*n-1)
	backtrack(cur, 0, n, 0, 0)
	return ans
}

func backtrack(arr []byte, k, n, val, mul int) {
	if k == n {
		if val == target {
			ans = append(ans, string(arr))
		}
		return
	}
	signIndex := len(arr)
	if k != 0 {
		arr = append(arr, 0)
	}
	for v, j := 0, k; j < n && (j == k || bytes[k] != '0'); j++ {
		v = v*10 + int(bytes[j]-'0')
		arr = append(arr, bytes[j])
		if k == 0 {
			backtrack(arr, j+1, n, v, v)
			continue
		}
		arr[signIndex] = '+'
		backtrack(arr, j+1, n, val+v, v)
		arr[signIndex] = '-'
		backtrack(arr, j+1, n, val-v, -v)
		arr[signIndex] = '*'
		backtrack(arr, j+1, n, val-mul+v*mul, v*mul)
	}
}
