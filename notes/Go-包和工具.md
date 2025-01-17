# Go - 包和工具



## 包

在Go语言中，包的作用类似C#中的命名空间。

一个代码包中可以包含任意个以.go 为扩展名的源码文件。

代码包的名称一般会与源码文件所在的目录同名。如果不同名，那么在构建、安装的过程中会以代码包名称为准。在文件系统中，这些代码包其实是与目录一一对应的。由于目录可以有子目录，所以代码包也可以有子包。

- 同一个目录里的 Go 代码的 package 要保持一致。
- 代码的 package 可以和所在的目录不一致

### 导入其他包

在Go程序里，每一个包通过称为`导入路径`的唯一字符串来标识。

文件所在的完整目录名的尾部（$GOPATH/src/之后的）就是包导入的路径。

```go
package main

import (
	"chapter2/tempconv"
	"fmt"
)
```

一个导入路径标注一个目录，目录中包含构成包的一个或多个Go源文件。

除了导入路径之外，每个包还有一个包名，它以短名字的形式（且不必是唯一的）出现在包的声明中。

按约定，包名匹配导入路径的最后一段，例如“chapter2/tempconv”的包名是tempconv。这种由导入声明给导入的包绑定的短名字，可以用来在整个文件中引用包的内容。为了避免冲突，导入声明还可以设定一个可选的名字。

如果导入一个没有被引用的包，将会触发一个编译错误。这个检查有助于消除代码演进过程中不再需要的依赖。

每个代码包都会有导入路径。代码包的导入路径是其他代码在使用该包中的程序实体时，需要引入的路径。在实际使用程序实体之前，我们必须先导入其所在的代码包。

```go
import "github.com/labstack/echo"
```

在工作区中，一个代码包的导入路径实际上就是从 src 子目录，到该包的实际存储位置的相对路径。

一般情况下，Go 语言的源码文件都需要被存放在环境变量 GOPATH 包含的某个工作区（目录）中的 src 目录下的某个代码包（目录）中。

### 包的初始化

包的初始化按照如下过程进行：

1. 首先初始化包级别的变量，这些变量按照声明顺序初始化，在依赖已解析完毕的情况下，根据依赖的顺序进行。例如：

   ```go
   var a = b + c //3：最后把a初始化为3
   var b = f()   //2：接着通过调用f()将b初始化为2
   var c = 1     //1：首先初始化c为1
   func f() int  { return c + 1 }
   ```

2. 如果包由多个.go文件组成，初始化按照编译器收到文件的顺序进行：go工具会在调用编译器前将.go文件进行排序。

   对于包级别的每一个变量，生命周期从其值被初始化开始，但是对于其他一些变量，比如数据表，初始化表达式不是简单地设置它的初始化值，而是需要调用init()函数（见下述）。

3. 包的初始化按照在程序中导入的顺序来进行，依赖顺序优先，每次初始化一个包。例如，如果包p导入了包q，可以确保q在p之前已完全初始化。初始化过程是自下向上的，main包最后初始化。在这种方式下，在程序的main函数开始执行前，所有的包已初始化完毕。

#### init()

- 在 main 被执行前，所有依赖的 package 的 init 方法都会被执行；
- 不同包的 init 函数按照包导入的依赖关系决定执行顺序
- 每个包都可以有多个 init 函数
- 包的每个源文件也可以有多个 init 函数，这点比较特殊。

任何文件都可以包含任意数量的init函数，它的声明形式如下：

```go
func init(){...}
```

init函数不能被手动的调用和被引用，另一方面，它也是普通的函数。在每一个文件里，当程序启动的时候，init函数按照它们**声明的顺序**自动执行。



### 补充

同一个目录下，只允许有一个包，不允许有多个不同名的包声明。

每一个含有main函数的文件，必须单独放在一个目录里。同一个目录下的多个go文件中，只允许出现一个main方法。如果多个文件有多个main方法，必须使用目录给隔开，这些文件可以都使用同一个包名是允许的。可参考官方的tools。



## 工具

### go run

用于检测数据访问冲突问题，可以使用如下命令：

```shell
go run -race wy.go
```



### go tool pprof

查看性能成本，CPU使用率等