/*
 * @lc app=leetcode.cn id=alphabet-board-path lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1138] 字母板上的路径
 *
 * https://leetcode.cn/problems/alphabet-board-path/description/
 *
 * algorithms
 * Medium (44.42%)
 * Total Accepted:    19.5K
 * Total Submissions: 38.8K
 * Testcase Example:  '"leet"'
 *
 * 我们从一块字母板上的位置 (0, 0) 出发，该坐标对应的字符为 board[0][0]。
 * 
 * 在本题里，字母板为board = ["abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"]，如下所示。
 * 
 * 
 * 
 * 我们可以按下面的指令规则行动：
 * 
 * 
 * 如果方格存在，'U' 意味着将我们的位置上移一行；
 * 如果方格存在，'D' 意味着将我们的位置下移一行；
 * 如果方格存在，'L' 意味着将我们的位置左移一列；
 * 如果方格存在，'R' 意味着将我们的位置右移一列；
 * '!' 会把在我们当前位置 (r, c) 的字符 board[r][c] 添加到答案中。
 * 
 * 
 * （注意，字母板上只存在有字母的位置。）
 * 
 * 返回指令序列，用最小的行动次数让答案和目标 target 相同。你可以返回任何达成目标的路径。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：target = "leet"
 * 输出："DDR!UURRR!!DDD!"
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：target = "code"
 * 输出："RR!DDRR!UUL!R!"
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= target.length <= 100
 * target 仅含有小写英文字母。
 * 
 * 
 */
public class Solution
{
    private const int ColCount = 5;

    public string AlphabetBoardPath(string target)
    {
        var ans = "";
        var cur = 'a';
        foreach (var ch in target)
        {
            var oldX = (int)(ch - 'a') % ColCount;
            var newX = (int)(cur - 'a') % ColCount;
            var oldY = (int)(ch - 'a') / ColCount;
            var newY = (int)(cur - 'a') / ColCount;
            var horizontal = new string("LR"[oldX < newX ? 0 : 1], Math.Abs(oldX - newX));
            var vertical = new string("UD"[oldY < newY ? 0 : 1], Math.Abs(oldY - newY));
            ans += (cur == 'z' ? vertical + horizontal : horizontal + vertical) + "!";
            cur = ch;
        }
        return ans;
    }
}

/*
 * public class Solution
 * {
 *     private const int RowCount = 6;
 *     private const int ColCount = 5;
 *     private const int AlphabetCount = 26;
 *     private string[][] Paths;
 *     private (string, Func<int, int>)[] Actions = new (string, Func<int, int>)[]
 *     {
 *         ("U", x => x - ColCount),
 *         ("D", x => x + ColCount),
 *         ("R", x => (x + 1) % ColCount == 0 ? -1 : x + 1),
 *         ("L", x => x % ColCount == 0 ? -1 : x - 1),
 *     };
 * 
 *     public Solution()
 *     {
 *         Paths = new string[AlphabetCount][];
 *         for (var i = 0; i < AlphabetCount; i++)
 *         {
 *             Paths[i] = new string[AlphabetCount];
 *             Queue<(int, string)> Q = new();
 *             var visit = new bool[AlphabetCount];
 *             Q.Enqueue((i, ""));
 *             visit[i] = true;
 *             bool valid(int k) => 0 <= k && k < AlphabetCount && (!visit[k]);
 *             while (Q.Count > 0)
 *             {
 *                 var (k, s) = Q.Dequeue();
 *                 Paths[i][k] = s;
 *                 foreach (var (act, f) in Actions)
 *                 {
 *                     if (valid(f(k)))
 *                     {
 *                         Q.Enqueue((f(k), s + act));
 *                         visit[f(k)] = true;
 *                     }
 *                 }
 *             }
 *         }
 *     }
 * 
 *     public string AlphabetBoardPath(string target)
 *     {
 *         var ans = "";
 *         var cur = 'a';
 *         foreach (var ch in target)
 *         {
 *             ans += Paths[cur - 'a'][ch - 'a'] + "!";
 *             cur = ch;
 *         }
 *         return ans;
 *     }
 * }
 */
