## [1185\. 单词游戏](https://www.acwing.com/problem/content/1187/)

### 题目

有 $N$ 个盘子，每个盘子上写着一个仅由小写字母组成的英文单词。

你需要给这些盘子安排一个合适的顺序，使得相邻两个盘子中，前一个盘子上单词的末字母等于后一个盘子上单词的首字母。

请你编写一个程序，判断是否能达到这一要求。

### 输入格式

第一行包含整数 $T$，表示共有 $T$ 组测试数据。

每组数据第一行包含整数 $N$，表示盘子数量。

接下来 $N$ 行，每行包含一个小写字母字符串，表示一个盘子上的单词。

一个单词可能出现多次。

### 输出格式

如果存在合法解，则输出”Ordering is possible.”，否则输出”The door cannot be opened.”。

### 数据范围

$1 \\le N \\le 10^5$,

单词长度均不超过1000

### 输入样例：

```
3
2
acm
ibm
3
acm
malform
mouse
2
ok
ok
```

### 输出样例：

```
The door cannot be opened.
Ordering is possible.
The door cannot be opened.
```

### 题解

