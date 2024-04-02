/*
 * Medium
 *
 * 二进制数转字符串。给定一个介于0和1之间的实数（如0.72），类型为double，打印它的二进制表达式。如果该数字无法精确地用32位以内的二进制表示，则打印“ERROR”。
 * 
 * 示例1:
 * 
 *  输入：0.625
 *  输出："0.101"
 * 示例2:
 * 
 *  输入：0.1
 *  输出："ERROR"
 *  提示：0.1无法被二进制准确表示
 *  
 * 
 * 提示：
 * 
 * 32位包括输出中的 "0." 这两位。
 * 题目保证输入用例的小数位数最多只有 6 位
 */

public class Solution
{
    private const int Maxn = (int)1e6;
    private int Gcd(int x, int y) => y == 0 ? x : Gcd(y, x % y);

    public string PrintBin(double val)
    {
        var S = new HashSet<(int, int)>();
        var sb = new StringBuilder();
        var k = (int)(val * Maxn);
        var gcd = Gcd(k, Maxn);
        for (var (num, den) = (k / gcd, Maxn / gcd); num != 0; num %= den)
        {
            if (S.Contains((num, den)))
            {
                return "ERROR";
            }
            S.Add((num, den));
            num <<= 1;
            sb.Append(num >= den ? '1' : '0');
            num %= den;
        }
        return "0." + sb.ToString();
    }
}
