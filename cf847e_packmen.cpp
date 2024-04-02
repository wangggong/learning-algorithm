#include <iostream>

using namespace std;

bool check(string s, int k) {
    int p = 0, q = 0, n = s.size();
    for (; p < n; p++) {
        for (; p < n && s[p] != 'P'; p++);
        for (; q < n && s[q] != '*'; q++);
        if (p >= n || q >= n)
            break;
        if (q < p) {
            if (q < p - k)
                return false;
            // left: t = q + \delta = q + (k - (p - q)) = k - p + 2 * q
            // right: k = (t - p) + (t - q) = (t - p) * 2 + (p - q) => t = (k + q - p) / 2 + p
            q = max(k - p + 2 * q, (k + q - p) / 2 + p) + 1;
        } else
            q = max(q, p + k + 1);
    }
    return q >= n;
}

int main() {
    int n;
    string s;
    cin >> n >> s;
    int p = 0, q = n << 1;
    while (p < q) {
        int mid = p + q >> 1;
        if (check(s, mid))
            q = mid;
        else
            p = mid + 1;
    }
    cout << p << endl;
    return 0;
}
