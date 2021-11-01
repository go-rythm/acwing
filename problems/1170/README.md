## [1170\. 排队布局](https://www.acwing.com/problem/content/1172/)

### 题目

当排队等候喂食时，奶牛喜欢和它们的朋友站得靠近些。

农夫约翰有 $N$ 头奶牛，编号从 $1$ 到 $N$，沿一条直线站着等候喂食。

奶牛排在队伍中的顺序和它们的编号是相同的。

因为奶牛相当苗条，所以可能有两头或者更多奶牛站在同一位置上。

如果我们想象奶牛是站在一条数轴上的话，允许有两头或更多奶牛拥有相同的横坐标。

一些奶牛相互间存有好感，它们希望两者之间的距离不超过一个给定的数 $L$。

另一方面，一些奶牛相互间非常反感，它们希望两者间的距离不小于一个给定的数 $D$。

给出 $M\_L$ 条关于两头奶牛间有好感的描述，再给出 $M\_D$ 条关于两头奶牛间存有反感的描述。

你的工作是：如果不存在满足要求的方案，输出-1；如果 $1$ 号奶牛和 $N$ 号奶牛间的距离可以任意大，输出-2；否则，计算出在满足所有要求的情况下，$1$ 号奶牛和 $N$ 号奶牛间可能的最大距离。

### 输入格式

第一行包含三个整数 $N,M\_L,M\_D$。

接下来 $M\_L$ 行，每行包含三个正整数 $A,B,L$，表示奶牛 $A$ 和奶牛 $B$ 至多相隔 $L$ 的距离。

再接下来 $M\_D$ 行，每行包含三个正整数 $A,B,D$，表示奶牛 $A$ 和奶牛 $B$ 至少相隔 $D$ 的距离。

### 输出格式

输出一个整数，如果不存在满足要求的方案，输出-1；如果 $1$ 号奶牛和 $N$ 号奶牛间的距离可以任意大，输出-2；否则，输出在满足所有要求的情况下，$1$ 号奶牛和 $N$ 号奶牛间可能的最大距离。

### 数据范围

$2 \\le N \\le 1000$,

$1 \\le M\_L,M\_D \\le 10^4$,

$1 \\le L,D \\le 10^6$

### 输入样例：

```
4 2 1
1 3 10
2 4 20
2 3 3
```

### 输出样例：

```
27
```

### 题解
