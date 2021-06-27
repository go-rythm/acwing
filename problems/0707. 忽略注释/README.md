## [707. 忽略注释](https://www.acwing.com/problem/content/709/)

### 题目

优秀的程序员能写出优秀的注释。

伊戈尔是一名程序员，他喜欢 *C* 语言风格形式的注释 `/* ... */`。

对他而言，如果他能够将这种风格用作所有编程语言甚至文档的统一注释格式，那将是理想的，例如用于 Python，Haskell 或 HTML / XML 文档等。

实现这种想法对于他而言并不算难。

他需要的是一个注释预处理器，用于删除所有的 `/* ... */` 内容。

然后，处理后的文本再移交给它所属的编译器 / 文档渲染器作进一步处理。

不过，伊戈尔的预处理器并没有那么简单。

以下是预处理器的一些强大的特点：

1、预处理器读取的注释可以嵌套，就像在大多数编程语言中的括号嵌套一样。可以在注释中添加注释。例如，以下代码块具有应由注释预处理器删除的外部注释，其中还包含两个内部注释。

```
printf("Hello /* a comment /* a comment inside comment */
        inside /* another comment inside comment */
        string */ world");
```

在预处理步骤之后，它变为：

```
printf("Hello  world");
```

2、注释可以出现在文本的任何地方，包括字符串中 `"/*...*/"`，常数中 `12/*...*/34`，甚至是字符转义符中 `/*...*/n`。

不管在任何地方出现的这些注释，都会被预处理器删除掉。

注意：

1、处理器只进行一轮删除工作。由于删除任何注释块而产生的新注释块不会被删除掉。例如：

```
//*no recursion*/* file header */
```

应该生成：

```
/* file header */
```

任何 `/*` 或 `*/` 中的 `*` 字符都不能在另一个 `/*` 或 `*/` 中重复使用。

例如，以下内容不构成正确的注释块

```
/*/
```

### 输入格式

输入文件是一个包含 `/* ... */` 注释块的文本文档。

输入文件保证有效。

它遵循问题陈述中的文本规范。

输入文件始终以换行符号终止。

### 输出格式

本题只包含一个测试用例。

首先，我们需要输出以下行：

```
Case #1:
```

然后，按照问题陈述中指定的方式输出文档，并删除所有注释。

请勿删除注释外的任何空格或空行。

### 数据范围

输入文本中只包含：

- 字母： `a-z, A-Z`
- 数字： `0-9`
- 标点符号： `~ ! @ # % ^ & * ( ) - + = : ; " ' < > , . ? | /  { } [ ] _`
- 空白字符：换行、空格

文本大小不超过 *100 kb*

### 输入样例：

```
//*no recursion*/* file header
***********/************
* Sample input program *
**********/*************
*/
int spawn_workers(int worker_count) {
/* The block below is supposed to spawn 100 workers.
     But it creates many more.
     Commented until I figure out why.
for (int i = 0; i < worker_count; ++i) {
    if(!fork()) {
      /* This is the worker. Start working. */
      do_work();
    }
}
*/
return 0; /* successfully spawned 100 workers */
}

int main() {
printf("Hello /*a comment inside string*/ world");
int worker_count = 0/*octal number*/144;
if (spawn_workers(worker_count) != 0) {
    exit(-1);
}
return 0;
}
```

### 输出样例：

```
Case #1:
/* file header
************************
*/
int spawn_workers(int worker_count) {

return 0;
}

int main() {
printf("Hello  world");
int worker_count = 0144;
if (spawn_workers(worker_count) != 0) {
    exit(-1);
}
return 0;
}
```
