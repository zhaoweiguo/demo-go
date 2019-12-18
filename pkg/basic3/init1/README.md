
1. 一个包中最好只出现一个init方法
2. 不同package的init函数, 打印的顺序:
    先执行import里 的包内部的init 再执行本身init，最后执行main包的init最后执行main函数
3. 同一package内执行顺序:
    同文件中的 init 确实是按照声明顺序（从上到下）执行的
    不同文件中的 init 的执行顺序则由编译器先收到先执行这个规则来定
        规范并没有定义应该先送哪个文件，仅仅是鼓励按文件名称词法序（即字母序）

spec定义说明:
```
A package with no imports is initialized
    by assigning initial values to
    all its package-level variables
    followed by calling all init functions in the order they appear in the source,
     possibly in multiple files,
     as presented to the compiler
```