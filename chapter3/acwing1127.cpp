
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
auto printvector = [](vector<int> nums) { for(auto x:nums) {cout << x << " ";} cout << endl;};


struct node {
    int to, val;
};

int n, p, c;
int from, to, val;

vector<vector<node>> g;
vi dist;
vi cows;

auto cmp = [](node a, node b){ return a.val > b.val; };

// 求农场到其他农场的距离
int dijkstra(int i) {
    dist = vector<int>(p+1, INT_MAX);
    pq<node, vector<node>, decltype(cmp)> q(cmp);

    q.push(node{i, 0});
    dist[i] = 0;

    while(q.size()) {
        auto cur = q.top(); q.pop();

        if(dist[cur.to] < cur.val) continue;

        for(auto e : g[cur.to]) {
            int d = dist[cur.to] + e.val;
            if(d < dist[e.to]) {
                dist[e.to] = d;
                q.push(node{e.to, d});
            }
        }
    }
    int ans = 0;
    // 求这n头奶牛到牧场的最小距离
    for(int j = 0; j < n; j++) {
        if(dist[cows[j]] == INT_MAX) return INT_MAX;
        ans += dist[cows[j]];
    }

    // cout << ans << endl;
    return ans;
}


int main() {
    ios::sync_with_stdio(false);
    std::cin.tie(nullptr);
    std::cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    cin >> n >> p >> c; // 奶牛数 农场数 道路数
    g = vector<vector<node>>(p+1);
    dist = vector<int>(p+1, INT_MAX);

    int cow;
    for(int i = 0; i < n; i++) {
        cin >> cow;
        cows.push_back(cow);
    }

    while(c--) {
        cin >> from >> to >> val;
        g[from].push_back(node{to, val});
        g[to].push_back(node{from, val});
    }

    int ans = INT_MAX;
    for(int i = 1; i <= p; i++) {   // i是农场的编号 求i到其他各个农场的最小距离
        ans = min(ans, dijkstra(i));
    }

    cout << ans << endl;

    return 0;
}
