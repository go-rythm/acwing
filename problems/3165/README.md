## [3165\. 第一类斯特林数](https://www.acwing.com/problem/content/3168/)

### 题目

**第一类斯特林数**（斯特林轮换数）$\\begin{bmatrix} n \\newline k \\end{bmatrix}$ 表示将 $n$ 个两两不同的元素，划分为 $k$ 个非空圆排列的方案数。

现在，给定 $n$ 和 $k$，请你求方案数。

> 圆排列定义：圆排列是排列的一种，指从 $n$ 个不同元素中取出 $m$（$1≤m≤n$）个不同的元素排列成一个环形，既无头也无尾。两个圆排列相同当且仅当所取元素的个数相同并且元素取法一致，在环上的排列顺序一致。

### 输入格式

两个整数 $n$ 和 $k$。

### 输出格式

输出一个整数表示划分方案数。

答案对 $10^9+7$ 取模。

### 数据范围

$1 \\le k \\le n \\le 1000$

### 输入样例：

```
3 2
```

### 输出样例：

```
3
```

### 题解

