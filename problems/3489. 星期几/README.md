## [3489. 星期几](https://www.acwing.com/problem/content/3492/)

### 题目

已知 *1* 年 *1* 月 *1* 日是星期一。

现在给定一个日期，请你判断是星期几。

注意闰年的 *2* 月有 *29* 天。

满足下面条件之一的是闰年：

1. 年份是 *4* 的整数倍，而且不是 *100* 的整数倍；
2. 年份是 *400* 的整数倍。

### 输入格式

输入包含多组测试数据。

每组数据占一行，包含一个整数 *d* 表示日，一个字符串 *m* 表示月，一个整数 *y* 表示年。

月份 *1 ~ 12*，依次如下所示：

```
January, February, March, April, May, June, July, August, September, October, November, December
```

### 输出格式

每组数据输出一行结果，输出一个字符串表示给定日期是星期几。

周一至周日依次如下所示：

```
Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
```

### 数据范围

*1000 ≤ y ≤ 3000*,

给定日期保证合法。

每个输入最多包含 *100* 组数据。

### 输入样例：

```
9 October 2001
14 October 2001
```

### 输出样例：

```
Tuesday
Sunday
```
