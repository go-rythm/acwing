## [1767. COWBASIC](https://www.acwing.com/problem/content/1769/)

### 题目

贝茜发明了一种新的编程语言，但是由于还没有编译器，她需要你的帮助才能真正运行她的程序。

COWBASIC 是一种简单，优雅的语言。

它具有两个关键功能：加法和 MOO 循环。

贝茜设计了一个聪明的溢出解决方案：所有加法都对 *10^9 + 7* 取模。

但是贝茜的真正成功之处在于 MOO 循环，该循环可以将一段代码运行固定次数。

MOO 循环和加法当然可以嵌套。

给定一个 COWBASIC 程序，请帮助贝茜确定程序返回的数字是多少。

### 输入格式

给定你的 COWBASIC 程序最长不超过 *100* 行，每行最多不超过 *350* 个字符。

COWBASIC 程序由语句列表（list of statements）构成。

共有三种类型的语句（statement）：

```
<variable> = <expression>

<literal> MOO {
<list of statements>
}

RETURN <variable>
```

共有三种类型的表示方式（expression）：

```
<literal>

<variable>

( <expression> ) + ( <expression> )
```

`literal` 是不超过 *100000* 的正整数。

`variable` 是长度不超过 *10* 的由小写英文字母构成的字符串。

保证所有变量（variable）在被定义之前，不会被使用或返回。

保证 RETURN 命令只会在最后一行出现一次。

### 输出格式

输出一个正整数，表示程序返回值。

### 输入样例1：

```
x = 1
10 MOO {
x = ( x ) + ( x )
}
RETURN x
```

### 输出样例1：

```
1024
```

### 样例1解释

这个 COWBASIC 程序计算了 *2^{10}*。

### 输入样例2：

```
n = 1
nsq = 1
100000 MOO {
100000 MOO {
    nsq = ( nsq ) + ( ( n ) + ( ( n ) + ( 1 ) ) )
    n = ( n ) + ( 1 )
}
}
RETURN nsq
```

### 输出样例2：

```
4761
```

### 样例2解释

这个 COWBASIC 程序计算了 *(10^5 × 10^5 + 1)^2 (bmod 10^9+7)*。
