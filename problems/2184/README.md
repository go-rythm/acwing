## [2184\. 餐巾计划问题](https://www.acwing.com/problem/content/2186/)

### 题目

一个餐厅在相继的 $N$ 天里，每天需用的餐巾数不尽相同。

假设第 $i$ 天需要 $r\_i$ 块餐巾 $(i=1,2,…,N)$。

餐厅可以购买新的餐巾，每块餐巾的费用为 $p$ 分；或者把旧餐巾送到快洗部，洗一块需 $m$ 天，其费用为 $f$ 分；或者送到慢洗部，洗一块需 $n$ 天，其费用为 $s$ 分。

餐厅每天使用的餐巾必须是今天刚购买的，或者是今天刚洗好的，且必须恰好提供 $r\_i$ 块毛巾，不能多也不能少。

每天结束时，餐厅必须决定将多少块脏的餐巾送到快洗部，多少块餐巾送到慢洗部，以及多少块保存起来延期送洗。

但是每天洗好的餐巾和购买的新餐巾数之和，要满足当天的需求量。

试设计一个算法为餐厅合理地安排好 $N$ 天中餐巾使用计划，使总的花费最小。

### 输入格式

第 $1$ 行有 $6$ 个正整数 $N,p,m,f,n,s$。$N$ 是要安排餐巾使用计划的天数；$p$ 是每块新餐巾的费用；$m$ 是快洗部洗一块餐巾需用天数；$f$ 是快洗部洗一块餐巾需要的费用；$n$ 是慢洗部洗一块餐巾需用天数；$s$ 是慢洗部洗一块餐巾需要的费用。

接下来的 $N$ 行是餐厅在相继的 $N$ 天里，每天需用的餐巾数。

### 输出格式

输出餐厅在相继的 $N$ 天里使用餐巾的最小总花费。

### 数据范围

$1 \\le N \\le 800$,

$1 \\le s < f < p \\le 50$,

$1 \\le m \\le n \\le 20$,

每天需用的餐巾数不超过 $1000$。

### 输入样例：

```
3 10 2 3 3 2
5
6
7
```

### 输出样例：

```
145
```

### 题解
