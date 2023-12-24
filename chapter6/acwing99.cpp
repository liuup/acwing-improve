
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
void printb(int a) { cout << bitset<sizeof(a)*8>(a) << endl; }  // 打印数字对应的二进制

struct node {
    // int from;
    int to;
    int val;
};

int n, r;
int x, y, w;

vector<vector<int>> sums(5010, vector<int>(5010));

int main(void) {
    ios::sync_with_stdio(false); cin.tie(nullptr); cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    cin >> n >> r;    
    r = min(5001, r);

    for(int i = 0; i < n; i++) {
        cin >> x >> y >> w;

        sums[x+1][y+1] += w;
    }

    for(int i = 1; i <= 5001; i++) {
        for(int j = 1; j <= 5001; j++) {
            sums[i][j] += sums[i-1][j] + sums[i][j-1] - sums[i-1][j-1];
        }
    }

    int ans = 0;

    for(int i = r; i <= 5001; i++) {
        for(int j = r; j <= 5001; j++) {
            ans = max(ans, sums[i][j] - sums[i-r][j] - sums[i][j-r] + sums[i-r][j-r]);
        }
    }

    cout << ans << endl;


    return 0;
}