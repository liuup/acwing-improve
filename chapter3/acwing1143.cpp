
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


struct node {
    int p;  // 种类
    int from;
    int to;
    int val;
};

int n, m;
int a, b, c, d; // a是类别


int main() {
    ios::sync_with_stdio(false); cin.tie(nullptr); cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    cin >> n >> m;
    
    vector<node> g;
    while(m--) {
        cin >> a >> b >> c >> d;
        g.pb(node{a, b, c, d});
    }

    auto cmp = [](node a, node b){
        if(a.p == b.p) {
            return a.val < b.val;
        }
        return a.p < b.p;
    };

    sort(all(g), cmp);

    int ans = 0;

    auto us = Unionset(n);
    
    for(auto x : g) {
        if(x.p == 1) {
            us.Union(x.from, x.to); // 必选
            ans += x.val;
        } else if(x.p == 2) {
            if(!us.isConnected(x.from, x.to)) {
                us.Union(x.from, x.to);
                ans += x.val;
            }
        }
    }

    cout << ans << endl;

    return 0;
}
