#include <iostream>
#include <algorithm>

using namespace std;
typedef long long ll;
const int N = 3;
ll n[N], tot[N], minVal[N] = {INT_MAX, INT_MAX, INT_MAX}, k;

int main() {
    for (int i = 0; i < N; i++)
        cin >> n[i];
    for (int i = 0; i < N; i++)
        for (int j = 0; j < n[i]; j++) {
            cin >> k;
            minVal[i] = min(minVal[i], k);
            tot[i] += k;
        }
    sort(tot, tot + N);
    sort(minVal, minVal + N);
    cout << tot[0] + tot[1] + tot[2] - min(minVal[0] + minVal[1], tot[0]) * 2 << endl;
    return 0;
}
