## [1308\. 方程的解](https://www.acwing.com/problem/content/1310/)

### 题目

佳佳碰到了一个难题，请你来帮忙解决。

对于不定方程 $a\_1+a\_2+ \\cdots +a\_{k−1}+a\_k=g(x)$，其中 $k \\ge 1$ 且 $k \\in N^\*$，$x$ 是正整数，$g(x)=x^x \\bmod 1000$（即 $x^x$ 除以 $1000$ 的余数），$x,k$ 是给定的数。

我们要求的是这个不定方程的正整数解组数。

举例来说，当 $k=3,x=2$ 时，方程的解分别为：

$$\\begin{cases}a\_1=1 \\newline a\_2=1 \\newline a\_3=2 \\end{cases} \ \\begin{cases}a\_1=1 \\newline a\_2=2 \\newline a\_3=1 \\end{cases} \ \\begin{cases}a\_1=2 \\newline a\_2=1 \\newline a\_3=1 \\end{cases}$$

### 输入格式

有且只有一行，为用空格隔开的两个正整数，依次为 $k,x$。

### 输出格式

有且只有一行，为方程的正整数解组数。

### 数据范围

$1 \\le k \\le 100$,

$1 \\le x < 2^{31}$,

$k \\le g(x)$

### 输入样例：

```
3 2
```

### 输出样例：

```
3
```

### 题解

