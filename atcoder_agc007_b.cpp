#include <iostream>

using namespace std;

const int maxn = 2e4 + 10;

int n;
int P[maxn + 10], a[maxn + 10], b[maxn + 10];

int main() {
    int x;
    scanf("%d", &n);
    for (int i = 0; i < n; i++) {
        scanf("%d", &x);
        P[i] = x - 1;
    }
    for (int i = 0; i < n; i++)
        a[i] = maxn * i + 1;
    for (int i = 0; i < n; i++)
        b[P[i]] = i + maxn * maxn - maxn * P[i];
    for (int i = 0; i < n; i++)
        printf("%d ", a[i]);
    puts("");
    for (int i = 0; i < n; i++)
        printf("%d ", b[i]);
    puts("");
    // for (int i = 0; i < n; i++)
    //     cout << a[P[i]] + b[P[i]] << " ";
    // cout << endl;
    return 0;
}
