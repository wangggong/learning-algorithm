/*
 * @lc app=leetcode.cn id=brace-expansion-ii lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1096] 花括号展开 II
 *
 * https://leetcode.cn/problems/brace-expansion-ii/description/
 *
 * algorithms
 * Hard (57.87%)
 * Total Accepted:    12.3K
 * Total Submissions: 16.7K
 * Testcase Example:  '"{a,b}{c,{d,e}}"'
 *
 * 如果你熟悉 Shell 编程，那么一定了解过花括号展开，它可以用来生成任意字符串。
 * 
 * 花括号展开的表达式可以看作一个由 花括号、逗号 和 小写英文字母 组成的字符串，定义下面几条语法规则：
 * 
 * 
 * 如果只给出单一的元素 x，那么表达式表示的字符串就只有 "x"。R(x) = {x}
 * 
 * 
 * 例如，表达式 "a" 表示字符串 "a"。
 * 而表达式 "w" 就表示字符串 "w"。
 * 
 * 
 * 当两个或多个表达式并列，以逗号分隔，我们取这些表达式中元素的并集。R({e_1,e_2,...}) = R(e_1) ∪ R(e_2) ∪
 * ...
 * 
 * 例如，表达式 "{a,b,c}" 表示字符串 "a","b","c"。
 * 而表达式 "{{a,b},{b,c}}" 也可以表示字符串 "a","b","c"。
 * 
 * 
 * 要是两个或多个表达式相接，中间没有隔开时，我们从这些表达式中各取一个元素依次连接形成字符串。R(e_1 + e_2) = {a + b for (a,
 * b) in R(e_1) × R(e_2)}
 * 
 * 例如，表达式 "{a,b}{c,d}" 表示字符串 "ac","ad","bc","bd"。
 * 
 * 
 * 表达式之间允许嵌套，单一元素与表达式的连接也是允许的。
 * 
 * 例如，表达式 "a{b,c,d}" 表示字符串 "ab","ac","ad"​​​​​​。
 * 例如，表达式 "a{b,c}{d,e}f{g,h}" 可以表示字符串 "abdfg", "abdfh", "abefg", "abefh",
 * "acdfg", "acdfh", "acefg", "acefh"。
 * 
 * 
 * 
 * 
 * 给出表示基于给定语法规则的表达式 expression，返回它所表示的所有字符串组成的有序列表。
 * 
 * 假如你希望以「集合」的概念了解此题，也可以通过点击 “显示英文描述” 获取详情。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：expression = "{a,b}{c,{d,e}}"
 * 输出：["ac","ad","ae","bc","bd","be"]
 * 
 * 示例 2：
 * 
 * 
 * 输入：expression = "{{a,z},a{b,c},{ab,z}}"
 * 输出：["a","ab","ac","z"]
 * 解释：输出中 不应 出现重复的组合结果。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= expression.length <= 60
 * expression[i] 由 '{'，'}'，',' 或小写英文字母组成
 * 给出的表达式 expression 用以表示一组基于题目描述中语法构造的字符串
 * 
 * 
 */

/*
 * // 奇怪的递归
 * public class Solution
 * {
 *     public IList<string> BraceExpansionII(string expression)
 *     {
 *         var n = expression.Length;
 *         if (n == 0)
 *         {
 *             return new List<string> { "" };
 *         }
 *         if (expression.IndexOf("{") < 0)
 *         {
 *             return expression.Split(",").ToHashSet().ToList();
 *         }
 *         var sb = new StringBuilder();
 *         var left = 0;
 *         var right = 0;
 *         for (; expression[left] != '{'; left++)
 *         {
 *             sb.Append(expression[left]);
 *         }
 *         var prefix = sb.ToString();
 *         var count = 1;
 *         for (right = left + 1; count != 0; right++)
 *         {
 *             count += expression[right] == '{' ? 1 : (expression[right] == '}' ? -1 : 0);
 *         }
 *         var inffixs = new HashSet<string>();
 *         for (int p = left + 1, q = left + 1; p < right - 1; p = q + 1)
 *         {
 *             for (q = p; q < right - 1 && expression[q] != ','; q++)
 *             {
 *                 if (expression[q] == '{')
 *                 {
 *                     for ((q, count) = (q + 1, 1); count != 0; q++)
 *                     {
 *                         count += expression[q] == '{' ? 1 : (expression[q] == '}' ? -1 : 0);
 *                     }
 *                     q--;
 *                 }
 *             }
 *             foreach (var s in BraceExpansionII(expression.Substring(p, q - p)))
 *             {
 *                 inffixs.Add(s);
 *             }
 *         }
 *         var S = new HashSet<string>();
 *         foreach (var inffix in inffixs)
 *         {
 *             foreach (var suffix in BraceExpansionII(right >= n ? "" : expression.Substring(right)))
 *             {
 *                 S.Add($"{prefix}{inffix}{suffix}");
 *             }
 *         }
 *         return S.OrderBy(s => s).ToList();
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/brace-expansion-ii/solution/hua-gua-hao-zhan-kai-ii-by-leetcode-solu-1s1y/
// 官解, 双栈
public class Solution
{
    Stack<char> OperatorStack = new();
    Stack<HashSet<string>> StringsStack = new();

    private void Calculate()
    {
        var op = OperatorStack.Pop();
        var rhs = StringsStack.Pop();
        var lhs = StringsStack.Pop();
        var ans = new HashSet<string>();
        switch (op)
        {
            case '+':
                foreach (var l in lhs)
                {
                    ans.Add(l);
                }
                foreach (var r in rhs)
                {
                    ans.Add(r);
                }
                break;
            default:
                foreach (var l in lhs)
                {
                    foreach (var r in rhs)
                    {
                        ans.Add(l + r);
                    }
                }
                break;
        }
        StringsStack.Push(ans);
    }

    public IList<string> BraceExpansionII(string expression)
    {
        void addMultiplyOperator(int k)
        {
            if (k > 0 && (expression[k - 1] == '}' || ('a' <= expression[k - 1] && expression[k - 1] <= 'z')))
            {
                OperatorStack.Push('*');
            }
        }
        for (int i = 0, n = expression.Length; i < n; i++)
        {
            switch (expression[i])
            {
                case ',':
                    while (OperatorStack.Count > 0 && OperatorStack.Peek() == '*')
                    {
                        Calculate();
                    }
                    OperatorStack.Push('+');
                    break;
                case '{':
                    addMultiplyOperator(i);
                    OperatorStack.Push('{');
                    break;
                case '}':
                    while (OperatorStack.Count > 0 && OperatorStack.Peek() != '{')
                    {
                        Calculate();
                    }
                    OperatorStack.Pop();
                    break;
                default:
                    addMultiplyOperator(i);
                    StringsStack.Push(new HashSet<string>() { new String(expression[i], 1) });
                    break;
            }
        }
        while (OperatorStack.Count > 0)
        {
            Calculate();
        }
        return StringsStack.Peek().OrderBy(s => s).ToList();
    }
}
