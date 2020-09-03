## Protocol buffers是什么？
Protocol buffers 是一种语言中立，平台无关，可扩展的序列化数据的格式，可用于通信协议，数据存储等。

Protocol buffers 在序列化数据方面，它是灵活的，高效的。相比于 XML 来说，Protocol buffers 更加小巧，更加快速，更加简单。一旦定义了要处理的数据的数据结构之后，就可以利用 Protocol buffers 的代码生成工具生成相关的代码。甚至可以在无需重新部署程序的情况下更新数据结构。只需使用 Protobuf 对数据结构进行一次描述，即可利用各种不同语言或从各种不同数据流中对你的结构化数据轻松读写。

Protocol buffers 很适合做数据存储或 RPC 数据交换格式。可用于通讯协议、数据存储等领域的语言无关、平台无关、可扩展的序列化结构数据格式。

## 为什么要发明Protocol buffers？
大家可能会觉得 Google 发明 protocol buffers 是为了解决序列化速度的，其实真实的原因并不是这样的。

protocol buffers 最先开始是 google 用来解决索引服务器 request/response 协议的。没有 protocol buffers 之前，google 已经存在了一种 request/response 格式，用于手动处理 request/response 的编组和反编组。它也能支持多版本协议，不过代码比较丑陋：

```java
if (version == 3) {
   ...
 } else if (version > 4) {
   if (version == 5) {
     ...
   }
   ...
 }
```
如果非常明确的格式化协议，会使新协议变得非常复杂。因为开发人员必须确保请求发起者与处理请求的实际服务器之间的所有服务器都能理解新协议，然后才能切换开关以开始使用新协议。

这也就是每个服务器开发人员都遇到过的低版本兼容、新旧协议兼容相关的问题。

protocol buffers 为了解决这些问题，于是就诞生了。protocol buffers 被寄予一下 2 个特点：

- 可以很容易地引入新的字段，并且不需要检查数据的中间服务器可以简单地解析并传递数据，而无需了解所有字段。
- 数据格式更加具有自我描述性，可以用各种语言来处理(C++, Java 等各种语言)

这个版本的 protocol buffers 仍需要自己手写解析的代码。

不过随着系统慢慢发展，演进，protocol buffers 目前具有了更多的特性：

- 自动生成的序列化和反序列化代码避免了手动解析的需要。（官方提供自动生成代码工具，各个语言平台的基本都有）
- 除了用于 RPC（远程过程调用）请求之外，人们开始将 protocol buffers 用作持久存储数据的便捷自描述格式（例如，在Bigtable中）。
- 服务器的 RPC 接口可以先声明为协议的一部分，然后用 protocol compiler 生成基类，用户可以使用服务器接口的实际实现来覆盖它们。

protocol buffers 现在是 Google 用于数据的通用语言。在撰写本文时，谷歌代码树中定义了 48162 种不同的消息类型，包括 12183 个 .proto 文件。它们既用于 RPC 系统，也用于在各种存储系统中持久存储数据。

**小结：** protocol buffers 诞生之初是为了解决服务器端新旧协议(高低版本)兼容性问题，名字也很体贴，“协议缓冲区”。只不过后期慢慢发展成用于传输数据。

## 如何使用？todo

### 数据格式

### 各种语言提供的API

### 一些规范

## Protocol Buffer 编码原理

###  Varints 编码

### Message Structure 编码

protocol buffer 中 message 是一系列键值对。message 的二进制版本只是使用字段号(field's number 和 wire_type)作为 key。每个字段的名称和声明类型只能在解码端通过引用消息类型的定义（即 .proto 文件）来确定。这一点也是人们常常说的 protocol buffer 比 JSON，XML 安全一点的原因，如果没有数据结构描述 .proto 文件，拿到数据以后是无法解释成正常的数据的。

由于采用了 tag-value 的形式，所以 option 的 field 如果有，就存在在这个 message buffer 中，如果没有，就不会在这里，这一点也算是压缩了 message 的大小了。

当消息编码时，键和值被连接成一个字节流。当消息被解码时，解析器需要能够跳过它无法识别的字段。这样，可以将新字段添加到消息中，而不会破坏不知道它们的旧程序。这就是所谓的 “向后”兼容性。

为此，线性的格式消息中每对的“key”实际上是两个值，其中一个是来自.proto文件的字段编号，加上提供正好足够的信息来查找下一个值的长度。在大多数语言实现中，这个 key 被称为 tag。

Tag - Length - Value

