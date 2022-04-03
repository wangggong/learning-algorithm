/*
 * @lc app=leetcode.cn id=stock-price-fluctuation lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2034] 股票价格波动
 *
 * https://leetcode-cn.com/problems/stock-price-fluctuation/description/
 *
 * algorithms
 * Medium (31.01%)
 * Total Accepted:    20.1K
 * Total Submissions: 42.5K
 * Testcase Example:  '["StockPrice","update","update","current","maximum","update","maximum","update","minimum"]\n' +
  '[[],[1,10],[2,5],[],[],[1,3],[],[4,2],[]]'
 *
 * 给你一支股票价格的数据流。数据流中每一条记录包含一个 时间戳 和该时间点股票对应的 价格 。
 * 
 * 
 * 不巧的是，由于股票市场内在的波动性，股票价格记录可能不是按时间顺序到来的。某些情况下，有的记录可能是错的。如果两个有相同时间戳的记录出现在数据流中，前一条记录视为错误记录，后出现的记录
 * 更正 前一条错误的记录。
 * 
 * 请你设计一个算法，实现：
 * 
 * 
 * 更新 股票在某一时间戳的股票价格，如果有之前同一时间戳的价格，这一操作将 更正 之前的错误价格。
 * 找到当前记录里 最新股票价格 。最新股票价格 定义为时间戳最晚的股票价格。
 * 找到当前记录里股票的 最高价格 。
 * 找到当前记录里股票的 最低价格 。
 * 
 * 
 * 请你实现 StockPrice 类：
 * 
 * 
 * StockPrice() 初始化对象，当前无股票价格记录。
 * void update(int timestamp, int price) 在时间点 timestamp 更新股票价格为 price 。
 * int current() 返回股票 最新价格 。
 * int maximum() 返回股票 最高价格 。
 * int minimum() 返回股票 最低价格 。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：
 * ["StockPrice", "update", "update", "current", "maximum", "update",
 * "maximum", "update", "minimum"]
 * [[], [1, 10], [2, 5], [], [], [1, 3], [], [4, 2], []]
 * 输出：
 * [null, null, null, 5, 10, null, 5, null, 2]
 * 
 * 解释：
 * StockPrice stockPrice = new StockPrice();
 * stockPrice.update(1, 10); // 时间戳为 [1] ，对应的股票价格为 [10] 。
 * stockPrice.update(2, 5);  // 时间戳为 [1,2] ，对应的股票价格为 [10,5] 。
 * stockPrice.current();     // 返回 5 ，最新时间戳为 2 ，对应价格为 5 。
 * stockPrice.maximum();     // 返回 10 ，最高价格的时间戳为 1 ，价格为 10 。
 * stockPrice.update(1, 3);  // 之前时间戳为 1 的价格错误，价格更新为 3 。
 * ⁠                         // 时间戳为 [1,2] ，对应股票价格为 [3,5] 。
 * stockPrice.maximum();     // 返回 5 ，更正后最高价格为 5 。
 * stockPrice.update(4, 2);  // 时间戳为 [1,2,4] ，对应价格为 [3,5,2] 。
 * stockPrice.minimum();     // 返回 2 ，最低价格时间戳为 4 ，价格为 2 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= timestamp, price <= 10^9
 * update，current，maximum 和 minimum 总 调用次数不超过 10^5 。
 * current，maximum 和 minimum 被调用时，update 操作 至少 已经被调用过 一次 。
 * 
 * 
 */

const int MAXN = 1e9 + 5;

class StockPrice {
public:
    unordered_map<int, int> priceMap;
    multiset<int> priceSet;
    int currentTimestamp;
    int currentPrice;
    int maximumPrice;
    int minimumPrice;

    StockPrice() {
        // assert price > 0 && price <= 1e9;
        minimumPrice = MAXN;
    }

    void _string() {
        cout << currentTimestamp << "\t" << currentPrice << "\t" << maximumPrice << "\t" << minimumPrice << "\t";
        for (auto it = priceMap.begin(); it != priceMap.end(); ++it) {
            cout << "(" << it->first << "," << it->second << "),";
        }
        cout << endl;
    }

    void update(int timestamp, int price) {
        if (priceMap.find(timestamp) != priceMap.end())
            _update(timestamp, price);
        else
            _insert(timestamp, price);
        // _string();
    }

    void _insert(int timestamp, int price) {
        priceMap[timestamp] = price;
        if (timestamp >= currentTimestamp) {
            currentTimestamp = timestamp;
            currentPrice = price;
        }
        maximumPrice = max(maximumPrice, price);
        minimumPrice = min(minimumPrice, price);
        priceSet.emplace(price);
    }

    void _update(int timestamp, int price) {
        auto formerPrice = priceMap[timestamp];
        // 去除 multiset 中的原始记录.
        auto it = priceSet.find(formerPrice);
        if (it != priceSet.end()) priceSet.erase(it);
        _insert(timestamp, price);
        // 在出现价格更新行为时需要额外更新最大 / 最小值.
        // `O(n)` 的更新方式会 TLE, 所以需要额外维护一个 *multiset* (对应更新时间 `O(logn)`).
        if (formerPrice == maximumPrice) {
            /*
             * maximumPrice = 0;
             * for (auto it = priceMap.begin(); it != priceMap.end(); ++it) {
             *     maximumPrice = max(maximumPrice, it->second);
             * }
             */
            maximumPrice = *priceSet.rbegin();
        }
        if (formerPrice == minimumPrice) {
            /*
             * minimumPrice = MAXN;
             * for (auto it = priceMap.begin(); it != priceMap.end(); ++it) {
             *     minimumPrice = min(minimumPrice, it->second);
             * }
             */
            minimumPrice = *priceSet.begin();
        }
    }

    int current() {
        return currentPrice;
    }

    int maximum() {
        return maximumPrice;
    }

    int minimum() {
        return minimumPrice;
    }
};

/**
 * Your StockPrice object will be instantiated and called as such:
 * StockPrice* obj = new StockPrice();
 * obj->update(timestamp,price);
 * int param_2 = obj->current();
 * int param_3 = obj->maximum();
 * int param_4 = obj->minimum();
 */
