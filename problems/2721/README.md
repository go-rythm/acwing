## [2721\. K-单调](https://www.acwing.com/problem/content/2723/)

### 题目

如果一个整数序列中的每一项都严格大于前一项，则称该序列是严格单调递增的。

类似地，如果一个整数序列中的每一项都严格小于前一项，则称该序列是严格单调递减的。

如果一个序列能够被分为 $k$ 个不相交的相邻子序列，并且每个子序列都严格单调，那么就称这个序列是 $k-$单调的。

例如，一个严格单调递增序列是 $1-$单调的----事实上，对于任意 $k \\in [1,序列包含元素个数]$，它都是 $k-$单调的。

序列 $\\{1,2,3,2,1\\}$ 是 $2-$单调的，因为它可以被拆分为 $\\{1,2,3\\}$ 和 $\\{2,1\\}$。

如果一个序列不是 $k-$单调的，则可以通过执行一次或多次以下操作将其转换为 $k-$单调序列：

选择序列中的任意一项，将其加一或减一。

你可以对序列中的任意项执行任意次操作。

给定一个整数序列 $A\_1,A\_2,…,A\_n$ 和整数 $k$，请计算将给定序列转换为 $k-$单调序列，所需要的最少操作数。

### 输入格式

第一行包含两个整数 $n$ 和 $k$。

第二行包含 $n$ 个整数 $A\_1,A\_2,…,A\_n$。

### 输出格式

一个整数，表示最少操作数。

### 数据范围

$1 \\le n \\le 1000$,

$1 \\le k \\le \\min(n,10)$,

$-10^5 \\le A\_i \\le 10^5$

### 输入样例：

```
6 1
1 2 3 3 2 1
```

### 输出样例：

```
9
```

### 题解
