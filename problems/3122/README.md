## [3122\. 多项式乘法](https://www.acwing.com/problem/content/3125/)

### 题目

给定一个 $n$ 次多项式 $F(x) = a\_0 + a\_1x + a\_2x^2+…+a\_nx^n$。

以及一个 $m$ 次多项式 $G(x) = b\_0 + b\_1x + b\_2x^2+…+b\_mx^m$。

已知 $H(x) = F(x) \\cdot G(x) = c\_0 + c\_1x + c\_2x^2+…+c\_{n+m}x^{n+m}$。

请你计算并输出 $c\_0,c\_1,…,c\_{n+m}$。

### 输入格式

第一行包含两个整数 $n,m$。

第二行包含 $n+1$ 个整数 $a\_0,a\_1,…,a\_n$。

第三行包含 $m+1$ 个整数 $b\_0,b\_1,…,b\_m$。

### 输出格式

共一行，依次输出 $c\_0,c\_1,…,c\_{n+m}$。

### 数据范围

$1 \\le n,m \\le 10^5$,

$0 \\le a\_i,b\_i \\le 9$

### 输入样例：

```
1 2
1 3
2 2 1
```

### 输出样例：

```
2 8 7 3
```

### 题解

