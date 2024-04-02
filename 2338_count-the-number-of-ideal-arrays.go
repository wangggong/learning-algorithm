/*
 * @lc app=leetcode.cn id=count-the-number-of-ideal-arrays lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2338] 统计理想数组的数目
 *
 * https://leetcode.cn/problems/count-the-number-of-ideal-arrays/description/
 *
 * algorithms
 * Hard (28.23%)
 * Total Accepted:    2.2K
 * Total Submissions: 7.7K
 * Testcase Example:  '2\n5'
 *
 * 给你两个整数 n 和 maxValue ，用于描述一个 理想数组 。
 *
 * 对于下标从 0 开始、长度为 n 的整数数组 arr ，如果满足以下条件，则认为该数组是一个 理想数组 ：
 *
 *
 * 每个 arr[i] 都是从 1 到 maxValue 范围内的一个值，其中 0 <= i < n 。
 * 每个 arr[i] 都可以被 arr[i - 1] 整除，其中 0 < i < n 。
 *
 *
 * 返回长度为 n 的 不同 理想数组的数目。由于答案可能很大，返回对 10^9 + 7 取余的结果。
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 2, maxValue = 5
 * 输出：10
 * 解释：存在以下理想数组：
 * - 以 1 开头的数组（5 个）：[1,1]、[1,2]、[1,3]、[1,4]、[1,5]
 * - 以 2 开头的数组（2 个）：[2,2]、[2,4]
 * - 以 3 开头的数组（1 个）：[3,3]
 * - 以 4 开头的数组（1 个）：[4,4]
 * - 以 5 开头的数组（1 个）：[5,5]
 * 共计 5 + 2 + 1 + 1 + 1 = 10 个不同理想数组。
 *
 *
 * 示例 2：
 *
 * 输入：n = 5, maxValue = 3
 * 输出：11
 * 解释：存在以下理想数组：
 * - 以 1 开头的数组（9 个）：
 * ⁠  - 不含其他不同值（1 个）：[1,1,1,1,1]
 * ⁠  - 含一个不同值 2（4 个）：[1,1,1,1,2], [1,1,1,2,2], [1,1,2,2,2], [1,2,2,2,2]
 * ⁠  - 含一个不同值 3（4 个）：[1,1,1,1,3], [1,1,1,3,3], [1,1,3,3,3], [1,3,3,3,3]
 * - 以 2 开头的数组（1 个）：[2,2,2,2,2]
 * - 以 3 开头的数组（1 个）：[3,3,3,3,3]
 * 共计 9 + 1 + 1 = 11 个不同理想数组。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= n <= 10^4
 * 1 <= maxValue <= 10^4
 *
 *
 */

// 参考: https://leetcode.cn/problems/count-the-number-of-ideal-arrays/solution/shu-lun-zu-he-shu-xue-zuo-fa-by-endlessc-iouh/
//
// 给我看懵了真的是...
const (
	mod  int = 1e9 + 7
	maxn int = 1e4
	K    int = 13
)

var fcs, C [][]int

func idealArrays(n int, m int) int {
	ans := 1
	for i := 2; i <= m; i++ {
		mul := 1
		for _, fc := range fcs[i] {
			mul = (mul * C[n+fc-1][fc]) % mod
		}
		ans = (ans + mul) % mod
	}
	return ans
}

func init() {
	fcs = make([][]int, maxn+1)
	for k := 2; k <= maxn; k++ {
		fcs[k] = factorCnt(k)
	}

	C = make([][]int, maxn+K)
	for i := range C {
		C[i] = make([]int, K+1)
	}
	C[0][0], C[1][0], C[1][1] = 1, 1, 1
	for i := 2; i < maxn+K; i++ {
		C[i][0] = 1
		for j := 1; j <= min(i, K); j++ {
			C[i][j] = (C[i-1][j-1] + C[i-1][j]) % mod
		}
	}
	return
}

func factorCnt(k int) []int {
	var ans []int
	for i := 2; i <= k/i; i++ {
		c := 0
		for ; k%i == 0; k /= i {
			c++
		}
		if c > 0 {
			ans = append(ans, c)
		}
	}
	if k > 1 {
		ans = append(ans, 1)
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
