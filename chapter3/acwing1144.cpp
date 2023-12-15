
// #define ACM_DEBUG   // comment this line when upload !!!

#pragma GCC optimize(2)

#include <bits/stdc++.h>
using namespace std;

#define LL long long
#define ULL unsigned long long

#define PII pair<int,int>
#define all(a) a.begin(), a.end()

#define umap unordered_map
#define pq priority_queue

#define vi vector<int>
#define vvi vector<vector<int>>
#define pb push_back

#define inf 0x3f3f3f3f

auto printvector = [](vector<int> nums) { for(auto x:nums) {cout << x << " ";} cout << endl;};

// 并查集模板
struct Unionset {
    vector<int> father;
    int count;

    Unionset(int n);
    int Find(int i);
    void Union(int i, int j);
    bool isConnected(int i, int j);
};

Unionset::Unionset(int n) {
    father = vector<int>(n+1);
    for(int i = 1; i <= n; i++) father[i] = i;
    count = n;
}

int Unionset::Find(int i) {
    if(this->father[i] == i) {
        return i;
    }
    this->father[i] = this->Find(this->father[i]);  // 路径压缩
    return this->father[i];
}

void Unionset::Union(int i, int j) {
    int i_fa = this->Find(i);
    int j_fa = this->Find(j);
    if(i_fa == j_fa) {
        return;
    }
    this->father[i_fa] = j_fa;
    this->count--;
}

bool Unionset::isConnected(int i, int j) {
    return this->Find(i) == this->Find(j);
}

int dts[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

struct node {
    int from;
    int to;
    int val;
};

int m, n;
int a, b, c, d;

int nint(int x, int y) {
    return (x-1) * m + (y-1);
}

int main() {
    ios::sync_with_stdio(false); cin.tie(nullptr); cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    cin >> n >> m;

    vector<node> g;

    auto us = Unionset(m*n);

    while(cin >> a >> b >> c >> d) {   // 若干行
        us.Union(nint(a, b), nint(c, d));
    }

    int ans = 0;

    // 竖向合并一遍
    for(int i = 1; i < n; i++) {
        for(int j = 1; j <= m; j++) {
            if(!us.isConnected(nint(i, j), nint(i+1, j))) {
                us.Union(nint(i, j), nint(i+1, j));
                ans++;
            }
        }
    }

    // 横向合并一遍
    for(int i = 1; i <= n; i++) {
        for(int j = 1; j < m; j++) {
            if(!us.isConnected(nint(i, j), nint(i, j+1))) {
                us.Union(nint(i, j), nint(i, j+1));
                ans += 2;
            }
        }
    }

    cout << ans << endl;

    return 0;
}
