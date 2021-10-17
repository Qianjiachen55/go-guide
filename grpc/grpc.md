# grpc

### demo

xxx.proto

```protobuf
指定 远程调用方法
指定 请求 & 返回消息格式

```



xxx.server

```
- server端实现远程调用方法

- 
	1. 监听端口
	2. 创建grpc服务器
	3. 注册grpc服务
	4. 给定的grpc服务器上注册服务器反射服务
	

```



client

```shell
1. 连接服务器
2. 创建连接客户端
3. 实例化message对象
4. 远程调用
```

