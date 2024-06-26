/*
 * Medium
 * 
 * 给定一个放有字母和数字的数组，找到最长的子数组，且包含的字母和数字的个数相同。
 * 
 * 返回该子数组，若存在多个最长子数组，返回左端点下标值最小的子数组。若不存在这样的数组，返回一个空数组。
 * 
 * 示例 1:
 * 
 * 输入: ["A","1","B","C","D","2","3","4","E","5","F","G","6","7","H","I","J","K","L","M"]
 * 
 * 输出: ["A","1","B","C","D","2","3","4","E","5","F","G","6","7"]
 * 示例 2:
 * 
 * 输入: ["A","A"]
 * 
 * 输出: []
 * 提示：
 * 
 * array.length <= 100000
 */

public class Solution
{
    public string[] FindLongestSubarray(string[] array)
    {
        var ans = new string[0];
        var n = array.Length;
        var cur = 0;
        var d = new Dictionary<int, int>();
        d[0] = 0;
        for (var i = 0; i < n; i++)
        {
            var ch = array[i][0];
            if (('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z'))
            {
                cur++;
            }
            else
            {
                cur--;
            }
            if (d.ContainsKey(cur))
            {
                if (ans.Length < i - d[cur] + 1)
                {
                    ans = array[d[cur]..(i + 1)];
                }
            }
            else
            {
                d[cur] = i + 1;
            }
        }
        return ans;
    }
}
