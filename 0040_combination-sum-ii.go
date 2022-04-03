/*
 * @lc app=leetcode.cn id=combination-sum-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [40] 组合总和 II
 *
 * https://leetcode-cn.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (61.34%)
 * Total Accepted:    238.7K
 * Total Submissions: 389.1K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的每个数字在每个组合中只能使用 一次 。
 *
 * 注意：解集不能包含重复的组合。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 输出:
 * [
 * [1,1,6],
 * [1,2,5],
 * [1,7],
 * [2,6]
 * ]
 *
 * 示例 2:
 *
 *
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 输出:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= candidates.length <= 100
 * 1 <= candidates[i] <= 50
 * 1 <= target <= 30
 *
 *
 */
import "sort"

var ans [][]int
var cur []int

func dfs(arr []int, n, k int) {
	if n < k {
		return
	}
	if n == k {
		v := make([]int, len(cur))
		copy(v, cur)
		ans = append(ans, v)
		return
	}
	for i, v := range arr {
		if i > 0 && arr[i-1] == v {
			continue
		}
		cur = append(cur, v)
		dfs(arr[i+1:], n, k+v)
		cur = cur[:len(cur)-1]
	}
}

/*
 * func combinationSum2(candidates []int, target int) [][]int {
 * 	sort.Ints(candidates)
 * 	ans, cur = nil, nil
 * 	dfs(candidates, target, 0)
 * 	return ans
 * }
 */

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ans, cur = nil, nil
	backtrack(candidates, target, 0, 0)
	return ans
}

func backtrack(arr []int, target, start, sum int) {
	if sum >= target {
		if sum == target {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			ans = append(ans, tmp)
		}
		return
	}
	for i := start; i < len(arr); i++ {
		if len(cur) > 0 && cur[len(cur)-1] > arr[i] {
			continue
		}
		// NOTE: 每轮回溯仅选取第一个重复元素, 不会选取后面的. 就不重了.
		//
		// e.g.: `1 1 2 5 6 7 10`, 第一次不选第二个 `1`. `1 1 6` 这个解的第二个 `1` 在 `start=2` 时被取到.
		if i > start && arr[i] == arr[i-1] {
			continue
		}
		if sum+arr[i] > target {
			break
		}
		cur = append(cur, arr[i])
		backtrack(arr, target, i+1, sum+arr[i])
		cur = cur[:len(cur)-1]
	}
}
