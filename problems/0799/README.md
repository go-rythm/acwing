## [799\. 最长连续不重复子序列](https://www.acwing.com/problem/content/801/)

### 题目

给定一个长度为 $n$ 的整数序列，请找出最长的不包含重复的数的连续区间，输出它的长度。

### 输入格式

第一行包含整数 $n$。

第二行包含 $n$ 个整数（均在 $0 \\sim 10^5$ 范围内），表示整数序列。

### 输出格式

共一行，包含一个整数，表示最长的不包含重复的数的连续区间的长度。

### 数据范围

$1 \\le n \\le 10^5$

### 输入样例：

```
5
1 2 2 3 5
```

### 输出样例：

```
3
```

### 题解

前置题目：0798

前置知识：语法

本题知识：基础算法-双指针

#### 题目分析

有点类似桶排序的方法，用一个数组记录每个数字出现的次数，这样就可以判断出是否重复

```c
// 每次 j 的位置都是向左最远能到达的满足条件（不重复）的位置
1  2  2  3  5  最大值1
i
j

1  2  2  3  5  最大值2
      i
j

1  2  2  3  5
         i
      j

1  2  2  3  5  最大值 (1 < 2) 还是2
      i
      j 

1  2  2  3  5  最大值 (2 < 3) 是3
            i
	  j
```

#### 模板

```go
// 两根指针常用解题代码
for i, j := 0, 0; i < n; i++ {
	for j < i && check(i, j) {
		j++
	}
}
```

