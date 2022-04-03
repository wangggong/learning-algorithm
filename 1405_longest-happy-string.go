/*
 * @lc app=leetcode.cn id=longest-happy-string lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1405] 最长快乐字符串
 *
 * https://leetcode-cn.com/problems/longest-happy-string/description/
 *
 * algorithms
 * Medium (51.17%)
 * Total Accepted:    6.4K
 * Total Submissions: 12.4K
 * Testcase Example:  '1\n1\n7'
 *
 * 如果字符串中不含有任何 'aaa'，'bbb' 或 'ccc' 这样的字符串作为子串，那么该字符串就是一个「快乐字符串」。
 *
 * 给你三个整数 a，b ，c，请你返回 任意一个 满足下列全部条件的字符串 s：
 *
 *
 * s 是一个尽可能长的快乐字符串。
 * s 中 最多 有a 个字母 'a'、b 个字母 'b'、c 个字母 'c' 。
 * s 中只含有 'a'、'b' 、'c' 三种字母。
 *
 *
 * 如果不存在这样的字符串 s ，请返回一个空字符串 ""。
 *
 *
 *
 * 示例 1：
 *
 * 输入：a = 1, b = 1, c = 7
 * 输出："ccaccbcc"
 * 解释："ccbccacc" 也是一种正确答案。
 *
 *
 * 示例 2：
 *
 * 输入：a = 2, b = 2, c = 1
 * 输出："aabbc"
 *
 *
 * 示例 3：
 *
 * 输入：a = 7, b = 1, c = 0
 * 输出："aabaa"
 * 解释：这是该测试用例的唯一正确答案。
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= a, b, c <= 100
 * a + b + c > 0
 *
 *
 */
import "container/heap"

type cell struct {
	x     byte
	count int
}

type hp struct {
	c []cell
}

func (h hp) Len() int            { return len(h.c) }
func (h hp) Less(p, q int) bool  { return h.c[p].count > h.c[q].count }
func (h hp) Swap(p, q int)       { h.c[p], h.c[q] = h.c[q], h.c[p] }
func (h *hp) Push(c interface{}) { h.c = append(h.c, c.(cell)) }
func (h *hp) Pop() interface{}   { l := len(h.c); v := h.c[l-1]; h.c = h.c[:l-1]; return v }

func longestDiverseString(a int, b int, c int) string {
	var ans []byte
	h := &hp{}
	if a > 0 {
		heap.Push(h, cell{'a', a})
	}
	if b > 0 {
		heap.Push(h, cell{'b', b})
	}
	if c > 0 {
		heap.Push(h, cell{'c', c})
	}
	for h.Len() > 0 {
		v := heap.Pop(h).(cell)
		duplicate := false
		if l := len(ans); l >= 2 && (v.x == ans[l-1] && v.x == ans[l-2]) {
			duplicate = true
		}
		if !duplicate {
			ans = append(ans, v.x)
			v.count--
			if v.count > 0 {
				heap.Push(h, v)
			}
			continue
		}
		if h.Len() == 0 {
			break
		}
		w := heap.Pop(h).(cell)
		ans = append(ans, w.x)
		w.count--
		if v.count > 0 {
			heap.Push(h, v)
		}
		if w.count > 0 {
			heap.Push(h, w)
		}
	}
	return string(ans)
}
