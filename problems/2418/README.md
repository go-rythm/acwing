## [2418\. 光之大陆](https://www.acwing.com/problem/content/2420/)

### 题目

在光之大陆的土地上，各种势力盘根错节。

来自光之峡谷的精灵，来自黑暗森林的亡灵，来自古老东方的人类共同生活在一起。

善于打造装置的矮人，善于发明的侏儒，隐匿于山林的巨人也坚守着属于自己的领土。

这些种族之间关系错综复杂，构成了极其庞大的关系网络。

大魔法师小 $P$ 想要研究其中的种族关系。

两个物种之间可以是盟友，也可以不是盟友，如果 $a\_1,a\_2..a\_n$ 满足 $a\_i$ 和 $a\_{i+1}$ 是盟友，且 $a\_n$ 和 $a\_1$ 是盟友，则他们构成了一个联盟。

由于光之大陆正处于微妙的和平之中。所以一个合理的物种关系应满足如下条件:

1. 对于任意两个物种 $A,B$，都存在一个序列 $A,a\_1,a\_2..a\_n,B$，使得任意相邻两个种族是盟友(注意 $A,B$ 不一定是盟友)。
2. 对于任意两个联盟 $S\_a,S\_b$，都不存在一个物种既参加了联盟 $S\_a$，又参加了联盟 $S\_b$。

小 $P$ 想知道，大陆上的 $N$ 个种族一共有多少种可能的结盟关系，由于结果可能很大，你只需要输出答案 $\\bmod M$ 的值。

### 输入格式

一行，两个正整数 $N,M$。

### 输出格式

一个整数，表示方案数 $\\bmod M$ 的值。

### 数据范围

$3 \\le N \\le 200$,

$1 \\le M \\le 10^6$

### 输入样例：

```
4 1000000
```

### 输出样例：

```
31
```

### 题解
