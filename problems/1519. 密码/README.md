## [1519. 密码](https://www.acwing.com/problem/content/1521/)

### 题目

为了准备 PAT，系统不得不为用户生成随机密码。

但是有时一些数字和字母之间总是难以区分，比如 `1`（数字一）和 `l`（*L* 的小写）， `0`（数字零）和 `O`（*o* 的大写）。

一种解决办法是将 `1`（数字一）替换为 `@`，将 `0`（数字零）替换为 `%`，将 `l`（*L* 的小写）替换为 `L`，将 `O`（*o* 的大写）替换为 `o`。

现在，你的任务就是帮助系统检查这些用户的密码，并对难以区分的部分加以修改。

### 输入格式

第一行包含一个整数 *N*，表示用户数量。

接下来 *N* 行，每行包含一个用户名和一个密码，都是长度不超过 *10* 且不含空格的字符串。

### 输出格式

首先输出一个整数 *M*，表示已修改的用户密码数量。

接下来 *M* 行，每行输出一个用户名称和其修改后的密码。

用户的输出顺序和读入顺序必须相同。

如果没有用户的密码被修改，则输出 `There are N accounts and no account is modified`，其中 *N* 是用户总数。

如果 *N = 1*，则应该输出 `There is 1 account and no account is modified`。

### 数据范围

*1 ≤ N ≤ 1000*

### 输入样例1：

```
3
Team000002 Rlsp0dfa
Team000003 perfectpwd
Team000001 R1spOdfa
```

### 输出样例1：

```
2
Team000002 RLsp%dfa
Team000001 R@spodfa
```

### 输入样例2：

```
1
team110 abcdefg332
```

### 输出样例2：

```
There is 1 account and no account is modified
```

### 输入样例3：

```
2
team110 abcdefg222
team220 abcdefg333
```

### 输出样例3：

```
There are 2 accounts and no account is modified
```
