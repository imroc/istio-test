# echoserver

## 简介
echoserver 是一个极简的 grpc 服务端，可用于一般 grpc 场景测试

## 端口
* 9000: grpc 接口地址，实现了 chat 服务端
* 9001: http 接口，暴露 prometheus metrics

## 编译镜像
``` bash
./build.sh
```

## ghz 压测
``` bash
./ghz --insecure -d '{"body": "hello"}' --call chat.ChatService/SayHello --concurrency=1 --rps=3000 --total=50000 echoserver:9000
```