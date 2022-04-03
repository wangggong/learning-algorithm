/*
 * @lc app=leetcode.cn id=sliding-window-maximum lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [239] 滑动窗口最大值
 *
 * https://leetcode-cn.com/problems/sliding-window-maximum/description/
 *
 * algorithms
 * Hard (49.76%)
 * Total Accepted:    244K
 * Total Submissions: 490.1K
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k
 * 个数字。滑动窗口每次只向右移动一位。
 *
 * 返回 滑动窗口中的最大值 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
 * 输出：[3,3,5,5,6,7]
 * 解释：
 * 滑动窗口的位置                最大值
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       3
 * ⁠1 [3  -1  -3] 5  3  6  7       3
 * ⁠1  3 [-1  -3  5] 3  6  7       5
 * ⁠1  3  -1 [-3  5  3] 6  7       5
 * ⁠1  3  -1  -3 [5  3  6] 7       6
 * ⁠1  3  -1  -3  5 [3  6  7]      7
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1], k = 1
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 * 1 <= k <= nums.length
 *
 *
 */

const (
	Ind = iota
	Val
)

type Heap struct {
	Arr [][2]int
}

func (h Heap) Len() int            { return len(h.Arr) }
func (h Heap) Less(p, q int) bool  { return h.Arr[p][Val] > h.Arr[q][Val] }
func (h Heap) Swap(p, q int)       { h.Arr[p], h.Arr[q] = h.Arr[q], h.Arr[p] }
func (h *Heap) Push(v interface{}) { h.Arr = append(h.Arr, v.([2]int)) }
func (h *Heap) Pop() interface{}   { l := len(h.Arr); v := h.Arr[l-1]; h.Arr = h.Arr[:l-1]; return v }

/*
 * // 优先队列
 * func maxSlidingWindow(nums []int, k int) []int {
 * 	n := len(nums)
 * 	h := &Heap{}
 * 	var ans []int
 * 	for i := 0; i < k; i++ {
 * 		heap.Push(h, [2]int{i, nums[i]})
 * 	}
 * 	ans = append(ans, h.Arr[0][Val])
 * 	for i := k; i < n; i++ {
 * 		heap.Push(h, [2]int{i, nums[i]})
 * 		for h.Len() > 0 && h.Arr[0][Ind]+k <= i {
 * 			heap.Pop(h)
 * 		}
 * 		ans = append(ans, h.Arr[0][Val])
 * 	}
 * 	return ans
 * }
 */

// 单调队列
// 1. 不同于单调栈;
// 2. 队首 -> 队尾值逐渐减小, 但每次稳定放入当前值;
// 3. 想象一下如果队列 size > k 会怎样? 不怎样, 除了首次计算, 每次计算时会把 stale 数据吐出;
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	var Q, ans []int
	push := func(ind int) {
		for len(Q) > 0 && nums[Q[len(Q)-1]] <= nums[ind] {
			Q = Q[:len(Q)-1]
		}
		Q = append(Q, ind)
	}
	for i := 0; i < k; i++ {
		push(i)
	}
	ans = append(ans, nums[Q[0]])
	for i := k; i < n; i++ {
		push(i)
		for len(Q) > 0 && Q[0]+k <= i {
			Q = Q[1:]
		}
		ans = append(ans, nums[Q[0]])
	}
	return ans
}

