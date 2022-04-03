/*
 * Medium
 *
 * 给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
 *
 *
 *
 * 示例 1:
 *
 * 输入: 12258
 * 输出: 5
 * 解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
 *
 *
 * 提示：
 *
 * 0 <= num < 231
 * 通过次数149,486
 * 提交次数282,339
 */

func translateNum(num int) int {
	arr := itoa(num)
	prev2, prev1, curr := 0, 1, 0
	for i, a := range arr {
		curr = prev1
		if i == 0 {
			prev2, prev1 = prev1, curr
			continue
		}
		// assert i > 0;
		if (arr[i-1] == '2' && '0' <= a && a < '6') || arr[i-1] == '1' {
			curr += prev2
		}
		prev2, prev1 = prev1, curr
	}
	return curr
}

func itoa(n int) []byte {
	// assert n >= 0;
	var ans []byte
	if n == 0 {
		return []byte("0")
	}
	for n > 0 {
		ans = append(ans, byte(n%10)+'0')
		n /= 10
	}
	reverse(ans)
	return ans
}

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
