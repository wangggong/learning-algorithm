/*
 * @lc app=leetcode.cn id=three-equal-parts lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [927] 三等分
 *
 * https://leetcode.cn/problems/three-equal-parts/description/
 *
 * algorithms
 * Hard (35.40%)
 * Total Accepted:    9.6K
 * Total Submissions: 23.9K
 * Testcase Example:  '[1,0,1,0,1]'
 *
 * 给定一个由 0 和 1 组成的数组 arr ，将数组分成  3 个非空的部分 ，使得所有这些部分表示相同的二进制值。
 *
 * 如果可以做到，请返回任何 [i, j]，其中 i+1 < j，这样一来：
 *
 *
 * arr[0], arr[1], ..., arr[i] 为第一部分；
 * arr[i + 1], arr[i + 2], ..., arr[j - 1] 为第二部分；
 * arr[j], arr[j + 1], ..., arr[arr.length - 1] 为第三部分。
 * 这三个部分所表示的二进制值相等。
 *
 *
 * 如果无法做到，就返回 [-1, -1]。
 *
 * 注意，在考虑每个部分所表示的二进制时，应当将其看作一个整体。例如，[1,1,0] 表示十进制中的 6，而不会是 3。此外，前导零也是被允许的，所以
 * [0,1,1] 和 [1,1] 表示相同的值。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：arr = [1,0,1,0,1]
 * 输出：[0,3]
 *
 *
 * 示例 2：
 *
 *
 * 输入：arr = [1,1,0,1,1]
 * 输出：[-1,-1]
 *
 * 示例 3:
 *
 *
 * 输入：arr = [1,1,0,0,1]
 * 输出：[0,2]
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 * 3 <= arr.length <= 3 * 10^4
 * arr[i] 是 0 或 1
 *
 *
 */

// 参考: https://leetcode.cn/problems/three-equal-parts/solution/bu-fang-huan-ge-si-lu-kan-kan-lei-si-yu-fz1ih/
// 没想到真的能用字符串哈希完成.
const P uint64 = 13331

var hash, mul []uint64

func get(p, q int) uint64 {
	return hash[q] - hash[p]*mul[q-p]
}

func threeEqualParts(arr []int) []int {
	if all0(arr) {
		return []int{0, 2}
	}
	n := len(arr)
	hash, mul = make([]uint64, n+1), make([]uint64, n+1)
	mul[0] = 1
	for i, a := range arr {
		hash[i+1], mul[i+1] = hash[i]*P+uint64(a), mul[i]*P
	}
	memo := make(map[uint64]int)
	for i := 1; i <= n; i++ {
		h := get(0, i)
		if _, ok := memo[h]; !ok {
			memo[h] = i
		}
	}
	for j := n - 1; j >= 0; j-- {
		h := get(j, n)
		if i, ok := memo[h]; ok && i < j {
			if get(i, j) == h {
				return []int{i - 1, j}
			}
		}
	}
	return []int{-1, -1}
}

func all0(arr []int) bool {
	for _, a := range arr {
		if a != 0 {
			return false
		}
	}
	return true
}
