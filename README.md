# micro-cli
micro-cli是对外提供接口去中转调用micro grpc的服务，是不需要注册到服务中心去的，因为它有可能是外部服务在调用，当然如果是内部业务服务也需要注册到注册中心去,再nginx做反向代理限流等