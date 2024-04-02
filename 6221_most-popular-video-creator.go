/*
 * @lc app=leetcode.cn id=most-popular-video-creator lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6221] 最流行的视频创作者
 *
 * https://leetcode.cn/problems/most-popular-video-creator/description/
 *
 * algorithms
 * Medium (28.15%)
 * Total Accepted:    3.5K
 * Total Submissions: 12.4K
 * Testcase Example:  '["alice","bob","alice","chris"]\n["one","two","three","four"]\n[5,10,5,4]'
 *
 * 给你两个字符串数组 creators 和 ids ，和一个整数数组 views ，所有数组的长度都是 n 。平台上第 i 个视频者是
 * creator[i] ，视频分配的 id 是 ids[i] ，且播放量为 views[i] 。
 *
 * 视频创作者的 流行度 是该创作者的 所有 视频的播放量的 总和 。请找出流行度 最高 创作者以及该创作者播放量 最大 的视频的 id 。
 *
 *
 * 如果存在多个创作者流行度都最高，则需要找出所有符合条件的创作者。
 * 如果某个创作者存在多个播放量最高的视频，则只需要找出字典序最小的 id 。
 *
 *
 * 返回一个二维字符串数组 answer ，其中 answer[i] = [creatori, idi] 表示 creatori 的流行度 最高
 * 且其最流行的视频 id 是 idi ，可以按任何顺序返回该结果。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：creators = ["alice","bob","alice","chris"], ids =
 * ["one","two","three","four"], views = [5,10,5,4]
 * 输出：[["alice","one"],["bob","two"]]
 * 解释：
 * alice 的流行度是 5 + 5 = 10 。
 * bob 的流行度是 10 。
 * chris 的流行度是 4 。
 * alice 和 bob 是流行度最高的创作者。
 * bob 播放量最高的视频 id 为 "two" 。
 * alice 播放量最高的视频 id 是 "one" 和 "three" 。由于 "one" 的字典序比 "three" 更小，所以结果中返回的 id 是
 * "one" 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：creators = ["alice","alice","alice"], ids = ["a","b","c"], views =
 * [1,2,2]
 * 输出：[["alice","b"]]
 * 解释：
 * id 为 "b" 和 "c" 的视频都满足播放量最高的条件。
 * 由于 "b" 的字典序比 "c" 更小，所以结果中返回的 id 是 "b" 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == creators.length == ids.length == views.length
 * 1 <= n <= 10^5
 * 1 <= creators[i].length, ids[i].length <= 5
 * creators[i] 和 ids[i] 仅由小写英文字母组成
 * 0 <= views[i] <= 10^5
 *
 *
 */
import "sort"

type info struct {
	id   string
	view int
}

func mostPopularCreator(creators []string, ids []string, views []int) (ans [][]string) {
	n := len(creators)
	cnt := make(map[string]int)
	idx := make(map[string][]info)
	mxView := 0
	for i := 0; i < n; i++ {
		c, id, v := creators[i], ids[i], views[i]
		cnt[c] += v
		idx[c] = append(idx[c], info{id, v})
		mxView = max(mxView, cnt[c])
	}
	for k, v := range cnt {
		if v == mxView {
			cur := idx[k]
			sort.Slice(cur, func(p, q int) bool {
				if cur[p].view != cur[q].view {
					return cur[p].view > cur[q].view
				}
				return cur[p].id < cur[q].id
			})
			ans = append(ans, []string{k, cur[0].id})
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
