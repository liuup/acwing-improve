
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

#define INF 0x3f3f3f3f

auto printvector = [](vector<int> nums) { for(auto x:nums) {cout << x << " ";} cout << endl;};

struct edge {
    int to;
    int val;
};

const int MAXN = 50010;

int n, m;
int x, y, t;    // 边 权

vector<vector<edge>> g;

vi rel(7); // 亲戚

int d[7][MAXN];
int ans = INF;

bool vis[7];

void dijkstra(int s) {
    auto cmp = [](edge a, edge b){ return a.val > b.val; };
    pq<edge, vector<edge>, decltype(cmp)> q(cmp);

    d[s][rel[s]] = 0;
    q.push(edge{rel[s], 0});

    while(q.size()) {
        edge cur = q.top(); q.pop();
        if(d[s][cur.to] < cur.val) {
            continue;
        }

        for(auto e : g[cur.to]) {   // 遍历所有孩子
            int dist = d[s][cur.to] + e.val;
            if(dist < d[s][e.to]) {
                d[s][e.to] = dist;
                q.push(edge{e.to, dist});
            }
        }
    }
}

void dfs(int cur, int cost, int pos) {
    if(cost > ans) return ;
    if(cur == 5) {
        ans = min(ans, cost);
        return;
    }
    for(int i = 1; i <= 5; i++) {
        if(!vis[i]) {
            vis[i] = true;
            dfs(cur+1, cost+d[pos][rel[i]], i);
            vis[i] = false;
        }
    }
}

int main() {
    ios::sync_with_stdio(false); cin.tie(nullptr); cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    cin >> n >> m;

    g = vector<vector<edge>>(n+1);

    for(int i = 1; i <= 5; i++) cin >> rel[i];   // 保存一下亲戚 
    rel[6] = 1;

    while(m--) {    // 先建个无向图
        cin >> x >> y >> t;
        g[x].push_back(edge{y, t});
        g[y].push_back(edge{x, t});
    }

    memset(d, 0x3f, sizeof(d)); // 初始化一波

    for(int i = 1; i <= 6; i++) {   // 求每个亲戚到各个点的最短距离
        dijkstra(i);
    }


    // 全排列
    dfs(0, 0, 6);

    cout << ans << endl;

    return 0;
}
