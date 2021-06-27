## [2630. 警告](https://www.acwing.com/problem/content/2632/)

### 题目

有一个简单的脚本语言，只有赋值语句、条件语句和返回语句三种。

变量名必须是单个大写字母，且变量都是 *32* 位带符号整数。

该语言的每条语句必须单独占一行。

程序中不含空行，且每行的行首行末均无空格。

每行的不同 token 之间用单个空格隔开。

该语言的 BNF 如下：

```
<line>          :: <head> | <assignment> | <if> | ELSE | END IF | <return>
<head>          :: PARAM <paramlist> | PARAM
<assignment>    :: <variable> = <rvalue>
<if>            :: IF <variable> <relation> <value> THEN
<return>        :: RETURN <value>
<paramlist>     :: <variable> | <variable> <paramlist>
<rvalue>        :: <value> | <value> <operator> <value>
<value>         :: <variable> | <integer>
<operator>      :: + | - | * | /
<relation>      :: < | = | >
<variable>      :: A | B | ... | Z
<integer>       :: 不含前导 0 的 32 位带符号整数
```

程序的第一行是一条 `<head>` 语句，定义了函数的参数，而最后一行一定是 `<return>` 语句。

`<head>` 语句不能在除了第一行之外的其他任何地方出现，但 `<return>` 语句可以在程序中多次出现。

行号从 *1* 开始编号。

每条 `IF` 语句一定有一个配套的 `END IF` 语句，还有一个可选的 `ELSE` 语句（注意没有 `ELSE IF` 语句）。

`IF` 语句可以嵌套，它总是比较一个变量和一个整数或者另一个变量。

你应该分析一个给定的程序，并输出两类警告信息（格式见样例输出）：

- 第一类警告：无法到达的代码行。不管各条 `IF` 语句的布尔表达式是真还是假（假设每条 `IF` 语句的布尔表达式都是既可能为真也可能为假，不受其他 `IF` 语句结果影响）。
- 第二类警告：可能未初始化的变量。该语句用到了某个变量的值，但这个变量既不在第一行指定的参数列表里，也没有在此之前由赋值语句赋值过。如果这条语句无法到达，那么不应该给出这类警告。

注意，语句 `ELSE` 和 `END IF` 不是可执行语句，因此不应该收到任何警告信息。

### 输入格式

最多包含 *50* 行，即你要处理的程序。

保证该程序合法。

### 输出格式

警告按照行号从小到大排序。

如果同一行内有多个可能未初始化的变量，按照字母顺序从小到大排列。

如果没有任何警告信息，你的输出应该为空。

### 输入样例1：

```
PARAM A B
IF A > 5 THEN
C = B * A
END IF
D = B - C
Z = Y + X
E = T
F = E + E
V = G + G
RETURN F
```

### 输出样例1：

```
Line 5: variable C might not have been initialized
Line 6: variable X might not have been initialized
Line 6: variable Y might not have been initialized
Line 7: variable T might not have been initialized
Line 9: variable G might not have been initialized
```

### 输入样例2：

```
PARAM G
RETURN G
B = K
RETURN C
```

### 输出样例2：

```
Line 3: unreachable code
Line 4: unreachable code
```

### 输入样例3：

```
PARAM T C
B = T
A = 4
IF A < 4 THEN
IF B > 3 THEN
Q = 100 + F
ELSE
IF C = -1111111111 THEN
Q = T - A
IF Q = 0 THEN
V = V - 1
END IF
ELSE
RETURN I
E = A
END IF
END IF
ELSE
Q = 1
END IF
RETURN Q
```

### 输出样例3：

```
Line 6: variable F might not have been initialized
Line 11: variable V might not have been initialized
Line 14: variable I might not have been initialized
Line 15: unreachable code
```
