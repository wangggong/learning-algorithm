/*
 * @lc app=leetcode.cn id=number-of-digit-one lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [233] 数字 1 的个数
 *
 * https://leetcode-cn.com/problems/number-of-digit-one/description/
 *
 * algorithms
 * Hard (47.94%)
 * Total Accepted:    38K
 * Total Submissions: 79.2K
 * Testcase Example:  '13'
 *
 * 给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 13
 * 输出：6
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 0
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= n <= 10^9
 *
 *
 */


// 参考: 剑指 offer, 32
func countDigitOne(n int) int {
	// assert n >= 0 && n <= 1e9;
	if n <= 0 {
		return 0
	}
	return _cntDgtOne(itoa(n))
}

func _cntDgtOne(arr []byte) int {
	size := len(arr)
	if size == 0 {
		return 0
	}
	firstDigit := int(arr[0] - '0')
	if size == 1 {
		switch firstDigit {
		case 0:
			return 0
		default:
			return 1
		}
	}
	var firstDgtCnt, otherDgtCnt int
	switch firstDigit {
	case 0:
		// Do nothing
	case 1:
		// 凑不齐 1e(size-1) 个 '1'
		firstDgtCnt = atoi(arr[1:]) + 1
	default:
		firstDgtCnt = pow(10, size-1)
	}
	otherDgtCnt = firstDigit * (size-1) * pow(10, size-2)
	return firstDgtCnt + otherDgtCnt + _cntDgtOne(arr[1:])
}

func itoa(n int) []byte {
	// assert n >= 0;
	if n == 0 {
		return []byte{'0'}
	}
	var ans []byte
	for n != 0 {
		ans = append(ans, byte(n % 10) + '0')
		n /= 10
	}
	for p, q := 0, len(ans)-1; p < q; p, q = p+1, q-1 {
		ans[p], ans[q] = ans[q], ans[p]
	}
	return ans
}

func atoi(arr []byte) int {
	ans := 0
	for _, a := range arr {
		ans = ans*10 + int(a - '0')
	}
	return ans
}

func pow(base, n int) int {
	// assert base >= 0 && n >= 0;
	ans := 1
	for ; n > 0; n >>= 1 {
		if n & 1 != 0 {
			ans *= base
		}
		base *= base
	}
	return ans
}
