## [1543. 栈](https://www.acwing.com/problem/content/1545/)

### 题目

栈是最基本的数据结构之一，它基于后进先出（LIFO）的原理。

基本操作包括“入栈”（将元素插入到顶部位置）和“出栈”（删除顶部元素）。

现在，你需要实现一个栈，该栈要具有一个额外的操作：PeekMedian-返回栈中所有元素的中值。

对于 *N* 个元素，如果 *N* 为偶数，则中值定义 **从小到大** 第 *frac{N}{2}* 个元素；如果 *N* 为奇数，则中值定义为 **从小到大** 第 *frac{N+1}{2}* 个元素。

### 输入格式

第一行包含整数 *N*，表示命令数。

接下来 *N* 行，每行包含以下三种命令中的一种：

```
Push key
Pop
PeekMedian
```

其中 `key` 是一个不超过 *10^5* 的正整数。

### 输出格式

对于每个 `Push` 操作，在顶部插入一个 `key` 值，不输出任何内容。

对于每个 `Pop` 或 `PeekMedian` 命令，在一行中输出相应的返回值。

如果命令无效，则输出 `Invalid`。

### 数据范围

*1 ≤ N ≤ 10^5*

### 输入样例：

```
17
Pop
PeekMedian
Push 3
PeekMedian
Push 2
PeekMedian
Push 1
PeekMedian
Pop
Pop
Push 5
Push 4
PeekMedian
Pop
Pop
Pop
Pop
```

### 输出样例：

```
Invalid
Invalid
3
2
2
1
2
4
4
5
3
Invalid
```
