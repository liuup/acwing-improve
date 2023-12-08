
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


int n, m;
int from, to, val;

vector<vector<node>> g;
vector<int> dist;


int main() {
    ios::sync_with_stdio(false);
    std::cin.tie(nullptr);
    std::cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    cin >> n >> m;

    g = vector<vector<node>>(n+1);
    dist = vector<int>(n+1, INT_MAX);

    while(m--) {
        cin >> from >> to >> val;
        g[from].push_back(node{to, val});
        g[to].push_back(node{from, val});
    }

    auto cmp = [](node a, node b){ return a.val > b.val; };
    pq<node, vector<node>, decltype(cmp)> q(cmp);

    q.push(node{1, 0});
    dist[1] = 0;

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

    auto ans = -1;
    for(int i = 1; i <= n; i++) {
        if(dist[i] == INT_MAX) {
            cout << -1 << endl;
            return 0;
        }
        ans = max(ans, dist[i]);
    }

    cout << ans << endl;


    return 0;
}
