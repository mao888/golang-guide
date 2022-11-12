# Protocol Buffers V3中文语法指南[翻译]

本文是官方protocol buffers v3指南的翻译。

本文翻译自https://developers.google.com/protocol-buffers/docs/proto3。

### 定义一个消息类型

首先让我们看一个非常简单的例子。假设你想要定义一个搜索请求消息格式，其中每个搜索请求都包含一个查询词字符串、你感兴趣的查询结果所在的特定页码数和每一页应展示的结果数。

下面是用于定义这个消息类型的 `.proto` 文件。

```protobuf
syntax = "proto3";

message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}
```

- 文件的第一行指定使用 `proto3` 语法: 如果不这样写，protocol buffer编译器将假定你使用 `proto2`。这个声明必须是文件的第一个非空非注释行。
- `SearchRequest` 消息定义指定了三个字段(名称/值对) ，每个字段表示希望包含在此类消息中的每一段数据。每个字段都有一个名称和一个类型。

#### 指定字段类型

在上面的示例中，所有字段都是标量类型（scalar types）: 两个整数(`page_number`和 `result_per_page`)和一个字符串(`query`)。但是也可以为字段指定组合类型，包括枚举和其他消息类型。

#### 分配字段编号

如你所见，消息定义中的每个字段都有一个**唯一的编号**。这些字段编号用来在消息二进制格式中标识字段，在消息类型使用后就不能再更改。注意，范围1到15中的字段编号需要一个字节进行编码，包括字段编号和字段类型。范围16到2047的字段编号采用两个字节。因此，应该为经常使用的消息元素保留数字1到15的编号。切记为将来可能添加的经常使用的元素留出一些编号。

你可以指定的最小字段数是1，最大的字段数是 229−1229−1 ，即536,870,911。你也不能使用19000到19999 (`FieldDescriptor::kFirstReservedNumber` 到`FieldDescriptor::kLastReservedNumber`)的编号，它们是预留给Protocol Buffers协议实现的。如果你在你的`.proto`文件中使用了预留的编号Protocol Buffers编译器就会报错。同样，你也不能使用任何之前保留的字段编号。

#### 指定字段规则

消息字段可以是下列字段之一:

- `singular`: 格式正确的消息可以有这个字段的零个或一个(但不能多于一个)。这是 proto3语法的默认字段规则。
- `repeated`: 该字段可以在格式正确的消息中重复任意次数(包括零次)。重复值的顺序将被保留。

在 proto3中，标量数值类型的`repeated`字段默认使用`packed`编码。

你可以在 [Protocol Buffer Encoding](https://developers.google.com/protocol-buffers/docs/encoding#packed) 中找到关于`packed`编码的更多信息。

#### 添加更多消息类型

可以在一个`.proto` 文件中定义多个消息类型。如果你正在定义多个相关的消息，这是非常有用的——例如，如果想定义与 `SearchRequest` 消息类型对应的应答消息格式`SearchResponse`，你就可以将其添加到同一个`.proto`文件中。

```protobuf
message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}

message SearchResponse {
 ...
}
```

#### 添加注释

要给你的`.proto`文件添加注释，需要使用C/C++风格的`//`和`/* ... */`语法。

```protobuf
/* SearchRequest 表示一个分页查询 
 * 其中有一些字段指示响应中包含哪些结果 */

message SearchRequest {
  string query = 1;
  int32 page_number = 2;  // 页码数
  int32 result_per_page = 3;  // 每页返回的结果数
}
```

#### 保留字段

如果你通过完全删除字段或将其注释掉来**更新**消息类型，那么未来的用户在对该类型进行自己的更新时可以重用字段号。如果其他人以后加载旧版本的相同`.proto`文件，这可能会导致严重的问题，包括数据损坏，隐私漏洞等等。确保这种情况不会发生的一种方法是指定已删除字段的字段编号(和/或名称，这也可能导致 JSON 序列化问题)是保留的（`reserved`）。如果将来有任何用户尝试使用这些字段标识符，protocol buffer编译器将发出提示。

```protobuf
message Foo {
  reserved 2, 15, 9 to 11;
  reserved "foo", "bar";
}
```

注意，不能在同一个`reserved`语句中混合字段名和字段编号。

#### 从你的`.proto`文件生成了什么？

当你使用 protocol buffer 编译器来运行`.proto`文件时，编译器用你选择的语言生成你需要使用文件中描述的消息类型，包括获取和设置字段值，将消息序列化为输出流，以及从输入流解析消息的代码。

- 对**C++**来说，编译器会为每个`.proto`文件生成一个`.h`文件和一个`.cc`文件，`.proto`文件中的每一个消息有一个对应的类。
- 对于 **Java**，编译器生成一个`.java` 文件，每种消息类型都有一个类，还有一个特殊的 `Builder` 类用于创建消息类实例。
- 对于 **Kotlin**，除了 Java 生成的代码之外，编译器还生成一个每种消息类型的 `.kt` 文件，包含一个 DSL，可用于简化消息实例的创建。
- **Python** 稍有不同ー Python 编译器为`.proto`文件中的每个消息类型生成一个带静态描述符的模块，然后与 *metaclass* 一起使用，在运行时创建必要的 Python 数据访问类。
- 对于 **Go**，编译器为文件中的每种消息类型生成一个类型（type）到一个`.pb.go` 文件。
- 对于 **Ruby**，编译器生成一个`.rb` 文件，其中包含一个包含消息类型的 Ruby 模块。
- 对于 **Objective-C**，编译器从每个`.proto`文件生成一个 `pbobjc.h` 和 `pbobjc.m` 文件，`.proto`文件中描述的每种消息类型都有一个类。
- 对于 **C#** ，编译器生从每个`.proto`文件生成一个`.cs` 文件。`.proto`文件中描述的每种消息类型都有一个类。
- 对于 **Dart**，编译器为文件中的每种消息类型生成一个`.pb.dart` 文件。

你可以通过学习所选语言的教程(proto3版本即将推出)了解更多关于使用每种语言的 API 的信息。有关 API 的更多细节，请参阅相关的 [API reference](https://developers.google.com/protocol-buffers/docs/reference/overview)(proto3版本也即将推出)。

### 标量值类型

标量消息字段可以具有以下类型之一——该表显示了`.proto`文件，以及自动生成类中的对应类型（省略了Ruby、C#和Dart）:

| .proto Type |                            Notes                             | C++ Type | Java/Kotlin Type[1] |         Python Type[3]          | Go Type |     PHP Type      |
| :---------: | :----------------------------------------------------------: | :------: | :-----------------: | :-----------------------------: | :-----: | :---------------: |
|   double    |                                                              |  double  |       double        |              float              | float64 |       float       |
|    float    |                                                              |  float   |        float        |              float              | float32 |       float       |
|    int32    | 使用可变长度编码。编码负数效率低下——如果你的字段可能有负值，则使用 sint32代替。 |  int32   |         int         |               int               |  int32  |      integer      |
|    int64    | 使用可变长度编码。编码负数效率低下——如果你的字段可能有负值，则使用 sint64代替。 |  int64   |        long         |           int/long[4]           |  int64  | integer/string[6] |
|   uint32    |                        使用变长编码。                        |  uint32  |       int[2]        |           int/long[4]           | uint32  |      integer      |
|   uint64    |                        使用变长编码。                        |  uint64  |       long[2]       |           int/long[4]           | uint64  | integer/string[6] |
|   sint32    | 使用可变长度编码。带符号的 int 值。这些编码比普通的 int32更有效地编码负数。 |  int32   |         int         |               int               |  int32  |      integer      |
|   sint64    | 使用可变长度编码。带符号的 int 值。这些编码比普通的 int64更有效地编码负数。 |  int64   |        long         |           int/long[4]           |  int64  | integer/string[6] |
|   fixed32   |    总是四个字节。如果值经常大于228，则比 uint32更有效率。    |  uint32  |       int[2]        |           int/long[4]           | uint32  |      integer      |
|   fixed64   |     总是8字节。如果值经常大于256，则比 uint64更有效率。      |  uint64  |  integer/string[6]  |                                 |         |                   |
|  sfixed32   |                        总是四个字节。                        |  int32   |         int         |               int               |  int32  |      integer      |
|  sfixed64   |                        总是八个字节。                        |  int64   |  integer/string[6]  |                                 |         |                   |
|    bool     |                                                              |   bool   |       boolean       |              bool               |  bool   |      boolean      |
|   string    | 字符串必须始终包含 UTF-8编码的或7位 ASCII 文本，且不能长于232。 |  string  |       String        |         str/unicode[5]          | string  |      string       |
|    bytes    |          可以包含任何不超过232字节的任意字节序列。           |  string  |     ByteString      | str (Python 2) bytes (Python 3) | []byte  |      string       |

在使用 [Protocol Buffer Encoding](https://developers.google.com/protocol-buffers/docs/encoding) 对消息进行序列化时，可以了解有关这些类型如何编码的更多信息。

[1] Kotlin 使用来自 Java 的相应类型，甚至是无符号类型，以确保混合 Java/Kotlin 代码库的兼容性。

[2] 在 Java 中，无符号的32位和64位整数使用它们的有符号对应项来表示，最高位存储在有符号位中。

[3] 在任何情况下，为字段设置值都将执行类型检查，以确保其有效。

[4] 64位或无符号的32位整数在解码时总是表示为 long ，但如果在设置字段时给出 int，则可以表示为 int。在任何情况下，值必须与设置时表示的类型相匹配。见[2]。

[5] Python 字符串在解码时表示为 unicode，但如果给出了 ASCII 字符串，则可以表示为 str (这可能会更改)。

[6] 整数用于64位机器，字符串用于32位机器。

### 默认值

当解析消息时，如果编码消息不包含特定的 singular 元素，则解析对象中的相应字段将设置为该字段的默认值。

- 对于字符串，默认值为空字符串。
- 对于字节，默认值为空字节。
- 对于布尔值，默认值为 false。
- 对于数值类型，默认值为零。
- 对于枚举，默认值是第一个定义的枚举值，该值必须为0。
- 对于消息字段，未设置该字段。其确切值与语言有关。详细信息请参阅[生成的代码指南](https://developers.google.com/protocol-buffers/docs/reference/overview)。

repeated 字段的默认值为空(通常是适当语言中的空列表)。

请注意，对于标量消息字段，一旦消息被解析，就无法判断字段是显式设置为默认值(例如，是否一个布尔值是被设置为 `false`)还是根本没有设置: 在定义消息类型时应该牢记这一点。例如，如果你不希望某个行为在默认情况下也发生，那么就不要设置一个布尔值，该布尔值在设置为 `false` 时会开启某些行为。还要注意，如果将标量消息字段**设置**为默认值，则该值将不会在传输过程中序列化。

有关生成的代码的默认工作方式的更多详细信息，请参阅所选语言的[生成代码指南](https://developers.google.com/protocol-buffers/docs/reference/overview)。

### 枚举

在定义消息类型时，你可能希望其中一个字段只能是预定义的值列表中的一个值。例如，假设你想为每个 `SearchRequest` 添加一个语料库字段，其中语料库可以是 `UNIVERSAL`、 `WEB`、 `IMAGES`、 `LOCAL`、 `NEWS`、 `PRODUCTS` 或 `VIDEO`。你可以通过在消息定义中添加一个枚举，为每个可能的值添加一个常量来非常简单地完成这项工作。

在下面的例子中，我们添加了一个名为 `Corpus` 的`enum`，包含所有可能的值，以及一个类型为 `Corpus` 的字段:

```protobuf
message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
  enum Corpus {
    UNIVERSAL = 0;
    WEB = 1;
    IMAGES = 2;
    LOCAL = 3;
    NEWS = 4;
    PRODUCTS = 5;
    VIDEO = 6;
  }
  Corpus corpus = 4;
}
```

如你所见，Corpus enum 的第一个常量映射为零: 每个 enum 定义必须包含一个常量，该常量映射为零作为它的第一个元素。这是因为:

1. 必须有一个零值，这样我们就可以使用0作为数值默认值。
2. 零值必须是第一个元素，以便与 proto2语义兼容，其中第一个枚举值总是默认值。

你可以通过将相同的值分配给不同的枚举常量来定义别名。为此，你需要将 `allow _ alias` 选项设置为 `true`，否则，当发现别名时，protocol 编译器将生成错误消息。

```protobuf
message MyMessage1 {
  enum EnumAllowingAlias {
    option allow_alias = true;
    UNKNOWN = 0;
    STARTED = 1;
    RUNNING = 1;
  }
}
message MyMessage2 {
  enum EnumNotAllowingAlias {
    UNKNOWN = 0;
    STARTED = 1;
    // RUNNING = 1;  // Uncommenting this line will cause a compile error inside Google and a warning message outside.
  }
}
```

枚举的常数必须在32位整数的范围内。由于枚举值在传输时使用[变长编码](https://developers.google.com/protocol-buffers/docs/encoding)，因此负值效率低，因此不推荐使用。可以在消息定义中定义枚举，如上面的例子所示，也可以在外面定义——这样就可以在`.proto`文件中的消息定义中重用这些枚举。你还可以使用`_MessageType_._EnumType_` 语法，使用在一个消息中声明的`enum`类型作为不同消息中的字段类型。

当对一个使用了枚举的`.proto`文件运行 protocol buffer 编译器的时候，对于 Java, Kotlin,或 C++ 生成的代码中将有一个对应的enum，或者对于 Python 会生成一个特殊的`EnumDescriptor`类，它被用于在运行时生成的类中创建一组带有整数值的符号常量。

**注意：**生成的代码可能会受到特定于语言的枚举数限制(单种语言的数量低于千)。请检查你计划使用的语言的限制。

在反序列化过程中，不可识别的枚举值将保留在消息中，尽管当消息被反序列化时，这种值的表示方式依赖于语言。在支持值超出指定符号范围(如 C++ 和 Go)的开放枚举类型的语言中，未知枚举值仅存储为其底层的整数表示形式。在具有闭合枚举类型(如 Java)的语言中，枚举中的一个类型将用于表示一个无法识别的值，并且可以使用特殊的访问器访问底层的整数。在这两种情况下，如果消息被序列化，那么不可识别的值仍然会与消息一起被序列化。

有关如何在应用程序中使用消息`enum`的详细信息，请参阅为所选语言[生成的代码指南](https://developers.google.com/protocol-buffers/docs/reference/overview)。

#### 预留值

如果通过完全删除枚举条目或注释掉枚举类型来更新枚举类型，那么未来的用户在自己更新该类型时可以重用该数值。这可能会导致严重的问题，如果以后有人加载旧版本的相同`.proto`文件，包括数据损坏，隐私漏洞等等。确保不发生这种情况的一种方法是指定已删除条目的数值(和/或名称，这也可能导致 JSON 序列化问题)为 `reserved`。如果任何未来的用户试图使用这些标识符，protocol buffer 编译器将报错。你可以使用 `max`关键字指定保留的数值范围最大为可能的值。

```protobuf
enum Foo {
  reserved 2, 15, 9 to 11, 40 to max;
  reserved "FOO", "BAR";
}
```

注意，不能在同一个保留语句中混合字段名和数值。

### 使用其他消息类型

你可以使用其他消息类型作为字段类型。例如，假设你希望在每个 `SearchResponse`消息中包含UI个 `Result`消息——为了做到这一点，你可以在同一个`.proto`文件中定义 `Result`消息类型。然后在 `SearchResponse`中指定 `Result` 类型的字段。

```protobuf
message SearchResponse {
  repeated Result results = 1;
}

message Result {
  string url = 1;
  string title = 2;
  repeated string snippets = 3;
}
```

#### 导入定义

在上面的示例中，`Result`消息类型定义在与 `SearchResponse`相同的文件中——如果你希望用作字段类型的消息类型已经在另一个`.proto`文件中定义了，该怎么办？

你可以通过 *import* 来使用来自其他 `.proto` 文件的定义。要导入另一个`.proto` 的定义，你需要在文件顶部添加一个 import 语句:

```protobuf
import "myproject/other_protos.proto";
```

默认情况下，只能从直接导入的 `.proto` 文件中使用定义。但是，有时你可能需要将 `.proto` 文件移动到新的位置。你可以在旧目录放一个占位的`.proto`文件使用`import public` 概念将所有导入转发到新位置，而不必直接移动`.proto`文件并修改所有的地方。

**注意，Java 中没有 import public 功能。**

`import public`依赖项可以被任何导入包含`import public`语句的 proto 的代码传递依赖。例如：

```protobuf
// new.proto
// All definitions are moved here
// old.proto
// This is the proto that all clients are importing.
import public "new.proto";
import "other.proto";
// client.proto
import "old.proto";
// You use definitions from old.proto and new.proto, but not other.proto
```

protocol 编译器使用命令行`-I`/`--proto_path`参数指定的一组目录中搜索导入的文件。如果没有给该命令行参数，则查看调用编译器的目录。一般来说，你应该将 `--proto_path` 参数设置为项目的根目录并为所有导入使用正确的名称。

#### 使用proto2消息类型

导入 proto2消息类型并在 proto3消息中使用它们是可能的，反之亦然。然而，proto2 enum 不能直接在 proto3语法中使用(如果一个导入的 proto2消息使用了它们，那没问题)。

### 嵌套类型

你可以在其他消息类型中定义和使用消息类型，如下面的例子——这里的`Result`消息是在 `SearchResponse`消息中定义的:

```protobuf
message SearchResponse {
  message Result {
    string url = 1;
    string title = 2;
    repeated string snippets = 3;
  }
  repeated Result results = 1;
}
```

如果要在其父消息类型之外重用此消息类型，请通过`_Parent_._Type_`使用:

```protobuf
message SomeOtherMessage {
  SearchResponse.Result result = 1;
}
```

你可以随心所欲地将信息一层又一层嵌入其中:

```protobuf
message Outer {                  // Level 0
  message MiddleAA {  // Level 1
    message Inner {   // Level 2
      int64 ival = 1;
      bool  booly = 2;
    }
  }
  message MiddleBB {  // Level 1
    message Inner {   // Level 2
      int32 ival = 1;
      bool  booly = 2;
    }
  }
}
```

### 更新消息类型

如果现有的消息类型不再满足你的所有需要——例如，你希望消息格式有一个额外的字段——但是你仍然希望使用用旧格式创建的代码，不要担心！在不破坏任何现有代码的情况下更新消息类型非常简单，只需记住以下规则:

- 不要更改任何现有字段的字段编号
- 如果添加新字段，那么任何使用“旧”消息格式通过代码序列化的消息仍然可以通过新生成的代码进行解析。你应该记住这些元素的默认值，以便新代码能够正确地与旧代码生成的消息交互。类似地，新代码创建的消息可以通过旧代码解析: 旧的二进制文件在解析时直接忽略新字段。有关详细信息，请参阅 未知字段 部分。
- 字段可以被删除，只要字段编号不再用于你更新的消息类型。你可能希望改为重命名字段，或者为其添加”OBSOLETE_“前缀，或者声明字段编号为`reserved`，以便`.proto`的未来用户不可能不小心重复使用这个编号。
- `int32`、 `uint32`、 `int64`、 `uint64`和 `bool`都是兼容的——这意味着你可以在不破坏向前或向后兼容性的情况下将一个字段从这些类型中的一个更改为另一个。
- 如果一个数字被解析到一个并不适当的类型中，你会得到与在 C++ 中将数字转换为该类型相同的效果(例如，如果一个64位的数字被读作 int32，它将被截断为32位)
- `sint32`和 `sint64`相互兼容，但与其他整数类型不兼容。
- `string`和`bytes`是兼容的，只要字节是有效的 UTF-8。
- 如果字节包含消息的编码版本，则嵌入的消息与`bytes`兼容。
- `fixed32`与 `sfixed32`兼容 `fixed64`与 `sfixed64`兼容。
- 对于`string`、`bytes`和消息字段，`optional`字段与`repeated`字段兼容。给定重复字段的序列化数据作为输入，如果该字段是基本类型字段，期望该字段为可选字段的客户端将接受最后一个输入值; 如果该字段是消息类型字段，则合并所有输入元素。注意，这对于数字类型(包括 bools 和 enums)通常是不安全的。重复的数值类型字段可以按`packed`的格式序列化，如果是`optional`字段，则无法正确解析这些字段。
- Enum 在格式方面与 int32、 uint32、 int64和 uint64兼容(请注意，如果不适合，值将被截断)。但是要注意，当消息被反序列化时，客户端代码可能会区别对待它们: 例如，未被识别的 proto3 `enum`类型将保留在消息中，但是当消息被反序列化时，这种类型的表示方式依赖于语言。Int 字段总是保留它们的值。
- 将单个值更改为**新**的`oneof`成员是安全的，并且二进制兼容。如果确保没有代码一次设置多个字段，那么将多个字段移动到新的`oneof`字段中可能是安全的。将任何字段移动到现有的字段中都是不安全的。

### 未知字段

未知字段是格式良好的协议缓冲区序列化数据，表示解析器不识别的字段。例如，当旧二进制解析由新二进制发送的带有新字段的数据时，这些新字段将成为旧二进制中的未知字段。

最初，proto3消息在解析过程中总是丢弃未知字段，但在3.5版本中，我们重新引入了未知字段的保存来匹配 proto2行为。在3.5及以后的版本中，解析期间保留未知字段，并将其包含在序列化输出中。

## Any

`Any` 消息类型允许你将消息作为嵌入类型使用，而不需要其 `.proto` 定义。`Any`包含一个任意序列化的字节消息，以及一个解析为该消息的类型作为消息的全局唯一标识符的URL。要使用 `Any`类型，需要导入`google/protobuf/any.proto`。

```protobuf
import "google/protobuf/any.proto";

message ErrorStatus {
  string message = 1;
  repeated google.protobuf.Any details = 2;
}
```

给定消息类型的默认类型 URL 是`type.googleapis.com/_packagename_._messagename_`。

不同的语言实现将支持运行库助手以类型安全的方式打包和解包 `Any`值。例如在java中，Any类型会有特殊的`pack()`和`unpack()`访问器，在C++中会有`PackFrom()`和`UnpackTo()`方法。

```protobuf
// Storing an arbitrary message type in Any.
NetworkErrorDetails details = ...;
ErrorStatus status;
status.add_details()->PackFrom(details);

// Reading an arbitrary message from Any.
ErrorStatus status = ...;
for (const Any& detail : status.details()) {
  if (detail.Is<NetworkErrorDetails>()) {
    NetworkErrorDetails network_error;
    detail.UnpackTo(&network_error);
    ... processing network_error ...
  }
}
```

**目前正在开发用于处理任何类型的运行时库。**

如果你已经熟悉 proto2语法，Any 可以保存任意的 proto3消息，类似于 proto2消息，可以允许扩展。

### oneof

如果你有一条包含多个字段的消息，并且最多同时设置其中一个字段，那么你可以通过使用`oneof`来实现并节省内存。

`oneof`字段类似于常规字段，只不过`oneof`中的所有字段共享内存，而且最多可以同时设置一个字段。设置其中的任何成员都会自动清除所有其他成员。根据所选择的语言，可以使用特殊 `case()`或 `WhichOneof()` 方法检查 one of 中的哪个值被设置(如果有的话)。

#### 使用oneof

要定义 oneof 字段需要在你的`.proto`文件中使用`oneof`关键字并在后面跟上名称，在下面的例子中字段名称为`test_oneof`。

```protobuf
message SampleMessage {
  oneof test_oneof {
    string name = 4;
    SubMessage sub_message = 9;
  }
}
```

然后将其中一个字段添加到该字段的定义中。你可以添加任何类型的字段，除了`map`字段和`repeated`字段。

在生成的代码中，其中一个字段具有与常规字段相同的 getter 和 setter。你还可以获得一个特殊的方法来检查其中一个设置了哪个值(如果有的话)。你可以在相关的 [API 参考文献](https://developers.google.com/protocol-buffers/docs/reference/overview)中找到更多关于所选语言的 API。

#### oneof 特性

- 设置一个字段将自动清除该字段的所有其他成员。因此，如果你设置了多个 oneof字段，那么只有最后设置的字段仍然具有值。

```protobuf
  SampleMessage message;
  message.set_name("name");
  CHECK(message.has_name());
  message.mutable_sub_message();   // Will clear name field.
  CHECK(!message.has_name());
```

- 如果解析器在连接中遇到同一个成员的多个成员，则只有最后看到的成员用于解析消息。
- oneof 不支持`repeated`。
- 反射 api 适用于 oneof 字段。
- 如果将 oneof 字段设置为默认值(例如将 int32 oneof 字段设置为0) ，则将设置该字段的“ case”，并在连接上序列化该值。
- 如果你使用 C++ ，确保你的代码不会导致内存崩溃。下面的示例代码将崩溃，因为通过调用 `set_name()`方法已经删除了 `sub_message`。

```protobuf
  SampleMessage message;
  SubMessage* sub_message = message.mutable_sub_message();
  message.set_name("name");      // Will delete sub_message
  sub_message->set_...            // Crashes here
```

- 在C++中，如果你使用`Swap()`两个 oneof 消息，每个消息，两个消息将拥有对方的值，例如在下面的例子中，`msg1`会拥有`sub_message`并且`msg2`会有`name`。

```protobuf
  SampleMessage msg1;
  msg1.set_name("name");
  SampleMessage msg2;
  msg2.mutable_sub_message();
  msg1.swap(&msg2);
  CHECK(msg1.has_sub_message());
  CHECK(msg2.has_name());
```

#### 向后兼容性问题

添加或删除一个字段时要小心。如果检查 one of 的值返回`None`/`NOT_SET`，这可能意味着 one of 没有被设置，或者它已经被设置为 one of 的不同版本中的一个字段。这没有办法区分，因为没有办法知道未知字段是否是 oneof 的成员。

标签重用问题

- 将字段移入或移出 oneof：在序列化和解析消息之后，你可能会丢失一些信息(某些字段将被清除)。但是，你可以安全地将单个字段移动到新的 oneof 字段中，并且如果已知只设置了一个字段，则可以移动多个字段。
- 删除一个oneof 字段再添加回来：这可能会在消息被序列化和解析后清除当前设置的 oneof 字段。
- 拆分或合并oneof：这与移动常规字段有类似的问题。

### Maps

如果你想创建一个关联映射作为你数据定义的一部分，protocol buffers提供了一个方便的快捷语法:

```protobuf
map<key_type, value_type> map_field = N;
```

…其中`key_type`可以是任何整型或字符串类型(因此，除了浮点类型和字节以外的任何标量类型) 。注意，枚举不是有效的`key_type`。`value_type`可以是除另一个映射以外的任何类型。

例如，如果你想创建一个项目映射，其中每个`Project`消息都与一个字符串键相关联，你可以这样定义:

```protobuf
map<string, Project> projects = 3;
```

- 映射字段不能重复。
- 映射值的有线格式排序和映射迭代排序是未定义的，因此不能依赖于映射项的特定排序。
- 当为 `.proto` 生成文本格式时，映射按键排序。数字键按数字排序。
- 当从连接解析或合并时，如果有重复的映射键，则使用最后看到的键。当从文本格式解析映射时，如果有重复的键，解析可能会失败。
- 如果为映射字段提供了键但没有值，则该字段序列化时的行为与语言相关。在 C++ 、 Java、 Kotlin 和 Python 中，类型的默认值是序列化的，而在其他语言中，没有任何值是序列化的。

生成的映射 API 目前可用于所有支持 proto3的语言。你可以在相关的 [API 参考](https://developers.google.com/protocol-buffers/docs/reference/overview)中找到更多关于所选语言的映射 API 的信息。

#### 向后兼容性

map语法序列化后等同于如下内容，因此即使是不支持map语法的protocol buffer实现也是可以处理你的数据的：

```protobuf
message MapFieldEntry {
  key_type key = 1;
  value_type value = 2;
}

repeated MapFieldEntry map_field = N;
```

任何支持映射的protocol buffers实现都必须生成并接受上述定义可以接受的数据。

## Packages

可以向 `.proto` 文件添加一个可选`package`说明符，以防止协议消息类型之间的名称冲突。

```protobuf
package foo.bar;
message Open { ... }
```

然后，你可以在定义消息类型的字段时使用package说明符:

```protobuf
message Foo {
  ...
  foo.bar.Open open = 1;
  ...
}
```

package 说明符影响生成代码的方式取决于你选择的语言:

- 对于**C++**，产生的类会被包装在C++的命名空间中，如上例中的`Open`会被封装在 `foo::bar`空间中；
- 对于**Java**和**Kotlin**，包声明符会变为java的一个包，除非在`.proto`文件中提供了一个明确有`option java_package`；
- 对于 **Python**，这个包声明符是被忽略的，因为Python模块是按照其在文件系统中的位置进行组织的
- 对于**Go**，包可以被用做Go包名称，除非你显式的提供一个`option go_package`在你的`.proto`文件中。
- 对于**Ruby**，生成的类可以被包装在内置的Ruby名称空间中，转换成Ruby所需的大小写样式 （首字母大写；如果第一个符号不是一个字母，则使用PB_前缀），例如Open会在`Foo::Bar`名称空间中。
- 在 **C#** 中，包在转换到 PascalCase 后被用作名称空间，除非你在`.proto`文件中提供`option csharp_namespace`。例如，`Open` 将位于`Foo.Bar`名称空间中。

#### package和名称解析

在 protocol buffer 语言中，类型名称解析的工作原理类似于 C++ : 首先搜索最内层的作用域，然后搜索下一个最内层的作用域，依此类推，每个包都被认为是其父包的“ inner”。前导的“ .”(例如，`.foo.bar.Baz`)表示从最外侧的范围开始。

protocol buffer 通过解析导入的 `.proto` 文件来解析所有类型名称。每种语言的代码生成器都知道如何引用该语言中的每种类型，即使它有不同的作用域规则。

### 定义服务

如果希望将消息类型与 RPC (远程过程调用)系统一起使用，可以在`.proto` 文件和 protocol buffer 编译器将用你选择的语言生成服务接口代码和存根。因此，例如你希望定义一个 RPC 服务，其方法接受你的 `SearchRequest`并返回一个 `SearchResponse`，则可以在`.proto`文件如下定义。

```protobuf
service SearchService {
  rpc Search(SearchRequest) returns (SearchResponse);
}
```

使用 protocol buffers 最直接的 RPC 系统是 gRPC，这是 Google 开发的一个语言和平台中立的开源 RPC 系统，可以与 protocol buffers 一起使用。gRPC 特别适用于protocol buffers ，它可以让你直接从你的`.proto`文件使用特殊的 protocol buffers 编译器插件。

如果你不想使用 gRPC，你也可以在你自己的 RPC 实现中使用协议缓冲。你可以在[《proto2语言指南》](https://developers.google.com/protocol-buffers/docs/proto#services)中找到更多相关信息。

还有一些正在进行的第三方项目正在开发 RPC 的实施协议缓冲。有关我们所知道的项目的链接列表，请参阅[第三方添加项 wiki 页面](https://github.com/protocolbuffers/protobuf/blob/master/docs/third_party.md)。

## JSON 映射

proto3支持 JSON 的规范编码，使得系统之间更容易共享数据。下表按类型逐一描述了编码。

如果 json 编码的数据中缺少某个值，或者该值为 null，那么在解析为 protocol buffer 时，该值将被解释为适当的默认值。如果一个字段在 protocol buffer 中具有默认值，为了节省空间，默认情况下 json 编码的数据中将省略该字段。具体实现可以提供在JSON编码中可选的默认值。

|         proto3         |     JSON      |               JSON example                |                            Notes                             |
| :--------------------: | :-----------: | :---------------------------------------: | :----------------------------------------------------------: |
|        message         |    object     |       `{"fooBar": v, "g": null, …}`       | 生成 JSON 对象。消息字段名映射到 lowerCamelCase 并成为 JSON 对象键。如果指定了 `json_name` 字段选项，则将使用指定的值作为键。解析器接受 lowerCamelCase 名称(或 `json_name` 选项指定的名称)和原始 proto 字段名称。 `null` 是所有字段类型的接受值，并被视为相应字段类型的默认值。 |
|          enum          |    string     |                `"FOO_BAR"`                | 使用 proto 中指定的枚举值的名称。解析器接受枚举名称和整数值。 |
|          map           |    object     |               `{"k": v, …}`               |                    所有键都转换为字符串。                    |
|       repeated V       |     array     |                 `[v, …]`                  |                 `null` 被接受为空列表 `[]`。                 |
|          bool          |  true, false  |               `true, false`               |                                                              |
|         string         |    string     |             `"Hello World!"`              |                                                              |
|         bytes          | base64 string |       `"YWJjMTIzIT8kKiYoKSctPUB+"`        | JSON 值将是使用带填充的标准 base64编码方式编码为字符串的数据。接受带/不带填充的标准或 URL 安全的 base64编码。 |
| int32, fixed32, uint32 |    number     |                `1, -10, 0`                |        JSON 值将是一个十进制数字。接受数字或字符串。         |
| int64, fixed64, uint64 |    string     |               `"1", "-10"`                |       JSON 值将是一个十进制字符串。接受数字或字符串。        |
|     float, double      |    number     |    `1.1, -10.0, 0, "NaN", "Infinity"`     | JSON 值将是一个数字或一个特殊的字符串值“NaN”、“ Infinity”和“-Infinity”。接受数字或字符串。也接受指数表示法。-0被认为等效于0。 |
|          Any           |   `object`    |      `{"@type": "url", "f": v, … }`       | 如果`Any`包含一个具有特殊 JSON 映射的值，它将被转换如下: `{"@type": xxx, "value": yyy}`. 否则，该值将转换为 JSON 对象，并插入`"@type"`字段以指示实际的数据类型。 |
|       Timestamp        |    string     |       `"1972-01-01T10:00:20.021Z"`        | 使用 RFC3339，其中生成的输出将始终是 Z 标准化的，并使用0、3、6或9个小数位。也接受“ Z”以外的偏移量。 |
|        Duration        |    string     |          `"1.000340012s", "1s"`           | 生成的输出总是包含0、3、6或9个小数位，具体取决于所需的精度，后缀“ s”。接受任何小数位(也可以没有) ，只要他们符合纳秒精度和后缀“ s”是必需的。 |
|         Struct         |   `object`    |                  `{ … }`                  |            任何JSON对象。请参见 `struct.proto`。             |
|     Wrapper types      | various types | `2, "2", "foo", true, "true", null, 0, …` | Wrappers 使用与包装原语类型相同的 JSON 表示，只是在数据转换和传输期间允许并保留 `null`。 |
|       FieldMask        |    string     |              `"f.fooBar,h"`               |                  请参见 `field_mask.proto`.                  |
|       ListValue        |     array     |              `[foo, bar, …]`              |                                                              |
|         Value          |     value     |                                           | 任何 JSON 值。请检查 [google.protobuf.Value](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Value) 以获取详细信息。 |
|       NullValue        |     null      |                                           |                          JSON null                           |
|         Empty          |    object     |                   `{}`                    |                       一个空的JSON对象                       |

#### JSON选项

一个proto3协议 JSON 实现可能提供以下选项:

- 提供默认值的字段：在proto3 JSON 输出中，值为默认值的字段被省略。可以提供一个选项，用默认值覆盖此行为和输出字段。
- 忽略位置字段：在缺省情况下，Proto3 JSON 解析器应该拒绝未知字段，但在解析过程中可能会提供一个忽略未知字段的选项。
- 使用 proto 字段名而不是小驼峰名称：默认情况下，proto3 JSON 打印机应该将字段名转换为 lowerCamelCase，并使用它作为 JSON 名称。可以提供一个选项，用原型字段名作为 JSON 名。需要协议3 JSON 解析器同时接受转换后的 lowerCamelCase 名称和原始字段名称。
- 以整数而不是字符串形式展示枚举值：在 JSON 输出中，默认情况下使用枚举值的名称。可以提供一个选项来代替使用枚举值的数值。

剩下options等内容本文略。