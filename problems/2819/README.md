## [2819\. 动态逆序对](https://www.acwing.com/problem/content/2821/)

### 题目

对于序列 $A$，它的逆序对数定义为满足 $i < j$，且 $A\_i > A\_j$ 的数对 $(i,j)$ 的个数。

给 $1$ 到 $n$ 的一个排列，按照某种顺序依次删除 $m$ 个元素，你的任务是在每次删除一个元素之前统计整个序列的逆序对数。

### 输入格式

第一行包含两个整数 $n$ 和 $m$，即初始元素的个数和删除的元素个数。

以下 $n$ 行每行包含一个 $1$ 到 $n$ 之间的正整数，即初始排列。

以下 $m$ 行每行一个正整数，依次为每次删除的元素。

### 输出格式

输出包含 $m$ 行，依次为删除每个元素之前，逆序对的个数。

### 数据范围

$1 \\le n \\le 10^5$,

$1 \\le m \\le 50000$

### 输入样例：

```
5 4
1
5
3
4
2
5
1
4
2
```

### 输出样例：

```
5
2
2
1
```

### 样例解释

给定序列变化依次为 $(1,5,3,4,2) \\to (1,3,4,2) \\to (3,4,2) \\to (3,2) \\to (3)$。

### 题解

