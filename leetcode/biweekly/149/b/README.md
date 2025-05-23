![lc3439.png](https://pic.leetcode.cn/1738458316-BLfZXa-lc3439.png)

看示例 1，把会议区间 $[1,2]$ 移动到 $[0,1]$ 或者 $[2,3]$，会产生空余时间段 $[1,3]$ 或者 $[0,2]$，相当于把两个**相邻**的长为 $1$ 空余时间段 $[0,1]$ 和 $[2,3]$ **合并**成一个更大的长为 $1+1=2$ 的空余时间段。

题目要求会议之间的相对顺序需要保持不变，这意味着我们只能合并相邻的空余时间段，所以重新安排至多 $k$ 个会议等价于如下问题：

- 给你 $n+1$ 个空余时间段，合并其中 $k+1$ 个**连续**的空余时间段，得到的最大长度是多少？

这可以用**定长滑动窗口**解决，原理见[【套路】教你解决定长滑窗！适用于所有定长滑窗题目！](https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/solutions/2809359/tao-lu-jiao-ni-jie-jue-ding-chang-hua-ch-fzfo/)

代码实现时，可以用一个数组保存所有的空余时间段的长度。也可以直接写一个函数获取各个空余时间段的长度，从而做到 $\mathcal{O}(1)$ 空间复杂度。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1eUF6eaERQ/?t=38s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxFreeTime(self, eventTime: int, k: int, startTime: List[int], endTime: List[int]) -> int:
        def get(i: int) -> int:
            if i == 0:
                return startTime[0]  # 最左边的空余时间段
            if i == n:
                return eventTime - endTime[-1]  # 最右边的空余时间段
            return startTime[i] - endTime[i - 1]  # 中间的空余时间段

        n = len(startTime)
        ans = s = 0
        for i in range(n + 1):
            s += get(i)
            if i < k:
                continue
            ans = max(ans, s)
            s -= get(i - k)
        return ans
```

```py [sol-Python3 写法二]
class Solution:
    def maxFreeTime(self, eventTime: int, k: int, startTime: List[int], endTime: List[int]) -> int:
        free = [startTime[0]] + [s - e for s, e in zip(startTime[1:], endTime)] + [eventTime - endTime[-1]]
        ans = s = 0
        for i, f in enumerate(free):
            s += f
            if i < k:
                continue
            ans = max(ans, s)
            s -= free[i - k]
        return ans
```

```java [sol-Java]
class Solution {
    public int maxFreeTime(int eventTime, int k, int[] startTime, int[] endTime) {
        int ans = 0;
        int s = 0;
        for (int i = 0; i <= startTime.length; i++) {
            s += get(i, eventTime, startTime, endTime);
            if (i < k) {
                continue;
            }
            ans = Math.max(ans, s);
            s -= get(i - k, eventTime, startTime, endTime);
        }
        return ans;
    }

    private int get(int i, int eventTime, int[] startTime, int[] endTime) {
        if (i == 0) {
            return startTime[0]; // 最左边的空余时间段
        }
        int n = startTime.length;
        if (i == n) {
            return eventTime - endTime[n - 1]; // 最右边的空余时间段
        }
        return startTime[i] - endTime[i - 1]; // 中间的空余时间段
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxFreeTime(int eventTime, int k, vector<int>& startTime, vector<int>& endTime) {
        int n = startTime.size();
        auto get = [&](int i) -> int {
            if (i == 0) {
                return startTime[0]; // 最左边的空余时间段
            }
            if (i == n) {
                return eventTime - endTime[n - 1]; // 最右边的空余时间段
            }
            return startTime[i] - endTime[i - 1]; // 中间的空余时间段
        };

        int s = 0, ans = 0;
        for (int i = 0; i <= n; i++) {
            s += get(i);
            if (i < k) {
                continue;
            }
            ans = max(ans, s);
            s -= get(i - k);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxFreeTime(eventTime, k int, startTime, endTime []int) (ans int) {
	n := len(startTime)
	get := func(i int) int {
		if i == 0 {
			return startTime[0] // 最左边的空余时间段
		}
		if i == n {
			return eventTime - endTime[n-1] // 最右边的空余时间段
		}
		return startTime[i] - endTime[i-1] // 中间的空余时间段
	}

	s := 0
	for i := range n + 1 {
		s += get(i)
		if i < k {
			continue
		}
		ans = max(ans, s)
		s -= get(i - k)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{startTime}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

更多相似题目，见下面滑动窗口题单。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
