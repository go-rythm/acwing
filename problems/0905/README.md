## [905\. 区间选点](https://www.acwing.com/problem/content/907/)

### 题目

给定 $N$ 个闭区间 $[a\_i,b\_i]$，请你在数轴上选择尽量少的点，使得每个区间内至少包含一个选出的点。

输出选择的点的最小数量。

位于区间端点上的点也算作区间内。

### 输入格式

第一行包含整数 $N$，表示区间数。

接下来 $N$ 行，每行包含两个整数 $a\_i,b\_i$，表示一个区间的两个端点。

### 输出格式

输出一个整数，表示所需的点的最小数量。

### 数据范围

$1 \\le N \\le 10^5$,

$-10^9 \\le a\_i \\le b\_i \\le 10^9$

### 输入样例：

```
3
-1 1
2 4
3 5
```

### 输出样例：

```
2
```

### 题解

前置题目：0000

前置知识：无

本题知识：贪心-区间问题

#### 题目分析

思路：

1. 将每个区间按右端点从小到大排序
2. 从前往后依次枚举每个区间
    * 如果当前区间中已经包含点，则pass
    * 否则，选择当前区间的右端点

贪心问题思路不复杂，难点在证明算法的正确性

证明：设题目的答案为 i，算法求出来的是 j，证明 i == j

1. i <= j
    按照上述思路求出的解，可以覆盖所有区间满足题意，i 是满足题意的最小值，故 i <= j
2. i >= j
    按照上述思路求得的 j 个区间**两两不交**，所以覆盖所有区间最少需要 j 个点，故 i >= j
3. i == j
    ∎

#### 
