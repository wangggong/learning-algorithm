/*
 * @lc app=leetcode.cn id=remove-sub-folders-from-the-filesystem lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1233] 删除子文件夹
 *
 * https://leetcode-cn.com/problems/remove-sub-folders-from-the-filesystem/description/
 *
 * algorithms
 * Medium (48.89%)
 * Total Accepted:    9.1K
 * Total Submissions: 18.5K
 * Testcase Example:  '["/a","/a/b","/c/d","/c/d/e","/c/f"]'
 *
 * 你是一位系统管理员，手里有一份文件夹列表 folder，你的任务是要删除该列表中的所有 子文件夹，并以 任意顺序 返回剩下的文件夹。
 *
 * 如果文件夹 folder[i] 位于另一个文件夹 folder[j] 下，那么 folder[i] 就是 folder[j] 的 子文件夹 。
 *
 * 文件夹的「路径」是由一个或多个按以下格式串联形成的字符串：'/' 后跟一个或者多个小写英文字母。
 *
 *
 * 例如，"/leetcode" 和 "/leetcode/problems" 都是有效的路径，而空字符串和 "/" 不是。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：folder = ["/a","/a/b","/c/d","/c/d/e","/c/f"]
 * 输出：["/a","/c/d","/c/f"]
 * 解释："/a/b/" 是 "/a" 的子文件夹，而 "/c/d/e" 是 "/c/d" 的子文件夹。
 *
 *
 * 示例 2：
 *
 *
 * 输入：folder = ["/a","/a/b/c","/a/b/d"]
 * 输出：["/a"]
 * 解释：文件夹 "/a/b/c" 和 "/a/b/d/" 都会被删除，因为它们都是 "/a" 的子文件夹。
 *
 *
 * 示例 3：
 *
 *
 * 输入: folder = ["/a/b/c","/a/b/ca","/a/b/d"]
 * 输出: ["/a/b/c","/a/b/ca","/a/b/d"]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= folder.length <= 4 * 10^4
 * 2 <= folder[i].length <= 100
 * folder[i] 只包含小写字母和 '/'
 * folder[i] 总是以字符 '/' 起始
 * 每个文件夹名都是 唯一 的
 *
 *
 */

import (
	"sort"
)

type PathNode struct {
	Val       string
	HasFolder bool
	Children  map[string]*PathNode
}

var pn *PathNode

/*
 * func removeSubfolders(folder []string) []string {
 * 	sort.Strings(folder)
 * 	n := len(folder)
 * 	ans := make([]string, 0, n)
 * 	pn = &PathNode{Children: make(map[string]*PathNode)}
 * 	for _, f := range folder {
 * 		path := strings.Split(f, "/")
 * 		hasParent := false
 * 		curr := pn
 * 		for _, p := range path {
 * 			if _, ok := curr.Children[p]; !ok {
 * 				curr.Children[p] = &PathNode{Children: make(map[string]*PathNode)}
 * 			}
 * 			if curr.HasFolder {
 * 				hasParent = true
 * 				break
 * 			}
 * 			curr = curr.Children[p]
 * 		}
 * 		if hasParent {
 * 			continue
 * 		} else {
 * 			curr.HasFolder = true
 * 		}
 * 		ans = append(ans, f)
 * 	}
 * 	return ans
 * }
 */

// 只排序不建 Trie:
// 参考 https://leetcode-cn.com/problems/remove-sub-folders-from-the-filesystem/solution/zi-dian-xu-sheng-xu-bu-yong-trie-by-boil-fjgt/
// 核心思路: 排序后 父文件夹 一定在 子文件夹 前面, 连续跟着 (这里是因为分隔符 '/' < 'a', 如果分隔符比较大就得换成 '/').
func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	n := len(folder)
	ans := make([]string, 0, n)
	if len(folder) == 0 {
		return ans
	}
	ans = append(ans, folder[0])
	for i, j := 0, 1; j < n; j++ {
		if li, lj := len(folder[i]), len(folder[j]); li < lj && folder[j][:li] == folder[i] && folder[j][li] == '/' {
			continue
		}
		i = j
		ans = append(ans, folder[j])
	}
	return ans
}
