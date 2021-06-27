## [1499. 数字图书馆](https://www.acwing.com/problem/content/1501/)

### 题目

数字图书馆包含数以百万计的书籍。

每本书的书名，作者，摘要关键词，出版商和出版年限这五类关键信息都在数据库中有所记录。

每本书被分配一个唯一的 *7* 位数字作为其 `ID`。

当读者查询某一关键信息时，你应该找到所有与查询相关的书籍，并将它们按 `ID` 的升序排序输出。

### 输入格式

第一行包含整数 *N*，表示共有 *N* 本书。

接下来包含这 *N* 本书的具体信息，每本书的相关信息占 *6* 行：

- 第一行：书的 `ID`，一个 *7* 位数字。
- 第二行：书名，一个长度不超过 *80* 的字符串。
- 第三行：作者，一个长度不超过 *80* 的字符串。
- 第四行：关键词，每个关键词的长度均不超过 *10*，且关键词中不包含空格，关键词之间用空格隔开。
- 第五行：出版商，一个长度不超过 *80* 的字符串。
- 第六行：出版年限，一个在 *[1000,3000]* 范围内的 *4* 位数字。

一本书，只有一位作者，包含的关键词不超过 *5* 个。

总共不超过 *1000* 个不同的关键词，不超过 *1000* 个不同的出版商。

图书信息介绍完毕后，有一行包含一个整数 *M*，表示查询次数。

接下来 *M* 行，每行包含一个查询，具体格式如下：

- `1: a book title`，查询书名。
- `2: name of an author`，查询作者名。
- `3: a key word`，查询关键词。
- `4: name of a publisher`，查询出版商。
- `5: a 4-digit number representing the year`，查询出版年限。注意，这个年限可能包含前导 *0*。

### 输出格式

对于每个查询，首先将查询信息输出在一行中。

接下来若干行，每行输出一个查询到的相关书籍的 `ID`，按升序顺序排列。

如果查询不到相关书籍，则输出 `Not Found`。

### 数据范围

*1 ≤ N ≤ 10^4*,

*1 ≤ M ≤ 1000*

### 输入样例：

```
3
1111111
The Testing Book
Yue Chen
test code debug sort keywords
ZUCS Print
2011
3333333
Another Testing Book
Yue Chen
test code sort keywords
ZUCS Print2
2012
2222222
The Testing Book
CYLL
keywords debug book
ZUCS Print2
2011
6
1: The Testing Book
2: Yue Chen
3: keywords
4: ZUCS Print
5: 2011
3: blablabla
```

### 输出样例：

```
1: The Testing Book
1111111
2222222
2: Yue Chen
1111111
3333333
3: keywords
1111111
2222222
3333333
4: ZUCS Print
1111111
5: 2011
1111111
2222222
3: blablabla
Not Found
```
