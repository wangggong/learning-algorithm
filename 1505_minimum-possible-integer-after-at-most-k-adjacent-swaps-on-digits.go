/*
 * @lc app=leetcode.cn id=minimum-possible-integer-after-at-most-k-adjacent-swaps-on-digits lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1505] 最多 K 次交换相邻数位后得到的最小整数
 *
 * https://leetcode-cn.com/problems/minimum-possible-integer-after-at-most-k-adjacent-swaps-on-digits/description/
 *
 * algorithms
 * Hard (36.88%)
 * Total Accepted:    3.6K
 * Total Submissions: 9.8K
 * Testcase Example:  '"4321"\n4'
 *
 * 给你一个字符串 num 和一个整数 k 。其中，num 表示一个很大的整数，字符串中的每个字符依次对应整数上的各个 数位 。
 *
 * 你可以交换这个整数相邻数位的数字 最多 k 次。
 *
 * 请你返回你能得到的最小整数，并以字符串形式返回。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：num = "4321", k = 4
 * 输出："1342"
 * 解释：4321 通过 4 次交换相邻数位得到最小整数的步骤如上图所示。
 *
 *
 * 示例 2：
 *
 *
 * 输入：num = "100", k = 1
 * 输出："010"
 * 解释：输出可以包含前导 0 ，但输入保证不会有前导 0 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：num = "36789", k = 1000
 * 输出："36789"
 * 解释：不需要做任何交换。
 *
 *
 * 示例 4：
 *
 *
 * 输入：num = "22", k = 22
 * 输出："22"
 *
 *
 * 示例 5：
 *
 *
 * 输入：num = "9438957234785635408", k = 23
 * 输出："0345989723478563548"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= num.length <= 30000
 * num 只包含 数字 且不含有 前导 0 。
 * 1 <= k <= 10^9
 *
 *
 */

/*
 * // 先来个模拟面试时的超时版本:
 * func minInteger(num string, k int) string {
 * 	return string(process(nil, []byte(num), k))
 * }
 *
 * func process(prefix, arr []byte, k int) []byte {
 * 	if k == 0 {
 * 		prefix = append(prefix, arr...)
 * 		return prefix
 * 	}
 * 	n := len(arr)
 * 	inc := true
 * 	for i := 0; i+1 < n; i++ {
 * 		if arr[i] > arr[i+1] {
 * 			inc = false
 * 		}
 * 	}
 * 	if inc {
 * 		prefix = append(prefix, arr...)
 * 		return prefix
 * 	}
 * 	minInd := 0
 * 	for i := 0; i <= k && i < n; i++ {
 * 		if arr[i] < arr[minInd] {
 * 			minInd = i
 * 		}
 * 	}
 * 	prefix = append(prefix, arr[minInd])
 * 	ns := append(arr[:minInd], arr[minInd+1:]...)
 * 	return process(prefix, ns, k-minInd)
 * }
 */

// 官解是 BIT:
import (
	"fmt"
	"strconv"
	"strings"
)

const digitCnt int = 1e1

func lowBit(x int) int { return x & (-x) }

type BIT []int

func NewBIT(n int) *BIT {
	b := BIT(make([]int, n+1))
	return &b
}

func (b BIT) query(x int) int {
	ans := 0
	for ; x > 0; x -= lowBit(x) {
		ans += b[x]
	}
	return ans
}

func (b BIT) between(p, q int) int {
	return b.query(q) - b.query(p-1)
}

func (b BIT) String() string {
	n := len(b)
	arr := make([]string, 0, n)
	for i := 1; i < n; i++ {
		arr = append(arr, strconv.Itoa(b.query(i)))
	}
	return strings.Join(arr, ",")
}

func (b *BIT) add(x, v int) {
	for n := len(*b) - 1; x <= n; x += lowBit(x) {
		(*b)[x] += v
	}
}

func minInteger(num string, k int) string {
	arr := []byte(num)
	n := len(arr)
	var index [digitCnt][]int
	for i, a := range arr {
		ind := int(a - '0')
		index[ind] = append(index[ind], i+1)
	}
	b := NewBIT(n)
	var ans []byte
	for i := 0; i < n; i++ {
		for j := 0; j < digitCnt; j++ {
			if len(index[j]) == 0 {
				continue
			}
			pos := index[j][0]
			// NOTE: 查询的时候只查询 `[pos+1, n]` 区间的总和, 因为我们关心的是该位 **后面** 有多少数字挪到前面去了.
			behind := b.between(pos+1, n)
			dist := pos + behind - (i + 1)
			if dist <= k {
				index[j] = index[j][1:]
				// NOTE: 移动的当位也 +1, 相当于 `[pos, n]` 区间整体 +1 了. 该位本身也挪到前面去了.
				b.add(pos, 1)
				fmt.Println("%v\n", b)
				k -= dist
				ans = append(ans, byte(j)+'0')
				break
			}
		}
	}
	return string(ans)
}

