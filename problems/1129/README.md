## [1129\. 热浪](https://www.acwing.com/problem/content/1131/)

### 题目

德克萨斯纯朴的民众们这个夏天正在遭受巨大的热浪！！！

他们的德克萨斯长角牛吃起来不错，可是它们并不是很擅长生产富含奶油的乳制品。

农夫John此时身先士卒地承担起向德克萨斯运送大量的营养冰凉的牛奶的重任，以减轻德克萨斯人忍受酷暑的痛苦。

John已经研究过可以把牛奶从威斯康星运送到德克萨斯州的路线。

这些路线包括起始点和终点一共有 $T$ 个城镇，为了方便标号为 $1$ 到 $T$。

除了起点和终点外的每个城镇都由 **双向道路** 连向至少两个其它的城镇。

每条道路有一个通过费用（包括油费，过路费等等）。

给定一个地图，包含 $C$ 条直接连接 $2$ 个城镇的道路。

每条道路由道路的起点 $R\_s$，终点 $R\_e$ 和花费 $C\_i$ 组成。

求从起始的城镇 $T\_s$ 到终点的城镇 $T\_e$ 最小的总费用。

### 输入格式

第一行: $4$ 个由空格隔开的整数: $T, C, T\_s, T\_e$;

第 $2$ 到第 $C+1$ 行: 第 $i+1$ 行描述第 $i$ 条道路，包含 $3$ 个由空格隔开的整数: $R\_s, R\_e, C\_i$。

### 输出格式

一个单独的整数表示从 $T\_s$ 到 $T\_e$ 的最小总费用。

数据保证至少存在一条道路。

### 数据范围

$1 \\le T \\le 2500$,

$1 \\le C \\le 6200$,

$1 \\le T\_s,T\_e,R\_s,R\_e \\le T$,

$1 \\le C\_i \\le 1000$

### 输入样例：

```
7 11 5 4
2 4 2
1 4 3
7 2 2
3 4 3
5 7 5
7 3 3
6 1 1
6 3 4
2 4 3
5 6 3
7 2 1
```

### 输出样例：

```
7
```

### 题解

