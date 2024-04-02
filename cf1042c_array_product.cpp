#include <iostream>

using namespace std;

const int N = 2e5;
int a[N + 10], zeros[N + 10];

int main() {
    int n, zero = 0, neg = 0, zcnt = 0, maxNeg = INT_MIN, negInd = -1;
    scanf("%d", &n);
    for (int i = 1; i <= n; i++) {
        scanf("%d", a + i);
        if (a[i] == 0)
            zero++, zeros[zcnt++] = i;
        else if (a[i] < 0)
            neg++;
    }
    if (neg % 2 != 0) {
        for (int i = 1; i <= n; i++)
            if (a[i] < 0 && maxNeg < a[i])
                maxNeg = a[i], negInd = i;
        zeros[zcnt++] = negInd;
    }
    for (int i = zcnt - 1; i > 0; i--)
        printf("1 %d %d\n", zeros[i], zeros[i - 1]);
    if (zcnt > 0 && zero + (neg == 1 ? 1 : 0) < n)
        printf("2 %d\n", zeros[0]);
    for (int p = 1, q = 2; p <= n; p = q) {
        for (; p <= n && (a[p] == 0 || p == negInd); p++);
        for (q = p + 1; q <= n && (a[q] == 0 || q == negInd); q++);
        if (q > n)
            break;
        printf("1 %d %d\n", p, q);
    }
    return 0;
}
