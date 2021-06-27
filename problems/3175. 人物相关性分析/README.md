## [3175. 人物相关性分析](https://www.acwing.com/problem/content/3178/)

### 题目

小明正在分析一本小说中的人物相关性。

他想知道在小说中 `Alice` 和 `Bob` 有多少次同时出现。

更准确的说，小明定义 `Alice` 和 `Bob` “同时出现”的意思是：在小说文本中 `Alice` 和 `Bob` 之间不超过 *K* 个字符。

例如以下文本：

`This is a story about Alice and Bob. Alice wants to send a private message to Bob.`

假设 *K = 20*，则 `Alice` 和 `Bob` 同时出现了 *2* 次，分别是 `Alice and Bob` 和 `Bob. Alice`。

前者 `Alice` 和 `Bob` 之间有 *5* 个字符，后者有 *2* 个字符。

注意:

1. `Alice` 和 `Bob` 是大小写敏感的， `alice` 或 `bob` 等并不计算在内。
2. `Alice` 和 `Bob` 应为单独的单词，前后可以有标点符号和空格，但是不能有字母。例如 `Bobbi` 並不算出现了 `Bob`。

### 输入格式

第一行包含一个整数 *K*。

第二行包含一行字符串，只包含大小写字母、标点符号和空格。长度不超过 *1000000*。

### 输出格式

输出一个整数，表示 `Alice` 和 `Bob` 同时出现的次数。

### 数据范围

*1 ≤ K ≤ 1000000*

### 输入样例：

```
20
This is a story about Alice and Bob. Alice wants to send a private message to Bob.
```

### 输出样例：

```
2
```
