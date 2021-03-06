## [358\. 岛屿](https://www.acwing.com/problem/content/360/)

### 题目

你准备游览一个公园，该公园由 $N$ 个岛屿组成，当地管理部门从每个岛屿出发向另外一个岛屿建了一座桥，不过桥是可以双向行走的。

同时，每对岛屿之间都有一艘专用的往来两岛之间的渡船。

相对于乘船而言，你更喜欢步行。

你希望所经过的桥的总长度尽可能的长，但受到以下的限制：

1. 可以自行挑选一个岛开始游览。
2. 任何一个岛都不能游览一次以上。
3. 无论任何时间你都可以由你现在所在的岛 $S$ 去另一个你从未到过的岛 $D$。由 $S$ 到 $D$ 可以有以下方法：


    （1）步行：仅当两个岛之间有一座桥时才有可能。对于这种情况，桥的长度会累加到你步行的总距离中。


    （2）渡船：你可以选择这种方法，仅当没有任何桥和以前使用过的渡船的组合可以由 $S$ 走到 $D$（当检查是否可到达时，你应该考虑所有的路径，包括经过你曾游览过的那些岛）。

注意，你不必游览所有的岛，也可能无法走完所有的桥。

请你编写一个程序，给定 $N$ 座桥以及它们的长度，按照上述的规则，计算你可以走过的桥的最大长度。

### 输入格式

第 $1$ 行包含整数 $N$。

第 $2.. N+1$ 行，每行包含两个整数 $a$ 和 $L$，第 $i+1$ 行表示岛屿 $i$ 上建了一座通向岛屿 $a$ 的桥，桥的长度为 $L$。

### 输出格式

输出一个整数，表示结果。

对某些测试，答案可能无法放进 $32-bit$ 整数。

### 数据范围

$2 \\le N \\le 10^6$,

$1 \\le L \\le 10^8$

### 输入样例：

```
7
3 8
7 2
4 2
1 4
1 9
3 4
2 3
```

### 输出样例：

```
24
```

### 题解

