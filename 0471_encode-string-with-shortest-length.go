/*
 * Medium
 *
 * 给定一个 非空 字符串，将其编码为具有最短长度的字符串。
 *
 * 编码规则是：k[encoded_string]，其中在方括号 encoded_string 中的内容重复 k 次。
 *
 * 注：
 *
 * k 为正整数
 * 如果编码的过程不能使字符串缩短，则不要对其进行编码。如果有多种编码方式，返回 任意一种 即可
 *
 *
 * 示例 1：
 *
 * 输入：s = "aaa"
 * 输出："aaa"
 * 解释：无法将其编码为更短的字符串，因此不进行编码。
 * 示例 2：
 *
 * 输入：s = "aaaaa"
 * 输出："5[a]"
 * 解释："5[a]" 比 "aaaaa" 短 1 个字符。
 * 示例 3：
 *
 * 输入：s = "aaaaaaaaaa"
 * 输出："10[a]"
 * 解释："a9[a]" 或 "9[a]a" 都是合法的编码，和 "10[a]" 一样长度都为 5。
 * 示例 4：
 *
 * 输入：s = "aabcaabcd"
 * 输出："2[aabc]d"
 * 解释："aabc" 出现两次，因此一种答案可以是 "2[aabc]d"。
 * 示例 5：
 *
 * 输入：s = "abbbabbbcabbbabbbc"
 * 输出："2[2[abbb]c]"
 * 解释："abbbabbbc" 出现两次，但是 "abbbabbbc" 可以编码为 "2[abbb]c"，因此一种答案可以是 "2[2[abbb]c]"。
 *
 *
 * 提示：
 *
 * 1 <= s.length <= 150
 * s 由小写英文字母组成
 * 通过次数1,735
 * 提交次数2,876
 */

import (
	"fmt"
	"strings"
)

// 记忆化搜索
//
// 参考: https://leetcode-cn.com/problems/encode-string-with-shortest-length/solution/liang-chong-jie-fa-by-jason-2-65/

const MAXN int = 150

var bytes []byte
var dp [MAXN + 5][MAXN + 5]string

// func encode(s string) string {
// 	bytes = []byte(s)
// 	n := len(bytes)
// 	// assert n > 0;
// 	for i := 0; i < MAXN; i++ {
// 		for j := 0; j < MAXN; j++ {
// 			dp[i][j] = ""
// 		}
// 	}
// 	return dfs(0, n)
// }

// DP, 就是把记忆化搜索的逻辑翻译了一遍.
func encode(s string) string {
	bytes := []byte(s)
	n := len(bytes)
	for i := 0; i < MAXN; i++ {
		for j := 0; j < MAXN; j++ {
			dp[i][j] = ""
		}
	}
	// for i := 0; i < n; i++ {
	for i := n-1; i >= 0; i-- {
		for j := i+1; j <= n; j++ {
			s := string(bytes[i:j])
			dp[i][j] = s
			// 1. 不需要压缩的情况;
			if j-i < 5 {
				continue
			}
			// 2. 当前串本身含有重复节;
			t := s[1:] + s
			ind := strings.Index(t, s)
			if ind+1 < j-i {
				cur := fmt.Sprintf("%v[%s]", (j-i) / (ind+1), dp[i][i + ind + 1])
				if len(cur) < len(dp[i][j]) {
					dp[i][j] = cur
				}
				// continue /* 这里其实确实可以直接 continue 的, 枚举了分割点长度也不会更好... 吧? */
			}
			// 3. 枚举每个分割点, 拼接左右两部分内容;
			for k := i+1; k < j; k++ {
				if len(dp[i][k]) + len(dp[k][j]) < len(dp[i][j]) {
					dp[i][j] = dp[i][k] + dp[k][j]
				}
			}
		}
	}
	return dp[0][n]
}


func dfs(p, q int) string {
	if v := dp[p][q]; len(v) > 0 {
		return v
	}
	var ans string
	defer func() { dp[p][q] = ans }()
	ans = string(bytes[p:q])
	// 给出的压缩写法最短要 4 个字符, 如果过短, 不需要压缩.
	if len(ans) < 5 {
		return ans
	}
	// 如果能找到, 说明当前串本身可压缩.
	//
	// NOTE:
	// 1. `(t[1:]+t).indexOf(t)` 一定非负, 因为一定能匹配上; 仅在 `(t[1:]+t).indexOf(t) < len(t)-1` 才算真正有重复节;
	// 2. 这里也套着一层递归, 见下方注释;
	tmp := ans[1:] + ans
	if ind := strings.Index(tmp, ans); ind < len(ans)-1 {
		// cur := fmt.Sprintf("%v[%s]", len(ans)/(ind+1), ans[:ind+1]) /* 这样会导致无法嵌套压缩 */
		// cur := fmt.Sprintf("%v[%s]", len(ans)/(ind+1), dfs(0, ind+1)) /* 这样也不对, 压缩返回值取的都是统一状态的. 反例: "aabcccccddd", want "aab5[c]ddd", got "aab5[a]ddd" */
		cur := fmt.Sprintf("%v[%s]", len(ans)/(ind+1), dfs(p, p+ind+1))
		if len(cur) < len(ans) {
			ans = cur
		}
	}
	// 枚举每个分割点, 看左右各部分的最优解和一起会不会比当前的解更优.
	for k := p + 1; k < q; k++ {
		s1, s2 := dfs(p, k), dfs(k, q)
		if len(s1)+len(s2) < len(ans) {
			ans = s1 + s2
		}
	}
	return ans
}
