/*
 * Medium
 *
 * 给你一个二维整数数组 tiles ，其中 tiles[i] = [li, ri] ，表示所有在 li <= j <= ri 之间的每个瓷砖位置 j 都被涂成了白色。
 *
 * 同时给你一个整数 carpetLen ，表示可以放在 任何位置 的一块毯子。
 *
 * 请你返回使用这块毯子，最多 可以盖住多少块瓷砖。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：tiles = [[1,5],[10,11],[12,18],[20,25],[30,32]], carpetLen = 10
 * 输出：9
 * 解释：将毯子从瓷砖 10 开始放置。
 * 总共覆盖 9 块瓷砖，所以返回 9 。
 * 注意可能有其他方案也可以覆盖 9 块瓷砖。
 * 可以看出，瓷砖无法覆盖超过 9 块瓷砖。
 * 示例 2：
 *
 *
 *
 * 输入：tiles = [[10,11],[1,1]], carpetLen = 2
 * 输出：2
 * 解释：将毯子从瓷砖 10 开始放置。
 * 总共覆盖 2 块瓷砖，所以我们返回 2 。
 *
 *
 * 提示：
 *
 * 1 <= tiles.length <= 5 * 104
 * tiles[i].length == 2
 * 1 <= li <= ri <= 109
 * 1 <= carpetLen <= 109
 * tiles 互相 不会重叠 。
 * 通过次数2,135
 * 提交次数8,632
 */

import "sort"

// 参考: https://leetcode.cn/problems/maximum-white-tiles-covered-by-a-carpet/solution/by-tsreaper-gde3/
//
// 贪心+滑窗, 每次只保证整块毯子被覆盖了.
func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	sort.Slice(tiles, func(p, q int) bool {
		return tiles[p][0] < tiles[q][0]
	})
	cur, ans := 0, 0
	for p, q, n := 0, 0, len(tiles); p < n; p++ {
		for ; q < n && tiles[q][1]+1-tiles[p][0] <= carpetLen; q++ {
			cur += tiles[q][1] - tiles[q][0] + 1
		}
		if q < n {
			ans = max(ans, cur+max(0, tiles[p][0]+carpetLen-tiles[q][0]))
		} else {
			ans = max(ans, cur)
		}
		cur -= tiles[p][1] - tiles[p][0] + 1
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
