#include <iostream>

using namespace std;

int n, d, h;

int main() {
    cin >> n >> d >> h;
    if ((h == 1 && d == 1 && n > 2) || d > 2 * h) {
        cout << -1 << endl;
        return 0;
    }
    for (int i = 1; i <= h; i++)
        printf("%d %d\n", i, i + 1);
    int t = h + 2;
    if (d > h) {
        printf("%d %d\n", 1, h + 2);
        for (int i = h + 2; i <= d; i++)
            printf("%d %d\n", i, i + 1);
        t = d + 2;
    }
    for (int i = t; i <= n; i++)
        printf("%d %d\n", (d > h ? 1 : 2), i);
    return 0;
}
