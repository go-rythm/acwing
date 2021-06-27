## [1944. 记录保存](https://www.acwing.com/problem/content/1946/)

### 题目

农夫约翰一直在详细记录奶牛进入牛棚挤奶的情况。

每小时会有一组三头牛进入牛棚产奶，约翰会记录它们的名字。

例如，在 *5* 小时内，他可能会记录下如下列表，每行对应进入牛棚的一组牛：

```
BESSIE ELSIE MATILDA
FRAN BESSIE INGRID
BESSIE ELSIE MATILDA
MATILDA INGRID FRAN
ELSIE BESSIE MATILDA
```

约翰发现同一组奶牛可能会多次出现在他记录的名单中，在上面的例子中， `BESSIE ELSIE MATILDA` 这个组合出现了 *3* 次（尽管约翰不一定每次都按同样的顺序记录它们的名字），

请帮助约翰计算进入牛棚次数最多的一组牛的进入次数。

### 输入格式

第一行包含整数 *N*。

接下来 *N* 行，每行都包含一组三头牛的名字，每个名字都是一个长度在 *1 ~ 10* 之间的由大写字母构成的字符串。

### 输出格式

输出进入牛棚次数最多的一组牛的进入次数。

### 数据范围

*1 ≤ N ≤ 1000*

### 输入样例：

```
5
BESSIE ELSIE MATILDA
FRAN BESSIE INGRID
BESSIE ELSIE MATILDA
MATILDA INGRID FRAN
ELSIE BESSIE MATILDA
```

### 输出样例：

```
3
```
