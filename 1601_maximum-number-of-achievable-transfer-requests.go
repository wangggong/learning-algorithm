/*
 * @lc app=leetcode.cn id=maximum-number-of-achievable-transfer-requests lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1601] 最多可达成的换楼请求数目
 *
 * https://leetcode-cn.com/problems/maximum-number-of-achievable-transfer-requests/description/
 *
 * algorithms
 * Hard (50.28%)
 * Total Accepted:    8.1K
 * Total Submissions: 13.3K
 * Testcase Example:  '5\n[[0,1],[1,0],[0,1],[1,2],[2,0],[3,4]]'
 *
 * 我们有 n 栋楼，编号从 0 到 n - 1 。每栋楼有若干员工。由于现在是换楼的季节，部分员工想要换一栋楼居住。
 *
 * 给你一个数组 requests ，其中 requests[i] = [fromi, toi] ，表示一个员工请求从编号为 fromi 的楼搬到编号为
 * toi 的楼。
 *
 * 一开始 所有楼都是满的，所以从请求列表中选出的若干个请求是可行的需要满足 每栋楼员工净变化为 0 。意思是每栋楼 离开 的员工数目 等于 该楼 搬入
 * 的员工数数目。比方说 n = 3 且两个员工要离开楼 0 ，一个员工要离开楼 1 ，一个员工要离开楼 2 ，如果该请求列表可行，应该要有两个员工搬入楼
 * 0 ，一个员工搬入楼 1 ，一个员工搬入楼 2 。
 *
 * 请你从原请求列表中选出若干个请求，使得它们是一个可行的请求列表，并返回所有可行列表中最大请求数目。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：n = 5, requests = [[0,1],[1,0],[0,1],[1,2],[2,0],[3,4]]
 * 输出：5
 * 解释：请求列表如下：
 * 从楼 0 离开的员工为 x 和 y ，且他们都想要搬到楼 1 。
 * 从楼 1 离开的员工为 a 和 b ，且他们分别想要搬到楼 2 和 0 。
 * 从楼 2 离开的员工为 z ，且他想要搬到楼 0 。
 * 从楼 3 离开的员工为 c ，且他想要搬到楼 4 。
 * 没有员工从楼 4 离开。
 * 我们可以让 x 和 b 交换他们的楼，以满足他们的请求。
 * 我们可以让 y，a 和 z 三人在三栋楼间交换位置，满足他们的要求。
 * 所以最多可以满足 5 个请求。
 *
 * 示例 2：
 *
 *
 *
 * 输入：n = 3, requests = [[0,0],[1,2],[2,1]]
 * 输出：3
 * 解释：请求列表如下：
 * 从楼 0 离开的员工为 x ，且他想要回到原来的楼 0 。
 * 从楼 1 离开的员工为 y ，且他想要搬到楼 2 。
 * 从楼 2 离开的员工为 z ，且他想要搬到楼 1 。
 * 我们可以满足所有的请求。
 *
 * 示例 3：
 *
 * 输入：n = 4, requests = [[0,3],[3,1],[1,2],[2,0]]
 * 输出：4
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 20
 * 1 <= requests.length <= 16
 * requests[i].length == 2
 * 0 <= fromi, toi < n
 *
 *
 */

const MAXN = 20 + 10

var ans, cur, cnt0 int
var count [MAXN]int

func maximumRequests(n int, requests [][]int) int {
	ans, cur, cnt0 = 0, 0, 0
	for i := range count {
		count[i] = 0
	}
	backtrack(requests, 0)
	return ans
}

// 暴力回溯, 坑死了.
//
// 之前还一直想找环, 太麻烦了感觉.
func backtrack(arr [][]int, k int) {
	if k == len(arr) {
		if cnt0 == 0 {
			ans = max(ans, cur)
		}
		return
	}
	backtrack(arr, k+1)
	c := cnt0
	cur++
	// NOTE: 这里只能这么写, x, y 的更新后值需要现场维护. 傻逼啊.
	x, y := arr[k][0], arr[k][1]
	if count[x] == 0 {
		cnt0--
	}
	count[x]--
	if count[x] == 0 {
		cnt0++
	}
	if count[y] == 0 {
		cnt0--
	}
	count[y]++
	if count[y] == 0 {
		cnt0++
	}
	backtrack(arr, k+1)
	count[x]++
	count[y]--
	cur--
	cnt0 = c
}

func all0(arr [MAXN]int) bool {
	for _, a := range arr {
		if a != 0 {
			return false
		}
	}
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

