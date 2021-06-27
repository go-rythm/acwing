## [1949. 农夫约翰没有大棕牛](https://www.acwing.com/problem/content/1951/)

### 题目

农夫约翰喜欢收集尽可能多的不同类型的奶牛。

实际上，他收集了几乎所有可能的奶牛类型，并将他未收集到的 *N* 种奶牛记录在了一个列表中。

列表格式如下：

```
Farmer John has no large brown noisy cow.
Farmer John has no small white silent cow.
Farmer John has no large spotted noisy cow.
```

列表中的每一行都使用一系列的形容词来描述一种未收集到的奶牛类型。

描述每种奶牛所使用的形容词数量均相同（在本例中形容词数量为 *3*）。

每行的形容词数量将在 *2..30* 范围内。

对于每个未在列表中出现的可能的形容词组合描述出的奶牛类型，约翰都有收集。

在本例中，第一个形容词可以是 `large` 或 `small`，第二个形容词可以是 `brown`、 `white` 或 `spotted`，第三个形容词可以是 `noisy` 或 `silent`。

这就给出了 *2 × 3 × 2 = 12* 种不同的形容词组合，对应 *12* 种不同类型的奶牛。

除了在列表中提及的奶牛以外，其他的每种类型的奶牛约翰都有收集。

在本例中，除去他没有的 *3* 种奶牛，他共收集了 *9* 种奶牛，例如， `large white noisy` 的奶牛就是他收集到的奶牛类型之一。

约翰保证他最多有不超过 *10^9* 种奶牛。

如果约翰按照字典序列出所有他拥有的奶牛类型，那么请你确定排在第 *K* 个的奶牛类型是什么。

### 输入格式

第一行包含两个整数 *N* 和 *K*。

接下来 *N* 行，每行都包含一个句子形如 `Farmer John has no large spotted noisy cow.`。

句子中的每个形容词由小写字母构成的，长度不超过 *10* 的字符串。

每个句子都以 `cow.` 作为结尾。

### 输出格式

输出描述约翰拥有的第 *K* 种奶牛类型的形容词组合。

### 数据范围

*1 ≤ N ≤ 100*,

*K* 保证合法。

### 输入样例：

```
3 7
Farmer John has no large brown noisy cow.
Farmer John has no small white silent cow.
Farmer John has no large spotted noisy cow.
```

### 输出样例：

```
small spotted noisy
```

### 样例解释

约翰拥有的奶牛类型按字典序排列如下：

```
large brown silent
large spotted silent
large white noisy
large white silent
small brown noisy
small brown silent
small spotted noisy
small spotted silent
small white noisy
```

第 *7* 种类型为 `small spotted noisy`。
