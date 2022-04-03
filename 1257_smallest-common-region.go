/*
 * Medium
 *
 * 给你一些区域列表 regions ，每个列表的第一个区域都包含这个列表内所有其他区域。
 *
 * 很自然地，如果区域 X 包含区域 Y ，那么区域 X  比区域 Y 大。
 *
 * 给定两个区域 region1 和 region2 ，找到同时包含这两个区域的 最小 区域。
 *
 * 如果区域列表中 r1 包含 r2 和 r3 ，那么数据保证 r2 不会包含 r3 。
 *
 * 数据同样保证最小公共区域一定存在。
 *
 *
 *
 * 示例 1：
 *
 * 输入：
 * regions = [["Earth","North America","South America"],
 * ["North America","United States","Canada"],
 * ["United States","New York","Boston"],
 * ["Canada","Ontario","Quebec"],
 * ["South America","Brazil"]],
 * region1 = "Quebec",
 * region2 = "New York"
 * 输出："North America"
 *
 *
 * 提示：
 *
 * 2 <= regions.length <= 10^4
 * region1 != region2
 * 所有字符串只包含英文字母和空格，且最多只有 20 个字母。
 * 通过次数2,041
 * 提交次数3,600
 */

var fa map[string]string

// NOTE: 直接建立 `fa` 哈希即可, 没必要构建字典树结构啊...
//
// 傻逼吧?!
func findSmallestRegion(regions [][]string, region1 string, region2 string) string {
	fa = make(map[string]string)
	for _, r := range regions {
		// assert len(r) > 1;
		for _, s := range r[1:] {
			fa[s] = r[0]
		}
	}
	path1 := find(region1)
	path2 := find(region1)
	p, q := len(path1)-1, len(path2)-1
	for ; p >= 0 && q >= 0; p, q = p-1, q-1 {
		if path1[p] != path2[q] {
			break
		}
	}
	return path1[p+1]
}

func find(s string) []string {
	var ans []string
	for ; len(s) > 0; s = fa[s] {
		ans = append(ans, s)
	}
	return ans
}
