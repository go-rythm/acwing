## [1419. 牛的密码学](https://www.acwing.com/problem/content/1421/)

### 题目

农夫布朗和农夫约翰的奶牛们正在计划着一起从农场中逃跑，为了保证暗中联络的信息不被他人获知，它们还设计了一种消息的加密方法。

具体来说，当一头奶牛收到了一个消息比如 `International Olympiad in Informatics`，则可以通过在消息中的任意位置插入字母 *C,O,W* 来对其进行更改。

插入时，应使得 *C* 的位置在 *O* 之前，*O* 的位置在 *W* 之前，并且 *C* 和 *O* 之间的内容还要与 *O* 和 *W* 之间的内容进行互换，从而得到加密后的消息。

下面给出两个具体加密的例子：

```
International Olympiad in Informatics --> CnOIWternational Olympiad in Informatics

International Olympiad in Informatics --> International Cin InformaticsOOlympiad W
```

甚至，为了让消息所代表的含义更加难以被人破解，奶牛们还会采取多重加密的方式，对用上述加密方法加密过的消息再次使用该方法进行加密，如此重复多次，从而让消息变得极其难以捉摸。

一天晚上，农夫约翰的奶牛收到了这种经过多重加密的消息。

编写一个程序来计算未加密的原始消息是否可能是字符串：

```
Begin the Escape execution at the Break of Dawn
```

### 输入格式

共一行，包含一个字符串，表示加密后的消息。

### 输出格式

输出两个整数，第一个整数表示是否能够得到题目所述的原始消息，如果能得到则输出 *1*，否则输出 *0*。

如果第一个整数为 *1*，则第二个整数表示加密次数，如果第一个整数为 *0*，则第二个整数直接输出 *0*。

### 数据范围

输入字符串长度不超过 *75*。

### 输入样例：

```
Begin the EscCution at the BreOape execWak of Dawn
```

### 输出样例：

```
1 1
```
