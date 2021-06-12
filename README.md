# micro-cli
micro-cli是对外提供接口去中转调用micro grpc的服务，是不需要注册到服务中心去的，因为它有可能是外部服务在调用，当然如果是内部业务服务也需要注册到注册中心去,再nginx做反向代理限流等

## 说明
micro-cli项目使用go-micro v3版本 ，是grpc的客户端,可理解网关服务，服务名称：go.micro.http 监听端口为：8002  
服务调用：客户端[http请求] => micro-cli[rpc请求] => micro  
