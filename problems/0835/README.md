## [835\. Trie字符串统计](https://www.acwing.com/problem/content/837/)

### 题目

维护一个字符串集合，支持两种操作：

1. `I x` 向集合中插入一个字符串 $x$；
2. `Q x` 询问一个字符串在集合中出现了多少次。

共有 $N$ 个操作，输入的字符串总长度不超过 $10^5$，字符串仅包含小写英文字母。

### 输入格式

第一行包含整数 $N$，表示操作数。

接下来 $N$ 行，每行包含一个操作指令，指令为 `I x` 或 `Q x` 中的一种。

### 输出格式

对于每个询问指令 `Q x`，都要输出一个整数作为结果，表示 $x$ 在集合中出现的次数。

每个结果占一行。

### 数据范围

$1 \\le N \\le 2\*10^4$

### 输入样例：

```
5
I abc
Q abc
Q ab
I ab
Q ab
```

### 输出样例：

```
1
0
1
```

### 题解

前置题目：0831

前置知识：邻接表

本题知识：数据结构-Trie

#### 题目分析

Trie树，取自`retrieval`，是一种以空间换时间，高效查找字符串集合的数据结构

![trie](https://raw.githubusercontent.com/luxcgo/imgs4md/master/img/trie.gif)

可以通过邻接表存储整棵树，并在字符串结尾做一个标记表明是一条字符串。

下标为 0 的点，不存储值，只作为根节点
