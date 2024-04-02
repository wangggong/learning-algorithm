/*
 * @lc app=leetcode.cn id=parsing-a-boolean-expression lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1106] 解析布尔表达式
 *
 * https://leetcode.cn/problems/parsing-a-boolean-expression/description/
 *
 * algorithms
 * Hard (58.78%)
 * Total Accepted:    10.9K
 * Total Submissions: 16.6K
 * Testcase Example:  '"&(|(f))"'
 *
 * 给你一个以字符串形式表述的 布尔表达式（boolean） expression，返回该式的运算结果。
 * 
 * 有效的表达式需遵循以下约定：
 * 
 * 
 * "t"，运算结果为 True
 * "f"，运算结果为 False
 * "!(expr)"，运算过程为对内部表达式 expr 进行逻辑 非的运算（NOT）
 * "&(expr1,expr2,...)"，运算过程为对 2 个或以上内部表达式 expr1, expr2, ... 进行逻辑 与的运算（AND）
 * "|(expr1,expr2,...)"，运算过程为对 2 个或以上内部表达式 expr1, expr2, ... 进行逻辑 或的运算（OR）
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：expression = "!(f)"
 * 输出：true
 * 
 * 
 * 示例 2：
 * 
 * 输入：expression = "|(f,t)"
 * 输出：true
 * 
 * 
 * 示例 3：
 * 
 * 输入：expression = "&(t,f)"
 * 输出：false
 * 
 * 
 * 示例 4：
 * 
 * 输入：expression = "|(&(t,f,t),!(t))"
 * 输出：false
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= expression.length <= 20000
 * expression[i] 由 {'(', ')', '&', '|', '!', 't', 'f', ','} 中的字符组成。
 * expression 是以上述形式给出的有效表达式，表示一个布尔值。
 * 
 * 
 */
public class Solution
{
    public bool ParseBoolExpr(string expression) => ParseBoolExpr(expression.ToCharArray(), 0, expression.Count());

    private bool ParseBoolExpr(char[] expr, int l, int r)
    {
        // assert valid expr
        // assert l < r
        if (expr[l] == 't' || expr[l] == 'f')
        {
            return expr[l] == 't';
        }
        if (expr[l] == '!')
        {
            return !ParseBoolExpr(expr, l + 2, r - 1);
        }
        List<bool> subExpr = new();
        for (int p = l + 2, q = l + 2; p < r - 1; p = q + 1)
        {
            if (expr[p] == 't' || expr[p] == 'f')
            {
                q = p + 1;
            }
            else
            {
                int cur = 1;
                for (q = p + 2; cur != 0; q++)
                {
                    if (expr[q] != '(' && expr[q] != ')')
                    {
                        continue;
                    }
                    cur += expr[q] == '(' ? 1 : -1;
                }
            }
            subExpr.Add(ParseBoolExpr(expr, p, q));
        }
        // assert subExpr.Count() > 0
        bool ans = subExpr.FirstOrDefault();
        foreach (var se in subExpr)
        {
            ans = expr[l] == '|' ? ans || se : ans && se;
        }
        return ans;
    }
}
