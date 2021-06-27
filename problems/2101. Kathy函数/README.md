## [2101. Kathy函数](https://www.acwing.com/problem/content/2103/)

### 题目

Tiger非常喜欢数学，所以他参加了学校组织的数学课外兴趣小组。

在兴趣小组的学习当中，老师向Tiger介绍了Kathy函数，Kathy函数是这样定义的：

*begin{cases} f(1) = 1 newline f(3) = 3 newline f(2n) = f(n) newline f(4n+1) = 2f(2n+1) - f(n) newline f(4n+3) = 3f(2n+1) - 2f(n) end{cases}*

Tiger对Kathy函数产生了浓厚的兴趣，他通过研究发现有很多的数 *n* 都满足 *f(n) = n*。

对于一个给定的数 *m*，他希望你求出所有的满足 *f(n) = n,(n ≤ m)* 的自然数 *n* 的个数。

### 输入格式

共一行，一个正整数 *m*。

### 输出格式

输出一个整数，表示所有的满足 *f(n) = n,(n ≤ m)* 的自然数 *n* 的个数。

### 数据范围

*1 ≤ m ≤ 10^{100}*

### 输入样例：

```
5
```

### 输出样例：

```
3
```
