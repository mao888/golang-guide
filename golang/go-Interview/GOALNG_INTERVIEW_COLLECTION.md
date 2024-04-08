## 零、go与其他语言

### 0、什么是[面向对象](https://so.csdn.net/so/search?q=面向对象&spm=1001.2101.3001.7020)

在了解 Go 语言是不是面向对象（简称：OOP） 之前，我们必须先知道 OOP 是啥，得先给他 “下定义”

根据 Wikipedia 的定义，我们梳理出 OOP 的几个基本认知：

- 面向对象编程（OOP）是一种基于 “对象” 概念的编程范式，它可以包含数据和代码：数据以字段的形式存在（通常称为属性或属性），代码以程序的形式存在（通常称为方法）。
- 对象自己的程序可以访问并经常修改自己的数据字段。
- 对象经常被定义为类的一个实例。
- 对象利用属性和方法的私有/受保护/公共可见性，对象的内部状态受到保护，不受外界影响（被封装）。

基于这几个基本认知进行一步延伸出，面向对象的三大基本特性：

- 封装
- 继承
- 多态

### 1、Go语言和Java有什么区别? 

1、Go上不允许函数重载，必须具有方法和函数的唯一名称，而Java允许函数重载。

2、在速度方面，Go的速度要比Java快。

3、Java默认允许多态，而Go没有。

4、Go语言使用HTTP协议进行路由配置，而Java使用Akka.routing.ConsistentHashingRouter和Akka.routing.ScatterGatherFirstCompletedRouter进行路由配置。

5、Go代码可以自动扩展到多个核心，而Java并不总是具有足够的可扩展性。

6、Go语言的继承通过匿名组合完成，基类以Struct的方式定义，子类只需要把基类作为成员放在子类的定义中，支持多继承;而Java的继承通过extends关键字完成，不支持多继承。

### 2、Go 是面向对象的语言吗？

是的，也不是。原因是：

1. Go 有类型和方法，并且允许面向对象的编程风格，但没有类型层次。
2. Go 中的 "接口 "概念提供了一种不同的方法，我们认为这种方法易于使用，而且在某些方面更加通用。还有一些方法可以将类型嵌入到其他类型中，以提供类似的东西，但不等同于子类。
3. Go 中的方法比 C++ 或 Java 中的方法更通用：它们可以为任何类型的数据定义，甚至是内置类型，如普通的、"未装箱的 "整数。它们并不局限于结构（类）。
4. Go 由于缺乏类型层次，Go 中的 "对象 "比 C++ 或 Java 等语言更轻巧。

### 3、Go 实现面向对象编程

#### 封装

面向对象中的 “封装” 指的是可以隐藏对象的内部属性和实现细节，仅对外提供公开接口调用，这样子用户就不需要关注你内部是怎么实现的。

**在 Go 语言中的属性访问权限，通过首字母大小写来控制：**

- 首字母大写，代表是公共的、可被外部访问的。
- 首字母小写，代表是私有的，不可以被外部访问。

Go 语言的例子如下：

```go
type Animal struct {
	name string
}

func NewAnimal() *Animal {
 	return &Animal{}
}

func (p *Animal) SetName(name string) {
 	p.name = name
}

func (p *Animal) GetName() string {
 	return p.name
}
```



在上述例子中，我们声明了一个结构体 Animal，其属性 name 为小写。没法通过外部方法，在配套上存在 Setter 和 Getter 的方法，用于统一的访问和设置控制。

以此实现在 Go 语言中的基本封装。

#### 继承

面向对象中的 “继承” 指的是子类继承父类的特征和行为，使得子类对象（实例）具有父类的实例域和方法，或子类从父类继承方法，使得子类具有父类相同的行为。

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661876060993-91ac09c4-3a79-4b38-a09e-6450552a3bfe.png)

从实际的例子来看，就是动物是一个大父类，下面又能细分为 “食草动物”、“食肉动物”，这两者会包含 “动物” 这个父类的基本定义。

从实际的例子来看，就是动物是一个大父类，下面又能细分为 “食草动物”、“食肉动物”，这两者会包含 “动物” 这个父类的基本定义。

**在 Go 语言中，是没有类似** **extends** **关键字的这种继承的方式，在语言设计上采取的是组合的方式**：

```go
type Animal struct {
 	Name string
}

type Cat struct {
 	Animal
 	FeatureA string
}

type Dog struct {
 	Animal
 	FeatureB string
}
```

在上述例子中，我们声明了 Cat 和 Dog 结构体，其在内部匿名组合了 Animal 结构体。因此 Cat 和 Dog 的实例都可以调用 Animal 结构体的方法：

```go
func main() {
 	p := NewAnimal()
 	p.SetName("我是搬运工，去给煎鱼点赞~")

 	dog := Dog{Animal: *p}
 	fmt.Println(dog.GetName())
}
```

同时 Cat 和 Dog 的实例可以拥有自己的方法：

```go
func (dog *Dog) HelloWorld() {
 	fmt.Println("脑子进煎鱼了")
}

func (cat *Cat) HelloWorld() {
 	fmt.Println("煎鱼进脑子了")
}
```

上述例子能够正常包含调用 Animal 的相关属性和方法，也能够拥有自己的独立属性和方法，在 Go 语言中达到了类似继承的效果。

#### 多态

多态

面向对象中的 “多态” 指的同一个行为具有多种不同表现形式或形态的能力，具体是指一个类实例（对象）的相同方法在不同情形有不同表现形式。

多态也使得不同内部结构的对象可以共享相同的外部接口，也就是都是一套外部模板，内部实际是什么，只要符合规格就可以。

**在 Go 语言中，多态是通过接口来实现的：**

```go
type AnimalSounder interface {
 	MakeDNA()
}

func MakeSomeDNA(animalSounder AnimalSounder) {		// 参数是AnimalSounder接口类型
 	animalSounder.MakeDNA()
}
```

在上述例子中，我们声明了一个接口类型 AnimalSounder，配套一个 MakeSomeDNA 方法，其接受 AnimalSounder 接口类型作为入参。

因此在 Go 语言中。只要配套的 Cat 和 Dog 的实例也实现了 MakeSomeDNA 方法，那么我们就可以认为他是 AnimalSounder 接口类型：

```go
type AnimalSounder interface {
 	MakeDNA()
}

func MakeSomeDNA(animalSounder AnimalSounder) {
 	animalSounder.MakeDNA()
}

func (c *Cat) MakeDNA() {
 	fmt.Println("煎鱼是煎鱼")
}

func (c *Dog) MakeDNA() {
 	fmt.Println("煎鱼其实不是煎鱼")
}

func main() {
 	MakeSomeDNA(&Cat{})
 	MakeSomeDNA(&Dog{})
}
```

当 Cat 和 Dog 的实例实现了 AnimalSounder 接口类型的约束后，就意味着满足了条件，他们在 Go 语言中就是一个东西。能够作为入参传入 MakeSomeDNA 方法中，再根据不同的实例实现多态行为。

------

在日常工作中，基本了解这些概念就可以了。若是面试，可以针对三大特性：“封装、继承、多态” 和 五大原则 “单一职责原则（SRP）、开放封闭原则（OCP）、里氏替换原则（LSP）、依赖倒置原则（DIP）、接口隔离原则（ISP）” 进行深入理解和说明。

### 4、go语言和python的区别：

1、范例

Python是一种基于面向对象编程的多范式，命令式和函数式编程语言。它坚持这样一种观点，即如果一种语言在某些情境中表现出某种特定的方式，理想情况下它应该在所有情境中都有相似的作用。但是，它又不是纯粹的OOP语言，它不支持强封装，这是OOP的主要原则之一。

Go是一种基于并发编程范式的过程编程语言，它与C具有表面相似性。实际上，Go更像是C的更新版本。

2、类型化

Python是动态类型语言，而Go是一种静态类型语言，它实际上有助于在编译时捕获错误，这可以进一步减少生产后期的严重错误。

3、并发

Python没有提供内置的并发机制，而Go有内置的并发机制。

4、安全性

Python是一种强类型语言，它是经过编译的，因此增加了一层安全性。Go具有分配给每个变量的类型，因此，它提供了安全性。但是，如果发生任何错误，用户需要自己运行整个代码。

5、管理内存

Go允许程序员在很大程度上管理内存。而，Python中的内存管理完全自动化并由Python VM管理；它不允许程序员对内存管理负责。

6、库

与Go相比，Python提供的库数量要大得多。然而，Go仍然是新的，并且还没有取得很大进展。

7、语法

Python的语法使用缩进来指示代码块。Go的语法基于打开和关闭括号。

8、详细程度

为了获得相同的功能，Golang代码通常需要编写比Python代码更多的字符。

### 5、go 与 node.js

深入对比Node.js和Golang 到底谁才是NO.1 : https://zhuanlan.zhihu.com/p/421352168

从 Node 到 Go：一个粗略的比较 : https://zhuanlan.zhihu.com/p/29847628

## **一、基础部分**

### 0、为什么选择golang

**0、高性能-协程**
golang 源码级别支持协程，实现简单；对比进程和线程，协程占用资源少，能够简洁高效地处理高并发问题。

**1、学习曲线容易-****代码极简**

Go语言语法简单，包含了类C语法。因为Go语言容易学习，所以一个普通的大学生花几个星期就能写出来可以上手的、高性能的应用。在国内大家都追求快，这也是为什么国内Go流行的原因之一。

Go 语言的语法特性简直是太简单了，简单到你几乎玩不出什么花招，直来直去的，学习曲线很低，上手非常快。

**2、效率：快速的编译时间，开发效率和运行效率高**

开发过程中相较于 Java 和 C++呆滞的编译速度，Go 的快速编译时间是一个主要的效率优势。Go拥有接近C的运行效率和接近PHP的开发效率。

C 语言的理念是信任程序员，保持语言的小巧，不屏蔽底层且底层友好，关注语言的执行效率和性能。而 Python 的姿态是用尽量少的代码完成尽量多的事。于是我能够感觉到，Go 语言想要把 C 和 Python 统一起来，这是多棒的一件事啊。

**3、出身名门、血统纯正**

之所以说Go出身名门，从Go语言的创造者就可见端倪，Go语言绝对血统纯正。其次Go语言出自Google公司，Google在业界的知名度和实力自然不用多说。Google公司聚集了一批牛人，在各种编程语言称雄争霸的局面下推出新的编程语言，自然有它的战略考虑。而且从Go语言的发展态势来看，Google对它这个新的宠儿还是很看重的，Go自然有一个良好的发展前途。

**4、自由高效：组合的思想、无侵入式的接口**

Go语言可以说是开发效率和运行效率二者的完美融合，天生的并发编程支持。Go语言支持当前所有的编程范式，包括过程式编程、面向对象编程、面向接口编程、函数式编程。程序员们可以各取所需、自由组合、想怎么玩就怎么玩。

**5、强大的标准库-****生态**

背靠谷歌，生态丰富，轻松 go get 获取各种高质量轮子。用户可以专注于业务逻辑，避免重复造轮子。

这包括互联网应用、系统编程和网络编程。Go里面的标准库基本上已经是非常稳定了，特别是我这里提到的三个，网络层、系统层的库非常实用。Go 语言的 lib 库麻雀虽小五脏俱全。Go 语言的 lib 库中基本上有绝大多数常用的库，虽然有些库还不是很好，但我觉得不是问题，因为我相信在未来的发展中会把这些问题解决掉。

**6、部署方便：二进制文件，Copy部署**

部署简单，源码编译成执行文件后，可以直接运行，减少了对其它插件依赖。不像其它语言，执行文件依赖各种插件，各种库，研发机器运行正常，部署到生产环境，死活跑不起来 。

**7、简单的并发**

并行和异步编程几乎无痛点。Go 语言的 Goroutine 和 Channel 这两个神器简直就是并发和异步编程的巨大福音。像 C、C++、Java、Python 和 JavaScript 这些语言的并发和异步方式太控制就比较复杂了，而且容易出错，而 Go 解决这个问题非常地优雅和流畅。这对于编程多年受尽并发和异步折磨的编程者来说，完全就是让人眼前一亮的感觉。Go 是一种非常高效的语言，高度支持并发性。Go是为大数据、微服务、并发而生的一种编程语言。

Go 作为一门语言致力于使事情简单化。它并未引入很多新概念，而是聚焦于打造一门简单的语言，它使用起来异常快速并且简单。其唯一的创新之处是 goroutines 和通道。Goroutines 是 Go 面向线程的轻量级方法，而通道是 goroutines 之间通信的优先方式。

创建 Goroutines 的成本很低，只需几千个字节的额外内存，正由于此，才使得同时运行数百个甚至数千个 goroutines 成为可能。可以借助通道实现 goroutines 之间的通信。Goroutines 以及基于通道的并发性方法使其非常容易使用所有可用的 CPU 内核，并处理并发的 IO。相较于 Python/Java，在一个 goroutine 上运行一个函数需要最小的代码。

**8、稳定性**

Go拥有强大的编译检查、严格的编码规范和完整的软件生命周期工具，具有很强的稳定性，稳定压倒一切。那么为什么Go相比于其他程序会更稳定呢？这是因为Go提供了软件生命周期（开发、测试、部署、维护等等）的各个环节的工具，如go tool、gofmt、go test。

**9、跨平台**
很多语言都支持跨平台，把这个优点单独拿出来，貌似没有什么值得称道的，但是结合上述优点，它的综合能力就非常强了。

### golang 缺点

**①右大括号不允许换行，否则编译报错**

**②不允许有未使用的包或变量**

**③错误处理原始，虽然引入了defer、panic、recover处理出错后的逻辑，函数可以返回多个值，但基本依靠返回错误是否为空来判断函数是否执行成功，if err != nil语句较多，比较繁琐，程序没有java美观。**(官方解释：提供了多个返回值，处理错误方便，如加入异常机制会要求记住一些常见异常，例如IOException，go的错误Error类型较统一方便)
**④[]interface{}不支持下标操作**

**⑤struct没有构造和析构，一些资源申请和释放动作不太方便**

**⑥仍然保留C/C++的指针操作，取地址&，取值\***

### **1、golang 中 make 和 new 的区别？（基本必问）**

**共同点：**给变量分配内存

**不同点：**

1）作用变量类型不同，new给string,int和数组分配内存，make给切片，map，channel分配内存；

2）返回类型不一样，new返回指向变量的指针，make返回变量本身；

3）new 分配的空间被清零。make 分配空间后，会进行初始化；

\4) 字节的面试官还说了另外一个区别，就是分配的位置，在堆上还是在栈上？这块我比较模糊，大家可以自己探究下，我搜索出来的答案是golang会弱化分配的位置的概念，因为编译的时候会自动内存逃逸处理，懂的大佬帮忙补充下：make、new内存分配是在堆上还是在栈上？

new和make都在堆上分配内存
new 函数分配内存，make 函数初始化
https://www.cnblogs.com/chenpingzhao/p/9918062.html

### 2、[IO多路复用](https://zhuanlan.zhihu.com/p/115220699)

### **3、for range 的时候它的地址会发生变化么？**

答：(go1.21之前)在 for a,b := range c 遍历中， a 和 b 在内存中只会存在一份，即之后每次循环时遍历到的数据都是以值覆盖的方式赋给 a 和 b，a，b 的内存地址始终不变。由于有这个特性，for 循环里面如果开协程，不要直接把 a 或者 b 的地址传给协程。解决办法：在每次循环时，创建一个临时变量。
(go1.21开启GOEXPERIMENT=loopvar，go1.22默认开启)在 for range 时，循环时变量不再共享同样的内存地址，即不再需要创建临时变量。

### **4、go defer，多个 defer 的顺序，defer 在什么时机会修改返回值？**

https://www.topgoer.cn/docs/golangxiuyang/golangxiuyang-1cmee0q64ij5p

作用：defer延迟函数，释放资源，收尾工作；如释放锁，关闭文件，关闭链接；捕获panic;

避坑指南：defer函数紧跟在资源打开后面，否则defer可能得不到执行，导致内存泄露。

多个 defer 调用顺序是 LIFO（后入先出），defer后的操作可以理解为压入栈中

defer，return，return value（函数返回值） 执行顺序：首先return，其次return value，最后defer。defer可以修改函数最终返回值，修改时机：**有名返回值或者函数返回指针** 参考：

[【Golang】Go语言defer用法大总结(含return返回机制)__奶酪的博客-CSDN博客blog.csdn.net/Cassie_zkq/article/details/108567205](https://link.zhihu.com/?target=https%3A//blog.csdn.net/Cassie_zkq/article/details/108567205)

**有名返回值**

```go
func b() (i int) { 	
    defer func() { 		
        i++ 		
        fmt.Println("defer2:", i) 	
    }() 	
    defer func() { 		
        i++ 		
        fmt.Println("defer1:", i) 	
    }() 	
    return i 
    //或者直接写成
    return 
} 
func main() { 	
    fmt.Println("return:", b()) 
} 
```



**函数返回指针**

```go
func c() *int { 	
    var i int 	
    defer func() { 		
        i++ 		
        fmt.Println("defer2:", i) 	
    }() 	
    defer func() { 		
        i++ 		
        fmt.Println("defer1:", i) 	
    }() 	
    return &i 
} 
func main() { 	
    fmt.Println("return:", *(c())) 
}
```

### **5、uint 类型溢出问题**

超过最大存储值如uint8最大是255

var a uint8 =255

var b uint8 =1

a+b = 0总之类型溢出会出现难以意料的事

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1659259378774-f5ecf978-5d67-4b9a-bd37-47d569ba7353.png)

### **6、能介绍下 rune 类型吗？**

相当int32

golang中的字符串底层实现是通过byte数组的，中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8

byte 等同于uint8，常用来处理ascii字符

rune 等同于int32,常用来处理unicode或utf-8字符

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1659259378747-48538a44-1ccb-47ac-9492-0b569d219e2b.png)

### **7、 golang 中解析 tag 是怎么实现的？反射原理是什么？(中高级肯定会问，比较难，需要自己多去总结)**

**参考如下连接**

[golang中struct关于反射tag_paladinosment的博客-CSDN博客_golang 反射tagblog.csdn.net/paladinosment/article/details/42570937](https://link.zhihu.com/?target=https%3A//blog.csdn.net/paladinosment/article/details/42570937)

type User struct { 	name string `json:name-field` 	age  int } func main() { 	user := &User{"John Doe The Fourth", 20} 	field, ok := reflect.TypeOf(user).Elem().FieldByName("name") 	if !ok { 		panic("Field not found") 	} 	fmt.Println(getStructTag(field)) } func getStructTag(f reflect.StructField) string { 	return string(f.Tag) }

Go 中解析的 tag 是通过反射实现的，反射是指计算机程序在运行时（Run time）可以访问、检测和修改它本身状态或行为的一种能力或动态知道给定数据对象的类型和结构，并有机会修改它。反射将接口变量转换成反射对象 Type 和 Value；反射可以通过反射对象 Value 还原成原先的接口变量；反射可以用来修改一个变量的值，前提是这个值可以被修改；tag是啥:结构体支持标记，name string `json:name-field` 就是 `json:name-field` 这部分

**gorm json yaml gRPC protobuf gin.Bind()都是通过反射来实现的**

### **8、调用函数传入结构体时，应该传值还是指针？ （Golang 都是传值）**

Go 的函数参数传递都是值传递。所谓值传递：指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。参数传递还有引用传递，所谓引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数

因为 Go 里面的 map，slice，chan 是引用类型。变量区分值类型和引用类型。所谓值类型：变量和变量的值存在同一个位置。所谓引用类型：变量和变量的值是不同的位置，变量的值存储的是对值的引用。但并不是 map，slice，chan 的所有的变量在函数内都能被修改，不同数据类型的底层存储结构和实现可能不太一样，情况也就不一样。

### 9、goroutine什么情况下会阻塞

在 Go 里面阻塞主要分为以下 4 种场景：

1. 由于原子、互斥量或通道操作调用导致 Goroutine 阻塞，调度器将把当前阻塞的 Goroutine 切换出去，重新调度 LRQ 上的其他 Goroutine；
2. 由于网络请求和 IO 操作导致 Goroutine 阻塞。Go 程序提供了网络轮询器（NetPoller）来处理网络请求和 IO 操作的问题，其后台通过 kqueue（MacOS），epoll（Linux）或 iocp（Windows）来实现 IO 多路复用。通过**使用 NetPoller 进行网络系统调用**，调度器可以防止 Goroutine 在进行这些系统调用时阻塞 M。这可以让 M 执行 P 的 LRQ 中其他的 Goroutines，而不需要创建新的 M。执行网络系统调用不需要额外的 M，**网络轮询器使用系统线程**，它时刻处理一个有效的事件循环，有助于减少操作系统上的调度负载。用户层眼中看到的 Goroutine 中的“block socket”，实现了 goroutine-per-connection 简单的网络编程模式。实际上是通过 Go runtime 中的 netpoller 通过 Non-block socket + I/O 多路复用机制“模拟”出来的。
3. 当调用一些系统方法的时候（如文件 I/O），如果系统方法调用的时候发生阻塞，这种情况下，网络轮询器（NetPoller）无法使用，而进行系统调用的 G1 将阻塞当前 M1。调度器引入 其它M 来服务 M1 的P。
4. 如果在 Goroutine 去执行一个 sleep 操作，导致 M 被阻塞了。Go 程序后台有一个监控线程 sysmon，它监控那些长时间运行的 G 任务然后设置可以强占的标识符，别的 Goroutine 就可以抢先进来执行。

### **10、讲讲 Go 的 select 底层数据结构和一些特性？（难点，没有项目经常可能说不清，面试一般会问你项目中怎么使用select）**

答：go 的 select 为 golang 提供了多路 IO 复用机制，和其他 IO 复用一样，用于检测是否有读写事件是否 ready。linux 的系统 IO 模型有 select，poll，epoll，go 的 select 和 linux 系统 select 非常相似。

select 结构组成主要是由 case 语句和执行的函数组成 select 实现的多路复用是：每个线程或者进程都先到注册和接受的 channel（装置）注册，然后阻塞，然后只有一个线程在运输，当注册的线程和进程准备好数据后，装置会根据注册的信息得到相应的数据。

**select 的特性**

1）select 操作至少要有一个 case 语句，出现读写 nil 的 channel 该分支会忽略，在 nil 的 channel 上操作则会报错。

2）select 仅支持管道，而且是单协程操作。

3）每个 case 语句仅能处理一个管道，要么读要么写。

4）多个 case 语句的执行顺序是随机的。

5）存在 default 语句，select 将不会阻塞，但是存在 default 会影响性能。

### **11、讲讲 Go 的 defer 底层数据结构和一些特性？**

答：每个 defer 语句都对应一个_defer 实例，多个实例使用指针连接起来形成一个单连表，保存在 gotoutine 数据结构中，每次插入_defer 实例，均插入到链表的头部，函数结束再一次从头部取出，从而形成后进先出的效果。

**defer 的规则总结**：

延迟函数的参数是 defer 语句出现的时候就已经确定了的。

延迟函数执行按照后进先出的顺序执行，即先出现的 defer 最后执行。

延迟函数可能操作主函数的返回值。

申请资源后立即使用 defer 关闭资源是个好习惯。

### **12、单引号，双引号，反引号的区别？**

单引号，表示byte类型或rune类型，对应 uint8和int32类型，默认是 rune 类型。byte用来强调数据是raw data，而不是数字；而rune用来表示Unicode的code point。

双引号，才是字符串，实际上是字符数组。可以用索引号访问某字节，也可以用len()函数来获取字符串所占的字节长度。

反引号，表示字符串字面量，但不支持任何转义序列。字面量 raw literal string 的意思是，你定义时写的啥样，它就啥样，你有换行，它就换行。你写转义字符，它也就展示转义字符。

### 13、go出现panic的场景

### https://www.cnblogs.com/paulwhw/p/15585467.html

- - 数组/切片越界
  - 空指针调用。比如访问一个 nil 结构体指针的成员
  - 过早关闭 HTTP 响应体
  - 除以 0
  - 向已经关闭的 channel 发送消息
  - 重复关闭 channel
  - 关闭未初始化的 channel
  - 未初始化 map。注意访问 map 不存在的 key 不会 panic，而是返回 map 类型对应的零值，但是不能直接赋值
  - 跨协程的 panic 处理
  - sync 计数为负数。
  - 类型断言不匹配。`var a interface{} = 1; fmt.Println(a.(string))` 会 panic，建议用 `s,ok := a.(string)`

### 14、go是否支持while循环，如何实现这种机制

https://blog.csdn.net/chengqiuming/article/details/115573947

### 15、go里面如何实现set？

Go中是不提供Set类型的，Set是一个集合，其本质就是一个List，只是List里的元素不能重复。

Go提供了map类型，但是我们知道，map类型的key是不能重复的，因此，我们可以利用这一点，来实现一个set。那value呢？value我们可以用一个常量来代替，比如一个空结构体，实际上空结构体不占任何内存，使用空结构体，能够帮我们节省内存空间，提高性能

代码实现：https://blog.csdn.net/haodawang/article/details/80006059

### 16、go如何实现类似于java当中的继承机制？

https://zhuanlan.zhihu.com/p/88480107

说到继承我们都知道，在Go中没有extends关键字，也就意味着Go并没有原生级别的继承支持。这也是为什么我在文章开头用了**伪继承**这个词。本质上，Go使用interface实现的功能叫组合，Go是使用组合来实现的继承，说的更精确一点，是使用组合来代替的继承，举个很简单的例子:

**通过组合实现了继承：**

```go
type Animal struct {
    Name string
}

func (a *Animal) Eat() {
    fmt.Printf("%v is eating", a.Name)
    fmt.Println()
}

type Cat struct {
    *Animal
}

cat := &Cat{
    Animal: &Animal{
        Name: "cat",
    },
}
cat.Eat() // cat is eating
```

首先，我们实现了一个Animal的结构体，代表动物类。并声明了Name字段，用于描述动物的名字。

然后，实现了一个以Animal为receiver的Eat方法，来描述动物进食的行为。

最后，声明了一个Cat结构体，组合了Cat字段。再实例化一个猫，调用Eat方法，可以看到会正常的输出。

可以看到，Cat结构体本身没有Name字段，也没有去实现Eat方法。唯一有的就是组合了Animal父类，至此，我们就证明了已经通过组合实现了继承。

**总结：**

- 如果一个 struct 嵌套了另一个匿名结构体，那么这个结构可以直接访问匿名结构体的属性和方法，从而实现继承。
- 如果一个 struct 嵌套了另一个有名的结构体，那么这个模式叫做组合。
- 如果一个 struct 嵌套了多个匿名结构体，那么这个结构可以直接访问多个匿名结构体的属性和方法，从而实现多重继承。

### 17、怎么去复用一个接口的方法？

https://www.yisu.com/zixun/452409.html

### 18、go里面的 _ 

1. **忽略返回值**

1. 1. 比如某个函数返回三个参数，但是我们只需要其中的两个，另外一个参数可以忽略，这样的话代码可以这样写：

```go
v1, v2, _ := function(...)
v1, _, _ := function(...)
```

1. **用在变量(特别是接口断言)**

```go
type T struct{}
var _ X = T{}
//其中 I为interface
```

上面用来判断 type T是否实现了X,用作类型断言，如果T没有实现接口X，则编译错误.

1. **用在import package**

```go
import _ "test/food"
```

引入包时，会先调用包中的初始化函数，这种使用方式仅让导入的包做初始化，而不使用包中其他功能

### 19、goroutine创建的时候如果要传一个参数进去有什么要注意的点？

https://www.cnblogs.com/waken-captain/p/10496454.html

### 20、写go单元测试的规范？

1.  **单元测试文件命名规则 ：**

单元测试需要创建单独的测试文件，不能在原有文件中书写，名字规则为 xxx_test.go。这个规则很好理解。

1.  **单元测试包命令规则** 

单元测试文件的包名为原文件的包名添加下划线接test，举例如下：

```go
// 原文件包名：

package xxx

// 单元测试文件包名：

package xxx_test
```

1.  **单元测试方法命名规则** 

单元测试文件中的测试方法和原文件中的待测试的方法名相对应，以Test开头，举例如下：

```go
// 原文件方法：
func Xxx(name string) error 
 
// 单元测试文件方法：
func TestXxx()
```

1.  **单元测试方法参数** 

单元测试方法的参数必须是t *testing.T，举例如下：

```go
func TestZipFiles(t *testing.T) { ...
```

### 21、单步调试？

https://www.jianshu.com/p/21ed30859d80

### 22、导入一个go的工程，有些依赖找不到，改怎么办？

https://www.cnblogs.com/niuben/p/16182001.html

### 23、[值拷贝 与 引用拷贝，深拷贝 与 浅拷贝](https://www.cnblogs.com/yizhixiaowenzi/p/14664222.html)

map，slice，chan 是引用拷贝；引用拷贝 是 浅拷贝

其余的，都是 值拷贝；值拷贝 是 深拷贝

#### 深浅拷贝的本质区别：

是否真正获取对象实体，而不是引用

**深拷贝：**

拷贝的是数据本身，创造一个新的对象，并在内存中开辟一个新的内存地址，与原对象是完全独立的，不共享内存，修改新对象时不会影响原对象的值。释放内存时，也没有任何关联。

**值拷贝：**

接收的是  整个array的值拷贝，所以方法对array中元素的重新赋值不起作用。

```go
package main  

import "fmt"  

func modify(a [3]int) {  
    a[0] = 4  
    fmt.Println("modify",a)             // modify [4 2 3]
}  

func main() {  
    a := [3]int{1, 2, 3}  
    modify(a)  
    fmt.Println("main",a)                  // main [1 2 3]
}  
```

**浅拷贝：**

拷贝的是数据地址，只复制指向的对象的指针，新旧对象的内存地址是一样的，修改一个另一个也会变。释放内存时，同时释放。

**引用拷贝：**

函数的引用拷贝与原始的引用指向同一个数组，所以对数组中元素的修改，是有效的

```go
package main  
  
import "fmt"  
  
func modify(s []int) {  
    s[0] = 4  
    fmt.Println("modify",s)          // modify [4 2 3]
}  
  
func main() {  
    s := []int{1, 2, 3}  
    modify(s)  
    fmt.Println("main",s)              // main [4 2 3]
}
```

### 24、[精通Golang项目依赖Go modules](https://www.topgoer.cn/docs/golangxiuyang/golangxiuyang-1cmee13oek1e8)

### 25、Go 多返回值怎么实现的？

答：Go 传参和返回值是通过 FP+offset 实现，并且存储在调用函数的栈帧中。FP 栈底寄存器，指向一个函数栈的顶部;PC 程序计数器，指向下一条执行指令;SB 指向静态数据的基指针，全局符号;SP 栈顶寄存器。

### 26、Go 语言中不同的类型如何比较是否相等？

答：像 string，int，float interface 等可以通过 reflect.DeepEqual 和等于号进行比较，像 slice，struct，map 则一般使用 reflect.DeepEqual 来检测是否相等。

### 27、Go中init 函数的特征?

答：一个包下可以有多个 init 函数，每个文件也可以有多个 init 函数。多个 init 函数按照它们的文件名顺序逐个初始化。应用初始化时初始化工作的顺序是，从被导入的最深层包开始进行初始化，层层递出最后到 main 包。不管包被导入多少次，包内的 init 函数只会执行一次。应用初始化时初始化工作的顺序是，从被导入的最深层包开始进行初始化，层层递出最后到 main 包。但包级别变量的初始化先于包内 init 函数的执行。

### 28、Go中 uintptr和 unsafe.Pointer 的区别？

答：unsafe.Pointer 是通用指针类型，它不能参与计算，任何类型的指针都可以转化成 unsafe.Pointer，unsafe.Pointer 可以转化成任何类型的指针，uintptr 可以转换为 unsafe.Pointer，unsafe.Pointer 可以转换为 uintptr。uintptr 是指针运算的工具，但是它不能持有指针对象（意思就是它跟指针对象不能互相转换），unsafe.Pointer 是指针对象进行运算（也就是 uintptr）的桥梁。

## 二、slice

### **1、数组和切片的区别 （基本必问）**

**相同点：**

1)只能存储一组相同类型的数据结构

2)都是通过下标来访问，并且有容量长度，长度通过 len 获取，容量通过 cap 获取

**区别：**

1）数组是定长，访问和复制不能超过数组定义的长度，否则就会下标越界，切片长度和容量可以自动扩容

2）数组是值类型，切片是引用类型，每个切片都引用了一个底层数组，切片本身不能存储任何数据，都是这底层数组存储数据，所以修改切片的时候修改的是底层数组中的数据。切片一旦扩容，指向一个新的底层数组，内存地址也就随之改变

**简洁的回答：**

1）定义方式不一样 2）初始化方式不一样，数组需要指定大小，大小不改变 3）在函数传递中，数组切片都是值传递。

**数组的定义**

var a1 [3]int

var a2 [...]int{1,2,3}

**切片的定义**

var a1 []int

var a2 :=make([]int,3,5)

**数组的初始化**

a1 := [...]int{1,2,3}

a2 := [5]int{1,2,3}

**切片的初始化**

b:= make([]int,3,5)



### **2、**[**讲讲 Go 的 slice 底层数据结构和一些特性？**](https://www.topgoer.cn/docs/gozhuanjia/gozhuanjiaslice)

答：Go 的 slice 底层数据结构是由一个 array 指针指向底层数组，len 表示切片长度，cap 表示切片容量。slice 的主要实现是扩容。对于 append 向 slice 添加元素时，假如 slice 容量够用，则追加新元素进去，slice.len++，返回原来的 slice。当原容量不够，则 slice 先扩容，扩容之后 slice 得到新的 slice，将元素追加进新的 slice，slice.len++，返回新的 slice。对于切片的扩容规则：当切片比较小时（容量小于 1024），则采用较大的扩容倍速进行扩容（新的扩容会是原来的 2 倍），避免频繁扩容，从而减少内存分配的次数和数据拷贝的代价。当切片较大的时（原来的 slice 的容量大于或者等于 1024），采用较小的扩容倍速（新的扩容将扩大大于或者等于原来 1.25 倍），主要避免空间浪费，网上其实很多总结的是 1.25 倍，那是在不考虑内存对齐的情况下，实际上还要考虑内存对齐，扩容是大于或者等于 1.25 倍。

（关于刚才问的 slice 为什么传到函数内可能被修改，如果 slice 在函数内没有出现扩容，函数外和函数内 slice 变量指向是同一个数组，则函数内复制的 slice 变量值出现更改，函数外这个 slice 变量值也会被修改。如果 slice 在函数内出现扩容，则函数内变量的值会新生成一个数组（也就是新的 slice，而函数外的 slice 指向的还是原来的 slice，则函数内的修改不会影响函数外的 slice。）

### 3、golang中数组和slice作为参数的区别？slice作为参数传递有什么问题？

https://blog.csdn.net/weixin_44387482/article/details/119763558

1. 当使用数组作为参数和返回值的时候，传进去的是值，在函数内部对数组进行修改并不会影响原数据
2. 当切片作为参数的时候穿进去的是值，也就是值传递，但是当我在函数里面修改切片的时候，我们发现源数据也会被修改，这是因为我们在切片的底层维护这一个匿名的数组，当我们把切片当成参数的时候，会重现创建一个切片，但是创建的这个切片和我们原来的数据是共享数据源的，所以在函数内被修改，源数据也会被修改
3. 数组还是切片，在函数中传递的时候如果没有指定为指针传递的话，都是值传递，但是切片在传递的过程中，有着共享底层数组的风险，所以如果在函数内部进行了更改的时候，会修改到源数据，所以我们需要根据不同的需求来处理，如果我们不希望源数据被修改话的我们可以使用copy函数复制切片后再传入，如果希望源数据被修改的话我们应该使用指针传递的方式

### 4、从数组中取一个相同大小的slice有成本吗？

或者这么问：从切片中取一个相同大小的数组有成本吗？



从数组中截取切片：https://blog.csdn.net/weixin_42117918/article/details/81913036



## **三、map相关**

### 1、map 使用注意的点，是否并发安全？

map的类型是map[key]，key类型的ke必须是可比较的，通常情况，会选择内建的基本类型，比如整数、字符串做key的类型。如果要使用struct作为key，要保证struct对象在逻辑上是不可变的。在Go语言中，map[key]函数返回结果可以是一个值，也可以是两个值。map是无序的，如果我们想要保证遍历map时元素有序，可以使用辅助的数据结构，例如orderedmap。

**第一，**一定要先初始化，否则panic

**第二，**map类型是容易发生并发访问问题的。不注意就容易发生程序运行时并发读写导致的panic。 Go语言内建的map对象不是线程安全的，并发读写的时候运行时会有检查，遇到并发问题就会导致panic。

### 2、map 循环是有序的还是无序的？

无序的, map 因扩张⽽重新哈希时，各键值项存储位置都可能会发生改变，顺序自然也没法保证了，所以官方避免大家依赖顺序，直接打乱处理。就是 for range map 在开始处理循环逻辑的时候，就做了随机播种

### 3、 map 中删除一个 key，它的内存会释放么？（常问）

如果删除的元素是值类型，如int，float，bool，string以及数组和struct，map的内存不会自动释放

如果删除的元素是引用类型，如指针，slice，map，chan等，map的内存会自动释放，但释放的内存是子元素应用类型的内存占用

将map设置为nil后，内存被回收。

**这个问题还需要大家去搜索下答案，我记得有不一样的说法，谨慎采用本题答案。**

### 4、怎么处理对 map 进行并发访问？有没有其他方案？ 区别是什么？

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1659259378747-2e47defb-6941-4481-af64-a2ca6dda832d.png)

**方式一、使用内置sync.Map，详细参考**

[https://mbd.baidu.com/ma/s/7Hwd9yMcmbd.baidu.com/ma/s/7Hwd9yMc](https://link.zhihu.com/?target=https%3A//mbd.baidu.com/ma/s/7Hwd9yMc)

**方式二、使用读写锁实现并发安全map**

[https://mbd.baidu.com/ma/s/qO7b0VQUmbd.baidu.com/ma/s/qO7b0VQU](https://link.zhihu.com/?target=https%3A//mbd.baidu.com/ma/s/qO7b0VQU)



https://cloud.tencent.com/developer/article/1539049

### 5、 nil map 和空 map 有何不同？

1）可以对未初始化的map进行取值，但取出来的东西是空：

```go
var m1 map[string]string
fmt.Println(m1["1"])
```

2）不能对未初始化的map进行赋值，这样将会抛出一个异常：

未初始化的map是nil，它与一个空map基本等价，只是nil的map不允许往里面添加值。

```go
var m1 map[string]string
m1["1"] = "1"
panic: assignment to entry in nil map

因此，map是nil时，取值是不会报错的（取不到而已），但增加值会报错。

其实，还有一个区别，delete一个nil map会panic，
但是delete 空map是一个空操作（并不会panic）
（这个区别在最新的Go tips中已经没有了，即：delete一个nil map也不会panic）
```

\3) 通过fmt打印map时，空map和nil map结果是一样的，都为map[]。所以，这个时候别断定map是空还是nil，而应该通过map == nil来判断。

**nil map 未初始化，空map是长度为空**

### 6、map 的数据结构是什么？

https://www.topgoer.cn/docs/gozhuanjia/gozhuanjiamap

答：golang 中 map 是一个 kv 对集合。底层使用 hash table，用链表来解决冲突 ，出现冲突时，不是每一个 key 都申请一个结构通过链表串起来，而是以 bmap 为最小粒度挂载，一个 bmap 可以放 8 个 kv。在哈希函数的选择上，会在程序启动时，检测 cpu 是否支持 aes，如果支持，则使用 aes hash，否则使用 memhash。每个 map 的底层结构是 hmap，是有若干个结构为 bmap 的 bucket 组成的数组。每个 bucket 底层都采用链表结构。

#### hmap 的结构如下：

```go
type hmap struct {     
    count     int                  // 元素个数     
    flags     uint8     
    B         uint8                // 扩容常量相关字段B是buckets数组的长度的对数 2^B     
    noverflow uint16               // 溢出的bucket个数     
    hash0     uint32               // hash seed     
    buckets    unsafe.Pointer      // buckets 数组指针     
    oldbuckets unsafe.Pointer      // 结构扩容的时候用于赋值的buckets数组     
    nevacuate  uintptr             // 搬迁进度     
    extra *mapextra                // 用于扩容的指针 
}
```

**下图展示一个拥有4个bucket的map：**

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661789793109-401b7c75-c26b-4893-bbf7-1f2dfa69316b.png)

本例中, hmap.B=2， 而hmap.buckets长度是2^B为4. 元素经过哈希运算后会落到某个bucket中进行存储。查找过程类似。

bucket很多时候被翻译为桶，所谓的哈希桶实际上就是bucket。

#### bucket数据结构

bucket数据结构由runtime/map.go:bmap定义：

```go
// A bucket for a Go map.
type bmap struct {
    // tophash generally contains the top byte of the hash value
    // for each key in this bucket. If tophash[0] < minTopHash,
    // tophash[0] is a bucket evacuation state instead.
    tophash [bucketCnt]uint8
    // Followed by bucketCnt keys and then bucketCnt elems.
    // NOTE: packing all the keys together and then all the elems together makes the
    // code a bit more complicated than alternating key/elem/key/elem/... but it allows
    // us to eliminate padding which would be needed for, e.g., map[int64]int8.
    // Followed by an overflow pointer.
}
```

这里只有一个 tophash 字段，而实际上在使用中值的类型是不固定的，甚至可以是一个自定义结构体的指针类型。这个结构体看起来可能有点让人费解，其实编译器在编译期间会动态创建一个新的同名数据结构，如下所示

```go
type bmap struct {
    topbits  [8]uint8
    keys     [8]keytype
    values   [8]valuetype
    pad      uintptr
    overflow uintptr
}
```

bmap 即 bucket map 的缩写。

每个bucket可以存储8个键值对。

- topbits 是个长度为8的数组，哈希值相同的键（准确的说是哈希值低位相同的键）存入当前bucket时会将哈希值的高位存储在该数组中，以方便后续匹配。
- keys 长度为8的数组，[]keytype，元素为：具体的key值。
- values 长度为8的数组，[]valuetype，元素为：键值对的key对应的值。
- pad 对齐内存使用的，不是每个bmap都有会这个字段，需要满足一定条件
- overflow 指针指向的是下一个bucket，据此将所有冲突的键连接起来。

下图展示bucket存放8个key-value对：

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661789834784-c60b0cb4-96be-4c4c-8978-2bfc9ca716b9.png)

#### [解决哈希冲突（四种方法）](https://blog.csdn.net/qq_48241564/article/details/118613312)

#### 哈希冲突

当有两个或以上数量的键被哈希到了同一个bucket时，我们称这些键发生了冲突。Go使用链地址法来解决键冲突。
由于每个bucket可以存放8个键值对，所以同一个bucket存放超过8个键值对时就会再创建一个键值对，用类似链表的方式将bucket连接起来。

下图展示产生冲突后的map：

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661789900886-a77838be-46c8-4254-999b-b6e217721fbf.png)

bucket数据结构指示下一个bucket的指针称为overflow bucket，意为当前bucket盛不下而溢出的部分。事实上哈希冲突并不是好事情，它降低了存取效率，好的哈希算法可以保证哈希值的随机性，但冲突过多也是要控制的，后面会再详细介绍。

#### 负载因子

负载因子用于衡量一个哈希表冲突情况，公式为：

负载因子 = 键数量/bucket数量

例如，对于一个bucket数量为4，包含4个键值对的哈希表来说，这个哈希表的负载因子为1.

哈希表需要将负载因子控制在合适的大小，超过其阀值需要进行rehash，也即键值对重新组织：

- 哈希因子过小，说明空间利用率低
- 哈希因子过大，说明冲突严重，存取效率低

每个哈希表的实现对负载因子容忍程度不同，比如Redis实现中负载因子大于1时就会触发rehash，而Go则在在负载因子达到6.5时才会触发rehash，因为Redis的每个bucket只能存1个键值对，而Go的bucket可能存8个键值对，所以Go可以容忍更高的负载因子。

### 7、是怎么实现扩容？

#### map 的容量大小

底层调用 makemap 函数，计算得到合适的 B，map 容量最多可容纳 6.52^B 个元素，6.5 为装载因子阈值常量。装载因子的计算公式是：装载因子=填入表中的元素个数/散列表的长度，装载因子越大，说明空闲位置越少，冲突越多，散列表的性能会下降。

#### 触发 map 扩容的条件

为了保证访问效率，当新元素将要添加进map时，都会检查是否需要扩容，扩容实际上是以空间换时间的手段。
触发扩容的条件有二个：

1. 负载因子 > 6.5时，也即平均每个bucket存储的键值对达到6.5个。
2. overflow数量 > 2^15时，也即overflow数量超过32768时。

#### 增量扩容

当负载因子过大时，就新建一个bucket，新的bucket长度是原来的2倍，然后旧bucket数据搬迁到新的bucket。
考虑到如果map存储了数以亿计的key-value，一次性搬迁将会造成比较大的延时，Go采用逐步搬迁策略，即每次访问map时都会触发一次搬迁，每次搬迁2个键值对。

下图展示了包含一个bucket满载的map(为了描述方便，图中bucket省略了value区域):

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661789723150-6a635c5e-5d5a-4173-972f-ac0fd0326ffe.png)

当前map存储了7个键值对，只有1个bucket。此地负载因子为7。再次插入数据时将会触发扩容操作，扩容之后再将新插入键写入新的bucket。

当第8个键值对插入时，将会触发扩容，扩容后示意图如下：

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661789723181-66b62c5f-34bb-4427-8c68-446e7e05b4de.png)

hmap数据结构中oldbuckets成员指身原bucket，而buckets指向了新申请的bucket。新的键值对被插入新的bucket中。
后续对map的访问操作会触发迁移，将oldbuckets中的键值对逐步的搬迁过来。当oldbuckets中的键值对全部搬迁完毕后，删除oldbuckets。

搬迁完成后的示意图如下：

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661789723183-d1c03c9d-b6a9-4dd7-8410-a2674f1f1c0c.png)

数据搬迁过程中原bucket中的键值对将存在于新bucket的前面，新插入的键值对将存在于新bucket的后面。
实际搬迁过程中比较复杂，将在后续源码分析中详细介绍。

#### 等量扩容

所谓等量扩容，实际上并不是扩大容量，buckets数量不变，重新做一遍类似增量扩容的搬迁动作，把松散的键值对重新排列一次，以使bucket的使用率更高，进而保证更快的存取。
在极端场景下，比如不断地增删，而键值对正好集中在一小部分的bucket，这样会造成overflow的bucket数量增多，但负载因子又不高，从而无法执行增量搬迁的情况，如下图所示：

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661789747828-6f31463b-a48d-4a4d-877b-828f7f6abc9d.png)

上图可见，overflow的bucket中大部分是空的，访问效率会很差。此时进行一次等量扩容，即buckets数量不变，经过重新组织后overflow的bucket数量会减少，即节省了空间又会提高访问效率。

### 8、查找过程

查找过程如下：

1. 根据key值算出哈希值
2. 取哈希值低位与hmap.B取模确定bucket位置
3. 取哈希值高位在tophash数组中查询
4. 如果tophash[i]中存储值也哈希值相等，则去找到该bucket中的key值进行比较
5. 当前bucket没有找到，则继续从下个overflow的bucket中查找。
6. 如果当前处于搬迁过程，则优先从oldbuckets查找

注：如果查找不到，也不会返回空值，而是返回相应类型的0值。

### 9、插入过程

新元素插入过程如下：

1. 根据key值算出哈希值
2. 取哈希值低位与hmap.B取模确定bucket位置
3. 查找该key是否已经存在，如果存在则直接更新值
4. 如果没找到将key，将key插入

### 10、slices能作为map类型的key吗？

当时被问的一脸懵逼，其实是这个问题的变种：golang 哪些类型可以作为map key？

答案是：**在golang规范中，可比较的类型都可以作为map key；**这个问题又延伸到在：golang规范中，哪些数据类型可以比较？

**不能作为map key 的类型包括：**

- slices
- maps
- functions

详细参考：

[golang 哪些类型可以作为map keyblog.csdn.net/lanyang123456/article/details/123765745](https://link.zhihu.com/?target=https%3A//blog.csdn.net/lanyang123456/article/details/123765745)

## 四、接口

### 1、[Go 语言与鸭子类型的关系](http://golang.design/go-questions/interface/duck-typing/)

总结一下，鸭子类型是一种动态语言的风格，在这种风格中，一个对象有效的语义，不是由继承自特定的类或实现特定的接口，而是由它"当前方法和属性的集合"决定。Go 作为一种静态语言，通过接口实现了 鸭子类型，实际上是 Go 的编译器在其中作了隐匿的转换工作。

### 2、[值接收者和指针接收者的区别](http://golang.design/go-questions/interface/receiver/)

#### 方法

方法能给用户自定义的类型添加新的行为。它和函数的区别在于方法有一个接收者，给一个函数添加一个接收者，那么它就变成了方法。接收者可以是值接收者，也可以是指针接收者。

在调用方法的时候，值类型既可以调用值接收者的方法，也可以调用指针接收者的方法；指针类型既可以调用指针接收者的方法，也可以调用值接收者的方法。

也就是说，不管方法的接收者是什么类型，该类型的值和指针都可以调用，不必严格符合接收者的类型。

实际上，当类型和方法的接收者类型不同时，其实是编译器在背后做了一些工作，用一个表格来呈现：

| **-**          | **值接收者**                                                 | **指针接收者**                                               |
| -------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| 值类型调用者   | 方法会使用调用者的一个副本，类似于“传值”                     | 使用值的引用来调用方法，上例中，qcrao.growUp() 实际上是 (&qcrao).growUp() |
| 指针类型调用者 | 指针被解引用为值，上例中，stefno.howOld() 实际上是 (*stefno).howOld() | 实际上也是“传值”，方法里的操作会影响到调用者，类似于指针传参，拷贝了一份指针 |

#### 值接收者和指针接收者

前面说过，不管接收者类型是值类型还是指针类型，都可以通过值类型或指针类型调用，这里面实际上通过语法糖起作用的。

先说结论：实现了接收者是值类型的方法，相当于自动实现了接收者是指针类型的方法；而实现了接收者是指针类型的方法，不会自动生成对应接收者是值类型的方法。

所以，当实现了一个接收者是值类型的方法，就可以自动生成一个接收者是对应指针类型的方法，因为两者都不会影响接收者。但是，当实现了一个接收者是指针类型的方法，如果此时自动生成一个接收者是值类型的方法，原本期望对接收者的改变（通过指针实现），现在无法实现，因为值类型会产生一个拷贝，不会真正影响调用者。

最后，只要记住下面这点就可以了：

如果实现了接收者是值类型的方法，会隐含地也实现了接收者是指针类型的方法。

#### 两者分别在何时使用

如果方法的接收者是值类型，无论调用者是对象还是对象指针，修改的都是对象的副本，不影响调用者；如果方法的接收者是指针类型，则调用者修改的是指针指向的对象本身。

使用指针作为方法的接收者的理由：

- 方法能够修改接收者指向的值。
- 避免在每次调用方法时复制该值，在值的类型为大型结构体时，这样做会更加高效。

是使用值接收者还是指针接收者，不是由该方法是否修改了调用者（也就是接收者）来决定，而是应该基于该类型的本质。

如果类型具备“原始的本质”，也就是说它的成员都是由 Go 语言里内置的原始类型，如字符串，整型值等，那就定义值接收者类型的方法。像内置的引用类型，如 slice，map，interface，channel，这些类型比较特殊，声明他们的时候，实际上是创建了一个 header， 对于他们也是直接定义值接收者类型的方法。这样，调用函数时，是直接 copy 了这些类型的 header，而 header 本身就是为复制设计的。

如果类型具备非原始的本质，不能被安全地复制，这种类型总是应该被共享，那就定义指针接收者的方法。比如 go 源码里的文件结构体（struct File）就不应该被复制，应该只有一份实体。

### 3、[iface 和 eface 的区别是什么](http://golang.design/go-questions/interface/iface-eface/)

iface 和 eface 都是 Go 中描述接口的底层结构体，区别在于 iface 描述的接口包含方法，而 eface 则是不包含任何方法的空接口：interface{}。

从源码层面看一下：

```plain
type iface struct {
    tab  *itab
    data unsafe.Pointer
}

type itab struct {
    inter  *interfacetype
    _type  *_type
    link   *itab
    hash   uint32 // copy of _type.hash. Used for type switches.
    bad    bool   // type does not implement interface
    inhash bool   // has this itab been added to hash?
    unused [2]byte
    fun    [1]uintptr // variable sized
}
```

iface 内部维护两个指针，tab 指向一个 itab 实体， 它表示接口的类型以及赋给这个接口的实体类型。data 则指向接口具体的值，一般而言是一个指向堆内存的指针。

再来仔细看一下 itab 结构体：_type 字段描述了实体的类型，包括内存对齐方式，大小等；inter 字段则描述了接口的类型。fun 字段放置和接口方法对应的具体数据类型的方法地址，实现接口调用方法的动态分派，一般在每次给接口赋值发生转换时会更新此表，或者直接拿缓存的 itab。

这里只会列出实体类型和接口相关的方法，实体类型的其他方法并不会出现在这里。

另外，你可能会觉得奇怪，为什么 fun 数组的大小为 1，要是接口定义了多个方法可怎么办？实际上，这里存储的是第一个方法的函数指针，如果有更多的方法，在它之后的内存空间里继续存储。从汇编角度来看，通过增加地址就能获取到这些函数指针，没什么影响。顺便提一句，这些方法是按照函数名称的字典序进行排列的。

再看一下 interfacetype 类型，它描述的是接口的类型：

```plain
type interfacetype struct {
    typ     _type
    pkgpath name
    mhdr    []imethod
}
```

可以看到，它包装了 _type 类型，_type 实际上是描述 Go 语言中各种数据类型的结构体。我们注意到，这里还包含一个 mhdr 字段，表示接口所定义的函数列表， pkgpath 记录定义了接口的包名。

这里通过一张图来看下 iface 结构体的全貌：

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1671113733638-8e2e9037-11a8-49af-8dd3-dfd37d7f5d21.png)

接着来看一下 eface 的源码：

```plain
type eface struct {
    _type *_type
    data  unsafe.Pointer
}
```

相比 iface，eface 就比较简单了。只维护了一个 _type 字段，表示空接口所承载的具体的实体类型。data 描述了具体的值。

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1671113735267-6bcdb7c8-dd73-432c-b933-d218fc1b7480.png)

### 4、[接口的动态类型和动态值](http://golang.design/go-questions/interface/dynamic-typing/)

从源码里可以看到：iface包含两个字段：tab 是接口表指针，指向类型信息；data 是数据指针，则指向具体的数据。它们分别被称为动态类型和动态值。而接口值包括动态类型和动态值。

【引申1】接口类型和 nil 作比较

接口值的零值是指动态类型和动态值都为 nil。当仅且当这两部分的值都为 nil 的情况下，这个接口值就才会被认为 接口值 == nil。

### 5、[编译器自动检测类型是否实现接口](http://golang.design/go-questions/interface/detect-impl/)

### 6、[接口的构造过程是怎样的](http://golang.design/go-questions/interface/construct/)

### 7、[类型转换和断言的区别](http://golang.design/go-questions/interface/assert/)

我们知道，Go 语言中不允许隐式类型转换，也就是说 = 两边，不允许出现类型不相同的变量。

类型转换、类型断言本质都是把一个类型转换成另外一个类型。不同之处在于，类型断言是对接口变量进行的操作。

#### **类型转换**

对于类型转换而言，转换前后的两个类型要相互兼容才行。类型转换的语法为：

<结果类型> := <目标类型> ( <表达式> )

```plain
func main() {
    var i int = 9

    var f float64
    f = float64(i)
    fmt.Printf("%T, %v\n", f, f)

    f = 10.8
    a := int(f)
    fmt.Printf("%T, %v\n", a, a)
}
```

#### 断言

前面说过，因为空接口 interface{} 没有定义任何函数，因此 Go 中所有类型都实现了空接口。当一个函数的形参是 interface{}，那么在函数中，需要对形参进行断言，从而得到它的真实类型。

断言的语法为：

<目标类型的值>，<布尔参数> := <表达式>.( 目标类型 ) // 安全类型断言 

<目标类型的值> := <表达式>.( 目标类型 ) //非安全类型断言

类型转换和类型断言有些相似，不同之处，在于类型断言是对接口进行的操作。

```plain
type Student struct {
    Name string
    Age int
}

func main() {
    var i interface{} = new(Student)
    s, ok := i.(Student)
    if ok {
        fmt.Println(s)
    }
}
```

断言其实还有另一种形式，就是用在利用 switch 语句判断接口的类型。每一个 case 会被顺序地考虑。当命中一个 case 时，就会执行 case 中的语句，因此 case 语句的顺序是很重要的，因为很有可能会有多个 case 匹配的情况。

### 8、[接口转换的原理](http://golang.design/go-questions/interface/convert/)

通过前面提到的 iface 的源码可以看到，实际上它包含接口的类型 interfacetype 和 实体类型的类型 _type，这两者都是 iface 的字段 itab 的成员。也就是说生成一个 itab 同时需要接口的类型和实体的类型。

<interface 类型， 实体类型> ->itable

当判定一种类型是否满足某个接口时，Go 使用类型的方法集和接口所需要的方法集进行匹配，如果类型的方法集完全包含接口的方法集，则可认为该类型实现了该接口。

例如某类型有 m 个方法，某接口有 n 个方法，则很容易知道这种判定的时间复杂度为 O(mn)，Go 会对方法集的函数按照函数名的字典序进行排序，所以实际的时间复杂度为 O(m+n)。

这里我们来探索将一个接口转换给另外一个接口背后的原理，当然，能转换的原因必然是类型兼容。

1. 具体类型转空接口时，_type 字段直接复制源类型的 _type；调用 mallocgc 获得一块新内存，把值复制进去，data 再指向这块新内存。
2. 具体类型转非空接口时，入参 tab 是编译器在编译阶段预先生成好的，新接口 tab 字段直接指向入参 tab 指向的 itab；调用 mallocgc 获得一块新内存，把值复制进去，data 再指向这块新内存。
3. 而对于接口转接口，itab 调用 getitab 函数获取。只用生成一次，之后直接从 hash 表中获取。

### 9、[如何用 interface 实现多态](http://golang.design/go-questions/interface/polymorphism/)

Go 语言并没有设计诸如虚函数、纯虚函数、继承、多重继承等概念，但它通过接口却非常优雅地支持了面向对象的特性。

多态是一种运行期的行为，它有以下几个特点：

1. 一种类型具有多种类型的能力
2. 允许不同的对象对同一消息做出灵活的反应
3. 以一种通用的方式对待个使用的对象
4. 非动态语言必须通过继承和接口的方式来实现

main 函数里先生成 Student 和 Programmer 的对象，再将它们分别传入到函数 whatJob 和 growUp。函数中，直接调用接口函数，实际执行的时候是看最终传入的实体类型是什么，调用的是实体类型实现的函数。于是，不同对象针对同一消息就有多种表现，多态就实现了。

### 10、[Go 接口与 C++ 接口有何异同](http://golang.design/go-questions/interface/compare-to-cpp/)

接口定义了一种规范，描述了类的行为和功能，而不做具体实现。

C++ 的接口是使用抽象类来实现的，如果类中至少有一个函数被声明为纯虚函数，则这个类就是抽象类。纯虚函数是通过在声明中使用 “= 0” 来指定的。例如：

```plain
class Shape
{
   public:
      // 纯虚函数
      virtual double getArea() = 0;
   private:
      string name;      // 名称
};
```

设计抽象类的目的，是为了给其他类提供一个可以继承的适当的基类。抽象类不能被用于实例化对象，它只能作为接口使用。

派生类需要明确地声明它继承自基类，并且需要实现基类中所有的纯虚函数。

C++ 定义接口的方式称为“侵入式”，而 Go 采用的是 “非侵入式”，不需要显式声明，只需要实现接口定义的函数，编译器自动会识别。

C++ 和 Go 在定义接口方式上的不同，也导致了底层实现上的不同。C++ 通过虚函数表来实现基类调用派生类的函数；而 Go 通过 itab 中的 fun 字段来实现接口变量调用实体类型的函数。C++ 中的虚函数表是在编译期生成的；而 Go 的 itab 中的 fun 字段是在运行期间动态生成的。原因在于，Go 中实体类型可能会无意中实现 N 多接口，很多接口并不是本来需要的，所以不能为类型实现的所有接口都生成一个 itab， 这也是“非侵入式”带来的影响；这在 C++ 中是不存在的，因为派生需要显示声明它继承自哪个基类。



## 五**、context相关**

https://www.topgoer.cn/docs/gozhuanjia/chapter055.3-context

### **1、context 结构是什么样的？context 使用场景和用途？**

**（难，也常常问你项目中怎么用，光靠记答案很难让面试官满意，反正有各种结合实际的问题）**

**参考链接：**

[go context详解 - 卷毛狒狒 - 博客园www.cnblogs.com/juanmaofeifei/p/14439957.html](https://link.zhihu.com/?target=https%3A//www.cnblogs.com/juanmaofeifei/p/14439957.html)

答：Go 的 Context 的数据结构包含 Deadline，Done，Err，Value，Deadline 方法返回一个 time.Time，表示当前 Context 应该结束的时间，ok 则表示有结束时间，Done 方法当 Context 被取消或者超时时候返回的一个 close 的 channel，告诉给 context 相关的函数要停止当前工作然后返回了，Err 表示 context 被取消的原因，Value 方法表示 context 实现共享数据存储的地方，是协程安全的。context 在业务中是经常被使用的，

**其主要的应用 ：**

1：上下文控制，2：多个 goroutine 之间的数据交互等，3：超时控制：到某个时间点超时，过多久超时。

## **六、channel相关**

### **1、channel 是否线程安全？锁用在什么地方？**

1. Golang的Channel,发送一个数据到Channel 和 从Channel接收一个数据 都是 原子性的。
2. 而且Go的设计思想就是:不要通过共享内存来通信，而是通过通信来共享内存，前者就是传统的加锁，后者就是Channel。
3. 也就是说，设计Channel的主要目的就是在多任务间传递数据的，这当然是安全的

### **2、go channel 的底层实现原理 （数据结构）**

https://juejin.cn/post/7037656471210819614

https://www.topgoer.cn/docs/gozhuanjia/gochan4

#### 数据结构

```go
type hchan struct {
    //channel分为无缓冲和有缓冲两种。
    //对于有缓冲的channel存储数据，借助的是如下循环数组的结构
    qcount   uint           // 循环数组中的元素数量
    dataqsiz uint           // 循环数组的长度
    buf      unsafe.Pointer // 指向底层循环数组的指针
    elemsize uint16 //能够收发元素的大小
    
    
    closed   uint32   //channel是否关闭的标志
    elemtype *_type //channel中的元素类型
    
    //有缓冲channel内的缓冲数组会被作为一个“环型”来使用。
    //当下标超过数组容量后会回到第一个位置，所以需要有两个字段记录当前读和写的下标位置
    sendx    uint   // 下一次发送数据的下标位置
    recvx    uint   // 下一次读取数据的下标位置
    
    //当循环数组中没有数据时，收到了接收请求，那么接收数据的变量地址将会写入读等待队列
    //当循环数组中数据已满时，收到了发送请求，那么发送数据的变量地址将写入写等待队列
    recvq    waitq  // 读等待队列
    sendq    waitq  // 写等待队列
    
    
    lock mutex //互斥锁，保证读写channel时不存在并发竞争问题
}
```

![img](https://cdn.nlark.com/yuque/0/2022/webp/22219483/1661787750459-2608e3a8-f5f9-4d1c-a97f-314d4d83fecf.webp)

总结hchan结构体的主要组成部分有四个：

- 用来保存goroutine之间传递数据的循环链表。=====> buf。
- 用来记录此循环链表当前发送或接收数据的下标值。=====> sendx和recvx。
- 用于保存向该chan发送和从改chan接收数据的goroutine的队列。=====> sendq 和 recvq
- 保证channel写入和读取数据时线程安全的锁。 =====> lock



### **3、nil、关闭的 channel、有数据的 channel，再进行读、写、关闭会怎么样？（各类变种题型，重要）**

#### Channel读写特性(15字口诀)

首先，我们先复习一下Channel都有哪些特性？

- 给一个 nil channel 发送数据，造成永远阻塞
- 从一个 nil channel 接收数据，造成永远阻塞
- 给一个已经关闭的 channel 发送数据，引起 panic
- 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值
- 无缓冲的channel是同步的，而有缓冲的channel是非同步的

以上5个特性是死东西，也可以通过口诀来记忆：“空读写阻塞，写关闭异常，读关闭空零”。

### **4、向 channel 发送数据和从 channel 读数据的流程是什么样的？**

#### 发送流程：

向一个channel中写数据简单过程如下：

1. 如果等待接收队列recvq不为空，说明缓冲区中没有数据或者没有缓冲区，此时直接从recvq取出G,并把数据写入，最后把该G唤醒，结束发送过程；
2. 如果缓冲区中有空余位置，将数据写入缓冲区，结束发送过程；
3. 如果缓冲区中没有空余位置，将待发送数据写入G，将当前G加入sendq，进入睡眠，等待被读goroutine唤醒；

简单流程图如下：

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661788117541-f82a3d7e-8b22-46cd-9bd9-dde26f0d290c.png)

#### 接收流程：

从一个channel读数据简单过程如下：

1. 如果等待发送队列sendq不为空，且没有缓冲区，直接从sendq中取出G，把G中数据读出，最后把G唤醒，结束读取过程；
2. 如果等待发送队列sendq不为空，此时说明缓冲区已满，从缓冲区中首部读出数据，把G中数据写入缓冲区尾部，把G唤醒，结束读取过程；
3. 如果缓冲区中有数据，则从缓冲区取出数据，结束读取过程；
4. 将当前goroutine加入recvq，进入睡眠，等待被写goroutine唤醒；

简单流程图如下：

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661788153163-c386fedf-84b2-42ed-9965-d5d80743650c.png)

#### 关闭channel

关闭channel时会把recvq中的G全部唤醒，本该写入G的数据位置为nil。把sendq中的G全部唤醒，但这些G会panic。

除此之外，panic出现的常见场景还有：

1. 关闭值为nil的channel
2. 关闭已经被关闭的channel
3. 向已经关闭的channel写数据

### **5、讲讲 Go 的 chan 底层数据结构和主要使用场景**

答：channel 的数据结构包含 qccount 当前队列中剩余元素个数，dataqsiz 环形队列长度，即可以存放的元素个数，buf 环形队列指针，elemsize 每个元素的大小，closed 标识关闭状态，elemtype 元素类型，sendx 队列下表，指示元素写入时存放到队列中的位置，recv 队列下表，指示元素从队列的该位置读出。recvq 等待读消息的 goroutine 队列，sendq 等待写消息的 goroutine 队列，lock 互斥锁，chan 不允许并发读写。

**无缓冲和有缓冲区别：** 管道没有缓冲区，从管道读数据会阻塞，直到有协程向管道中写入数据。同样，向管道写入数据也会阻塞，直到有协程从管道读取数据。管道有缓冲区但缓冲区没有数据，从管道读取数据也会阻塞，直到协程写入数据，如果管道满了，写数据也会阻塞，直到协程从缓冲区读取数据。

**channel 的一些特点** 1）、读写值 nil 管道会永久阻塞 2）、关闭的管道读数据仍然可以读数据 3）、往关闭的管道写数据会 panic 4）、关闭为 nil 的管道 panic 5）、关闭已经关闭的管道 panic

**向 channel 写数据的流程：** 如果等待接收队列 recvq 不为空，说明缓冲区中没有数据或者没有缓冲区，此时直接从 recvq 取出 G,并把数据写入，最后把该 G 唤醒，结束发送过程； 如果缓冲区中有空余位置，将数据写入缓冲区，结束发送过程； 如果缓冲区中没有空余位置，将待发送数据写入 G，将当前 G 加入 sendq，进入睡眠，等待被读 goroutine 唤醒；

**向 channel 读数据的流程：** 如果等待发送队列 sendq 不为空，且没有缓冲区，直接从 sendq 中取出 G，把 G 中数据读出，最后把 G 唤醒，结束读取过程； 如果等待发送队列 sendq 不为空，此时说明缓冲区已满，从缓冲区中首部读出数据，把 G 中数据写入缓冲区尾部，把 G 唤醒，结束读取过程； 如果缓冲区中有数据，则从缓冲区取出数据，结束读取过程；将当前 goroutine 加入 recvq，进入睡眠，等待被写 goroutine 唤醒；

**使用场景：** 消息传递、消息过滤，信号广播，事件订阅与广播，请求、响应转发，任务分发，结果汇总，并发控制，限流，同步与异步

### 6、有缓存channel和无缓存channel

https://zhuanlan.zhihu.com/p/355487940

无缓存channel适用于数据要求同步的场景，而有缓存channel适用于无数据同步的场景。可以根据实现项目需求选择。

## **七、GMP相关**

- https://www.topgoer.cn/docs/golangxiuyang/golangxiuyang-1cmeduvk27bo0
- https://www.mubucm.com/doc/7pukUL_nuCI-刘超

### 0、进程、线程、协程有什么区别？（必问）

进程：是应用程序的启动实例，每个进程都有独立的内存空间，不同的进程通过进程间的通信方式来通信。

线程：从属于进程，每个进程至少包含一个线程，线程是 CPU 调度的基本单位，多个线程之间可以共享进程的资源并通过共享内存等线程间的通信方式来通信。

协程：为轻量级线程，与线程相比，协程不受操作系统的调度，协程的调度器由用户应用程序提供，协程调度器按照调度策略把协程调度到线程中运行

### 1、什么是 GMP？（必问）

答：G 代表着 goroutine，P 代表着上下文处理器，M 代表 thread 线程，

在 GPM 模型，有一个全局队列（Global Queue）：存放等待运行的 G，还有一个 P 的本地队列：也是存放等待运行的 G，但数量有限，不超过 256 个。

GPM 的调度流程从 go func()开始创建一个 goroutine，新建的 goroutine 优先保存在 P 的本地队列中，如果 P 的本地队列已经满了，则会保存到全局队列中。

M 会从 P 的队列中取一个可执行状态的 G 来执行，如果 P 的本地队列为空，就会从其他的 MP 组合偷取一个可执行的 G 来执行，

当 M 执行某一个 G 时候发生系统调用或者阻塞，M 阻塞，

如果这个时候 G 在执行，runtime 会把这个线程 M 从 P 中摘除，然后创建一个新的操作系统线程来服务于这个 P，当 M 系统调用结束时，这个 G 会尝试获取一个空闲的 P 来执行，并放入到这个 P 的本地队列，如果这个线程 M 变成休眠状态，加入到空闲线程中，然后整个 G 就会被放入到全局队列中。

 **G,P,M 的个数问题**：

1. G 的个数理论上是无限制的，但是受内存限制，
2. P 的数量一般建议是逻辑 CPU 数量的 2 倍，

1. 1. 由启动时环境变量$GOMAXPROCS或者是由runtime的方法GOMAXPROCS()决定。这意味着在程序执行的任意时刻都只有$GOMAXPROCS个goroutine在同时运行。

1. M 的数量

1. 1. go语言本身的限制：go程序启动时，会设置M的最大数量，默认10000.但是内核很难支持这么多的线程数，所以这个限制可以忽略。
   2. runtime/debug中的SetMaxThreads函数，设置M的最大数量
   3. 一个M阻塞了，会创建新的M。

1. M与P的数量没有绝对关系，一个M阻塞，P就会去创建或者切换另一个M，所以，即使P的默认数量是1，也有可能会创建很多个M出来。

**work stealing（工作量窃取） 机制**：会优先从全局队列里进行窃取，之后会从其它的P队列里窃取一半的G，放入到本地P队列里。
**hand off （移交）机制**：当前线程的G进行阻塞调用时，例如睡眠，则当前线程就会释放P，然后把P转交给其它空闲的线程执行，如果没有闲置的线程，则创建新的线程

### 2、[为什么要有 P？](https://segmentfault.com/a/1190000040092613)

**带来什么改变**

加了 P 之后会带来什么改变呢？我们再更显式的讲一下。

- 每个 P 有自己的本地队列，大幅度的减轻了对全局队列的直接依赖，所带来的效果就是锁竞争的减少。而 GM 模型的性能开销大头就是锁竞争。
- 每个 P 相对的平衡上，在 GMP 模型中也实现了 Work Stealing （工作量窃取机制）算法，如果 P 的本地队列为空，则会从全局队列或其他 P 的本地队列中窃取可运行的 G 来运行，减少空转，提高了资源利用率。

**为什么要有 P**

这时候就有小伙伴会疑惑了，如果是想实现本地队列、Work Stealing 算法，那为什么不直接在 M 上加呢，M 也照样可以实现类似的组件。为什么又再加多一个 P 组件？

结合 M（系统线程） 的定位来看，若这么做，有以下问题：

- 一般来讲，M 的数量都会多于 P。像在 Go 中，M 的数量默认是 10000，P 的默认数量的 CPU 核数。另外由于 M 的属性，也就是如果存在系统阻塞调用，阻塞了 M，又不够用的情况下，M 会不断增加。
- M 不断增加的话，如果本地队列挂载在 M 上，那就意味着本地队列也会随之增加。这显然是不合理的，因为本地队列的管理会变得复杂，且 Work Stealing 性能会大幅度下降。
- M 被系统调用阻塞后，我们是期望把他既有未执行的任务分配给其他继续运行的，而不是一阻塞就导致全部停止。

因此使用 M 是不合理的，那么引入新的组件 P，把本地队列关联到 P 上，就能很好的解决这个问题。

### 3、调度器的设计策略

**复用线程**：避免频繁的创建、销毁线程，而是对线程的复用。

1）work stealing（工作量窃取）机制

当本线程无可运行的G时，尝试从其他线程绑定的P偷取G，而不是销毁线程。

2）hand off（移交）机制

当本线程因为G进行系统调用阻塞时，线程释放绑定的P，把P转移给其他空闲的线程执行。

**利用并行**：GOMAXPROCS设置P的数量，最多有GOMAXPROCS个线程分布在多个CPU上同时运行。GOMAXPROCS也限制了并发的程度，比如GOMAXPROCS = 核数/2，则最多利用了一半的CPU核进行并行。

**抢占**：在coroutine中要等待一个协程主动让出CPU才执行下一个协程，在Go中，一个goroutine最多占用CPU 10ms，防止其他goroutine被饿死，这就是goroutine不同于coroutine的一个地方。

**全局G队列**：在新的调度器中依然有全局G队列，但功能已经被弱化了，当M执行work stealing从其他P偷不到G时，它可以从全局G队列获取G。

### **3、抢占式调度是如何抢占的？**

**基于协作式抢占**

**基于信号量抢占**

就像操作系统要负责线程的调度一样，Go的runtime要负责goroutine的调度。现代操作系统调度线程都是抢占式的，我们不能依赖用户代码主动让出CPU，或者因为IO、锁等待而让出，这样会造成调度的不公平。基于经典的时间片算法，当线程的时间片用完之后，会被时钟中断给打断，调度器会将当前线程的执行上下文进行保存，然后恢复下一个线程的上下文，分配新的时间片令其开始执行。这种抢占对于线程本身是无感知的，系统底层支持，不需要开发人员特殊处理。

基于时间片的抢占式调度有个明显的优点，能够避免CPU资源持续被少数线程占用，从而使其他线程长时间处于饥饿状态。goroutine的调度器也用到了时间片算法，但是和操作系统的线程调度还是有些区别的，因为整个Go程序都是运行在用户态的，所以不能像操作系统那样利用时钟中断来打断运行中的goroutine。也得益于完全在用户态实现，goroutine的调度切换更加轻量。

**上面这两段文字只是对调度的一个概括，具体的协作式调度、信号量调度大家还需要去详细了解，这偏底层了，大厂或者中高级开发会问。（字节就问了）**

### 4、调度器的生命周期

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1671108479128-e538cce4-0911-4683-ba0a-8a7866e4e2c1.png)

特殊的M0和G0

#### M0

M0是启动程序后的编号为0的主线程，这个M对应的实例会在全局变量runtime.m0中，不需要在heap上分配，M0负责执行初始化操作和启动第一个G， 在之后M0就和其他的M一样了。

#### G0

G0是每次启动一个M都会第一个创建的gourtine，G0仅用于负责调度的G，G0不指向任何可执行的函数, 每个M都会有一个自己的G0。在调度或系统调用时会使用G0的栈空间, 全局变量的G0是M0的G0。

我们来跟踪一段代码

```go
package main 
import "fmt" 
func main() {
    fmt.Println("Hello world") 
}
```

接下来我们来针对上面的代码对调度器里面的结构做一个分析。

也会经历如上图所示的过程：

1. runtime创建最初的线程m0和goroutine g0，并把2者关联。
2. 调度器初始化：初始化m0、栈、垃圾回收，以及创建和初始化由GOMAXPROCS个P构成的P列表。
3. 示例代码中的main函数是main.main，runtime中也有1个main函数——runtime.main，代码经过编译后，runtime.main会调用main.main，程序启动时会为runtime.main创建goroutine，称它为main goroutine吧，然后把main goroutine加入到P的本地队列。
4. 启动m0，m0已经绑定了P，会从P的本地队列获取G，获取到main goroutine。
5. G拥有栈，M根据G中的栈信息和调度信息设置运行环境
6. M运行G
7. G退出，再次回到M获取可运行的G，这样重复下去，直到main.main退出，runtime.main执行Defer和Panic处理，或调用runtime.exit退出程序。

调度器的生命周期几乎占满了一个Go程序的一生，runtime.main的goroutine执行之前都是为调度器做准备工作，runtime.main的goroutine运行，才是调度器的真正开始，直到runtime.main结束而结束。

## 八、锁相关

https://www.topgoer.cn/docs/gozhuanjia/gozhuanjiamutex

https://www.topgoer.cn/docs/gozhuanjia/gozhuanjiarwmutex

### 1、除了 mutex 以外还有那些方式安全读写共享变量？

\* 将共享变量的读写放到一个 goroutine 中，其它 goroutine 通过 channel 进行读写操作。

\* 可以用个数为 1 的信号量（semaphore）实现互斥

\* 通过 Mutex 锁实现

### 2、Go 如何实现原子操作？

答：原子操作就是不可中断的操作，外界是看不到原子操作的中间状态，要么看到原子操作已经完成，要么看到原子操作已经结束。在某个值的原子操作执行的过程中，CPU 绝对不会再去执行其他针对该值的操作，那么其他操作也是原子操作。

Go 语言的标准库代码包 sync/atomic 提供了原子的读取（Load 为前缀的函数）或写入（Store 为前缀的函数）某个值（这里细节还要多去查查资料）。

**原子操作与互斥锁的区别**

1）、互斥锁是一种数据结构，用来让一个线程执行程序的关键部分，完成互斥的多个操作。

2）、原子操作是针对某个值的单个互斥操作。

### 3、Mutex 是悲观锁还是乐观锁？悲观锁、乐观锁是什么？

**悲观锁**

悲观锁：当要对数据库中的一条数据进行修改的时候，为了避免同时被其他人修改，最好的办法就是直接对该数据进行加锁以防止并发。这种借助数据库锁机制，在修改数据之前先锁定，再修改的方式被称之为悲观并发控制【Pessimistic Concurrency Control，缩写“PCC”，又名“悲观锁”】。

**乐观锁**

乐观锁是相对悲观锁而言的，乐观锁假设数据一般情况不会造成冲突，所以在数据进行提交更新的时候，才会正式对数据的冲突与否进行检测，如果冲突，则返回给用户异常信息，让用户决定如何去做。乐观锁适用于读多写少的场景，这样可以提高程序的吞吐量

### 4、Mutex 有几种模式？

**1）正常模式**

1. 当前的mutex只有一个goruntine来获取，那么没有竞争，直接返回。
2. 新的goruntine进来，如果当前mutex已经被获取了，则该goruntine进入一个先入先出的waiter队列，在mutex被释放后，waiter按照先进先出的方式获取锁。该goruntine会处于自旋状态(不挂起，继续占有cpu)。
3. 新的goruntine进来，mutex处于空闲状态，将参与竞争。新来的 goroutine 有先天的优势，它们正在 CPU 中运行，可能它们的数量还不少，所以，在高并发情况下，被唤醒的 waiter 可能比较悲剧地获取不到锁，这时，它会被插入到队列的前面。如果 waiter 获取不到锁的时间超过阈值 1 毫秒，那么，这个 Mutex 就进入到了饥饿模式。

**2）饥饿模式**

在饥饿模式下，Mutex 的拥有者将直接把锁交给队列最前面的 waiter。新来的 goroutine 不会尝试获取锁，即使看起来锁没有被持有，它也不会去抢，也不会 spin（自旋），它会乖乖地加入到等待队列的尾部。 如果拥有 Mutex 的 waiter 发现下面两种情况的其中之一，它就会把这个 Mutex 转换成正常模式:

1. 此 waiter 已经是队列中的最后一个 waiter 了，没有其它的等待锁的 goroutine 了；
2. 此 waiter 的等待时间小于 1 毫秒。

### 5、goroutine 的自旋占用资源如何解决

自旋锁是指当一个线程在获取锁的时候，如果锁已经被其他线程获取，那么该线程将循环等待，然后不断地判断是否能够被成功获取，直到获取到锁才会退出循环。

**自旋的条件如下：**

1）还没自旋超过 4 次,

2）多核处理器，

3）GOMAXPROCS > 1，

4）p 上本地 goroutine 队列为空。

mutex 会让当前的 goroutine 去空转 CPU，在空转完后再次调用 CAS 方法去尝试性的占有锁资源，直到不满足自旋条件，则最终会加入到等待队列里。

## **九、并发相关**

### 0、讲讲 Go 中主协程如何等待其余协程退出?

答：Go 的 sync.WaitGroup 是等待一组协程结束，sync.WaitGroup 只有 3 个方法，Add()是添加计数，Done()减去一个计数，Wait()阻塞直到所有的任务完成。Go 里面还能通过有缓冲的 channel 实现其阻塞等待一组协程结束，这个不能保证一组 goroutine 按照顺序执行，可以并发执行协程。Go 里面能通过无缓冲的 channel 实现其阻塞等待一组协程结束，这个能保证一组 goroutine 按照顺序执行，但是不能并发执行。

**啰嗦一句：**循环智能二面，手写代码部分时，三个协程按交替顺序打印数字，最后题目做出来了，问我代码中Add()是什么意思，我回答的不是很清晰，这家公司就没有然后了。Add()表示协程计数，可以一次Add多个，如Add(3),可以多次Add(1);然后每个子协程必须调用done（）,这样才能保证所有子协程结束，主协程才能结束。

### 1、怎么控制并发数？

**第一，有缓冲通道**

根据通道中没有数据时读取操作陷入阻塞和通道已满时继续写入操作陷入阻塞的特性，正好实现控制并发数量。

```go
func main() {
    count := 10                     // 最大支持并发
    sum := 100                      // 任务总数
    wg := sync.WaitGroup{}          //控制主协程等待所有子协程执行完之后再退出。
    c := make(chan struct{}, count) // 控制任务并发的chan
    defer close(c)
    for i := 0; i < sum; i++ {
        wg.Add(1)
        c <- struct{}{} // 作用类似于waitgroup.Add(1)
        go func(j int) {
            defer wg.Done()
            fmt.Println(j)
            <-c // 执行完毕，释放资源
        }(i)
    }
    wg.Wait()
}
```

**第二，三方库实现的协程池**

```go
import (
    "github.com/Jeffail/tunny"
    "log"
    "time"
)
func main() {
    pool := tunny.NewFunc(10, func(i interface{}) interface{} {
        log.Println(i)
        time.Sleep(time.Second)
        return nil
    })
    defer pool.Close()
    for i := 0; i < 500; i++ {
        go pool.Process(i)
    }
    time.Sleep(time.Second * 4)
}
```

### 2、多个 goroutine 对同一个 map 写会 panic，异常是否可以用 defer 捕获？

可以捕获异常，但是只能捕获一次，Go语言，可以使用多值返回来返回错误。不要用异常代替错误，更不要用来控制流程。在极个别的情况下，才使用Go中引入的Exception处理：defer, panic, recover Go中，对异常处理的原则是：多用error包，少用panic

```go
defer func() {
    if err := recover(); err != nil {
        // 打印异常，关闭资源，退出此函数
        fmt.Println(err)
    }
}()
```

### 3、如何优雅的实现一个 goroutine 池

（百度、手写代码，本人面传音控股被问道：请求数大于消费能力怎么设计协程池）

这一块能啃下来，offer满天飞，这应该是保证高并发系统稳定性、高可用的核心部分之一。

**建议参考：**

[Golang学习篇--协程池_Word哥的博客-CSDN博客_golang协程池blog.csdn.net/finghting321/article/details/106492915/](https://link.zhihu.com/?target=https%3A//blog.csdn.net/finghting321/article/details/106492915/)

**这篇文章的目录是：**

\1. 为什么需要协程池？

\2. 简单的协程池

\3. go-playground/pool

\4. ants（推荐）

**所以直接研究ants底层吧，省的造轮子。**

### 4、golang实现多并发请求（发送多个get请求）

在[go语言](https://so.csdn.net/so/search?q=go语言&spm=1001.2101.3001.7020)中其实有两种方法进行协程之间的通信。**一个是共享内存、一个是消息传递**

[**共享内存（互斥锁）**](https://blog.csdn.net/m0_43432638/article/details/108359182)

```go
//基本的GET请求
package main
 
import (
    "fmt"
    "io/ioutil"
    "net/http"
	"time"
	"sync"
	"runtime"
)
 
// 计数器
var counter int = 0
 
func httpget(lock *sync.Mutex){
	lock.Lock()
	counter++
    resp, err := http.Get("http://localhost:8000/rest/api/user")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
    fmt.Println(resp.StatusCode)
    if resp.StatusCode == 200 {
        fmt.Println("ok")
    }
	lock.Unlock()
}
 
func main() {
    start := time.Now()
	lock := &sync.Mutex{}
    for i := 0; i < 800; i++ {
        go httpget(lock)
    }
    for  {
        lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
        if c >= 800 {
            break
        }
    }
    end := time.Now()
    consume := end.Sub(start).Seconds()
    fmt.Println("程序执行耗时(s)：", consume)
}
```

问题

我们可以看到共享内存的方式是可以做到并发，但是我们需要利用共享变量来进行[协程](https://so.csdn.net/so/search?q=协程&spm=1001.2101.3001.7020)的通信，也就需要使用互斥锁来确保数据安全性，导致代码啰嗦，复杂话，不易维护。我们后续使用go的[消息传递](https://blog.csdn.net/m0_43432638/article/details/108349384)方式避免这些问题。

[**消息传递（管道）**](https://blog.csdn.net/m0_43432638/article/details/108349384)

```go
//基本的GET请求
package main
 
import (
    "fmt"
    "io/ioutil"
    "net/http"
	"time"
)
// HTTP get请求
func httpget(ch chan int){
    resp, err := http.Get("http://localhost:8000/rest/api/user")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
    fmt.Println(resp.StatusCode)
    if resp.StatusCode == 200 {
        fmt.Println("ok")
    }
	ch <- 1
}
// 主方法
func main() {
    start := time.Now()
    // 注意设置缓冲区大小要和开启协程的个人相等
    chs := make([]chan int, 2000)
    for i := 0; i < 2000; i++ {
        chs[i] = make(chan int)
        go httpget(chs[i])
    }
    for _, ch := range chs {
        <- ch
    }
    end := time.Now()
    consume := end.Sub(start).Seconds()
    fmt.Println("程序执行耗时(s)：", consume)
}
```

**总结：**

我们通过[go语言](https://so.csdn.net/so/search?q=go语言&spm=1001.2101.3001.7020)的管道channel来实现并发请求，能够解决如何避免传统共享内存实现并发的很多问题而且效率会高于共享内存的方法。 

## **十、GC相关**

https://www.topgoer.cn/docs/gozhuanjia/chapter044.2-garbage_collection

https://www.topgoer.cn/docs/golangxiuyang/golangxiuyang-1cmee076rjgk7

### 1、go gc 是怎么实现的？（必问）

答：

**细分常见的三个问题：1、GC机制随着golang版本变化如何变化的？2、三色标记法的流程？3、插入屏障、删除屏障，混合写屏障（具体的实现比较难描述，但你要知道屏障的作用：避免程序运行过程中，变量被误回收；减少STW的时间）4、虾皮还问了个开放性的题目：你觉得以后GC机制会怎么优化？**

**Go 的 GC 回收有三次演进过程，Go V1.3 之前普通标记清除（mark and sweep）方法，整体过程需要启动 STW，效率极低。GoV1.5 三色标记法，堆空间启动写屏障，栈空间不启动，全部扫描之后，需要重新扫描一次栈(需要 STW)，效率普通。GoV1.8 三色标记法，混合写屏障机制：栈空间不启动（全部标记成黑色），堆空间启用写屏障，整个过程不要 STW，效率高。**

Go1.3 之前的版本所谓**标记清除**是先启动 STW 暂停，然后执行标记，再执行数据回收，最后停止 STW。Go1.3 版本标记清除做了点优化，流程是：先启动 STW 暂停，然后执行标记，停止 STW，最后再执行数据回收。

Go1.5 **三色标记**主要是插入屏障和删除屏障，写入屏障的流程：程序开始，全部标记为白色，1）所有的对象放到白色集合，2）遍历一次根节点，得到灰色节点，3）遍历灰色节点，将可达的对象，从白色标记灰色，遍历之后的灰色标记成黑色，4）由于并发特性，此刻外界向在堆中的对象发生添加对象，以及在栈中的对象添加对象，在堆中的对象会触发插入屏障机制，栈中的对象不触发，5）由于堆中对象插入屏障，则会把堆中黑色对象添加的白色对象改成灰色，栈中的黑色对象添加的白色对象依然是白色，6）循环第 5 步，直到没有灰色节点，7）在准备回收白色前，重新遍历扫描一次栈空间，加上 STW 暂停保护栈，防止外界干扰（有新的白色会被添加成黑色）在 STW 中，将栈中的对象一次三色标记，直到没有灰色，8）停止 STW，清除白色。至于删除写屏障，则是遍历灰色节点的时候出现可达的节点被删除，这个时候触发删除写屏障，这个可达的被删除的节点也是灰色，等循环三色标记之后，直到没有灰色节点，然后清理白色，删除写屏障会造成一个对象即使被删除了最后一个指向它的指针也依旧可以活过这一轮，在下一轮 GC 中被清理掉。

GoV1.8 **混合写屏障**规则是：

1）GC 开始将栈上的对象全部扫描并标记为黑色(之后不再进行第二次重复扫描，无需 STW)，2）GC 期间，任何在栈上创建的新对象，均为黑色。3）被删除的对象标记为灰色。4）被添加的对象标记为灰色。

### 2、go 是 gc 算法是怎么实现的？ （得物，出现频率低）

```go
func GC() {
    n := atomic.Load(&work.cycles)
    gcWaitOnMark(n)
    gcStart(gcTrigger{kind: gcTriggerCycle, n: n + 1})
    gcWaitOnMark(n + 1)
    for atomic.Load(&work.cycles) == n+1 && sweepone() != ^uintptr(0) {
        sweep.nbgsweep++
        Gosched()
    }
    for atomic.Load(&work.cycles) == n+1 && atomic.Load(&mheap_.sweepers) != 0 {
        Gosched()
    }
    mp := acquirem()
    cycle := atomic.Load(&work.cycles)
    if cycle == n+1 || (gcphase == _GCmark && cycle == n+2) {
        mProf_PostSweep()
    }
    releasem(mp)
}
```

底层原理了，可能大厂，中高级才会问，参考：

[Golang GC算法解读_suchy_sz的博客-CSDN博客_go的gc算法blog.csdn.net/shudaqi2010/article/details/90025192](https://link.zhihu.com/?target=https%3A//blog.csdn.net/shudaqi2010/article/details/90025192)

### 3、GC 中 stw 时机，各个阶段是如何解决的？ （百度）

**底层原理，自行百度一下，我等渣渣简历都过不了BAT，字节，虾皮，特使拉以及一些国Q还能收到面试邀约**。

1）在开始新的一轮 GC 周期前，需要调用 gcWaitOnMark 方法上一轮 GC 的标记结束（含扫描终止、标记、或标记终止等）。

2）开始新的一轮 GC 周期，调用 gcStart 方法触发 GC 行为，开始扫描标记阶段。

3）需要调用 gcWaitOnMark 方法等待，直到当前 GC 周期的扫描、标记、标记终止完成。

4）需要调用 sweepone 方法，扫描未扫除的堆跨度，并持续扫除，保证清理完成。在等待扫除完毕前的阻塞时间，会调用 Gosched 让出。

5）在本轮 GC 已经基本完成后，会调用 mProf_PostSweep 方法。以此记录最后一次标记终止时的堆配置文件快照。

6）结束，释放 M。

### 4、GC 的触发时机？

初级必问，分为系统触发和主动触发。

1）gcTriggerHeap：当所分配的堆大小达到阈值（由控制器计算的触发堆的大小）时，将会触发。

2）gcTriggerTime：当距离上一个 GC 周期的时间超过一定时间时，将会触发。时间周期以runtime.forcegcperiod 变量为准，默认 2 分钟。

3）gcTriggerCycle：如果没有开启 GC，则启动 GC。

4）手动触发的 runtime.GC 方法。

## **十一、内存相关**

### [内存分配原理](https://www.topgoer.cn/docs/gozhuanjia/gozhuanjiachapter044.1-memory_alloc)

### [垃圾回收原理](https://www.topgoer.cn/docs/gozhuanjia/chapter044.2-garbage_collection)

### [逃逸分析](https://www.topgoer.cn/docs/gozhuanjia/chapter044.3-escape_analysis)

### [Go语言的内存模型及堆的分配管理](https://zhuanlan.zhihu.com/p/76802887)

### 1、谈谈内存泄露，什么情况下内存会泄露？怎么定位排查内存泄漏问题？

答：go 中的内存泄漏一般都是 goroutine 泄漏，就是 goroutine 没有被关闭，或者没有添加超时控制，让 goroutine 一只处于阻塞状态，不能被 GC。

**内存泄露有下面一些情况**

1）如果 goroutine 在执行时被阻塞而无法退出，就会导致 goroutine 的内存泄漏，一个 goroutine 的最低栈大小为 2KB，在高并发的场景下，对内存的消耗也是非常恐怖的。

2）互斥锁未释放或者造成死锁会造成内存泄漏

3）time.Ticker 是每隔指定的时间就会向通道内写数据。作为循环触发器，必须调用 stop 方法才会停止，从而被 GC 掉，否则会一直占用内存空间。

4）字符串的截取引发临时性的内存泄漏

```go
func main() {
    var str0 = "12345678901234567890"
    str1 := str0[:10]
}
```

5）切片截取引起子切片内存泄漏

```go
func main() {
    var s0 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    s1 := s0[:3]
}
```

6）函数数组传参引发内存泄漏【如果我们在函数传参的时候用到了数组传参，且这个数组够大（我们假设数组大小为 100 万，64 位机上消耗的内存约为 800w 字节，即 8MB 内存），或者该函数短时间内被调用 N 次，那么可想而知，会消耗大量内存，对性能产生极大的影响，如果短时间内分配大量内存，而又来不及 GC，那么就会产生临时性的内存泄漏，对于高并发场景相当可怕。】

**排查方式：**

一般通过 pprof 是 Go 的性能分析工具，在程序运行过程中，可以记录程序的运行信息，可以是 CPU 使用情况、内存使用情况、goroutine 运行情况等，当需要性能调优或者定位 Bug 时候，这些记录的信息是相当重要。

**当然你能说说具体的分析指标更加分咯，有的面试官就喜欢他问什么，你简洁的回答什么，不喜欢巴拉巴拉详细解释一通，比如虾P面试官，不过他考察的内容特别多，可能是为了节约时间。**

### 2、golang 的内存逃逸吗？什么情况下会发生内存逃逸？（必问）

答：1)本该分配到栈上的变量，跑到了堆上，这就导致了内存逃逸。2)栈是高地址到低地址，栈上的变量，函数结束后变量会跟着回收掉，不会有额外性能的开销。3)变量从栈逃逸到堆上，如果要回收掉，需要进行 gc，那么 gc 一定会带来额外的性能开销。编程语言不断优化 gc 算法，主要目的都是为了减少 gc 带来的额外性能开销，变量一旦逃逸会导致性能开销变大。

**内存逃逸的情况如下：**

1）方法内返回局部变量指针。

2）向 channel 发送指针数据。

3）在闭包中引用包外的值。

4）在 slice 或 map 中存储指针。

5）切片（扩容后）长度太大。

6）在 interface 类型上调用方法。

### 3、请简述 Go 是如何分配内存的？

mcache mcentral mheap mspan

Go 程序启动的时候申请一大块内存，并且划分 spans，bitmap，areana 区域；arena 区域按照页划分成一个个小块，span 管理一个或者多个页，mcentral 管理多个 span 供现场申请使用；mcache 作为线程私有资源，来源于 mcentral。

**这里描述的比较简单，你可以自己再去搜索下更简洁完整的答案。**

### 4、[go内存分配器](https://zhuanlan.zhihu.com/p/410317967)

### 5、Channel 分配在栈上还是堆上？哪些对象分配在堆上，哪些对象分配在栈上？

Channel 被设计用来实现协程间通信的组件，其作用域和生命周期不可能仅限于某个函数内部，所以 golang 直接将其分配在堆上

准确地说，你并不需要知道。Golang 中的变量只要被引用就一直会存活，存储在堆上还是栈上由内部实现决定而和具体的语法没有关系。

知道变量的存储位置确实和效率编程有关系。如果可能，Golang 编译器会将函数的局部变量分配到函数栈帧（stack frame）上。然而，如果编译器不能确保变量在函数 return 之后不再被引用，编译器就会将变量分配到堆上。而且，如果一个局部变量非常大，那么它也应该被分配到堆上而不是栈上。

当前情况下，如果一个变量被取地址，那么它就有可能被分配到堆上,然而，还要对这些变量做逃逸分析，如果函数 return 之后，变量不再被引用，则将其分配到栈上。

### 6、介绍一下大对象小对象，为什么小对象多了会造成 gc 压力？

小于等于 32k 的对象就是小对象，其它都是大对象。一般小对象通过 mspan 分配内存；大对象则直接由 mheap 分配内存。通常小对象过多会导致 GC 三色法消耗过多的 CPU。优化思路是，减少对象分配。

小对象：如果申请小对象时，发现当前内存空间不存在空闲跨度时，将会需要调用 nextFree 方法获取新的可用的对象，可能会触发 GC 行为。

大对象：如果申请大于 32k 以上的大对象时，可能会触发 GC 行为。

## 十二、编译

### [逃逸分析是怎么进行的](http://golang.design/go-questions/compile/escape/)



在编译原理中，分析指针动态范围的方法称之为逃逸分析。通俗来讲，当一个对象的指针被多个方法或线程引用时，我们称这个指针发生了逃逸。



Go语言的逃逸分析是编译器执行静态代码分析后，对内存管理进行的优化和简化，它可以决定一个变量是分配到堆还栈上。



写过C/C++的同学都知道，调用著名的malloc和new函数可以在堆上分配一块内存，这块内存的使用和销毁的责任都在程序员。一不小心，就会发生内存泄露。



Go语言里，基本不用担心内存泄露了。虽然也有new函数，但是使用new函数得到的内存不一定就在堆上。堆和栈的区别对程序员“模糊化”了，当然这一切都是Go编译器在背后帮我们完成的。



Go语言逃逸分析最基本的原则是：如果一个函数返回对一个变量的引用，那么它就会发生逃逸。



简单来说，编译器会分析代码的特征和代码生命周期，Go中的变量只有在编译器可以证明在函数返回后不会再被引用的，才分配到栈上，其他情况下都是分配到堆上。



Go语言里没有一个关键字或者函数可以直接让变量被编译器分配到堆上，相反，编译器通过分析代码来决定将变量分配到何处。



对一个变量取地址，可能会被分配到堆上。但是编译器进行逃逸分析后，如果考察到在函数返回后，此变量不会被引用，那么还是会被分配到栈上。



编译器会根据变量是否被外部引用来决定是否逃逸：



1. 如果函数外部没有引用，则优先放到栈中；
2. 如果函数外部存在引用，则必定放到堆中；



Go的垃圾回收，让堆和栈对程序员保持透明。真正解放了程序员的双手，让他们可以专注于业务，“高效”地完成代码编写。把那些内存管理的复杂机制交给编译器，而程序员可以去享受生活。



逃逸分析这种“骚操作”把变量合理地分配到它该去的地方。即使你是用new申请到的内存，如果我发现你竟然在退出函数后没有用了，那么就把你丢到栈上，毕竟栈上的内存分配比堆上快很多；反之，即使你表面上只是一个普通的变量，但是经过逃逸分析后发现在退出函数之后还有其他地方在引用，那我就把你分配到堆上。



如果变量都分配到堆上，堆不像栈可以自动清理。它会引起Go频繁地进行垃圾回收，而垃圾回收会占用比较大的系统开销（占用CPU容量的25%）。



堆和栈相比，堆适合不可预知大小的内存分配。但是为此付出的代价是分配速度较慢，而且会形成内存碎片。栈内存分配则会非常快。栈分配内存只需要两个CPU指令：“PUSH”和“RELEASE”，分配和释放；而堆分配内存首先需要去找到一块大小合适的内存块，之后要通过垃圾回收才能释放。



通过逃逸分析，可以尽量把那些不需要分配到堆上的变量直接分配到栈上，堆上的变量少了，会减轻分配堆内存的开销，同时也会减少gc的压力，提高程序的运行速度。



### [GoRoot 和 GoPath 有什么用](http://golang.design/go-questions/compile/gopath/)



GoRoot 是 Go 的安装路径。mac 或 unix 是在 `/usr/local/go` 路径上，来看下这里都装了些什么：



![img](https://golang.design/go-questions/compile/assets/1.png)



bin 目录下面：



![img](https://golang.design/go-questions/compile/assets/2.png)



pkg 目录下面：



![img](https://golang.design/go-questions/compile/assets/3.png)



Go 工具目录如下，其中比较重要的有编译器 `compile`，链接器 `link`：



![img](https://golang.design/go-questions/compile/assets/4.png)



GoPath 的作用在于提供一个可以寻找 `.go` 源码的路径，它是一个工作空间的概念，可以设置多个目录。Go 官方要求，GoPath 下面需要包含三个文件夹：



```go
src
pkg
bin
```



src 存放源文件，pkg 存放源文件编译后的库文件，后缀为 `.a`；bin 则存放可执行文件。



### [Go 编译链接过程概述](http://golang.design/go-questions/compile/link-process/)



Go 程序并不能直接运行，每条 Go 语句必须转化为一系列的低级机器语言指令，将这些指令打包到一起，并以二进制磁盘文件的形式存储起来，也就是可执行目标文件。



从源文件到可执行目标文件的转化过程：



![img](https://golang.design/go-questions/compile/assets/7.png)



完成以上各个阶段的就是 Go 编译系统。你肯定知道大名鼎鼎的 GCC（GNU Compile Collection），中文名为 GNU 编译器套装，它支持像 C，C++，Java，Python，Objective-C，Ada，Fortran，Pascal，能够为很多不同的机器生成机器码。



可执行目标文件可以直接在机器上执行。一般而言，先执行一些初始化的工作；找到 main 函数的入口，执行用户写的代码；执行完成后，main 函数退出；再执行一些收尾的工作，整个过程完毕。



在接下来的文章里，我们将探索`编译`和`运行`的过程。



Go 源码里的编译器源码位于 `src/cmd/compile` 路径下，链接器源码位于 `src/cmd/link` 路径下。



### [Go 编译相关的命令详解](http://golang.design/go-questions/compile/cmd/)



和编译相关的命令主要是：



```go
go build
go install
go run
```



#### go build



`go build` 用来编译指定 packages 里的源码文件以及它们的依赖包，编译的时候会到 `$GoPath/src/package` 路径下寻找源码文件。`go build` 还可以直接编译指定的源码文件，并且可以同时指定多个。



通过执行 `go help build` 命令得到 `go build` 的使用方法：



```go
usage: go build [-o output] [-i] [build flags] [packages]
```



`-o` 只能在编译单个包的时候出现，它指定输出的可执行文件的名字。



`-i` 会安装编译目标所依赖的包，安装是指生成与代码包相对应的 `.a` 文件，即静态库文件（后面要参与链接），并且放置到当前工作区的 pkg 目录下，且库文件的目录层级和源码层级一致。



至于 build flags 参数，`build, clean, get, install, list, run, test` 这些命令会共用一套：

| 参数  | 作用                                                         |
| ----- | ------------------------------------------------------------ |
| -a    | 强制重新编译所有涉及到的包，包括标准库中的代码包，这会重写 /usr/local/go 目录下的 `.a` 文件 |
| -n    | 打印命令执行过程，不真正执行                                 |
| -p n  | 指定编译过程中命令执行的并行数，n 默认为 CPU 核数            |
| -race | 检测并报告程序中的数据竞争问题                               |
| -v    | 打印命令执行过程中所涉及到的代码包名称                       |
| -x    | 打印命令执行过程中所涉及到的命令，并执行                     |
| -work | 打印编译过程中的临时文件夹。通常情况下，编译完成后会被删除   |



我们知道，Go 语言的源码文件分为三类：命令源码、库源码、测试源码。



命令源码文件：是 Go 程序的入口，包含 `func main()` 函数，且第一行用 `package main` 声明属于 main 包。



库源码文件：主要是各种函数、接口等，例如工具类的函数。



测试源码文件：以 `_test.go` 为后缀的文件，用于测试程序的功能和性能。



注意，`go build` 会忽略 `*_test.go` 文件。



#### go install



`go install` 用于编译并安装指定的代码包及它们的依赖包。相比 `go build`，它只是多了一个“安装编译后的结果文件到指定目录”的步骤。



还是使用之前 hello-world 项目的例子，我们先将 pkg 目录删掉，在项目根目录执行：



```go
go install src/main.go

或者

go install util
```



两者都会在根目录下新建一个 `pkg` 目录，并且生成一个 `util.a` 文件。



并且，在执行前者的时候，会在 GOBIN 目录下生成名为 main 的可执行文件。



所以，运行 `go install` 命令，库源码包对应的 `.a` 文件会被放置到 `pkg` 目录下，命令源码包生成的可执行文件会被放到 GOBIN 目录。



`go install` 在 GoPath 有多个目录的时候，会产生一些问题，具体可以去看郝林老师的 `Go 命令教程`，这里不展开了。



#### go run



`go run` 用于编译并运行命令源码文件。

### [Go 程序启动过程是怎样的](http://golang.design/go-questions/compile/booting/)

## 十三、框架

### Gin

文档：https://gin-gonic.com/zh-cn/docs/introduction/

#### 0、特性

1. **快速**

1. 1. 基于 Radix 树的路由，小内存占用。没有反射。可预测的 API 性能。

1. **支持中间件**

1. 1. 传入的 HTTP 请求可以由一系列中间件和最终操作来处理。 例如：Logger，Authorization，GZIP，最终操作 DB。

1. **Crash 处理**

1. 1. Gin 可以 catch 一个发生在 HTTP 请求中的 panic 并 recover 它。这样，你的服务器将始终可用。例如，你可以向 Sentry 报告这个 panic！

1. **JSON 验证**

1. 1. Gin 可以解析并验证请求的 JSON，例如检查所需值的存在。

1. **路由组**

1. 1. 更好地组织路由。是否需要授权，不同的 API 版本…… 此外，这些组可以无限制地嵌套而不会降低性能。

1. **错误管理**

1. 1. Gin 提供了一种方便的方法来收集 HTTP 请求期间发生的所有错误。最终，中间件可以将它们写入日志文件，数据库并通过网络发送。

1. **内置渲染**

1. 1. Gin 为 JSON，XML 和 HTML 渲染提供了易于使用的 API。

1. **可扩展性**

1. 1. 新建一个中间件非常简单，去查看[示例代码](https://gin-gonic.com/zh-cn/docs/examples/using-middleware/)吧。

#### 1、[gin目录结构](https://blog.csdn.net/qq_34877350/article/details/107959381)

文档：https://blog.csdn.net/qq_34877350/article/details/107959381

```plain
├── gin
│   ├──  Router
│          └── router.go
│   ├──  Middlewares
│          └── corsMiddleware.go
│   ├──  Controllers
│          └── testController.go
│   ├──  Services
│          └── testService.go
│   ├──  Models
│          └── testModel.go
│   ├──  Databases
│          └── mysql.go
│   ├──  Sessions
│          └── session.go
└── main.go
```

- 使用gorm访问数据库
- gin 为项目根目录
- main.go 为入口文件
- Router 为路由目录
- Middlewares 为中间件目录
- Controllers 为控制器目录（MVC）
- Services 为服务层目录，这里把DAO逻辑也写入其中，如果分开也可以
- Models 为模型目录
- Databases 为数据库初始化目录
- Sessions 为session初始化目录
- 文件 引用顺序 大致如下：
- main.go(在main中关闭数据库) - router(Middlewares) - Controllers - Services(sessions) - Models - Databases

#### 2、[Gin框架介绍及使用 - 李文周的博客](https://www.liwenzhou.com/posts/Go/Gin_framework/#autoid-0-0-0)

文档：https://www.liwenzhou.com/posts/Go/Gin_framework/#autoid-0-0-0

#### 3、源码

[Gin源码阅读与分析](https://www.yuque.com/iveryimportantpig/huchao/zd24cb3z2bco5304)：https://www.yuque.com/iveryimportantpig/huchao/zd24cb3z2bco5304

### go-zero

文档：https://go-zero.dev/cn/docs/introduction

go-zero 是一个集成了各种工程实践的 web 和 rpc 框架。通过弹性设计保障了大并发服务端的稳定性，经受了充分的实战检验。

go-zero 包含极简的 API 定义和生成工具 goctl，可以根据定义的 api 文件一键生成 Go, iOS, Android, Kotlin, Dart, TypeScript, JavaScript 代码，并可直接运行。

使用 go-zero 的好处：

- 轻松获得支撑千万日活服务的稳定性
- 内建级联超时控制、限流、自适应熔断、自适应降载等微服务治理能力，无需配置和额外代码
- 微服务治理中间件可无缝集成到其它现有框架使用
- 极简的 API 描述，一键生成各端代码
- 自动校验客户端请求参数合法性
- 大量微服务治理和并发工具包

### 字节-CloudWeGo

文档：https://www.cloudwego.io/zh/docs/

### HTTP-Hertz

文档：https://www.cloudwego.io/zh/docs/hertz/overview/

是一个 Golang 微服务 HTTP 框架，在设计之初参考了其他开源框架 [fasthttp](https://github.com/valyala/fasthttp)、[gin](https://github.com/gin-gonic/gin)、[echo](https://github.com/labstack/echo) 的优势， 并结合字节跳动内部的需求，使其具有高易用性、高性能、高扩展性等特点，目前在字节跳动内部已广泛使用。 如今越来越多的微服务选择使用 Golang，如果对微服务性能有要求，又希望框架能够充分满足内部的可定制化需求，Hertz 会是一个不错的选择。

**特点**

- 高易用性在开发过程中，快速写出来正确的代码往往是更重要的。因此，在 Hertz 在迭代过程中，积极听取用户意见，持续打磨框架，希望为用户提供一个更好的使用体验，帮助用户更快的写出正确的代码。
- 高性能Hertz 默认使用自研的高性能网络库 Netpoll，在一些特殊场景相较于 go net，Hertz 在 QPS、时延上均具有一定优势。关于性能数据，可参考下图 Echo 数据。四个框架的对比:![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1675414683589-8ae9d18c-b2e6-43bd-943f-7392415e0e74.png)三个框架的对比:![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1675414685005-e51955bc-2290-48b8-8782-11f6a26f4efc.png)关于详细的性能数据，可参考 https://github.com/cloudwego/hertz-benchmark。
- 高扩展性Hertz 采用了分层设计，提供了较多的接口以及默认的扩展实现，用户也可以自行扩展。同时得益于框架的分层设计，框架的扩展性也会大很多。目前仅将稳定的能力开源给社区，更多的规划参考 [RoadMap](https://github.com/cloudwego/hertz/blob/main/ROADMAP.md)。
- 多协议支持Hertz 框架原生提供 HTTP1.1、ALPN 协议支持。除此之外，由于分层设计，Hertz 甚至支持自定义构建协议解析逻辑，以满足协议层扩展的任意需求。
- 网络层切换能力Hertz 实现了 Netpoll 和 Golang 原生网络库 间按需切换能力，用户可以针对不同的场景选择合适的网络库，同时也支持以插件的方式为 Hertz 扩展网络库实现。

### RPC-Kitex 

文档：https://www.cloudwego.io/zh/docs/kitex/overview/

字节跳动内部的 Golang 微服务 RPC 框架，具有**高性能**、**强可扩展**的特点，在字节内部已广泛使用。如果对微服务性能有要求，又希望定制扩展融入自己的治理体系，Kitex 会是一个不错的选择。

**框架特点**

- **高性能**使用自研的高性能网络库 [Netpoll](https://github.com/cloudwego/netpoll)，性能相较 go net 具有显著优势。
- **扩展性**提供了较多的扩展接口以及默认扩展实现，使用者也可以根据需要自行定制扩展，具体见下面的框架扩展。
- **多消息协议**RPC 消息协议默认支持 **Thrift**、**Kitex Protobuf**、**gRPC**。Thrift 支持 Buffered 和 Framed 二进制协议；Kitex Protobuf 是 Kitex 自定义的 Protobuf 消息协议，协议格式类似 Thrift；gRPC 是对 gRPC 消息协议的支持，可以与 gRPC 互通。除此之外，使用者也可以扩展自己的消息协议。
- **多传输协议**传输协议封装消息协议进行 RPC 互通，传输协议可以额外透传元信息，用于服务治理，Kitex 支持的传输协议有 **TTHeader**、**HTTP2**。TTHeader 可以和 Thrift、Kitex Protobuf 结合使用；HTTP2 目前主要是结合 gRPC 协议使用，后续也会支持 Thrift。
- **多种消息类型**支持 **PingPong**、**Oneway**、**双向 Streaming**。其中 Oneway 目前只对 Thrift 协议支持，双向 Streaming 只对 gRPC 支持，后续会考虑支持 Thrift 的双向 Streaming。
- **服务治理**支持服务注册/发现、负载均衡、熔断、限流、重试、监控、链路跟踪、日志、诊断等服务治理模块，大部分均已提供默认扩展，使用者可选择集成。
- **代码生成**Kitex 内置代码生成工具，可支持生成 **Thrift**、**Protobuf** 以及脚手架代码。

## 参考并致谢

1、可可酱 [可可酱：Golang常见面试题](https://zhuanlan.zhihu.com/p/434629143)

2、Bel_Ami同学 [golang 面试题(从基础到高级)](https://link.zhihu.com/?target=https%3A//blog.csdn.net/Bel_Ami_n/article/details/123352478)
