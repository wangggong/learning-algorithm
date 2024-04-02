#include <iostream>
#include <stack>

using namespace std;

const int N = 1e5;
int a[N + 10], n, ans;

int main() {
    cin >> n;
    for (int i = 0; i < n; i++)
        cin >> a[i];
    stack<int> S;
    for (int i = 0; i < n; i++) {
        while (!S.empty() && S.top() < a[i])
            ans = max(ans, S.top() ^ a[i]), S.pop();
        if (!S.empty())
            ans = max(ans, S.top() ^ a[i]);
        S.push(a[i]);
    }
    cout << ans << endl;
    return 0;
}
