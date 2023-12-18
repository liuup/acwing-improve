
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

struct node {
    // int from;
    int to;
    int val;
};


int main() {
    ios::sync_with_stdio(false); cin.tie(nullptr); cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    double x1, y1, x2, y2;
    double ans = 0.0;
   
    cin >> x1 >> y1;

    while(cin >> x1 >> y1 >> x2 >> y2) {
        double dx = x1-x2, dy = y1 - y2;
        ans += sqrt(dx * dx + dy * dy) * 2;
    }

    int m = round(ans / 1000 / 20 * 60);
    int h = m / 60;
    m %= 60;

    printf("%d:%02d\n", h, m);

    return 0;
}
