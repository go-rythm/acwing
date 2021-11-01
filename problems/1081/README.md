## [1081\. 度的数量](https://www.acwing.com/problem/content/1083/)

### 题目

求给定区间 $[X,Y]$ 中满足下列条件的整数个数：这个数恰好等于 $K$ 个互不相等的 $B$ 的整数次幂之和。

例如，设 $X = 15, Y = 20, K = 2, B = 2$，则有且仅有下列三个数满足题意：

$17 = 2^4 + 2^0$

$18 = 2^4 + 2^1$

$20 = 2^4 + 2^2$

### 输入格式

第一行包含两个整数 $X$ 和 $Y$，接下来两行包含整数 $K$ 和 $B$。

### 输出格式

只包含一个整数，表示满足条件的数的个数。

### 数据范围

$1 \\le X \\le Y \\le 2^{31}-1$,

$1 \\le K \\le 20$,

$2 \\le B \\le 10$

### 输入样例：

```
15 20
2
2
```

### 输出样例：

```
3
```

### 题解
