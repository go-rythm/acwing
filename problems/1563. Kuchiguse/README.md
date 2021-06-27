## [1563. Kuchiguse](https://www.acwing.com/problem/content/1565/)

### 题目

日语以其句尾助词而臭名昭著。

这种助词的个人偏爱可以被认为是说话人个性的反映。

这种偏好被称为 “Kuchiguse”，在动画和漫画中经常被艺术地夸大。

例如，“nyan〜”（喵~） 通常用于具有猫式个性的角色：

- Itai nyan~（好痛啊，喵~)
- Ninjin wa iyada nyan~（人家不喜欢胡萝卜，喵~）

现在给定同一个人物说的几句话，你能找到她的 Kuchiguse 吗？

### 输入格式

第一行包含整数 *N*，表示共有 *N* 句话。

接下来 *N* 行，每行包含一个长度范围在 [0,256] 的字符串，表示一句话，注意 **区分大小写**。

### 输出格式

输出找到的 kuchiguse，即所有 *N* 行字符串的最长公共后缀。

如果不存在公共后缀，则输出 `nai`。

### 数据范围

*2 ≤ N ≤ 100*

### 输入样例1：

```
3
Itai nyan~
Ninjin wa iyadanyan~
uhhh nyan~
```

### 输出样例1：

```
nyan~
```

### 输入样例2：

```
3
Itai!
Ninjinnwaiyada T_T
T_T
```

### 输出样例2：

```
nai
```
