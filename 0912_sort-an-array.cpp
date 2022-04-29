/*
 * @lc app=leetcode.cn id=sort-an-array lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [912] 排序数组
 *
 * https://leetcode-cn.com/problems/sort-an-array/description/
 *
 * algorithms
 * Medium (55.62%)
 * Total Accepted:    353.9K
 * Total Submissions: 636.8K
 * Testcase Example:  '[5,2,3,1]'
 *
 * 给你一个整数数组 nums，请你将该数组升序排列。
 * 
 * 
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [5,2,3,1]
 * 输出：[1,2,3,5]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [5,1,1,2,0,0]
 * 输出：[0,0,1,1,2,5]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 5 * 10^4
 * -5 * 10^4 <= nums[i] <= 5 * 10^4
 * 
 * 
 */

struct Heap {
    vector<int> arr;
    int size;

    Heap() {
        arr = vector<int>(1, 0);
        size = 0;
    }

    void swim(int k) {
        // 傻逼! 这个 `index` 是从 `1` 开始的!
        while (k > 1 && arr[k] < arr[k >> 1]) {
            swap(arr[k], arr[k >> 1]);
            k >>= 1;
        }
        return;
    }

    void sink(int k) {
        while ((k << 1) <= size) {
            int j = k << 1;
            if (j < size && arr[j] > arr[j + 1]) j++;
            if (arr[j] > arr[k]) j = k;
            if (j == k) break;
            swap(arr[k], arr[j]);
            k = j;
        }
        return;
    }

    void push(int n) {
        arr.push_back(n);
        size++;
        swim(size);
        return;
    }

    int pop() {
        int ans = arr[1];
        swap(arr[1], arr[size]);
        arr.pop_back();
        size--;
        sink(1);
        return ans;
    }

    bool empty() { return size > 0; }
};

class Solution {
    void qsort(vector<int> &arr, int l, int r) {
        if (l + 1 >= r) return;
        auto k = partition(arr, l, r);
        qsort(arr, l, k);
        qsort(arr, k + 1, r);
    }

    int partition(vector<int> &arr, int l, int r) {
        auto v = arr[l];
        int start = l;
        for (auto i = l + 1; i < r; i++) {
            if (arr[i] < v) ++start, swap(arr[i], arr[start]);
        }
        swap(arr[l], arr[start]);
        return start;
    }

    vector<int> tmp;

    void mergeSort(vector<int> &arr, int l, int r) {
        if (l + 1 >= r) return;
        auto mid = (l + r) >> 1;
        mergeSort(arr, l, mid);
        mergeSort(arr, mid, r);
        merge(arr, l, r, mid);
        return;
    }

    void merge(vector<int> &arr, int l, int r, int mid) {
        int p = l, q = mid, k = 0;
        while (p < mid && q < r) {
            if (arr[p] < arr[q]) tmp[k] = arr[p], k++, p++;
            else tmp[k] = arr[q], k++, q++;
        }
        while (p < mid) tmp[k] = arr[p], k++, p++;
        while (q < r) tmp[k] = arr[q], k++, q++;
        for (auto i = 0; i < r - l; i++) arr[l + i] = tmp[i];
        return;
    }

public:
    /*
     * vector<int> sortArray(vector<int>& nums) {
     *     // 傻逼! 这里得 shuffle! 不然遇到已排序数组会退化!
     *     random_shuffle(nums.begin(), nums.end());
     *     qsort(nums, 0, nums.size());
     *     return nums;
     * }
     */

    /*
     * vector<int> sortArray(vector<int>& nums) {
     *     // 傻逼! 临时数组需要预开区间! 每次分配不浪费时间嘛!
     *     tmp.resize(nums.size(), 0);
     *     random_shuffle(nums.begin(), nums.end());
     *     mergeSort(nums, 0, nums.size());
     *     return nums;
     * }
     */

    vector<int> sortArray(vector<int>& nums) {
        auto h = Heap();
        for (int n : nums) h.push(n);
        for (int i = 0; i < nums.size(); i++) nums[i] = h.pop();
        return nums;
    }

};
