## [1598. 求平均值](https://www.acwing.com/problem/content/1600/)

### 题目

基本任务非常简单：给定 *N* 个实数，请你计算它们的平均值。

但是令事情变得复杂的是，某些输入数字可能不合法。

合法输入数字指 *[−1000,1000]* 范围内的精确到 **不超过** *2* 个小数位的实数。

在计算平均值时，不得将这些非法数字计算在内。

### 输入格式

第一行包含整数 *N*。

第二行包含 *N* 个输入数字，数字之间用空格隔开。

### 输出格式

按照输入的顺序，对于每个输入的非法数字，在一行中输出 `ERROR: X is not a legal number`，其中 `X` 是非法输入。

最后一行，输出 `The average of K numbers is Y`，其中 `K` 是合法输入的数量， `Y` 是它们的平均值，注意保留两位小数。

如果平均值无法计算，则将 `Y` 替换为 `Undefined`。

如果 `K` 仅为 *1*，则输出 `The average of 1 number is Y`。

### 数据范围

*1 ≤ N ≤ 1000*

### 输入样例1：

```
7
5 -3.2 aaa 9999 2.3.4 7.123 2.35
```

### 输出样例1：

```
ERROR: aaa is not a legal number
ERROR: 9999 is not a legal number
ERROR: 2.3.4 is not a legal number
ERROR: 7.123 is not a legal number
The average of 3 numbers is 1.38
```

### 输入样例2：

```
2
aaa -9999
```

### 输出样例2：

```
ERROR: aaa is not a legal number
ERROR: -9999 is not a legal number
The average of 0 numbers is Undefined
```
