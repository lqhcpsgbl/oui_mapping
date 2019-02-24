# Go defer 笔记

defer用于资源的释放，会在函数返回之前进行调用。比如，官网实例：

```golang
package main

import "fmt"

func main() {
    fmt.Println("hello")
    defer fmt.Println("world")
}

```

描述如下：

> A defer statement defers the execution of a function until the surrounding function returns.
The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

defer 的执行按照 后进先出 的顺序

> It is very important to remember that deferred functions are executed in Last In First Out
(LIFO) order after the return of the surrounding function. 

实例代码：

```golang

package main

import (
    "fmt"
)

func test() {
    for i := 1; i <= 3; i++ {
        fmt.Println(i)
        defer fmt.Println(i)
    }
    fmt.Println("first execute here!")
}

func main() {
    test()
}


```

执行结果如下：

```bash
λ go run test.go
1
2
3
first execute here!
3
2
1
```

结果 test() 中 for 循环 中的 defer 按照 后进先出 执行。

```golang

package main

import (
    "fmt"
)

func test() {
    i := 3
    for i = 3; i > 0; i-- {
        // 会先输出 3 2 1
        fmt.Print(i, " ")
        // deferred anonymous function is evaluated after the for loop ends
        // because it has no parameters
        // 这里的i作用域是整个函数
        defer func() {
            fmt.Print("in loop: ", i, " ")
        }()
    }
    i = 3
    defer fmt.Print("tttt", i, "ffff\n") // 此处的i 作用域是 i = 3
    i = 100
    fmt.Println("first execute here? No!")
    fmt.Println("i is", i)
}

func main() {
    test()
}


```

执行结果如下：

```bash

λ go run test.go
3 2 1 first execute here? No!
i is 100
tttt3ffff
in loop: 100 in loop: 100 in loop: 100 %
```

```golang

package main

import (
    "fmt"
)

func test() {
    i := 3
    for i = 3; i > 0; i-- {
        // 会先输出 3 2 1
        fmt.Print(i, " ")
        defer func(n int) {
            fmt.Print("in loop: ", n, " ")
        }(i)
    }
    i = 3
    defer fmt.Print("tttt", i, "ffff\n") // 此处的i 作用域是 i = 3
    i = 100
    fmt.Println("first execute here? No!")
    fmt.Println("i is", i)
}

func main() {
    test()
}

```

运行结果：

```bash
λ go run test.go
3 2 1 first execute here? No!
i is 100
tttt3ffff
in loop: 1 in loop: 2 in loop: 3 %
```

`defer` 关键字只是注册回调，并不保存栈中的变量，而注册时候的值就是最后打印出来的值。
