## [245\. 你能回答这些问题吗](https://www.acwing.com/problem/content/246/)

### 题目

给定长度为 $N$ 的数列 $A$，以及 $M$ 条指令，每条指令可能是以下两种之一：

1. `1 x y`，查询区间 $[x,y]$ 中的最大连续子段和，即 $\\max\_{x \\le l \\le r \\le y}${$\\sum\\limits^r\_{i=l} A[i]$}。
2. `2 x y`，把 $A[x]$ 改成 $y$。

对于每个查询指令，输出一个整数表示答案。

### 输入格式

第一行两个整数 $N,M$。

第二行 $N$ 个整数 $A[i]$。

接下来 $M$ 行每行 $3$ 个整数 $k,x,y$，$k=1$ 表示查询（此时如果 $x>y$，请交换 $x,y$），$k=2$ 表示修改。

### 输出格式

对于每个查询指令输出一个整数表示答案。

每个答案占一行。

### 数据范围

$N \\le 500000, M \\le 100000$,

$-1000 \\le A[i] \\le 1000$

### 输入样例：

```
5 3
1 2 -3 4 5
1 2 3
2 2 -1
1 3 2
```

### 输出样例：

```
2
-1
```

### 题解

