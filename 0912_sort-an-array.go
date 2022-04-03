/*
 * @lc app=leetcode.cn id=sort-an-array lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [912] 排序数组
 *
 * https://leetcode-cn.com/problems/sort-an-array/description/
 *
 * algorithms
 * Medium (55.58%)
 * Total Accepted:    318.8K
 * Total Submissions: 573.5K
 * Testcase Example:  '[5,2,3,1]'
 *
 * 给你一个整数数组 nums，请你将该数组升序排列。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [5,2,3,1]
 * 输出：[1,2,3,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [5,1,1,2,0,0]
 * 输出：[0,0,1,1,2,5]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 5 * 10^4
 * -5 * 10^4 <= nums[i] <= 5 * 10^4
 *
 *
 */

/*
 * func sortArray(nums []int) []int {
 * 	// 无脑玩赖
 * 	sort.Ints(nums)
 * 	return nums
 * }
 */

// quicksort
func sortArray(nums []int) []int {
	// quicksort(nums, 0, len(nums))
	// mergesort(nums, 0, len(nums))
	heapsort(nums, 0, len(nums))
	return nums
}

func quicksort(arr []int, l, r int) {
	if r-l <= 1 {
		return
	}
	if r-l == 2 {
		if arr[l] > arr[l+1] {
			arr[l], arr[l+1] = arr[l+1], arr[l]
		}
	}
	p := partition(arr, l, r)
	quicksort(arr, l, p)
	quicksort(arr, p+1, r)
}

func partition(arr []int, l, r int) int {
	k := l
	v := arr[l]
	k++
	start := l
	for k < r {
		if arr[k] < v {
			start++
			arr[k], arr[start] = arr[start], arr[k]
		}
		k++
	}
	arr[l], arr[start] = arr[start], arr[l]
	return start
}

func mergesort(arr []int, l, r int) {
	if r-l <= 1 {
		return
	}
	if r-l == 2 {
		if arr[l] > arr[l+1] {
			arr[l], arr[l+1] = arr[l+1], arr[l]
		}
		return
	}
	mid := (r + l) >> 1
	mergesort(arr, l, mid)
	mergesort(arr, mid, r)
	merge(arr, l, mid, r)
}

func merge(arr []int, l, m, r int) {
	left := make([]int, m-l)
	right := make([]int, r-m)
	copy(left, arr[l:m])
	copy(right, arr[m:r])
	p, q, k := 0, 0, l
	// 傻逼! left 和 right 得走自己的索引值!
	for p < m-l && q < r-m {
		if left[p] < right[q] {
			arr[k] = left[p]
			p++
		} else {
			arr[k] = right[q]
			q++
		}
		k++
	}
	for p < m-l {
		arr[k] = left[p]
		p++
		k++
	}
	for q < r-m {
		arr[k] = right[q]
		q++
		k++
	}
	return
}

// 希望我不会再害怕堆排序吧.
func heapsort(arr []int, l, r int) {
	heap := append([]int{r - l}, arr...)
	n := len(heap)
	// build heap
	for i := n >> 1; i > 0; i-- {
		heapify(heap, i, n)
	}
	// sort by heap
	for i := n - 1; i > 0; i-- {
		heap[i], heap[1] = heap[1], heap[i]
		heapify(heap, 1, i)
	}
	// copy result
	copy(arr, heap[1:])
	return
}

func heapify(arr []int, k, n int) {
	for k<<1 < n {
		p := k << 1
		q := k<<1 | 1
		mx := k
		if arr[p] > arr[mx] {
			mx = p
		}
		if q < n && arr[q] > arr[mx] {
			mx = q
		}
		if mx != k {
			arr[k], arr[mx] = arr[mx], arr[k]
			k = mx
		} else {
			break
		}
	}
	return
}
