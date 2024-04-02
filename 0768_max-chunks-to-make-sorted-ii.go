/*
 * @lc app=leetcode.cn id=max-chunks-to-make-sorted-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [768] 最多能完成排序的块 II
 *
 * https://leetcode.cn/problems/max-chunks-to-make-sorted-ii/description/
 *
 * algorithms
 * Hard (52.37%)
 * Total Accepted:    14.5K
 * Total Submissions: 26K
 * Testcase Example:  '[5,4,3,2,1]'
 *
 * 这个问题和“最多能完成排序的块”相似，但给定数组中的元素可以重复，输入数组最大长度为2000，其中的元素最大为10**8。
 *
 *
 * arr是一个可能包含重复元素的整数数组，我们将这个数组分割成几个“块”，并将这些块分别进行排序。之后再连接起来，使得连接的结果和按升序排序后的原数组相同。
 *
 * 我们最多能将数组分成多少块？
 *
 * 示例 1:
 *
 *
 * 输入: arr = [5,4,3,2,1]
 * 输出: 1
 * 解释:
 * 将数组分成2块或者更多块，都无法得到所需的结果。
 * 例如，分成 [5, 4], [3, 2, 1] 的结果是 [4, 5, 1, 2, 3]，这不是有序的数组。
 *
 *
 * 示例 2:
 *
 *
 * 输入: arr = [2,1,3,4,4]
 * 输出: 4
 * 解释:
 * 我们可以把它分成两块，例如 [2, 1], [3, 4, 4]。
 * 然而，分成 [2, 1], [3], [4], [4] 可以得到最多的块数。
 *
 *
 * 注意:
 *
 *
 * arr的长度在[1, 2000]之间。
 * arr[i]的大小在[0, 10**8]之间。
 *
 *
 */

/*
 * import "sort"
 *
 * // 排序 + 哈希
 * func maxChunksToSorted(arr []int) (ans int) {
 * 	n := len(arr)
 * 	sorted := make([]int, n)
 * 	copy(sorted, arr)
 * 	sort.Ints(sorted)
 * 	count := make(map[int]int)
 * 	for i := 0; i < n; i++ {
 * 		count[arr[i]]--
 * 		count[sorted[i]]++
 * 		if count[arr[i]] == 0 {
 * 			delete(count, arr[i])
 * 		}
 * 		if count[sorted[i]] == 0 {
 * 			delete(count, sorted[i])
 * 		}
 * 		if len(count) == 0 {
 * 			ans++
 * 		}
 * 	}
 * 	return
 * }
 */

// 单调栈
func maxChunksToSorted(arr []int) int {
	var S []int
	for _, a := range arr {
		if len(S) == 0 {
			S = append(S, a)
			continue
		}
		if a >= S[len(S)-1] {
			S = append(S, a)
			continue
		}
		top := S[len(S)-1]
		for len(S) > 0 && a < S[len(S)-1] {
			S = S[:len(S)-1]
		}
		S = append(S, top)
	}
	return len(S)
}
