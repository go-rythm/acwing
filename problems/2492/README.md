## [2492\. HH的项链](https://www.acwing.com/problem/content/2494/)

### 题目

HH 有一串由各种漂亮的贝壳组成的项链。

HH 相信不同的贝壳会带来好运，所以每次散步完后，他都会随意取出一段贝壳，思考它们所表达的含义。

HH 不断地收集新的贝壳，因此他的项链变得越来越长。

有一天，他突然提出了一个问题：某一段贝壳中，包含了多少种不同的贝壳？

这个问题很难回答，因为项链实在是太长了。

于是，他只好求助睿智的你，来解决这个问题。

### 输入格式

第一行：一个整数 $N$，表示项链的长度。

第二行：$N$ 个整数，表示依次表示项链中贝壳的编号（编号为 $0$ 到 $1000000$ 之间的整数）。

第三行：一个整数 $M$，表示 HH 询问的个数。

接下来 $M$ 行：每行两个整数，$L$ 和 $R$，表示询问的区间。

### 输出格式

$M$ 行，每行一个整数，依次表示询问对应的答案。

### 数据范围

$1 \\le N \\le 50000$,

$1 \\le M \\le 2 \\times 10^5$,

$1 \\le L \\le R \\le N$

### 输入样例：

```
6
1 2 3 4 3 5
3
1 2
3 5
2 6
```

### 输出样例：

```
2
2
4
```

### 题解

