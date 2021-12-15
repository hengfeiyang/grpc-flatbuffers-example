# grpc-flatbuffers-example

gRPC默认使用Protocol Buffers编码，同时也支持其他编码如：JSON、FlatBuffers等。

FlatBuffers是一个跨平台的序列化库，旨在实现最大的内存效率。它允许您直接访问序列化数据，而无需首先对其进行解析/解包，同时仍具有良好的向前/向后兼容性。

项目地址：[https://github.com/google/flatbuffers](https://github.com/google/flatbuffers)

FlatBuffers在解编码性能上要比Protocol Buffers快很多，这里有两篇详细介绍Protocol Buffers和FlatBuffers对比的文章：

[https://blog.csdn.net/chosen0ne/article/details/43033575](https://blog.csdn.net/chosen0ne/article/details/43033575)
[https://juzii.gitee.io/2020/03/02/protobuf-vs-flatbuffer/](https://juzii.gitee.io/2020/03/02/protobuf-vs-flatbuffer/)

这里有一篇文章详细介绍了FlatBuffers以及schema的编写:

[https://halfrost.com/flatbuffers_schema/](https://halfrost.com/flatbuffers_schema/)

这里主要来演示一下如何在gRPC中使用FlatBuffers.


## 提示

先安装`flatc` 从下面地址中下载对应版本的flatc即可

[https://github.com/google/flatbuffers/releases/tag/v2.0.0](https://github.com/google/flatbuffers/releases/tag/v2.0.0)

编译生成指令：

```shell
flatc --go --grpc -o api/ api/fbs/greeter.fbs
```

参数说明：

> --go 指定生成的语言是go
> --grpc 指定生成grpc代码
> -o 可选，指定要生成的目标文件目录前缀
> --go-namespace 可选，指定生成的包名，覆盖 fbs 文件中的定义

会在指定目录下生成一个`models`目录，里面即是生成的代码，这个目录名就是`fbs`文件中定义的`namespace`，也可以通过参数`'--go-namespace`来覆盖这个值，以指定新的目录，如：

```shell
flatc --go --go-namespace newmodels --grpc -o api/ api/fbs/greeter.fbs
```

建议通过`fbs`定义`namespace`，这个`namespace`也是Go文件的`package`名称。


参考链接：

[https://github.com/google/flatbuffers/tree/master/grpc/examples/go/greeter](https://github.com/google/flatbuffers/tree/master/grpc/examples/go/greeter)

