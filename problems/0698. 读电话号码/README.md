## [698. 读电话号码](https://www.acwing.com/problem/content/700/)

### 题目

你知道怎么用英文阅读电话号码吗？

现在让我告诉你。

例如，在中国，电话号码是 *11* 位数字，如： `15012233444`。

有人将号码划分为 *3-4-4* 格式，即 `150 1223 3444`，有人将号码划分为 *3-3-5* 格式，即 `150 122 33444`，不同的格式导致阅读这些数字的方式不同：

`150 1223 3444` 读作 `one five zero one double two three three triple four`。

`150 122 33444` 读作 `one five zero one double two double three triple four`。

问题出现了：

给定电话号码列表和号码划分格式，输出这些号码的正确英文读法。

规则：

- 单个数字单独读即可
- 两个连续数字使用 `double`
- 三个连续数字使用 `triple`
- 四个连续数字使用 `quadruple`
- 五个连续数字使用 `quintuple`
- 六个连续数字使用 `sextuple`
- 七个连续数字使用 `septuple`
- 八个连续数字使用 `octuple`
- 九个连续数字使用 `nonuple`
- 十个连续数字使用 `decuple`
- 超过十个连续数字则分别读每个数字

### 输入格式

第一行包含整数 *T*，表示共有 *T* 组测试数据。

每组数据占一行，首先包含一个电话号码 *N*，然后包含一个划分格式 *F*，*F* 为一个整数或多个用 `-` 连接起来的整数，不存在前导零，并且 *F* 中整数的总和等于 *N* 的长度。

### 输出格式

每组数据输出一个结果，每个结果占一行。

结果表示为 `Case #x: y`，其中 *x* 是组别编号（从 *1* 开始），*y* 是电话号码的正确读法。

### 数据范围

*1 ≤ T ≤ 100*,

电话号码长度不超过 *100*。

### 输入样例：

```
3
15012233444 3-4-4
15012233444 3-3-5
12223 2-3
```

### 输出样例：

```
Case #1:  one five zero one double two three three triple four
Case #2:  one five zero one double two double three triple four
Case #3:  one two double two three
```
