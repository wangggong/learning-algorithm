/*
 * @lc app=leetcode.cn id=letter-combinations-of-a-phone-number lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (57.64%)
 * Total Accepted:    404.5K
 * Total Submissions: 701.7K
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：digits = "23"
 * 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：digits = ""
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：digits = "2"
 * 输出：["a","b","c"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= digits.length <= 4
 * digits[i] 是范围 ['2', '9'] 的一个数字。
 *
 *
 */
var ans []string
var cur []byte
var alphaMap = map[byte][]byte{
	'2': []byte("abc"),
	'3': []byte("def"),
	'4': []byte("ghi"),
	'5': []byte("jkl"),
	'6': []byte("mno"),
	'7': []byte("pqrs"),
	'8': []byte("tuv"),
	'9': []byte("wxyz"),
}

func dfs(bytes []byte) {
	if len(bytes) == 0 {
		ans = append(ans, string(cur))
		return
	}
	// assert len(bytes) > 0;
	for _, b := range alphaMap[bytes[0]] {
		cur = append(cur, b)
		dfs(bytes[1:])
		cur = cur[:len(cur)-1]
	}
}

func letterCombinations(digits string) []string {
	ans, cur = nil, nil
	if len(digits) == 0 {
		return []string{}
	}
	// assert len(digits) > 0 && len(digits) <= 4;
	// assert digits[i] >= '2' && digits[i] <= '9';
	dfs([]byte(digits))
	return ans
}
