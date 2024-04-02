#include <iostream>
#include <vector>

using namespace std;

int main() {
    int n, k, p, x;
    cin >> n >> k >> p;
    vector<int> odd, even;
    for (int i = 0; i < n; i++) {
        cin >> x;
        if (x % 2)
            odd.push_back(x);
        else
            even.push_back(x);
    }
    if (odd.size() < k - p || (odd.size() - (k - p)) % 2 || even.size() + (odd.size() - (k - p)) / 2 < p) {
        cout << "NO" << endl;
        return 0;
    }
    cout << "YES" << endl;
    int s = 0, t = 0, i = 0;
    for (; s < odd.size() && i < min(k - p, k - 1); s++, i++)
        cout << 1 << " " << odd[s] << endl;
    for (; t < even.size() && i < k - 1; t++, i++)
        cout << 1 << " " << even[t] << endl;
    for (; s < odd.size() && i < k - 1; s += 2, i++)
        cout << 2 << " " << odd[s] << " " << odd[s + 1] << endl;
    cout << ((odd.size() - s) + (even.size() - t)) << " ";
    for (; s < odd.size(); s++)
        cout << odd[s] << " ";
    for (; t < even.size(); t++)
        cout << even[t] << " ";
    cout << endl;
    return 0;
}
