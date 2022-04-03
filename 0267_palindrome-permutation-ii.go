/*
 * Medium
 *
 * 给定一个字符串 s ，返回 其重新排列组合后可能构成的所有回文字符串，并去除重复的组合 。
 *
 * 你可以按 任意顺序 返回答案。如果 s 不能形成任何回文排列时，则返回一个空列表。
 *
 *
 *
 * 示例 1：
 *
 * 输入: s = "aabb"
 * 输出: ["abba", "baab"]
 * 示例 2：
 *
 * 输入: s = "abc"
 * 输出: []
 *
 *
 * 提示：
 *
 * 1 <= s.length <= 16
 * s 仅由小写英文字母组成
 * 通过次数4,070
 * 提交次数8,999
 */

// import "sort"

var counter [1 << 8]int
var ans []string
var cur []byte
var single byte
var visit [20]bool

func generatePalindromes(s string) []string {
	bytes := []byte(s)
	ans, cur = nil, nil
	for i := range counter {
		counter[i] = 0
	}
	for i := range visit {
		visit[i] = false
	}
	for _, b := range bytes {
		counter[int(b)]++
	}
	single = byte(0)
	var newBytes []byte
	for b := byte('a'); b <= 'z'; b++ {
		c := counter[int(b)]
		if c%2 != 0 {
			if single != 0 {
				return ans
			}
			single = b
			c--
		}
		for i := 0; i*2 < c; i++ {
			newBytes = append(newBytes, b)
		}
	}
	// 千古谜题: 为啥不需要排序啊?
	// sort.Slice(newBytes, func(p, q int) bool { return newBytes[p] < newBytes[q] })
	backtrack(newBytes, 0, len(newBytes))
	return ans
}

func backtrack(arr []byte, k, n int) {
	if k == n {
		tmp := make([]byte, n)
		copy(tmp, cur)
		if single != 0 {
			tmp = append(tmp, single)
		}
		for i := n - 1; i >= 0; i-- {
			tmp = append(tmp, tmp[i])
		}
		ans = append(ans, string(tmp))
		return
	}
	for i := 0; i < n; i++ {
		if visit[i] {
			continue
		}
		if i > 0 && arr[i] == arr[i-1] && !visit[i-1] {
			continue
		}
		visit[i] = true
		cur = append(cur, arr[i])
		backtrack(arr, k+1, n)
		visit[i] = false
		cur = cur[:len(cur)-1]
	}
}

