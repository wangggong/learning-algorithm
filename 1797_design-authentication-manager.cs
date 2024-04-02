/*
 * @lc app=leetcode.cn id=design-authentication-manager lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1797] 设计一个验证系统
 *
 * https://leetcode.cn/problems/design-authentication-manager/description/
 *
 * algorithms
 * Medium (58.09%)
 * Total Accepted:    11.7K
 * Total Submissions: 19.8K
 * Testcase Example:  '["AuthenticationManager","renew","generate","countUnexpiredTokens","generate","renew","renew","countUnexpiredTokens"]\n' +
  '[[5],["aaa",1],["aaa",2],[6],["bbb",7],["aaa",8],["bbb",10],[15]]'
 *
 * 你需要设计一个包含验证码的验证系统。每一次验证中，用户会收到一个新的验证码，这个验证码在 currentTime 时刻之后 timeToLive
 * 秒过期。如果验证码被更新了，那么它会在 currentTime （可能与之前的 currentTime 不同）时刻延长 timeToLive 秒。
 * 
 * 请你实现 AuthenticationManager 类：
 * 
 * 
 * AuthenticationManager(int timeToLive) 构造 AuthenticationManager 并设置
 * timeToLive 参数。
 * generate(string tokenId, int currentTime) 给定 tokenId ，在当前时间 currentTime
 * 生成一个新的验证码。
 * renew(string tokenId, int currentTime) 将给定 tokenId 且 未过期 的验证码在 currentTime
 * 时刻更新。如果给定 tokenId 对应的验证码不存在或已过期，请你忽略该操作，不会有任何更新操作发生。
 * countUnexpiredTokens(int currentTime) 请返回在给定 currentTime 时刻，未过期 的验证码数目。
 * 
 * 
 * 如果一个验证码在时刻 t 过期，且另一个操作恰好在时刻 t 发生（renew 或者 countUnexpiredTokens 操作），过期事件 优先于
 * 其他操作。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：
 * ["AuthenticationManager", "renew", "generate", "countUnexpiredTokens",
 * "generate", "renew", "renew", "countUnexpiredTokens"]
 * [[5], ["aaa", 1], ["aaa", 2], [6], ["bbb", 7], ["aaa", 8], ["bbb", 10],
 * [15]]
 * 输出：
 * [null, null, null, 1, null, null, null, 0]
 * 
 * 解释：
 * AuthenticationManager authenticationManager = new AuthenticationManager(5);
 * // 构造 AuthenticationManager ，设置 timeToLive = 5 秒。
 * authenticationManager.renew("aaa", 1); // 时刻 1 时，没有验证码的 tokenId 为 "aaa"
 * ，没有验证码被更新。
 * authenticationManager.generate("aaa", 2); // 时刻 2 时，生成一个 tokenId 为 "aaa"
 * 的新验证码。
 * authenticationManager.countUnexpiredTokens(6); // 时刻 6 时，只有 tokenId 为 "aaa"
 * 的验证码未过期，所以返回 1 。
 * authenticationManager.generate("bbb", 7); // 时刻 7 时，生成一个 tokenId 为 "bbb"
 * 的新验证码。
 * authenticationManager.renew("aaa", 8); // tokenId 为 "aaa" 的验证码在时刻 7 过期，且 8
 * >= 7 ，所以时刻 8 的renew 操作被忽略，没有验证码被更新。
 * authenticationManager.renew("bbb", 10); // tokenId 为 "bbb" 的验证码在时刻 10
 * 没有过期，所以 renew 操作会执行，该 token 将在时刻 15 过期。
 * authenticationManager.countUnexpiredTokens(15); // tokenId 为 "bbb" 的验证码在时刻
 * 15 过期，tokenId 为 "aaa" 的验证码在时刻 7 过期，所有验证码均已过期，所以返回 0 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 1 
 * 1 
 * tokenId 只包含小写英文字母。
 * 所有 generate 函数的调用都会包含独一无二的 tokenId 值。
 * 所有函数调用中，currentTime 的值 严格递增 。
 * 所有函数的调用次数总共不超过 2000 次。
 * 
 * 
 */

class TokenInfoListNode
{
    public string Token;
    public int Time;
    public TokenInfoListNode Prev;
    public TokenInfoListNode Next;

    public TokenInfoListNode(string token = "", int time = 0)
    {
        Token = token;
        Time = time;
    }
}

class TokenInfoList
{
    public TokenInfoListNode Head = new();
    public TokenInfoListNode Tail = new();
    public int Count;

    public TokenInfoList()
    {
        Head.Next = Tail;
        Tail.Prev = Head;
        Count = 0;
    }

    public void Append(TokenInfoListNode node)
    {
        node.Prev = Tail.Prev;
        node.Next = Tail;
        Tail.Prev.Next = node;
        Tail.Prev = node;
        Count++;
    }

    public void Remove(TokenInfoListNode node)
    {
        node.Prev.Next = node.Next;
        node.Next.Prev = node.Prev;
        Count--;
    }
}

public class AuthenticationManager
{
    private Dictionary<string, TokenInfoListNode> TokenMap = new();
    private int TimeToLive;
    private TokenInfoList List = new();

    public AuthenticationManager(int timeToLive)
        => TimeToLive = timeToLive;

    public void Generate(string tokenId, int currentTime)
    {
        var node = new TokenInfoListNode(tokenId, currentTime + TimeToLive);
        TokenMap[tokenId] = node;
        List.Append(node);
    }

    public void Renew(string tokenId, int currentTime)
    {
        TokenMap.TryGetValue(tokenId, out var node);
        if (node != null && node.Time > currentTime)
        {
            node.Time = currentTime + TimeToLive;
            List.Remove(node);
            List.Append(node);
        }
    }

    public int CountUnexpiredTokens(int currentTime)
    {
        while (List.Count > 0 && List.Head.Next.Time <= currentTime)
        {
            List.Remove(List.Head.Next);
        }
        return List.Count;
    }
}

/*
 * public class AuthenticationManager
 * {
 *     private int TimeToLive { get; set; }
 *     private List<(string, int)> TokenInfo = new();
 * 
 *     public AuthenticationManager(int timeToLive)
 *         => TimeToLive = timeToLive;
 * 
 *     public void Generate(string tokenId, int currentTime)
 *         => TokenInfo.Add((tokenId, currentTime + TimeToLive));
 * 
 *     public void Renew(string tokenId, int currentTime)
 *     {
 *         for (int i = 0, n = TokenInfo.Count(); i < n; i++)
 *         {
 *             var (curTokenId, time) = TokenInfo[i];
 *             if (curTokenId == tokenId && time > currentTime)
 *             {
 *                 TokenInfo[i] = (curTokenId, currentTime + TimeToLive);
 *                 break;
 *             }
 *         }
 *     }
 * 
 *     public int CountUnexpiredTokens(int currentTime)
 *     {
 *         TokenInfo.Sort((x, y) => x.Item2.CompareTo(y.Item2));
 *         while (TokenInfo.Count() > 0 && TokenInfo[0].Item2 <= currentTime)
 *         {
 *             TokenInfo.RemoveAt(0);
 *         }
 *         return TokenInfo.Count();
 *     }
 * }
 */

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * AuthenticationManager obj = new AuthenticationManager(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * int param_3 = obj.CountUnexpiredTokens(currentTime);
 */
