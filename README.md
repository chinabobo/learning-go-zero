# learning-go-zero

a demo to learn go-zero with zrpc

## Feature

1. 利用 go-zero 实现一个最小化的 HTTP 服务 以及 简单的api

2. 使用 zrpc 通过 proto文件 生成 rpc 代码

## Useage

启动服务

```shell
$ go run demo.go
```

终端中请求 HTTP 服务接口

```shell
$ curl --request GET 'http://127.0.0.1:8888/from/me'
{"message":"me"}

$ curl --request POST 'http://127.0.0.1:8888/to/herbert'
{"message":"give herbert"}
```

生成rpc代码

```shell
goctl rpc protoc common/proto/train_proxy.proto --style=gozero --home="/home/herbert/workspace/golang/bin/goctl" --go_out=/home/herbert/workspace/learning-go-zero/demo/rpc -
-go-grpc_out=/home/herbert/workspace/learning-go-zero/demo/rpc --zrpc_out=/home/herbert/workspace/learning-go-zero/demo/rpc
```

其中参数

--home goctl 的安装目录

--go_out 生成的文件路径

--go-grpc_out 生成的 grpc 文件路径

--zrpc_out 生成的 zrpc 文件路径



在 /home/herbert/workspace/learning-go-zero/demo/rpc 这个路径下的所有文件都是生成的
