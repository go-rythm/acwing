## [3575. 编排字符串](https://www.acwing.com/problem/content/3578/)

### 题目

请输入字符串，最多输入 *4* 个字符串，要求后输入的字符串排在前面，例如

```
输入：EricZ
输出：1=EricZ

输入：David
输出：1=David 2=EricZ

输入：Peter
输出：1=Peter 2=David 3=EricZ

输入：Alan
输出：1=Alan 2=Peter 3=David 4=EricZ

输入：Jane
输出：1=Jane 2=Alan 3=Peter 4=David
```

### 输入格式

第一行为字符串个数 *m*。

接下来 *m* 行每行一个字符串。

### 输出格式

输出 *m* 行，每行按照样例格式输出，注意用一个空格隔开。

### 数据范围

*m* 不超过 *100*，每个字符串长度不超过 *20*。

### 输入样例：

```
5
EricZ
David
Peter
Alan
Jane
```

### 输出样例：

```
1=EricZ
1=David 2=EricZ
1=Peter 2=David 3=EricZ
1=Alan 2=Peter 3=David 4=EricZ
1=Jane 2=Alan 3=Peter 4=David
```
