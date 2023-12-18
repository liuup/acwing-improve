
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
    int from;
    int to;
    int val;
};

int t;
int n;
int x, y, z;


int main() {
    ios::sync_with_stdio(false); cin.tie(nullptr); cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    cin >> t;
    while(t--) {
        cin >> n;
        vector<node> g;
        for(int i = 0; i < n-1; i++) {
            cin >> x >> y >> z;
            g.pb(node{x, y, z});
        }

        auto us = Unionset(n);

        vector<LL> cize(n+10, 1); // 表示其集合大小 保证对根正确
        
        auto cmp = [](node a, node b){ return a.val < b.val; };
        sort(all(g), cmp);

        LL ans = 0;
        for(auto x : g) {
            int u = us.Find(x.from);
            int v = us.Find(x.to);

            if(!us.isConnected(u, v)) {
                ans += (LL)(cize[u] * cize[v] - 1) * (x.val + 1);
                
                us.Union(x.from, x.to);
                cize[u] = cize[v] = cize[u] + cize[v];
            }
        }

        cout << ans << endl;
    }

    return 0;
}