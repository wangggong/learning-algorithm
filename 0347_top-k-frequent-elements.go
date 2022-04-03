/*
 * @lc app=leetcode.cn id=top-k-frequent-elements lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [347] 前 K 个高频元素
 *
 * https://leetcode-cn.com/problems/top-k-frequent-elements/description/
 *
 * algorithms
 * Medium (62.52%)
 * Total Accepted:    252.6K
 * Total Submissions: 403.3K
 * Testcase Example:  '[1,1,1,2,2,3]\n2'
 *
 * 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,1,1,2,2,3], k = 2
 * 输出: [1,2]
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [1], k = 1
 * 输出: [1]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 1e5
 * k 的取值范围是 [1, 数组中不相同的元素的个数]
 * 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
 *
 *
 *
 *
 * 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
 *
 */

var count map[int]int

func topKFrequent(nums []int, k int) []int {
	count = make(map[int]int)
	for _, n := range nums {
		count[n]++
	}
	var candidates []int
	for k := range count {
		candidates = append(candidates, k)
	}
	n := len(candidates)
	p, q := 0, n-1
	for p < q {
		ind := partition(candidates, p, q)
		// fmt.Printf("%v %v %v %v\n", candidates, p, q, ind)
		if ind == n-k {
			break
		} else if ind > n-k {
			q = ind - 1
		} else {
			p = ind + 1
		}
	}
	return candidates[n-k:]
}

func partition(arr []int, p, q int) int {
	// assert 0 <= p && p <= q && q < len(arr);
	v, start := arr[p], p
	for i := p + 1; i <= q; i++ {
		if count[arr[i]] < count[v] {
			start++
			arr[start], arr[i] = arr[i], arr[start]
		}
	}
	arr[start], arr[p] = arr[p], arr[start]
	// fmt.Printf("%v %v %v %v", arr, p, q, start)
	return start
}
