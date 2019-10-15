
## From

* https://github.com/DrmagicE/wire-examples
* https://studygolang.com/articles/22266

## Provider & Injector

provider和injector是wire的两个核心概念:

```
provider: a function that can produce a value. 
        These functions are ordinary Go code.
injector: a function that calls providers in dependency order. 
        With Wire, you write the injector's signature, 
        then Wire generates the function's body.
```
通过提供provider函数，让wire知道如何产生这些依赖对象。
wire根据我们定义的injector函数签名，生成完整的injector函数，
injector函数是最终我们需要的函数，它将按依赖顺序调用provider。

在quickstart的例子中，NewMessage,NewGreeter,NewEvent都是provider，
wire_gen.go中的InitializeEvent函数是injector，
可以看到injector通过按依赖顺序调用provider来生成我们需要的对象Event。

## 接口绑定

根据依赖倒置原则（Dependence Inversion Principle），
对象应当依赖于接口，而不是直接依赖于具体实现。

抽象成接口依赖更有助于单元测试！

UserService依赖UserRepository接口，
其中mockUserRepo是UserRepository的一个实现，
由于在Go的最佳实践中，更推荐返回具体实现而不是接口。
所以mockUserRepo的provider函数返回的是*mockUserRepo这一具体类型。

wire无法自动将具体实现与接口进行关联，我们需要显示声明它们之间的关联关系。
通过wire.NewSet和wire.Bind将*mockUserRepo与UserRepository进行绑定

## 返回错误

我们的provider函数均只有一个返回值，
但在某些情况下，provider函数可能会对入参做校验，
如果参数错误，则需要返回error。
wire也考虑了这种情况，provider函数可以将返回值的第二个参数设置成error

## Cleanup functions

跟返回错误类似，将provider的第二个返回参数设置成func()用于返回cleanup function

wire对provider的返回值个数和顺序有所规定：

第一个参数是需要生成的依赖对象
如果返回2个返回值，第二个参数必须是func()或者error
如果返回3个返回值，第二个参数必须是func()，第三个参数则必须是error


## Provider set

当一些provider通常是一起使用的时候，可以使用provider set将它们组织起来

如:
wire.Build(NewEvent, NewGreeter, NewMessage)
改成:
wire.Build(EventSet) 


## 结构体provider

除了函数外，结构体也可以充当provider的角色，类似于setter注入


var Set = wire.NewSet(
	ProvideFoo,
	ProvideBar,
	wire.Struct(new(FooBar), "MyFoo", "MyBar"),
)
通过wire.Struct来指定那些字段要被注入到结构体中，如果是全部字段，也可以简写成：

var Set = wire.NewSet(
	ProvideFoo,
	ProvideBar,
	wire.Struct(new(FooBar), "*"), // * 表示注入全部字段
)


## 区分类型

由于injector的函数中，不允许出现重复的参数类型，否则wire将无法区分这些相同的参数类型
其中基础类型和通用接口类型是最容易发生冲突的类型，
    如果它们在provider函数中出现，最好统一新建一个别名来代替它(尽管还未发生冲突)

## Options Structs

如果一个provider方法包含了许多依赖，可以将这些依赖放在一个options结构体中，从而避免构造函数的参数太多：




