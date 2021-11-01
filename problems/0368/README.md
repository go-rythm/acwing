## [368\. 银河](https://www.acwing.com/problem/content/370/)

### 题目

银河中的恒星浩如烟海，但是我们只关注那些最亮的恒星。

我们用一个正整数来表示恒星的亮度，数值越大则恒星就越亮，恒星的亮度最暗是 $1$。

现在对于 $N$ 颗我们关注的恒星，有 $M$ 对亮度之间的相对关系已经判明。

你的任务就是求出这 $N$ 颗恒星的亮度值总和至少有多大。

### 输入格式

第一行给出两个整数 $N$ 和 $M$。

之后 $M$ 行，每行三个整数 $T, A, B$，表示一对恒星 $(A, B)$ 之间的亮度关系。恒星的编号从 $1$ 开始。

如果 $T = 1$，说明 $A$ 和 $B$ 亮度相等。

如果 $T = 2$，说明 $A$ 的亮度小于 $B$ 的亮度。

如果 $T = 3$，说明 $A$ 的亮度不小于 $B$ 的亮度。

如果 $T = 4$，说明 $A$ 的亮度大于 $B$ 的亮度。

如果 $T = 5$，说明 $A$ 的亮度不大于 $B$ 的亮度。

### 输出格式

输出一个整数表示结果。

若无解，则输出 $-1$。

### 数据范围

$N \\le 100000,M \\le 100000$

### 输入样例：

```
5 7
1 1 2
2 3 2
4 4 1
3 4 5
5 4 5
2 3 5
4 5 1
```

### 输出样例：

```
11
```

### 题解
